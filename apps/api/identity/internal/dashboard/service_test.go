package dashboard

import (
	"context"
	"testing"
	"time"

	"github.com/kaua-nasc/gymtrack-go/apps/api/identity/internal/domain"
	"github.com/kaua-nasc/gymtrack-go/apps/api/identity/internal/metrics"
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
	mockMetricsRepo := metrics.NewMockRepository(ctrl)
	mockTPClient := domain.NewMockTrainingPlanClient(ctrl)

	service := NewService(mockUserRepo, mockTrainerRepo, mockMetricsRepo, mockTPClient)

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
}

func TestService_GetStudentBiometrics(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockUserRepo := user.NewMockRepository(ctrl)
	mockTrainerRepo := trainer.NewMockRepository(ctrl)
	mockMetricsRepo := metrics.NewMockRepository(ctrl)
	mockTPClient := domain.NewMockTrainingPlanClient(ctrl)

	service := NewService(mockUserRepo, mockTrainerRepo, mockMetricsRepo, mockTPClient)

	ctx := context.Background()
	trainerId := "trainer-1"
	studentId := "student-1"
	start := time.Now().AddDate(0, 0, -30)
	end := time.Now()

	t.Run("Success", func(t *testing.T) {
		linkedAt := time.Now()
		mockTrainerRepo.EXPECT().GetTrainerLinkDate(ctx, trainerId, studentId).Return(&linkedAt, nil)
		mockUserRepo.EXPECT().GetPrivacySettings(ctx, studentId).Return(&domain.UserPrivacySettings{
			ShareWeightLogs:       true,
			ShareBodyMeasurements: true,
		}, nil)
		mockUserRepo.EXPECT().Find(ctx, studentId, studentId).Return(&domain.User{Height: floatPtr(180)}, nil)

		weightHistory := []*domain.WeightLog{
			{Weight: 80.0, MeasuredAt: start},
			{Weight: 78.5, MeasuredAt: end},
		}
		mockMetricsRepo.EXPECT().ListWeightHistory(ctx, studentId, start, end).Return(weightHistory, nil)
		mockMetricsRepo.EXPECT().ListMeasurementsHistory(ctx, studentId, start, end).Return([]*domain.BodyMeasurement{}, nil)

		res, err := service.GetStudentBiometrics(ctx, trainerId, studentId, start, end)

		assert.NoError(t, err)
		assert.NotNil(t, res.Weight)
		assert.Equal(t, 78.5, res.Weight.Current)
		assert.Equal(t, -1.5, res.Weight.Delta7Days) // simplified delta logic in test
	})

	t.Run("Invalid Period", func(t *testing.T) {
		startLong := time.Now().AddDate(-2, 0, 0)
		res, err := service.GetStudentBiometrics(ctx, trainerId, studentId, startLong, end)
		assert.ErrorIs(t, err, domain.ErrInvalidPeriod)
		assert.Nil(t, res)
	})
}

func floatPtr(f float64) *float64 { return &f }
