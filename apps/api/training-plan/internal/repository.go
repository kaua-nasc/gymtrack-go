package internal

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"time"

	"github.com/kaua-nasc/gymtrack-go/libs/utils"
	"github.com/lib/pq"
)

type TrainingPlanRepository interface {
	Create(ctx context.Context, p *TrainingPlan) error
	CountByAuthor(ctx context.Context, authorId string) (int, error)
	CreateDay(ctx context.Context, d *Day) error
	CreateExercise(ctx context.Context, e *Exercise) error
	DeleteDay(ctx context.Context, id string) error
	DeleteExercise(ctx context.Context, id string) error
	Find(ctx context.Context, id string) (*TrainingPlan, error)
	Update(ctx context.Context, p *TrainingPlan) error
	DeletePlan(ctx context.Context, id string) error
	ListDaysByPlan(ctx context.Context, planID string) ([]*Day, error)
	ListExercisesByDay(ctx context.Context, dayID string) ([]*Exercise, error)
	ListExercisesByPlan(ctx context.Context, planID string) ([]*Exercise, error)
	List(ctx context.Context, authorId string, cursor *utils.CursorData, limit int) ([]*TrainingPlan, *utils.CursorData, error)
	ListByIds(ctx context.Context, ids []string) ([]*TrainingPlan, error)
	ListSubscription(ctx context.Context, userId string) ([]*PlanSubscription, error)
	FindSubscriptionByPlan(ctx context.Context, planId, userId string) (*PlanSubscription, error)
	FindSubscription(ctx context.Context, id string) (*PlanSubscription, error)
	CreatePlanSubscription(ctx context.Context, s *PlanSubscription) error
	DeletePlanSubscription(ctx context.Context, s *PlanSubscription) error
	UpdateSubscriptionStatus(ctx context.Context, s PlanSubscription, status PlanSubscriptionStatus) error
	UpdateSubscriptionPrivacy(ctx context.Context, s PlanSubscription, status PlanSubscriptionType) error
	CreateSubscriptionProgress(ctx context.Context, p *PlanDayProgress) error
	CountSubscriptionProgress(ctx context.Context, subsId string) (int, error)
	AddFeedback(ctx context.Context, f *TrainingPlanFeedback) error
	LogExercise(ctx context.Context, l *ExerciseLog) error
	ListActivityWeekly(ctx context.Context, userId string, start, end time.Time) ([]time.Time, error)
	GetSubscriptionEligibility(ctx context.Context, planId, userId string) (alreadySubscribed bool, isComplete bool, err error)
	IsPlanComplete(ctx context.Context, planId string) (bool, error)
	IsParticipant(ctx context.Context, planId, userId string) (bool, error)
}

type PostgresTrainingPlanRepository struct {
	db *sql.DB
}

func NewTrainingPlanRepository(database *sql.DB) TrainingPlanRepository {
	return &PostgresTrainingPlanRepository{
		db: database,
	}
}

func (r *PostgresTrainingPlanRepository) Create(ctx context.Context, p *TrainingPlan) error {
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

func (r *PostgresTrainingPlanRepository) Update(ctx context.Context, p *TrainingPlan) error {
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

func (r *PostgresTrainingPlanRepository) CountByAuthor(ctx context.Context, authorId string) (int, error) {
	query := `SELECT COUNT(*) FROM training_plans WHERE "authorId" = $1 AND "deletedAt" IS NULL`

	var count int
	err := r.db.QueryRowContext(ctx, query, authorId).Scan(&count)
	if err != nil {
		return 0, fmt.Errorf("could not count training plans by author id: %w", err)
	}

	return count, nil
}

func (r *PostgresTrainingPlanRepository) Find(ctx context.Context, id string) (*TrainingPlan, error) {
	query := `
		SELECT 
			id, name, "authorId", "timeInDays", type, visibility, 
			level, observation, pathology, "maxSubscriptions", 
				"imageUrl", description, "createdAt", "updatedAt"
		FROM training_plans 
		WHERE id = $1 AND "deletedAt" IS NULL
		LIMIT 1`

	var p TrainingPlan
	err := r.db.QueryRowContext(ctx, query, id).Scan(
		&p.Id, &p.Name, &p.AuthorId, &p.TimeInDays, &p.Type, &p.Visibility,
		&p.Level, &p.Observation, &p.Pathology, &p.MaxSubscriptions,
		&p.ImageUrl, &p.Description, &p.CreatedAt, &p.UpdatedAt,
	)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, fmt.Errorf("could not find training plan: %w", err)
	}

	return &p, nil
}

func (r *PostgresTrainingPlanRepository) List(ctx context.Context, authorId string, cursor *utils.CursorData, limit int) ([]*TrainingPlan, *utils.CursorData, error) {
	// sqlStr := `SELECT id, "authorId", name, visibility, "createdAt", image_url FROM training_plans WHERE 1=1`
	sqlStr := `SELECT id, "authorId", name, visibility, "createdAt", "imageUrl" FROM training_plans WHERE 1=1`

	var args []interface{}

	if authorId != "" {
		// sqlStr += ` AND (visibility = 'public' OR "authorId" = $1 OR id IN (SELECT "trainingPlanId" FROM private_participants WHERE user_id = $1))`
		sqlStr += ` AND "authorId" = $1 AND "deletedAt" IS NULL`
		args = append(args, authorId)
	} else {
		sqlStr += ` AND visibility = 'PUBLIC' AND "deletedAt" IS NULL`
	}

	if cursor != nil {
		argCount := len(args)
		// sqlStr += fmt.Sprintf(` AND ("createdAt", id) < ($%d, $%d)`, argCount+1, argCount+2)
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

	plans := make([]*TrainingPlan, 0)
	for rows.Next() {
		p := &TrainingPlan{}
		err := rows.Scan(&p.Id, &p.AuthorId, &p.Name, &p.Visibility, &p.CreatedAt, &p.ImageUrl)
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

func (r *PostgresTrainingPlanRepository) ListByIds(ctx context.Context, ids []string) ([]*TrainingPlan, error) {
	if len(ids) == 0 {
		return []*TrainingPlan{}, nil
	}

	query := `
		SELECT 
			id, name, "authorId", "timeInDays", type, visibility, 
			level, observation, pathology, "maxSubscriptions", 
			"imageUrl", description, "createdAt", "updatedAt"
		FROM training_plans 
		WHERE id = ANY($1) AND "deletedAt" IS NULL`

	rows, err := r.db.QueryContext(ctx, query, pq.Array(ids))
	if err != nil {
		return nil, fmt.Errorf("could not list training plans by ids: %w", err)
	}
	defer rows.Close()

	plans := make([]*TrainingPlan, 0)
	for rows.Next() {
		p := &TrainingPlan{}
		err := rows.Scan(
			&p.Id, &p.Name, &p.AuthorId, &p.TimeInDays, &p.Type, &p.Visibility,
			&p.Level, &p.Observation, &p.Pathology, &p.MaxSubscriptions,
			&p.ImageUrl, &p.Description, &p.CreatedAt, &p.UpdatedAt,
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

// Day Methods
func (r *PostgresTrainingPlanRepository) CreateDay(ctx context.Context, d *Day) error {
	query := `INSERT INTO days (id, name, "trainingPlanId", "createdAt", "updatedAt") VALUES ($1, $2, $3, $4, $5)`
	_, err := r.db.ExecContext(ctx, query, d.Id, d.Name, d.TrainingPlanId, d.CreatedAt, d.UpdatedAt)
	if err != nil {
		return fmt.Errorf("could not create day: %w", err)
	}
	return nil
}

func (r *PostgresTrainingPlanRepository) DeleteDay(ctx context.Context, id string) error {
	query := `UPDATE days SET "deletedAt" = NOW() WHERE id = $1`
	_, err := r.db.ExecContext(ctx, query, id)
	if err != nil {
		return fmt.Errorf("could not soft delete day: %w", err)
	}
	return nil
}

func (r *PostgresTrainingPlanRepository) DeleteExercise(ctx context.Context, id string) error {
	query := `UPDATE exercises SET "deletedAt" = NOW() WHERE id = $1`
	_, err := r.db.ExecContext(ctx, query, id)
	if err != nil {
		return fmt.Errorf("could not soft delete exercise: %w", err)
	}
	return nil
}

func (r *PostgresTrainingPlanRepository) ListDaysByPlan(ctx context.Context, planID string) ([]*Day, error) {
	query := `SELECT id, name, "trainingPlanId", "createdAt", "updatedAt" FROM days WHERE "trainingPlanId" = $1 AND "deletedAt" IS NULL ORDER BY "createdAt" ASC`
	rows, err := r.db.QueryContext(ctx, query, planID)
	if err != nil {
		return nil, fmt.Errorf("could not list days: %w", err)
	}
	defer rows.Close()

	days := make([]*Day, 0)
	for rows.Next() {
		d := &Day{}
		if err := rows.Scan(&d.Id, &d.Name, &d.TrainingPlanId, &d.CreatedAt, &d.UpdatedAt); err != nil {
			return nil, err
		}
		days = append(days, d)
	}
	return days, nil
}

// Exercise Methods
func (r *PostgresTrainingPlanRepository) CreateExercise(ctx context.Context, e *Exercise) error {
	query := `
		INSERT INTO exercises (
			id, name, "dayId", type, "setsNumber", "repsNumber", 
				description, observation, "createdAt", "updatedAt"
		) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10)`

	_, err := r.db.ExecContext(ctx, query,
		e.Id, e.Name, e.DayId, e.Type, e.SetsNumber, e.RepsNumber,
		e.Description, e.Observation, e.CreatedAt, e.UpdatedAt,
	)
	if err != nil {
		return fmt.Errorf("could not create exercise: %w", err)
	}
	return nil
}

func (r *PostgresTrainingPlanRepository) ListExercisesByDay(ctx context.Context, dayID string) ([]*Exercise, error) {
	query := `
		SELECT 
			id, name, "dayId", type, "setsNumber", "repsNumber", 
				description, observation, "createdAt", "updatedAt" 
		FROM exercises 
		WHERE "dayId" = $1 AND "deletedAt" IS NULL
			ORDER BY "createdAt" ASC`

	rows, err := r.db.QueryContext(ctx, query, dayID)
	if err != nil {
		return nil, fmt.Errorf("could not list exercises: %w", err)
	}
	defer rows.Close()

	exercises := make([]*Exercise, 0)
	for rows.Next() {
		e := &Exercise{}
		if err := rows.Scan(
			&e.Id, &e.Name, &e.DayId, &e.Type, &e.SetsNumber, &e.RepsNumber,
			&e.Description, &e.Observation, &e.CreatedAt, &e.UpdatedAt,
		); err != nil {
			return nil, err
		}
		exercises = append(exercises, e)
	}
	return exercises, nil
}

func (r *PostgresTrainingPlanRepository) ListExercisesByPlan(ctx context.Context, planID string) ([]*Exercise, error) {
	query := `
		SELECT 
			e.id, e.name, e."dayId", e.type, e."setsNumber", e."repsNumber", 
			e.description, e.observation, e."createdAt", e."updatedAt" 
		FROM exercises e
		INNER JOIN days d ON e."dayId" = d.id
		WHERE d."trainingPlanId" = $1 AND e."deletedAt" IS NULL AND d."deletedAt" IS NULL
		ORDER BY d."createdAt" ASC, e."createdAt" ASC`

	rows, err := r.db.QueryContext(ctx, query, planID)
	if err != nil {
		return nil, fmt.Errorf("could not list exercises by plan: %w", err)
	}
	defer rows.Close()

	exercises := make([]*Exercise, 0)
	for rows.Next() {
		e := &Exercise{}
		if err := rows.Scan(
			&e.Id, &e.Name, &e.DayId, &e.Type, &e.SetsNumber, &e.RepsNumber,
			&e.Description, &e.Observation, &e.CreatedAt, &e.UpdatedAt,
		); err != nil {
			return nil, err
		}
		exercises = append(exercises, e)
	}
	return exercises, nil
}

func (r *PostgresTrainingPlanRepository) ListSubscription(ctx context.Context, userId string) ([]*PlanSubscription, error) {
	query := `
		SELECT 
			subs.id, subs."createdAt", subs."updatedAt", subs."trainingPlanId", subs."userId", subs.status, subs."type", 
			plans.id, plans.name, plans."authorId", plans."timeInDays", plans.type, plans.visibility, plans.level, plans.observation, plans.pathology, plans."maxSubscriptions", plans."imageUrl", plans.description, plans."createdAt", plans."updatedAt"
		FROM plan_subscription subs LEFT JOIN training_plans plans ON subs."trainingPlanId" = plans.id WHERE subs."userId" = $1 AND subs."deletedAt" IS NULL AND plans."deletedAt" IS NULL`

	rows, err := r.db.QueryContext(ctx, query, userId)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	subscriptions := make([]*PlanSubscription, 0)
	for rows.Next() {
		c := &PlanSubscription{TrainingPlan: &TrainingPlan{}}
		err := rows.Scan(&c.Id, &c.CreatedAt, &c.UpdatedAt, &c.TrainingPlanId, &c.UserId, &c.Status, &c.Type, &c.TrainingPlan.Id,
			&c.TrainingPlan.Name, &c.TrainingPlan.AuthorId, &c.TrainingPlan.TimeInDays, &c.TrainingPlan.Type, &c.TrainingPlan.Visibility,
			&c.TrainingPlan.Level, &c.TrainingPlan.Observation, &c.TrainingPlan.Pathology, &c.TrainingPlan.MaxSubscriptions,
			&c.TrainingPlan.ImageUrl, &c.TrainingPlan.Description, &c.TrainingPlan.CreatedAt, &c.TrainingPlan.UpdatedAt)
		if err != nil {
			return nil, err
		}
		subscriptions = append(subscriptions, c)
	}

	return subscriptions, nil
}

func (r *PostgresTrainingPlanRepository) ListActivityWeekly(ctx context.Context, userId string, start, end time.Time) ([]time.Time, error) {
	query := `
		SELECT progress."updatedAt" 
		FROM plan_day_progress progress
		JOIN plan_subscription subs ON progress."planSubscriptionId" = subs.id
		WHERE subs."userId" = $1 
		  AND progress.status = 'COMPLETED'
		  AND progress."updatedAt" BETWEEN $2 AND $3
		  AND progress."deletedAt" IS NULL AND subs."deletedAt" IS NULL`

	rows, err := r.db.QueryContext(ctx, query, userId, start, end)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var dates []time.Time
	for rows.Next() {
		var t time.Time
		if err := rows.Scan(&t); err != nil {
			return nil, err
		}
		dates = append(dates, t)
	}

	return dates, nil
}

func (r *PostgresTrainingPlanRepository) FindSubscriptionByPlan(ctx context.Context, planId, userId string) (*PlanSubscription, error) {
	query := `SELECT id, "trainingPlanId", "userId", status, type, "createdAt", "updatedAt" FROM plan_subscription WHERE "trainingPlanId" = $1 AND "userId" = $2 AND "deletedAt" IS NULL LIMIT 1`
	var s PlanSubscription
	err := r.db.QueryRowContext(ctx, query, planId, userId).Scan(&s.Id, &s.TrainingPlanId, &s.UserId, &s.Status, &s.Type, &s.CreatedAt, &s.UpdatedAt)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, fmt.Errorf("could not find subscription: %w", err)
	}
	return &s, nil
}

func (r *PostgresTrainingPlanRepository) FindSubscription(ctx context.Context, id string) (*PlanSubscription, error) {
	query := `SELECT 
            subs.id, subs."createdAt", subs."updatedAt", subs."trainingPlanId", subs."userId", subs.status, subs."type", 
            plans.id, plans.name, plans."authorId", plans."timeInDays", plans.type, plans.visibility, plans.level, plans.observation, plans.pathology, plans."maxSubscriptions", plans."imageUrl", plans.description, plans."createdAt", plans."updatedAt"
        FROM plan_subscription subs 
        JOIN training_plans plans ON subs."trainingPlanId" = plans.id 
        WHERE subs.id = $1 AND subs."deletedAt" IS NULL AND plans."deletedAt" IS NULL
        LIMIT 1`

	c := &PlanSubscription{}
	p := &TrainingPlan{}

	err := r.db.QueryRowContext(ctx, query, id).Scan(
		&c.Id, &c.CreatedAt, &c.UpdatedAt, &c.TrainingPlanId, &c.UserId, &c.Status, &c.Type,
		&p.Id, &p.Name, &p.AuthorId, &p.TimeInDays, &p.Type, &p.Visibility,
		&p.Level, &p.Observation, &p.Pathology, &p.MaxSubscriptions,
		&p.ImageUrl, &p.Description, &p.CreatedAt, &p.UpdatedAt,
	)

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil
		}
		return nil, fmt.Errorf("could not find subscription: %w", err)
	}

	c.TrainingPlan = p

	return c, nil
}

func (r *PostgresTrainingPlanRepository) UpdateSubscriptionStatus(ctx context.Context, s PlanSubscription, status PlanSubscriptionStatus) error {
	query := `UPDATE plan_subscription SET status = $1, "updatedAt" = NOW() WHERE id = $2`
	_, err := r.db.ExecContext(ctx, query, status, s.Id)
	if err != nil {
		return fmt.Errorf("could not update subscription status: %w", err)
	}
	return nil
}

func (r *PostgresTrainingPlanRepository) UpdateSubscriptionPrivacy(ctx context.Context, s PlanSubscription, status PlanSubscriptionType) error {
	query := `UPDATE plan_subscription SET type = $1, "updatedAt" = NOW() WHERE id = $2`
	_, err := r.db.ExecContext(ctx, query, status, s.Id)
	if err != nil {
		return fmt.Errorf("could not update subscription status: %w", err)
	}
	return nil
}

func (r *PostgresTrainingPlanRepository) CreatePlanSubscription(ctx context.Context, s *PlanSubscription) error {
	query := `INSERT INTO plan_subscription (id, "trainingPlanId", "userId", status, type, "createdAt", "updatedAt") VALUES ($1, $2, $3, $4, $5, $6, $7)`
	_, err := r.db.ExecContext(ctx, query, s.Id, s.TrainingPlanId, s.UserId, s.Status, s.Type, s.CreatedAt, s.UpdatedAt)
	if err != nil {
		return fmt.Errorf("could not create subscription: %w", err)
	}
	return nil
}

func (r *PostgresTrainingPlanRepository) DeletePlan(ctx context.Context, id string) error {
	query := `UPDATE training_plans SET "deletedAt" = NOW() WHERE id = $1`
	_, err := r.db.ExecContext(ctx, query, id)
	if err != nil {
		return fmt.Errorf("could not soft delete training plan: %w", err)
	}
	return nil
}

func (r *PostgresTrainingPlanRepository) DeletePlanSubscription(ctx context.Context, s *PlanSubscription) error {
	query := `UPDATE plan_subscription SET "deletedAt" = NOW() WHERE id = $1`
	_, err := r.db.ExecContext(ctx, query, s.Id)
	if err != nil {
		return fmt.Errorf("could not soft delete subscription: %w", err)
	}
	return nil
}

func (r *PostgresTrainingPlanRepository) CreateSubscriptionProgress(ctx context.Context, p *PlanDayProgress) error {
	query := `INSERT INTO plan_day_progress (id, "dayId", "planSubscriptionId", status, "createdAt", "updatedAt") VALUES ($1, $2, $3, $4, $5, $6)`
	_, err := r.db.ExecContext(ctx, query, p.Id, p.DayId, p.PlanSubscriptionId, p.Status, p.CreatedAt, p.UpdatedAt)
	if err != nil {
		return fmt.Errorf("could not create progress: %w", err)
	}
	return nil
}

func (r *PostgresTrainingPlanRepository) CountSubscriptionProgress(ctx context.Context, subsId string) (int, error) {
	query := `SELECT COUNT(*) FROM plan_day_progress WHERE "deletedAt" IS NULL AND "planSubscriptionId" = $1`

	var count int
	err := r.db.QueryRowContext(ctx, query, subsId).Scan(&count)
	if err != nil {
		return 0, fmt.Errorf("could not count progress: %w", err)
	}

	return count, nil
}

// Access and Invite Methods
func (r *PostgresTrainingPlanRepository) CreateAccessRequest(ctx context.Context, req *PlanAccessRequest) error {
	query := `INSERT INTO plan_access_request (id, "userId", "trainingPlanId", status, "createdAt", "updatedAt") VALUES ($1, $2, $3, $4, $5, $6)`
	_, err := r.db.ExecContext(ctx, query, req.Id, req.UserId, req.TrainingPlanId, req.Status, req.CreatedAt, req.UpdatedAt)
	if err != nil {
		return fmt.Errorf("could not create access request: %w", err)
	}
	return nil
}

func (r *PostgresTrainingPlanRepository) CreateInvite(ctx context.Context, i *PlanInvite) error {
	query := `INSERT INTO plan_invites (id, "planId", "senderId", "recipientId", "recipientEmail", status, "createdAt", "updatedAt") VALUES ($1, $2, $3, $4, $5, $6, $7, $8)`
	_, err := r.db.ExecContext(ctx, query, i.Id, i.PlanId, i.SenderId, i.RecipientId, i.RecipientEmail, i.Status, i.CreatedAt, i.UpdatedAt)
	if err != nil {
		return fmt.Errorf("could not create invite: %w", err)
	}
	return nil
}

// Participant Methods
func (r *PostgresTrainingPlanRepository) AddParticipant(ctx context.Context, p *PlanParticipant) error {
	query := `INSERT INTO plan_participant (id, "userId", "trainingPlanId", expiration_date, approved_at, "createdAt", "updatedAt") VALUES ($1, $2, $3, $4, $5, $6, $7)`
	_, err := r.db.ExecContext(ctx, query, p.Id, p.UserId, p.TrainingPlanId, p.ExpirationDate, p.ApprovedAt, p.CreatedAt, p.UpdatedAt)
	if err != nil {
		return fmt.Errorf("could not add participant: %w", err)
	}
	return nil
}

func (r *PostgresTrainingPlanRepository) GetSubscriptionEligibility(ctx context.Context, planId, userId string) (bool, bool, error) {
	query := `
		SELECT 
			EXISTS (SELECT 1 FROM plan_subscription WHERE "trainingPlanId" = $1 AND "userId" = $2 AND "deletedAt" IS NULL) as already_subscribed,
			EXISTS (SELECT 1 FROM days d JOIN exercises e ON d.id = e."dayId" WHERE d."trainingPlanId" = $1 AND d."deletedAt" IS NULL AND e."deletedAt" IS NULL) as is_complete
	`
	var alreadySubscribed, isComplete bool
	err := r.db.QueryRowContext(ctx, query, planId, userId).Scan(&alreadySubscribed, &isComplete)
	if err != nil {
		return false, false, fmt.Errorf("could not check subscription eligibility: %w", err)
	}
	return alreadySubscribed, isComplete, nil
}

func (r *PostgresTrainingPlanRepository) IsParticipant(ctx context.Context, planId, userId string) (bool, error) {
	query := `SELECT EXISTS(SELECT 1 FROM plan_participant WHERE "trainingPlanId" = $1 AND "userId" = $2 AND "deletedAt" IS NULL)`
	var exists bool
	err := r.db.QueryRowContext(ctx, query, planId, userId).Scan(&exists)
	if err != nil {
		return false, fmt.Errorf("could not check participant status: %w", err)
	}
	return exists, nil
}

func (r *PostgresTrainingPlanRepository) IsPlanComplete(ctx context.Context, planId string) (bool, error) {
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

// Feedback and Log Methods
func (r *PostgresTrainingPlanRepository) AddFeedback(ctx context.Context, f *TrainingPlanFeedback) error {
	query := `INSERT INTO training_plan_feedbacks (id, "trainingPlanId", "userId", rating, message, "createdAt", "updatedAt") VALUES ($1, $2, $3, $4, $5, $6, $7)`
	_, err := r.db.ExecContext(ctx, query, f.Id, f.TrainingPlanId, f.UserId, f.Rating, f.Message, f.CreatedAt, f.UpdatedAt)
	if err != nil {
		return fmt.Errorf("could not add feedback: %w", err)
	}
	return nil
}

func (r *PostgresTrainingPlanRepository) LogExercise(ctx context.Context, l *ExerciseLog) error {
	query := `INSERT INTO exercise_logs (id, "userId", "exerciseId", reps, weight, notes, "createdAt", "updatedAt") VALUES ($1, $2, $3, $4, $5, $6, $7, $8)`
	_, err := r.db.ExecContext(ctx, query, l.Id, l.UserId, l.ExerciseId, pq.Array(l.Reps), pq.Array(l.Weight), l.Notes, l.CreatedAt, l.UpdatedAt)
	if err != nil {
		return fmt.Errorf("could not log exercise: %w", err)
	}
	return nil
}
