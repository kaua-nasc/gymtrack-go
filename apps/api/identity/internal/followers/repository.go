package followers

import (
	"context"
	"database/sql"
	"encoding/json"
	"fmt"
	"time"

	"github.com/kaua-nasc/gymtrack-go/apps/api/identity/internal/domain"
	"github.com/kaua-nasc/gymtrack-go/libs/cache"
	"github.com/kaua-nasc/gymtrack-go/libs/utils"
	"github.com/lib/pq"
)

//go:generate go run go.uber.org/mock/mockgen -source=repository.go -destination=mock_repository.go -package=followers
type Repository interface {
	ListByIDs(ctx context.Context, ids []string) ([]*domain.User, error)
	FollowUser(ctx context.Context, f domain.UserFollows) error
	UnfollowUser(ctx context.Context, followerId, followingId string) error
	CountFollowers(ctx context.Context, userId string) (int, error)
	CountFollowing(ctx context.Context, userId string) (int, error)
	ListFollower(ctx context.Context, id string, cursor *utils.CursorData, limit int) ([]*domain.User, *utils.CursorData, error)
	ListFollowing(ctx context.Context, id string, cursor *utils.CursorData, limit int) ([]*domain.User, *utils.CursorData, error)
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

func (r *PostgresRepository) FollowUser(ctx context.Context, f domain.UserFollows) error {
	query := `INSERT INTO user_follows (id, "followerId", "followingId", "createdAt", "updatedAt") VALUES ($1, $2, $3, $4, $5)`
	_, err := r.db.ExecContext(ctx, query, f.ID, f.FollowerId, f.FollowingId, f.CreatedAt, f.UpdatedAt)
	if err != nil {
		return fmt.Errorf("could not follow user: %w", err)
	}
	return nil
}

func (r *PostgresRepository) UnfollowUser(ctx context.Context, followerId, followingId string) error {
	query := `UPDATE user_follows SET "deletedAt" = $1 WHERE "followerId" = $2 AND "followingId" = $3`
	_, err := r.db.ExecContext(ctx, query, time.Now().UTC(), followerId, followingId)
	if err != nil {
		return fmt.Errorf("could not unfollow user: %w", err)
	}
	return nil
}

func (r *PostgresRepository) CountFollowers(ctx context.Context, userId string) (int, error) {
	query := `SELECT count(*) FROM user_follows WHERE "followingId" = $1 AND "deletedAt" IS NULL`
	var count int
	err := r.db.QueryRowContext(ctx, query, userId).Scan(&count)
	if err != nil {
		return 0, fmt.Errorf("could not count followers: %w", err)
	}
	return count, nil
}

func (r *PostgresRepository) CountFollowing(ctx context.Context, userId string) (int, error) {
	query := `SELECT count(*) FROM user_follows WHERE "followerId" = $1 AND "deletedAt" IS NULL`
	var count int
	err := r.db.QueryRowContext(ctx, query, userId).Scan(&count)
	if err != nil {
		return 0, fmt.Errorf("could not count following: %w", err)
	}
	return count, nil
}

func (r *PostgresRepository) ListFollower(ctx context.Context, id string, cursor *utils.CursorData, limit int) ([]*domain.User, *utils.CursorData, error) {
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

	users := make([]*domain.User, 0)
	for rows.Next() {
		var u domain.User
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

func (r *PostgresRepository) ListFollowing(ctx context.Context, id string, cursor *utils.CursorData, limit int) ([]*domain.User, *utils.CursorData, error) {
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

	users := make([]*domain.User, 0)
	for rows.Next() {
		var u domain.User
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
