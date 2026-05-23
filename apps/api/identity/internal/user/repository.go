package user

import (
	"context"
	"database/sql"
	"encoding/json"
	"fmt"

	"github.com/kaua-nasc/gymtrack-go/apps/api/identity/internal/domain"
	"github.com/lib/pq"
)

type Repository interface {
	Find(ctx context.Context, id string, currentUserId string) (*domain.User, error)
	Update(ctx context.Context, u *domain.User) error
	ListByIDs(ctx context.Context, ids []string) ([]*domain.User, error)
	ChangeUserType(ctx context.Context, u domain.User, newType domain.UserType) error
	RemoveProfilePicture(ctx context.Context, userId string) error
	ChangeProfileImage(ctx context.Context, u domain.User, pictureUrl string) error
	GetPrivacySettings(ctx context.Context, userId string) (*domain.UserPrivacySettings, error)
	UpsertPrivacySettings(ctx context.Context, settings domain.UserPrivacySettings) error
}

type PostgresRepository struct {
	db *sql.DB
}

func NewRepository(database *sql.DB) Repository {
	return &PostgresRepository{
		db: database,
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

func (r *PostgresRepository) ListByIDs(ctx context.Context, ids []string) ([]*domain.User, error) {
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

func (r *PostgresRepository) ChangeUserType(ctx context.Context, u domain.User, newType domain.UserType) error {
	query := `UPDATE users SET type = $2 WHERE id = $1`
	_, err := r.db.ExecContext(ctx, query, u.ID, newType)
	if err != nil {
		return fmt.Errorf("could not update type: %w", err)
	}
	return nil
}

func (r *PostgresRepository) RemoveProfilePicture(ctx context.Context, userId string) error {
	query := `UPDATE users SET "profilePictureUrl" = NULL WHERE id = $1`
	_, err := r.db.ExecContext(ctx, query, userId)
	if err != nil {
		return fmt.Errorf("could not remove profile picture: %w", err)
	}
	return nil
}

func (r *PostgresRepository) ChangeProfileImage(ctx context.Context, u domain.User, pictureUrl string) error {
	query := `UPDATE users SET "profilePictureUrl" = $2 WHERE id = $1`
	_, err := r.db.ExecContext(ctx, query, u.ID, pictureUrl)
	if err != nil {
		return fmt.Errorf("could not change profile image: %w", err)
	}
	return nil
}

func (r *PostgresRepository) GetPrivacySettings(ctx context.Context, userId string) (*domain.UserPrivacySettings, error) {
	query := `
		SELECT 
			id, "createdAt", "updatedAt", "shareEmail", "shareTrainingProgress", 
			"sharePastDataWithTrainer", "shareBodyMeasurements", "shareWeightLogs", 
			"shareMetricGoals", "allowTrainerNotes", "userId"
		FROM user_privacy_settings 
		WHERE "userId" = $1 AND "deletedAt" IS NULL`

	var s domain.UserPrivacySettings
	err := r.db.QueryRowContext(ctx, query, userId).Scan(
		&s.ID, &s.CreatedAt, &s.UpdatedAt, &s.ShareEmail, &s.ShareTrainingProgress,
		&s.SharePastDataWithTrainer, &s.ShareBodyMeasurements, &s.ShareWeightLogs,
		&s.ShareMetricGoals, &s.AllowTrainerNotes, &s.UserId,
	)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}

	return &s, nil
}

func (r *PostgresRepository) UpsertPrivacySettings(ctx context.Context, s domain.UserPrivacySettings) error {
	query := `
		INSERT INTO user_privacy_settings (
			"shareEmail", "shareTrainingProgress", "sharePastDataWithTrainer", 
			"shareBodyMeasurements", "shareWeightLogs", "shareMetricGoals", 
			"allowTrainerNotes", "userId", "updatedAt", "deletedAt"
		) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, NOW(), NULL)
		ON CONFLICT ("userId") DO UPDATE SET
			"shareEmail" = EXCLUDED."shareEmail",
			"shareTrainingProgress" = EXCLUDED."shareTrainingProgress",
			"sharePastDataWithTrainer" = EXCLUDED."sharePastDataWithTrainer",
			"shareBodyMeasurements" = EXCLUDED."shareBodyMeasurements",
			"shareWeightLogs" = EXCLUDED."shareWeightLogs",
			"shareMetricGoals" = EXCLUDED."shareMetricGoals",
			"allowTrainerNotes" = EXCLUDED."allowTrainerNotes",
			"updatedAt" = NOW(),
			"deletedAt" = NULL`

	_, err := r.db.ExecContext(ctx, query,
		s.ShareEmail, s.ShareTrainingProgress, s.SharePastDataWithTrainer,
		s.ShareBodyMeasurements, s.ShareWeightLogs, s.ShareMetricGoals,
		s.AllowTrainerNotes, s.UserId,
	)

	return err
}
