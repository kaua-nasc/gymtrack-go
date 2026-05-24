package auth

import (
	"context"
	"database/sql"
	"encoding/json"
	"fmt"
	"time"

	"github.com/kaua-nasc/gymtrack-go/apps/api/identity/internal/domain"
	"github.com/kaua-nasc/gymtrack-go/libs/cache"
)

//go:generate go run go.uber.org/mock/mockgen -source=repository.go -destination=mock_repository.go -package=auth
type Repository interface {
	FindByEmail(ctx context.Context, email string) (*domain.User, error)
	Create(ctx context.Context, u *domain.User) error
	SaveResetCode(ctx context.Context, code, email string) error
	GetResetCode(ctx context.Context, email string) (string, error)
	SaveVerificationCode(ctx context.Context, code, email string) error
	GetVerificationCode(ctx context.Context, email string) (string, error)
	Update(ctx context.Context, u *domain.User) error
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

func (r *PostgresRepository) Create(ctx context.Context, u *domain.User) error {
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

func (r *PostgresRepository) FindByEmail(ctx context.Context, email string) (*domain.User, error) {
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
	var u domain.User
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

func (r *PostgresRepository) SaveResetCode(ctx context.Context, code, email string) error {
	return r.cache.Set(ctx, fmt.Sprintf("reset_password:%s", email), code, time.Minute*5)
}

func (r *PostgresRepository) GetResetCode(ctx context.Context, email string) (string, error) {
	return r.cache.Get(ctx, fmt.Sprintf("reset_password:%s", email))
}

func (r *PostgresRepository) SaveVerificationCode(ctx context.Context, code, email string) error {
	return r.cache.Set(ctx, fmt.Sprintf("email_verification:%s", email), code, time.Minute*10)
}

func (r *PostgresRepository) GetVerificationCode(ctx context.Context, email string) (string, error) {
	return r.cache.Get(ctx, fmt.Sprintf("email_verification:%s", email))
}

func (r *PostgresRepository) Update(ctx context.Context, u *domain.User) error {
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
