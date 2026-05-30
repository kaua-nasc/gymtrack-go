package subscription

import (
	"context"
	"testing"

	"github.com/kaua-nasc/gymtrack-go/apps/api/training-plan/internal/domain"
	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
)

func TestService_ListSubscription(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := NewMockRepository(ctrl)
	mockIdentity := domain.NewMockIdentityClient(ctrl)
	service := NewService(mockRepo, mockIdentity)

	tests := []struct {
		name         string
		userId       string
		mockBehavior func()
		wantErr      bool
	}{
		{
			name:   "Success list subscriptions",
			userId: "user-123",
			mockBehavior: func() {
				mockRepo.EXPECT().ListSubscription(gomock.Any(), "user-123").Return([]*domain.PlanSubscription{}, nil)
				mockIdentity.EXPECT().FindUser(gomock.Any(), "user-123", gomock.Any()).Return(&domain.User{}, nil)
			},
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.mockBehavior()
			_, err := service.ListSubscription(context.Background(), tt.userId)
			if tt.wantErr {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
			}
		})
	}
}

func TestService_Subscribe(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := NewMockRepository(ctrl)
	mockIdentity := domain.NewMockIdentityClient(ctrl)
	service := NewService(mockRepo, mockIdentity)

	tests := []struct {
		name         string
		planId       string
		userId       string
		mockBehavior func()
		wantErr      bool
	}{
		{
			name:   "Success subscribe",
			planId: "plan-123",
			userId: "user-123",
			mockBehavior: func() {
				mockRepo.EXPECT().GetSubscriptionEligibility(gomock.Any(), "plan-123", "user-123").Return(false, true, nil)
				mockRepo.EXPECT().CreatePlanSubscription(gomock.Any(), gomock.Any()).Return(nil)
			},
			wantErr: false,
		},
		{
			name:   "Already subscribed",
			planId: "plan-123",
			userId: "user-123",
			mockBehavior: func() {
				mockRepo.EXPECT().GetSubscriptionEligibility(gomock.Any(), "plan-123", "user-123").Return(true, true, nil)
			},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.mockBehavior()
			err := service.Subscribe(context.Background(), tt.planId, tt.userId, domain.TotalAccessSubscription)
			if tt.wantErr {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
			}
		})
	}
}
