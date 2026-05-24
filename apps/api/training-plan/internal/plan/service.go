package plan

import (
	"context"
	"errors"
	"fmt"
	"log/slog"
	"time"

	"path/filepath"

	"github.com/go-playground/validator/v10"
	"github.com/kaua-nasc/gymtrack-go/apps/api/training-plan/internal/domain"
	"github.com/kaua-nasc/gymtrack-go/libs/auth"
	"github.com/kaua-nasc/gymtrack-go/libs/storage"
	"github.com/kaua-nasc/gymtrack-go/libs/utils"
	"golang.org/x/sync/errgroup"
)

type Service struct {
	repo     Repository
	identity domain.IdentityClient
	validate *validator.Validate
}

func NewService(
	repo Repository,
	identity domain.IdentityClient,
) *Service {
	return &Service{
		repo:     repo,
		identity: identity,
		validate: validator.New(),
	}
}

func (s *Service) CreatePlan(ctx context.Context, plan domain.TrainingPlan, user auth.AuthUser) (*domain.TrainingPlan, error) {
	slog.InfoContext(ctx, "creating training plan", slog.String("authorId", user.ID))

	if err := s.validate.Struct(plan); err != nil {
		slog.ErrorContext(ctx, "failed to validate training plan", slog.Any("error", err))
		return nil, err
	}

	token, _ := ctx.Value(string(auth.TokenContextKey)).(string)

	// 1. Check if user exists in identity service
	exists, err := s.identity.ExistsUser(ctx, user.ID, token)
	if err != nil || !exists {
		slog.ErrorContext(ctx, "user not found in identity service", slog.String("user_id", user.ID), slog.Any("error", err))
		return nil, errors.New("user not found")
	}

	// 2. Business Rule: Clients can only have one personal plan, and it must be private
	count, err := s.repo.CountByAuthor(ctx, user.ID)
	if err != nil {
		slog.ErrorContext(ctx, "failed to count plans by author", slog.String("authorId", user.ID), slog.Any("error", err))
		return nil, err
	}
	if count >= 1 {
		slog.WarnContext(ctx, "client attempted to create more than one plan", slog.String("authorId", user.ID))
		return nil, errors.New("clients can only have one personal training plan")
	}
	plan.Visibility = domain.Private

	id, err := utils.GenerateUUIDV7(ctx)
	if err != nil {
		return nil, err
	}
	now := time.Now().UTC()
	plan.Id = id
	plan.AuthorId = user.ID
	plan.CreatedAt = now
	plan.UpdatedAt = now

	// 3. Save the plan
	if err := s.repo.Create(ctx, &plan); err != nil {
		slog.ErrorContext(ctx, "failed to save training plan", slog.Any("error", err))
		return nil, err
	}

	// 4. Save nested days and exercises (Cascade)
	for dayIdx := range plan.Days {
		plan.Days[dayIdx].TrainingPlanId = *plan.Id
		plan.Days[dayIdx].CreatedAt = now
		plan.Days[dayIdx].UpdatedAt = now
		plan.Days[dayIdx].Sequence = dayIdx
		if err := s.repo.CreateDay(ctx, &plan.Days[dayIdx]); err != nil {
			slog.ErrorContext(ctx, "failed to save plan day", slog.String("plan_id", *plan.Id), slog.Any("error", err))
			return nil, fmt.Errorf("could not save plan day: %w", err)
		}

		for exerciseIdx := range plan.Days[dayIdx].Exercises {
			plan.Days[dayIdx].Exercises[exerciseIdx].DayId = plan.Days[dayIdx].Id
			plan.Days[dayIdx].Exercises[exerciseIdx].CreatedAt = now
			plan.Days[dayIdx].Exercises[exerciseIdx].UpdatedAt = now
			plan.Days[dayIdx].Exercises[exerciseIdx].Sequence = exerciseIdx
			if err := s.repo.CreateExercise(ctx, &plan.Days[dayIdx].Exercises[exerciseIdx]); err != nil {
				slog.ErrorContext(ctx, "failed to save exercise", slog.String("day_id", plan.Days[dayIdx].Id), slog.Any("error", err))
				return nil, fmt.Errorf("could not save exercise: %w", err)
			}
		}
	}

	slog.InfoContext(ctx, "training plan created successfully", slog.String("plan_id", *plan.Id))
	return &plan, nil
}

func (s *Service) CreateDay(ctx context.Context, day *domain.Day) error {
	id, err := utils.GenerateUUIDV7(ctx)
	if err != nil {
		return err
	}

	now := time.Now().UTC()
	day.Id = *id
	day.CreatedAt = now
	day.UpdatedAt = now
	if err := s.repo.CreateDay(ctx, day); err != nil {
		slog.ErrorContext(ctx, "failed to save plan day", slog.String("day_id", day.Id), slog.Any("error", err))
		return fmt.Errorf("could not save plan day: %w", err)
	}

	return nil
}

func (s *Service) CreateExercise(ctx context.Context, e *domain.Exercise) error {
	id, err := utils.GenerateUUIDV7(ctx)
	if err != nil {
		return err
	}

	now := time.Now().UTC()
	e.Id = *id
	e.CreatedAt = now
	e.UpdatedAt = now
	if err := s.repo.CreateExercise(ctx, e); err != nil {
		slog.ErrorContext(ctx, "failed to save plan day", slog.Any("error", err))
		return fmt.Errorf("could not save plan day: %w", err)
	}

	return nil
}

func (s *Service) UploadExerciseMedia(ctx context.Context, exerciseId string, videoFile *domain.UploadFile, imageFile *domain.UploadFile) (*string, *string, error) {
	var videoUrl *string
	var imageUrl *string

	if videoFile != nil {
		ext := filepath.Ext(videoFile.Filename)
		videoPath := fmt.Sprintf("exercises/%s/video%s", exerciseId, ext)
		if err := storage.UploadBuffer(ctx, videoPath, videoFile.Data); err != nil {
			return nil, nil, fmt.Errorf("failed to upload video: %w", err)
		}
		url := storage.GetBlobURL(videoPath)
		videoUrl = url
	}

	if imageFile != nil {
		ext := filepath.Ext(imageFile.Filename)
		imagePath := fmt.Sprintf("exercises/%s/image%s", exerciseId, ext)
		if err := storage.UploadBuffer(ctx, imagePath, imageFile.Data); err != nil {
			return nil, nil, fmt.Errorf("failed to upload image: %w", err)
		}
		url := storage.GetBlobURL(imagePath)
		imageUrl = url
	}

	if err := s.repo.UpdateExerciseMedia(ctx, exerciseId, videoUrl, imageUrl); err != nil {
		return nil, nil, fmt.Errorf("failed to update exercise media in repository: %w", err)
	}

	return videoUrl, imageUrl, nil
}

func (s *Service) DeleteDay(ctx context.Context, id string) error {
	if err := s.repo.DeleteDay(ctx, id); err != nil {
		slog.ErrorContext(ctx, "failed to save plan day", slog.Any("error", err))
		return fmt.Errorf("could not save plan day: %w", err)
	}

	return nil
}

func (s *Service) DeleteExercise(ctx context.Context, id string) error {
	if err := s.repo.DeleteExercise(ctx, id); err != nil {
		slog.ErrorContext(ctx, "failed to save plan day", slog.Any("error", err))
		return fmt.Errorf("could not save plan day: %w", err)
	}

	return nil
}

func (s *Service) UpdatePlan(ctx context.Context, id string, data domain.TrainingPlan) (*domain.TrainingPlan, error) {
	slog.InfoContext(ctx, "updating training plan", slog.String("plan_id", id))

	plan, err := s.repo.Find(ctx, id)
	if err != nil {
		slog.ErrorContext(ctx, "failed to find training plan", slog.String("plan_id", id), slog.Any("error", err))
		return nil, fmt.Errorf("error finding training plan: %w", err)
	}

	if plan == nil {
		slog.WarnContext(ctx, "training plan not found for update", slog.String("plan_id", id))
		return nil, errors.New("training plan not found")
	}

	if err := s.authorizeAccess(ctx, plan); err != nil {
		slog.WarnContext(ctx, "unauthorized attempt to update plan", slog.String("plan_id", id))
		return nil, err
	}

	// Update fields
	plan.Name = data.Name
	plan.TimeInDays = data.TimeInDays
	plan.Type = data.Type
	plan.Level = data.Level
	plan.Visibility = data.Visibility
	plan.Observation = data.Observation
	plan.Pathology = data.Pathology
	plan.Description = data.Description
	plan.UpdatedAt = time.Now().UTC()

	if err := s.repo.Update(ctx, plan); err != nil {
		slog.ErrorContext(ctx, "failed to update training plan", slog.String("plan_id", id), slog.Any("error", err))
		return nil, err
	}

	slog.InfoContext(ctx, "training plan updated successfully", slog.String("plan_id", id))
	return plan, nil
}

func (s *Service) DeletePlan(ctx context.Context, id string) error {
	if err := s.repo.DeletePlan(ctx, id); err != nil {
		slog.ErrorContext(ctx, "failed to save plan day", slog.Any("error", err))
		return fmt.Errorf("could not save plan day: %w", err)
	}

	return nil
}

func (s *Service) ExistsPlan(ctx context.Context, id string) (bool, error) {
	plan, err := s.repo.Find(ctx, id)
	if err != nil {
		slog.ErrorContext(ctx, "failed to search for training plan", slog.String("plan_id", id), slog.Any("error", err))
		return false, fmt.Errorf("error searching for training plan")
	}
	return plan != nil, nil
}

func (s *Service) GetPlan(ctx context.Context, id string) (*domain.TrainingPlan, error) {
	slog.InfoContext(ctx, "fetching training plan complete", slog.String("plan_id", id))

	plan, err := s.repo.FindComplete(ctx, id)
	if err != nil {
		slog.ErrorContext(ctx, "failed to search for training plan", slog.String("plan_id", id), slog.Any("error", err))
		return nil, fmt.Errorf("error searching for training plan")
	}
	if plan == nil {
		slog.WarnContext(ctx, "training plan not found", slog.String("plan_id", id))
		return nil, nil
	}

	user, ok := ctx.Value(string(auth.UserContextKey)).(auth.AuthUser)
	if ok {
		sub, err := s.repo.FindSubscriptionByPlan(ctx, id, user.ID)
		if err == nil && sub != nil {
			plan.PlanSubscriptionStatus = &sub.Status
		}
	}

	token, _ := ctx.Value(string(auth.TokenContextKey)).(string)
	author, err := s.identity.FindUser(ctx, plan.AuthorId, token)
	if err != nil {
		slog.WarnContext(ctx, "failed to fetch author details", slog.String("authorId", plan.AuthorId), slog.Any("error", err))
	}

	plan.Author = author

	return plan, nil
}

func (s *Service) ListPlansByIds(ctx context.Context, ids []string) ([]*domain.TrainingPlan, error) {
	slog.InfoContext(ctx, "listing training plans by ids", slog.Any("ids", ids))

	plans, err := s.repo.ListByIds(ctx, ids)
	if err != nil {
		slog.ErrorContext(ctx, "failed to list training plans by ids", slog.Any("error", err))
		return nil, fmt.Errorf("error listing training plans by ids")
	}

	return plans, nil
}

func (s *Service) ListPlan(ctx context.Context, authorId, cursor string, limit int) ([]*domain.TrainingPlan, string, error) {
	slog.InfoContext(ctx, "listing training plans", slog.String("authorId", authorId), slog.Int("limit", limit))

	var decodedCursor *utils.CursorData
	if err := utils.DecodeCursor(cursor, &decodedCursor); err != nil {
		slog.WarnContext(ctx, "failed to decode cursor", slog.String("cursor", cursor), slog.Any("error", err))
	}

	plans, rawNextCursor, err := s.repo.List(ctx, authorId, decodedCursor, limit)
	if err != nil {
		slog.ErrorContext(ctx, "failed to list training plans", slog.Any("error", err))
		return nil, "", fmt.Errorf("error listing training plans")
	}

	if len(plans) == 0 {
		return []*domain.TrainingPlan{}, "", nil
	}

	planIDs := make([]string, len(plans))
	authorIDMap := make(map[string]bool)
	for i, p := range plans {
		planIDs[i] = *p.Id
		authorIDMap[p.AuthorId] = true
	}
	var authorIDs []string
	for id := range authorIDMap {
		authorIDs = append(authorIDs, id)
	}

	g, ctx := errgroup.WithContext(ctx)

	var authorsMap map[string]*any

	token, _ := ctx.Value(string(auth.TokenContextKey)).(string)

	g.Go(func() error {
		var err error
		authorsMap, err = s.identity.ListUser(ctx, &authorIDs, token)
		if err != nil {
			slog.WarnContext(ctx, "failed to fetch authors", slog.Any("error", err))
		}
		return err
	})

	if err := g.Wait(); err != nil {
		slog.ErrorContext(ctx, "error during parallel hydration of plans", slog.Any("error", err))
		return nil, "", err
	}

	// Hydration
	for _, p := range plans {
		p.Author = authorsMap[p.AuthorId]
	}

	nextCursorStr, _ := utils.EncodeCursor(rawNextCursor)
	return plans, nextCursorStr, nil
}

func (s *Service) authorizeAccess(ctx context.Context, plan *domain.TrainingPlan) error {
	user, ok := ctx.Value(string(auth.UserContextKey)).(auth.AuthUser)
	if !ok {
		return errors.New("unauthorized")
	}

	if plan.AuthorId == user.ID {
		return nil
	}

	return errors.New("you are not authorized to modify this training plan")
}
