package like_test

import (
	"context"
	"testing"

	"github.com/kaua-nasc/gymtrack-go/apps/api/social/internal/like"
	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
)

func TestService_ToggleLike(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := like.NewMockRepository(ctrl)
	service := like.NewService(mockRepo)

	postId := "post-123"
	userId := "user-123"

	mockRepo.EXPECT().ToggleLike(gomock.Any(), gomock.Any(), postId, userId).Return(nil)

	err := service.ToggleLike(context.Background(), postId, userId)
	assert.NoError(t, err)
}
