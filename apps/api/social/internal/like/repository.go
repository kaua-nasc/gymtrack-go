package like

import (
	"context"
	"database/sql"
)

//go:generate go run go.uber.org/mock/mockgen -source=repository.go -destination=mock_repository.go -package=like
type Repository interface {
	ToggleLike(ctx context.Context, id, postId, userId string) error
}

type PostgresRepository struct {
	db *sql.DB
}

func NewRepository(db *sql.DB) Repository {
	return &PostgresRepository{db: db}
}

func (r *PostgresRepository) ToggleLike(ctx context.Context, id, postId, userId string) error {
	query := `
		WITH deleted AS (
			DELETE FROM post_likes 
			WHERE "postId" = $2 AND "userId" = $3 
			RETURNING *
		)
		INSERT INTO post_likes (id, "postId", "userId")
		SELECT $1, $2, $3
		WHERE NOT EXISTS (SELECT 1 FROM deleted);
	`
	_, err := r.db.ExecContext(ctx, query, id, postId, userId)
	return err
}
