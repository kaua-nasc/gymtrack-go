package metrics

import (
	"context"
	"testing"

	"github.com/kaua-nasc/gymtrack-go/apps/api/identity/internal/domain"
	"github.com/kaua-nasc/gymtrack-go/libs/utils"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type MockRepository struct {
	mock.Mock
}

func (m *MockRepository) AddBodyMeasurementNote(ctx context.Context, id, note string) error {
	args := m.Called(ctx, id, note)
	return args.Error(0)
}

func (m *MockRepository) FindLastBodyMeasurementNote(ctx context.Context, userId string) (*domain.BodyMeasurement, error) {
	args := m.Called(ctx, userId)
	bm, _ := args.Get(0).(*domain.BodyMeasurement)
	return bm, args.Error(1)
}

func (m *MockRepository) ListBodyMeasurements(ctx context.Context, userId string, cursor *utils.CursorData, limit int) ([]*domain.BodyMeasurement, *utils.CursorData, error) {
	args := m.Called(ctx, userId, cursor, limit)
	bms, _ := args.Get(0).([]*domain.BodyMeasurement)
	nextCursor, _ := args.Get(1).(*utils.CursorData)
	return bms, nextCursor, args.Error(2)
}

func (m *MockRepository) AddWeightLogNote(ctx context.Context, id, note string) error {
	args := m.Called(ctx, id, note)
	return args.Error(0)
}

func (m *MockRepository) ListGoalsMetric(ctx context.Context, id string, cursor *utils.CursorData, limit int) ([]*domain.MetricGoal, *utils.CursorData, error) {
	args := m.Called(ctx, id, cursor, limit)
	goals, _ := args.Get(0).([]*domain.MetricGoal)
	nextCursor, _ := args.Get(1).(*utils.CursorData)
	return goals, nextCursor, args.Error(2)
}

func (m *MockRepository) ListWeightLogs(ctx context.Context, userId string, cursor *utils.CursorData, limit int) ([]*domain.WeightLog, *utils.CursorData, error) {
	args := m.Called(ctx, userId, cursor, limit)
	logs, _ := args.Get(0).([]*domain.WeightLog)
	nextCursor, _ := args.Get(1).(*utils.CursorData)
	return logs, nextCursor, args.Error(2)
}

func (m *MockRepository) AddGoalMetric(ctx context.Context, g domain.MetricGoal) error {
	args := m.Called(ctx, g)
	return args.Error(0)
}

func TestUserService_BodyMeasurements(t *testing.T) {
	mockRepo := new(MockRepository)
	service := NewService(mockRepo)

	// Test AddBodyMeasurementNote
	mockRepo.On("AddBodyMeasurementNote", mock.Anything, "meas-123", "Nice progress").Return(nil).Once()
	err := service.AddBodyMeasurementNote(context.Background(), "meas-123", "Nice progress")
	assert.NoError(t, err)

	// Test FindLastBodyMeasurementNote
	expectedMeasurement := &domain.BodyMeasurement{ID: "meas-123", Value: 85.5}
	mockRepo.On("FindLastBodyMeasurementNote", mock.Anything, "user-123").Return(expectedMeasurement, nil).Once()
	meas, err := service.FindLastBodyMeasurementNote(context.Background(), "user-123")
	assert.NoError(t, err)
	assert.Equal(t, expectedMeasurement, meas)

	// Test ListBodyMeasurements
	mockRepo.On("ListBodyMeasurements", mock.Anything, "user-123", mock.Anything, 10).Return([]*domain.BodyMeasurement{
		expectedMeasurement,
	}, nil, nil).Once()
	measurements, nextCursor, err := service.ListBodyMeasurements(context.Background(), "user-123", "", 10)
	assert.NoError(t, err)
	assert.Len(t, measurements, 1)
	assert.Equal(t, "bnVsbA==", nextCursor)

	mockRepo.AssertExpectations(t)
}

func TestUserService_WeightLogs(t *testing.T) {
	mockRepo := new(MockRepository)
	service := NewService(mockRepo)

	// Test AddWeightLogNote
	mockRepo.On("AddWeightLogNote", mock.Anything, "log-123", "Good day").Return(nil).Once()
	err := service.AddWeightLogNote(context.Background(), "log-123", "Good day")
	assert.NoError(t, err)

	// Test ListWeightLogs
	mockRepo.On("ListWeightLogs", mock.Anything, "user-123", mock.Anything, 10).Return([]*domain.WeightLog{
		{ID: "log-123", Weight: 80.0},
	}, nil, nil).Once()
	logs, nextCursor, err := service.ListWeightLogs(context.Background(), "user-123", "", 10)
	assert.NoError(t, err)
	assert.Len(t, logs, 1)
	assert.Equal(t, "bnVsbA==", nextCursor)

	mockRepo.AssertExpectations(t)
}

func TestUserService_GoalsMetric(t *testing.T) {
	mockRepo := new(MockRepository)
	service := NewService(mockRepo)

	// Test AddGoalMetric
	goal := &domain.MetricGoal{
		UserId: "user-123",
	}
	mockRepo.On("AddGoalMetric", mock.Anything, mock.MatchedBy(func(g domain.MetricGoal) bool {
		return g.UserId == "user-123" && g.Status == domain.MetricGoalActive
	})).Return(nil).Once()

	err := service.AddGoalMetric(context.Background(), goal)
	assert.NoError(t, err)
	assert.NotEmpty(t, goal.ID)

	// Test ListGoalsMetric
	mockRepo.On("ListGoalsMetric", mock.Anything, "user-123", mock.Anything, 10).Return([]*domain.MetricGoal{
		{ID: "goal-1", UserId: "user-123"},
	}, nil, nil).Once()

	goals, nextCursor, err := service.ListGoalsMetric(context.Background(), "user-123", "", 10)
	assert.NoError(t, err)
	assert.Len(t, goals, 1)
	assert.Equal(t, "bnVsbA==", nextCursor)

	mockRepo.AssertExpectations(t)
}
