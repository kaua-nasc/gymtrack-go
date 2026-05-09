package internal

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	"github.com/lib/pq"
)

type UserRepository struct {
	db *sql.DB
}

func NewUserRepository(database *sql.DB) *UserRepository {
	return &UserRepository{
		db: database,
	}
}

func (r *UserRepository) Create(ctx context.Context, u *User) error {
	query := `INSERT INTO users (id, "firstName", "lastName", email, password, type, "createdAt", "updatedAt") VALUES ($1, $2, $3, $4, $5, $6, $7)`
	_, err := r.db.ExecContext(ctx, query, u.ID, u.FirstName, u.LastName, u.Email, u.Password, u.Type, u.CreatedAt, u.UpdatedAt)
	if err != nil {
		return fmt.Errorf("could not create user: %w", err)
	}
	return nil
}

func (r *UserRepository) FindByEmail(ctx context.Context, email string) (*User, error) {
	query := `SELECT id, "firstName", "lastName", email, password, type, "createdAt", "updatedAt" FROM users WHERE email = $1 LIMIT 1`
	var u User
	err := r.db.QueryRowContext(ctx, query, email).Scan(&u.ID, &u.FirstName, &u.LastName, &u.Email, &u.Password, &u.Type, &u.CreatedAt, &u.UpdatedAt)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}
	return &u, nil
}

func (r *UserRepository) Find(ctx context.Context, id string) (*User, error) {
	query := `SELECT id, "firstName", "lastName", email, password, type, "createdAt", "updatedAt" FROM users WHERE id = $1 LIMIT 1`
	var u User
	err := r.db.QueryRowContext(ctx, query, id).Scan(&u.ID, &u.FirstName, &u.LastName, &u.Email, &u.Password, &u.Type, &u.CreatedAt, &u.UpdatedAt)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}
	return &u, nil
}

func (r *UserRepository) ListByIDs(ctx context.Context, ids []string) ([]*User, error) {
	query := `SELECT id, "firstName", "lastName", email, type, "createdAt", "updatedAt" FROM users WHERE id = ANY($1)`
	rows, err := r.db.QueryContext(ctx, query, pq.Array(ids))
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var users []*User
	for rows.Next() {
		var u User
		err := rows.Scan(&u.ID, &u.FirstName, &u.LastName, &u.Email, &u.Type, &u.CreatedAt, &u.UpdatedAt)
		if err != nil {
			return nil, err
		}
		users = append(users, &u)
	}
	return users, nil
}

func (r *UserRepository) FollowUser(ctx context.Context, f UserFollows) error {
	query := `INSERT INTO user_follows (id, "followerId", "followingId", "createdAt", "updatedAt") VALUES ($1, $2, $3, $4, $5)`
	_, err := r.db.ExecContext(ctx, query, f.ID, f.FollowerId, f.FollowingId, f.CreatedAt, f.UpdatedAt)
	if err != nil {
		return fmt.Errorf("could not create user: %w", err)
	}
	return nil
}

func (r *UserRepository) UnfollowUser(ctx context.Context, followerId, followingId string) error {
	query := `UPDATE user_follows SET "deletedAt" = $1 WHERE "followerId" = $2 AND "followingId" = $3`
	_, err := r.db.ExecContext(ctx, query, time.Now().UTC(), followerId, followingId)
	if err != nil {
		return fmt.Errorf("could not create user: %w", err)
	}
	return nil
}

func (r *UserRepository) ListFollowing(ctx context.Context, id string) ([]*UserFollows, error) {
	query := `SELECT id, "followerId", "followingId", "createdAt", "updatedAt" FROM user_follows WHERE followerId = ANY($1)`
	rows, err := r.db.QueryContext(ctx, query, id)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var follows []*UserFollows
	for rows.Next() {
		var u UserFollows
		err := rows.Scan(&u.ID, &u.FollowerId, &u.FollowingId, &u.CreatedAt, &u.UpdatedAt)
		if err != nil {
			return nil, err
		}
		follows = append(follows, &u)
	}
	return follows, nil
}

func (r *UserRepository) ListFollower(ctx context.Context, id string) ([]*UserFollows, error) {
	query := `SELECT id, "followerId", "followingId", "createdAt", "updatedAt" FROM user_follows WHERE followingId = ANY($1)`
	rows, err := r.db.QueryContext(ctx, query, id)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var follows []*UserFollows
	for rows.Next() {
		var u UserFollows
		err := rows.Scan(&u.ID, &u.FollowerId, &u.FollowingId, &u.CreatedAt, &u.UpdatedAt)
		if err != nil {
			return nil, err
		}
		follows = append(follows, &u)
	}
	return follows, nil
}

func (r *UserRepository) FindByTrainerCode(ctx context.Context, code string) (*User, error) {
	query := `SELECT id, "firstName", "lastName", email, password, type, "createdAt", "updatedAt" FROM users WHERE "trainerInviteCode" = $1 LIMIT 1`
	var u User
	err := r.db.QueryRowContext(ctx, query, code).Scan(&u.ID, &u.FirstName, &u.LastName, &u.Email, &u.Password, &u.Type, &u.CreatedAt, &u.UpdatedAt)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}
	return &u, nil
}

func (r *UserRepository) CreateTrainerCode(ctx context.Context, id, code string) error {
	query := `UPDATE users SET trainerInviteCode = $2 WHERE id = $1`
	_, err := r.db.ExecContext(ctx, query, id, code)
	if err != nil {
		return fmt.Errorf("could not create user: %w", err)
	}
	return nil
}

func (r *UserRepository) LinkTrainer(ctx context.Context, relation TrainerStudentRelation) error {
	query := `INSERT INTO trainer_student_relationships (id, "createdAt", "updatedAt", "trainerId", "studentId", "linkedAt") VALUES ($1, $2, $3, $4, $5, $6)`
	_, err := r.db.ExecContext(ctx, query, relation.ID, relation.CreatedAt, relation.UpdatedAt, relation.TrainerId, relation.StudentId, relation.LinkedAt)
	if err != nil {
		return fmt.Errorf("could not link trainer: %w", err)
	}
	return nil
}

func (r *UserRepository) UnlinkTrainer(ctx context.Context, studentId string) error {
	query := `DELETE FROM trainer_student_relationships WHERE "studentId" = $1`
	_, err := r.db.ExecContext(ctx, query, studentId)
	if err != nil {
		return fmt.Errorf("could not unlink trainer: %w", err)
	}
	return nil
}

func (r *UserRepository) AddBodyMeasurementNote(ctx context.Context, id, note string) error {
	query := `UPDATE body_measurements SET "trainerNote" = $2, "trainerNoteAt" = $3 WHERE id = $1`
	_, err := r.db.ExecContext(ctx, query, id, note, time.Now().UTC())
	if err != nil {
		return fmt.Errorf("could not add body measurement note: %w", err)
	}
	return nil
}

func (r *UserRepository) FindLastBodyMeasurementNote(ctx context.Context, userId string) (*BodyMeasurement, error) {
	query := `SELECT id, "createdAt", "updatedAt", "type", value, "measuredAt", "userId", "trainerNote", "trainerNoteAt" FROM body_measurements WHERE "userId" = $1 ORDER BY "measuredAt" DESC LIMIT 1`
	var m BodyMeasurement
	err := r.db.QueryRowContext(ctx, query, userId).Scan(&m.ID, &m.CreatedAt, &m.UpdatedAt, &m.Type, &m.Value, &m.MeasuredAt, &m.UserId, &m.TrainerNote, &m.TrainerNoteAt)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, fmt.Errorf("could not find last body measurement: %w", err)
	}
	return &m, nil
}

func (r *UserRepository) AddWeightLogNote(ctx context.Context, id, note string) error {
	query := `UPDATE weight_logs SET "trainerNote" = $2, "trainerNoteAt" = $3 WHERE id = $1`
	_, err := r.db.ExecContext(ctx, query, id, note, time.Now().UTC())
	if err != nil {
		return fmt.Errorf("could not add weight log note: %w", err)
	}
	return nil
}

func (r *UserRepository) ChangeUserType(ctx context.Context, u User, newType UserType) error {
	query := `UPDATE users SET type = $2 WHERE id = $1`
	_, err := r.db.ExecContext(ctx, query, u.ID, newType)
	if err != nil {
		return fmt.Errorf("could not create user: %w", err)
	}
	return nil
}

func (r *UserRepository) ListGoalsMetric(ctx context.Context, id string) ([]*MetricGoal, error) {
	query := `SELECT id, "createdAt", "updatedAt", "type", "startingValue", "targetValue", deadline, "achievedAt", status, "userId" FROM metric_goals WHERE "userId" = $1`
	rows, err := r.db.QueryContext(ctx, query, id)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var goals []*MetricGoal
	for rows.Next() {
		var g MetricGoal
		err := rows.Scan(&g.ID, &g.CreatedAt, &g.UpdatedAt, &g.Type, &g.StartingValue, &g.TargetValue, &g.Deadline, &g.AchievedAt, &g.Status, &g.UserId)
		if err != nil {
			return nil, err
		}
		goals = append(goals, &g)
	}
	return goals, nil
}

func (r *UserRepository) AddGoalMetric(ctx context.Context, g MetricGoal) error {
	query := `INSERT INTO metric_goals (id, "createdAt", "updatedAt", "type", "startingValue", "targetValue", deadline, "achievedAt", status, "userId") VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10)`
	_, err := r.db.ExecContext(ctx, query, g.ID, g.CreatedAt, g.UpdatedAt, g.Type, g.StartingValue, g.TargetValue, g.Deadline, g.AchievedAt, g.Status, g.UserId)
	if err != nil {
		return fmt.Errorf("could not add goal metric: %w", err)
	}
	return nil
}

func (r *UserRepository) ListWeightHistory(ctx context.Context, userId string, cursor *CursorData, limit int) ([]*WeightLog, *CursorData, error) {
	query := `SELECT id, "createdAt", "updatedAt", weight, "measuredAt", "userId", "trainerNote", "trainerNoteAt" FROM weight_logs WHERE "userId" = $1`

	var args []interface{}
	args = append(args, userId)

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

	var logs []*WeightLog
	for rows.Next() {
		l := &WeightLog{}
		err := rows.Scan(&l.ID, &l.CreatedAt, &l.UpdatedAt, &l.Weight, &l.MeasuredAt, &l.UserId, &l.TrainerNote, &l.TrainerNoteAt)
		if err != nil {
			return nil, nil, err
		}
		logs = append(logs, l)
	}

	var nextCursor *CursorData
	if len(logs) > limit {
		nextCursor = &CursorData{
			ID:        logs[limit].ID,
			CreatedAt: logs[limit].CreatedAt,
		}
		logs = logs[:limit]
	}

	return logs, nextCursor, nil
}
