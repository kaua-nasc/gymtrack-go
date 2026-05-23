package user

import (
	"context"
	"fmt"
	"io"
	"strings"
	"time"

	"github.com/kaua-nasc/gymtrack-go/apps/api/identity/internal/domain"
	"github.com/kaua-nasc/gymtrack-go/libs/auth"
	"github.com/kaua-nasc/gymtrack-go/libs/storage"
)

type Service struct {
	repo Repository
}

func NewService(repo Repository) *Service {
	return &Service{repo: repo}
}

func (s *Service) ChangePassword(ctx context.Context, id, currentPassword, newPassword string) error {
	u, err := s.repo.Find(ctx, id, "")
	if err != nil {
		return err
	}
	if u == nil {
		return domain.ErrUserNotFound
	}

	if !u.IsVerified {
		return domain.ErrNotVerified
	}

	ok, err := auth.VerifyArgon2Password(currentPassword, u.Password)
	if err != nil || !ok {
		return domain.ErrInvalidCredentials
	}

	hashedPassword, err := auth.HashArgon2Password(newPassword)
	if err != nil {
		return err
	}
	u.Password = hashedPassword
	u.UpdatedAt = time.Now().UTC()

	return s.repo.Update(ctx, u)
}

func (s *Service) UpdateProfile(
	ctx context.Context,
	id string,
	firstName, lastName, bio *string,
	height *float64,
	weightUnit *domain.WeightUnit,
	heightUnit *domain.HeightUnit,
	currentWeight *float64,
) error {
	u, err := s.repo.Find(ctx, id, "")
	if err != nil {
		return err
	}
	if u == nil {
		return domain.ErrUserNotFound
	}

	if firstName != nil {
		u.FirstName = *firstName
	}
	if lastName != nil {
		u.LastName = *lastName
	}
	if bio != nil {
		u.Bio = bio
	}
	if height != nil {
		u.Height = height
	}
	if weightUnit != nil {
		u.WeightUnit = *weightUnit
	}
	if heightUnit != nil {
		u.HeightUnit = *heightUnit
	}
	if currentWeight != nil {
		u.CurrentWeight = currentWeight
	}

	u.UpdatedAt = time.Now().UTC()

	return s.repo.Update(ctx, u)
}

func (s *Service) GetUser(ctx context.Context, id string, requesterId string) (*domain.User, error) {
	u, err := s.repo.Find(ctx, id, requesterId)
	if err != nil || u == nil {
		return u, err
	}

	if id != requesterId {
		settings, err := s.repo.GetPrivacySettings(ctx, id)
		if err == nil && settings != nil {
			u.ApplyPrivacy(settings)
		}
	}

	u.Sanitize()
	return u, nil
}

func (s *Service) ListUsers(ctx context.Context, requesterId string, ids []string) ([]*domain.User, error) {
	users, err := s.repo.ListByIDs(ctx, ids)
	if err != nil {
		return nil, err
	}
	for _, u := range users {
		if requesterId != *u.ID {
			settings, err := s.repo.GetPrivacySettings(ctx, *u.ID)
			if err == nil && settings != nil {
				u.ApplyPrivacy(settings)
			}
		}
		u.Sanitize()
	}
	return users, nil
}

func (s *Service) ChangeToTrainer(ctx context.Context, id, cref string) error {
	userVal, err := s.repo.Find(ctx, id, "")
	if err != nil {
		return err
	}
	if userVal == nil {
		return domain.ErrUserNotFound
	}

	return s.repo.ChangeUserType(ctx, *userVal, domain.Trainer)
}

func (s *Service) ChangeToClient(ctx context.Context, id string) error {
	userVal, err := s.repo.Find(ctx, id, "")
	if err != nil {
		return err
	}
	if userVal == nil {
		return domain.ErrUserNotFound
	}

	return s.repo.ChangeUserType(ctx, *userVal, domain.Client)
}

func (s *Service) RemoveProfilePicture(ctx context.Context, id string) error {
	return s.repo.RemoveProfilePicture(ctx, id)
}

func (s *Service) UploadProfilePicture(ctx context.Context, id string, file io.Reader) error {
	userVal, err := s.repo.Find(ctx, id, "")
	if err != nil {
		return err
	}
	if userVal == nil {
		return domain.ErrUserNotFound
	}

	bytesVal, err := io.ReadAll(file)
	if err != nil {
		return err
	}

	timestamp := strings.ReplaceAll(time.Now().UTC().String(), ":", "-")
	timestamp = strings.ReplaceAll(timestamp, ".", "-")
	filename := `identity/user/profile/user-` + id + `_` + timestamp + `.png`

	if err := storage.UploadBuffer(ctx, filename, bytesVal); err != nil {
		return fmt.Errorf("falha no upload: %w", err)
	}

	return s.repo.ChangeProfileImage(ctx, *userVal, filename)
}

func (s *Service) GetPrivacySettings(ctx context.Context, userId string) (*domain.UserPrivacySettings, error) {
	u, err := s.repo.Find(ctx, userId, "")
	if err != nil {
		return nil, err
	}
	if u == nil {
		return nil, domain.ErrUserNotFound
	}

	settings, err := s.repo.GetPrivacySettings(ctx, userId)
	if err != nil {
		return nil, err
	}

	if settings == nil {
		// Default settings if not created yet
		return &domain.UserPrivacySettings{
			UserId:                userId,
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

func (s *Service) UpdatePrivacySettings(
	ctx context.Context,
	userId string,
	shareEmail, shareTrainingProgress, sharePastDataWithTrainer,
	shareBodyMeasurements, shareWeightLogs, shareMetricGoals,
	allowTrainerNotes bool,
) error {
	u, err := s.repo.Find(ctx, userId, "")
	if err != nil {
		return err
	}
	if u == nil {
		return domain.ErrUserNotFound
	}

	settings := domain.UserPrivacySettings{
		UserId:                   userId,
		ShareEmail:               shareEmail,
		ShareTrainingProgress:    shareTrainingProgress,
		SharePastDataWithTrainer: sharePastDataWithTrainer,
		ShareBodyMeasurements:    shareBodyMeasurements,
		ShareWeightLogs:          shareWeightLogs,
		ShareMetricGoals:         shareMetricGoals,
		AllowTrainerNotes:        allowTrainerNotes,
	}

	return s.repo.UpsertPrivacySettings(ctx, settings)
}
