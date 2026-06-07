package like_test

import (
	"context"
	"testing"

	"github.com/kaua-nasc/gymtrack-go/apps/api/social/internal/domain"
	"github.com/kaua-nasc/gymtrack-go/apps/api/social/internal/like"
	"github.com/kaua-nasc/gymtrack-go/apps/api/social/internal/post"
	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
)

func TestService_ToggleLike(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := like.NewMockRepository(ctrl)
	mockPostRepo := post.NewMockRepository(ctrl)
	service := like.NewService(mockRepo, mockPostRepo)

	postId := "post-123"
	userId := "user-123"

	mockPostRepo.EXPECT().FindById(gomock.Any(), postId).Return(&domain.Post{
		Id:     &postId,
		Status: domain.PostApproved,
	}, nil)
	mockRepo.EXPECT().ToggleLike(gomock.Any(), gomock.Any(), postId, userId).Return(nil)

	err := service.ToggleLike(context.Background(), postId, userId)
	assert.NoError(t, err)
}
