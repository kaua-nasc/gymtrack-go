package metrics

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	"github.com/kaua-nasc/gymtrack-go/apps/api/identity/internal/domain"
	"github.com/kaua-nasc/gymtrack-go/libs/cache"
	"github.com/kaua-nasc/gymtrack-go/libs/utils"
)

type Repository interface {
	AddBodyMeasurementNote(ctx context.Context, id, note string) error
	FindBodyMeasurement(ctx context.Context, id string) (*domain.BodyMeasurement, error)
	FindLastBodyMeasurementNote(ctx context.Context, userId string) (*domain.BodyMeasurement, error)
	ListBodyMeasurements(ctx context.Context, userId string, since *time.Time, cursor *utils.CursorData, limit int) ([]*domain.BodyMeasurement, *utils.CursorData, error)
	AddWeightLogNote(ctx context.Context, id, note string) error
	FindWeightLog(ctx context.Context, id string) (*domain.WeightLog, error)
	ListGoalsMetric(ctx context.Context, id string, since *time.Time, cursor *utils.CursorData, limit int) ([]*domain.MetricGoal, *utils.CursorData, error)
	ListWeightLogs(ctx context.Context, userId string, since *time.Time, cursor *utils.CursorData, limit int) ([]*domain.WeightLog, *utils.CursorData, error)
	AddGoalMetric(ctx context.Context, g domain.MetricGoal) error
}

type PostgresRepository struct {
	db    *sql.DB
	cache cache.Cache
}

func NewRepository(database *sql.DB, cache cache.Cache) Repository {
	return &PostgresRepository{
		db:    database,
		cache: cache,
	}
}

func (r *PostgresRepository) AddBodyMeasurementNote(ctx context.Context, id, note string) error {
	query := `UPDATE body_measurements SET "trainerNote" = $2, "trainerNoteAt" = $3 WHERE id = $1`
	_, err := r.db.ExecContext(ctx, query, id, note, time.Now().UTC())
	if err != nil {
		return fmt.Errorf("could not add body measurement note: %w", err)
	}
	return nil
}

func (r *PostgresRepository) FindBodyMeasurement(ctx context.Context, id string) (*domain.BodyMeasurement, error) {
	query := `SELECT id, "createdAt", "updatedAt", "type", value, "measuredAt", "userId", "trainerNote", "trainerNoteAt" FROM body_measurements WHERE id = $1 AND "deletedAt" IS NULL`
	var m domain.BodyMeasurement
	err := r.db.QueryRowContext(ctx, query, id).Scan(&m.ID, &m.CreatedAt, &m.UpdatedAt, &m.Type, &m.Value, &m.MeasuredAt, &m.UserId, &m.TrainerNote, &m.TrainerNoteAt)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, fmt.Errorf("could not find body measurement: %w", err)
	}
	return &m, nil
}

func (r *PostgresRepository) FindLastBodyMeasurementNote(ctx context.Context, userId string) (*domain.BodyMeasurement, error) {
	query := `SELECT id, "createdAt", "updatedAt", "type", value, "measuredAt", "userId", "trainerNote", "trainerNoteAt" FROM body_measurements WHERE "userId" = $1 AND "deletedAt" IS NULL ORDER BY "measuredAt" DESC LIMIT 1`
	var m domain.BodyMeasurement
	err := r.db.QueryRowContext(ctx, query, userId).Scan(&m.ID, &m.CreatedAt, &m.UpdatedAt, &m.Type, &m.Value, &m.MeasuredAt, &m.UserId, &m.TrainerNote, &m.TrainerNoteAt)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, fmt.Errorf("could not find last body measurement: %w", err)
	}
	return &m, nil
}

func (r *PostgresRepository) ListBodyMeasurements(ctx context.Context, userId string, since *time.Time, cursor *utils.CursorData, limit int) ([]*domain.BodyMeasurement, *utils.CursorData, error) {
	query := `SELECT id, "createdAt", "updatedAt", "type", value, "measuredAt", "userId", "trainerNote", "trainerNoteAt" FROM body_measurements WHERE "userId" = $1 AND "deletedAt" IS NULL`

	var args []interface{}
	args = append(args, userId)

	if since != nil {
		query += fmt.Sprintf(` AND "createdAt" >= $%d`, len(args)+1)
		args = append(args, *since)
	}

	if cursor != nil {
		query += fmt.Sprintf(` AND ("createdAt", id) < ($%d, $%d)`, len(args)+1, len(args)+2)
		args = append(args, cursor.CreatedAt, cursor.ID)
	}

	query += fmt.Sprintf(` ORDER BY "createdAt" DESC, id DESC LIMIT $%d`, len(args)+1)
	args = append(args, limit+1)

	rows, err := r.db.QueryContext(ctx, query, args...)
	if err != nil {
		return nil, nil, err
	}
	defer rows.Close()

	measurements := make([]*domain.BodyMeasurement, 0)
	for rows.Next() {
		m := &domain.BodyMeasurement{}
		err := rows.Scan(&m.ID, &m.CreatedAt, &m.UpdatedAt, &m.Type, &m.Value, &m.MeasuredAt, &m.UserId, &m.TrainerNote, &m.TrainerNoteAt)
		if err != nil {
			return nil, nil, err
		}
		measurements = append(measurements, m)
	}

	var nextCursor *utils.CursorData
	if len(measurements) > limit {
		nextCursor = &utils.CursorData{
			ID:        measurements[limit].ID,
			CreatedAt: measurements[limit].CreatedAt,
		}
		measurements = measurements[:limit]
	}

	return measurements, nextCursor, nil
}

func (r *PostgresRepository) AddWeightLogNote(ctx context.Context, id, note string) error {
	query := `UPDATE weight_logs SET "trainerNote" = $2, "trainerNoteAt" = $3 WHERE id = $1`
	_, err := r.db.ExecContext(ctx, query, id, note, time.Now().UTC())
	if err != nil {
		return fmt.Errorf("could not add weight log note: %w", err)
	}
	return nil
}

func (r *PostgresRepository) FindWeightLog(ctx context.Context, id string) (*domain.WeightLog, error) {
	query := `SELECT id, "createdAt", "updatedAt", weight, "measuredAt", "userId", "trainerNote", "trainerNoteAt" FROM weight_logs WHERE id = $1 AND "deletedAt" IS NULL`
	var l domain.WeightLog
	err := r.db.QueryRowContext(ctx, query, id).Scan(&l.ID, &l.CreatedAt, &l.UpdatedAt, &l.Weight, &l.MeasuredAt, &l.UserId, &l.TrainerNote, &l.TrainerNoteAt)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, fmt.Errorf("could not find weight log: %w", err)
	}
	return &l, nil
}

func (r *PostgresRepository) ListGoalsMetric(ctx context.Context, id string, since *time.Time, cursor *utils.CursorData, limit int) ([]*domain.MetricGoal, *utils.CursorData, error) {
	query := `SELECT id, "createdAt", "updatedAt", "type", "startingValue", "targetValue", deadline, "achievedAt", status, "userId" FROM metric_goals WHERE "userId" = $1 AND "deletedAt" IS NULL`

	var args []interface{}
	args = append(args, id)

	if since != nil {
		query += fmt.Sprintf(` AND "createdAt" >= $%d`, len(args)+1)
		args = append(args, *since)
	}

	if cursor != nil {
		query += fmt.Sprintf(` AND ("createdAt", id) < ($%d, $%d)`, len(args)+1, len(args)+2)
		args = append(args, cursor.CreatedAt, cursor.ID)
	}

	query += fmt.Sprintf(` ORDER BY "createdAt" DESC, id DESC LIMIT $%d`, len(args)+1)
	args = append(args, limit+1)

	rows, err := r.db.QueryContext(ctx, query, args...)
	if err != nil {
		return nil, nil, err
	}
	defer rows.Close()

	goals := make([]*domain.MetricGoal, 0)
	for rows.Next() {
		var g domain.MetricGoal
		err := rows.Scan(&g.ID, &g.CreatedAt, &g.UpdatedAt, &g.Type, &g.StartingValue, &g.TargetValue, &g.Deadline, &g.AchievedAt, &g.Status, &g.UserId)
		if err != nil {
			return nil, nil, err
		}
		goals = append(goals, &g)
	}

	var nextCursor *utils.CursorData
	if len(goals) > limit {
		nextCursor = &utils.CursorData{
			ID:        goals[limit].ID,
			CreatedAt: goals[limit].CreatedAt,
		}
		goals = goals[:limit]
	}

	return goals, nextCursor, nil
}

func (r *PostgresRepository) ListWeightLogs(ctx context.Context, userId string, since *time.Time, cursor *utils.CursorData, limit int) ([]*domain.WeightLog, *utils.CursorData, error) {
	query := `SELECT id, "createdAt", "updatedAt", weight, "measuredAt", "userId", "trainerNote", "trainerNoteAt" FROM weight_logs WHERE "userId" = $1 AND "deletedAt" IS NULL`

	var args []interface{}
	args = append(args, userId)

	if since != nil {
		query += fmt.Sprintf(` AND "createdAt" >= $%d`, len(args)+1)
		args = append(args, *since)
	}

	if cursor != nil {
		query += fmt.Sprintf(` AND ("createdAt", id) < ($%d, $%d)`, len(args)+1, len(args)+2)
		args = append(args, cursor.CreatedAt, cursor.ID)
	}

	query += fmt.Sprintf(` ORDER BY "createdAt" DESC, id DESC LIMIT $%d`, len(args)+1)
	args = append(args, limit+1)

	rows, err := r.db.QueryContext(ctx, query, args...)
	if err != nil {
		return nil, nil, err
	}
	defer rows.Close()

	logs := make([]*domain.WeightLog, 0)
	for rows.Next() {
		l := &domain.WeightLog{}
		err := rows.Scan(&l.ID, &l.CreatedAt, &l.UpdatedAt, &l.Weight, &l.MeasuredAt, &l.UserId, &l.TrainerNote, &l.TrainerNoteAt)
		if err != nil {
			return nil, nil, err
		}
		logs = append(logs, l)
	}

	var nextCursor *utils.CursorData
	if len(logs) > limit {
		nextCursor = &utils.CursorData{
			ID:        logs[limit].ID,
			CreatedAt: logs[limit].CreatedAt,
		}
		logs = logs[:limit]
	}

	return logs, nextCursor, nil
}

func (r *PostgresRepository) AddGoalMetric(ctx context.Context, g domain.MetricGoal) error {
	query := `INSERT INTO metric_goals (id, "createdAt", "updatedAt", "type", "startingValue", "targetValue", deadline, "achievedAt", status, "userId") VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10)`
	_, err := r.db.ExecContext(ctx, query, g.ID, g.CreatedAt, g.UpdatedAt, g.Type, g.StartingValue, g.TargetValue, g.Deadline, g.AchievedAt, g.Status, g.UserId)
	if err != nil {
		return fmt.Errorf("could not add goal metric: %w", err)
	}
	return nil
}
