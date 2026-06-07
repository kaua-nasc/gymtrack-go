package post

import (
	"context"
	"fmt"
	"io"
	"net/http"
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
	repo         Repository
	identity     domain.IdentityClient
	trainingPlan domain.TrainingPlanClient
}

func NewService(repo Repository, identity domain.IdentityClient, trainingPlan domain.TrainingPlanClient) *Service {
	return &Service{
		repo:         repo,
		identity:     identity,
		trainingPlan: trainingPlan,
	}
}

func (s *Service) CreatePost(ctx context.Context, post *domain.Post, authorId string) error {
	token, _ := ctx.Value(string(auth.TokenContextKey)).(string)

	g, egCtx := errgroup.WithContext(ctx)

	// Validate author existence
	g.Go(func() error {
		if _, err := s.identity.FindUser(egCtx, authorId, token); err != nil {
			return fmt.Errorf("failed to verify author existence: %w", err)
		}
		return nil
	})

	// Validate training plan existence (if applicable)
	if post.EntityType != nil && *post.EntityType == domain.TrainingPlanPost && post.EntityId != nil {
		g.Go(func() error {
			exists, err := s.trainingPlan.ExistsPublicPlan(egCtx, *post.EntityId, token)
			if err != nil {
				return fmt.Errorf("failed to verify training plan visibility: %w", err)
			}
			if !exists {
				return fmt.Errorf("training plan with id %s not found or is not public", *post.EntityId)
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
	post.Status = domain.PostPending

	if post.MediaUrls == nil {
		post.MediaUrls = []string{}
	}

	if err := s.repo.Create(ctx, post); err != nil {
		return err
	}

	return nil
}

func (s *Service) GetFeed(ctx context.Context, userId, cursor string, limit int) ([]domain.Post, string, error) {
	var decodedCursor *utils.CursorData
	utils.DecodeCursor(cursor, &decodedCursor)

	posts, rawNextCursor, err := s.repo.FindAll(ctx, userId, decodedCursor, limit)
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

		g, egCtx := errgroup.WithContext(ctx)
		var authorsMap map[string]any
		var plansMap map[string]any

		g.Go(func() error {
			var err error
			authorsMap, err = s.identity.ListUser(egCtx, authorIDs, token)
			return err
		})

		g.Go(func() error {
			var err error
			plansMap, err = s.trainingPlan.ListPlans(egCtx, plansIDs, token)
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

	for i, url := range p.MediaUrls {
		if !strings.HasPrefix(url, "http") {
			p.MediaUrls[i] = *storage.GetBlobURL(url)
		}
	}
}

func (s *Service) UpdatePost(ctx context.Context, postId, userId, content string) error {
	post, err := s.repo.FindById(ctx, postId)
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
	post.Status = domain.PostPending

	return s.repo.Update(ctx, post)
}

func (s *Service) DeletePost(ctx context.Context, postId, userId string) error {
	post, err := s.repo.FindById(ctx, postId)
	if err != nil {
		return err
	}
	if post == nil {
		return fmt.Errorf("post not found")
	}

	if post.AuthorId != userId {
		return fmt.Errorf("you are not the author of this post")
	}

	return s.repo.Delete(ctx, postId)
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
		ext := "png" // default
		if parts := strings.Split(filenames[i], "."); len(parts) > 1 {
			ext = parts[len(parts)-1]
		}
		path := fmt.Sprintf("social/posts/media/%s.%s", filenames[i], ext)

		storageFilename, err := storage.UploadBuffer(ctx, path, content)
		if err != nil {
			return nil, fmt.Errorf("failed to upload file %s: %w", filenames[i], err)
		}

		mediaUrls = append(mediaUrls, storageFilename)
	}

	return mediaUrls, nil
}

func (s *Service) GetPostsByAuthor(ctx context.Context, authorId, userId, cursor string, limit int) ([]domain.Post, string, error) {
	var decodedCursor *utils.CursorData
	utils.DecodeCursor(cursor, &decodedCursor)

	posts, rawNextCursor, err := s.repo.FindByAuthor(ctx, authorId, userId, decodedCursor, limit)
	if err != nil {
		return nil, "", err
	}

	token, _ := ctx.Value(string(auth.TokenContextKey)).(string)

	if len(posts) > 0 {
		plansIDsMap := make(map[string]bool)
		for _, p := range posts {
			if p.EntityType != nil && *p.EntityType == domain.TrainingPlanPost && p.EntityId != nil && *p.EntityId != "" {
				plansIDsMap[*p.EntityId] = true
			}
		}
		var plansIDs []string
		for id := range plansIDsMap {
			plansIDs = append(plansIDs, id)
		}

		g, egCtx := errgroup.WithContext(ctx)
		var plansMap map[string]any

		g.Go(func() error {
			var err error
			plansMap, err = s.trainingPlan.ListPlans(egCtx, plansIDs, token)
			return err
		})

		if err := g.Wait(); err != nil {
			// We might want to log the error but still return the posts even if hydration fails partially
		}

		for i := range posts {
			s.sanitizePost(&posts[i])

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

func (s *Service) GetPendingPosts(ctx context.Context, adminId, cursor string, limit int) ([]domain.Post, string, error) {
	// Verify admin role
	token, _ := ctx.Value(string(auth.TokenContextKey)).(string)
	userAny, err := s.identity.FindUser(ctx, adminId, token)
	if err != nil {
		return nil, "", fmt.Errorf("failed to verify admin: %w", err)
	}

	userMap, ok := userAny.(map[string]any)
	if !ok || userMap["type"] != string(auth.Admin) {
		return nil, "", fmt.Errorf("forbidden: only admins can view pending posts")
	}

	var decodedCursor *utils.CursorData
	utils.DecodeCursor(cursor, &decodedCursor)

	posts, rawNextCursor, err := s.repo.FindPending(ctx, decodedCursor, limit)
	if err != nil {
		return nil, "", err
	}

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

		g, egCtx := errgroup.WithContext(ctx)
		var authorsMap map[string]any
		var plansMap map[string]any

		g.Go(func() error {
			var err error
			authorsMap, err = s.identity.ListUser(egCtx, authorIDs, token)
			return err
		})

		g.Go(func() error {
			var err error
			plansMap, err = s.trainingPlan.ListPlans(egCtx, plansIDs, token)
			return err
		})

		g.Wait()

		for i := range posts {
			s.sanitizePost(&posts[i])
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

func (s *Service) UpdatePostStatus(ctx context.Context, adminId, postId string, status domain.PostStatus, reason *string) error {
	// Verify admin role
	token, _ := ctx.Value(string(auth.TokenContextKey)).(string)
	userAny, err := s.identity.FindUser(ctx, adminId, token)
	if err != nil {
		return fmt.Errorf("failed to verify admin: %w", err)
	}

	userMap, ok := userAny.(map[string]any)
	if !ok || userMap["type"] != string(auth.Admin) {
		return fmt.Errorf("forbidden: only admins can audit posts")
	}

	post, err := s.repo.FindById(ctx, postId)
	if err != nil {
		return err
	}
	if post == nil {
		return fmt.Errorf("post not found")
	}

	if post.AuthorId == adminId {
		return fmt.Errorf("admins cannot audit their own posts")
	}

	return s.repo.AuditPost(ctx, postId, status, reason, adminId)
}

func (s *Service) GetAuditHistory(ctx context.Context, adminId, statusStr, startDateStr, endDateStr, cursor string, limit int) ([]domain.AuditLog, string, error) {
	// Verify admin role
	token, _ := ctx.Value(string(auth.TokenContextKey)).(string)
	userAny, err := s.identity.FindUser(ctx, adminId, token)
	if err != nil {
		return nil, "", fmt.Errorf("failed to verify admin: %w", err)
	}
	userMap, ok := userAny.(map[string]any)
	if !ok || userMap["type"] != string(auth.Admin) {
		return nil, "", fmt.Errorf("forbidden: only admins can view audit history")
	}

	var newStatus domain.PostStatus
	if statusStr != "" {
		if statusStr != string(domain.PostApproved) && statusStr != string(domain.PostRejected) {
			return nil, "", fmt.Errorf("invalid status filter")
		}
		newStatus = domain.PostStatus(statusStr)
	}

	var startDate, endDate *time.Time
	if startDateStr != "" {
		t, err := time.Parse(time.RFC3339, startDateStr)
		if err != nil {
			return nil, "", fmt.Errorf("invalid start date format, use RFC3339")
		}
		startDate = &t
	}
	if endDateStr != "" {
		t, err := time.Parse(time.RFC3339, endDateStr)
		if err != nil {
			return nil, "", fmt.Errorf("invalid end date format, use RFC3339")
		}
		endDate = &t
	}

	var decodedCursor *AuditCursorData
	if cursor != "" {
		decodedCursor = &AuditCursorData{}
		utils.DecodeCursor(cursor, decodedCursor)
	}

	logs, rawNextCursor, err := s.repo.FindAuditLogs(ctx, newStatus, startDate, endDate, decodedCursor, limit)
	if err != nil {
		return nil, "", err
	}

	if len(logs) > 0 {
		postIDsMap := make(map[string]bool)
		adminIDsMap := make(map[string]bool)
		for _, l := range logs {
			postIDsMap[l.PostId] = true
			adminIDsMap[l.AdminId] = true
		}

		var postIDs []string
		for id := range postIDsMap {
			postIDs = append(postIDs, id)
		}
		var adminIDs []string
		for id := range adminIDsMap {
			adminIDs = append(adminIDs, id)
		}

		// Fetch posts and collect author IDs (sequential, no race condition)
		postsMap := make(map[string]*domain.Post)
		authorIDsMap := make(map[string]bool)
		for _, pid := range postIDs {
			post, err := s.repo.FindById(ctx, pid)
			if err != nil {
				continue
			}
			if post != nil {
				postsMap[pid] = post
				authorIDsMap[post.AuthorId] = true
			}
		}

		var authorIDs []string
		for id := range authorIDsMap {
			authorIDs = append(authorIDs, id)
		}

		// Fetch authors and admins in parallel
		var authorsMap map[string]any
		var adminsMap map[string]any
		g, egCtx := errgroup.WithContext(ctx)

		if len(authorIDs) > 0 {
			g.Go(func() error {
				var err error
				authorsMap, err = s.identity.ListUser(egCtx, authorIDs, token)
				return err
			})
		}
		if len(adminIDs) > 0 {
			g.Go(func() error {
				var err error
				adminsMap, err = s.identity.ListUser(egCtx, adminIDs, token)
				return err
			})
		}

		g.Wait()

		for i := range logs {
			if post, ok := postsMap[logs[i].PostId]; ok {
				s.sanitizePost(post)
				if author, ok := authorsMap[post.AuthorId]; ok {
					post.Author = author
				}
				logs[i].Post = post
			}
			if admin, ok := adminsMap[logs[i].AdminId]; ok {
				logs[i].Admin = admin
			}
		}
	}

	nextCursor, _ := utils.EncodeCursor(rawNextCursor)
	return logs, nextCursor, nil
}
