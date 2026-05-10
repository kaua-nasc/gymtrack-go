package internal

import (
	"context"
	"database/sql"
	"fmt"
	"strings"
	"time"

	"github.com/lib/pq"
)

type TrainingPlanRepository interface {
	Create(ctx context.Context, p *TrainingPlan) error
	CountByAuthor(ctx context.Context, authorId string) (int, error)
	CreateDay(ctx context.Context, d *Day) error
	CreateExercise(ctx context.Context, e *Exercise) error
	CreateDays(ctx context.Context, days []Day) error
	DeleteDay(ctx context.Context, id string) error
	DeleteExercise(ctx context.Context, id string) error
	Find(ctx context.Context, id string) (*TrainingPlan, error)
	Update(ctx context.Context, p *TrainingPlan) error
	DeletePlan(ctx context.Context, id string) error
	ListDaysByPlan(ctx context.Context, planID string) ([]*Day, error)
	ListExercisesByDay(ctx context.Context, dayID string) ([]*Exercise, error)
	List(ctx context.Context, authorId string, cursor *CursorData, limit int) ([]*TrainingPlan, *CursorData, error)
	LikesCount(ctx context.Context, ids *[]string) (map[string]int, error)
	LikesByUser(ctx context.Context, ids *[]string, userId *string) (map[string]bool, error)
	LikePlan(ctx context.Context, like *TrainingPlanLike) error
	UnlikePlan(ctx context.Context, planId, userId string) error
	AddPlanComment(ctx context.Context, c *TrainingPlanComment) error
	FindComment(ctx context.Context, id string) (*TrainingPlanComment, error)
	RemovePlanComment(ctx context.Context, commentId string) error
	ListPlanComments(ctx context.Context, planId string, cursor *CursorData, limit int) ([]*TrainingPlanComment, *CursorData, error)
	ListSubscription(ctx context.Context, userId string) ([]*PlanSubscription, error)
	FindSubscription(ctx context.Context, planId, userId string) (*PlanSubscription, error)
	CreatePlanSubscription(ctx context.Context, s *PlanSubscription) error
	DeletePlanSubscription(ctx context.Context, s *PlanSubscription) error
	UpdateSubscriptionStatus(ctx context.Context, s PlanSubscription, status PlanSubscriptionStatus) error
	UpdateSubscriptionPrivacy(ctx context.Context, s PlanSubscription, status PlanSubscriptionType) error
	CreateSubscriptionProgress(ctx context.Context, p *PlanDayProgress) error
	AddFeedback(ctx context.Context, f *TrainingPlanFeedback) error
	LogExercise(ctx context.Context, l *ExerciseLog) error
	ListActivityWeekly(ctx context.Context, userId string, start, end time.Time) ([]time.Time, error)
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
	query := `SELECT COUNT(*) FROM training_plans WHERE "authorId" = $1`

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
		WHERE id = $1 
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

func (r *PostgresTrainingPlanRepository) List(ctx context.Context, authorId string, cursor *CursorData, limit int) ([]*TrainingPlan, *CursorData, error) {
	// sqlStr := `SELECT id, "authorId", name, visibility, "createdAt", image_url FROM training_plans WHERE 1=1`
	sqlStr := `SELECT id, "authorId", name, visibility, "createdAt", "imageUrl" FROM training_plans WHERE 1=1`

	var args []interface{}

	if authorId != "" {
		// sqlStr += ` AND (visibility = 'public' OR "authorId" = $1 OR id IN (SELECT "trainingPlanId" FROM private_participants WHERE user_id = $1))`
		sqlStr += ` AND (visibility = 'PUBLIC' OR "authorId" = $1 OR id IN (SELECT "trainingPlanId" FROM private_participants WHERE "userId" = $1))`
		args = append(args, authorId)
	} else {
		sqlStr += ` AND visibility = 'PUBLIC'`
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

	var nextCursor *CursorData
	if len(plans) > limit {
		lastItem := plans[limit-1]
		nextCursor = &CursorData{
			ID:        lastItem.Id,
			CreatedAt: lastItem.CreatedAt,
		}
		plans = plans[:limit]
	}

	return plans, nextCursor, nil
}

func (r *PostgresTrainingPlanRepository) LikesCount(ctx context.Context, ids *[]string) (map[string]int, error) {
	if ids == nil || len(*ids) == 0 {
		return map[string]int{}, nil
	}

	query := `
		SELECT "trainingPlanId", COUNT(*) 
		FROM training_plan_likes 
		WHERE "trainingPlanId" = ANY($1) 
		GROUP BY "trainingPlanId"`

	rows, err := r.db.QueryContext(ctx, query, pq.Array(*ids))
	if err != nil {
		return nil, fmt.Errorf("could not count likes: %w", err)
	}
	defer rows.Close()

	counts := make(map[string]int)
	for rows.Next() {
		var id string
		var count int
		if err := rows.Scan(&id, &count); err != nil {
			return nil, err
		}
		counts[id] = count
	}

	return counts, nil
}

func (r *PostgresTrainingPlanRepository) LikesByUser(ctx context.Context, ids *[]string, userId *string) (map[string]bool, error) {
	if ids == nil || len(*ids) == 0 || userId == nil || *userId == "" {
		return map[string]bool{}, nil
	}

	query := `
		SELECT "trainingPlanId" 
		FROM training_plan_likes 
		WHERE "likedBy" = $1 AND "trainingPlanId" = ANY($2)`

	rows, err := r.db.QueryContext(ctx, query, *userId, pq.Array(*ids))
	if err != nil {
		return nil, fmt.Errorf("could not check likes by user: %w", err)
	}
	defer rows.Close()

	likes := make(map[string]bool)
	for rows.Next() {
		var id string
		if err := rows.Scan(&id); err != nil {
			return nil, err
		}
		likes[id] = true
	}

	return likes, nil
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

func (r *PostgresTrainingPlanRepository) CreateDays(ctx context.Context, days []Day) error {
	if len(days) == 0 {
		return nil
	}

	numFields := 5
	placeholderCount := len(days) * numFields
	placeholders := make([]string, 0, len(days))
	values := make([]any, 0, placeholderCount)

	for i, d := range days {
		offset := i * numFields
		placeholders = append(placeholders, fmt.Sprintf(
			"($%d, $%d, $%d, $%d, $%d)",
			offset+1, offset+2, offset+3, offset+4, offset+5,
		))

		values = append(values, d.Id, d.Name, d.TrainingPlanId, d.CreatedAt, d.UpdatedAt)
	}

	query := fmt.Sprintf(
		"INSERT INTO days (id, name, \"trainingPlanId\", \"createdAt\", \"updatedAt\") VALUES %s",
		strings.Join(placeholders, ","),
	)

	_, err := r.db.ExecContext(ctx, query, values...)
	if err != nil {
		return fmt.Errorf("could not batch create days: %w", err)
	}

	return nil
}

func (r *PostgresTrainingPlanRepository) DeleteDay(ctx context.Context, id string) error {
	query := `delete from days where id = $1`
	_, err := r.db.ExecContext(ctx, query, id)
	if err != nil {
		return fmt.Errorf("could not delete day: %w", err)
	}
	return nil
}

func (r *PostgresTrainingPlanRepository) DeleteExercise(ctx context.Context, id string) error {
	query := `delete from exercises where id = $1`
	_, err := r.db.ExecContext(ctx, query, id)
	if err != nil {
		return fmt.Errorf("could not delete day: %w", err)
	}
	return nil
}

func (r *PostgresTrainingPlanRepository) ListDaysByPlan(ctx context.Context, planID string) ([]*Day, error) {
	query := `SELECT id, name, "trainingPlanId", "createdAt", "updatedAt" FROM days WHERE "trainingPlanId" = $1 ORDER BY "createdAt" ASC`
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
		WHERE "dayId" = $1 
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

func (r *PostgresTrainingPlanRepository) LikePlan(ctx context.Context, like *TrainingPlanLike) error {
	query := `INSERT INTO training_plan_likes (id, "likedBy", "trainingPlanId", "createdAt", "updatedAt") VALUES ($1, $2, $3, $4, $5)`
	_, err := r.db.ExecContext(ctx, query, like.Id, like.LikedBy, like.TrainingPlanId, like.CreatedAt, like.UpdatedAt)
	if err != nil {
		return fmt.Errorf("could not like plan: %w", err)
	}
	return nil
}

func (r *PostgresTrainingPlanRepository) UnlikePlan(ctx context.Context, planId, userId string) error {
	query := `DELETE FROM training_plan_likes WHERE "trainingPlanId" = $1 AND "likedBy" = $2`
	_, err := r.db.ExecContext(ctx, query, planId, userId)
	if err != nil {
		return fmt.Errorf("could not unlike plan: %w", err)
	}
	return nil
}

func (r *PostgresTrainingPlanRepository) AddPlanComment(ctx context.Context, c *TrainingPlanComment) error {
	query := `INSERT INTO training_plan_comments (id, content, "authorId", "trainingPlanId", "createdAt", "updatedAt") VALUES ($1, $2, $3, $4, $5, $6)`
	_, err := r.db.ExecContext(ctx, query, c.Id, c.Content, c.AuthorId, c.TrainingPlanId, c.CreatedAt, c.UpdatedAt)
	if err != nil {
		return fmt.Errorf("could not add comment: %w", err)
	}
	return nil
}

func (r *PostgresTrainingPlanRepository) RemovePlanComment(ctx context.Context, commentId string) error {
	query := `DELETE FROM training_plan_comments WHERE id = $1`
	_, err := r.db.ExecContext(ctx, query, commentId)
	if err != nil {
		return fmt.Errorf("could not remove comment: %w", err)
	}
	return nil
}

func (r *PostgresTrainingPlanRepository) FindComment(ctx context.Context, id string) (*TrainingPlanComment, error) {
	query := `SELECT id, content, "authorId", "trainingPlanId", "createdAt", "updatedAt" FROM training_plan_comments WHERE id = $1 LIMIT 1`
	var c TrainingPlanComment
	err := r.db.QueryRowContext(ctx, query, id).Scan(&c.Id, &c.Content, &c.AuthorId, &c.TrainingPlanId, &c.CreatedAt, &c.UpdatedAt)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, fmt.Errorf("could not find comment: %w", err)
	}
	return &c, nil
}

func (r *PostgresTrainingPlanRepository) ListPlanComments(ctx context.Context, planId string, cursor *CursorData, limit int) ([]*TrainingPlanComment, *CursorData, error) {
	query := `SELECT id, content, "authorId", "trainingPlanId", "createdAt", "updatedAt" FROM training_plan_comments WHERE "trainingPlanId" = $1`

	var args []interface{}
	args = append(args, planId)

	if cursor != nil {
		query += ` AND ("createdAt", id) < ($2, $3)`
		args = append(args, cursor.CreatedAt, cursor.ID)
	}

	query += fmt.Sprintf(` ORDER BY "createdAt" DESC, id DESC LIMIT $%d`, len(args)+1)
	args = append(args, limit+1)

	rows, err := r.db.QueryContext(ctx, query, args...)
	if err != nil {
		return nil, nil, err
	}
	defer rows.Close()

	comments := make([]*TrainingPlanComment, 0)
	for rows.Next() {
		c := &TrainingPlanComment{}
		err := rows.Scan(&c.Id, &c.Content, &c.AuthorId, &c.TrainingPlanId, &c.CreatedAt, &c.UpdatedAt)
		if err != nil {
			return nil, nil, err
		}
		comments = append(comments, c)
	}

	var nextCursor *CursorData
	if len(comments) > limit {
		lastItem := comments[limit-1]
		nextCursor = &CursorData{
			ID:        lastItem.Id,
			CreatedAt: lastItem.CreatedAt,
		}
		comments = comments[:limit]
	}

	return comments, nextCursor, nil
}

func (r *PostgresTrainingPlanRepository) ListSubscription(ctx context.Context, userId string) ([]*PlanSubscription, error) {
	query := `
		SELECT 
			subs.id, subs."createdAt", subs."updatedAt", subs."trainingPlanId", subs."userId", subs.status, subs."type", 
			plans.id, plans.name, plans."authorId", plans."timeInDays", plans.type, plans.visibility, plans.level, plans.observation, plans.pathology, plans."maxSubscriptions", plans."imageUrl", plans.description, plans."createdAt", plans."updatedAt"
		FROM plan_subscription subs LEFT JOIN training_plans plans ON subs."trainingPlanId" = plans.id WHERE subs."userId" = $1`

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
		  AND progress."updatedAt" BETWEEN $2 AND $3`

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

func (r *PostgresTrainingPlanRepository) FindSubscription(ctx context.Context, planId, userId string) (*PlanSubscription, error) {
	query := `SELECT id, "trainingPlanId", "userId", status, type, "createdAt", "updatedAt" FROM plan_subscription WHERE "trainingPlanId" = $1 AND "userId" = $2 LIMIT 1`
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
	query := `delete from training_plans where id = $1`
	_, err := r.db.ExecContext(ctx, query, id)
	if err != nil {
		return fmt.Errorf("could not delete day: %w", err)
	}
	return nil
}

func (r *PostgresTrainingPlanRepository) DeletePlanSubscription(ctx context.Context, s *PlanSubscription) error {
	query := `DELETE FROM plan_subscription WHERE id = $1`
	_, err := r.db.ExecContext(ctx, query, s.Id)
	if err != nil {
		return fmt.Errorf("could not delete subscription: %w", err)
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

func (r *PostgresTrainingPlanRepository) IsParticipant(ctx context.Context, planId, userId string) (bool, error) {
	query := `SELECT EXISTS(SELECT 1 FROM plan_participant WHERE "trainingPlanId" = $1 AND "userId" = $2)`
	var exists bool
	err := r.db.QueryRowContext(ctx, query, planId, userId).Scan(&exists)
	if err != nil {
		return false, fmt.Errorf("could not check participant status: %w", err)
	}
	return exists, nil
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
