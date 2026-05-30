package subscription

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"log/slog"
	"time"

	"github.com/kaua-nasc/gymtrack-go/apps/api/training-plan/internal/domain"
)

//go:generate go run go.uber.org/mock/mockgen -source=repository.go -destination=mock_repository.go -package=subscription
type Repository interface {
	ListSubscription(ctx context.Context, userId string) ([]*domain.PlanSubscription, error)
	FindSubscriptionByPlan(ctx context.Context, planId, userId string) (*domain.PlanSubscription, error)
	FindSubscription(ctx context.Context, id string) (*domain.PlanSubscription, error)
	CreatePlanSubscription(ctx context.Context, s *domain.PlanSubscription) error
	DeletePlanSubscription(ctx context.Context, s *domain.PlanSubscription) error
	UpdateSubscriptionStatus(ctx context.Context, s domain.PlanSubscription, status domain.PlanSubscriptionStatus) error
	UpdateSubscriptionPrivacy(ctx context.Context, s domain.PlanSubscription, status domain.PlanSubscriptionType) error
	CreateSubscriptionProgress(ctx context.Context, p *domain.PlanDayProgress) error
	FindInProgressDayProgress(ctx context.Context, subsId, dayId string) (*domain.PlanDayProgress, error)
	UpdateSubscriptionProgressStatus(ctx context.Context, id string, status domain.PlanDayProgressStatus) error
	CountSubscriptionProgress(ctx context.Context, subsId string) (int, error)
	ListWeeklyDayProgress(ctx context.Context, userId string) ([]domain.PlanDayProgress, error)
	FindLastDayProgressByUser(ctx context.Context, userId string) (*domain.PlanDayProgress, error)
	GetSubscriptionEligibility(ctx context.Context, planId, userId string) (alreadySubscribed bool, isComplete bool, err error)
	CreateAccessRequest(ctx context.Context, req *domain.PlanAccessRequest) error
	CreateInvite(ctx context.Context, i *domain.PlanInvite) error
	AddParticipant(ctx context.Context, p *domain.PlanParticipant) error
	IsParticipant(ctx context.Context, planId, userId string) (bool, error)
	FindActiveSubscription(ctx context.Context, userId string) (*domain.PlanSubscription, error)
	FindInProgressSubscription(ctx context.Context, userId string) (*domain.PlanSubscription, error)
	FindFirstDay(ctx context.Context, planId string) (*domain.Day, error)
	FindNextDayInSequence(ctx context.Context, planId string, currentSequence int) (*domain.Day, error)
	FindDayWithExercises(ctx context.Context, dayId string) (*domain.Day, error)
	CountActiveSubscriptionsByPlan(ctx context.Context, planId string) (int, error)
}

type PostgresRepository struct {
	db *sql.DB
}

func NewRepository(database *sql.DB) Repository {
	return &PostgresRepository{
		db: database,
	}
}

func (r *PostgresRepository) ListSubscription(ctx context.Context, userId string) ([]*domain.PlanSubscription, error) {
	query := `
		SELECT 
			subs.id, subs."createdAt", subs."updatedAt", subs."trainingPlanId", subs."userId", subs.status, subs."type", 
			plans.id, plans.name, plans."authorId", plans."timeInDays", plans.type, plans.visibility, plans.level, plans.observation, plans.pathology, plans."maxSubscriptions", plans."imageUrl", plans.description, plans."createdAt", plans."updatedAt",
			COALESCE((SELECT COUNT(*) FROM plan_day_progress WHERE "planSubscriptionId" = subs.id AND status IN ('COMPLETED', 'CANCELLED') AND "deletedAt" IS NULL), 0) as completed_days_count
		FROM plan_subscription subs LEFT JOIN training_plans plans ON subs."trainingPlanId" = plans.id WHERE subs."userId" = $1 AND subs."deletedAt" IS NULL`

	rows, err := r.db.QueryContext(ctx, query, userId)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	subscriptions := make([]*domain.PlanSubscription, 0)
	for rows.Next() {
		c := &domain.PlanSubscription{TrainingPlan: &domain.TrainingPlan{}}
		var completedCount int
		err := rows.Scan(
			&c.Id, &c.CreatedAt, &c.UpdatedAt, &c.TrainingPlanId, &c.UserId, &c.Status, &c.Type,
			&c.TrainingPlan.Id, &c.TrainingPlan.Name, &c.TrainingPlan.AuthorId, &c.TrainingPlan.TimeInDays,
			&c.TrainingPlan.Type, &c.TrainingPlan.Visibility, &c.TrainingPlan.Level, &c.TrainingPlan.Observation,
			&c.TrainingPlan.Pathology, &c.TrainingPlan.MaxSubscriptions, &c.TrainingPlan.ImageUrl,
			&c.TrainingPlan.Description, &c.TrainingPlan.CreatedAt, &c.TrainingPlan.UpdatedAt,
			&completedCount,
		)
		if err != nil {
			return nil, err
		}
		c.CompletedDaysCount = &completedCount
		subscriptions = append(subscriptions, c)
	}

	return subscriptions, nil
}

func (r *PostgresRepository) ListWeeklyDayProgress(ctx context.Context, userId string) ([]domain.PlanDayProgress, error) {
	now := time.Now().UTC()
	today := time.Date(now.Year(), now.Month(), now.Day(), 0, 0, 0, 0, time.UTC)
	weekday := int(today.Weekday()) // Sun=0, Mon=1, ...
	daysSinceMonday := (weekday + 6) % 7
	startOfWeek := today.AddDate(0, 0, -daysSinceMonday)
	endOfWeek := startOfWeek.AddDate(0, 0, 7).Add(-time.Nanosecond)

	query := `
		SELECT progress.id, progress."dayId", progress."planSubscriptionId", progress.status, progress."createdAt", progress."updatedAt"
		FROM plan_day_progress progress
		JOIN plan_subscription subs ON progress."planSubscriptionId" = subs.id
		WHERE subs."userId" = $1 
		  AND progress."updatedAt" BETWEEN $2 AND $3
		  AND progress."deletedAt" IS NULL 
		  AND subs."deletedAt" IS NULL`
	rows, err := r.db.QueryContext(ctx, query, userId, startOfWeek, endOfWeek)
	if err != nil {
		return nil, fmt.Errorf("could not query weekly day progress: %w", err)
	}
	defer rows.Close()

	var progresses []domain.PlanDayProgress
	for rows.Next() {
		var p domain.PlanDayProgress
		err := rows.Scan(
			&p.Id,
			&p.DayId,
			&p.PlanSubscriptionId,
			&p.Status,
			&p.CreatedAt,
			&p.UpdatedAt,
		)
		if err != nil {
			return nil, fmt.Errorf("could not scan weekly day progress: %w", err)
		}
		progresses = append(progresses, p)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}

	return progresses, nil
}

func (r *PostgresRepository) FindSubscriptionByPlan(ctx context.Context, planId, userId string) (*domain.PlanSubscription, error) {
	query := `SELECT id, "trainingPlanId", "userId", status, type, "createdAt", "updatedAt" FROM plan_subscription WHERE "trainingPlanId" = $1 AND "userId" = $2 AND "deletedAt" IS NULL ORDER BY "createdAt" DESC LIMIT 1`
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

func (r *PostgresRepository) FindSubscription(ctx context.Context, id string) (*domain.PlanSubscription, error) {
	query := `SELECT 
            subs.id, subs."createdAt", subs."updatedAt", subs."trainingPlanId", subs."userId", subs.status, subs."type", 
            plans.id, plans.name, plans."authorId", plans."timeInDays", plans.type, plans.visibility, plans.level, plans.observation, plans.pathology, plans."maxSubscriptions", plans."imageUrl", plans.description, plans."createdAt", plans."updatedAt"
        FROM plan_subscription subs 
        JOIN training_plans plans ON subs."trainingPlanId" = plans.id 
        WHERE subs.id = $1 AND subs."deletedAt" IS NULL AND plans."deletedAt" IS NULL
        LIMIT 1`

	c := &domain.PlanSubscription{}
	p := &domain.TrainingPlan{}

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

func (r *PostgresRepository) UpdateSubscriptionStatus(ctx context.Context, s domain.PlanSubscription, status domain.PlanSubscriptionStatus) error {
	query := `UPDATE plan_subscription SET status = $1, "updatedAt" = NOW() WHERE id = $2`
	_, err := r.db.ExecContext(ctx, query, status, s.Id)
	if err != nil {
		return fmt.Errorf("could not update subscription status: %w", err)
	}
	return nil
}

func (r *PostgresRepository) UpdateSubscriptionPrivacy(ctx context.Context, s domain.PlanSubscription, status domain.PlanSubscriptionType) error {
	query := `UPDATE plan_subscription SET type = $1, "updatedAt" = NOW() WHERE id = $2`
	_, err := r.db.ExecContext(ctx, query, status, s.Id)
	if err != nil {
		return fmt.Errorf("could not update subscription status: %w", err)
	}
	return nil
}

func (r *PostgresRepository) CreatePlanSubscription(ctx context.Context, s *domain.PlanSubscription) error {
	query := `INSERT INTO plan_subscription (id, "trainingPlanId", "userId", status, type, "createdAt", "updatedAt") VALUES ($1, $2, $3, $4, $5, $6, $7)`
	_, err := r.db.ExecContext(ctx, query, s.Id, s.TrainingPlanId, s.UserId, s.Status, s.Type, s.CreatedAt, s.UpdatedAt)
	if err != nil {
		return fmt.Errorf("could not create subscription: %w", err)
	}
	return nil
}

func (r *PostgresRepository) DeletePlanSubscription(ctx context.Context, s *domain.PlanSubscription) error {
	query := `UPDATE plan_subscription SET "deletedAt" = NOW() WHERE id = $1`
	_, err := r.db.ExecContext(ctx, query, s.Id)
	if err != nil {
		return fmt.Errorf("could not soft delete subscription: %w", err)
	}
	return nil
}

func (r *PostgresRepository) CreateSubscriptionProgress(ctx context.Context, p *domain.PlanDayProgress) error {
	slog.InfoContext(ctx, "creating subscription progress", slog.String("progressId", p.Id), slog.String("subsId", p.PlanSubscriptionId), slog.String("dayId", p.DayId), slog.String("status", string(p.Status)))
	query := `INSERT INTO plan_day_progress (id, "dayId", "planSubscriptionId", status, "createdAt", "updatedAt") VALUES ($1, $2, $3, $4, $5, $6)`
	_, err := r.db.ExecContext(ctx, query, p.Id, p.DayId, p.PlanSubscriptionId, p.Status, p.CreatedAt, p.UpdatedAt)
	if err != nil {
		return fmt.Errorf("could not create progress: %w", err)
	}
	return nil
}

func (r *PostgresRepository) FindInProgressDayProgress(ctx context.Context, subsId, dayId string) (*domain.PlanDayProgress, error) {
	slog.InfoContext(ctx, "searching for in-progress day progress", slog.String("subsId", subsId), slog.String("dayId", dayId))
	query := `SELECT id, "dayId", "planSubscriptionId", status, "createdAt", "updatedAt" 
	          FROM plan_day_progress 
	          WHERE "planSubscriptionId" = $1 AND "dayId" = $2 AND status = 'IN_PROGRESS' AND "deletedAt" IS NULL 
	          LIMIT 1`

	var p domain.PlanDayProgress
	err := r.db.QueryRowContext(ctx, query, subsId, dayId).Scan(
		&p.Id, &p.DayId, &p.PlanSubscriptionId, &p.Status, &p.CreatedAt, &p.UpdatedAt,
	)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			slog.InfoContext(ctx, "in-progress day progress not found in db", slog.String("subsId", subsId), slog.String("dayId", dayId))
			return nil, nil
		}
		slog.ErrorContext(ctx, "error querying in-progress day progress", slog.Any("error", err), slog.String("subsId", subsId), slog.String("dayId", dayId))
		return nil, fmt.Errorf("could not find in-progress day progress: %w", err)
	}
	slog.InfoContext(ctx, "in-progress day progress found in db", slog.String("progressId", p.Id), slog.String("status", string(p.Status)), slog.String("subsId", subsId), slog.String("dayId", dayId))
	return &p, nil
}

func (r *PostgresRepository) UpdateSubscriptionProgressStatus(ctx context.Context, id string, status domain.PlanDayProgressStatus) error {
	query := `UPDATE plan_day_progress SET status = $1, "updatedAt" = NOW() WHERE id = $2`
	_, err := r.db.ExecContext(ctx, query, status, id)
	if err != nil {
		return fmt.Errorf("could not update subscription progress status: %w", err)
	}
	return nil
}

func (r *PostgresRepository) CountSubscriptionProgress(ctx context.Context, subsId string) (int, error) {
	query := `SELECT COUNT(*) FROM plan_day_progress WHERE "deletedAt" IS NULL AND "planSubscriptionId" = $1 AND status IN ('COMPLETED', 'CANCELLED')`

	var count int
	err := r.db.QueryRowContext(ctx, query, subsId).Scan(&count)
	if err != nil {
		return 0, fmt.Errorf("could not count progress: %w", err)
	}

	return count, nil
}

func (r *PostgresRepository) CreateAccessRequest(ctx context.Context, req *domain.PlanAccessRequest) error {
	query := `INSERT INTO plan_access_request (id, "userId", "trainingPlanId", status, "createdAt", "updatedAt") VALUES ($1, $2, $3, $4, $5, $6)`
	_, err := r.db.ExecContext(ctx, query, req.Id, req.UserId, req.TrainingPlanId, req.Status, req.CreatedAt, req.UpdatedAt)
	if err != nil {
		return fmt.Errorf("could not create access request: %w", err)
	}
	return nil
}

func (r *PostgresRepository) CreateInvite(ctx context.Context, i *domain.PlanInvite) error {
	query := `INSERT INTO plan_invites (id, "planId", "senderId", "recipientId", "recipientEmail", status, "createdAt", "updatedAt") VALUES ($1, $2, $3, $4, $5, $6, $7, $8)`
	_, err := r.db.ExecContext(ctx, query, i.Id, i.PlanId, i.SenderId, i.RecipientId, i.RecipientEmail, i.Status, i.CreatedAt, i.UpdatedAt)
	if err != nil {
		return fmt.Errorf("could not create invite: %w", err)
	}
	return nil
}

func (r *PostgresRepository) AddParticipant(ctx context.Context, p *domain.PlanParticipant) error {
	query := `INSERT INTO plan_participant (id, "userId", "trainingPlanId", expiration_date, approved_at, "createdAt", "updatedAt") VALUES ($1, $2, $3, $4, $5, $6, $7)`
	_, err := r.db.ExecContext(ctx, query, p.Id, p.UserId, p.TrainingPlanId, p.ExpirationDate, p.ApprovedAt, p.CreatedAt, p.UpdatedAt)
	if err != nil {
		return fmt.Errorf("could not add participant: %w", err)
	}
	return nil
}

func (r *PostgresRepository) GetSubscriptionEligibility(ctx context.Context, planId, userId string) (bool, bool, error) {
	query := `
		SELECT 
			EXISTS (SELECT 1 FROM plan_subscription WHERE "trainingPlanId" = $1 AND "userId" = $2 AND status != 'CANCELED' AND "deletedAt" IS NULL) as already_subscribed,
			EXISTS (SELECT 1 FROM days d JOIN exercises e ON d.id = e."dayId" WHERE d."trainingPlanId" = $1 AND d."deletedAt" IS NULL AND e."deletedAt" IS NULL) as is_complete
	`
	var alreadySubscribed, isComplete bool
	err := r.db.QueryRowContext(ctx, query, planId, userId).Scan(&alreadySubscribed, &isComplete)
	if err != nil {
		return false, false, fmt.Errorf("could not check subscription eligibility: %w", err)
	}
	return alreadySubscribed, isComplete, nil
}

func (r *PostgresRepository) IsParticipant(ctx context.Context, planId, userId string) (bool, error) {
	query := `SELECT EXISTS(SELECT 1 FROM plan_participant WHERE "trainingPlanId" = $1 AND "userId" = $2 AND "deletedAt" IS NULL)`
	var exists bool
	err := r.db.QueryRowContext(ctx, query, planId, userId).Scan(&exists)
	if err != nil {
		return false, fmt.Errorf("could not check participant status: %w", err)
	}
	return exists, nil
}

func (r *PostgresRepository) FindLastDayProgressByUser(ctx context.Context, userId string) (*domain.PlanDayProgress, error) {
	query := `
		SELECT progress.id, progress."dayId", progress."planSubscriptionId", progress.status, progress."createdAt", progress."updatedAt"
		FROM plan_day_progress progress
		JOIN plan_subscription subs ON progress."planSubscriptionId" = subs.id
		WHERE subs."userId" = $1 
		  AND progress."deletedAt" IS NULL 
		  AND subs."deletedAt" IS NULL
		ORDER BY 
			progress."updatedAt" DESC
		LIMIT 1`

	var p domain.PlanDayProgress

	err := r.db.QueryRowContext(ctx, query, userId).Scan(&p.Id, &p.DayId, &p.PlanSubscriptionId, &p.Status, &p.CreatedAt, &p.UpdatedAt)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, fmt.Errorf("could not find subscription: %w", err)
	}
	return &p, nil
}

func (r *PostgresRepository) FindActiveSubscription(ctx context.Context, userId string) (*domain.PlanSubscription, error) {
	query := `
		SELECT 
			subs.id, subs."createdAt", subs."updatedAt", subs."trainingPlanId", subs."userId", subs.status, subs."type",
			plans.id, plans.name, plans."authorId", plans."timeInDays", plans.type, plans.visibility, plans.level, plans.observation, plans.pathology, plans."maxSubscriptions", plans."imageUrl", plans.description, plans."createdAt", plans."updatedAt"
		FROM plan_subscription subs 
		LEFT JOIN training_plans plans ON subs."trainingPlanId" = plans.id 
		WHERE subs."userId" = $1 
		  AND subs.status NOT IN ('COMPLETED', 'CANCELED') 
		  AND subs."deletedAt" IS NULL 
		  AND plans."deletedAt" IS NULL
		ORDER BY subs."createdAt" DESC
		LIMIT 1`

	var c domain.PlanSubscription
	c.TrainingPlan = &domain.TrainingPlan{}
	err := r.db.QueryRowContext(ctx, query, userId).Scan(
		&c.Id, &c.CreatedAt, &c.UpdatedAt, &c.TrainingPlanId, &c.UserId, &c.Status, &c.Type,
		&c.TrainingPlan.Id, &c.TrainingPlan.Name, &c.TrainingPlan.AuthorId, &c.TrainingPlan.TimeInDays,
		&c.TrainingPlan.Type, &c.TrainingPlan.Visibility, &c.TrainingPlan.Level, &c.TrainingPlan.Observation,
		&c.TrainingPlan.Pathology, &c.TrainingPlan.MaxSubscriptions, &c.TrainingPlan.ImageUrl,
		&c.TrainingPlan.Description, &c.TrainingPlan.CreatedAt, &c.TrainingPlan.UpdatedAt,
	)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, fmt.Errorf("could not find active subscription: %w", err)
	}

	return &c, nil
}

func (r *PostgresRepository) FindInProgressSubscription(ctx context.Context, userId string) (*domain.PlanSubscription, error) {
	query := `
		SELECT id, "trainingPlanId", "userId", status, type, "createdAt", "updatedAt"
		FROM plan_subscription
		WHERE "userId" = $1 AND status = 'IN_PROGRESS' AND "deletedAt" IS NULL
		LIMIT 1`

	var s domain.PlanSubscription
	err := r.db.QueryRowContext(ctx, query, userId).Scan(&s.Id, &s.TrainingPlanId, &s.UserId, &s.Status, &s.Type, &s.CreatedAt, &s.UpdatedAt)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, fmt.Errorf("could not find in progress subscription: %w", err)
	}
	return &s, nil
}

func (r *PostgresRepository) loadExercisesForDay(ctx context.Context, dayId string) ([]domain.Exercise, error) {
	query := `
		SELECT 
			id, name, "dayId", type, "setsNumber", "repsNumber", 
			description, observation, sequence, "createdAt", "updatedAt",
			"videoUrl", "imageUrl"
		FROM exercises 
		WHERE "dayId" = $1 AND "deletedAt" IS NULL
		ORDER BY sequence ASC, "createdAt" ASC`

	rows, err := r.db.QueryContext(ctx, query, dayId)
	if err != nil {
		return nil, fmt.Errorf("could not load exercises: %w", err)
	}
	defer rows.Close()

	exercises := make([]domain.Exercise, 0)
	for rows.Next() {
		var e domain.Exercise
		err := rows.Scan(
			&e.Id, &e.Name, &e.DayId, &e.Type, &e.SetsNumber, &e.RepsNumber,
			&e.Description, &e.Observation, &e.Sequence, &e.CreatedAt, &e.UpdatedAt,
			&e.VideoUrl, &e.ImageUrl,
		)
		if err != nil {
			return nil, err
		}
		exercises = append(exercises, e)
	}
	return exercises, nil
}

func (r *PostgresRepository) FindFirstDay(ctx context.Context, planId string) (*domain.Day, error) {
	query := `
		SELECT id, name, "trainingPlanId", sequence, "createdAt", "updatedAt"
		FROM days
		WHERE "trainingPlanId" = $1 AND "deletedAt" IS NULL
		ORDER BY sequence ASC, "createdAt" ASC
		LIMIT 1`

	var d domain.Day
	err := r.db.QueryRowContext(ctx, query, planId).Scan(&d.Id, &d.Name, &d.TrainingPlanId, &d.Sequence, &d.CreatedAt, &d.UpdatedAt)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, fmt.Errorf("could not find first day: %w", err)
	}

	exercises, err := r.loadExercisesForDay(ctx, d.Id)
	if err != nil {
		return nil, err
	}
	d.Exercises = exercises

	return &d, nil
}

func (r *PostgresRepository) FindNextDayInSequence(ctx context.Context, planId string, currentSequence int) (*domain.Day, error) {
	query := `
		SELECT id, name, "trainingPlanId", sequence, "createdAt", "updatedAt"
		FROM days
		WHERE "trainingPlanId" = $1 AND sequence > $2 AND "deletedAt" IS NULL
		ORDER BY sequence ASC, "createdAt" ASC
		LIMIT 1`

	var d domain.Day
	err := r.db.QueryRowContext(ctx, query, planId, currentSequence).Scan(&d.Id, &d.Name, &d.TrainingPlanId, &d.Sequence, &d.CreatedAt, &d.UpdatedAt)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, fmt.Errorf("could not find next day in sequence: %w", err)
	}

	exercises, err := r.loadExercisesForDay(ctx, d.Id)
	if err != nil {
		return nil, err
	}
	d.Exercises = exercises

	return &d, nil
}

func (r *PostgresRepository) FindDayWithExercises(ctx context.Context, dayId string) (*domain.Day, error) {
	query := `
		SELECT id, name, "trainingPlanId", sequence, "createdAt", "updatedAt"
		FROM days
		WHERE id = $1 AND "deletedAt" IS NULL
		LIMIT 1`

	var d domain.Day
	err := r.db.QueryRowContext(ctx, query, dayId).Scan(&d.Id, &d.Name, &d.TrainingPlanId, &d.Sequence, &d.CreatedAt, &d.UpdatedAt)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, fmt.Errorf("could not find day: %w", err)
	}

	exercises, err := r.loadExercisesForDay(ctx, d.Id)
	if err != nil {
		return nil, err
	}
	d.Exercises = exercises

	return &d, nil
}

func (r *PostgresRepository) CountActiveSubscriptionsByPlan(ctx context.Context, planId string) (int, error) {
	query := `SELECT COUNT(*) FROM plan_subscription WHERE "trainingPlanId" = $1 AND status IN ('NOT_STARTED', 'IN_PROGRESS') AND "deletedAt" IS NULL`

	var count int
	err := r.db.QueryRowContext(ctx, query, planId).Scan(&count)
	if err != nil {
		return 0, fmt.Errorf("could not count active subscriptions: %w", err)
	}

	return count, nil
}
