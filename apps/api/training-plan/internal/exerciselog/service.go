package exerciselog

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

func (s *Service) LogExercise(ctx context.Context, exerciseId, userId string, reps []int, weight []float64, notes *string) error {
	slog.InfoContext(ctx, "logging exercise", slog.String("exercise_id", exerciseId), slog.String("user_id", userId))

	id, err := utils.GenerateUUIDV7(ctx)
	if err != nil {
		return err
	}

	now := time.Now().UTC()
	l := &domain.ExerciseLog{
		Id:         *id,
		UserId:     userId,
		ExerciseId: exerciseId,
		Reps:       reps,
		Weight:     weight,
		Notes:      notes,
		CreatedAt:  now,
		UpdatedAt:  now,
	}
	if err := s.repo.LogExercise(ctx, l); err != nil {
		slog.ErrorContext(ctx, "failed to log exercise", slog.String("exercise_id", exerciseId), slog.Any("error", err))
		return err
	}
	return nil
}
