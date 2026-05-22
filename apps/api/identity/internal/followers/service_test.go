package followers

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

func (m *MockRepository) ListByIDs(ctx context.Context, ids []string) ([]*domain.User, error) {
	args := m.Called(ctx, ids)
	users, _ := args.Get(0).([]*domain.User)
	return users, args.Error(1)
}

func (m *MockRepository) FollowUser(ctx context.Context, f domain.UserFollows) error {
	args := m.Called(ctx, f)
	return args.Error(0)
}

func (m *MockRepository) UnfollowUser(ctx context.Context, followerId, followingId string) error {
	args := m.Called(ctx, followerId, followingId)
	return args.Error(0)
}

func (m *MockRepository) CountFollowers(ctx context.Context, userId string) (int, error) {
	args := m.Called(ctx, userId)
	return args.Int(0), args.Error(1)
}

func (m *MockRepository) CountFollowing(ctx context.Context, userId string) (int, error) {
	args := m.Called(ctx, userId)
	return args.Int(0), args.Error(1)
}

func (m *MockRepository) ListFollower(ctx context.Context, id string, cursor *utils.CursorData, limit int) ([]*domain.User, *utils.CursorData, error) {
	args := m.Called(ctx, id, cursor, limit)
	users, _ := args.Get(0).([]*domain.User)
	nextCursor, _ := args.Get(1).(*utils.CursorData)
	return users, nextCursor, args.Error(2)
}

func (m *MockRepository) ListFollowing(ctx context.Context, id string, cursor *utils.CursorData, limit int) ([]*domain.User, *utils.CursorData, error) {
	args := m.Called(ctx, id, cursor, limit)
	users, _ := args.Get(0).([]*domain.User)
	nextCursor, _ := args.Get(1).(*utils.CursorData)
	return users, nextCursor, args.Error(2)
}

func newString(s string) *string {
	return &s
}

func TestUserService_FollowUser(t *testing.T) {
	mockRepo := new(MockRepository)
	service := NewService(mockRepo)

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
				mockRepo.On("ListByIDs", mock.Anything, []string{followerID, followingID}).Return([]*domain.User{
					{ID: newString(followerID)},
					{ID: newString(followingID)},
				}, nil).Once()
				mockRepo.On("FollowUser", mock.Anything, mock.MatchedBy(func(f domain.UserFollows) bool {
					return *f.FollowerId == followerID && *f.FollowingId == followingID
				})).Return(nil).Once()
			},
			wantErr: false,
		},
		{
			name:        "Error: Following user not found",
			followerID:  followerID,
			followingID: followingID,
			mockBehavior: func() {
				mockRepo.On("ListByIDs", mock.Anything, []string{followerID, followingID}).Return([]*domain.User{
					{ID: newString(followerID)},
				}, nil).Once()
			},
			wantErr:       true,
			expectedError: "usuario para seguir não encontrado",
		},
		{
			name:        "Error: Follower not found",
			followerID:  followerID,
			followingID: followingID,
			mockBehavior: func() {
				mockRepo.On("ListByIDs", mock.Anything, []string{followerID, followingID}).Return([]*domain.User{
					{ID: newString(followingID)},
				}, nil).Once()
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
			mockRepo.AssertExpectations(t)
		})
	}
}

func TestUserService_UnfollowUser(t *testing.T) {
	mockRepo := new(MockRepository)
	service := NewService(mockRepo)

	mockRepo.On("UnfollowUser", mock.Anything, "follower-123", "following-456").Return(nil).Once()

	err := service.UnfollowUser(context.Background(), "follower-123", "following-456")
	assert.NoError(t, err)
	mockRepo.AssertExpectations(t)
}
