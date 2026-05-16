package internal

import (
	"context"
	"fmt"
	"log/slog"
	"time"

	"github.com/google/uuid"
	"github.com/kaua-nasc/gymtrack-go/libs/auth"
)

type PostService struct {
	repo         *PostRepository
	identity     *IdentityService
	trainingPlan *TrainingPlanClient
}

func NewPostService(repo *PostRepository, identity *IdentityService, trainingPlan *TrainingPlanClient) *PostService {
	return &PostService{
		repo:         repo,
		identity:     identity,
		trainingPlan: trainingPlan,
	}
}

func (s *PostService) CreatePost(ctx context.Context, post *Post, authorId string) error {
	token, _ := ctx.Value(string(auth.TokenContextKey)).(string)

	// Validate author existence
	if _, err := s.identity.FindUser(ctx, authorId, token); err != nil {
		return fmt.Errorf("failed to verify author existence: %w", err)
	}

	if post.EntityType == TrainingPlanPost {
		plan, err := s.trainingPlan.FindPlan(ctx, post.EntityId, token)
		if err != nil {
			return fmt.Errorf("failed to verify training plan existence: %w", err)
		}
		if plan == nil {
			return fmt.Errorf("training plan with id %s not found", post.EntityId)
		}
	}

	id, err := s.generateUuid(ctx)
	if err != nil {
		return err
	}

	now := time.Now().UTC()
	post.Id = id
	post.AuthorId = authorId
	post.CreatedAt = now
	post.UpdatedAt = now

	return s.repo.Create(post)
}

func (s *PostService) GetFeed(ctx context.Context, userId string) ([]Post, error) {
	posts, err := s.repo.FindAll(userId)
	if err != nil {
		return nil, err
	}

	token, _ := ctx.Value(string(auth.TokenContextKey)).(string)

	if len(posts) > 0 {
		authorIDsMap := make(map[string]bool)
		for _, p := range posts {
			authorIDsMap[p.AuthorId] = true
		}
		var authorIDs []string
		for id := range authorIDsMap {
			authorIDs = append(authorIDs, id)
		}

		authorsMap, err := s.identity.ListUser(ctx, authorIDs, token)

		for i := range posts {
			// Hydrate Author
			if err == nil && authorsMap != nil {
				if author, ok := authorsMap[posts[i].AuthorId]; ok {
					posts[i].Author = author
				}
			}

			// Hydrate TrainingPlan if entityType is TRAINING_PLAN
			if posts[i].EntityType == TrainingPlanPost {
				plan, err := s.trainingPlan.FindPlan(ctx, posts[i].EntityId, token)
				if err == nil && plan != nil {
					posts[i].TrainingPlan = plan
				}
			}
		}
	}

	return posts, nil
}

func (s *PostService) ToggleLike(ctx context.Context, postId, userId string) error {
	return s.repo.ToggleLike(postId, userId)
}

func (s *PostService) AddComment(ctx context.Context, comment *Comment, authorId string) error {
	token, _ := ctx.Value(string(auth.TokenContextKey)).(string)

	// Validate author existence
	if _, err := s.identity.FindUser(ctx, authorId, token); err != nil {
		return fmt.Errorf("failed to verify author existence: %w", err)
	}

	id, err := s.generateUuid(ctx)
	if err != nil {
		return err
	}

	now := time.Now().UTC()
	comment.Id = id
	comment.AuthorId = authorId
	comment.CreatedAt = now
	comment.UpdatedAt = now
	return s.repo.AddComment(comment)
}

func (s *PostService) GetComments(ctx context.Context, postId string) ([]Comment, error) {
	comments, err := s.repo.GetComments(postId)
	if err != nil {
		return nil, err
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

	return comments, nil
}

func (s *PostService) generateUuid(ctx context.Context) (*string, error) {
	id, err := uuid.NewV7()
	if err != nil {
		slog.ErrorContext(ctx, "failed to generate uuid for day", slog.Any("error", err))
		return nil, fmt.Errorf("error on generate uuid")
	}

	idStr := id.String()

	return &idStr, nil
}
