package domain

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"strings"
)

//go:generate go run go.uber.org/mock/mockgen -source=clients.go -destination=mock_clients.go -package=domain
type IdentityClient interface {
	FindUser(ctx context.Context, id string, token string) (any, error)
	ListUser(ctx context.Context, ids []string, token string) (map[string]any, error)
}

type TrainingPlanClient interface {
	FindPlan(ctx context.Context, id string, token string) (any, error)
	ListPlans(ctx context.Context, ids []string, token string) (map[string]any, error)
}

type trainingPlanClient struct {
	baseURL    string
	httpClient *http.Client
}

func NewTrainingPlanClient() TrainingPlanClient {
	apiURL := os.Getenv("TRAINING_PLAN_API_URL")

	if apiURL == "" {
		apiURL = "http://localhost:8081"
	}

	return &trainingPlanClient{
		baseURL:    apiURL,
		httpClient: &http.Client{},
	}
}

func (s *trainingPlanClient) FindPlan(ctx context.Context, id string, token string) (any, error) {
	req, err := http.NewRequestWithContext(ctx, "GET", fmt.Sprintf("%s/training-plans/%s", s.baseURL, id), nil)
	if err != nil {
		return nil, err
	}

	if token != "" {
		req.Header.Set("Authorization", "Bearer "+token)
	}

	resp, err := s.httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode == http.StatusNotFound {
		return nil, nil
	}

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("training-plan service error: %d", resp.StatusCode)
	}

	var plan any
	if err := json.NewDecoder(resp.Body).Decode(&plan); err != nil {
		return nil, fmt.Errorf("failed to decode training plan: %w", err)
	}

	return plan, nil
}

func (s *trainingPlanClient) ListPlans(ctx context.Context, ids []string, token string) (map[string]any, error) {
	if len(ids) == 0 {
		return map[string]any{}, nil
	}

	body, err := json.Marshal(ids)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal ids: %w", err)
	}

	url := fmt.Sprintf("%s/training-plans/by-ids", s.baseURL)
	req, err := http.NewRequestWithContext(ctx, "POST", url, strings.NewReader(string(body)))
	if err != nil {
		return nil, err
	}

	req.Header.Set("Content-Type", "application/json")
	if token != "" {
		req.Header.Set("Authorization", "Bearer "+token)
	}

	resp, err := s.httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("training-plan service error: %d", resp.StatusCode)
	}

	var plans []map[string]any
	if err := json.NewDecoder(resp.Body).Decode(&plans); err != nil {
		return nil, err
	}

	result := make(map[string]any)
	for _, p := range plans {
		id, _ := p["id"].(string)
		result[id] = p
	}

	return result, nil
}
