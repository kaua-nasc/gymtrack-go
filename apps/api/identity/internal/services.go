package internal

import (
	"context"
	"crypto/subtle"
	"encoding/base64"
	"errors"
	"fmt"
	"log/slog"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
	"golang.org/x/crypto/scrypt"
)

type UserService struct {
	repo *UserRepository
}

func NewUserService(repo *UserRepository) *UserService {
	return &UserService{repo: repo}
}

func (s *UserService) Register(ctx context.Context, u User) (*User, error) {
	existing, err := s.repo.FindByEmail(ctx, u.Email)
	if err != nil {
		return nil, err
	}
	if existing != nil {
		return nil, errors.New("user already exists")
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(u.Password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}
	u.Password = string(hashedPassword)

	now := time.Now().UTC()
	u.CreatedAt = now
	u.UpdatedAt = now

	if err := s.repo.Create(ctx, &u); err != nil {
		return nil, err
	}

	u.Password = "" // Clear password before returning
	return &u, nil
}

func (s *UserService) Login(ctx context.Context, email, password string) (string, error) {
	u, err := s.repo.FindByEmail(ctx, email)
	if err != nil || u == nil {
		return "", errors.New("invalid credentials")
	}

	ok, err := VerifyScryptPassword(password, u.Password)
	if err != nil || !ok {
		return "", errors.New("invalid credentials")
	}

	now := time.Now().UTC()
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub":  u.ID,
		"type": u.Type,
		"iat":  now.Unix(),
		"exp":  now.Add(time.Minute * 60).Unix(), // Aligned with NestJS 60m
	})

	secret := os.Getenv("JWT_SECRET")
	if secret == "" {
		secret = "default_secret" // In production, this must be set
	}

	tokenString, err := token.SignedString([]byte(secret))
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func (s *UserService) GetUser(ctx context.Context, id string) (*User, error) {
	u, err := s.repo.FindByID(ctx, id)
	if err != nil || u == nil {
		return u, err
	}
	u.Password = ""
	return u, nil
}

func (s *UserService) ListUsers(ctx context.Context, ids []string) ([]*User, error) {
	users, err := s.repo.ListByIDs(ctx, ids)
	if err != nil {
		return nil, err
	}
	for _, u := range users {
		u.Password = ""
	}
	return users, nil
}

func VerifyScryptPassword(plainPassword, stored string) (bool, error) {
	parts := strings.Split(stored, "$")
	if len(parts) != 6 {
		return false, errors.New("invalid hash format")
	}

	if parts[0] != "scrypt" {
		return false, errors.New("unsupported hash type")
	}

	N, err := strconv.Atoi(parts[1])
	if err != nil {
		return false, err
	}

	r, err := strconv.Atoi(parts[2])
	if err != nil {
		return false, err
	}

	p, err := strconv.Atoi(parts[3])
	if err != nil {
		return false, err
	}

	salt := []byte(parts[4])

	expectedHash, err := base64.RawURLEncoding.DecodeString(parts[5])
	if err != nil {
		return false, err
	}

	derivedKey, err := scrypt.Key(
		[]byte(plainPassword),
		salt,
		N,
		r,
		p,
		len(expectedHash),
	)
	if err != nil {
		return false, err
	}

	if len(derivedKey) != len(expectedHash) {
		return false, nil
	}

	return subtle.ConstantTimeCompare(derivedKey, expectedHash) == 1, nil
}

func (s *UserService) ListFollowing(ctx context.Context, id string) ([]*UserFollows, error) {
	follows, err := s.repo.ListFollowing(ctx, id)
	if err != nil {
		return nil, err
	}

	return follows, nil
}

func (s *UserService) ListFollower(ctx context.Context, id string) ([]*UserFollows, error) {
	follows, err := s.repo.ListFollower(ctx, id)
	if err != nil {
		return nil, err
	}

	return follows, nil
}

func (s *UserService) FollowUser(ctx context.Context, followerId, followingId string) error {
	users, err := s.repo.ListByIDs(ctx, []string{followerId, followingId})

	var follower, following *User
	for _, user := range users {
		if user.ID == followerId {
			follower = user
		} else {
			following = user
		}
	}

	if following == nil {
		return fmt.Errorf("usuario nao existe")
	}

	id, err := uuid.NewV7()
	if err != nil {
		slog.ErrorContext(ctx, "failed to generate uuid for comment", slog.Any("error", err))
		return fmt.Errorf("error on generate uuid")
	}

	now := time.Now().UTC()
	follow := &UserFollows{
		ID:          id.String(),
		CreatedAt:   now,
		UpdatedAt:   now,
		FollowerId:  follower.ID,
		FollowingId: following.ID,
	}

	return s.repo.FollowUser(ctx, *follow)
}

func (s *UserService) UnfollowUser(ctx context.Context, followerId, followingId string) error {
	return s.repo.UnfollowUser(ctx, followerId, followingId)
}

func (s *UserService) CreateTrainerCode(ctx context.Context, id, code string) error {
	user, err := s.repo.FindByID(ctx, code)
	if err != nil {
		return fmt.Errorf("erro")
	}
	if user == nil {
		return fmt.Errorf("trainer not found")
	}

	return s.repo.CreateTrainerCode(ctx, id, code)
}

func (s *UserService) LinkTrainer(ctx context.Context, id, code string) error {
	trainer, err := s.repo.FindByTrainerCode(ctx, code)
	if err != nil {
		return fmt.Errorf("erro")
	}
	if trainer == nil {
		return fmt.Errorf("trainer not found")
	}

	now := time.Now().UTC()
	createdId, err := uuid.NewV7()
	if err != nil {
		slog.ErrorContext(ctx, "failed to generate uuid for day", slog.Any("error", err))
		return fmt.Errorf("error on generate uuid")
	}

	relation := &TrainerStudentRelation{
		ID:        createdId.String(),
		CreatedAt: now,
		UpdatedAt: now,
		TrainerId: trainer.ID,
		StudentId: id,
		LinkedAt:  now,
	}

	return s.repo.LinkTrainer(ctx, *relation)
}

func (s *UserService) UnlinkTrainer(ctx context.Context, studentId string) error {
	return s.repo.UnlinkTrainer(ctx, studentId)
}

func (s *UserService) UnlinkStudant(ctx context.Context, studentId string) error {
	return s.repo.UnlinkTrainer(ctx, studentId)
}
