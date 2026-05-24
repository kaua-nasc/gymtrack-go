package like

import (
	"context"
	"database/sql"
)

//go:generate go run go.uber.org/mock/mockgen -source=repository.go -destination=mock_repository.go -package=like
type Repository interface {
	ToggleLike(ctx context.Context, postId, userId string) error
}

type PostgresRepository struct {
	db *sql.DB
}

func NewRepository(db *sql.DB) Repository {
	return &PostgresRepository{db: db}
}

func (r *PostgresRepository) ToggleLike(ctx context.Context, postId, userId string) error {
	query := `
		DO $$
		BEGIN
			IF EXISTS (SELECT 1 FROM post_likes WHERE "postId" = $1 AND "userId" = $2) THEN
				DELETE FROM post_likes WHERE "postId" = $1 AND "userId" = $2;
			ELSE
				INSERT INTO post_likes ("postId", "userId") VALUES ($1, $2);
			END IF;
		END $$;
	`
	_, err := r.db.ExecContext(ctx, query, postId, userId)
	return err
}
