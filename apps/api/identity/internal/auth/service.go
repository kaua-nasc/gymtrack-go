package auth

import (
	"context"
	"fmt"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/kaua-nasc/gymtrack-go/apps/api/identity/internal/domain"
	"github.com/kaua-nasc/gymtrack-go/libs/auth"
	"github.com/kaua-nasc/gymtrack-go/libs/email"
	"github.com/kaua-nasc/gymtrack-go/libs/utils"
)

type Service struct {
	repo Repository
}

func NewService(repo Repository) *Service {
	return &Service{repo: repo}
}

func (s *Service) Register(ctx context.Context, u domain.User) error {
	if u.Email == nil {
		return domain.ErrInvalidCredentials
	}
	existing, err := s.repo.FindByEmail(ctx, *u.Email)
	if err != nil {
		return err
	}
	if existing != nil {
		return domain.ErrUserAlreadyExists
	}

	hashedPassword, err := auth.HashArgon2Password(u.Password)
	if err != nil {
		return err
	}
	u.Password = hashedPassword

	id, err := utils.GenerateUUIDV7(ctx)
	if err != nil {
		return err
	}

	now := time.Now().UTC()
	u.ID = id
	u.CreatedAt = now
	u.UpdatedAt = now
	u.Type = domain.Client
	u.WeightUnit = domain.KG
	u.HeightUnit = domain.CM

	if err := s.repo.Create(ctx, &u); err != nil {
		return err
	}

	u.Password = ""
	return nil
}

func (s *Service) Login(ctx context.Context, emailVal, password string) (string, error) {
	u, err := s.repo.FindByEmail(ctx, emailVal)
	if err != nil || u == nil {
		return "", domain.ErrInvalidCredentials
	}

	ok, err := auth.VerifyArgon2Password(password, u.Password)
	if err != nil || !ok {
		return "", domain.ErrInvalidCredentials
	}

	now := time.Now().UTC()
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub":  u.ID,
		"type": u.Type,
		"iat":  now.Unix(),
		"exp":  now.AddDate(0, 0, 30).Unix(), // 30 days for better consumer UX
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

func (s *Service) SendVerificationEmail(ctx context.Context, userEmail string) error {
	u, err := s.repo.FindByEmail(ctx, userEmail)
	if err != nil {
		return err
	}
	if u == nil {
		return domain.ErrEmailNotFound
	}

	if u.IsVerified {
		return domain.ErrAlreadyVerified
	}

	code, err := auth.GenerateCode(6)
	if err != nil {
		return err
	}

	if err := s.repo.SaveVerificationCode(ctx, code, *u.Email); err != nil {
		return err
	}

	return email.Send(*u.Email, email.EmailRequestContent{
		Subject:   "Verificação de E-mail - Gymtrack",
		PlainText: fmt.Sprintf("Seu código de verificação é: %s", code),
		HTML: fmt.Sprintf(`
			<!DOCTYPE html>
			<html lang="pt-BR">
			<head>
			<meta charset="UTF-8">
			<meta name="viewport" content="width=device-width, initial-scale=1.0">
			<title>Verificação de E-mail</title>
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
				background: linear-gradient(135deg, #10b981, #059669);
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
				<h1>Verifique seu e-mail</h1>

				<p>Olá %s,</p>

				<p>
				Obrigado por se juntar ao Gymtrack!
				Para confirmar sua conta e começar a treinar, utilize o código abaixo:
				</p>

				<div class="code-box">%s</div>

				<p>
				Este código é válido por <strong>10 minutos</strong>.
				Se você não solicitou este código, por favor ignore este email.
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

func (s *Service) VerifyEmail(ctx context.Context, userEmail string, userCode string) error {
	code, err := s.repo.GetVerificationCode(ctx, userEmail)
	if err != nil {
		return err
	}

	if userCode != code {
		return domain.ErrInvalidCode
	}

	user, err := s.repo.FindByEmail(ctx, userEmail)
	if err != nil {
		return err
	}

	if user == nil {
		return domain.ErrUserNotFound
	}

	if user.IsVerified {
		return domain.ErrAlreadyVerified
	}

	user.IsVerified = true
	user.UpdatedAt = time.Now().UTC()

	return s.repo.Update(ctx, user)
}

func (s *Service) ResetPasswordSendToken(ctx context.Context, userEmail string) error {
	u, err := s.repo.FindByEmail(ctx, userEmail)
	if err != nil {
		return err
	}
	if u == nil {
		return domain.ErrEmailNotFound
	}

	code, err := auth.GenerateCode(6)
	if err != nil {
		return err
	}

	if err := s.repo.SaveResetCode(ctx, code, *u.Email); err != nil {
		return err
	}

	return email.Send(*u.Email, email.EmailRequestContent{
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

func (s *Service) ResetPasswordVerifyToken(ctx context.Context, userEmail string, userCode string) (bool, error) {
	code, err := s.repo.GetResetCode(ctx, userEmail)
	if err != nil {
		return false, err
	}

	if userCode != code {
		return false, domain.ErrInvalidCode
	}

	return true, nil
}

func (s *Service) ResetPassword(ctx context.Context, userEmail, userCode, newPassword string) error {
	code, err := s.repo.GetResetCode(ctx, userEmail)
	if err != nil {
		return err
	}

	if userCode != code {
		return domain.ErrInvalidCode
	}

	user, err := s.repo.FindByEmail(ctx, userEmail)
	if err != nil {
		return err
	}

	if user == nil {
		return domain.ErrUserNotFound
	}

	hashedPassword, err := auth.HashArgon2Password(newPassword)
	if err != nil {
		return err
	}
	user.Password = hashedPassword

	return s.repo.Update(ctx, user)
}
