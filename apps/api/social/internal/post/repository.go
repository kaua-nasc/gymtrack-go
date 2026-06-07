package post

import (
	"context"
	"database/sql"
	"time"

	"github.com/kaua-nasc/gymtrack-go/apps/api/social/internal/domain"
	"github.com/kaua-nasc/gymtrack-go/libs/utils"
	"github.com/lib/pq"
)

type AuditCursorData struct {
	ID        string    `json:"id"`
	UpdatedAt time.Time `json:"updatedAt"`
}

//go:generate go run go.uber.org/mock/mockgen -source=repository.go -destination=mock_repository.go -package=post
type Repository interface {
	Create(ctx context.Context, post *domain.Post) error
	FindAll(ctx context.Context, currentUserId string, cursor *utils.CursorData, limit int) ([]domain.Post, *utils.CursorData, error)
	FindById(ctx context.Context, id string) (*domain.Post, error)
	Update(ctx context.Context, post *domain.Post) error
	Delete(ctx context.Context, id string) error
	FindByAuthor(ctx context.Context, authorId, currentUserId string, cursor *utils.CursorData, limit int) ([]domain.Post, *utils.CursorData, error)
	FindPending(ctx context.Context, cursor *utils.CursorData, limit int) ([]domain.Post, *utils.CursorData, error)
	AuditPost(ctx context.Context, id string, status domain.PostStatus, reason *string, adminId string) error
	FindAuditLogs(ctx context.Context, newStatus domain.PostStatus, startDate, endDate *time.Time, cursor *AuditCursorData, limit int) ([]domain.AuditLog, *AuditCursorData, error)
}

type PostgresRepository struct {
	db *sql.DB
}

func NewRepository(db *sql.DB) Repository {
	return &PostgresRepository{db: db}
}

func (r *PostgresRepository) Create(ctx context.Context, post *domain.Post) error {
	query := `
        INSERT INTO posts (id, "createdAt", "updatedAt", "authorId", "content", "entityId", "entityType", "mediaUrls", status)
        VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9)
    `
	_, err := r.db.ExecContext(ctx, query, *post.Id, post.CreatedAt, post.UpdatedAt, post.AuthorId, post.Content, post.EntityId, post.EntityType, pq.Array(post.MediaUrls), post.Status)
	return err
}

func (r *PostgresRepository) FindAll(ctx context.Context, currentUserId string, cursor *utils.CursorData, limit int) ([]domain.Post, *utils.CursorData, error) {
	query := `
		SELECT 
			p.id, p."createdAt", p."updatedAt", p."authorId", p."content", p."entityId", p."entityType", p."mediaUrls", p.status,
			(SELECT COUNT(*) FROM public.post_likes WHERE "postId" = p.id) as likes_count,
			(SELECT COUNT(*) FROM public.post_comments WHERE "postId" = p.id) as comments_count,
			EXISTS(SELECT 1 FROM public.post_likes WHERE "postId" = p.id AND "userId" = $1) as liked_by_me
		FROM posts p
		WHERE ("deletedAt" IS NULL) AND (status = 'APPROVED') AND ($2::timestamp IS NULL OR p."createdAt" < $2)
		ORDER BY p."createdAt" DESC
		LIMIT $3
	`
	var cursorTime interface{}
	if cursor != nil {
		cursorTime = cursor.CreatedAt
	}

	rows, err := r.db.QueryContext(ctx, query, currentUserId, cursorTime, limit)
	if err != nil {
		return nil, nil, err
	}
	defer rows.Close()

	posts := make([]domain.Post, 0)
	for rows.Next() {
		var p domain.Post
		err := rows.Scan(
			&p.Id, &p.CreatedAt, &p.UpdatedAt, &p.AuthorId, &p.Content, &p.EntityId, &p.EntityType, pq.Array(&p.MediaUrls), &p.Status,
			&p.LikesCount, &p.CommentsCount, &p.LikedByCurrentUser,
		)
		if err != nil {
			return nil, nil, err
		}
		posts = append(posts, p)
	}

	var nextCursor *utils.CursorData
	if len(posts) > 0 && len(posts) == limit {
		lastPost := posts[len(posts)-1]
		nextCursor = &utils.CursorData{
			ID:        *lastPost.Id,
			CreatedAt: lastPost.CreatedAt,
		}
	}

	return posts, nextCursor, nil
}

func (r *PostgresRepository) FindPending(ctx context.Context, cursor *utils.CursorData, limit int) ([]domain.Post, *utils.CursorData, error) {
	query := `
		SELECT 
			p.id, p."createdAt", p."updatedAt", p."authorId", p."content", p."entityId", p."entityType", p."mediaUrls", p.status,
			(SELECT COUNT(*) FROM public.post_likes WHERE "postId" = p.id) as likes_count,
			(SELECT COUNT(*) FROM public.post_comments WHERE "postId" = p.id) as comments_count,
			false as liked_by_me
		FROM posts p
		WHERE ("deletedAt" IS NULL) AND (status = 'PENDING') AND ($1::timestamp IS NULL OR p."createdAt" < $1)
		ORDER BY p."createdAt" DESC
		LIMIT $2
	`
	var cursorTime interface{}
	if cursor != nil {
		cursorTime = cursor.CreatedAt
	}

	rows, err := r.db.QueryContext(ctx, query, cursorTime, limit)
	if err != nil {
		return nil, nil, err
	}
	defer rows.Close()

	posts := make([]domain.Post, 0)
	for rows.Next() {
		var p domain.Post
		err := rows.Scan(
			&p.Id, &p.CreatedAt, &p.UpdatedAt, &p.AuthorId, &p.Content, &p.EntityId, &p.EntityType, pq.Array(&p.MediaUrls), &p.Status,
			&p.LikesCount, &p.CommentsCount, &p.LikedByCurrentUser,
		)
		if err != nil {
			return nil, nil, err
		}
		posts = append(posts, p)
	}

	var nextCursor *utils.CursorData
	if len(posts) > 0 && len(posts) == limit {
		lastPost := posts[len(posts)-1]
		nextCursor = &utils.CursorData{
			ID:        *lastPost.Id,
			CreatedAt: lastPost.CreatedAt,
		}
	}

	return posts, nextCursor, nil
}

func (r *PostgresRepository) AuditPost(ctx context.Context, id string, status domain.PostStatus, reason *string, adminId string) error {
	tx, err := r.db.BeginTx(ctx, nil)
	if err != nil {
		return err
	}
	defer tx.Rollback()

	updateQuery := `UPDATE posts SET status = $1, "updatedAt" = NOW(), "rejectedReason" = $2 WHERE id = $3`
	_, err = tx.ExecContext(ctx, updateQuery, status, reason, id)
	if err != nil {
		return err
	}

	previousStatusQuery := `SELECT status FROM posts WHERE id = $1`
	var previousStatus domain.PostStatus
	err = tx.QueryRowContext(ctx, previousStatusQuery, id).Scan(&previousStatus)
	if err != nil {
		return err
	}

	logId, err := utils.GenerateUUIDV7String(ctx)
	if err != nil {
		return err
	}
	insertQuery := `
		INSERT INTO post_audit_logs (id, "postId", "adminId", "previousStatus", "newStatus", reason, "createdAt")
		VALUES ($1, $2, $3, $4, $5, $6, NOW())
	`
	_, err = tx.ExecContext(ctx, insertQuery, logId, id, adminId, previousStatus, status, reason)
	if err != nil {
		return err
	}

	return tx.Commit()
}

func (r *PostgresRepository) FindAuditLogs(ctx context.Context, newStatus domain.PostStatus, startDate, endDate *time.Time, cursor *AuditCursorData, limit int) ([]domain.AuditLog, *AuditCursorData, error) {
	query := `
		SELECT 
			l.id, l."postId", l."adminId", l."previousStatus", l."newStatus", l.reason, l."createdAt"
		FROM post_audit_logs l
		WHERE ($1::post_status_enum IS NULL OR l."newStatus" = $1)
			AND ($2::timestamp IS NULL OR l."createdAt" >= $2)
			AND ($3::timestamp IS NULL OR l."createdAt" <= $3)
			AND ($4::timestamp IS NULL OR l."createdAt" < $4)
		ORDER BY l."createdAt" DESC
		LIMIT $5
	`
	var cursorTime interface{}
	if cursor != nil {
		cursorTime = cursor.UpdatedAt
	}

	var statusArg interface{} = newStatus
	if newStatus == "" {
		statusArg = nil
	}

	rows, err := r.db.QueryContext(ctx, query, statusArg, startDate, endDate, cursorTime, limit)
	if err != nil {
		return nil, nil, err
	}
	defer rows.Close()

	logs := make([]domain.AuditLog, 0)
	for rows.Next() {
		var l domain.AuditLog
		err := rows.Scan(
			&l.Id, &l.PostId, &l.AdminId, &l.PreviousStatus, &l.NewStatus, &l.Reason, &l.CreatedAt,
		)
		if err != nil {
			return nil, nil, err
		}
		logs = append(logs, l)
	}

	var nextCursor *AuditCursorData
	if len(logs) > 0 && len(logs) == limit {
		lastLog := logs[len(logs)-1]
		nextCursor = &AuditCursorData{
			ID:        lastLog.Id,
			UpdatedAt: lastLog.CreatedAt,
		}
	}

	return logs, nextCursor, nil
}

func (r *PostgresRepository) FindById(ctx context.Context, id string) (*domain.Post, error) {
	query := `
		SELECT id, "createdAt", "updatedAt", "authorId", "content", "entityId", "entityType", "mediaUrls", status, "rejectedReason"
		FROM posts WHERE id = $1 AND "deletedAt" IS NULL
	`
	var p domain.Post
	err := r.db.QueryRowContext(ctx, query, id).Scan(&p.Id, &p.CreatedAt, &p.UpdatedAt, &p.AuthorId, &p.Content, &p.EntityId, &p.EntityType, pq.Array(&p.MediaUrls), &p.Status, &p.RejectedReason)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}
	return &p, nil
}

func (r *PostgresRepository) Update(ctx context.Context, post *domain.Post) error {
	query := `UPDATE posts SET content = $1, "updatedAt" = $2, status = $3 WHERE id = $4`
	_, err := r.db.ExecContext(ctx, query, post.Content, post.UpdatedAt, post.Status, *post.Id)
	return err
}

func (r *PostgresRepository) Delete(ctx context.Context, id string) error {
	query := `UPDATE posts SET "deletedAt" = NOW() WHERE id = $1`
	_, err := r.db.ExecContext(ctx, query, id)
	return err
}

func (r *PostgresRepository) FindByAuthor(ctx context.Context, authorId, currentUserId string, cursor *utils.CursorData, limit int) ([]domain.Post, *utils.CursorData, error) {
	query := `
		SELECT 
			p.id, p."createdAt", p."updatedAt", p."authorId", p."content", p."entityId", p."entityType", p."mediaUrls", p.status, p."rejectedReason",
			(SELECT COUNT(*) FROM public.post_likes WHERE "postId" = p.id) as likes_count,
			(SELECT COUNT(*) FROM public.post_comments WHERE "postId" = p.id) as comments_count,
			EXISTS(SELECT 1 FROM public.post_likes WHERE "postId" = p.id AND "userId" = $1) as liked_by_me
		FROM posts p
		WHERE p."authorId" = $2 AND ("deletedAt" IS NULL) AND (status = 'APPROVED' OR $1 = $2) AND ($3::timestamp IS NULL OR p."createdAt" < $3)
		ORDER BY p."createdAt" DESC
		LIMIT $4
	`
	var cursorTime interface{}
	if cursor != nil {
		cursorTime = cursor.CreatedAt
	}

	rows, err := r.db.QueryContext(ctx, query, currentUserId, authorId, cursorTime, limit)
	if err != nil {
		return nil, nil, err
	}
	defer rows.Close()

	posts := make([]domain.Post, 0)
	for rows.Next() {
		var p domain.Post
		err := rows.Scan(
			&p.Id, &p.CreatedAt, &p.UpdatedAt, &p.AuthorId, &p.Content, &p.EntityId, &p.EntityType, pq.Array(&p.MediaUrls), &p.Status, &p.RejectedReason,
			&p.LikesCount, &p.CommentsCount, &p.LikedByCurrentUser,
		)
		if err != nil {
			return nil, nil, err
		}
		posts = append(posts, p)
	}

	var nextCursor *utils.CursorData
	if len(posts) > 0 && len(posts) == limit {
		lastPost := posts[len(posts)-1]
		nextCursor = &utils.CursorData{
			ID:        *lastPost.Id,
			CreatedAt: lastPost.CreatedAt,
		}
	}

	return posts, nextCursor, nil
}
