package internal

import (
	"context"
	"encoding/base64"
	"encoding/json"
	"errors"
	"fmt"
	"log/slog"
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"
	"github.com/kaua-nasc/gymtrack-go/libs/auth"
	"golang.org/x/sync/errgroup"
)

type TrainingPlanService struct {
	repo     *TrainingPlanRepository
	identity *IdentityService
	storage  StorageService
	validate *validator.Validate
}

func NewTrainingPlanService(
	repo *TrainingPlanRepository,
	identity *IdentityService,
	storage StorageService,
) *TrainingPlanService {
	return &TrainingPlanService{
		repo:     repo,
		identity: identity,
		storage:  storage,
		validate: validator.New(),
	}
}

func (s *TrainingPlanService) CreatePlan(ctx context.Context, plan TrainingPlan, user auth.AuthUser) (*TrainingPlan, error) {
	slog.InfoContext(ctx, "creating training plan", slog.String("author_id", user.ID))

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
	if user.Type == "client" {
		count, err := s.repo.CountByAuthor(ctx, user.ID)
		if err != nil {
			slog.ErrorContext(ctx, "failed to count plans by author", slog.String("author_id", user.ID), slog.Any("error", err))
			return nil, err
		}
		if count >= 1 {
			slog.WarnContext(ctx, "client attempted to create more than one plan", slog.String("author_id", user.ID))
			return nil, errors.New("clients can only have one personal training plan")
		}
		plan.Visibility = Private
	} else if user.Type != "trainer" && user.Type != "admin" {
		slog.WarnContext(ctx, "unauthorized user type attempted to create plan", slog.String("user_id", user.ID), slog.String("user_type", user.Type))
		return nil, errors.New("only trainers and admins can create training plans")
	}

	now := time.Now().UTC()
	plan.AuthorId = user.ID
	plan.CreatedAt = now
	plan.UpdatedAt = now

	// 3. Save the plan
	if err := s.repo.Create(ctx, &plan); err != nil {
		slog.ErrorContext(ctx, "failed to save training plan", slog.Any("error", err))
		return nil, err
	}

	// 4. Save nested days and exercises (Cascade)
	for _, day := range plan.Days {
		day.TrainingPlanId = plan.Id
		day.CreatedAt = now
		day.UpdatedAt = now
		if err := s.repo.CreateDay(ctx, &day); err != nil {
			slog.ErrorContext(ctx, "failed to save plan day", slog.String("plan_id", plan.Id), slog.Any("error", err))
			return nil, fmt.Errorf("could not save plan day: %w", err)
		}

		for _, exercise := range day.Exercises {
			exercise.DayId = day.Id
			exercise.CreatedAt = now
			exercise.UpdatedAt = now
			if err := s.repo.CreateExercise(ctx, &exercise); err != nil {
				slog.ErrorContext(ctx, "failed to save exercise", slog.String("day_id", day.Id), slog.Any("error", err))
				return nil, fmt.Errorf("could not save exercise: %w", err)
			}
		}
	}

	slog.InfoContext(ctx, "training plan created successfully", slog.String("plan_id", plan.Id))
	return &plan, nil
}

func (s *TrainingPlanService) CreateDay(ctx context.Context, day Day) error {
	id, err := uuid.NewV7()
	if err != nil {
		slog.ErrorContext(ctx, "failed to generate uuid for day", slog.Any("error", err))
		return fmt.Errorf("error on generate uuid")
	}

	now := time.Now().UTC()
	day.Id = id.String()
	day.CreatedAt = now
	day.UpdatedAt = now
	if err := s.repo.CreateDay(ctx, &day); err != nil {
		slog.ErrorContext(ctx, "failed to save plan day", slog.String("day_id", day.Id), slog.Any("error", err))
		return fmt.Errorf("could not save plan day: %w", err)
	}

	return nil
}

func (s *TrainingPlanService) CreateExercise(ctx context.Context, e Exercise) error {
	id, err := uuid.NewV7()
	if err != nil {
		slog.ErrorContext(ctx, "failed to generate uuid for day", slog.Any("error", err))
		return fmt.Errorf("error on generate uuid")
	}

	now := time.Now().UTC()
	e.Id = id.String()
	e.CreatedAt = now
	e.UpdatedAt = now
	if err := s.repo.CreateExercise(ctx, &e); err != nil {
		slog.ErrorContext(ctx, "failed to save plan day", slog.Any("error", err))
		return fmt.Errorf("could not save plan day: %w", err)
	}

	return nil
}

func (s *TrainingPlanService) CreateDays(ctx context.Context, days []Day) error {
	now := time.Now().UTC()
	for i := range days {
		id, err := uuid.NewV7()
		if err != nil {
			slog.ErrorContext(ctx, "failed to generate uuid for day", slog.Any("error", err))
			return fmt.Errorf("error on generate uuid")
		}

		days[i].Id = id.String()
		days[i].CreatedAt = now
		days[i].UpdatedAt = now
	}

	if err := s.repo.CreateDays(ctx, days); err != nil {
		slog.ErrorContext(ctx, "failed to save plan day", slog.Any("error", err))
		return fmt.Errorf("could not save plan day: %w", err)
	}

	return nil
}

func (s *TrainingPlanService) DeleteDay(ctx context.Context, id string) error {
	if err := s.repo.DeleteDay(ctx, id); err != nil {
		slog.ErrorContext(ctx, "failed to save plan day", slog.Any("error", err))
		return fmt.Errorf("could not save plan day: %w", err)
	}

	return nil
}

func (s *TrainingPlanService) DeleteExercise(ctx context.Context, id string) error {
	if err := s.repo.DeleteExercise(ctx, id); err != nil {
		slog.ErrorContext(ctx, "failed to save plan day", slog.Any("error", err))
		return fmt.Errorf("could not save plan day: %w", err)
	}

	return nil
}

func (s *TrainingPlanService) UpdatePlan(ctx context.Context, id string, data TrainingPlan) (*TrainingPlan, error) {
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

func (s *TrainingPlanService) DeletePlan(ctx context.Context, id string) error {
	if err := s.repo.DeletePlan(ctx, id); err != nil {
		slog.ErrorContext(ctx, "failed to save plan day", slog.Any("error", err))
		return fmt.Errorf("could not save plan day: %w", err)
	}

	return nil
}

func (s *TrainingPlanService) ExistsPlan(ctx context.Context, id string) (bool, error) {
	plan, err := s.repo.Find(ctx, id)
	if err != nil {
		slog.ErrorContext(ctx, "failed to search for training plan", slog.String("plan_id", id), slog.Any("error", err))
		return false, fmt.Errorf("error searching for training plan")
	}
	return plan != nil, nil
}

func (s *TrainingPlanService) GetPlan(ctx context.Context, id string) (*TrainingPlan, error) {
	slog.InfoContext(ctx, "fetching training plan", slog.String("plan_id", id))

	plan, err := s.repo.Find(ctx, id)
	if err != nil {
		slog.ErrorContext(ctx, "failed to search for training plan", slog.String("plan_id", id), slog.Any("error", err))
		return nil, fmt.Errorf("error searching for training plan")
	}
	if plan == nil {
		slog.WarnContext(ctx, "training plan not found", slog.String("plan_id", id))
		return nil, nil
	}

	token, _ := ctx.Value(string(auth.TokenContextKey)).(string)
	authorIds := []string{plan.AuthorId}
	authorsMap, err := s.identity.ListUser(ctx, &authorIds, token)
	if err == nil {
		plan.Author = authorsMap[plan.AuthorId]
	} else {
		slog.WarnContext(ctx, "failed to fetch author details", slog.String("author_id", plan.AuthorId), slog.Any("error", err))
	}

	// Hydrate with days and exercises
	days, err := s.repo.ListDaysByPlan(ctx, id)
	if err == nil {
		for i := range days {
			exercises, err := s.repo.ListExercisesByDay(ctx, days[i].Id)
			if err == nil {
				days[i].Exercises = make([]Exercise, len(exercises))
				for j, e := range exercises {
					days[i].Exercises[j] = *e
				}
			} else {
				slog.WarnContext(ctx, "failed to list exercises for day", slog.String("day_id", days[i].Id), slog.Any("error", err))
			}
		}
		plan.Days = make([]Day, len(days))
		for i, d := range days {
			plan.Days[i] = *d
		}
	} else {
		slog.WarnContext(ctx, "failed to list days for plan", slog.String("plan_id", id), slog.Any("error", err))
	}

	return plan, nil
}

func (s *TrainingPlanService) ListPlan(ctx context.Context, authorId, cursor string, limit int) ([]*TrainingPlan, string, error) {
	slog.InfoContext(ctx, "listing training plans", slog.String("author_id", authorId), slog.Int("limit", limit))

	var decodedCursor *CursorData
	if cursor != "" {
		b, err := base64.StdEncoding.DecodeString(cursor)
		if err == nil {
			json.Unmarshal(b, &decodedCursor)
		} else {
			slog.WarnContext(ctx, "failed to decode cursor", slog.String("cursor", cursor), slog.Any("error", err))
		}
	}

	plans, rawNextCursor, err := s.repo.List(ctx, authorId, decodedCursor, limit)
	if err != nil {
		slog.ErrorContext(ctx, "failed to list training plans", slog.Any("error", err))
		return nil, "", fmt.Errorf("error listing training plans")
	}

	if len(plans) == 0 {
		return []*TrainingPlan{}, "", nil
	}

	planIDs := make([]string, len(plans))
	authorIDMap := make(map[string]bool)
	for i, p := range plans {
		planIDs[i] = p.Id
		authorIDMap[p.AuthorId] = true
	}
	var authorIDs []string
	for id := range authorIDMap {
		authorIDs = append(authorIDs, id)
	}

	g, ctx := errgroup.WithContext(ctx)

	var likesMap map[string]int
	var userLikesSet map[string]bool
	var authorsMap map[string]*any

	token, _ := ctx.Value(string(auth.TokenContextKey)).(string)

	g.Go(func() error {
		var err error
		likesMap, err = s.repo.LikesCount(ctx, &planIDs)
		if err != nil {
			slog.WarnContext(ctx, "failed to fetch likes count", slog.Any("error", err))
		}
		return err
	})

	g.Go(func() error {
		if authorId == "" {
			return nil
		}
		var err error
		userLikesSet, err = s.repo.LikesByUser(ctx, &planIDs, &authorId)
		if err != nil {
			slog.WarnContext(ctx, "failed to fetch user likes", slog.String("user_id", authorId), slog.Any("error", err))
		}
		return err
	})

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
		p.LikesCount = likesMap[p.Id]
		p.LikedByCurrentUser = userLikesSet[p.Id]
	}

	// Encode cursor
	var nextCursorStr string
	if rawNextCursor != nil {
		b, _ := json.Marshal(rawNextCursor)
		nextCursorStr = base64.StdEncoding.EncodeToString(b)
	}

	return plans, nextCursorStr, nil
}

func (s *TrainingPlanService) LikePlan(ctx context.Context, planId string, userId string) error {
	slog.InfoContext(ctx, "liking training plan", slog.String("plan_id", planId), slog.String("user_id", userId))

	now := time.Now().UTC()
	like := &TrainingPlanLike{
		Id:             fmt.Sprintf("%s:%s", planId, userId),
		LikedBy:        userId,
		TrainingPlanId: planId,
		CreatedAt:      now,
		UpdatedAt:      now,
	}
	if err := s.repo.LikePlan(ctx, like); err != nil {
		slog.ErrorContext(ctx, "failed to like plan", slog.String("plan_id", planId), slog.String("user_id", userId), slog.Any("error", err))
		return err
	}
	return nil
}

func (s *TrainingPlanService) UnlikePlan(ctx context.Context, planId string, userId string) error {
	slog.InfoContext(ctx, "unliking training plan", slog.String("plan_id", planId), slog.String("user_id", userId))

	if err := s.repo.UnlikePlan(ctx, planId, userId); err != nil {
		slog.ErrorContext(ctx, "failed to unlike plan", slog.String("plan_id", planId), slog.String("user_id", userId), slog.Any("error", err))
		return err
	}
	return nil
}

func (s *TrainingPlanService) AddPlanComment(ctx context.Context, planId, content, userId string) (*TrainingPlanComment, error) {
	slog.InfoContext(ctx, "adding comment to plan", slog.String("plan_id", planId), slog.String("user_id", userId))

	id, err := uuid.NewV7()
	if err != nil {
		slog.ErrorContext(ctx, "failed to generate uuid for comment", slog.Any("error", err))
		return nil, fmt.Errorf("error on generate uuid")
	}
	now := time.Now().UTC()
	comment := &TrainingPlanComment{
		Id:             id.String(),
		Content:        content,
		AuthorId:       userId,
		TrainingPlanId: planId,
		CreatedAt:      now,
		UpdatedAt:      now,
	}
	if err := s.repo.AddPlanComment(ctx, comment); err != nil {
		slog.ErrorContext(ctx, "failed to add comment to plan", slog.String("plan_id", planId), slog.Any("error", err))
		return nil, err
	}
	return comment, nil
}

func (s *TrainingPlanService) RemovePlanComment(ctx context.Context, commentId, userId string) error {
	slog.InfoContext(ctx, "removing comment from plan", slog.String("comment_id", commentId), slog.String("user_id", userId))

	comment, err := s.repo.FindComment(ctx, commentId)
	if err != nil {
		slog.ErrorContext(ctx, "failed to find comment", slog.String("comment_id", commentId), slog.Any("error", err))
		return err
	}
	if comment == nil {
		slog.WarnContext(ctx, "comment not found for removal", slog.String("comment_id", commentId))
		return errors.New("comment not found")
	}

	// Check if the user is the author of the comment
	if comment.AuthorId != userId {
		// Optional: Check if the user is the author of the plan
		plan, err := s.repo.Find(ctx, comment.TrainingPlanId)
		if err != nil || plan == nil || plan.AuthorId != userId {
			slog.WarnContext(ctx, "unauthorized attempt to remove comment", slog.String("comment_id", commentId), slog.String("user_id", userId))
			return errors.New("you are not authorized to remove this comment")
		}
	}

	if err := s.repo.RemovePlanComment(ctx, commentId); err != nil {
		slog.ErrorContext(ctx, "failed to remove comment", slog.String("comment_id", commentId), slog.Any("error", err))
		return err
	}
	return nil
}

func (s *TrainingPlanService) ListPlanComments(ctx context.Context, planId, cursor string, limit int) ([]*TrainingPlanComment, string, error) {
	slog.InfoContext(ctx, "listing plan comments", slog.String("plan_id", planId), slog.Int("limit", limit))

	var decodedCursor *CursorData
	if cursor != "" {
		b, err := base64.StdEncoding.DecodeString(cursor)
		if err == nil {
			json.Unmarshal(b, &decodedCursor)
		} else {
			slog.WarnContext(ctx, "failed to decode cursor for comments", slog.String("cursor", cursor), slog.Any("error", err))
		}
	}

	comments, rawNextCursor, err := s.repo.ListPlanComments(ctx, planId, decodedCursor, limit)
	if err != nil {
		slog.ErrorContext(ctx, "failed to list plan comments", slog.String("plan_id", planId), slog.Any("error", err))
		return nil, "", err
	}

	if len(comments) > 0 {
		authorIDsMap := make(map[string]bool)
		for _, c := range comments {
			authorIDsMap[c.AuthorId] = true
		}
		var authorIDs []string
		for id := range authorIDsMap {
			authorIDs = append(authorIDs, id)
		}

		token, _ := ctx.Value(string(auth.TokenContextKey)).(string)
		authorsMap, err := s.identity.ListUser(ctx, &authorIDs, token)
		if err == nil {
			for _, c := range comments {
				c.Author = authorsMap[c.AuthorId]
			}
		} else {
			slog.WarnContext(ctx, "failed to fetch authors for comments", slog.Any("error", err))
		}
	}

	// Encode cursor
	var nextCursorStr string
	if rawNextCursor != nil {
		b, _ := json.Marshal(rawNextCursor)
		nextCursorStr = base64.StdEncoding.EncodeToString(b)
	}

	return comments, nextCursorStr, nil
}

func (s *TrainingPlanService) authorizeAccess(ctx context.Context, plan *TrainingPlan) error {
	user, ok := ctx.Value(string(auth.UserContextKey)).(auth.AuthUser)
	if !ok {
		return errors.New("unauthorized")
	}

	if plan.AuthorId == user.ID {
		return nil
	}

	return errors.New("you are not authorized to modify this training plan")
}

// Subscription Logic

func (s *TrainingPlanService) ListSubscription(ctx context.Context, userId string) ([]*PlanSubscription, error) {
	slog.InfoContext(ctx, "listing subscriptions", slog.String("user_id", userId))

	subscriptions, err := s.repo.ListSubscription(ctx, userId)
	if err != nil {
		slog.ErrorContext(ctx, "failed to list subscriptions", slog.String("user_id", userId), slog.Any("error", err))
		return nil, err
	}

	return subscriptions, nil
}

func (s *TrainingPlanService) Subscribe(ctx context.Context, planId, userId string, subType PlanSubscriptionType) error {
	slog.InfoContext(ctx, "subscribing user to plan", slog.String("plan_id", planId), slog.String("user_id", userId))

	existing, err := s.repo.FindSubscription(ctx, planId, userId)
	if err != nil {
		slog.ErrorContext(ctx, "failed to check existing subscription", slog.String("plan_id", planId), slog.String("user_id", userId), slog.Any("error", err))
		return err
	}
	if existing != nil {
		slog.WarnContext(ctx, "user already subscribed to plan", slog.String("plan_id", planId), slog.String("user_id", userId))
		return errors.New("already subscribed")
	}

	now := time.Now().UTC()
	sub := &PlanSubscription{
		Id:             fmt.Sprintf("%s:%s", planId, userId),
		TrainingPlanId: planId,
		UserId:         userId,
		Status:         InProgress,
		Type:           subType,
		CreatedAt:      now,
		UpdatedAt:      now,
	}

	if err := s.repo.CreatePlanSubscription(ctx, sub); err != nil {
		slog.ErrorContext(ctx, "failed to create plan subscription", slog.String("plan_id", planId), slog.String("user_id", userId), slog.Any("error", err))
		return err
	}
	return nil
}

// Unsubscription Logic
func (s *TrainingPlanService) Unsubscribe(ctx context.Context, planId, userId string) error {
	slog.InfoContext(ctx, "unsubscribing user from plan", slog.String("plan_id", planId), slog.String("user_id", userId))

	existing, err := s.repo.FindSubscription(ctx, planId, userId)
	if err != nil {
		slog.ErrorContext(ctx, "failed to find subscription for unsubscription", slog.String("plan_id", planId), slog.String("user_id", userId), slog.Any("error", err))
		return err
	}
	if existing == nil {
		slog.WarnContext(ctx, "subscription not found for unsubscription", slog.String("plan_id", planId), slog.String("user_id", userId))
		return errors.New("subscription not found")
	}

	if err := s.repo.DeletePlanSubscription(ctx, existing); err != nil {
		slog.ErrorContext(ctx, "failed to delete plan subscription", slog.String("plan_id", planId), slog.String("user_id", userId), slog.Any("error", err))
		return err
	}
	return nil
}

func (s *TrainingPlanService) ChangeSubscriptionStatus(ctx context.Context, planId, userId string, status PlanSubscriptionStatus) error {
	subscription, err := s.repo.FindSubscription(ctx, planId, userId)
	if err != nil {
		slog.ErrorContext(ctx, "failed to find subscription for unsubscription", slog.String("plan_id", planId), slog.String("user_id", userId), slog.Any("error", err))
		return err
	}

	switch status {
	case NotStarted:
		if subscription.Status != Canceled {
			return fmt.Errorf("o status deve ser cancelado")
		}
	case InProgress:
		if subscription.Status != NotStarted {
			return fmt.Errorf("o status deve ser cancelado")
		}
	case Completed:
		if subscription.Status != InProgress {
			return fmt.Errorf("o status deve ser cancelado")
		}
	case Canceled:
		if subscription.Status != InProgress {
			return fmt.Errorf("o status deve ser cancelado")
		}
	}

	return s.repo.UpdateSubscriptionStatus(ctx, *subscription, status)
}

func (s *TrainingPlanService) ChangeSubscriptionPrivacy(ctx context.Context, planId, userId string, subsType PlanSubscriptionType) error {
	subscription, err := s.repo.FindSubscription(ctx, planId, userId)
	if err != nil {
		slog.ErrorContext(ctx, "failed to find subscription for unsubscription", slog.String("plan_id", planId), slog.String("user_id", userId), slog.Any("error", err))
		return err
	}

	return s.repo.UpdateSubscriptionPrivacy(ctx, *subscription, subsType)
}

func (s *TrainingPlanService) CompleteDay(ctx context.Context, planId, userId, dayId string) error {
	slog.InfoContext(ctx, "completing plan day", slog.String("plan_id", planId), slog.String("user_id", userId), slog.String("day_id", dayId))

	sub, err := s.repo.FindSubscription(ctx, planId, userId)
	if err != nil {
		slog.ErrorContext(ctx, "failed to find subscription for day completion", slog.String("plan_id", planId), slog.String("user_id", userId), slog.Any("error", err))
		return err
	}
	if sub == nil {
		slog.WarnContext(ctx, "subscription not found for day completion", slog.String("plan_id", planId), slog.String("user_id", userId))
		return errors.New("subscription not found")
	}

	now := time.Now().UTC()
	progress := &PlanDayProgress{
		Id:                 fmt.Sprintf("%s:%s:%s", planId, userId, dayId),
		DayId:              dayId,
		PlanSubscriptionId: sub.Id,
		Status:             DayCompleted,
		CreatedAt:          now,
		UpdatedAt:          now,
	}

	if err := s.repo.CreateSubscriptionProgress(ctx, progress); err != nil {
		slog.ErrorContext(ctx, "failed to create subscription progress", slog.Any("error", err))
		return err
	}
	return nil
}

func (s *TrainingPlanService) AddFeedback(ctx context.Context, planId, userId string, rating float64, message *string) error {
	slog.InfoContext(ctx, "adding feedback to plan", slog.String("plan_id", planId), slog.String("user_id", userId), slog.Float64("rating", rating))

	now := time.Now().UTC()
	feedback := &TrainingPlanFeedback{
		Id:             fmt.Sprintf("%s:%s", planId, userId),
		TrainingPlanId: planId,
		UserId:         userId,
		Rating:         rating,
		Message:        message,
		CreatedAt:      now,
		UpdatedAt:      now,
	}
	if err := s.repo.AddFeedback(ctx, feedback); err != nil {
		slog.ErrorContext(ctx, "failed to add feedback", slog.String("plan_id", planId), slog.Any("error", err))
		return err
	}
	return nil
}

func (s *TrainingPlanService) LogExercise(ctx context.Context, exerciseId, userId string, reps []int, weight []float64, notes *string) error {
	slog.InfoContext(ctx, "logging exercise", slog.String("exercise_id", exerciseId), slog.String("user_id", userId))

	id, err := uuid.NewV7()
	if err != nil {
		slog.ErrorContext(ctx, "failed to generate uuid for day", slog.Any("error", err))
		return fmt.Errorf("error on generate uuid")
	}

	now := time.Now().UTC()
	log := &ExerciseLog{
		Id:         id.String(),
		UserId:     userId,
		ExerciseId: exerciseId,
		Reps:       reps,
		Weight:     weight,
		Notes:      notes,
		CreatedAt:  now,
		UpdatedAt:  now,
	}
	if err := s.repo.LogExercise(ctx, log); err != nil {
		slog.ErrorContext(ctx, "failed to log exercise", slog.String("exercise_id", exerciseId), slog.Any("error", err))
		return err
	}
	return nil
}
