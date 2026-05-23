package trainer

import (
	"context"
	"time"

	"github.com/kaua-nasc/gymtrack-go/apps/api/identity/internal/domain"
	"github.com/kaua-nasc/gymtrack-go/apps/api/identity/internal/user"
	"github.com/kaua-nasc/gymtrack-go/libs/utils"
)

type Service struct {
	repo     Repository
	userRepo user.Repository
}

func NewService(repo Repository, userRepo user.Repository) *Service {
	return &Service{
		repo:     repo,
		userRepo: userRepo,
	}
}

func (s *Service) CreateTrainerCode(ctx context.Context, id, code string) error {
	user, err := s.repo.Find(ctx, id, "")
	if err != nil {
		return err
	}
	if user == nil {
		return domain.ErrTrainerNotFound
	}

	return s.repo.CreateTrainerCode(ctx, id, code)
}

func (s *Service) LinkTrainer(ctx context.Context, id, code string) error {
	trainer, err := s.repo.FindByTrainerCode(ctx, code)
	if err != nil {
		return err
	}
	if trainer == nil {
		return domain.ErrTrainerNotFound
	}

	now := time.Now().UTC()
	newId, err := utils.GenerateUUIDV7(ctx)
	if err != nil {
		return err
	}

	relation := &domain.TrainerStudentRelation{
		ID:        newId,
		CreatedAt: now,
		UpdatedAt: now,
		TrainerId: trainer.ID,
		StudentId: id,
		LinkedAt:  now,
	}

	return s.repo.LinkTrainer(ctx, *relation)
}

func (s *Service) UnlinkTrainer(ctx context.Context, studentId string) error {
	return s.repo.UnlinkTrainer(ctx, studentId)
}

func (s *Service) UnlinkStudent(ctx context.Context, studentId string) error {
	return s.repo.UnlinkTrainer(ctx, studentId)
}

func (s *Service) ListStudents(ctx context.Context, trainerId, cursor string, limit int) ([]*domain.User, string, error) {
	var decodedCursor *utils.CursorData
	utils.DecodeCursor(cursor, &decodedCursor)

	users, rawNextCursor, err := s.repo.ListStudents(ctx, trainerId, decodedCursor, limit)
	if err != nil {
		return nil, "", err
	}

	for _, u := range users {
		settings, err := s.userRepo.GetPrivacySettings(ctx, *u.ID)
		if err == nil && settings != nil {
			u.ApplyPrivacy(settings)
		}
		u.Sanitize()
	}

	nextCursorStr, _ := utils.EncodeCursor(rawNextCursor)
	return users, nextCursorStr, nil
}

func (s *Service) GetStudentPrivacy(ctx context.Context, trainerId, studentId string) (*domain.UserPrivacySettings, error) {
	linkedAt, err := s.repo.GetTrainerLinkDate(ctx, trainerId, studentId)
	if err != nil {
		return nil, err
	}
	if linkedAt == nil {
		return nil, domain.ErrUnauthorizedTrainerAccess
	}

	settings, err := s.userRepo.GetPrivacySettings(ctx, studentId)
	if err != nil {
		return nil, err
	}

	if settings == nil {
		return &domain.UserPrivacySettings{
			UserId:                studentId,
			ShareEmail:            true,
			ShareTrainingProgress: false,
			SharePastDataWithTrainer: false,
			ShareBodyMeasurements: false,
			ShareWeightLogs:       false,
			ShareMetricGoals:      false,
			AllowTrainerNotes:     true,
		}, nil
	}

	return settings, nil
}
