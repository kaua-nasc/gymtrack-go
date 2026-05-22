package metrics

import (
	"context"
	"log/slog"
	"time"

	"github.com/kaua-nasc/gymtrack-go/apps/api/identity/internal/domain"
	"github.com/kaua-nasc/gymtrack-go/libs/utils"
)

type Service struct {
	repo Repository
}

func NewService(repo Repository) *Service {
	return &Service{repo: repo}
}

func (s *Service) AddBodyMeasurementNote(ctx context.Context, id, note string) error {
	return s.repo.AddBodyMeasurementNote(ctx, id, note)
}

func (s *Service) FindLastBodyMeasurementNote(ctx context.Context, userId string) (*domain.BodyMeasurement, error) {
	return s.repo.FindLastBodyMeasurementNote(ctx, userId)
}

func (s *Service) ListBodyMeasurements(ctx context.Context, userId, cursor string, limit int) ([]*domain.BodyMeasurement, string, error) {
	var decodedCursor *utils.CursorData
	utils.DecodeCursor(cursor, &decodedCursor)

	measurements, rawNextCursor, err := s.repo.ListBodyMeasurements(ctx, userId, decodedCursor, limit)
	if err != nil {
		return nil, "", err
	}

	nextCursorStr, _ := utils.EncodeCursor(rawNextCursor)
	return measurements, nextCursorStr, nil
}

func (s *Service) AddWeightLogNote(ctx context.Context, id, note string) error {
	return s.repo.AddWeightLogNote(ctx, id, note)
}

func (s *Service) AddGoalMetric(ctx context.Context, goal *domain.MetricGoal) error {
	now := time.Now().UTC()
	newId, err := utils.GenerateUUIDV7(ctx)
	if err != nil {
		return err
	}

	goal.ID = *newId
	goal.CreatedAt = now
	goal.UpdatedAt = now
	goal.Status = domain.MetricGoalActive

	return s.repo.AddGoalMetric(ctx, *goal)
}

func (s *Service) ListGoalsMetric(ctx context.Context, userId, cursor string, limit int) ([]*domain.MetricGoal, string, error) {
	var decodedCursor *utils.CursorData
	utils.DecodeCursor(cursor, &decodedCursor)

	goals, rawNextCursor, err := s.repo.ListGoalsMetric(ctx, userId, decodedCursor, limit)
	if err != nil {
		return nil, "", err
	}

	nextCursorStr, _ := utils.EncodeCursor(rawNextCursor)
	return goals, nextCursorStr, nil
}

func (s *Service) ListWeightLogs(ctx context.Context, userId, cursor string, limit int) ([]*domain.WeightLog, string, error) {
	slog.InfoContext(ctx, "listing weight logs", slog.String("user_id", userId), slog.Int("limit", limit))

	var decodedCursor *utils.CursorData
	if err := utils.DecodeCursor(cursor, &decodedCursor); err != nil {
		slog.WarnContext(ctx, "failed to decode cursor for weight history", slog.String("cursor", cursor), slog.Any("error", err))
	}

	logs, rawNextCursor, err := s.repo.ListWeightLogs(ctx, userId, decodedCursor, limit)
	if err != nil {
		slog.ErrorContext(ctx, "failed to list weight history", slog.String("user_id", userId), slog.Any("error", err))
		return nil, "", err
	}

	nextCursorStr, _ := utils.EncodeCursor(rawNextCursor)
	return logs, nextCursorStr, nil
}
