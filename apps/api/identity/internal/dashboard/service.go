package dashboard

import (
	"context"

	"github.com/kaua-nasc/gymtrack-go/apps/api/identity/internal/domain"
	"github.com/kaua-nasc/gymtrack-go/apps/api/identity/internal/trainer"
	"github.com/kaua-nasc/gymtrack-go/apps/api/identity/internal/user"
	"github.com/kaua-nasc/gymtrack-go/libs/auth"
)

type Service struct {
	userRepo           user.Repository
	trainerRepo        trainer.Repository
	trainingPlanClient domain.TrainingPlanClient
}

func NewService(userRepo user.Repository, trainerRepo trainer.Repository, trainingPlanClient domain.TrainingPlanClient) *Service {
	return &Service{
		userRepo:           userRepo,
		trainerRepo:        trainerRepo,
		trainingPlanClient: trainingPlanClient,
	}
}

func (s *Service) GetStudentEngagement(ctx context.Context, trainerId, studentId string) (*domain.EngagementSummary, error) {
	// 1. Check if they are linked
	linkedAt, err := s.trainerRepo.GetTrainerLinkDate(ctx, trainerId, studentId)
	if err != nil {
		return nil, err
	}
	if linkedAt == nil {
		return nil, domain.ErrUnauthorizedTrainerAccess
	}

	// 2. Check privacy settings
	privacy, err := s.userRepo.GetPrivacySettings(ctx, studentId)
	if err != nil {
		return nil, err
	}

	if privacy != nil && !privacy.ShareTrainingProgress {
		return nil, domain.ErrPrivacySettingsForbidden
	}

	// 3. Fetch summary from training-plan microservice
	token, _ := ctx.Value(string(auth.TokenContextKey)).(string)
	summary, err := s.trainingPlanClient.GetEngagementSummary(ctx, studentId, token)
	if err != nil {
		return nil, err
	}

	return summary, nil
}
