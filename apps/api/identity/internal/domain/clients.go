package domain

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
)

//go:generate go run go.uber.org/mock/mockgen -source=clients.go -destination=mock_clients.go -package=domain
type TrainingPlanClient interface {
	GetEngagementSummary(ctx context.Context, studentId string, token string) (*EngagementSummary, error)
}

type TrainingPlanService struct {
	baseURL    string
	httpClient *http.Client
}

func NewTrainingPlanService() TrainingPlanClient {
	trainingPlanAPI := os.Getenv("TRAINING_PLAN_API_URL")

	if trainingPlanAPI == "" {
		panic("TRAINING_PLAN_API_URL environment variable cannot be empty")
	}

	return &TrainingPlanService{
		baseURL:    trainingPlanAPI,
		httpClient: &http.Client{},
	}
}

func (s *TrainingPlanService) GetEngagementSummary(ctx context.Context, studentId string, token string) (*EngagementSummary, error) {
	url := fmt.Sprintf("%s/training-plans/subscriptions/engagement/%s", s.baseURL, studentId)
	req, err := http.NewRequestWithContext(ctx, "GET", url, nil)
	if err != nil {
		return nil, err
	}

	req.Header.Set("Authorization", "Bearer "+token)

	resp, err := s.httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("training-plan service error: %d", resp.StatusCode)
	}

	var summary EngagementSummary
	if err := json.NewDecoder(resp.Body).Decode(&summary); err != nil {
		return nil, fmt.Errorf("failed to decode engagement summary: %w", err)
	}

	return &summary, nil
}
