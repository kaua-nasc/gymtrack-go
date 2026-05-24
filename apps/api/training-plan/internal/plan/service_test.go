package plan

import (
	"context"
	"errors"
	"testing"

	"github.com/kaua-nasc/gymtrack-go/apps/api/training-plan/internal/domain"
	"github.com/kaua-nasc/gymtrack-go/libs/auth"
	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
)

func TestService_CreatePlan(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := NewMockRepository(ctrl)
	mockIdentity := domain.NewMockIdentityClient(ctrl)
	service := NewService(mockRepo, mockIdentity)

	user := auth.AuthUser{ID: "018f7a2b-8b5e-7a4b-9e3f-1d4e5f6a7b8c"}
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

func TestService_GetPlan(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := NewMockRepository(ctrl)
	mockIdentity := domain.NewMockIdentityClient(ctrl)
	service := NewService(mockRepo, mockIdentity)

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
				mockRepo.EXPECT().FindComplete(gomock.Any(), "plan-123").Return(&domain.TrainingPlan{AuthorId: "author-1"}, nil)
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
			_, err := service.GetPlan(context.Background(), tt.id)
			if tt.wantErr {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
			}
		})
	}
}
