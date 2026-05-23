package feedback

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/kaua-nasc/gymtrack-go/apps/api/training-plan/internal/domain"
	"github.com/kaua-nasc/gymtrack-go/libs/utils"
)

type Repository interface {
	AddFeedback(ctx context.Context, f *domain.TrainingPlanFeedback) error
	ListFeedback(ctx context.Context, planId string, cursor *utils.CursorData, limit int) ([]domain.TrainingPlanFeedback, *utils.CursorData, error)
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

func (r *PostgresRepository) ListFeedback(ctx context.Context, planId string, cursor *utils.CursorData, limit int) ([]domain.TrainingPlanFeedback, *utils.CursorData, error) {
	query := `SELECT id, "trainingPlanId", '' AS "userId", rating, message, "createdAt", "updatedAt" FROM training_plan_feedbacks WHERE "trainingPlanId" = $1 AND "deletedAt" IS NULL`

	var args []interface{}
	args = append(args, planId)

	if cursor != nil {
		argCount := len(args)
		query += fmt.Sprintf(` AND ("createdAt", id) < ($%d, $%d)`, argCount+1, argCount+2)
		args = append(args, cursor.CreatedAt, cursor.ID)
	}

	query += fmt.Sprintf(` ORDER BY "createdAt" DESC, id DESC LIMIT $%d`, len(args)+1)
	args = append(args, limit+1)

	rows, err := r.db.QueryContext(ctx, query, args...)
	if err != nil {
		return nil, nil, err
	}
	defer rows.Close()

	feedbacks := make([]domain.TrainingPlanFeedback, 0)
	for rows.Next() {
		var f domain.TrainingPlanFeedback
		err := rows.Scan(&f.Id, &f.TrainingPlanId, &f.UserId, &f.Rating, &f.Message, &f.CreatedAt, &f.UpdatedAt)
		if err != nil {
			return nil, nil, err
		}
		feedbacks = append(feedbacks, f)
	}

	var nextCursor *utils.CursorData
	if len(feedbacks) > limit {
		lastItem := feedbacks[limit-1]
		nextCursor = &utils.CursorData{
			ID:        lastItem.Id,
			CreatedAt: lastItem.CreatedAt,
		}
		feedbacks = feedbacks[:limit]
	}

	return feedbacks, nextCursor, nil
}
