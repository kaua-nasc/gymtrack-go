package domain

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"strings"
)

type IdentityService struct {
	baseURL    string
	httpClient *http.Client
}

func NewIdentityService() *IdentityService {
	identityAPI := os.Getenv("IDENTITY_API_URL")

	if identityAPI == "" {
		// Default for local development if not provided
		identityAPI = "http://localhost:8080"
	}

	return &IdentityService{
		baseURL:    identityAPI,
		httpClient: &http.Client{},
	}
}

func (s *IdentityService) FindUser(ctx context.Context, id string, token string) (any, error) {
	req, err := http.NewRequestWithContext(ctx, "GET", fmt.Sprintf("%s/identity/users/%s", s.baseURL, id), nil)
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
		return nil, fmt.Errorf("user not found")
	}

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("identity service error: %d", resp.StatusCode)
	}

	var user any
	if err := json.NewDecoder(resp.Body).Decode(&user); err != nil {
		return nil, fmt.Errorf("failed to decode user: %w", err)
	}

	return user, nil
}

func (s *IdentityService) ListUser(ctx context.Context, ids []string, token string) (map[string]any, error) {
	if len(ids) == 0 {
		return map[string]any{}, nil
	}

	url := fmt.Sprintf("%s/identity/users?ids=%s", s.baseURL, strings.Join(ids, ","))
	req, err := http.NewRequestWithContext(ctx, "GET", url, nil)
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

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("identity service error: %d", resp.StatusCode)
	}

	var users []map[string]any
	if err := json.NewDecoder(resp.Body).Decode(&users); err != nil {
		return nil, err
	}

	result := make(map[string]any)
	for _, u := range users {
		id, _ := u["id"].(string)
		result[id] = u
	}

	return result, nil
}

type TrainingPlanClient struct {
	baseURL    string
	httpClient *http.Client
}

func NewTrainingPlanClient() *TrainingPlanClient {
	apiURL := os.Getenv("TRAINING_PLAN_API_URL")

	if apiURL == "" {
		apiURL = "http://localhost:8081"
	}

	return &TrainingPlanClient{
		baseURL:    apiURL,
		httpClient: &http.Client{},
	}
}

func (s *TrainingPlanClient) FindPlan(ctx context.Context, id string, token string) (any, error) {
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

func (s *TrainingPlanClient) ListPlans(ctx context.Context, ids []string, token string) (map[string]any, error) {
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
