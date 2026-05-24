package trainer

import (
	"context"
	"testing"

	"github.com/kaua-nasc/gymtrack-go/apps/api/identity/internal/domain"
	"github.com/kaua-nasc/gymtrack-go/apps/api/identity/internal/user"
	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
)

func TestUserService_CreateTrainerCode(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := NewMockRepository(ctrl)
	mockUserRepo := user.NewMockRepository(ctrl)
	service := NewService(mockRepo, mockUserRepo)

	trainerIdStr := "trainer-123"

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
				mockRepo.EXPECT().Find(gomock.Any(), "trainer-123", "").Return(&domain.User{ID: &trainerIdStr}, nil)
				mockRepo.EXPECT().CreateTrainerCode(gomock.Any(), "trainer-123", "TRAIN12").Return(nil)
			},
			wantErr: false,
		},
		{
			name: "Error: Trainer not found",
			id:   "nonexistent",
			code: "TRAIN12",
			mockBehavior: func() {
				mockRepo.EXPECT().Find(gomock.Any(), "nonexistent", "").Return(nil, nil)
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
		})
	}
}

func TestUserService_LinkTrainer(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := NewMockRepository(ctrl)
	mockUserRepo := user.NewMockRepository(ctrl)
	service := NewService(mockRepo, mockUserRepo)

	trainerIdStr := "trainer-123"

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
				mockRepo.EXPECT().FindByTrainerCode(gomock.Any(), "TRAIN12").Return(&domain.User{ID: &trainerIdStr}, nil)
				mockRepo.EXPECT().LinkTrainer(gomock.Any(), gomock.Any()).Return(nil)
			},
			wantErr: false,
		},
		{
			name: "Error: Trainer not found",
			id:   "student-123",
			code: "TRAIN12",
			mockBehavior: func() {
				mockRepo.EXPECT().FindByTrainerCode(gomock.Any(), "TRAIN12").Return(nil, nil)
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
		})
	}
}

func TestUserService_UnlinkTrainerAndStudent(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := NewMockRepository(ctrl)
	mockUserRepo := user.NewMockRepository(ctrl)
	service := NewService(mockRepo, mockUserRepo)

	mockRepo.EXPECT().UnlinkTrainer(gomock.Any(), "student-123").Return(nil).Times(2)

	err := service.UnlinkTrainer(context.Background(), "student-123")
	assert.NoError(t, err)

	err = service.UnlinkStudent(context.Background(), "student-123")
	assert.NoError(t, err)
}

func TestUserService_ListStudents(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := NewMockRepository(ctrl)
	mockUserRepo := user.NewMockRepository(ctrl)
	service := NewService(mockRepo, mockUserRepo)

	studentIdStr := "student-1"

	mockRepo.EXPECT().ListStudents(gomock.Any(), "trainer-123", gomock.Any(), 10).Return([]*domain.User{
		{ID: &studentIdStr},
	}, nil, nil)

	mockUserRepo.EXPECT().GetPrivacySettings(gomock.Any(), "student-1").Return(nil, nil)

	users, nextCursor, err := service.ListStudents(context.Background(), "trainer-123", "", 10)
	assert.NoError(t, err)
	assert.Len(t, users, 1)
	assert.Equal(t, "bnVsbA==", nextCursor)
}
