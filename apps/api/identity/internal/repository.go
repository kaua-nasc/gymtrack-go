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
	query := `INSERT INTO trainer_student_relationships (id, "followerId", "followingId", "createdAt", "updatedAt") VALUES ($1, $2, $3, $4, $5)`
	_, err := r.db.ExecContext(ctx, query, relation.ID, relation.CreatedAt, relation.UpdatedAt, relation.LinkedAt, relation.StudentId, relation.TrainerId)
	if err != nil {
		return fmt.Errorf("could not create user: %w", err)
	}
	return nil
}

func (r *UserRepository) UnlinkTrainer(ctx context.Context, studentId string) error {
	query := `DELETE FROM trainer_student_relationships WHERE "studentId" = $1`
	_, err := r.db.ExecContext(ctx, query, studentId)
	if err != nil {
		return fmt.Errorf("could not create user: %w", err)
	}
	return nil
}

func (r *UserRepository) AddBodyMeasurementNote(ctx context.Context, id, note string) error {
	query := `UPDATE a SET trainerNote = $2, TrainerNoteAt = $3 WHERE id = $1`
	_, err := r.db.ExecContext(ctx, query, id, note, time.Now().UTC())
	if err != nil {
		return fmt.Errorf("could not create user: %w", err)
	}
	return nil
}

func (r *UserRepository) AddWeightLogNote(ctx context.Context, id, note string) error {
	query := `UPDATE a SET trainerNote = $2, TrainerNoteAt = $3 WHERE id = $1`
	_, err := r.db.ExecContext(ctx, query, id, note, time.Now().UTC())
	if err != nil {
		return fmt.Errorf("could not create user: %w", err)
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
