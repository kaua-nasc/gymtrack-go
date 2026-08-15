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
				mockRepo.EXPECT().ListSubscription(gomock.Any(), "user-123", domain.ListSubscriptionFilters{}).Return([]*domain.PlanSubscription{}, nil)
				mockIdentity.EXPECT().FindUser(gomock.Any(), "user-123", gomock.Any()).Return(&domain.User{}, nil)
			},
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.mockBehavior()
			_, err := service.ListSubscription(context.Background(), tt.userId, domain.ListSubscriptionFilters{})
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

	publicPlan := &domain.TrainingPlan{Id: new("plan-123"), AuthorId: "trainer-1", Visibility: domain.Public}
	protectedPlan := &domain.TrainingPlan{Id: new("plan-456"), AuthorId: "trainer-1", Visibility: domain.Protected}
	privatePlan := &domain.TrainingPlan{Id: new("plan-789"), AuthorId: "trainer-1", Visibility: domain.Private}
	limitedPlan := &domain.TrainingPlan{Id: new("plan-101"), AuthorId: "trainer-1", Visibility: domain.Public, MaxSubscriptions: new(1)}

	tests := []struct {
		name         string
		planId       string
		userId       string
		mockBehavior func()
		wantErr      bool
		wantErrIs    error
	}{
		{
			name:   "Plan not found",
			planId: "plan-missing",
			userId: "user-123",
			mockBehavior: func() {
				mockRepo.EXPECT().FindPlanForSubscription(gomock.Any(), "plan-missing").Return(nil, nil)
			},
			wantErr:   true,
			wantErrIs: domain.ErrPlanNotFound,
		},
		{
			name:   "Cannot subscribe to own plan",
			planId: "plan-123",
			userId: "trainer-1",
			mockBehavior: func() {
				mockRepo.EXPECT().FindPlanForSubscription(gomock.Any(), "plan-123").Return(publicPlan, nil)
			},
			wantErr:   true,
			wantErrIs: domain.ErrCannotSubscribeOwnPlan,
		},
		{
			name:   "Already subscribed",
			planId: "plan-123",
			userId: "user-123",
			mockBehavior: func() {
				mockRepo.EXPECT().FindPlanForSubscription(gomock.Any(), "plan-123").Return(publicPlan, nil)
				mockRepo.EXPECT().GetSubscriptionEligibility(gomock.Any(), "plan-123", "user-123").Return(true, true, nil)
			},
			wantErr:   true,
			wantErrIs: domain.ErrAlreadySubscribed,
		},
		{
			name:   "Plan incomplete",
			planId: "plan-123",
			userId: "user-123",
			mockBehavior: func() {
				mockRepo.EXPECT().FindPlanForSubscription(gomock.Any(), "plan-123").Return(publicPlan, nil)
				mockRepo.EXPECT().GetSubscriptionEligibility(gomock.Any(), "plan-123", "user-123").Return(false, false, nil)
			},
			wantErr:   true,
			wantErrIs: domain.ErrPlanIncomplete,
		},
		{
			name:   "Forbidden on private plan",
			planId: "plan-789",
			userId: "user-123",
			mockBehavior: func() {
				mockRepo.EXPECT().FindPlanForSubscription(gomock.Any(), "plan-789").Return(privatePlan, nil)
				mockRepo.EXPECT().GetSubscriptionEligibility(gomock.Any(), "plan-789", "user-123").Return(false, true, nil)
			},
			wantErr:   true,
			wantErrIs: domain.ErrSubscriptionForbidden,
		},
		{
			name:   "Forbidden on protected plan for non-student",
			planId: "plan-456",
			userId: "user-123",
			mockBehavior: func() {
				mockRepo.EXPECT().FindPlanForSubscription(gomock.Any(), "plan-456").Return(protectedPlan, nil)
				mockRepo.EXPECT().GetSubscriptionEligibility(gomock.Any(), "plan-456", "user-123").Return(false, true, nil)
				mockIdentity.EXPECT().FindUser(gomock.Any(), "user-123", gomock.Any()).Return(&domain.User{ID: "user-123"}, nil)
			},
			wantErr:   true,
			wantErrIs: domain.ErrSubscriptionForbidden,
		},
		{
			name:   "Protected plan allowed for student",
			planId: "plan-456",
			userId: "user-123",
			mockBehavior: func() {
				mockRepo.EXPECT().FindPlanForSubscription(gomock.Any(), "plan-456").Return(protectedPlan, nil)
				mockRepo.EXPECT().GetSubscriptionEligibility(gomock.Any(), "plan-456", "user-123").Return(false, true, nil)
				mockIdentity.EXPECT().FindUser(gomock.Any(), "user-123", gomock.Any()).Return(&domain.User{
					ID:        "user-123",
					StudentOf: &domain.TrainerStudentRelation{TrainerId: "trainer-1"},
				}, nil)
				mockRepo.EXPECT().CreatePlanSubscription(gomock.Any(), gomock.Any()).Return(nil)
			},
			wantErr: false,
		},
		{
			name:   "Max subscriptions reached",
			planId: "plan-101",
			userId: "user-123",
			mockBehavior: func() {
				mockRepo.EXPECT().FindPlanForSubscription(gomock.Any(), "plan-101").Return(limitedPlan, nil)
				mockRepo.EXPECT().GetSubscriptionEligibility(gomock.Any(), "plan-101", "user-123").Return(false, true, nil)
				mockRepo.EXPECT().CountActiveSubscriptionsByPlan(gomock.Any(), "plan-101").Return(1, nil)
			},
			wantErr:   true,
			wantErrIs: domain.ErrMaxSubscriptionsReached,
		},
		{
			name:   "Success subscribe",
			planId: "plan-123",
			userId: "user-123",
			mockBehavior: func() {
				mockRepo.EXPECT().FindPlanForSubscription(gomock.Any(), "plan-123").Return(publicPlan, nil)
				mockRepo.EXPECT().GetSubscriptionEligibility(gomock.Any(), "plan-123", "user-123").Return(false, true, nil)
				mockRepo.EXPECT().CreatePlanSubscription(gomock.Any(), gomock.Any()).Return(nil)
			},
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.mockBehavior()
			err := service.Subscribe(context.Background(), tt.planId, tt.userId)
			if tt.wantErr {
				assert.Error(t, err)
				if tt.wantErrIs != nil {
					assert.ErrorIs(t, err, tt.wantErrIs)
				}
			} else {
				assert.NoError(t, err)
			}
		})
	}
}
