package post_test

import (
	"context"
	"errors"
	"testing"

	"github.com/kaua-nasc/gymtrack-go/apps/api/social/internal/domain"
	"github.com/kaua-nasc/gymtrack-go/apps/api/social/internal/post"
	"github.com/kaua-nasc/gymtrack-go/libs/auth"
	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
)

func ptr[T any](v T) *T {
	return &v
}

func TestService_CreatePost(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := post.NewMockRepository(ctrl)
	mockIdentity := domain.NewMockIdentityClient(ctrl)
	mockPlan := domain.NewMockTrainingPlanClient(ctrl)

	service := post.NewService(mockRepo, mockIdentity, mockPlan)

	authorId := "018f7a2b-8b5e-7a4b-9e3f-1d4e5f6a7b8c"
	planId := "018f7a2b-8b5e-7a4b-9e3f-1d4e5f6a7b8d"
	ctx := context.WithValue(context.Background(), string(auth.TokenContextKey), "valid-token")

	planPostType := domain.TrainingPlanPost

	tests := []struct {
		name         string
		post         *domain.Post
		mockBehavior func()
		wantErr      bool
	}{
		{
			name: "Success create simple post",
			post: &domain.Post{
				Content: "Hello world",
			},
			mockBehavior: func() {
				mockIdentity.EXPECT().FindUser(gomock.Any(), authorId, "valid-token").Return(map[string]any{"id": authorId}, nil)
				mockRepo.EXPECT().Create(gomock.Any(), gomock.Any()).Return(nil)
			},
			wantErr: false,
		},
		{
			name: "Success create training plan post",
			post: &domain.Post{
				Content:    "My new plan",
				EntityType: &planPostType,
				EntityId:   &planId,
			},
			mockBehavior: func() {
				mockIdentity.EXPECT().FindUser(gomock.Any(), authorId, "valid-token").Return(map[string]any{"id": authorId}, nil)
				mockPlan.EXPECT().FindPlan(gomock.Any(), planId, "valid-token").Return(map[string]any{"id": planId}, nil)
				mockRepo.EXPECT().Create(gomock.Any(), gomock.Any()).Return(nil)
			},
			wantErr: false,
		},
		{
			name: "Author not found",
			post: &domain.Post{
				Content: "Hello world",
			},
			mockBehavior: func() {
				mockIdentity.EXPECT().FindUser(gomock.Any(), authorId, "valid-token").Return(nil, errors.New("not found"))
			},
			wantErr: true,
		},
		{
			name: "Plan not found",
			post: &domain.Post{
				Content:    "My new plan",
				EntityType: &planPostType,
				EntityId:   &planId,
			},
			mockBehavior: func() {
				mockIdentity.EXPECT().FindUser(gomock.Any(), authorId, "valid-token").Return(map[string]any{"id": authorId}, nil)
				mockPlan.EXPECT().FindPlan(gomock.Any(), planId, "valid-token").Return(nil, nil)
			},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.mockBehavior()
			err := service.CreatePost(ctx, tt.post, authorId)
			if tt.wantErr {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
			}
		})
	}
}

func TestService_GetFeed(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := post.NewMockRepository(ctrl)
	mockIdentity := domain.NewMockIdentityClient(ctrl)
	mockPlan := domain.NewMockTrainingPlanClient(ctrl)

	service := post.NewService(mockRepo, mockIdentity, mockPlan)

	userId := "user-123"
	ctx := context.WithValue(context.Background(), string(auth.TokenContextKey), "valid-token")

	posts := []domain.Post{
		{
			Id:       ptr("post-1"),
			AuthorId: "author-1",
		},
	}

	tests := []struct {
		name         string
		mockBehavior func()
		wantErr      bool
	}{
		{
			name: "Success get feed and hydrate",
			mockBehavior: func() {
				mockRepo.EXPECT().FindAll(gomock.Any(), userId, nil, 10).Return(posts, nil, nil)
				mockIdentity.EXPECT().ListUser(gomock.Any(), []string{"author-1"}, "valid-token").Return(map[string]any{
					"author-1": map[string]any{"id": "author-1", "name": "Author One"},
				}, nil)
				mockPlan.EXPECT().ListPlans(gomock.Any(), gomock.Any(), "valid-token").Return(map[string]any{}, nil)
			},
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.mockBehavior()
			result, _, err := service.GetFeed(ctx, userId, "", 10)
			if tt.wantErr {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
				assert.Len(t, result, 1)
				assert.NotNil(t, result[0].Author)
			}
		})
	}
}
