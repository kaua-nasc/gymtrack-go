package comment_test

import (
	"context"
	"errors"
	"testing"

	"github.com/kaua-nasc/gymtrack-go/apps/api/social/internal/comment"
	"github.com/kaua-nasc/gymtrack-go/apps/api/social/internal/domain"
	"github.com/kaua-nasc/gymtrack-go/apps/api/social/internal/post"
	"github.com/kaua-nasc/gymtrack-go/libs/auth"
	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
)

func TestService_AddComment(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := comment.NewMockRepository(ctrl)
	mockPostRepo := post.NewMockRepository(ctrl)
	mockIdentity := domain.NewMockIdentityClient(ctrl)

	service := comment.NewService(mockRepo, mockPostRepo, mockIdentity)

	authorId := "user-123"
	ctx := context.WithValue(context.Background(), string(auth.TokenContextKey), "valid-token")

	tests := []struct {
		name         string
		comment      *domain.Comment
		mockBehavior func()
		wantErr      bool
	}{
		{
			name: "Success add comment",
			comment: &domain.Comment{
				Content: "Nice post!",
				PostId:  "post-123",
			},
			mockBehavior: func() {
				mockPostRepo.EXPECT().FindById(gomock.Any(), "post-123").Return(&domain.Post{Status: domain.PostApproved}, nil)
				mockIdentity.EXPECT().FindUser(gomock.Any(), authorId, "valid-token").Return(map[string]any{"id": authorId}, nil)
				mockRepo.EXPECT().AddComment(gomock.Any(), gomock.Any()).Return(nil)
			},
			wantErr: false,
		},
		{
			name: "Post not approved",
			comment: &domain.Comment{
				Content: "Nice post!",
				PostId:  "post-123",
			},
			mockBehavior: func() {
				mockPostRepo.EXPECT().FindById(gomock.Any(), "post-123").Return(&domain.Post{Status: domain.PostPending}, nil)
			},
			wantErr: true,
		},
		{
			name: "Identity error",
			comment: &domain.Comment{
				Content: "Nice post!",
				PostId:  "post-123",
			},
			mockBehavior: func() {
				mockPostRepo.EXPECT().FindById(gomock.Any(), "post-123").Return(&domain.Post{Status: domain.PostApproved}, nil)
				mockIdentity.EXPECT().FindUser(gomock.Any(), authorId, "valid-token").Return(nil, errors.New("user not found"))
			},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.mockBehavior()
			err := service.AddComment(ctx, tt.comment, authorId)
			if tt.wantErr {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
			}
		})
	}
}

func TestService_DeleteComment(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := comment.NewMockRepository(ctrl)
	mockPostRepo := post.NewMockRepository(ctrl)
	mockIdentity := domain.NewMockIdentityClient(ctrl)

	service := comment.NewService(mockRepo, mockPostRepo, mockIdentity)

	commentId := "comm-123"
	userId := "user-123"

	tests := []struct {
		name         string
		mockBehavior func()
		wantErr      bool
	}{
		{
			name: "Success delete",
			mockBehavior: func() {
				mockRepo.EXPECT().FindCommentById(gomock.Any(), commentId).Return(&domain.Comment{Id: &commentId, AuthorId: userId}, nil)
				mockRepo.EXPECT().DeleteComment(gomock.Any(), commentId).Return(nil)
			},
			wantErr: false,
		},
		{
			name: "Unauthorized delete",
			mockBehavior: func() {
				mockRepo.EXPECT().FindCommentById(gomock.Any(), commentId).Return(&domain.Comment{Id: &commentId, AuthorId: "other-user"}, nil)
			},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.mockBehavior()
			err := service.DeleteComment(context.Background(), commentId, userId)
			if tt.wantErr {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
			}
		})
	}
}
