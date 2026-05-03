package internal

import (
	"context"
	"encoding/base64"
	"encoding/json"
	"errors"
	"fmt"
	"time"

	"github.com/go-playground/validator/v10"
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
	if err := s.validate.Struct(plan); err != nil {
		return nil, err
	}

	// 1. Check if user exists in identity service
	exists, err := s.identity.ExistsUser(ctx, user.ID)
	if err != nil || !exists {
		return nil, errors.New("user not found")
	}

	// 2. Business Rule: Clients can only have one personal plan, and it must be private
	if user.Type == "client" {
		count, err := s.repo.CountByAuthor(ctx, user.ID)
		if err != nil {
			return nil, err
		}
		if count >= 1 {
			return nil, errors.New("clients can only have one personal training plan")
		}
		plan.Visibility = Private
	} else if user.Type != "trainer" && user.Type != "admin" {
		return nil, errors.New("only trainers and admins can create training plans")
	}

	plan.AuthorId = user.ID
	plan.CreatedAt = time.Now()
	plan.UpdatedAt = time.Now()

	// 3. Save the plan
	if err := s.repo.Create(ctx, &plan); err != nil {
		return nil, err
	}

	// 4. Save nested days and exercises (Cascade)
	for _, day := range plan.Days {
		day.TrainingPlanId = plan.Id
		day.CreatedAt = time.Now()
		day.UpdatedAt = time.Now()
		if err := s.repo.CreateDay(ctx, &day); err != nil {
			return nil, fmt.Errorf("could not save plan day: %w", err)
		}

		for _, exercise := range day.Exercises {
			exercise.DayId = day.Id
			exercise.CreatedAt = time.Now()
			exercise.UpdatedAt = time.Now()
			if err := s.repo.CreateExercise(ctx, &exercise); err != nil {
				return nil, fmt.Errorf("could not save exercise: %w", err)
			}
		}
	}

	return &plan, nil
}

func (s *TrainingPlanService) UpdatePlan(ctx context.Context, id string, data TrainingPlan) (*TrainingPlan, error) {
	plan, err := s.repo.Find(ctx, id)
	if err != nil {
		return nil, fmt.Errorf("error finding training plan: %w", err)
	}

	if plan == nil {
		return nil, errors.New("training plan not found")
	}

	if err := s.authorizeAccess(ctx, plan); err != nil {
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
	plan.UpdatedAt = time.Now()

	if err := s.repo.Update(ctx, plan); err != nil {
		return nil, err
	}

	return plan, nil
}

func (s *TrainingPlanService) ExistsPlan(ctx context.Context, id string) (bool, error) {
	plan, err := s.repo.Find(ctx, id)
	if err != nil {
		return false, fmt.Errorf("error searching for training plan")
	}
	return plan != nil, nil
}

func (s *TrainingPlanService) GetPlan(ctx context.Context, id string) (*TrainingPlan, error) {
	plan, err := s.repo.Find(ctx, id)
	if err != nil {
		return nil, fmt.Errorf("error searching for training plan")
	}
	if plan == nil {
		return nil, nil
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
			}
		}
		plan.Days = make([]Day, len(days))
		for i, d := range days {
			plan.Days[i] = *d
		}
	}

	return plan, nil
}

func (s *TrainingPlanService) ListPlan(ctx context.Context, authorId, cursor string, limit int) ([]*TrainingPlan, string, error) {
	var decodedCursor *CursorData
	if cursor != "" {
		b, err := base64.StdEncoding.DecodeString(cursor)
		if err == nil {
			json.Unmarshal(b, &decodedCursor)
		}
	}

	plans, rawNextCursor, err := s.repo.List(ctx, authorId, decodedCursor, limit)
	if err != nil {
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

	g.Go(func() error {
		var err error
		likesMap, err = s.repo.LikesCount(ctx, &planIDs)
		return err
	})

	g.Go(func() error {
		if authorId == "" {
			return nil
		}
		var err error
		userLikesSet, err = s.repo.LikesByUser(ctx, &planIDs, &authorId)
		return err
	})

	g.Go(func() error {
		var err error
		authorsMap, err = s.identity.ListUser(ctx, &authorIDs)
		return err
	})

	if err := g.Wait(); err != nil {
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
	like := &TrainingPlanLike{
		Id:             fmt.Sprintf("%s:%s", planId, userId),
		LikedBy:        userId,
		TrainingPlanId: planId,
		CreatedAt:      time.Now(),
		UpdatedAt:      time.Now(),
	}
	return s.repo.LikePlan(ctx, like)
}

func (s *TrainingPlanService) UnlikePlan(ctx context.Context, planId string, userId string) error {
	return s.repo.UnlikePlan(ctx, planId, userId)
}

func (s *TrainingPlanService) AddPlanComment(ctx context.Context, planId, content, userId string) (*TrainingPlanComment, error) {
	comment := &TrainingPlanComment{
		Id:             fmt.Sprintf("%s:%s:%d", planId, userId, time.Now().UnixNano()),
		Content:        content,
		AuthorId:       userId,
		TrainingPlanId: planId,
		CreatedAt:      time.Now(),
		UpdatedAt:      time.Now(),
	}
	if err := s.repo.AddPlanComment(ctx, comment); err != nil {
		return nil, err
	}
	return comment, nil
}

func (s *TrainingPlanService) RemovePlanComment(ctx context.Context, commentId, userId string) error {
	comment, err := s.repo.FindComment(ctx, commentId)
	if err != nil {
		return err
	}
	if comment == nil {
		return errors.New("comment not found")
	}

	// Check if the user is the author of the comment
	if comment.AuthorId != userId {
		// Optional: Check if the user is the author of the plan
		plan, err := s.repo.Find(ctx, comment.TrainingPlanId)
		if err != nil || plan == nil || plan.AuthorId != userId {
			return errors.New("you are not authorized to remove this comment")
		}
	}

	return s.repo.RemovePlanComment(ctx, commentId)
}

func (s *TrainingPlanService) ListPlanComments(ctx context.Context, planId, cursor string, limit int) ([]*TrainingPlanComment, string, error) {
	var decodedCursor *CursorData
	if cursor != "" {
		b, err := base64.StdEncoding.DecodeString(cursor)
		if err == nil {
			json.Unmarshal(b, &decodedCursor)
		}
	}

	comments, rawNextCursor, err := s.repo.ListPlanComments(ctx, planId, decodedCursor, limit)
	if err != nil {
		return nil, "", err
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
func (s *TrainingPlanService) Subscribe(ctx context.Context, planId, userId string, subType PlanSubscriptionType) error {
	existing, err := s.repo.FindSubscription(ctx, planId, userId)
	if err != nil {
		return err
	}
	if existing != nil {
		return errors.New("already subscribed")
	}

	sub := &PlanSubscription{
		Id:             fmt.Sprintf("%s:%s", planId, userId),
		TrainingPlanId: planId,
		UserId:         userId,
		Status:         InProgress,
		Type:           subType,
		CreatedAt:      time.Now(),
		UpdatedAt:      time.Now(),
	}

	return s.repo.CreatePlanSubscription(ctx, sub)
}

func (s *TrainingPlanService) CompleteDay(ctx context.Context, planId, userId, dayId string) error {
	sub, err := s.repo.FindSubscription(ctx, planId, userId)
	if err != nil {
		return err
	}
	if sub == nil {
		return errors.New("subscription not found")
	}

	progress := &PlanDayProgress{
		Id:                 fmt.Sprintf("%s:%s:%s", planId, userId, dayId),
		DayId:              dayId,
		PlanSubscriptionId: sub.Id,
		Status:             DayCompleted,
		CreatedAt:          time.Now(),
		UpdatedAt:          time.Now(),
	}

	return s.repo.CreateSubscriptionProgress(ctx, progress)
}

func (s *TrainingPlanService) AddFeedback(ctx context.Context, planId, userId string, rating float64, message *string) error {
	feedback := &TrainingPlanFeedback{
		Id:             fmt.Sprintf("%s:%s", planId, userId),
		TrainingPlanId: planId,
		UserId:         userId,
		Rating:         rating,
		Message:        message,
		CreatedAt:      time.Now(),
		UpdatedAt:      time.Now(),
	}
	return s.repo.AddFeedback(ctx, feedback)
}

func (s *TrainingPlanService) LogExercise(ctx context.Context, exerciseId, userId string, reps []int, weight []float64, notes *string) error {
	log := &ExerciseLog{
		Id:         fmt.Sprintf("%s:%s:%d", exerciseId, userId, time.Now().UnixNano()),
		UserId:     userId,
		ExerciseId: exerciseId,
		Reps:       reps,
		Weight:     weight,
		Notes:      notes,
		CreatedAt:  time.Now(),
		UpdatedAt:  time.Now(),
	}
	return s.repo.LogExercise(ctx, log)
}
