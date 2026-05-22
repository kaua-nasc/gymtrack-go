package like

import "context"

type Service struct {
	repo *Repository
}

func NewService(repo *Repository) *Service {
	return &Service{repo: repo}
}

func (s *Service) ToggleLike(ctx context.Context, postId, userId string) error {
	return s.repo.ToggleLike(postId, userId)
}
