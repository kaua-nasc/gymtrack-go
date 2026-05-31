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

func TestService_GetStudentInsights(t *testing.T) {
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

	t.Run("Success with Multiple Insights", func(t *testing.T) {
		linkedAt := time.Now()
		mockTrainerRepo.EXPECT().GetTrainerLinkDate(ctx, trainerId, studentId).Return(&linkedAt, nil)
		mockUserRepo.EXPECT().GetPrivacySettings(ctx, studentId).Return(&domain.UserPrivacySettings{
			ShareWeightLogs:       true,
			ShareBodyMeasurements: true,
			ShareTrainingProgress: true,
		}, nil)

		// 1. Pending review
		mockMetricsRepo.EXPECT().CountUnreviewedMetrics(ctx, studentId).Return(2, 1, nil)

		// 2. Inactivity
		lastWorkout := time.Now().AddDate(0, 0, -5)
		mockTPClient.EXPECT().GetEngagementSummary(ctx, studentId, token).Return(&domain.EngagementSummary{
			LastWorkoutDate: &lastWorkout,
		}, nil)

		// 3. Stagnation
		mockMetricsRepo.EXPECT().ListWeightHistory(ctx, studentId, gomock.Any(), gomock.Any()).Return([]*domain.WeightLog{
			{Weight: 80.0}, {Weight: 80.1}, {Weight: 80.0}, {Weight: 80.1},
		}, nil)

		res, err := service.GetStudentInsights(ctx, trainerId, studentId)

		assert.NoError(t, err)
		assert.Len(t, res.Insights, 4)
		types := []string{}
		for _, i := range res.Insights {
			types = append(types, i.Type)
		}
		assert.Contains(t, types, "PENDING_REVIEW_WEIGHT")
		assert.Contains(t, types, "PENDING_REVIEW_MEASUREMENTS")
		assert.Contains(t, types, "INACTIVITY_ALERT")
		assert.Contains(t, types, "STAGNATION_WEIGHT")
	})
}

func floatPtr(f float64) *float64 { return &f }
