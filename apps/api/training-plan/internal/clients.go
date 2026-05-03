package internal

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
)

type IdentityService struct {
	baseURL    string
	httpClient *http.Client
}

func NewIdentityService() *IdentityService {
	return &IdentityService{
		baseURL:    "http://localhost:3334",
		httpClient: &http.Client{},
	}
}

func (s *IdentityService) ExistsUser(ctx context.Context, id string) (bool, error) {
	req, err := http.NewRequestWithContext(ctx, "GET", fmt.Sprintf("%s/users/%s", s.baseURL, id), nil)
	if err != nil {
		return false, err
	}

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

func (s *IdentityService) ListUser(ctx context.Context, ids *[]string) (map[string]*any, error) {
	if ids == nil || len(*ids) == 0 {
		return map[string]*any{}, nil
	}

	// In a real app, this would be a GET /users?ids=uuid1,uuid2...
	// For now, let's implement a placeholder that calls the service
	url := fmt.Sprintf("%s/users?ids=%s", s.baseURL, strings.Join(*ids, ","))
	req, err := http.NewRequestWithContext(ctx, "GET", url, nil)
	if err != nil {
		return nil, err
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

type StorageService interface {
	GenerateURL(ctx context.Context, path string) string
	Upload(ctx context.Context, path string, file []byte) error
	Delete(ctx context.Context, path string) error
	Copy(ctx context.Context, src, dst string) error
}

type LocalStorageService struct{}

func NewLocalStorageService() StorageService {
	return &LocalStorageService{}
}

func (s *LocalStorageService) GenerateURL(ctx context.Context, path string) string {
	return "http://localhost:3333/uploads/" + path
}

func (s *LocalStorageService) Upload(ctx context.Context, path string, file []byte) error {
	return nil
}

func (s *LocalStorageService) Delete(ctx context.Context, path string) error {
	return nil
}

func (s *LocalStorageService) Copy(ctx context.Context, src, dst string) error {
	return nil
}
