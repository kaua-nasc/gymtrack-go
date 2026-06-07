package like

import (
	"context"
	"fmt"

	"github.com/kaua-nasc/gymtrack-go/apps/api/social/internal/domain"
	"github.com/kaua-nasc/gymtrack-go/apps/api/social/internal/post"
	"github.com/kaua-nasc/gymtrack-go/libs/utils"
)

type Service struct {
	repo     Repository
	postRepo post.Repository
}

func NewService(repo Repository, postRepo post.Repository) *Service {
	return &Service{
		repo:     repo,
		postRepo: postRepo,
	}
}

func (s *Service) ToggleLike(ctx context.Context, postId, userId string) error {
	p, err := s.postRepo.FindById(ctx, postId)
	if err != nil {
		return err
	}
	if p == nil {
		return fmt.Errorf("post not found")
	}

	if p.Status != domain.PostApproved {
		return fmt.Errorf("cannot like a post that is not approved")
	}

	id, err := utils.GenerateUUIDV7String(ctx)
	if err != nil {
		return err
	}
	return s.repo.ToggleLike(ctx, id, postId, userId)
}
