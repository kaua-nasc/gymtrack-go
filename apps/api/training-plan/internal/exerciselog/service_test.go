package exerciselog

import (
	"context"
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
)

func TestService_LogExercise(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := NewMockRepository(ctrl)
	service := NewService(mockRepo)

	notes := "Felt good"

	tests := []struct {
		name         string
		exerciseId   string
		userId       string
		reps         []int
		weight       []float64
		notes        *string
		mockBehavior func()
		wantErr      bool
	}{
		{
			name:       "Success log exercise",
			exerciseId: "ex-123",
			userId:     "user-123",
			reps:       []int{10, 10},
			weight:     []float64{50.5, 50.5},
			notes:      &notes,
			mockBehavior: func() {
				mockRepo.EXPECT().LogExercise(gomock.Any(), gomock.Any()).Return(nil)
			},
			wantErr: false,
		},
		{
			name:       "Repo error",
			exerciseId: "ex-123",
			userId:     "user-123",
			reps:       []int{12},
			weight:     []float64{60},
			notes:      nil,
			mockBehavior: func() {
				mockRepo.EXPECT().LogExercise(gomock.Any(), gomock.Any()).Return(errors.New("db error"))
			},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.mockBehavior()
			err := service.LogExercise(context.Background(), tt.exerciseId, tt.userId, tt.reps, tt.weight, tt.notes)
			if tt.wantErr {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
			}
		})
	}
}
