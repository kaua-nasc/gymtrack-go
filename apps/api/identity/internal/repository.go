package internal

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	"github.com/kaua-nasc/gymtrack-go/libs/cache"
	"github.com/lib/pq"
)

type UserRepository struct {
	db    *sql.DB
	cache cache.Cache
}

func NewUserRepository(database *sql.DB, cache cache.Cache) *UserRepository {
	return &UserRepository{
		db:    database,
		cache: cache,
	}
}

func (r *UserRepository) Create(ctx context.Context, u *User) error {
	query := `
		INSERT INTO users (
			id, "firstName", "lastName", email, password, type, "createdAt", "updatedAt",
			bio, "profilePictureUrl", height, "currentWeight", "weightUnit", "heightUnit",
			"trainerInviteCode", cref, "isVerified"
		) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14, $15, $16, $17)`
	_, err := r.db.ExecContext(ctx, query,
		u.ID, u.FirstName, u.LastName, u.Email, u.Password, u.Type, u.CreatedAt, u.UpdatedAt,
		u.Bio, u.ProfilePictureUrl, u.Height, u.CurrentWeight, u.WeightUnit, u.HeightUnit,
		u.TrainerInviteCode, u.Cref, u.IsVerified,
	)
	if err != nil {
		return fmt.Errorf("could not create user: %w", err)
	}
	return nil
}

func (r *UserRepository) FindByEmail(ctx context.Context, email string) (*User, error) {
	query := `
		SELECT 
			id, "firstName", "lastName", email, password, type, "createdAt", "updatedAt",
			bio, "profilePictureUrl", height, "currentWeight", "weightUnit", "heightUnit",
			"trainerInviteCode", cref, "isVerified"
		FROM users WHERE email = $1 LIMIT 1`
	var u User
	err := r.db.QueryRowContext(ctx, query, email).Scan(
		&u.ID, &u.FirstName, &u.LastName, &u.Email, &u.Password, &u.Type, &u.CreatedAt, &u.UpdatedAt,
		&u.Bio, &u.ProfilePictureUrl, &u.Height, &u.CurrentWeight, &u.WeightUnit, &u.HeightUnit,
		&u.TrainerInviteCode, &u.Cref, &u.IsVerified,
	)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}
	return &u, nil
}

func (r *UserRepository) Find(ctx context.Context, id string) (*User, error) {
	query := `
		SELECT 
			id, "firstName", "lastName", email, password, type, "createdAt", "updatedAt",
			bio, "profilePictureUrl", height, "currentWeight", "weightUnit", "heightUnit",
			"trainerInviteCode", cref, "isVerified"
		FROM users WHERE id = $1 LIMIT 1`
	var u User
	err := r.db.QueryRowContext(ctx, query, id).Scan(
		&u.ID, &u.FirstName, &u.LastName, &u.Email, &u.Password, &u.Type, &u.CreatedAt, &u.UpdatedAt,
		&u.Bio, &u.ProfilePictureUrl, &u.Height, &u.CurrentWeight, &u.WeightUnit, &u.HeightUnit,
		&u.TrainerInviteCode, &u.Cref, &u.IsVerified,
	)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}
	return &u, nil
}

func (r *UserRepository) ListByIDs(ctx context.Context, ids []string) ([]*User, error) {
	query := `
		SELECT 
			id, "firstName", "lastName", email, type, "createdAt", "updatedAt",
			bio, "profilePictureUrl", height, "currentWeight", "weightUnit", "heightUnit",
			"trainerInviteCode", cref, "isVerified"
		FROM users WHERE id = ANY($1)`
	rows, err := r.db.QueryContext(ctx, query, pq.Array(ids))
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	users := make([]*User, 0)
	for rows.Next() {
		var u User
		err := rows.Scan(
			&u.ID, &u.FirstName, &u.LastName, &u.Email, &u.Type, &u.CreatedAt, &u.UpdatedAt,
			&u.Bio, &u.ProfilePictureUrl, &u.Height, &u.CurrentWeight, &u.WeightUnit, &u.HeightUnit,
			&u.TrainerInviteCode, &u.Cref, &u.IsVerified,
		)
		if err != nil {
			return nil, err
		}
		users = append(users, &u)
	}
	return users, nil
}

func (r *UserRepository) ListStudents(ctx context.Context, trainerId string, cursor *CursorData, limit int) ([]*User, *CursorData, error) {
	query := `
		SELECT 
			u.id, u."firstName", u."lastName", u.email, u.type, u."createdAt", u."updatedAt",
			u.bio, u."profilePictureUrl", u.height, u."currentWeight", u."weightUnit", u."heightUnit",
			u."trainerInviteCode", u.cref, u."isVerified"
		FROM users u
		INNER JOIN trainer_student_relationships tsr ON tsr."studentId" = u.id
		WHERE tsr."trainerId" = $1 AND tsr."deletedAt" IS NULL
	`

	var args []interface{}
	args = append(args, trainerId)

	if cursor != nil {
		query += ` AND ("createdAt", id) < ($2, $3)`
		args = append(args, cursor.CreatedAt, cursor.ID)
	}

	query += fmt.Sprintf(` ORDER BY u."createdAt" DESC, u.id DESC LIMIT $%d`, len(args)+1)
	args = append(args, limit+1)

	rows, err := r.db.QueryContext(ctx, query, args...)
	if err != nil {
		return nil, nil, err
	}
	defer rows.Close()

	users := make([]*User, 0)
	for rows.Next() {
		var u User
		err := rows.Scan(
			&u.ID, &u.FirstName, &u.LastName, &u.Email, &u.Type, &u.CreatedAt, &u.UpdatedAt,
			&u.Bio, &u.ProfilePictureUrl, &u.Height, &u.CurrentWeight, &u.WeightUnit, &u.HeightUnit,
			&u.TrainerInviteCode, &u.Cref, &u.IsVerified,
		)
		if err != nil {
			return nil, nil, err
		}
		users = append(users, &u)
	}

	var nextCursor *CursorData
	if len(users) > limit {
		nextCursor = &CursorData{
			ID:        users[limit].ID,
			CreatedAt: users[limit].CreatedAt,
		}
		users = users[:limit]
	}

	return users, nextCursor, nil
}

func (r *UserRepository) CountFollowers(ctx context.Context, userId string) (int, error) {
	query := `SELECT count(*) FROM user_follows WHERE "followingId" = $1 AND "deletedAt" IS NULL`
	var count int
	err := r.db.QueryRowContext(ctx, query, userId).Scan(&count)
	if err != nil {
		return 0, fmt.Errorf("could not count followers: %w", err)
	}
	return count, nil
}

func (r *UserRepository) CountFollowing(ctx context.Context, userId string) (int, error) {
	query := `SELECT count(*) FROM user_follows WHERE "followerId" = $1 AND "deletedAt" IS NULL`
	var count int
	err := r.db.QueryRowContext(ctx, query, userId).Scan(&count)
	if err != nil {
		return 0, fmt.Errorf("could not count following: %w", err)
	}
	return count, nil
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

	follows := make([]*UserFollows, 0)
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

	follows := make([]*UserFollows, 0)
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
	query := `
		SELECT 
			id, "firstName", "lastName", email, password, type, "createdAt", "updatedAt",
			bio, "profilePictureUrl", height, "currentWeight", "weightUnit", "heightUnit",
			"trainerInviteCode", cref, "isVerified"
		FROM users WHERE "trainerInviteCode" = $1 LIMIT 1`
	var u User
	err := r.db.QueryRowContext(ctx, query, code).Scan(
		&u.ID, &u.FirstName, &u.LastName, &u.Email, &u.Password, &u.Type, &u.CreatedAt, &u.UpdatedAt,
		&u.Bio, &u.ProfilePictureUrl, &u.Height, &u.CurrentWeight, &u.WeightUnit, &u.HeightUnit,
		&u.TrainerInviteCode, &u.Cref, &u.IsVerified,
	)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}
	return &u, nil
}

func (r *UserRepository) CreateTrainerCode(ctx context.Context, id, code string) error {
	query := `UPDATE users SET "trainerInviteCode" = $2 WHERE id = $1`
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

func (r *UserRepository) ListBodyMeasurements(ctx context.Context, userId string, cursor *CursorData, limit int) ([]*BodyMeasurement, *CursorData, error) {
	query := `SELECT id, "createdAt", "updatedAt", "type", value, "measuredAt", "userId", "trainerNote", "trainerNoteAt" FROM body_measurements WHERE "userId" = $1`

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

	measurements := make([]*BodyMeasurement, 0)
	for rows.Next() {
		m := &BodyMeasurement{}
		err := rows.Scan(&m.ID, &m.CreatedAt, &m.UpdatedAt, &m.Type, &m.Value, &m.MeasuredAt, &m.UserId, &m.TrainerNote, &m.TrainerNoteAt)
		if err != nil {
			return nil, nil, err
		}
		measurements = append(measurements, m)
	}

	var nextCursor *CursorData
	if len(measurements) > limit {
		nextCursor = &CursorData{
			ID:        measurements[limit].ID,
			CreatedAt: measurements[limit].CreatedAt,
		}
		measurements = measurements[:limit]
	}

	return measurements, nextCursor, nil
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

func (r *UserRepository) RemoveProfilePicture(ctx context.Context, userId string) error {
	query := `UPDATE users SET "profilePictureUrl" = NULL WHERE id = $1`
	_, err := r.db.ExecContext(ctx, query, userId)
	if err != nil {
		return fmt.Errorf("could not remove profile picture: %w", err)
	}
	return nil
}

func (r *UserRepository) ChangeProfileImage(ctx context.Context, u User, pictureUrl string) error {
	query := `UPDATE users SET "profilePictureUrl" = $2 WHERE id = $1`
	_, err := r.db.ExecContext(ctx, query, u.ID, pictureUrl)
	if err != nil {
		return fmt.Errorf("could not create user: %w", err)
	}
	return nil
}

func (r *UserRepository) ListGoalsMetric(ctx context.Context, id string, cursor *CursorData, limit int) ([]*MetricGoal, *CursorData, error) {
	query := `SELECT id, "createdAt", "updatedAt", "type", "startingValue", "targetValue", deadline, "achievedAt", status, "userId" FROM metric_goals WHERE "userId" = $1`

	var args []interface{}
	args = append(args, id)

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

	goals := make([]*MetricGoal, 0)
	for rows.Next() {
		var g MetricGoal
		err := rows.Scan(&g.ID, &g.CreatedAt, &g.UpdatedAt, &g.Type, &g.StartingValue, &g.TargetValue, &g.Deadline, &g.AchievedAt, &g.Status, &g.UserId)
		if err != nil {
			return nil, nil, err
		}
		goals = append(goals, &g)
	}

	var nextCursor *CursorData
	if len(goals) > limit {
		nextCursor = &CursorData{
			ID:        goals[limit].ID,
			CreatedAt: goals[limit].CreatedAt,
		}
		goals = goals[:limit]
	}

	return goals, nextCursor, nil
}

func (r *UserRepository) AddGoalMetric(ctx context.Context, g MetricGoal) error {
	query := `INSERT INTO metric_goals (id, "createdAt", "updatedAt", "type", "startingValue", "targetValue", deadline, "achievedAt", status, "userId") VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10)`
	_, err := r.db.ExecContext(ctx, query, g.ID, g.CreatedAt, g.UpdatedAt, g.Type, g.StartingValue, g.TargetValue, g.Deadline, g.AchievedAt, g.Status, g.UserId)
	if err != nil {
		return fmt.Errorf("could not add goal metric: %w", err)
	}
	return nil
}

func (r *UserRepository) ListWeightLogs(ctx context.Context, userId string, cursor *CursorData, limit int) ([]*WeightLog, *CursorData, error) {
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

	logs := make([]*WeightLog, 0)
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

func (r *UserRepository) SaveResetCode(ctx context.Context, code, email string) error {
	return r.cache.Set(ctx, email, code, time.Minute*5)
}

func (r *UserRepository) GetResetCode(ctx context.Context, email string) (string, error) {
	return r.cache.Get(ctx, email)
}

func (r *UserRepository) Update(ctx context.Context, u *User) error {
	query := `
		UPDATE users SET
			"firstName" = $2, "lastName" = $3, email = $4, password = $5, type = $6,
			"updatedAt" = $7, bio = $8, "profilePictureUrl" = $9, height = $10,
			"currentWeight" = $11, "weightUnit" = $12, "heightUnit" = $13,
			"trainerInviteCode" = $14, cref = $15, "isVerified" = $16
		WHERE id = $1`
	_, err := r.db.ExecContext(ctx, query,
		u.ID, u.FirstName, u.LastName, u.Email, u.Password, u.Type,
		u.UpdatedAt, u.Bio, u.ProfilePictureUrl, u.Height,
		u.CurrentWeight, u.WeightUnit, u.HeightUnit,
		u.TrainerInviteCode, u.Cref, u.IsVerified,
	)
	if err != nil {
		return fmt.Errorf("could not update user: %w", err)
	}
	return nil
}
