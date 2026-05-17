package internal

import (
	"context"
	"database/sql"
	"encoding/json"
	"fmt"
	"time"

	"github.com/kaua-nasc/gymtrack-go/libs/cache"
	"github.com/kaua-nasc/gymtrack-go/libs/utils"
	"github.com/lib/pq"
)

type UserRepository interface {
	Create(ctx context.Context, u *User) error
	Update(ctx context.Context, u *User) error
	Find(ctx context.Context, id string, currentUserId string) (*User, error)
	FindByEmail(ctx context.Context, email string) (*User, error)
	SaveResetCode(ctx context.Context, code, email string) error
	GetResetCode(ctx context.Context, email string) (string, error)
	SaveVerificationCode(ctx context.Context, code, email string) error
	GetVerificationCode(ctx context.Context, email string) (string, error)
	ListByIDs(ctx context.Context, ids []string) ([]*User, error)
	ListFollowing(ctx context.Context, id string, cursor *utils.CursorData, limit int) ([]*User, *utils.CursorData, error)
	ListFollower(ctx context.Context, id string, cursor *utils.CursorData, limit int) ([]*User, *utils.CursorData, error)
	CountFollowers(ctx context.Context, userId string) (int, error)
	CountFollowing(ctx context.Context, userId string) (int, error)
	FollowUser(ctx context.Context, f UserFollows) error
	UnfollowUser(ctx context.Context, followerId, followingId string) error
	CreateTrainerCode(ctx context.Context, id, code string) error
	FindByTrainerCode(ctx context.Context, code string) (*User, error)
	LinkTrainer(ctx context.Context, relation TrainerStudentRelation) error
	UnlinkTrainer(ctx context.Context, studentId string) error
	ListStudents(ctx context.Context, trainerId string, cursor *utils.CursorData, limit int) ([]*User, *utils.CursorData, error)
	AddBodyMeasurementNote(ctx context.Context, id, note string) error
	FindLastBodyMeasurementNote(ctx context.Context, userId string) (*BodyMeasurement, error)
	ListBodyMeasurements(ctx context.Context, userId string, cursor *utils.CursorData, limit int) ([]*BodyMeasurement, *utils.CursorData, error)
	AddWeightLogNote(ctx context.Context, id, note string) error
	ChangeUserType(ctx context.Context, u User, newType UserType) error
	RemoveProfilePicture(ctx context.Context, userId string) error
	ChangeProfileImage(ctx context.Context, u User, pictureUrl string) error
	ListGoalsMetric(ctx context.Context, id string, cursor *utils.CursorData, limit int) ([]*MetricGoal, *utils.CursorData, error)
	ListWeightLogs(ctx context.Context, userId string, cursor *utils.CursorData, limit int) ([]*WeightLog, *utils.CursorData, error)
	AddGoalMetric(ctx context.Context, g MetricGoal) error
}

type PostgresUserRepository struct {
	db    *sql.DB
	cache cache.Cache
}

func NewUserRepository(database *sql.DB, cache cache.Cache) UserRepository {
	return &PostgresUserRepository{
		db:    database,
		cache: cache,
	}
}

func (r PostgresUserRepository) Create(ctx context.Context, u *User) error {
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

func (r *PostgresUserRepository) FindByEmail(ctx context.Context, email string) (*User, error) {
	query := `
		SELECT 
			u.id, u."firstName", u."lastName", u.email, u.password, u.type, u."createdAt", u."updatedAt",
			u.bio, u."profilePictureUrl", u.height, u."currentWeight", u."weightUnit", u."heightUnit",
			u."trainerInviteCode", u.cref, u."isVerified",
			(SELECT json_build_object(
				'id', tsr.id, 'createdAt', tsr."createdAt"::timestamptz, 'updatedAt', tsr."updatedAt"::timestamptz,
				'trainerId', tsr."trainerId", 'studentId', tsr."studentId", 'linkedAt', tsr."linkedAt"::timestamptz,
				'trainer', json_build_object(
					'id', t.id, 'firstName', t."firstName", 'lastName', t."lastName",
					'email', t.email, 'type', t.type, 'profilePictureUrl', t."profilePictureUrl"
				)
			) FROM trainer_student_relationships tsr 
			  LEFT JOIN users t ON tsr."trainerId" = t.id 
			  WHERE tsr."studentId" = u.id AND tsr."deletedAt" IS NULL LIMIT 1) as student_of,
			COALESCE((SELECT json_agg(json_build_object(
				'id', tsr.id, 'createdAt', tsr."createdAt"::timestamptz, 'updatedAt', tsr."updatedAt"::timestamptz,
				'trainerId', tsr."trainerId", 'studentId', tsr."studentId", 'linkedAt', tsr."linkedAt"::timestamptz,
				'student', json_build_object(
					'id', s.id, 'firstName', s."firstName", 'lastName', s."lastName",
					'email', s.email, 'type', s.type, 'profilePictureUrl', s."profilePictureUrl"
				)
			)) FROM trainer_student_relationships tsr 
			   LEFT JOIN users s ON tsr."studentId" = s.id 
			   WHERE tsr."trainerId" = u.id AND tsr."deletedAt" IS NULL), '[]'::json) as trainer_of
		FROM users u WHERE email = $1 AND u."deletedAt" IS NULL LIMIT 1`
	var u User
	var studentOfJSON, trainerOfJSON []byte
	err := r.db.QueryRowContext(ctx, query, email).Scan(
		&u.ID, &u.FirstName, &u.LastName, &u.Email, &u.Password, &u.Type, &u.CreatedAt, &u.UpdatedAt,
		&u.Bio, &u.ProfilePictureUrl, &u.Height, &u.CurrentWeight, &u.WeightUnit, &u.HeightUnit,
		&u.TrainerInviteCode, &u.Cref, &u.IsVerified,
		&studentOfJSON, &trainerOfJSON,
	)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}
	if len(studentOfJSON) > 0 {
		json.Unmarshal(studentOfJSON, &u.StudentOf)
	}
	if len(trainerOfJSON) > 0 {
		json.Unmarshal(trainerOfJSON, &u.TrainerOf)
	}
	return &u, nil
}

func (r *PostgresUserRepository) Find(ctx context.Context, id string, currentUserId string) (*User, error) {
	query := `
		SELECT
			u.id, u."firstName", u."lastName", u.email, u.password, u.type, u."createdAt", u."updatedAt",
			u.bio, u."profilePictureUrl", u.height, u."currentWeight", u."weightUnit", u."heightUnit",
			u."trainerInviteCode", u.cref, u."isVerified",
			(SELECT json_build_object(
				'id', tsr.id, 'createdAt', tsr."createdAt"::timestamptz, 'updatedAt', tsr."updatedAt"::timestamptz,
				'trainerId', tsr."trainerId", 'studentId', tsr."studentId", 'linkedAt', tsr."linkedAt"::timestamptz,
				'trainer', json_build_object(
					'id', t.id, 'firstName', t."firstName", 'lastName', t."lastName",
					'email', t.email, 'type', t.type, 'profilePictureUrl', t."profilePictureUrl"
				)
			) FROM trainer_student_relationships tsr
			  LEFT JOIN users t ON tsr."trainerId" = t.id
			  WHERE tsr."studentId" = u.id AND tsr."deletedAt" IS NULL LIMIT 1) as student_of,
			COALESCE((SELECT json_agg(json_build_object(
				'id', tsr.id, 'createdAt', tsr."createdAt"::timestamptz, 'updatedAt', tsr."updatedAt"::timestamptz,
				'trainerId', tsr."trainerId", 'studentId', tsr."studentId", 'linkedAt', tsr."linkedAt"::timestamptz,
				'student', json_build_object(
					'id', s.id, 'firstName', s."firstName", 'lastName', s."lastName",
					'email', s.email, 'type', s.type, 'profilePictureUrl', s."profilePictureUrl"
				)
			)) FROM trainer_student_relationships tsr
			   LEFT JOIN users s ON tsr."studentId" = s.id
			   WHERE tsr."trainerId" = u.id AND tsr."deletedAt" IS NULL), '[]'::json) as trainer_of,
			EXISTS (SELECT 1 FROM user_follows WHERE "followerId" = NULLIF($2, '')::uuid AND "followingId" = u.id AND "deletedAt" IS NULL) as is_following
			FROM users u WHERE id = $1 AND u."deletedAt" IS NULL LIMIT 1`
	var u User
	var studentOfJSON, trainerOfJSON []byte
	var isFollowing bool
	err := r.db.QueryRowContext(ctx, query, id, currentUserId).Scan(
		&u.ID, &u.FirstName, &u.LastName, &u.Email, &u.Password, &u.Type, &u.CreatedAt, &u.UpdatedAt,
		&u.Bio, &u.ProfilePictureUrl, &u.Height, &u.CurrentWeight, &u.WeightUnit, &u.HeightUnit,
		&u.TrainerInviteCode, &u.Cref, &u.IsVerified,
		&studentOfJSON, &trainerOfJSON, &isFollowing,
	)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}

	u.IsFollowing = &isFollowing

	if len(studentOfJSON) > 0 {
		json.Unmarshal(studentOfJSON, &u.StudentOf)
	}
	if len(trainerOfJSON) > 0 {
		json.Unmarshal(trainerOfJSON, &u.TrainerOf)
	}
	return &u, nil
}

func (r *PostgresUserRepository) ListByIDs(ctx context.Context, ids []string) ([]*User, error) {
	query := `
		SELECT 
			u.id, u."firstName", u."lastName", u.email, u.type, u."createdAt", u."updatedAt",
			u.bio, u."profilePictureUrl", u.height, u."currentWeight", u."weightUnit", u."heightUnit",
			u."trainerInviteCode", u.cref, u."isVerified",
			(SELECT json_build_object(
				'id', tsr.id, 'createdAt', tsr."createdAt"::timestamptz, 'updatedAt', tsr."updatedAt"::timestamptz,
				'trainerId', tsr."trainerId", 'studentId', tsr."studentId", 'linkedAt', tsr."linkedAt"::timestamptz,
				'trainer', json_build_object(
					'id', t.id, 'firstName', t."firstName", 'lastName', t."lastName",
					'email', t.email, 'type', t.type, 'profilePictureUrl', t."profilePictureUrl"
				)
			) FROM trainer_student_relationships tsr 
			  LEFT JOIN users t ON tsr."trainerId" = t.id 
			  WHERE tsr."studentId" = u.id AND tsr."deletedAt" IS NULL LIMIT 1) as student_of,
			COALESCE((SELECT json_agg(json_build_object(
				'id', tsr.id, 'createdAt', tsr."createdAt"::timestamptz, 'updatedAt', tsr."updatedAt"::timestamptz,
				'trainerId', tsr."trainerId", 'studentId', tsr."studentId", 'linkedAt', tsr."linkedAt"::timestamptz,
				'student', json_build_object(
					'id', s.id, 'firstName', s."firstName", 'lastName', s."lastName",
					'email', s.email, 'type', s.type, 'profilePictureUrl', s."profilePictureUrl"
				)
			)) FROM trainer_student_relationships tsr 
			   LEFT JOIN users s ON tsr."studentId" = s.id 
			   WHERE tsr."trainerId" = u.id AND tsr."deletedAt" IS NULL), '[]'::json) as trainer_of
		FROM users u WHERE id = ANY($1) AND u."deletedAt" IS NULL`
	rows, err := r.db.QueryContext(ctx, query, pq.Array(ids))
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	users := make([]*User, 0)
	for rows.Next() {
		var u User
		var studentOfJSON, trainerOfJSON []byte
		err := rows.Scan(
			&u.ID, &u.FirstName, &u.LastName, &u.Email, &u.Type, &u.CreatedAt, &u.UpdatedAt,
			&u.Bio, &u.ProfilePictureUrl, &u.Height, &u.CurrentWeight, &u.WeightUnit, &u.HeightUnit,
			&u.TrainerInviteCode, &u.Cref, &u.IsVerified,
			&studentOfJSON, &trainerOfJSON,
		)
		if err != nil {
			return nil, err
		}
		if len(studentOfJSON) > 0 {
			json.Unmarshal(studentOfJSON, &u.StudentOf)
		}
		if len(trainerOfJSON) > 0 {
			json.Unmarshal(trainerOfJSON, &u.TrainerOf)
		}
		users = append(users, &u)
	}
	return users, nil
}

func (r *PostgresUserRepository) ListStudents(ctx context.Context, trainerId string, cursor *utils.CursorData, limit int) ([]*User, *utils.CursorData, error) {
	query := `
		SELECT 
			u.id, u."firstName", u."lastName", u.email, u.type, u."createdAt", u."updatedAt",
			u.bio, u."profilePictureUrl", u.height, u."currentWeight", u."weightUnit", u."heightUnit",
			u."trainerInviteCode", u.cref, u."isVerified",
			(SELECT json_build_object(
				'id', tsr_s.id, 'createdAt', tsr_s."createdAt"::timestamptz, 'updatedAt', tsr_s."updatedAt"::timestamptz,
				'trainerId', tsr_s."trainerId", 'studentId', tsr_s."studentId", 'linkedAt', tsr_s."linkedAt"::timestamptz,
				'trainer', json_build_object(
					'id', t.id, 'firstName', t."firstName", 'lastName', t."lastName",
					'email', t.email, 'type', t.type, 'profilePictureUrl', t."profilePictureUrl"
				)
			) FROM trainer_student_relationships tsr_s 
			  LEFT JOIN users t ON tsr_s."trainerId" = t.id 
			  WHERE tsr_s."studentId" = u.id AND tsr_s."deletedAt" IS NULL LIMIT 1) as student_of,
			COALESCE((SELECT json_agg(json_build_object(
				'id', tsr_t.id, 'createdAt', tsr_t."createdAt"::timestamptz, 'updatedAt', tsr_t."updatedAt"::timestamptz,
				'trainerId', tsr_t."trainerId", 'studentId', tsr_t."studentId", 'linkedAt', tsr_t."linkedAt"::timestamptz,
				'student', json_build_object(
					'id', s.id, 'firstName', s."firstName", 'lastName', s."lastName",
					'email', s.email, 'type', s.type, 'profilePictureUrl', s."profilePictureUrl"
				)
			)) FROM trainer_student_relationships tsr_t 
			   LEFT JOIN users s ON tsr_t."studentId" = s.id 
			   WHERE tsr_t."trainerId" = u.id AND tsr_t."deletedAt" IS NULL), '[]'::json) as trainer_of
		FROM users u
		INNER JOIN trainer_student_relationships tsr ON tsr."studentId" = u.id
		WHERE tsr."trainerId" = $1 AND tsr."deletedAt" IS NULL AND u."deletedAt" IS NULL
	`

	var args []interface{}
	args = append(args, trainerId)

	if cursor != nil {
		query += ` AND (u."createdAt", u.id) < ($2, $3)`
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
		var studentOfJSON, trainerOfJSON []byte
		err := rows.Scan(
			&u.ID, &u.FirstName, &u.LastName, &u.Email, &u.Type, &u.CreatedAt, &u.UpdatedAt,
			&u.Bio, &u.ProfilePictureUrl, &u.Height, &u.CurrentWeight, &u.WeightUnit, &u.HeightUnit,
			&u.TrainerInviteCode, &u.Cref, &u.IsVerified,
			&studentOfJSON, &trainerOfJSON,
		)
		if err != nil {
			return nil, nil, err
		}
		if len(studentOfJSON) > 0 {
			json.Unmarshal(studentOfJSON, &u.StudentOf)
		}
		if len(trainerOfJSON) > 0 {
			json.Unmarshal(trainerOfJSON, &u.TrainerOf)
		}
		users = append(users, &u)
	}

	var nextCursor *utils.CursorData
	if len(users) > limit {
		nextCursor = &utils.CursorData{
			ID:        *users[limit].ID,
			CreatedAt: users[limit].CreatedAt,
		}
		users = users[:limit]
	}

	return users, nextCursor, nil
}

func (r *PostgresUserRepository) CountFollowers(ctx context.Context, userId string) (int, error) {
	query := `SELECT count(*) FROM user_follows WHERE "followingId" = $1 AND "deletedAt" IS NULL`
	var count int
	err := r.db.QueryRowContext(ctx, query, userId).Scan(&count)
	if err != nil {
		return 0, fmt.Errorf("could not count followers: %w", err)
	}
	return count, nil
}

func (r *PostgresUserRepository) CountFollowing(ctx context.Context, userId string) (int, error) {
	query := `SELECT count(*) FROM user_follows WHERE "followerId" = $1 AND "deletedAt" IS NULL`
	var count int
	err := r.db.QueryRowContext(ctx, query, userId).Scan(&count)
	if err != nil {
		return 0, fmt.Errorf("could not count following: %w", err)
	}
	return count, nil
}

func (r *PostgresUserRepository) FollowUser(ctx context.Context, f UserFollows) error {
	query := `INSERT INTO user_follows (id, "followerId", "followingId", "createdAt", "updatedAt") VALUES ($1, $2, $3, $4, $5)`
	_, err := r.db.ExecContext(ctx, query, f.ID, f.FollowerId, f.FollowingId, f.CreatedAt, f.UpdatedAt)
	if err != nil {
		return fmt.Errorf("could not create user: %w", err)
	}
	return nil
}

func (r *PostgresUserRepository) UnfollowUser(ctx context.Context, followerId, followingId string) error {
	query := `UPDATE user_follows SET "deletedAt" = $1 WHERE "followerId" = $2 AND "followingId" = $3`
	_, err := r.db.ExecContext(ctx, query, time.Now().UTC(), followerId, followingId)
	if err != nil {
		return fmt.Errorf("could not create user: %w", err)
	}
	return nil
}

func (r *PostgresUserRepository) ListFollowing(ctx context.Context, id string, cursor *utils.CursorData, limit int) ([]*User, *utils.CursorData, error) {
	query := `
		SELECT u.id, u."firstName", u."lastName", u.email, u.type, u."createdAt", u."updatedAt", u."profilePictureUrl"
		FROM user_follows f LEFT JOIN users u ON f."followingId" = u.id
		WHERE f."followerId" = $1 AND f."deletedAt" IS NULL AND u."deletedAt" IS NULL
	`
	var args []interface{}
	args = append(args, id)

	if cursor != nil {
		query += ` AND (u."createdAt", u.id) < ($2, $3)`
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
		err := rows.Scan(&u.ID, &u.FirstName, &u.LastName, &u.Email, &u.Type, &u.CreatedAt, &u.UpdatedAt, &u.ProfilePictureUrl)
		if err != nil {
			return nil, nil, err
		}
		users = append(users, &u)
	}

	var nextCursor *utils.CursorData
	if len(users) > limit {
		nextCursor = &utils.CursorData{
			ID:        *users[limit].ID,
			CreatedAt: users[limit].CreatedAt,
		}
		users = users[:limit]
	}

	return users, nextCursor, nil
}

func (r *PostgresUserRepository) ListFollower(ctx context.Context, id string, cursor *utils.CursorData, limit int) ([]*User, *utils.CursorData, error) {
	query := `
		SELECT u.id, u."firstName", u."lastName", u.email, u.type, u."createdAt", u."updatedAt", u."profilePictureUrl"
		FROM user_follows f LEFT JOIN users u ON f."followerId" = u.id
		WHERE f."followingId" = $1 AND f."deletedAt" IS NULL AND u."deletedAt" IS NULL
	`
	var args []interface{}
	args = append(args, id)

	if cursor != nil {
		query += ` AND (u."createdAt", u.id) < ($2, $3)`
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
		err := rows.Scan(&u.ID, &u.FirstName, &u.LastName, &u.Email, &u.Type, &u.CreatedAt, &u.UpdatedAt, &u.ProfilePictureUrl)
		if err != nil {
			return nil, nil, err
		}
		users = append(users, &u)
	}

	var nextCursor *utils.CursorData
	if len(users) > limit {
		nextCursor = &utils.CursorData{
			ID:        *users[limit].ID,
			CreatedAt: users[limit].CreatedAt,
		}
		users = users[:limit]
	}

	return users, nextCursor, nil
}

func (r *PostgresUserRepository) FindByTrainerCode(ctx context.Context, code string) (*User, error) {
	query := `
		SELECT 
			u.id, u."firstName", u."lastName", u.email, u.password, u.type, u."createdAt", u."updatedAt",
			u.bio, u."profilePictureUrl", u.height, u."currentWeight", u."weightUnit", u."heightUnit",
			u."trainerInviteCode", u.cref, u."isVerified",
			(SELECT json_build_object(
				'id', tsr.id, 'createdAt', tsr."createdAt"::timestamptz, 'updatedAt', tsr."updatedAt"::timestamptz,
				'trainerId', tsr."trainerId", 'studentId', tsr."studentId", 'linkedAt', tsr."linkedAt"::timestamptz,
				'trainer', json_build_object(
					'id', t.id, 'firstName', t."firstName", 'lastName', t."lastName",
					'email', t.email, 'type', t.type, 'profilePictureUrl', t."profilePictureUrl"
				)
			) FROM trainer_student_relationships tsr 
			  LEFT JOIN users t ON tsr."trainerId" = t.id 
			  WHERE tsr."studentId" = u.id AND tsr."deletedAt" IS NULL LIMIT 1) as student_of,
			COALESCE((SELECT json_agg(json_build_object(
				'id', tsr.id, 'createdAt', tsr."createdAt"::timestamptz, 'updatedAt', tsr."updatedAt"::timestamptz,
				'trainerId', tsr."trainerId", 'studentId', tsr."studentId", 'linkedAt', tsr."linkedAt"::timestamptz,
				'student', json_build_object(
					'id', s.id, 'firstName', s."firstName", 'lastName', s."lastName",
					'email', s.email, 'type', s.type, 'profilePictureUrl', s."profilePictureUrl"
				)
			)) FROM trainer_student_relationships tsr 
			   LEFT JOIN users s ON tsr."studentId" = s.id 
			   WHERE tsr."trainerId" = u.id AND tsr."deletedAt" IS NULL), '[]'::json) as trainer_of
		FROM users u WHERE "trainerInviteCode" = $1 AND u."deletedAt" IS NULL LIMIT 1`
	var u User
	var studentOfJSON, trainerOfJSON []byte
	err := r.db.QueryRowContext(ctx, query, code).Scan(
		&u.ID, &u.FirstName, &u.LastName, &u.Email, &u.Password, &u.Type, &u.CreatedAt, &u.UpdatedAt,
		&u.Bio, &u.ProfilePictureUrl, &u.Height, &u.CurrentWeight, &u.WeightUnit, &u.HeightUnit,
		&u.TrainerInviteCode, &u.Cref, &u.IsVerified,
		&studentOfJSON, &trainerOfJSON,
	)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}
	if len(studentOfJSON) > 0 {
		json.Unmarshal(studentOfJSON, &u.StudentOf)
	}
	if len(trainerOfJSON) > 0 {
		json.Unmarshal(trainerOfJSON, &u.TrainerOf)
	}
	return &u, nil
}

func (r *PostgresUserRepository) CreateTrainerCode(ctx context.Context, id, code string) error {
	query := `UPDATE users SET "trainerInviteCode" = $2 WHERE id = $1`
	_, err := r.db.ExecContext(ctx, query, id, code)
	if err != nil {
		return fmt.Errorf("could not create user: %w", err)
	}
	return nil
}

func (r *PostgresUserRepository) LinkTrainer(ctx context.Context, relation TrainerStudentRelation) error {
	query := `INSERT INTO trainer_student_relationships (id, "createdAt", "updatedAt", "trainerId", "studentId", "linkedAt") VALUES ($1, $2, $3, $4, $5, $6)`
	_, err := r.db.ExecContext(ctx, query, relation.ID, relation.CreatedAt, relation.UpdatedAt, relation.TrainerId, relation.StudentId, relation.LinkedAt)
	if err != nil {
		return fmt.Errorf("could not link trainer: %w", err)
	}
	return nil
}

func (r *PostgresUserRepository) UnlinkTrainer(ctx context.Context, studentId string) error {
	query := `UPDATE trainer_student_relationships SET "deletedAt" = NOW() WHERE "studentId" = $1 AND "deletedAt" IS NULL`
	_, err := r.db.ExecContext(ctx, query, studentId)
	if err != nil {
		return fmt.Errorf("could not unlink trainer: %w", err)
	}
	return nil
}

func (r *PostgresUserRepository) AddBodyMeasurementNote(ctx context.Context, id, note string) error {
	query := `UPDATE body_measurements SET "trainerNote" = $2, "trainerNoteAt" = $3 WHERE id = $1`
	_, err := r.db.ExecContext(ctx, query, id, note, time.Now().UTC())
	if err != nil {
		return fmt.Errorf("could not add body measurement note: %w", err)
	}
	return nil
}

func (r *PostgresUserRepository) ListBodyMeasurements(ctx context.Context, userId string, cursor *utils.CursorData, limit int) ([]*BodyMeasurement, *utils.CursorData, error) {
	query := `SELECT id, "createdAt", "updatedAt", "type", value, "measuredAt", "userId", "trainerNote", "trainerNoteAt" FROM body_measurements WHERE "userId" = $1 AND "deletedAt" IS NULL`

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
func (r *PostgresUserRepository) FindLastBodyMeasurementNote(ctx context.Context, userId string) (*BodyMeasurement, error) {
	query := `SELECT id, "createdAt", "updatedAt", "type", value, "measuredAt", "userId", "trainerNote", "trainerNoteAt" FROM body_measurements WHERE "userId" = $1 AND "deletedAt" IS NULL ORDER BY "measuredAt" DESC LIMIT 1`
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

func (r *PostgresUserRepository) AddWeightLogNote(ctx context.Context, id, note string) error {
	query := `UPDATE weight_logs SET "trainerNote" = $2, "trainerNoteAt" = $3 WHERE id = $1`
	_, err := r.db.ExecContext(ctx, query, id, note, time.Now().UTC())
	if err != nil {
		return fmt.Errorf("could not add weight log note: %w", err)
	}
	return nil
}

func (r *PostgresUserRepository) ChangeUserType(ctx context.Context, u User, newType UserType) error {
	query := `UPDATE users SET type = $2 WHERE id = $1`
	_, err := r.db.ExecContext(ctx, query, u.ID, newType)
	if err != nil {
		return fmt.Errorf("could not create user: %w", err)
	}
	return nil
}

func (r *PostgresUserRepository) RemoveProfilePicture(ctx context.Context, userId string) error {
	query := `UPDATE users SET "profilePictureUrl" = NULL WHERE id = $1`
	_, err := r.db.ExecContext(ctx, query, userId)
	if err != nil {
		return fmt.Errorf("could not remove profile picture: %w", err)
	}
	return nil
}

func (r *PostgresUserRepository) ChangeProfileImage(ctx context.Context, u User, pictureUrl string) error {
	query := `UPDATE users SET "profilePictureUrl" = $2 WHERE id = $1`
	_, err := r.db.ExecContext(ctx, query, u.ID, pictureUrl)
	if err != nil {
		return fmt.Errorf("could not create user: %w", err)
	}
	return nil
}

func (r *PostgresUserRepository) ListGoalsMetric(ctx context.Context, id string, cursor *utils.CursorData, limit int) ([]*MetricGoal, *utils.CursorData, error) {
	query := `SELECT id, "createdAt", "updatedAt", "type", "startingValue", "targetValue", deadline, "achievedAt", status, "userId" FROM metric_goals WHERE "userId" = $1 AND "deletedAt" IS NULL`

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

func (r *PostgresUserRepository) AddGoalMetric(ctx context.Context, g MetricGoal) error {
	query := `INSERT INTO metric_goals (id, "createdAt", "updatedAt", "type", "startingValue", "targetValue", deadline, "achievedAt", status, "userId") VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10)`
	_, err := r.db.ExecContext(ctx, query, g.ID, g.CreatedAt, g.UpdatedAt, g.Type, g.StartingValue, g.TargetValue, g.Deadline, g.AchievedAt, g.Status, g.UserId)
	if err != nil {
		return fmt.Errorf("could not add goal metric: %w", err)
	}
	return nil
}

func (r *PostgresUserRepository) ListWeightLogs(ctx context.Context, userId string, cursor *utils.CursorData, limit int) ([]*WeightLog, *utils.CursorData, error) {
	query := `SELECT id, "createdAt", "updatedAt", weight, "measuredAt", "userId", "trainerNote", "trainerNoteAt" FROM weight_logs WHERE "userId" = $1 AND "deletedAt" IS NULL`

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

func (r *PostgresUserRepository) SaveResetCode(ctx context.Context, code, email string) error {
	return r.cache.Set(ctx, fmt.Sprintf("reset_password:%s", email), code, time.Minute*5)
}

func (r *PostgresUserRepository) GetResetCode(ctx context.Context, email string) (string, error) {
	return r.cache.Get(ctx, fmt.Sprintf("reset_password:%s", email))
}

func (r *PostgresUserRepository) SaveVerificationCode(ctx context.Context, code, email string) error {
	return r.cache.Set(ctx, fmt.Sprintf("email_verification:%s", email), code, time.Minute*10)
}

func (r *PostgresUserRepository) GetVerificationCode(ctx context.Context, email string) (string, error) {
	return r.cache.Get(ctx, fmt.Sprintf("email_verification:%s", email))
}

func (r *PostgresUserRepository) Update(ctx context.Context, u *User) error {
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
