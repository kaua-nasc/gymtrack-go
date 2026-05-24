package comment

import (
	"context"
	"fmt"
	"time"

	"github.com/kaua-nasc/gymtrack-go/apps/api/social/internal/domain"
	"github.com/kaua-nasc/gymtrack-go/libs/auth"
	"github.com/kaua-nasc/gymtrack-go/libs/utils"
)

type Service struct {
	repo     Repository
	identity domain.IdentityClient
}

func NewService(repo Repository, identity domain.IdentityClient) *Service {
	return &Service{
		repo:     repo,
		identity: identity,
	}
}

func (s *Service) AddComment(ctx context.Context, comment *domain.Comment, authorId string) error {
	token, _ := ctx.Value(string(auth.TokenContextKey)).(string)

	// Validate author existence
	if _, err := s.identity.FindUser(ctx, authorId, token); err != nil {
		return fmt.Errorf("failed to verify author existence: %w", err)
	}

	id, err := utils.GenerateUUIDV7(ctx)
	if err != nil {
		return err
	}

	now := time.Now().UTC()
	comment.Id = id
	comment.AuthorId = authorId
	comment.CreatedAt = now
	comment.UpdatedAt = now
	if err := s.repo.AddComment(ctx, comment); err != nil {
		return err
	}

	return nil
}

func (s *Service) DeleteComment(ctx context.Context, commentId, userId string) error {
	comment, err := s.repo.FindCommentById(ctx, commentId)
	if err != nil {
		return err
	}
	if comment == nil {
		return fmt.Errorf("comment not found")
	}

	if comment.AuthorId != userId {
		return fmt.Errorf("you are not the author of this comment")
	}

	return s.repo.DeleteComment(ctx, commentId)
}

func (s *Service) GetComments(ctx context.Context, postId, cursor string, limit int) ([]domain.Comment, string, error) {
	var decodedCursor *utils.CursorData
	utils.DecodeCursor(cursor, &decodedCursor)

	comments, rawNextCursor, err := s.repo.GetComments(ctx, postId, decodedCursor, limit)
	if err != nil {
		return nil, "", err
	}

	token, _ := ctx.Value(string(auth.TokenContextKey)).(string)

	if len(comments) > 0 {
		authorIDsMap := make(map[string]bool)
		for _, c := range comments {
			authorIDsMap[c.AuthorId] = true
		}
		var authorIDs []string
		for id := range authorIDsMap {
			authorIDs = append(authorIDs, id)
		}

		authorsMap, err := s.identity.ListUser(ctx, authorIDs, token)
		if err == nil {
			for i := range comments {
				if author, ok := authorsMap[comments[i].AuthorId]; ok {
					comments[i].Author = author
				}
			}
		}
	}

	nextCursor, _ := utils.EncodeCursor(rawNextCursor)
	return comments, nextCursor, nil
}
