package plan

import (
	"context"
	"errors"
	"testing"

	"github.com/kaua-nasc/gymtrack-go/apps/api/training-plan/internal/domain"
	"github.com/kaua-nasc/gymtrack-go/apps/api/training-plan/internal/subscription"
	"github.com/kaua-nasc/gymtrack-go/libs/auth"
	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
)

func TestService_CreatePlan(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := NewMockRepository(ctrl)
	mockSubRepo := subscription.NewMockRepository(ctrl)
	mockIdentity := domain.NewMockIdentityClient(ctrl)
	service := NewService(mockRepo, mockSubRepo, mockIdentity)

	user := auth.AuthUser{ID: "018f7a2b-8b5e-7a4b-9e3f-1d4e5f6a7b8c", Type: auth.Client}
	plan := domain.TrainingPlan{
		Name:       "My Plan",
		AuthorId:   "018f7a2b-8b5e-7a4b-9e3f-1d4e5f6a7b8c",
		TimeInDays: 30,
		Type:       domain.Hypertrophy,
		Visibility: domain.Public,
		Level:      domain.Beginner,
	}

	tests := []struct {
		name         string
		plan         domain.TrainingPlan
		user         auth.AuthUser
		mockBehavior func()
		wantErr      bool
	}{
		{
			name: "Success create plan",
			plan: plan,
			user: user,
			mockBehavior: func() {
				mockIdentity.EXPECT().ExistsUser(gomock.Any(), "018f7a2b-8b5e-7a4b-9e3f-1d4e5f6a7b8c", gomock.Any()).Return(true, nil)
				mockRepo.EXPECT().CountByAuthor(gomock.Any(), "018f7a2b-8b5e-7a4b-9e3f-1d4e5f6a7b8c").Return(0, nil)
				mockRepo.EXPECT().Create(gomock.Any(), gomock.Any()).Return(nil)
			},
			wantErr: false,
		},
		{
			name: "User not found",
			plan: plan,
			user: user,
			mockBehavior: func() {
				mockIdentity.EXPECT().ExistsUser(gomock.Any(), "018f7a2b-8b5e-7a4b-9e3f-1d4e5f6a7b8c", gomock.Any()).Return(false, nil)
			},
			wantErr: true,
		},
		{
			name: "Client already has a plan",
			plan: plan,
			user: user,
			mockBehavior: func() {
				mockIdentity.EXPECT().ExistsUser(gomock.Any(), "018f7a2b-8b5e-7a4b-9e3f-1d4e5f6a7b8c", gomock.Any()).Return(true, nil)
				mockRepo.EXPECT().CountByAuthor(gomock.Any(), "018f7a2b-8b5e-7a4b-9e3f-1d4e5f6a7b8c").Return(1, nil)
			},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.mockBehavior()
			_, err := service.CreatePlan(context.Background(), tt.plan, tt.user)
			if tt.wantErr {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
			}
		})
	}
}

func TestService_UpdateMaxSubscriptions(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := NewMockRepository(ctrl)
	mockSubRepo := subscription.NewMockRepository(ctrl)
	mockIdentity := domain.NewMockIdentityClient(ctrl)
	service := NewService(mockRepo, mockSubRepo, mockIdentity)

	const (
		planID   = "plan-123"
		authorID = "author-1"
		otherID  = "user-other"
	)

	plan := &domain.TrainingPlan{Id: new(planID), AuthorId: authorID}

	ctxWithUser := func(userID string) context.Context {
		return context.WithValue(context.Background(), string(auth.UserContextKey), auth.AuthUser{ID: userID})
	}

	tests := []struct {
		name         string
		userID       string
		max          int
		mockBehavior func()
		wantErr      bool
		wantErrIs    error
	}{
		{
			name:   "Plan not found",
			userID: authorID,
			max:    10,
			mockBehavior: func() {
				mockRepo.EXPECT().Find(gomock.Any(), planID).Return(nil, nil)
			},
			wantErr:   true,
			wantErrIs: domain.ErrPlanNotFound,
		},
		{
			name:   "Unauthorized",
			userID: otherID,
			max:    10,
			mockBehavior: func() {
				mockRepo.EXPECT().Find(gomock.Any(), planID).Return(plan, nil)
			},
			wantErr:   true,
			wantErrIs: domain.ErrPlanAccessForbidden,
		},
		{
			name:   "Success when max above count",
			userID: authorID,
			max:    10,
			mockBehavior: func() {
				mockRepo.EXPECT().Find(gomock.Any(), planID).Return(plan, nil)
				mockSubRepo.EXPECT().CountSubscriptionsByPlan(gomock.Any(), planID).Return(2, nil)
				mockRepo.EXPECT().UpdateMaxSubscriptions(gomock.Any(), planID, 10).Return(nil)
			},
			wantErr: false,
		},
		{
			name:   "Error when count above max",
			userID: authorID,
			max:    1,
			mockBehavior: func() {
				mockRepo.EXPECT().Find(gomock.Any(), planID).Return(plan, nil)
				mockSubRepo.EXPECT().CountSubscriptionsByPlan(gomock.Any(), planID).Return(3, nil)
				mockRepo.EXPECT().UpdateMaxSubscriptions(gomock.Any(), gomock.Any(), gomock.Any()).Times(0)
			},
			wantErr:   true,
			wantErrIs: domain.ErrMaxSubscriptionsBelowCurrent,
		},
		{
			name:   "Unlimited zero allowed",
			userID: authorID,
			max:    0,
			mockBehavior: func() {
				mockRepo.EXPECT().Find(gomock.Any(), planID).Return(plan, nil)
				mockSubRepo.EXPECT().CountSubscriptionsByPlan(gomock.Any(), gomock.Any()).Times(0)
				mockRepo.EXPECT().UpdateMaxSubscriptions(gomock.Any(), planID, 0).Return(nil)
			},
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.mockBehavior()
			err := service.UpdateMaxSubscriptions(ctxWithUser(tt.userID), planID, tt.max)
			if tt.wantErr {
				assert.ErrorIs(t, err, tt.wantErrIs)
			} else {
				assert.NoError(t, err)
			}
		})
	}
}

func TestService_GetPlan(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := NewMockRepository(ctrl)
	mockSubRepo := subscription.NewMockRepository(ctrl)
	mockIdentity := domain.NewMockIdentityClient(ctrl)
	service := NewService(mockRepo, mockSubRepo, mockIdentity)

	tests := []struct {
		name         string
		id           string
		mockBehavior func()
		wantErr      bool
	}{
		{
			name: "Success find plan",
			id:   "plan-123",
			mockBehavior: func() {
				planID := "plan-123"
				mockRepo.EXPECT().FindComplete(gomock.Any(), "plan-123").Return(&domain.TrainingPlan{Id: &planID, AuthorId: "author-1", Visibility: domain.Public}, nil)
				mockRepo.EXPECT().FindSubscriptionByPlan(gomock.Any(), "plan-123", "test-user").Return(nil, nil)
				mockIdentity.EXPECT().FindUser(gomock.Any(), "author-1", gomock.Any()).Return(&domain.User{}, nil)
			},
			wantErr: false,
		},
		{
			name: "Plan not found",
			id:   "plan-not-exists",
			mockBehavior: func() {
				mockRepo.EXPECT().FindComplete(gomock.Any(), "plan-not-exists").Return(nil, errors.New("not found"))
			},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.mockBehavior()
			ctx := context.WithValue(context.Background(), string(auth.UserContextKey), auth.AuthUser{ID: "test-user"})
			_, err := service.GetPlan(ctx, tt.id)
			if tt.wantErr {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
			}
		})
	}
}
