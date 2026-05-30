package feedback

import (
	"context"
	"log/slog"
	"time"

	"github.com/kaua-nasc/gymtrack-go/apps/api/training-plan/internal/domain"
	"github.com/kaua-nasc/gymtrack-go/apps/api/training-plan/internal/plan"
	"github.com/kaua-nasc/gymtrack-go/apps/api/training-plan/internal/subscription"
	"github.com/kaua-nasc/gymtrack-go/libs/utils"
)

type Service struct {
	repo     Repository
	subRepo  subscription.Repository
	planRepo plan.Repository
}

func NewService(repo Repository, subRepo subscription.Repository, planRepo plan.Repository) *Service {
	return &Service{
		repo:     repo,
		subRepo:  subRepo,
		planRepo: planRepo,
	}
}

func (s *Service) AddFeedback(ctx context.Context, planId, userId string, rating float64, message *string) error {
	slog.InfoContext(ctx, "adding feedback to plan", slog.String("plan_id", planId), slog.String("user_id", userId), slog.Float64("rating", rating))

	exists, err := s.subRepo.HasSubscription(ctx, planId, userId)
	if err != nil {
		return err
	}

	if !exists {
		return domain.ErrFeedbackNotAllowed
	}

	id, err := utils.GenerateUUIDV7(ctx)
	if err != nil {
		return err
	}
	now := time.Now().UTC()
	f := &domain.TrainingPlanFeedback{
		Id:             *id,
		TrainingPlanId: planId,
		UserId:         userId,
		Rating:         rating,
		Message:        message,
		CreatedAt:      now,
		UpdatedAt:      now,
	}
	if err := s.repo.AddFeedback(ctx, f); err != nil {
		slog.ErrorContext(ctx, "failed to add feedback", slog.String("plan_id", planId), slog.Any("error", err))
		return err
	}

	return s.updatePlanRatingStats(ctx, planId)
}

func (s *Service) ListFeedback(ctx context.Context, planId string, cursor string, limit int) ([]domain.TrainingPlanFeedback, string, error) {
	slog.InfoContext(ctx, "listing feedback for plan", slog.String("plan_id", planId))

	var decodedCursor *utils.CursorData
	utils.DecodeCursor(cursor, &decodedCursor)

	feedbacks, rawNextCursor, err := s.repo.ListFeedback(ctx, planId, decodedCursor, limit)
	if err != nil {
		slog.ErrorContext(ctx, "failed to list feedback", slog.String("plan_id", planId), slog.Any("error", err))
		return nil, "", err
	}

	nextCursor, _ := utils.EncodeCursor(rawNextCursor)
	return feedbacks, nextCursor, nil
}

func (s *Service) DeleteFeedback(ctx context.Context, feedbackId, userId string) error {
	slog.InfoContext(ctx, "deleting feedback", slog.String("feedback_id", feedbackId), slog.String("user_id", userId))

	f, err := s.repo.FindByID(ctx, feedbackId)
	if err != nil {
		return err
	}

	if f == nil {
		return domain.ErrFeedbackNotFound
	}

	if f.UserId != userId {
		return domain.ErrNotFeedbackAuthor
	}

	if err := s.repo.Delete(ctx, feedbackId); err != nil {
		return err
	}

	return s.updatePlanRatingStats(ctx, f.TrainingPlanId)
}

func (s *Service) updatePlanRatingStats(ctx context.Context, planId string) error {
	sum, count, err := s.repo.GetRatingStats(ctx, planId)
	if err != nil {
		slog.ErrorContext(ctx, "failed to get rating stats", slog.String("plan_id", planId), slog.Any("error", err))
		return err
	}

	if err := s.planRepo.UpdateRatingStats(ctx, planId, sum, count); err != nil {
		slog.ErrorContext(ctx, "failed to update plan rating stats", slog.String("plan_id", planId), slog.Any("error", err))
		return err
	}

	return nil
}
