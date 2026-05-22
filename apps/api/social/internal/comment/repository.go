package comment

import (
	"database/sql"

	"github.com/kaua-nasc/gymtrack-go/apps/api/social/internal/domain"
	"github.com/kaua-nasc/gymtrack-go/libs/utils"
)

type Repository struct {
	db *sql.DB
}

func NewRepository(db *sql.DB) *Repository {
	return &Repository{db: db}
}

func (r *Repository) AddComment(comment *domain.Comment) error {
	query := `
		INSERT INTO public.post_comments (id, "createdAt", "updatedAt", "content", "authorId", "postId")
		VALUES ($1, $2, $3, $4, $5, $6)
	`
	_, err := r.db.Exec(query, comment.Id, comment.CreatedAt, comment.UpdatedAt, comment.Content, comment.AuthorId, comment.PostId)
	return err
}

func (r *Repository) FindCommentById(id string) (*domain.Comment, error) {
	query := `SELECT id, "createdAt", "updatedAt", content, "authorId", "postId" FROM public.post_comments WHERE id = $1 AND "deletedAt" IS NULL`
	var c domain.Comment
	err := r.db.QueryRow(query, id).Scan(&c.Id, &c.CreatedAt, &c.UpdatedAt, &c.Content, &c.AuthorId, &c.PostId)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}
	return &c, nil
}

func (r *Repository) DeleteComment(id string) error {
	query := `UPDATE public.post_comments SET "deletedAt" = NOW() WHERE id = $1`
	_, err := r.db.Exec(query, id)
	return err
}

func (r *Repository) GetComments(postId string, cursor *utils.CursorData, limit int) ([]domain.Comment, *utils.CursorData, error) {
	query := `
		SELECT id, "createdAt", "updatedAt", content, "authorId", "postId"
		FROM public.post_comments
		WHERE "postId" = $1 AND "deletedAt" IS NULL
		AND ($2::timestamp IS NULL OR "createdAt" > $2)
		ORDER BY "createdAt" ASC
		LIMIT $3
	`
	var cursorTime interface{}
	if cursor != nil {
		cursorTime = cursor.CreatedAt
	}

	rows, err := r.db.Query(query, postId, cursorTime, limit)
	if err != nil {
		return nil, nil, err
	}
	defer rows.Close()

	comments := make([]domain.Comment, 0)
	for rows.Next() {
		var c domain.Comment
		err := rows.Scan(&c.Id, &c.CreatedAt, &c.UpdatedAt, &c.Content, &c.AuthorId, &c.PostId)
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
