package internal

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

// MockUserRepository is a mock implementation of the UserRepository interface
type MockUserRepository struct {
	mock.Mock
}

func (m *MockUserRepository) Create(ctx context.Context, u *User) error {
	args := m.Called(ctx, u)
	return args.Error(0)
}

func (m *MockUserRepository) Update(ctx context.Context, u *User) error {
	args := m.Called(ctx, u)
	return args.Error(0)
}

func (m *MockUserRepository) Find(ctx context.Context, id string, currentUserId string) (*User, error) {
	args := m.Called(ctx, id, currentUserId)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*User), args.Error(1)
}

func (m *MockUserRepository) FindByEmail(ctx context.Context, email string) (*User, error) {
	args := m.Called(ctx, email)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*User), args.Error(1)
}

func (m *MockUserRepository) SaveResetCode(ctx context.Context, code, email string) error {
	args := m.Called(ctx, code, email)
	return args.Error(0)
}

func (m *MockUserRepository) GetResetCode(ctx context.Context, email string) (string, error) {
	args := m.Called(ctx, email)
	return args.String(0), args.Error(1)
}

func (m *MockUserRepository) ListByIDs(ctx context.Context, ids []string) ([]*User, error) {
	args := m.Called(ctx, ids)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).([]*User), args.Error(1)
}

func (m *MockUserRepository) ListFollowing(ctx context.Context, id string, cursor *CursorData, limit int) ([]*User, *CursorData, error) {
	args := m.Called(ctx, id, cursor, limit)
	var users []*User
	if args.Get(0) != nil {
		users = args.Get(0).([]*User)
	}
	var nextCursor *CursorData
	if args.Get(1) != nil {
		nextCursor = args.Get(1).(*CursorData)
	}
	return users, nextCursor, args.Error(2)
}

func (m *MockUserRepository) ListFollower(ctx context.Context, id string, cursor *CursorData, limit int) ([]*User, *CursorData, error) {
	args := m.Called(ctx, id, cursor, limit)
	var users []*User
	if args.Get(0) != nil {
		users = args.Get(0).([]*User)
	}
	var nextCursor *CursorData
	if args.Get(1) != nil {
		nextCursor = args.Get(1).(*CursorData)
	}
	return users, nextCursor, args.Error(2)
}

func (m *MockUserRepository) CountFollowers(ctx context.Context, userId string) (int, error) {
	args := m.Called(ctx, userId)
	return args.Int(0), args.Error(1)
}

func (m *MockUserRepository) CountFollowing(ctx context.Context, userId string) (int, error) {
	args := m.Called(ctx, userId)
	return args.Int(0), args.Error(1)
}

func (m *MockUserRepository) FollowUser(ctx context.Context, f UserFollows) error {
	args := m.Called(ctx, f)
	return args.Error(0)
}

func (m *MockUserRepository) UnfollowUser(ctx context.Context, followerId, followingId string) error {
	args := m.Called(ctx, followerId, followingId)
	return args.Error(0)
}

func (m *MockUserRepository) CreateTrainerCode(ctx context.Context, id, code string) error {
	args := m.Called(ctx, id, code)
	return args.Error(0)
}

func (m *MockUserRepository) FindByTrainerCode(ctx context.Context, code string) (*User, error) {
	args := m.Called(ctx, code)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*User), args.Error(1)
}

func (m *MockUserRepository) LinkTrainer(ctx context.Context, relation TrainerStudentRelation) error {
	args := m.Called(ctx, relation)
	return args.Error(0)
}

func (m *MockUserRepository) UnlinkTrainer(ctx context.Context, studentId string) error {
	args := m.Called(ctx, studentId)
	return args.Error(0)
}

func (m *MockUserRepository) ListStudents(ctx context.Context, trainerId string, cursor *CursorData, limit int) ([]*User, *CursorData, error) {
	args := m.Called(ctx, trainerId, cursor, limit)
	var users []*User
	if args.Get(0) != nil {
		users = args.Get(0).([]*User)
	}
	var nextCursor *CursorData
	if args.Get(1) != nil {
		nextCursor = args.Get(1).(*CursorData)
	}
	return users, nextCursor, args.Error(2)
}

func (m *MockUserRepository) AddBodyMeasurementNote(ctx context.Context, id, note string) error {
	args := m.Called(ctx, id, note)
	return args.Error(0)
}

func (m *MockUserRepository) FindLastBodyMeasurementNote(ctx context.Context, userId string) (*BodyMeasurement, error) {
	args := m.Called(ctx, userId)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*BodyMeasurement), args.Error(1)
}

func (m *MockUserRepository) ListBodyMeasurements(ctx context.Context, userId string, cursor *CursorData, limit int) ([]*BodyMeasurement, *CursorData, error) {
	args := m.Called(ctx, userId, cursor, limit)
	var measurements []*BodyMeasurement
	if args.Get(0) != nil {
		measurements = args.Get(0).([]*BodyMeasurement)
	}
	var nextCursor *CursorData
	if args.Get(1) != nil {
		nextCursor = args.Get(1).(*CursorData)
	}
	return measurements, nextCursor, args.Error(2)
}

func (m *MockUserRepository) AddWeightLogNote(ctx context.Context, id, note string) error {
	args := m.Called(ctx, id, note)
	return args.Error(0)
}

func (m *MockUserRepository) ChangeUserType(ctx context.Context, u User, newType UserType) error {
	args := m.Called(ctx, u, newType)
	return args.Error(0)
}

func (m *MockUserRepository) RemoveProfilePicture(ctx context.Context, userId string) error {
	args := m.Called(ctx, userId)
	return args.Error(0)
}

func (m *MockUserRepository) ChangeProfileImage(ctx context.Context, u User, pictureUrl string) error {
	return m.Called(ctx, u, pictureUrl).Error(0)
}

func (m *MockUserRepository) ListGoalsMetric(ctx context.Context, id string, cursor *CursorData, limit int) ([]*MetricGoal, *CursorData, error) {
	args := m.Called(ctx, id, cursor, limit)
	var goals []*MetricGoal
	if args.Get(0) != nil {
		goals = args.Get(0).([]*MetricGoal)
	}
	var nextCursor *CursorData
	if args.Get(1) != nil {
		nextCursor = args.Get(1).(*CursorData)
	}
	return goals, nextCursor, args.Error(2)
}

func (m *MockUserRepository) AddGoalMetric(ctx context.Context, g MetricGoal) error {
	return m.Called(ctx, g).Error(0)
}

func (m *MockUserRepository) ListWeightLogs(ctx context.Context, userId string, cursor *CursorData, limit int) ([]*WeightLog, *CursorData, error) {
	args := m.Called(ctx, userId, cursor, limit)
	var logs []*WeightLog
	if args.Get(0) != nil {
		logs = args.Get(0).([]*WeightLog)
	}
	var nextCursor *CursorData
	if args.Get(1) != nil {
		nextCursor = args.Get(1).(*CursorData)
	}
	return logs, nextCursor, args.Error(2)
}

func TestUserService_Register(t *testing.T) {
	mockRepo := new(MockUserRepository)
	service := NewUserService(mockRepo)

	tests := []struct {
		name          string
		inputUser     User
		mockBehavior  func()
		wantErr       bool
		expectedError string
	}{
		{
			name: "Success registration",
			inputUser: User{
				Email:     "newuser@example.com",
				Password:  "strongpassword123",
				FirstName: "John",
				LastName:  "Doe",
			},
			mockBehavior: func() {
				mockRepo.On("FindByEmail", mock.Anything, "newuser@example.com").Return(nil, nil).Once()
				mockRepo.On("Create", mock.Anything, mock.AnythingOfType("*internal.User")).Return(nil).Once()
			},
			wantErr: false,
		},
		{
			name: "Error: User already exists",
			inputUser: User{
				Email: "existing@example.com",
			},
			mockBehavior: func() {
				mockRepo.On("FindByEmail", mock.Anything, "existing@example.com").Return(&User{ID: "123"}, nil).Once()
			},
			wantErr:       true,
			expectedError: "user already exists",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.mockBehavior()
			err := service.Register(context.Background(), tt.inputUser)

			if tt.wantErr {
				assert.Error(t, err)
				assert.Equal(t, tt.expectedError, err.Error())
			} else {
				assert.NoError(t, err)
			}
			mockRepo.AssertExpectations(t)
		})
	}
}
