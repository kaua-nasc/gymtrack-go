package auth

import (
	"context"
	"testing"

	"github.com/kaua-nasc/gymtrack-go/apps/api/identity/internal/domain"
	"github.com/kaua-nasc/gymtrack-go/libs/auth"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	gomock "go.uber.org/mock/gomock"
)

func TestService_Register(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	mockRepo := NewMockRepository(ctrl)
	service := NewService(mockRepo)

	tests := []struct {
		name          string
		inputUser     domain.User
		mockBehavior  func()
		wantErr       bool
		expectedError string
	}{
		{
			name: "Success registration",
			inputUser: domain.User{
				Email:     new("newuser@example.com"),
				Password:  "strongpassword123",
				FirstName: "John",
				LastName:  "Doe",
			},
			mockBehavior: func() {
				mockRepo.EXPECT().FindByEmail(gomock.Any(), "newuser@example.com").Return(nil, nil).Times(1)
				mockRepo.EXPECT().Create(gomock.Any(), gomock.Any()).Return(nil).Times(1)
			},
			wantErr: false,
		},
		{
			name: "Error: User already exists",
			inputUser: domain.User{
				Email: new("existing@example.com"),
			},
			mockBehavior: func() {
				userId := "123"
				mockRepo.EXPECT().FindByEmail(gomock.Any(), "existing@example.com").Return(&domain.User{ID: &userId}, nil).Times(1)
			},
			wantErr:       true,
			expectedError: "user already exists",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.mockBehavior()
			err := service.Register(context.Background(), tt.inputUser)

			if tt.wantErr {
				assert.Error(t, err)
				assert.Equal(t, tt.expectedError, err.Error())
			} else {
				assert.NoError(t, err)
			}
		})
	}
}

func TestService_SendVerificationEmail(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	mockRepo := NewMockRepository(ctrl)
	service := NewService(mockRepo)

	tests := []struct {
		name          string
		email         string
		mockBehavior  func()
		wantErr       bool
		expectedError string
	}{
		{
			name:  "Success send verification email",
			email: "user@example.com",
			mockBehavior: func() {
				mockRepo.EXPECT().FindByEmail(gomock.Any(), "user@example.com").Return(&domain.User{Email: new("user@example.com"), FirstName: "John", IsVerified: false}, nil).Times(1)
				mockRepo.EXPECT().SaveVerificationCode(gomock.Any(), gomock.Any(), gomock.Any()).Return(nil).Times(1)
			},
			wantErr: false,
		},
		{
			name:  "Error: User not found",
			email: "nonexistent@example.com",
			mockBehavior: func() {
				mockRepo.EXPECT().FindByEmail(gomock.Any(), "nonexistent@example.com").Return(nil, nil).Times(1)
			},
			wantErr:       true,
			expectedError: "email not found",
		},
		{
			name:  "Error: Already verified",
			email: "verified@example.com",
			mockBehavior: func() {
				mockRepo.EXPECT().FindByEmail(gomock.Any(), "verified@example.com").Return(&domain.User{Email: new("verified@example.com"), IsVerified: true}, nil).Times(1)
			},
			wantErr:       true,
			expectedError: "email already verified",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.mockBehavior()

			t.Setenv("AZURE_EMAIL_CONNECTION_STRING", "Endpoint=https://test.com;AccessKey=dGVzdA==")
			t.Setenv("AZURE_EMAIL_SENDER_ADDRESS", "test@test.com")

			ctx := context.Background()
			err := service.SendVerificationEmail(ctx, tt.email)

			if tt.wantErr {
				assert.Error(t, err)
				assert.Equal(t, tt.expectedError, err.Error())
			} else {
				assert.Error(t, err)
				assert.Contains(t, err.Error(), "failed to send email request")
			}
		})
	}
}

func TestService_VerifyEmail(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	mockRepo := NewMockRepository(ctrl)
	service := NewService(mockRepo)

	tests := []struct {
		name          string
		email         string
		code          string
		mockBehavior  func()
		wantErr       bool
		expectedError string
	}{
		{
			name:  "Success verify email",
			email: "user@example.com",
			code:  "123456",
			mockBehavior: func() {
				mockRepo.EXPECT().GetVerificationCode(gomock.Any(), "user@example.com").Return("123456", nil).Times(1)
				mockRepo.EXPECT().FindByEmail(gomock.Any(), "user@example.com").Return(&domain.User{Email: new("user@example.com"), IsVerified: false}, nil).Times(1)
				mockRepo.EXPECT().Update(gomock.Any(), mock.MatchedBy(func(u *domain.User) bool {
					return u.IsVerified == true
				})).Return(nil).Times(1)
			},
			wantErr: false,
		},
		{
			name:  "Error: Invalid code",
			email: "user@example.com",
			code:  "wrong",
			mockBehavior: func() {
				mockRepo.EXPECT().GetVerificationCode(gomock.Any(), "user@example.com").Return("123456", nil).Times(1)
			},
			wantErr:       true,
			expectedError: "invalid code",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.mockBehavior()
			err := service.VerifyEmail(context.Background(), tt.email, tt.code)

			if tt.wantErr {
				assert.Error(t, err)
				assert.Equal(t, tt.expectedError, err.Error())
			} else {
				assert.NoError(t, err)
			}
		})
	}
}

func TestService_Login(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	mockRepo := NewMockRepository(ctrl)
	service := NewService(mockRepo)

	hashedPassword, _ := auth.HashArgon2Password("secretpass123")

	tests := []struct {
		name          string
		email         string
		password      string
		mockBehavior  func()
		wantErr       bool
		expectedError string
	}{
		{
			name:     "Success login",
			email:    "user@example.com",
			password: "secretpass123",
			mockBehavior: func() {
				userId := "user-123"
				mockRepo.EXPECT().FindByEmail(gomock.Any(), "user@example.com").Return(&domain.User{
					ID:       &userId,
					Email:    new("user@example.com"),
					Password: hashedPassword,
					Type:     domain.Client,
				}, nil).Times(1)
			},
			wantErr: false,
		},
		{
			name:     "Error: Invalid credentials (user not found)",
			email:    "nonexistent@example.com",
			password: "password",
			mockBehavior: func() {
				mockRepo.EXPECT().FindByEmail(gomock.Any(), "nonexistent@example.com").Return(nil, nil).Times(1)
			},
			wantErr:       true,
			expectedError: "invalid credentials",
		},
		{
			name:     "Error: Invalid credentials (incorrect password)",
			email:    "user@example.com",
			password: "wrongpassword",
			mockBehavior: func() {
				userId := "user-123"
				mockRepo.EXPECT().FindByEmail(gomock.Any(), "user@example.com").Return(&domain.User{
					ID:       &userId,
					Email:    new("user@example.com"),
					Password: hashedPassword,
					Type:     domain.Client,
				}, nil).Times(1)
			},
			wantErr:       true,
			expectedError: "invalid credentials",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.mockBehavior()
			token, err := service.Login(context.Background(), tt.email, tt.password)

			if tt.wantErr {
				assert.Error(t, err)
				assert.Equal(t, tt.expectedError, err.Error())
				assert.Empty(t, token)
			} else {
				assert.NoError(t, err)
				assert.NotEmpty(t, token)
			}
		})
	}
}

func TestService_ResetPasswordSendToken(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	mockRepo := NewMockRepository(ctrl)
	service := NewService(mockRepo)

	tests := []struct {
		name          string
		email         string
		mockBehavior  func()
		wantErr       bool
		expectedError string
	}{
		{
			name:  "Success send reset token",
			email: "user@example.com",
			mockBehavior: func() {
				mockRepo.EXPECT().FindByEmail(gomock.Any(), "user@example.com").Return(&domain.User{
					Email:     new("user@example.com"),
					FirstName: "John",
				}, nil).Times(1)
				mockRepo.EXPECT().SaveResetCode(gomock.Any(), gomock.Any(), "user@example.com").Return(nil).Times(1)
			},
			wantErr: false,
		},
		{
			name:  "Error: Email not found",
			email: "nonexistent@example.com",
			mockBehavior: func() {
				mockRepo.EXPECT().FindByEmail(gomock.Any(), "nonexistent@example.com").Return(nil, nil).Times(1)
			},
			wantErr:       true,
			expectedError: "email not found",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.mockBehavior()

			t.Setenv("AZURE_EMAIL_CONNECTION_STRING", "Endpoint=https://test.com;AccessKey=dGVzdA==")
			t.Setenv("AZURE_EMAIL_SENDER_ADDRESS", "test@test.com")

			err := service.ResetPasswordSendToken(context.Background(), tt.email)

			if tt.wantErr {
				assert.Error(t, err)
				assert.Equal(t, tt.expectedError, err.Error())
			} else {
				assert.Error(t, err)
				assert.Contains(t, err.Error(), "failed to send email request")
			}
		})
	}
}

func TestService_ResetPasswordVerifyToken(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	mockRepo := NewMockRepository(ctrl)
	service := NewService(mockRepo)

	tests := []struct {
		name          string
		email         string
		code          string
		mockBehavior  func()
		wantErr       bool
		expectedError string
		expectedOk    bool
	}{
		{
			name:  "Success verify token",
			email: "user@example.com",
			code:  "123456",
			mockBehavior: func() {
				mockRepo.EXPECT().GetResetCode(gomock.Any(), "user@example.com").Return("123456", nil).Times(1)
			},
			wantErr:    false,
			expectedOk: true,
		},
		{
			name:  "Error: Invalid code",
			email: "user@example.com",
			code:  "wrong",
			mockBehavior: func() {
				mockRepo.EXPECT().GetResetCode(gomock.Any(), "user@example.com").Return("123456", nil).Times(1)
			},
			wantErr:       true,
			expectedError: "invalid code",
			expectedOk:    false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.mockBehavior()
			ok, err := service.ResetPasswordVerifyToken(context.Background(), tt.email, tt.code)

			if tt.wantErr {
				assert.Error(t, err)
				assert.Equal(t, tt.expectedError, err.Error())
				assert.False(t, ok)
			} else {
				assert.NoError(t, err)
				assert.True(t, ok)
			}
		})
	}
}

func TestService_ResetPassword(t *testing.T) {
	ctrl := gomock.NewController(t)
	mockRepo := NewMockRepository(ctrl)
	service := NewService(mockRepo)

	tests := []struct {
		name          string
		email         string
		code          string
		newPassword   string
		mockBehavior  func()
		wantErr       bool
		expectedError string
	}{
		{
			name:        "Success reset password",
			email:       "user@example.com",
			code:        "123456",
			newPassword: "newsecretpassword123",
			mockBehavior: func() {
				mockRepo.EXPECT().GetResetCode(gomock.Any(), "user@example.com").Return("123456", nil).Times(1)
				mockRepo.EXPECT().FindByEmail(gomock.Any(), "user@example.com").Return(&domain.User{
					Email: new("user@example.com"),
				}, nil).Times(1)
				mockRepo.EXPECT().Update(gomock.Any(), gomock.Any()).Return(nil).Times(1)
			},
			wantErr: false,
		},
		{
			name:        "Error: Invalid code",
			email:       "user@example.com",
			code:        "wrong",
			newPassword: "newsecretpassword123",
			mockBehavior: func() {
				mockRepo.EXPECT().GetResetCode(gomock.Any(), "user@example.com").Return("123456", nil).Times(1)
			},
			wantErr:       true,
			expectedError: "invalid code",
		},
		{
			name:        "Error: User not found",
			email:       "nonexistent@example.com",
			code:        "123456",
			newPassword: "newsecretpassword123",
			mockBehavior: func() {
				mockRepo.EXPECT().GetResetCode(gomock.Any(), "nonexistent@example.com").Return("123456", nil).Times(1)
				mockRepo.EXPECT().FindByEmail(gomock.Any(), "nonexistent@example.com").Return(nil, nil).Times(1)
			},
			wantErr:       true,
			expectedError: "user not found",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.mockBehavior()
			err := service.ResetPassword(context.Background(), tt.email, tt.code, tt.newPassword)

			if tt.wantErr {
				assert.Error(t, err)
				assert.Equal(t, tt.expectedError, err.Error())
			} else {
				assert.NoError(t, err)
			}
		})
	}
}
