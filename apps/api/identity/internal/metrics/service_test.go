package metrics

import (
	"context"
	"testing"
	"time"

	"github.com/kaua-nasc/gymtrack-go/apps/api/identity/internal/domain"
	"github.com/kaua-nasc/gymtrack-go/apps/api/identity/internal/trainer"
	"github.com/kaua-nasc/gymtrack-go/apps/api/identity/internal/user"
	"github.com/stretchr/testify/assert"
	gomock "go.uber.org/mock/gomock"
)

func TestUserService_BodyMeasurements(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	mockRepo := NewMockRepository(ctrl)
	mockUserRepo := user.NewMockRepository(ctrl)
	mockTrainerRepo := trainer.NewMockRepository(ctrl)
	service := NewService(mockRepo, mockUserRepo, mockTrainerRepo)

	// Test AddBodyMeasurementNote
	measurement := &domain.BodyMeasurement{ID: "meas-123", UserId: "user-123"}
	mockRepo.EXPECT().FindBodyMeasurement(gomock.Any(), "meas-123").Return(measurement, nil)
	now := time.Now()
	mockTrainerRepo.EXPECT().GetTrainerLinkDate(gomock.Any(), "123", "user-123").Return(&now, nil)
	mockUserRepo.EXPECT().GetPrivacySettings(gomock.Any(), "user-123").Return(&domain.UserPrivacySettings{AllowTrainerNotes: true}, nil)
	mockRepo.EXPECT().AddBodyMeasurementNote(gomock.Any(), "meas-123", "Nice progress").Return(nil)
	err := service.AddBodyMeasurementNote(context.Background(), "123", "meas-123", "Nice progress")
	assert.NoError(t, err)

	// Test FindLastBodyMeasurementNote
	expectedMeasurement := &domain.BodyMeasurement{ID: "meas-123", Value: 85.5}
	mockTrainerRepo.EXPECT().GetTrainerLinkDate(gomock.Any(), "123", "user-123").Return(&now, nil)
	mockUserRepo.EXPECT().GetPrivacySettings(gomock.Any(), "user-123").Return(&domain.UserPrivacySettings{ShareBodyMeasurements: true}, nil)
	mockRepo.EXPECT().FindLastBodyMeasurementNote(gomock.Any(), "user-123").Return(expectedMeasurement, nil)
	meas, err := service.FindLastBodyMeasurementNote(context.Background(), "123", "user-123")
	assert.NoError(t, err)
	assert.Equal(t, expectedMeasurement, meas)

	// Test ListBodyMeasurements
	mockTrainerRepo.EXPECT().GetTrainerLinkDate(gomock.Any(), "123", "user-123").Return(&now, nil)
	mockUserRepo.EXPECT().GetPrivacySettings(gomock.Any(), "user-123").Return(&domain.UserPrivacySettings{ShareBodyMeasurements: true, SharePastDataWithTrainer: true}, nil).Times(2)
	mockRepo.EXPECT().ListBodyMeasurements(gomock.Any(), "user-123", gomock.Any(), gomock.Any(), 10).Return([]*domain.BodyMeasurement{
		expectedMeasurement,
	}, nil, nil)
	measurements, nextCursor, err := service.ListBodyMeasurements(context.Background(), "123", "user-123", "", 10)
	assert.NoError(t, err)
	assert.Len(t, measurements, 1)
	assert.Equal(t, "bnVsbA==", nextCursor)
}

func TestUserService_WeightLogs(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	mockRepo := NewMockRepository(ctrl)
	mockUserRepo := user.NewMockRepository(ctrl)
	mockTrainerRepo := trainer.NewMockRepository(ctrl)
	service := NewService(mockRepo, mockUserRepo, mockTrainerRepo)

	now := time.Now()

	// Test AddWeightLogNote
	logEntry := &domain.WeightLog{ID: "log-123", UserId: "user-123"}
	mockRepo.EXPECT().FindWeightLog(gomock.Any(), "log-123").Return(logEntry, nil)
	mockTrainerRepo.EXPECT().GetTrainerLinkDate(gomock.Any(), "trainer", "user-123").Return(&now, nil)
	mockUserRepo.EXPECT().GetPrivacySettings(gomock.Any(), "user-123").Return(&domain.UserPrivacySettings{AllowTrainerNotes: true}, nil)
	mockRepo.EXPECT().AddWeightLogNote(gomock.Any(), "log-123", "Good day").Return(nil)
	err := service.AddWeightLogNote(context.Background(), "trainer", "log-123", "Good day")
	assert.NoError(t, err)

	// Test ListWeightLogs
	mockTrainerRepo.EXPECT().GetTrainerLinkDate(gomock.Any(), "a", "user-123").Return(&now, nil)
	mockUserRepo.EXPECT().GetPrivacySettings(gomock.Any(), "user-123").Return(&domain.UserPrivacySettings{ShareWeightLogs: true, SharePastDataWithTrainer: true}, nil).Times(2)
	mockRepo.EXPECT().ListWeightLogs(gomock.Any(), "user-123", gomock.Any(), gomock.Any(), 10).Return([]*domain.WeightLog{
		{ID: "log-123", Weight: 80.0},
	}, nil, nil)
	logs, nextCursor, err := service.ListWeightLogs(context.Background(), "a", "user-123", "", 10)
	assert.NoError(t, err)
	assert.Len(t, logs, 1)
	assert.Equal(t, "bnVsbA==", nextCursor)
}

func TestUserService_GoalsMetric(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	mockRepo := NewMockRepository(ctrl)
	mockUserRepo := user.NewMockRepository(ctrl)
	mockTrainerRepo := trainer.NewMockRepository(ctrl)
	service := NewService(mockRepo, mockUserRepo, mockTrainerRepo)

	now := time.Now()

	// Test AddGoalMetric
	goal := &domain.MetricGoal{
		UserId: "user-123",
	}
	mockRepo.EXPECT().AddGoalMetric(gomock.Any(), gomock.Any()).Return(nil)
	err := service.AddGoalMetric(context.Background(), "user-123", goal) // requester must match userId
	assert.NoError(t, err)
	assert.NotEmpty(t, goal.ID)

	// Test ListGoalsMetric
	mockTrainerRepo.EXPECT().GetTrainerLinkDate(gomock.Any(), "requester", "user-123").Return(&now, nil)
	mockUserRepo.EXPECT().GetPrivacySettings(gomock.Any(), "user-123").Return(&domain.UserPrivacySettings{ShareMetricGoals: true, SharePastDataWithTrainer: true}, nil).Times(2)
	mockRepo.EXPECT().ListGoalsMetric(gomock.Any(), "user-123", gomock.Any(), gomock.Any(), 10).Return([]*domain.MetricGoal{
		{ID: "goal-1", UserId: "user-123"},
	}, nil, nil)

	goals, nextCursor, err := service.ListGoalsMetric(context.Background(), "requester", "user-123", "", 10)
	assert.NoError(t, err)
	assert.Len(t, goals, 1)
	assert.Equal(t, "bnVsbA==", nextCursor)
}
