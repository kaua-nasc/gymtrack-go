package internal

import (
	"context"
	"crypto/rand"
	"crypto/subtle"
	"encoding/base64"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log/slog"
	"math/big"
	"os"
	"strings"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
	"github.com/kaua-nasc/gymtrack-go/libs/email"
	"github.com/kaua-nasc/gymtrack-go/libs/storage"
	"golang.org/x/crypto/argon2"
)

var (
	ErrUserAlreadyExists  = errors.New("user already exists")
	ErrInvalidCredentials = errors.New("invalid credentials")
	ErrUserNotFound       = errors.New("user not found")
	ErrEmailNotFound      = errors.New("email not found")
	ErrInvalidCode        = errors.New("invalid code")
	ErrTrainerNotFound    = errors.New("trainer not found")
)

type UserService struct {
	repo UserRepository
}

func NewUserService(repo UserRepository) *UserService {
	return &UserService{repo: repo}
}

func (s *UserService) Register(ctx context.Context, u User) error {
	existing, err := s.repo.FindByEmail(ctx, u.Email)
	if err != nil {
		return err
	}
	if existing != nil {
		return ErrUserAlreadyExists
	}

	hashedPassword, err := HashArgon2Password(u.Password)
	if err != nil {
		return err
	}
	u.Password = hashedPassword

	id, err := uuid.NewV7()
	if err != nil {
		slog.ErrorContext(ctx, "failed to generate uuid for day", slog.Any("error", err))
		return fmt.Errorf("error on generate uuid")
	}
	now := time.Now().UTC()
	u.ID = id.String()
	u.CreatedAt = now
	u.UpdatedAt = now
	u.Type = Client
	u.WeightUnit = KG
	u.HeightUnit = CM

	if err := s.repo.Create(ctx, &u); err != nil {
		return err
	}

	u.Password = ""
	return nil
}

func HashArgon2Password(password string) (string, error) {
	memory := uint32(64 * 1024)
	iterations := uint32(3)
	parallelism := uint8(2)
	saltLength := uint32(16)
	keyLength := uint32(32)

	salt := make([]byte, saltLength)
	if _, err := rand.Read(salt); err != nil {
		return "", err
	}

	hash := argon2.IDKey([]byte(password), salt, iterations, memory, parallelism, keyLength)

	b64Salt := base64.RawStdEncoding.EncodeToString(salt)
	b64Hash := base64.RawStdEncoding.EncodeToString(hash)

	encodedHash := fmt.Sprintf("$argon2id$v=%d$m=%d,t=%d,p=%d$%s$%s",
		argon2.Version, memory, iterations, parallelism, b64Salt, b64Hash)

	return encodedHash, nil
}

func (s *UserService) Login(ctx context.Context, email, password string) (string, error) {
	u, err := s.repo.FindByEmail(ctx, email)
	if err != nil || u == nil {
		return "", ErrInvalidCredentials
	}

	ok, err := VerifyArgon2Password(password, u.Password)

	if err != nil || !ok {
		return "", ErrInvalidCredentials
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

func (s *UserService) ResetPasswordSendToken(ctx context.Context, userEmail string) error {
	u, err := s.repo.FindByEmail(ctx, userEmail)
	if err != nil {
		return err
	}
	if u == nil {
		return ErrEmailNotFound
	}

	code, err := generateResetCode(6)
	if err != nil {
		return err
	}

	if err := s.repo.SaveResetCode(ctx, code, u.Email); err != nil {
		return err
	}

	return email.Send(u.Email, email.EmailRequestContent{
		Subject:   "Redefinição de Senha - Gymtrack",
		PlainText: fmt.Sprintf("Seu código de redefinição de senha é: %s", code),
		HTML: fmt.Sprintf(`
			<!DOCTYPE html>
			<html lang="pt-BR">
			<head>
			<meta charset="UTF-8">
			<meta name="viewport" content="width=device-width, initial-scale=1.0">
			<title>Redefinição de Senha</title>
			<style>
				body {
				font-family: Arial, sans-serif;
				background-color: #f4f8fc;
				color: #334155;
				margin: 0;
				padding: 0;
				}

				.container {
				max-width: 600px;
				margin: 40px auto;
				background-color: #ffffff;
				border-radius: 8px;
				padding: 20px;
				box-shadow: 0 4px 10px rgba(0,0,0,0.08);
				}

				h1 {
				color: #1d4ed8;
				font-size: 24px;
				margin-bottom: 10px;
				}

				p {
				font-size: 16px;
				line-height: 1.5;
				}

				.code-box {
				margin: 20px 0;
				padding: 15px;
				font-size: 22px;
				letter-spacing: 4px;
				font-weight: bold;
				text-align: center;
				color: #ffffff;
				background: linear-gradient(135deg, #2563eb, #0ea5e9);
				border-radius: 6px;
				}

				.footer {
				font-size: 12px;
				color: #94a3b8;
				text-align: center;
				margin-top: 30px;
				}

				a {
				color: #2563eb;
				text-decoration: none;
				}
			</style>
			</head>

			<body>
			<div class="container">
				<h1>Redefinição de Senha</h1>

				<p>Olá %s,</p>

				<p>
				Recebemos uma solicitação para redefinir sua senha.
				Utilize o código abaixo para prosseguir com a redefinição:
				</p>

				<div class="code-box">%s</div>

				<p>
				Este código é válido por <strong>5 minutos</strong>.
				Se você não solicitou a redefinição, por favor ignore este email.
				</p>

				<p>
				Atenciosamente,<br>
				Equipe GymTrack
				</p>

				<div class="footer">
				© 2026 GymTrack. Todos os direitos reservados.
				</div>
			</div>
			</body>
			</html>`, u.FirstName, code),
	})
}

func generateResetCode(length int) (string, error) {
	const chars = "ABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	code := make([]byte, length)

	for i := range length {
		index, err := rand.Int(rand.Reader, big.NewInt(int64(len(chars))))
		if err != nil {
			return "", err
		}
		code[i] = chars[index.Int64()]
	}

	return string(code), nil
}

func (s *UserService) ResetPasswordVerifyToken(ctx context.Context, userEmail string, userCode string) (bool, error) {
	code, err := s.repo.GetResetCode(ctx, userEmail)
	if err != nil {
		return false, err
	}

	if userCode != code {
		return false, ErrInvalidCode
	}

	return true, nil
}

func (s *UserService) ResetPassword(ctx context.Context, userEmail, userCode, newPassword string) error {
	code, err := s.repo.GetResetCode(ctx, userEmail)
	if err != nil {
		return err
	}

	if userCode != code {
		return ErrInvalidCode
	}

	user, err := s.repo.FindByEmail(ctx, userEmail)
	if err != nil {
		return err
	}

	if user == nil {
		return ErrUserNotFound
	}

	hashedPassword, err := HashArgon2Password(newPassword)
	if err != nil {
		return err
	}
	user.Password = hashedPassword

	return s.repo.Update(ctx, user)
}

func (s *UserService) UpdateProfile(
	ctx context.Context,
	id string,
	firstName, lastName, bio *string,
	height *float64,
	weightUnit *WeightUnit,
	heightUnit *HeightUnit,
	currentWeight *float64,
) error {
	u, err := s.repo.Find(ctx, id)
	if err != nil {
		return err
	}
	if u == nil {
		return ErrUserNotFound
	}

	if firstName != nil {
		u.FirstName = *firstName
	}
	if lastName != nil {
		u.LastName = *lastName
	}
	if bio != nil {
		u.Bio = bio
	}
	if height != nil {
		u.Height = height
	}
	if weightUnit != nil {
		u.WeightUnit = *weightUnit
	}
	if heightUnit != nil {
		u.HeightUnit = *heightUnit
	}
	if currentWeight != nil {
		u.CurrentWeight = currentWeight
	}

	u.UpdatedAt = time.Now().UTC()

	return s.repo.Update(ctx, u)
}

func (s *UserService) GetUser(ctx context.Context, id string) (*User, error) {
	u, err := s.repo.Find(ctx, id)
	if err != nil || u == nil {
		return u, err
	}
	s.sanitizeUser(u)
	return u, nil
}

func (s *UserService) ListUsers(ctx context.Context, ids []string) ([]*User, error) {
	users, err := s.repo.ListByIDs(ctx, ids)
	if err != nil {
		return nil, err
	}
	for _, u := range users {
		s.sanitizeUser(u)
	}
	return users, nil
}

func (s *UserService) sanitizeUser(u *User) {
	if u == nil {
		return
	}
	u.Password = ""
	if u.ProfilePictureUrl != nil && *u.ProfilePictureUrl != "" && !strings.HasPrefix(*u.ProfilePictureUrl, "http") {
		uri := os.Getenv("AZURE_STORAGE_URL")
		fullUrl := uri + "/" + *u.ProfilePictureUrl
		u.ProfilePictureUrl = &fullUrl
	}

	if u.StudentOf != nil && u.StudentOf.Trainer != nil {
		s.sanitizeUser(u.StudentOf.Trainer)
	}

	for i := range u.TrainerOf {
		if u.TrainerOf[i].Student != nil {
			s.sanitizeUser(u.TrainerOf[i].Student)
		}
	}
}

func VerifyArgon2Password(password, encodedHash string) (bool, error) {
	parts := strings.Split(encodedHash, "$")
	if len(parts) != 6 {
		return false, errors.New("invalid hash format")
	}

	var memory, iterations uint32
	var parallelism uint8
	_, err := fmt.Sscanf(parts[3], "m=%d,t=%d,p=%d", &memory, &iterations, &parallelism)
	if err != nil {
		return false, err
	}

	salt, err := base64.RawStdEncoding.DecodeString(parts[4])
	if err != nil {
		return false, err
	}

	decodedHash, err := base64.RawStdEncoding.DecodeString(parts[5])
	if err != nil {
		return false, err
	}

	comparisonHash := argon2.IDKey([]byte(password), salt, iterations, memory, parallelism, uint32(len(decodedHash)))

	return subtle.ConstantTimeCompare(decodedHash, comparisonHash) == 1, nil
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

func (s *UserService) CountFollowers(ctx context.Context, id string) (int, error) {
	count, err := s.repo.CountFollowers(ctx, id)
	if err != nil {
		return 0, err
	}

	return count, nil
}

func (s *UserService) CountFollowing(ctx context.Context, id string) (int, error) {
	count, err := s.repo.CountFollowing(ctx, id)
	if err != nil {
		return 0, err
	}

	return count, nil
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
	user, err := s.repo.Find(ctx, id)
	if err != nil {
		return err
	}
	if user == nil {
		return ErrTrainerNotFound
	}

	return s.repo.CreateTrainerCode(ctx, id, code)
}

func (s *UserService) LinkTrainer(ctx context.Context, id, code string) error {
	trainer, err := s.repo.FindByTrainerCode(ctx, code)
	if err != nil {
		return err
	}
	if trainer == nil {
		return ErrTrainerNotFound
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

func (s *UserService) UnlinkStudent(ctx context.Context, studentId string) error {
	return s.repo.UnlinkTrainer(ctx, studentId)
}

func (s *UserService) ListStudents(ctx context.Context, trainerId, cursor string, limit int) ([]*User, string, error) {
	var decodedCursor *CursorData
	if cursor != "" {
		b, err := base64.StdEncoding.DecodeString(cursor)
		if err == nil {
			json.Unmarshal(b, &decodedCursor)
		}
	}

	users, rawNextCursor, err := s.repo.ListStudents(ctx, trainerId, decodedCursor, limit)
	if err != nil {
		return nil, "", err
	}

	for _, u := range users {
		s.sanitizeUser(u)
	}

	var nextCursorStr string
	if rawNextCursor != nil {
		b, _ := json.Marshal(rawNextCursor)
		nextCursorStr = base64.StdEncoding.EncodeToString(b)
	}

	return users, nextCursorStr, nil
}

func (s *UserService) AddBodyMeasurementNote(ctx context.Context, id, note string) error {
	return s.repo.AddBodyMeasurementNote(ctx, id, note)
}

func (s *UserService) FindLastBodyMeasurementNote(ctx context.Context, userId string) (*BodyMeasurement, error) {
	return s.repo.FindLastBodyMeasurementNote(ctx, userId)
}

func (s *UserService) ListBodyMeasurements(ctx context.Context, userId, cursor string, limit int) ([]*BodyMeasurement, string, error) {
	var decodedCursor *CursorData
	if cursor != "" {
		b, err := base64.StdEncoding.DecodeString(cursor)
		if err == nil {
			json.Unmarshal(b, &decodedCursor)
		}
	}

	measurements, rawNextCursor, err := s.repo.ListBodyMeasurements(ctx, userId, decodedCursor, limit)
	if err != nil {
		return nil, "", err
	}

	var nextCursorStr string
	if rawNextCursor != nil {
		b, _ := json.Marshal(rawNextCursor)
		nextCursorStr = base64.StdEncoding.EncodeToString(b)
	}

	return measurements, nextCursorStr, nil
}

func (s *UserService) AddWeightLogNote(ctx context.Context, id, note string) error {
	return s.repo.AddWeightLogNote(ctx, id, note)
}

func (s *UserService) ChangeToTrainer(ctx context.Context, id, cref string) error {
	user, err := s.repo.Find(ctx, id)
	if err != nil {
		return err
	}
	if user == nil {
		return ErrUserNotFound
	}

	return s.repo.ChangeUserType(ctx, *user, Trainer)
}

func (s *UserService) ChangeToClient(ctx context.Context, id string) error {
	user, err := s.repo.Find(ctx, id)
	if err != nil {
		return err
	}
	if user == nil {
		return ErrUserNotFound
	}

	return s.repo.ChangeUserType(ctx, *user, Client)
}

func (s *UserService) RemoveProfilePicture(ctx context.Context, id string) error {
	return s.repo.RemoveProfilePicture(ctx, id)
}

func (s *UserService) UploadProfilePicture(ctx context.Context, id string, file io.Reader) error {
	user, err := s.repo.Find(ctx, id)
	if err != nil {
		return err
	}

	if user == nil {
		return ErrUserNotFound
	}

	bytes, err := io.ReadAll(file)
	if err != nil {
		return err
	}

	timestamp := strings.ReplaceAll(time.Now().UTC().String(), ":", "-")
	timestamp = strings.ReplaceAll(timestamp, ".", "-")
	filename := `identity/user/profile/user-` + id + `_` + timestamp + `.png`

	if err := storage.UploadBuffer(ctx, filename, bytes); err != nil {
		return fmt.Errorf("falha no upload: %w", err)
	}

	return s.repo.ChangeProfileImage(ctx, *user, filename)
}

func (s *UserService) ListGoalsMetric(ctx context.Context, userId, cursor string, limit int) ([]*MetricGoal, string, error) {
	var decodedCursor *CursorData
	if cursor != "" {
		b, err := base64.StdEncoding.DecodeString(cursor)
		if err == nil {
			json.Unmarshal(b, &decodedCursor)
		}
	}

	goals, rawNextCursor, err := s.repo.ListGoalsMetric(ctx, userId, decodedCursor, limit)
	if err != nil {
		return nil, "", err
	}

	var nextCursorStr string
	if rawNextCursor != nil {
		b, _ := json.Marshal(rawNextCursor)
		nextCursorStr = base64.StdEncoding.EncodeToString(b)
	}

	return goals, nextCursorStr, nil
}

func (s *UserService) AddGoalMetric(ctx context.Context, goal *MetricGoal) error {
	now := time.Now().UTC()
	createdId, err := uuid.NewV7()
	if err != nil {
		slog.ErrorContext(ctx, "failed to generate uuid for day", slog.Any("error", err))
		return fmt.Errorf("error on generate uuid")
	}

	goal.ID = createdId.String()
	goal.CreatedAt = now
	goal.UpdatedAt = now
	goal.Status = MetricGoalActive

	return s.repo.AddGoalMetric(ctx, *goal)
}

func (s *UserService) ListWeightLogs(ctx context.Context, userId, cursor string, limit int) ([]*WeightLog, string, error) {
	slog.InfoContext(ctx, "listing weight logs", slog.String("user_id", userId), slog.Int("limit", limit))

	var decodedCursor *CursorData
	if cursor != "" {
		b, err := base64.StdEncoding.DecodeString(cursor)
		if err == nil {
			json.Unmarshal(b, &decodedCursor)
		} else {
			slog.WarnContext(ctx, "failed to decode cursor for weight history", slog.String("cursor", cursor), slog.Any("error", err))
		}
	}

	logs, rawNextCursor, err := s.repo.ListWeightLogs(ctx, userId, decodedCursor, limit)
	if err != nil {
		slog.ErrorContext(ctx, "failed to list weight history", slog.String("user_id", userId), slog.Any("error", err))
		return nil, "", err
	}

	// Encode cursor
	var nextCursorStr string
	if rawNextCursor != nil {
		b, _ := json.Marshal(rawNextCursor)
		nextCursorStr = base64.StdEncoding.EncodeToString(b)
	}

	return logs, nextCursorStr, nil
}
