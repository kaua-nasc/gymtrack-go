package plan

import (
	"context"
	"database/sql"
	"encoding/json"
	"fmt"

	"github.com/kaua-nasc/gymtrack-go/apps/api/training-plan/internal/domain"
	"github.com/kaua-nasc/gymtrack-go/libs/utils"
	"github.com/lib/pq"
)

//go:generate go run go.uber.org/mock/mockgen -source=repository.go -destination=mock_repository.go -package=plan
type Repository interface {
	Create(ctx context.Context, p *domain.TrainingPlan) error
	CountByAuthor(ctx context.Context, authorId string) (int, error)
	CreateDay(ctx context.Context, d *domain.Day) error
	CreateExercise(ctx context.Context, e *domain.Exercise) error
	UpdateExerciseMedia(ctx context.Context, id string, videoUrl *string, imageUrl *string) error
	DeleteDay(ctx context.Context, id string) error
	DeleteExercise(ctx context.Context, id string) error
	Find(ctx context.Context, id string) (*domain.TrainingPlan, error)
	FindComplete(ctx context.Context, id string) (*domain.TrainingPlan, error)
	Update(ctx context.Context, p *domain.TrainingPlan) error
	DeletePlan(ctx context.Context, id string) error
	ListDaysByPlan(ctx context.Context, planID string) ([]*domain.Day, error)
	ListExercisesByDay(ctx context.Context, dayID string) ([]*domain.Exercise, error)
	ListExercisesByPlan(ctx context.Context, planID string) ([]*domain.Exercise, error)
	List(ctx context.Context, authorId string, cursor *utils.CursorData, limit int) ([]*domain.TrainingPlan, *utils.CursorData, error)
	ListByIds(ctx context.Context, ids []string) ([]*domain.TrainingPlan, error)
	IsPlanComplete(ctx context.Context, planId string) (bool, error)
	FindSubscriptionByPlan(ctx context.Context, planId, userId string) (*domain.PlanSubscription, error)
	ExistsPlan(ctx context.Context, id string, publicOnly bool) (bool, error)
	UpdateRatingStats(ctx context.Context, planId string, sum *float64, count int) error
}

type PostgresRepository struct {
	db *sql.DB
}

func NewRepository(database *sql.DB) Repository {
	return &PostgresRepository{
		db: database,
	}
}

func (r *PostgresRepository) Create(ctx context.Context, p *domain.TrainingPlan) error {
	query := `
		INSERT INTO training_plans (
			id, name, "authorId", "timeInDays", type, visibility, 
			level, observation, pathology, "maxSubscriptions", 
			"imageUrl", description, "createdAt", "updatedAt"
		) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14)`

	_, err := r.db.ExecContext(ctx, query,
		p.Id, p.Name, p.AuthorId, p.TimeInDays, p.Type, p.Visibility,
		p.Level, p.Observation, p.Pathology, p.MaxSubscriptions,
		p.ImageUrl, p.Description, p.CreatedAt, p.UpdatedAt,
	)
	if err != nil {
		return fmt.Errorf("could not create training plan: %w", err)
	}

	return nil
}

func (r *PostgresRepository) Update(ctx context.Context, p *domain.TrainingPlan) error {
	query := `
		UPDATE training_plans SET 
			name = $1, "timeInDays" = $2, type = $3, visibility = $4, 
			level = $5, observation = $6, pathology = $7, 
			"maxSubscriptions" = $8, "imageUrl" = $9, description = $10, 
			"updatedAt" = $11
		WHERE id = $12`

	_, err := r.db.ExecContext(ctx, query,
		p.Name, p.TimeInDays, p.Type, p.Visibility,
		p.Level, p.Observation, p.Pathology,
		p.MaxSubscriptions, p.ImageUrl, p.Description,
		p.UpdatedAt, p.Id,
	)
	if err != nil {
		return fmt.Errorf("could not update training plan: %w", err)
	}

	return nil
}

func (r *PostgresRepository) CountByAuthor(ctx context.Context, authorId string) (int, error) {
	query := `SELECT COUNT(*) FROM training_plans WHERE "authorId" = $1 AND "deletedAt" IS NULL`

	var count int
	err := r.db.QueryRowContext(ctx, query, authorId).Scan(&count)
	if err != nil {
		return 0, fmt.Errorf("could not count training plans by author id: %w", err)
	}

	return count, nil
}

func (r *PostgresRepository) Find(ctx context.Context, id string) (*domain.TrainingPlan, error) {
	query := `
		SELECT 
			id, name, "authorId", "timeInDays", type, visibility, 
			level, observation, pathology, "maxSubscriptions", 
			"imageUrl", description, "createdAt", "updatedAt",
			"totalRatingSum", "totalRatingsCount"
		FROM training_plans 
		WHERE id = $1 AND "deletedAt" IS NULL
		LIMIT 1`

	var p domain.TrainingPlan
	err := r.db.QueryRowContext(ctx, query, id).Scan(
		&p.Id, &p.Name, &p.AuthorId, &p.TimeInDays, &p.Type, &p.Visibility,
		&p.Level, &p.Observation, &p.Pathology, &p.MaxSubscriptions,
		&p.ImageUrl, &p.Description, &p.CreatedAt, &p.UpdatedAt,
		&p.TotalRatingSum, &p.TotalRatingsCount,
	)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, fmt.Errorf("could not find training plan: %w", err)
	}

	return &p, nil
}

func (r *PostgresRepository) FindComplete(ctx context.Context, id string) (*domain.TrainingPlan, error) {
	query := `
		SELECT 
			p.id, p.name, p."authorId", p."timeInDays", p.type, p.visibility, 
			p.level, p.observation, p.pathology, p."maxSubscriptions", 
			p."imageUrl", p.description, p."createdAt"::timestamptz, p."updatedAt"::timestamptz,
			p."totalRatingSum", p."totalRatingsCount",
			COALESCE(
				(SELECT json_agg(
					json_build_object(
						'id', d.id,
						'name', d.name,
						'trainingPlanId', d."trainingPlanId",
						'sequence', d.sequence,
						'createdAt', d."createdAt"::timestamptz,
						'updatedAt', d."updatedAt"::timestamptz,
						'exercises', COALESCE(
							(SELECT json_agg(
								json_build_object(
									'id', e.id,
									'name', e.name,
									'dayId', e."dayId",
									'type', e.type,
									'setsNumber', e."setsNumber",
									'repsNumber', e."repsNumber",
									'description', e.description,
									'observation', e.observation,
									'sequence', e.sequence,
									'videoUrl', e."videoUrl",
									'imageUrl', e."imageUrl",
									'createdAt', e."createdAt"::timestamptz,
									'updatedAt', e."updatedAt"::timestamptz
								) ORDER BY e.sequence ASC, e."createdAt" ASC
							) FROM exercises e WHERE e."dayId" = d.id AND e."deletedAt" IS NULL),
							'[]'::json
						)
					) ORDER BY d.sequence ASC, d."createdAt" ASC
				) FROM days d WHERE d."trainingPlanId" = p.id AND d."deletedAt" IS NULL),
				'[]'::json
			) as days
		FROM training_plans p
		WHERE p.id = $1 AND p."deletedAt" IS NULL
		LIMIT 1`

	var p domain.TrainingPlan
	var daysJSON []byte
	err := r.db.QueryRowContext(ctx, query, id).Scan(
		&p.Id, &p.Name, &p.AuthorId, &p.TimeInDays, &p.Type, &p.Visibility,
		&p.Level, &p.Observation, &p.Pathology, &p.MaxSubscriptions,
		&p.ImageUrl, &p.Description, &p.CreatedAt, &p.UpdatedAt,
		&p.TotalRatingSum, &p.TotalRatingsCount,
		&daysJSON,
	)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, fmt.Errorf("could not find complete training plan: %w", err)
	}

	if len(daysJSON) > 0 {
		if err := json.Unmarshal(daysJSON, &p.Days); err != nil {
			return nil, fmt.Errorf("failed to unmarshal days: %w", err)
		}
	}

	return &p, nil
}

func (r *PostgresRepository) List(ctx context.Context, authorId string, cursor *utils.CursorData, limit int) ([]*domain.TrainingPlan, *utils.CursorData, error) {
	sqlStr := `
		SELECT 
			id, name, "authorId", "timeInDays", type, visibility, 
			level, observation, pathology, "maxSubscriptions", 
			"imageUrl", description, "createdAt", "updatedAt",
			"totalRatingSum", "totalRatingsCount"
		FROM training_plans 
		WHERE 1=1`

	var args []interface{}

	if authorId != "" {
		sqlStr += ` AND "authorId" = $1 AND "deletedAt" IS NULL`
		args = append(args, authorId)
	} else {
		sqlStr += ` AND visibility = 'PUBLIC' AND "deletedAt" IS NULL`
	}

	if cursor != nil {
		argCount := len(args)
		sqlStr += fmt.Sprintf(` AND ("createdAt", id) < ($%d, $%d)`, argCount+1, argCount+2)
		args = append(args, cursor.CreatedAt, cursor.ID)
	}

	sqlStr += fmt.Sprintf(` ORDER BY "createdAt" DESC, id DESC LIMIT $%d`, len(args)+1)
	args = append(args, limit+1)

	rows, err := r.db.QueryContext(ctx, sqlStr, args...)
	if err != nil {
		return nil, nil, err
	}
	defer rows.Close()

	plans := make([]*domain.TrainingPlan, 0)
	for rows.Next() {
		p := &domain.TrainingPlan{}
		err := rows.Scan(
			&p.Id, &p.Name, &p.AuthorId, &p.TimeInDays, &p.Type, &p.Visibility,
			&p.Level, &p.Observation, &p.Pathology, &p.MaxSubscriptions,
			&p.ImageUrl, &p.Description, &p.CreatedAt, &p.UpdatedAt,
			&p.TotalRatingSum, &p.TotalRatingsCount,
		)
		if err != nil {
			return nil, nil, err
		}
		plans = append(plans, p)
	}

	var nextCursor *utils.CursorData
	if len(plans) > limit {
		lastItem := plans[limit-1]
		nextCursor = &utils.CursorData{
			ID:        *lastItem.Id,
			CreatedAt: lastItem.CreatedAt,
		}
		plans = plans[:limit]
	}

	return plans, nextCursor, nil
}

func (r *PostgresRepository) ListByIds(ctx context.Context, ids []string) ([]*domain.TrainingPlan, error) {
	if len(ids) == 0 {
		return []*domain.TrainingPlan{}, nil
	}

	query := `
		SELECT 
			id, name, "authorId", "timeInDays", type, visibility, 
			level, observation, pathology, "maxSubscriptions", 
			"imageUrl", description, "createdAt", "updatedAt",
			"totalRatingSum", "totalRatingsCount"
		FROM training_plans 
		WHERE id = ANY($1) AND "deletedAt" IS NULL`

	rows, err := r.db.QueryContext(ctx, query, pq.Array(ids))
	if err != nil {
		return nil, fmt.Errorf("could not list training plans by ids: %w", err)
	}
	defer rows.Close()

	plans := make([]*domain.TrainingPlan, 0)
	for rows.Next() {
		p := &domain.TrainingPlan{}
		err := rows.Scan(
			&p.Id, &p.Name, &p.AuthorId, &p.TimeInDays, &p.Type, &p.Visibility,
			&p.Level, &p.Observation, &p.Pathology, &p.MaxSubscriptions,
			&p.ImageUrl, &p.Description, &p.CreatedAt, &p.UpdatedAt,
			&p.TotalRatingSum, &p.TotalRatingsCount,
		)
		if err != nil {
			return nil, err
		}
		plans = append(plans, p)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}

	return plans, nil
}

func (r *PostgresRepository) CreateDay(ctx context.Context, d *domain.Day) error {
	query := `INSERT INTO days (id, name, "trainingPlanId", "createdAt", "updatedAt", sequence) VALUES ($1, $2, $3, $4, $5, $6)`
	_, err := r.db.ExecContext(ctx, query, d.Id, d.Name, d.TrainingPlanId, d.CreatedAt, d.UpdatedAt, d.Sequence)
	if err != nil {
		return fmt.Errorf("could not create day: %w", err)
	}
	return nil
}

func (r *PostgresRepository) DeleteDay(ctx context.Context, id string) error {
	query := `UPDATE days SET "deletedAt" = NOW() WHERE id = $1`
	_, err := r.db.ExecContext(ctx, query, id)
	if err != nil {
		return fmt.Errorf("could not soft delete day: %w", err)
	}
	return nil
}

func (r *PostgresRepository) DeleteExercise(ctx context.Context, id string) error {
	query := `UPDATE exercises SET "deletedAt" = NOW() WHERE id = $1`
	_, err := r.db.ExecContext(ctx, query, id)
	if err != nil {
		return fmt.Errorf("could not soft delete exercise: %w", err)
	}
	return nil
}

func (r *PostgresRepository) ListDaysByPlan(ctx context.Context, planID string) ([]*domain.Day, error) {
	query := `SELECT id, name, "trainingPlanId", "createdAt", "updatedAt", sequence FROM days WHERE "trainingPlanId" = $1 AND "deletedAt" IS NULL ORDER BY sequence ASC, "createdAt" ASC`
	rows, err := r.db.QueryContext(ctx, query, planID)
	if err != nil {
		return nil, fmt.Errorf("could not list days: %w", err)
	}
	defer rows.Close()

	days := make([]*domain.Day, 0)
	for rows.Next() {
		d := &domain.Day{}
		if err := rows.Scan(&d.Id, &d.Name, &d.TrainingPlanId, &d.CreatedAt, &d.UpdatedAt, &d.Sequence); err != nil {
			return nil, err
		}
		days = append(days, d)
	}
	return days, nil
}

func (r *PostgresRepository) CreateExercise(ctx context.Context, e *domain.Exercise) error {
	query := `
		INSERT INTO exercises (
			id, name, "dayId", type, "setsNumber", "repsNumber", 
			description, observation, "createdAt", "updatedAt", sequence,
			"videoUrl", "imageUrl"
		) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13)`

	_, err := r.db.ExecContext(ctx, query,
		e.Id, e.Name, e.DayId, e.Type, e.SetsNumber, e.RepsNumber,
		e.Description, e.Observation, e.CreatedAt, e.UpdatedAt, e.Sequence,
		e.VideoUrl, e.ImageUrl,
	)
	if err != nil {
		return fmt.Errorf("could not create exercise: %w", err)
	}
	return nil
}

func (r *PostgresRepository) UpdateExerciseMedia(ctx context.Context, id string, videoUrl *string, imageUrl *string) error {
	query := `UPDATE exercises SET "videoUrl" = $1, "imageUrl" = $2, "updatedAt" = NOW() WHERE id = $3 AND "deletedAt" IS NULL`
	_, err := r.db.ExecContext(ctx, query, videoUrl, imageUrl, id)
	if err != nil {
		return fmt.Errorf("could not update exercise media: %w", err)
	}
	return nil
}

func (r *PostgresRepository) ListExercisesByDay(ctx context.Context, dayID string) ([]*domain.Exercise, error) {
	query := `
		SELECT 
			id, name, "dayId", type, "setsNumber", "repsNumber", 
			description, observation, "createdAt", "updatedAt", sequence,
			"videoUrl", "imageUrl"
		FROM exercises 
		WHERE "dayId" = $1 AND "deletedAt" IS NULL
		ORDER BY sequence ASC, "createdAt" ASC`

	rows, err := r.db.QueryContext(ctx, query, dayID)
	if err != nil {
		return nil, fmt.Errorf("could not list exercises: %w", err)
	}
	defer rows.Close()

	exercises := make([]*domain.Exercise, 0)
	for rows.Next() {
		e := &domain.Exercise{}
		if err := rows.Scan(
			&e.Id, &e.Name, &e.DayId, &e.Type, &e.SetsNumber, &e.RepsNumber,
			&e.Description, &e.Observation, &e.CreatedAt, &e.UpdatedAt, &e.Sequence,
			&e.VideoUrl, &e.ImageUrl,
		); err != nil {
			return nil, err
		}
		exercises = append(exercises, e)
	}
	return exercises, nil
}

func (r *PostgresRepository) ListExercisesByPlan(ctx context.Context, planID string) ([]*domain.Exercise, error) {
	query := `
		SELECT 
			e.id, e.name, e."dayId", e.type, e."setsNumber", e."repsNumber", 
			e.description, e.observation, e."createdAt", e."updatedAt", e.sequence,
			e."videoUrl", e."imageUrl"
		FROM exercises e
		INNER JOIN days d ON e."dayId" = d.id
		WHERE d."trainingPlanId" = $1 AND e."deletedAt" IS NULL AND d."deletedAt" IS NULL
		ORDER BY d.sequence ASC, d."createdAt" ASC, e.sequence ASC, e."createdAt" ASC`

	rows, err := r.db.QueryContext(ctx, query, planID)
	if err != nil {
		return nil, fmt.Errorf("could not list exercises by plan: %w", err)
	}
	defer rows.Close()

	exercises := make([]*domain.Exercise, 0)
	for rows.Next() {
		e := &domain.Exercise{}
		if err := rows.Scan(
			&e.Id, &e.Name, &e.DayId, &e.Type, &e.SetsNumber, &e.RepsNumber,
			&e.Description, &e.Observation, &e.CreatedAt, &e.UpdatedAt, &e.Sequence,
			&e.VideoUrl, &e.ImageUrl,
		); err != nil {
			return nil, err
		}
		exercises = append(exercises, e)
	}
	return exercises, nil
}

func (r *PostgresRepository) DeletePlan(ctx context.Context, id string) error {
	query := `UPDATE training_plans SET "deletedAt" = NOW() WHERE id = $1`
	_, err := r.db.ExecContext(ctx, query, id)
	if err != nil {
		return fmt.Errorf("could not soft delete training plan: %w", err)
	}
	return nil
}

func (r *PostgresRepository) IsPlanComplete(ctx context.Context, planId string) (bool, error) {
	query := `
		SELECT EXISTS (
			SELECT 1 FROM days d
			JOIN exercises e ON d.id = e."dayId"
			WHERE d."trainingPlanId" = $1 AND d."deletedAt" IS NULL AND e."deletedAt" IS NULL
		)`
	var complete bool
	err := r.db.QueryRowContext(ctx, query, planId).Scan(&complete)
	if err != nil {
		return false, fmt.Errorf("could not check if plan is complete: %w", err)
	}
	return complete, nil
}

func (r *PostgresRepository) FindSubscriptionByPlan(ctx context.Context, planId, userId string) (*domain.PlanSubscription, error) {
	query := `SELECT id, "trainingPlanId", "userId", status, type, "createdAt", "updatedAt" FROM plan_subscription WHERE "trainingPlanId" = $1 AND "userId" = $2 AND "deletedAt" IS NULL LIMIT 1`
	var s domain.PlanSubscription
	err := r.db.QueryRowContext(ctx, query, planId, userId).Scan(&s.Id, &s.TrainingPlanId, &s.UserId, &s.Status, &s.Type, &s.CreatedAt, &s.UpdatedAt)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, fmt.Errorf("could not find subscription: %w", err)
	}
	return &s, nil
}

func (r *PostgresRepository) ExistsPlan(ctx context.Context, id string, publicOnly bool) (bool, error) {
	query := `SELECT EXISTS(SELECT 1 FROM training_plans WHERE id = $1 AND "deletedAt" IS NULL`
	if publicOnly {
		query += ` AND visibility = 'PUBLIC'`
	}
	query += `)`

	var exists bool
	err := r.db.QueryRowContext(ctx, query, id).Scan(&exists)
	if err != nil {
		return false, fmt.Errorf("could not check plan existence: %w", err)
	}

	return exists, nil
}

func (r *PostgresRepository) UpdateRatingStats(ctx context.Context, planId string, sum *float64, count int) error {
	query := `UPDATE training_plans SET "totalRatingSum" = $1, "totalRatingsCount" = $2 WHERE id = $3`
	_, err := r.db.ExecContext(ctx, query, sum, count, planId)
	if err != nil {
		return fmt.Errorf("could not update plan rating stats: %w", err)
	}
	return nil
}
