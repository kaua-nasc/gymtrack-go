package feedback

import (
	"context"
	"errors"
	"testing"

	"github.com/kaua-nasc/gymtrack-go/apps/api/training-plan/internal/domain"
	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
)

func TestService_AddFeedback(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := NewMockRepository(ctrl)
	service := NewService(mockRepo)

	msg := "Great plan!"

	tests := []struct {
		name         string
		planId       string
		userId       string
		rating       float64
		message      *string
		mockBehavior func()
		wantErr      bool
	}{
		{
			name:    "Success add feedback",
			planId:  "plan-123",
			userId:  "user-123",
			rating:  5.0,
			message: &msg,
			mockBehavior: func() {
				mockRepo.EXPECT().AddFeedback(gomock.Any(), gomock.Any()).Return(nil)
			},
			wantErr: false,
		},
		{
			name:    "Repo error",
			planId:  "plan-123",
			userId:  "user-123",
			rating:  4.0,
			message: nil,
			mockBehavior: func() {
				mockRepo.EXPECT().AddFeedback(gomock.Any(), gomock.Any()).Return(errors.New("db error"))
			},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.mockBehavior()
			err := service.AddFeedback(context.Background(), tt.planId, tt.userId, tt.rating, tt.message)
			if tt.wantErr {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
			}
		})
	}
}

func TestService_ListFeedback(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := NewMockRepository(ctrl)
	service := NewService(mockRepo)

	tests := []struct {
		name         string
		planId       string
		cursor       string
		limit        int
		mockBehavior func()
		wantErr      bool
	}{
		{
			name:   "Success list feedback",
			planId: "plan-123",
			cursor: "",
			limit:  10,
			mockBehavior: func() {
				mockRepo.EXPECT().ListFeedback(gomock.Any(), "plan-123", gomock.Any(), 10).Return([]domain.TrainingPlanFeedback{}, nil, nil)
			},
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.mockBehavior()
			_, _, err := service.ListFeedback(context.Background(), tt.planId, tt.cursor, tt.limit)
			if tt.wantErr {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
			}
		})
	}
}
