package trainer

import (
	"context"
	"database/sql"
	"encoding/json"
	"fmt"
	"time"

	"github.com/kaua-nasc/gymtrack-go/apps/api/identity/internal/domain"
	"github.com/kaua-nasc/gymtrack-go/libs/cache"
	"github.com/kaua-nasc/gymtrack-go/libs/utils"
)

//go:generate go run go.uber.org/mock/mockgen -source=repository.go -destination=mock_repository.go -package=trainer
type Repository interface {
	Find(ctx context.Context, id string, currentUserId string) (*domain.User, error)
	CreateTrainerCode(ctx context.Context, id, code string) error
	FindByTrainerCode(ctx context.Context, code string) (*domain.User, error)
	LinkTrainer(ctx context.Context, relation domain.TrainerStudentRelation) error
	UnlinkTrainer(ctx context.Context, studentId string) error
	ListStudents(ctx context.Context, trainerId string, cursor *utils.CursorData, limit int) ([]*domain.User, *utils.CursorData, error)
	GetTrainerLinkDate(ctx context.Context, trainerId, studentId string) (*time.Time, error)
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

func (r *PostgresRepository) Find(ctx context.Context, id string, currentUserId string) (*domain.User, error) {
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
	var u domain.User
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

func (r *PostgresRepository) FindByTrainerCode(ctx context.Context, code string) (*domain.User, error) {
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
	var u domain.User
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

func (r *PostgresRepository) CreateTrainerCode(ctx context.Context, id, code string) error {
	query := `UPDATE users SET "trainerInviteCode" = $2 WHERE id = $1`
	_, err := r.db.ExecContext(ctx, query, id, code)
	if err != nil {
		return fmt.Errorf("could not create trainer code: %w", err)
	}
	return nil
}

func (r *PostgresRepository) LinkTrainer(ctx context.Context, relation domain.TrainerStudentRelation) error {
	query := `INSERT INTO trainer_student_relationships (id, "createdAt", "updatedAt", "trainerId", "studentId", "linkedAt") VALUES ($1, $2, $3, $4, $5, $6)`
	_, err := r.db.ExecContext(ctx, query, relation.ID, relation.CreatedAt, relation.UpdatedAt, relation.TrainerId, relation.StudentId, relation.LinkedAt)
	if err != nil {
		return fmt.Errorf("could not link trainer: %w", err)
	}
	return nil
}

func (r *PostgresRepository) UnlinkTrainer(ctx context.Context, studentId string) error {
	query := `UPDATE trainer_student_relationships SET "deletedAt" = NOW() WHERE "studentId" = $1 AND "deletedAt" IS NULL`
	_, err := r.db.ExecContext(ctx, query, studentId)
	if err != nil {
		return fmt.Errorf("could not unlink trainer: %w", err)
	}
	return nil
}

func (r *PostgresRepository) ListStudents(ctx context.Context, trainerId string, cursor *utils.CursorData, limit int) ([]*domain.User, *utils.CursorData, error) {
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

	users := make([]*domain.User, 0)
	for rows.Next() {
		var u domain.User
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

func (r *PostgresRepository) GetTrainerLinkDate(ctx context.Context, trainerId, studentId string) (*time.Time, error) {
	query := `SELECT "linkedAt" FROM trainer_student_relationships WHERE "trainerId" = $1 AND "studentId" = $2 AND "deletedAt" IS NULL LIMIT 1`
	var linkedAt time.Time
	err := r.db.QueryRowContext(ctx, query, trainerId, studentId).Scan(&linkedAt)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}
	return &linkedAt, nil
}
