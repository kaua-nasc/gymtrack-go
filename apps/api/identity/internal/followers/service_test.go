package followers

import (
	"context"
	"testing"

	"github.com/kaua-nasc/gymtrack-go/apps/api/identity/internal/domain"
	"github.com/kaua-nasc/gymtrack-go/apps/api/identity/internal/user"
	"github.com/stretchr/testify/assert"
	gomock "go.uber.org/mock/gomock"
)

func TestUserService_FollowUser(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	mockRepo := NewMockRepository(ctrl)
	userRepo := user.NewMockRepository(ctrl)
	service := NewService(mockRepo, userRepo)

	followerID := "follower-123"
	followingID := "following-456"

	tests := []struct {
		name          string
		followerID    string
		followingID   string
		mockBehavior  func()
		wantErr       bool
		expectedError string
	}{
		{
			name:        "Success follow user",
			followerID:  followerID,
			followingID: followingID,
			mockBehavior: func() {
				mockRepo.EXPECT().ListByIDs(gomock.Any(), []string{followerID, followingID}).Return([]*domain.User{
					{ID: new(followerID)},
					{ID: new(followingID)},
				}, nil).Times(1)
				mockRepo.EXPECT().FollowUser(gomock.Any(), gomock.Any()).Return(nil).Times(1)
			},
			wantErr: false,
		},
		{
			name:        "Error: Following user not found",
			followerID:  followerID,
			followingID: followingID,
			mockBehavior: func() {
				mockRepo.EXPECT().ListByIDs(gomock.Any(), []string{followerID, followingID}).Return([]*domain.User{
					{ID: new(followerID)},
				}, nil).Times(1)
			},
			wantErr:       true,
			expectedError: "usuario para seguir não encontrado",
		},
		{
			name:        "Error: Follower not found",
			followerID:  followerID,
			followingID: followingID,
			mockBehavior: func() {
				mockRepo.EXPECT().ListByIDs(gomock.Any(), []string{followerID, followingID}).Return([]*domain.User{
					{ID: new(followingID)},
				}, nil).Times(1)
			},
			wantErr:       true,
			expectedError: "seguidor não encontrado",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.mockBehavior()
			err := service.FollowUser(context.Background(), tt.followerID, tt.followingID)

			if tt.wantErr {
				assert.Error(t, err)
				assert.Equal(t, tt.expectedError, err.Error())
			} else {
				assert.NoError(t, err)
			}
		})
	}
}

func TestUserService_UnfollowUser(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	mockRepo := NewMockRepository(ctrl)
	userRepo := user.NewMockRepository(ctrl)
	service := NewService(mockRepo, userRepo)

	mockRepo.EXPECT().UnfollowUser(gomock.Any(), "follower-123", "following-456").Return(nil).Times(1)

	err := service.UnfollowUser(context.Background(), "follower-123", "following-456")
	assert.NoError(t, err)
}
