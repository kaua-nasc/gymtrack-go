package feedback

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/kaua-nasc/gymtrack-go/apps/api/training-plan/internal/domain"
)

type Repository interface {
	AddFeedback(ctx context.Context, f *domain.TrainingPlanFeedback) error
}

type PostgresRepository struct {
	db *sql.DB
}

func NewRepository(database *sql.DB) Repository {
	return &PostgresRepository{
		db: database,
	}
}

func (r *PostgresRepository) AddFeedback(ctx context.Context, f *domain.TrainingPlanFeedback) error {
	query := `INSERT INTO training_plan_feedbacks (id, "trainingPlanId", "userId", rating, message, "createdAt", "updatedAt") VALUES ($1, $2, $3, $4, $5, $6, $7)`
	_, err := r.db.ExecContext(ctx, query, f.Id, f.TrainingPlanId, f.UserId, f.Rating, f.Message, f.CreatedAt, f.UpdatedAt)
	if err != nil {
		return fmt.Errorf("could not add feedback: %w", err)
	}
	return nil
}
