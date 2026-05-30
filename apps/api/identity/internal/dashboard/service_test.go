package dashboard

import (
	"context"
	"testing"
	"time"

	"github.com/kaua-nasc/gymtrack-go/apps/api/identity/internal/domain"
	"github.com/kaua-nasc/gymtrack-go/apps/api/identity/internal/trainer"
	"github.com/kaua-nasc/gymtrack-go/apps/api/identity/internal/user"
	"github.com/kaua-nasc/gymtrack-go/libs/auth"
	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
)

func TestService_GetStudentEngagement(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockUserRepo := user.NewMockRepository(ctrl)
	mockTrainerRepo := trainer.NewMockRepository(ctrl)
	mockTPClient := domain.NewMockTrainingPlanClient(ctrl)

	service := NewService(mockUserRepo, mockTrainerRepo, mockTPClient)

	ctx := context.Background()
	trainerId := "trainer-1"
	studentId := "student-1"
	token := "valid-token"

	ctx = context.WithValue(ctx, string(auth.TokenContextKey), token)

	t.Run("Success", func(t *testing.T) {
		linkedAt := time.Now()
		mockTrainerRepo.EXPECT().GetTrainerLinkDate(ctx, trainerId, studentId).Return(&linkedAt, nil)
		mockUserRepo.EXPECT().GetPrivacySettings(ctx, studentId).Return(&domain.UserPrivacySettings{
			ShareTrainingProgress: true,
		}, nil)

		expectedSummary := &domain.EngagementSummary{
			AdherenceRate:   85.0,
			CurrentPlanName: "Strong Body",
		}
		mockTPClient.EXPECT().GetEngagementSummary(ctx, studentId, token).Return(expectedSummary, nil)

		summary, err := service.GetStudentEngagement(ctx, trainerId, studentId)

		assert.NoError(t, err)
		assert.Equal(t, expectedSummary, summary)
	})

	t.Run("Unauthorized - Not Linked", func(t *testing.T) {
		mockTrainerRepo.EXPECT().GetTrainerLinkDate(ctx, trainerId, studentId).Return(nil, nil)

		summary, err := service.GetStudentEngagement(ctx, trainerId, studentId)

		assert.ErrorIs(t, err, domain.ErrUnauthorizedTrainerAccess)
		assert.Nil(t, summary)
	})

	t.Run("Forbidden - Privacy Settings", func(t *testing.T) {
		linkedAt := time.Now()
		mockTrainerRepo.EXPECT().GetTrainerLinkDate(ctx, trainerId, studentId).Return(&linkedAt, nil)
		mockUserRepo.EXPECT().GetPrivacySettings(ctx, studentId).Return(&domain.UserPrivacySettings{
			ShareTrainingProgress: false,
		}, nil)

		summary, err := service.GetStudentEngagement(ctx, trainerId, studentId)

		assert.ErrorIs(t, err, domain.ErrPrivacySettingsForbidden)
		assert.Nil(t, summary)
	})
}
