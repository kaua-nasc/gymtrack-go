package metrics

import (
	"context"
	"log/slog"
	"time"

	"github.com/kaua-nasc/gymtrack-go/apps/api/identity/internal/domain"
	"github.com/kaua-nasc/gymtrack-go/apps/api/identity/internal/trainer"
	"github.com/kaua-nasc/gymtrack-go/apps/api/identity/internal/user"
	"github.com/kaua-nasc/gymtrack-go/libs/utils"
)

type Service struct {
	repo        Repository
	userRepo    user.Repository
	trainerRepo trainer.Repository
}

func NewService(repo Repository, userRepo user.Repository, trainerRepo trainer.Repository) *Service {
	return &Service{
		repo:        repo,
		userRepo:    userRepo,
		trainerRepo: trainerRepo,
	}
}

func (s *Service) checkPrivacy(ctx context.Context, requesterId, userId string, checkFn func(*domain.UserPrivacySettings) bool) (*time.Time, error) {
	if requesterId == userId {
		return nil, nil
	}

	linkedAt, err := s.trainerRepo.GetTrainerLinkDate(ctx, requesterId, userId)
	if err != nil {
		return nil, err
	}
	if linkedAt == nil {
		return nil, domain.ErrUnauthorizedTrainerAccess
	}

	privacy, err := s.userRepo.GetPrivacySettings(ctx, userId)
	if err != nil {
		return nil, err
	}

	if privacy != nil && !checkFn(privacy) {
		return nil, domain.ErrPrivacySettingsForbidden
	}

	return linkedAt, nil
}

func (s *Service) CreateBodyMeasurement(ctx context.Context, requesterId string, m *domain.BodyMeasurement) error {
	if requesterId != m.UserId {
		return domain.ErrUnauthorizedAccess
	}

	id, err := utils.GenerateUUIDV7String(ctx)
	if err != nil {
		return err
	}

	m.ID = id
	m.CreatedAt = time.Now().UTC()
	m.UpdatedAt = time.Now().UTC()

	return s.repo.CreateBodyMeasurement(ctx, m)
}

func (s *Service) AddBodyMeasurementNote(ctx context.Context, requesterId, id, note string) error {
	m, err := s.repo.FindBodyMeasurement(ctx, id)
	if err != nil {
		return err
	}
	if m == nil {
		return domain.ErrUserNotFound
	}

	if _, err := s.checkPrivacy(ctx, requesterId, m.UserId, func(p *domain.UserPrivacySettings) bool {
		return p.AllowTrainerNotes
	}); err != nil {
		return err
	}

	return s.repo.AddBodyMeasurementNote(ctx, id, note)
}

func (s *Service) FindLastBodyMeasurement(ctx context.Context, requesterId, userId string) (*domain.BodyMeasurement, error) {
	if requesterId != userId {
		return nil, domain.ErrUnauthorizedAccess
	}

	return s.repo.FindLastBodyMeasurement(ctx, userId)
}

func (s *Service) ListBodyMeasurements(ctx context.Context, requesterId, userId, cursor string, limit int) ([]*domain.BodyMeasurement, string, error) {
	linkedAt, err := s.checkPrivacy(ctx, requesterId, userId, func(p *domain.UserPrivacySettings) bool {
		return p.ShareBodyMeasurements
	})
	if err != nil {
		return nil, "", err
	}

	var decodedCursor *utils.CursorData
	utils.DecodeCursor(cursor, &decodedCursor)

	var since *time.Time
	if linkedAt != nil {
		privacy, _ := s.userRepo.GetPrivacySettings(ctx, userId)
		if privacy != nil && !privacy.SharePastDataWithTrainer {
			since = linkedAt
		}
	}

	measurements, rawNextCursor, err := s.repo.ListBodyMeasurements(ctx, userId, since, decodedCursor, limit)
	if err != nil {
		return nil, "", err
	}

	nextCursorStr, _ := utils.EncodeCursor(rawNextCursor)
	return measurements, nextCursorStr, nil
}

func (s *Service) CreateWeightLog(ctx context.Context, requesterId string, l *domain.WeightLog) error {
	if requesterId != l.UserId {
		return domain.ErrUnauthorizedAccess
	}

	id, err := utils.GenerateUUIDV7String(ctx)
	if err != nil {
		return err
	}

	l.ID = id
	l.CreatedAt = time.Now().UTC()
	l.UpdatedAt = time.Now().UTC()

	return s.repo.CreateWeightLog(ctx, l)
}

func (s *Service) AddWeightLogNote(ctx context.Context, requesterId, id, note string) error {
	l, err := s.repo.FindWeightLog(ctx, id)
	if err != nil {
		return err
	}
	if l == nil {
		return domain.ErrUserNotFound
	}

	if _, err := s.checkPrivacy(ctx, requesterId, l.UserId, func(p *domain.UserPrivacySettings) bool {
		return p.AllowTrainerNotes
	}); err != nil {
		return err
	}

	return s.repo.AddWeightLogNote(ctx, id, note)
}

func (s *Service) FindLastWeightLog(ctx context.Context, requesterId, userId string) (*domain.WeightLog, error) {
	if requesterId != userId {
		return nil, domain.ErrUnauthorizedAccess
	}

	return s.repo.FindLastWeightLog(ctx, userId)
}

func (s *Service) ListWeightLogs(ctx context.Context, requesterId, userId, cursor string, limit int) ([]*domain.WeightLog, string, error) {
	slog.InfoContext(ctx, "listing weight logs", slog.String("user_id", userId), slog.Int("limit", limit))

	linkedAt, err := s.checkPrivacy(ctx, requesterId, userId, func(p *domain.UserPrivacySettings) bool {
		return p.ShareWeightLogs
	})
	if err != nil {
		return nil, "", err
	}

	var decodedCursor *utils.CursorData
	if err := utils.DecodeCursor(cursor, &decodedCursor); err != nil {
		slog.WarnContext(ctx, "failed to decode cursor for weight history", slog.String("cursor", cursor), slog.Any("error", err))
	}

	var since *time.Time
	if linkedAt != nil {
		privacy, _ := s.userRepo.GetPrivacySettings(ctx, userId)
		if privacy != nil && !privacy.SharePastDataWithTrainer {
			since = linkedAt
		}
	}

	logs, rawNextCursor, err := s.repo.ListWeightLogs(ctx, userId, since, decodedCursor, limit)
	if err != nil {
		slog.ErrorContext(ctx, "failed to list weight history", slog.String("user_id", userId), slog.Any("error", err))
		return nil, "", err
	}

	nextCursorStr, _ := utils.EncodeCursor(rawNextCursor)
	return logs, nextCursorStr, nil
}
