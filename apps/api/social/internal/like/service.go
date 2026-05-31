package like

import (
	"context"

	"github.com/kaua-nasc/gymtrack-go/libs/utils"
)

type Service struct {
	repo Repository
}

func NewService(repo Repository) *Service {
	return &Service{repo: repo}
}

func (s *Service) ToggleLike(ctx context.Context, postId, userId string) error {
	id, err := utils.GenerateUUIDV7String(ctx)
	if err != nil {
		return err
	}
	return s.repo.ToggleLike(ctx, id, postId, userId)
}
