package trainer

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

func (m *MockRepository) Find(ctx context.Context, id string, currentUserId string) (*domain.User, error) {
	args := m.Called(ctx, id, currentUserId)
	u, _ := args.Get(0).(*domain.User)
	return u, args.Error(1)
}

func (m *MockRepository) CreateTrainerCode(ctx context.Context, id, code string) error {
	args := m.Called(ctx, id, code)
	return args.Error(0)
}

func (m *MockRepository) FindByTrainerCode(ctx context.Context, code string) (*domain.User, error) {
	args := m.Called(ctx, code)
	u, _ := args.Get(0).(*domain.User)
	return u, args.Error(1)
}

func (m *MockRepository) LinkTrainer(ctx context.Context, relation domain.TrainerStudentRelation) error {
	args := m.Called(ctx, relation)
	return args.Error(0)
}

func (m *MockRepository) UnlinkTrainer(ctx context.Context, studentId string) error {
	args := m.Called(ctx, studentId)
	return args.Error(0)
}

func (m *MockRepository) ListStudents(ctx context.Context, trainerId string, cursor *utils.CursorData, limit int) ([]*domain.User, *utils.CursorData, error) {
	args := m.Called(ctx, trainerId, cursor, limit)
	users, _ := args.Get(0).([]*domain.User)
	nextCursor, _ := args.Get(1).(*utils.CursorData)
	return users, nextCursor, args.Error(2)
}

func newString(s string) *string {
	return &s
}

func TestUserService_CreateTrainerCode(t *testing.T) {
	mockRepo := new(MockRepository)
	service := NewService(mockRepo)

	tests := []struct {
		name          string
		id            string
		code          string
		mockBehavior  func()
		wantErr       bool
		expectedError string
	}{
		{
			name: "Success create trainer code",
			id:   "trainer-123",
			code: "TRAIN12",
			mockBehavior: func() {
				mockRepo.On("Find", mock.Anything, "trainer-123", "").Return(&domain.User{ID: newString("trainer-123")}, nil).Once()
				mockRepo.On("CreateTrainerCode", mock.Anything, "trainer-123", "TRAIN12").Return(nil).Once()
			},
			wantErr: false,
		},
		{
			name: "Error: Trainer not found",
			id:   "nonexistent",
			code: "TRAIN12",
			mockBehavior: func() {
				mockRepo.On("Find", mock.Anything, "nonexistent", "").Return(nil, nil).Once()
			},
			wantErr:       true,
			expectedError: "trainer not found",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.mockBehavior()
			err := service.CreateTrainerCode(context.Background(), tt.id, tt.code)

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

func TestUserService_LinkTrainer(t *testing.T) {
	mockRepo := new(MockRepository)
	service := NewService(mockRepo)

	tests := []struct {
		name          string
		id            string
		code          string
		mockBehavior  func()
		wantErr       bool
		expectedError string
	}{
		{
			name: "Success link trainer",
			id:   "student-123",
			code: "TRAIN12",
			mockBehavior: func() {
				mockRepo.On("FindByTrainerCode", mock.Anything, "TRAIN12").Return(&domain.User{ID: newString("trainer-123")}, nil).Once()
				mockRepo.On("LinkTrainer", mock.Anything, mock.MatchedBy(func(r domain.TrainerStudentRelation) bool {
					return *r.TrainerId == "trainer-123" && r.StudentId == "student-123"
				})).Return(nil).Once()
			},
			wantErr: false,
		},
		{
			name: "Error: Trainer not found",
			id:   "student-123",
			code: "TRAIN12",
			mockBehavior: func() {
				mockRepo.On("FindByTrainerCode", mock.Anything, "TRAIN12").Return(nil, nil).Once()
			},
			wantErr:       true,
			expectedError: "trainer not found",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.mockBehavior()
			err := service.LinkTrainer(context.Background(), tt.id, tt.code)

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

func TestUserService_UnlinkTrainerAndStudent(t *testing.T) {
	mockRepo := new(MockRepository)
	service := NewService(mockRepo)

	mockRepo.On("UnlinkTrainer", mock.Anything, "student-123").Return(nil).Twice()

	err := service.UnlinkTrainer(context.Background(), "student-123")
	assert.NoError(t, err)

	err = service.UnlinkStudent(context.Background(), "student-123")
	assert.NoError(t, err)

	mockRepo.AssertExpectations(t)
}

func TestUserService_ListStudents(t *testing.T) {
	mockRepo := new(MockRepository)
	service := NewService(mockRepo)

	mockRepo.On("ListStudents", mock.Anything, "trainer-123", mock.Anything, 10).Return([]*domain.User{
		{ID: newString("student-1")},
	}, nil, nil).Once()

	users, nextCursor, err := service.ListStudents(context.Background(), "trainer-123", "", 10)
	assert.NoError(t, err)
	assert.Len(t, users, 1)
	assert.Equal(t, "bnVsbA==", nextCursor)
	mockRepo.AssertExpectations(t)
}
