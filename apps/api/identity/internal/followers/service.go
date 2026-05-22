package followers

import (
	"context"
	"fmt"
	"time"

	"github.com/kaua-nasc/gymtrack-go/apps/api/identity/internal/domain"
	"github.com/kaua-nasc/gymtrack-go/libs/utils"
)

type Service struct {
	repo Repository
}

func NewService(repo Repository) *Service {
	return &Service{repo: repo}
}

func (s *Service) ListFollowing(ctx context.Context, id, cursor string, limit int) ([]*domain.User, string, error) {
	var decodedCursor *utils.CursorData
	if err := utils.DecodeCursor(cursor, &decodedCursor); err != nil {
		// Log error if needed, but maintain current behavior of continuing with nil cursor
	}

	users, rawNextCursor, err := s.repo.ListFollowing(ctx, id, decodedCursor, limit)
	if err != nil {
		return nil, "", err
	}

	for _, u := range users {
		u.Sanitize()
	}

	nextCursorStr, _ := utils.EncodeCursor(rawNextCursor)
	return users, nextCursorStr, nil
}

func (s *Service) ListFollower(ctx context.Context, id, cursor string, limit int) ([]*domain.User, string, error) {
	var decodedCursor *utils.CursorData
	if err := utils.DecodeCursor(cursor, &decodedCursor); err != nil {
		// Log error if needed
	}

	users, rawNextCursor, err := s.repo.ListFollower(ctx, id, decodedCursor, limit)
	if err != nil {
		return nil, "", err
	}

	for _, u := range users {
		u.Sanitize()
	}

	nextCursorStr, _ := utils.EncodeCursor(rawNextCursor)
	return users, nextCursorStr, nil
}

func (s *Service) CountFollowers(ctx context.Context, id string) (int, error) {
	count, err := s.repo.CountFollowers(ctx, id)
	if err != nil {
		return 0, err
	}

	return count, nil
}

func (s *Service) CountFollowing(ctx context.Context, id string) (int, error) {
	count, err := s.repo.CountFollowing(ctx, id)
	if err != nil {
		return 0, err
	}

	return count, nil
}

func (s *Service) FollowUser(ctx context.Context, followerId, followingId string) error {
	users, err := s.repo.ListByIDs(ctx, []string{followerId, followingId})
	if err != nil {
		return err
	}

	var follower, following *domain.User
	for _, user := range users {
		if user.ID != nil {
			if *user.ID == followerId {
				follower = user
			}
			if *user.ID == followingId {
				following = user
			}
		}
	}

	if follower == nil {
		return fmt.Errorf("seguidor não encontrado")
	}

	if following == nil {
		return fmt.Errorf("usuario para seguir não encontrado")
	}

	id, err := utils.GenerateUUIDV7(ctx)
	if err != nil {
		return err
	}

	now := time.Now().UTC()
	follow := &domain.UserFollows{
		ID:          id,
		CreatedAt:   now,
		UpdatedAt:   now,
		FollowerId:  follower.ID,
		FollowingId: following.ID,
	}

	return s.repo.FollowUser(ctx, *follow)
}

func (s *Service) UnfollowUser(ctx context.Context, followerId, followingId string) error {
	return s.repo.UnfollowUser(ctx, followerId, followingId)
}
