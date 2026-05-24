package comment

import (
	"context"
	"database/sql"

	"github.com/kaua-nasc/gymtrack-go/apps/api/social/internal/domain"
	"github.com/kaua-nasc/gymtrack-go/libs/utils"
)

//go:generate go run go.uber.org/mock/mockgen -source=repository.go -destination=mock_repository.go -package=comment
type Repository interface {
	AddComment(ctx context.Context, comment *domain.Comment) error
	FindCommentById(ctx context.Context, id string) (*domain.Comment, error)
	DeleteComment(ctx context.Context, id string) error
	GetComments(ctx context.Context, postId string, cursor *utils.CursorData, limit int) ([]domain.Comment, *utils.CursorData, error)
}

type PostgresRepository struct {
	db *sql.DB
}

func NewRepository(db *sql.DB) Repository {
	return &PostgresRepository{db: db}
}

func (r *PostgresRepository) AddComment(ctx context.Context, comment *domain.Comment) error {
	query := `INSERT INTO post_comments (id, "createdAt", "updatedAt", "postId", "authorId", content) VALUES ($1, $2, $3, $4, $5, $6)`
	_, err := r.db.ExecContext(ctx, query, *comment.Id, comment.CreatedAt, comment.UpdatedAt, comment.PostId, comment.AuthorId, comment.Content)
	return err
}

func (r *PostgresRepository) FindCommentById(ctx context.Context, id string) (*domain.Comment, error) {
	query := `SELECT id, "createdAt", "updatedAt", "postId", "authorId", content FROM post_comments WHERE id = $1`
	var c domain.Comment
	err := r.db.QueryRowContext(ctx, query, id).Scan(&c.Id, &c.CreatedAt, &c.UpdatedAt, &c.PostId, &c.AuthorId, &c.Content)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}
	return &c, nil
}

func (r *PostgresRepository) DeleteComment(ctx context.Context, id string) error {
	query := `DELETE FROM post_comments WHERE id = $1`
	_, err := r.db.ExecContext(ctx, query, id)
	return err
}

func (r *PostgresRepository) GetComments(ctx context.Context, postId string, cursor *utils.CursorData, limit int) ([]domain.Comment, *utils.CursorData, error) {
	query := `
		SELECT id, "createdAt", "updatedAt", "postId", "authorId", content 
		FROM post_comments 
		WHERE "postId" = $1 AND ($2::timestamp IS NULL OR "createdAt" < $2)
		ORDER BY "createdAt" DESC
		LIMIT $3
	`
	var cursorTime interface{}
	if cursor != nil {
		cursorTime = cursor.CreatedAt
	}

	rows, err := r.db.QueryContext(ctx, query, postId, cursorTime, limit)
	if err != nil {
		return nil, nil, err
	}
	defer rows.Close()

	comments := make([]domain.Comment, 0)
	for rows.Next() {
		var c domain.Comment
		err := rows.Scan(&c.Id, &c.CreatedAt, &c.UpdatedAt, &c.PostId, &c.AuthorId, &c.Content)
		if err != nil {
			return nil, nil, err
		}
		comments = append(comments, c)
	}

	var nextCursor *utils.CursorData
	if len(comments) > 0 && len(comments) == limit {
		lastComment := comments[len(comments)-1]
		nextCursor = &utils.CursorData{
			ID:        *lastComment.Id,
			CreatedAt: lastComment.CreatedAt,
		}
	}

	return comments, nextCursor, nil
}
