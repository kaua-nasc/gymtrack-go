package post

import (
	"context"
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/kaua-nasc/gymtrack-go/apps/api/social/internal/domain"
	"github.com/kaua-nasc/gymtrack-go/libs/auth"
	"github.com/kaua-nasc/gymtrack-go/libs/storage"
	"github.com/kaua-nasc/gymtrack-go/libs/utils"
	"golang.org/x/sync/errgroup"
)

const (
	MaxMediaFiles = 5
	MaxImageSize  = 5 * 1024 * 1024  // 5MB
	MaxVideoSize  = 20 * 1024 * 1024 // 20MB
	MaxGifSize    = 5 * 1024 * 1024  // 5MB
)

type Service struct {
	repo         *Repository
	identity     *domain.IdentityService
	trainingPlan *domain.TrainingPlanClient
}

func NewService(repo *Repository, identity *domain.IdentityService, trainingPlan *domain.TrainingPlanClient) *Service {
	return &Service{
		repo:         repo,
		identity:     identity,
		trainingPlan: trainingPlan,
	}
}

func (s *Service) CreatePost(ctx context.Context, post *domain.Post, authorId string) error {
	token, _ := ctx.Value(string(auth.TokenContextKey)).(string)

	g, ctx := errgroup.WithContext(ctx)

	// Validate author existence
	g.Go(func() error {
		if _, err := s.identity.FindUser(ctx, authorId, token); err != nil {
			return fmt.Errorf("failed to verify author existence: %w", err)
		}
		return nil
	})

	// Validate training plan existence (if applicable)
	if post.EntityType != nil && *post.EntityType == domain.TrainingPlanPost && post.EntityId != nil {
		g.Go(func() error {
			plan, err := s.trainingPlan.FindPlan(ctx, *post.EntityId, token)
			if err != nil {
				return fmt.Errorf("failed to verify training plan existence: %w", err)
			}
			if plan == nil {
				return fmt.Errorf("training plan with id %s not found", *post.EntityId)
			}
			return nil
		})
	}

	if err := g.Wait(); err != nil {
		return err
	}

	id, err := utils.GenerateUUIDV7(ctx)
	if err != nil {
		return err
	}

	now := time.Now().UTC()
	post.Id = id
	post.AuthorId = authorId
	post.CreatedAt = now
	post.UpdatedAt = now

	if post.MediaUrls == nil {
		post.MediaUrls = []string{}
	}

	if err := s.repo.Create(post); err != nil {
		return err
	}

	return nil
}

func (s *Service) GetFeed(ctx context.Context, userId, cursor string, limit int) ([]domain.Post, string, error) {
	var decodedCursor *utils.CursorData
	utils.DecodeCursor(cursor, &decodedCursor)

	posts, rawNextCursor, err := s.repo.FindAll(userId, decodedCursor, limit)
	if err != nil {
		return nil, "", err
	}

	token, _ := ctx.Value(string(auth.TokenContextKey)).(string)

	if len(posts) > 0 {
		authorIDsMap := make(map[string]bool)
		plansIDsMap := make(map[string]bool)
		for _, p := range posts {
			authorIDsMap[p.AuthorId] = true
			if p.EntityType != nil && *p.EntityType == domain.TrainingPlanPost && p.EntityId != nil && *p.EntityId != "" {
				plansIDsMap[*p.EntityId] = true
			}
		}
		var authorIDs []string
		for id := range authorIDsMap {
			authorIDs = append(authorIDs, id)
		}
		var plansIDs []string
		for id := range plansIDsMap {
			plansIDs = append(plansIDs, id)
		}

		g, ctx := errgroup.WithContext(ctx)
		var authorsMap map[string]any
		var plansMap map[string]any

		g.Go(func() error {
			var err error
			authorsMap, err = s.identity.ListUser(ctx, authorIDs, token)
			return err
		})

		g.Go(func() error {
			var err error
			plansMap, err = s.trainingPlan.ListPlans(ctx, plansIDs, token)
			return err
		})

		if err := g.Wait(); err != nil {
			// We might want to log the error but still return the posts even if hydration fails partially
			// For now, let's just proceed if we have some data
		}

		for i := range posts {
			s.sanitizePost(&posts[i])
			// Hydrate Author
			if authorsMap != nil {
				if author, ok := authorsMap[posts[i].AuthorId]; ok {
					posts[i].Author = author
				}
			}

			if plansMap != nil && posts[i].EntityId != nil {
				if plan, ok := plansMap[*posts[i].EntityId]; ok {
					posts[i].Entity = plan
				}
			}
		}
	}

	nextCursor, _ := utils.EncodeCursor(rawNextCursor)
	return posts, nextCursor, nil
}

func (s *Service) sanitizePost(p *domain.Post) {
	if p == nil {
		return
	}

	if p.MediaUrls == nil {
		p.MediaUrls = []string{}
	}

	storageURL := os.Getenv("AZURE_STORAGE_URL")
	if storageURL == "" {
		return
	}

	for i, url := range p.MediaUrls {
		if !strings.HasPrefix(url, "http") {
			p.MediaUrls[i] = storageURL + "/" + url
		}
	}
}

func (s *Service) UpdatePost(ctx context.Context, postId, userId, content string) error {
	post, err := s.repo.FindById(postId)
	if err != nil {
		return err
	}
	if post == nil {
		return fmt.Errorf("post not found")
	}

	if post.AuthorId != userId {
		return fmt.Errorf("you are not the author of this post")
	}

	post.Content = content
	post.UpdatedAt = time.Now().UTC()

	return s.repo.Update(post)
}

func (s *Service) DeletePost(ctx context.Context, postId, userId string) error {
	post, err := s.repo.FindById(postId)
	if err != nil {
		return err
	}
	if post == nil {
		return fmt.Errorf("post not found")
	}

	if post.AuthorId != userId {
		return fmt.Errorf("you are not the author of this post")
	}

	return s.repo.Delete(postId)
}

func (s *Service) UploadMedia(ctx context.Context, authorId string, files []io.Reader, filenames []string) ([]string, error) {
	if len(files) > MaxMediaFiles {
		return nil, fmt.Errorf("maximum of %d files allowed", MaxMediaFiles)
	}

	var mediaUrls []string
	for i, file := range files {
		content, err := io.ReadAll(file)
		if err != nil {
			return nil, fmt.Errorf("failed to read file %s: %w", filenames[i], err)
		}

		// Detect Content-Type
		mimeType := http.DetectContentType(content)
		var maxSize int64

		if strings.HasPrefix(mimeType, "image/") {
			if mimeType == "image/gif" {
				maxSize = MaxGifSize
			} else {
				maxSize = MaxImageSize
			}
		} else if strings.HasPrefix(mimeType, "video/") {
			maxSize = MaxVideoSize
		} else {
			return nil, fmt.Errorf("file type %s not supported", mimeType)
		}

		if int64(len(content)) > maxSize {
			return nil, fmt.Errorf("file %s exceeds size limit for its type", filenames[i])
		}

		// Generate unique filename
		timestamp := time.Now().UnixNano()
		ext := "png" // default
		if parts := strings.Split(filenames[i], "."); len(parts) > 1 {
			ext = parts[len(parts)-1]
		}
		storageFilename := fmt.Sprintf("social/posts/media/%s-%d-%d.%s", authorId, timestamp, i, ext)

		if err := storage.UploadBuffer(ctx, storageFilename, content); err != nil {
			return nil, fmt.Errorf("failed to upload file %s: %w", filenames[i], err)
		}

		mediaUrls = append(mediaUrls, storageFilename)
	}

	return mediaUrls, nil
}
