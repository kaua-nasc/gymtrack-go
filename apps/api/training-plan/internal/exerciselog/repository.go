package exerciselog

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/kaua-nasc/gymtrack-go/apps/api/training-plan/internal/domain"
	"github.com/lib/pq"
)

//go:generate go run go.uber.org/mock/mockgen -source=repository.go -destination=mock_repository.go -package=exerciselog
type Repository interface {
	LogExercise(ctx context.Context, l *domain.ExerciseLog) error
}

type PostgresRepository struct {
	db *sql.DB
}

func NewRepository(database *sql.DB) Repository {
	return &PostgresRepository{
		db: database,
	}
}

func (r *PostgresRepository) LogExercise(ctx context.Context, l *domain.ExerciseLog) error {
	query := `INSERT INTO exercise_logs (id, "userId", "exerciseId", reps, weight, notes, "createdAt", "updatedAt") VALUES ($1, $2, $3, $4, $5, $6, $7, $8)`
	_, err := r.db.ExecContext(ctx, query, l.Id, l.UserId, l.ExerciseId, pq.Array(l.Reps), pq.Array(l.Weight), l.Notes, l.CreatedAt, l.UpdatedAt)
	if err != nil {
		return fmt.Errorf("could not log exercise: %w", err)
	}
	return nil
}
