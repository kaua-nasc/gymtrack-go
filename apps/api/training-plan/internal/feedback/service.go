package feedback

import (
	"context"
	"log/slog"
	"time"

	"github.com/kaua-nasc/gymtrack-go/apps/api/training-plan/internal/domain"
	"github.com/kaua-nasc/gymtrack-go/libs/utils"
)

type Service struct {
	repo Repository
}

func NewService(repo Repository) *Service {
	return &Service{
		repo: repo,
	}
}

func (s *Service) AddFeedback(ctx context.Context, planId, userId string, rating float64, message *string) error {
	slog.InfoContext(ctx, "adding feedback to plan", slog.String("plan_id", planId), slog.String("user_id", userId), slog.Float64("rating", rating))

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
	return nil
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
