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
		panic("IDENTITY_API_URL environment variable cannot be empty")
	}

	return &IdentityService{
		baseURL:    identityAPI,
		httpClient: &http.Client{},
	}
}

func (s *IdentityService) ExistsUser(ctx context.Context, id string, token string) (bool, error) {
	req, err := http.NewRequestWithContext(ctx, "GET", fmt.Sprintf("%s/identity/users/%s", s.baseURL, id), nil)
	if err != nil {
		return false, err
	}

	req.Header.Set("Authorization", "Bearer "+token)

	resp, err := s.httpClient.Do(req)
	if err != nil {
		return false, err
	}
	defer resp.Body.Close()

	if resp.StatusCode == http.StatusNotFound {
		return false, nil
	}

	if resp.StatusCode != http.StatusOK {
		return false, fmt.Errorf("identity service error: %d", resp.StatusCode)
	}

	return true, nil
}

type UserType string

const (
	Trainer UserType = "PERSONAL_TRAINER"
	Client  UserType = "CLIENT"
)

type User struct {
	ID        string                   `json:"id" validate:"uuid"`
	Type      UserType                 `json:"type"`
	StudentOf *TrainerStudentRelation  `json:"studentOf,omitempty"`
	TrainerOf []TrainerStudentRelation `json:"trainerOf,omitempty"`
}

type TrainerStudentRelation struct {
	ID        string `json:"id"`
	TrainerId string `json:"trainerId"`
	StudentId string `json:"studentId"`
	Trainer   *User  `json:"trainer,omitempty"`
	Student   *User  `json:"student,omitempty"`
}

func (s *IdentityService) FindUser(ctx context.Context, id string, token string) (*User, error) {
	req, err := http.NewRequestWithContext(ctx, "GET", fmt.Sprintf("%s/identity/users/%s", s.baseURL, id), nil)
	if err != nil {
		return nil, err
	}

	req.Header.Set("Authorization", "Bearer "+token)

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

	var user User
	if err := json.NewDecoder(resp.Body).Decode(&user); err != nil {
		return nil, fmt.Errorf("failed to decode user: %w", err)
	}

	return &user, nil
}

func (s *IdentityService) ListUser(ctx context.Context, ids *[]string, token string) (map[string]*any, error) {
	if ids == nil || len(*ids) == 0 {
		return map[string]*any{}, nil
	}

	url := fmt.Sprintf("%s/identity/users?ids=%s", s.baseURL, strings.Join(*ids, ","))
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
		return nil, fmt.Errorf("identity service error: %d", resp.StatusCode)
	}

	var users []map[string]any
	if err := json.NewDecoder(resp.Body).Decode(&users); err != nil {
		return nil, err
	}

	result := make(map[string]*any)
	for _, u := range users {
		id, _ := u["id"].(string)
		var val any = u
		result[id] = &val
	}

	return result, nil
}

func (s *IdentityService) GetAuthorIdFromPlan(ctx context.Context, id string) (bool, error) {
	return false, nil
}

type UploadFile struct {
	Data     []byte
	Filename string
}
