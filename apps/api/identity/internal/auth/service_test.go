package auth

import (
	"context"
	"testing"

	"github.com/kaua-nasc/gymtrack-go/apps/api/identity/internal/domain"
	"github.com/kaua-nasc/gymtrack-go/libs/auth"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type MockRepository struct {
	mock.Mock
}

func (m *MockRepository) FindByEmail(ctx context.Context, email string) (*domain.User, error) {
	args := m.Called(ctx, email)
	user, _ := args.Get(0).(*domain.User)
	return user, args.Error(1)
}

func (m *MockRepository) Create(ctx context.Context, u *domain.User) error {
	args := m.Called(ctx, u)
	return args.Error(0)
}

func (m *MockRepository) SaveResetCode(ctx context.Context, code, email string) error {
	args := m.Called(ctx, code, email)
	return args.Error(0)
}

func (m *MockRepository) GetResetCode(ctx context.Context, email string) (string, error) {
	args := m.Called(ctx, email)
	return args.String(0), args.Error(1)
}

func (m *MockRepository) SaveVerificationCode(ctx context.Context, code, email string) error {
	args := m.Called(ctx, code, email)
	return args.Error(0)
}

func (m *MockRepository) GetVerificationCode(ctx context.Context, email string) (string, error) {
	args := m.Called(ctx, email)
	return args.String(0), args.Error(1)
}

func (m *MockRepository) Update(ctx context.Context, u *domain.User) error {
	args := m.Called(ctx, u)
	return args.Error(0)
}

func TestService_Register(t *testing.T) {
	mockRepo := new(MockRepository)
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
				Email:     "newuser@example.com",
				Password:  "strongpassword123",
				FirstName: "John",
				LastName:  "Doe",
			},
			mockBehavior: func() {
				mockRepo.On("FindByEmail", mock.Anything, "newuser@example.com").Return(nil, nil).Once()
				mockRepo.On("Create", mock.Anything, mock.AnythingOfType("*domain.User")).Return(nil).Once()
			},
			wantErr: false,
		},
		{
			name: "Error: User already exists",
			inputUser: domain.User{
				Email: "existing@example.com",
			},
			mockBehavior: func() {
				userId := "123"
				mockRepo.On("FindByEmail", mock.Anything, "existing@example.com").Return(&domain.User{ID: &userId}, nil).Once()
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
			mockRepo.AssertExpectations(t)
		})
	}
}

func TestService_SendVerificationEmail(t *testing.T) {
	mockRepo := new(MockRepository)
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
				mockRepo.On("FindByEmail", mock.Anything, "user@example.com").Return(&domain.User{Email: "user@example.com", FirstName: "John", IsVerified: false}, nil).Once()
				mockRepo.On("SaveVerificationCode", mock.Anything, mock.AnythingOfType("string"), "user@example.com").Return(nil).Once()
			},
			wantErr: false,
		},
		{
			name:  "Error: User not found",
			email: "nonexistent@example.com",
			mockBehavior: func() {
				mockRepo.On("FindByEmail", mock.Anything, "nonexistent@example.com").Return(nil, nil).Once()
			},
			wantErr:       true,
			expectedError: "email not found",
		},
		{
			name:  "Error: Already verified",
			email: "verified@example.com",
			mockBehavior: func() {
				mockRepo.On("FindByEmail", mock.Anything, "verified@example.com").Return(&domain.User{Email: "verified@example.com", IsVerified: true}, nil).Once()
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
			mockRepo.AssertExpectations(t)
		})
	}
}

func TestService_VerifyEmail(t *testing.T) {
	mockRepo := new(MockRepository)
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
				mockRepo.On("GetVerificationCode", mock.Anything, "user@example.com").Return("123456", nil).Once()
				mockRepo.On("FindByEmail", mock.Anything, "user@example.com").Return(&domain.User{Email: "user@example.com", IsVerified: false}, nil).Once()
				mockRepo.On("Update", mock.Anything, mock.MatchedBy(func(u *domain.User) bool {
					return u.IsVerified == true
				})).Return(nil).Once()
			},
			wantErr: false,
		},
		{
			name:  "Error: Invalid code",
			email: "user@example.com",
			code:  "wrong",
			mockBehavior: func() {
				mockRepo.On("GetVerificationCode", mock.Anything, "user@example.com").Return("123456", nil).Once()
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
			mockRepo.AssertExpectations(t)
		})
	}
}

func TestService_Login(t *testing.T) {
	mockRepo := new(MockRepository)
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
				mockRepo.On("FindByEmail", mock.Anything, "user@example.com").Return(&domain.User{
					ID:       &userId,
					Email:    "user@example.com",
					Password: hashedPassword,
					Type:     domain.Client,
				}, nil).Once()
			},
			wantErr: false,
		},
		{
			name:     "Error: Invalid credentials (user not found)",
			email:    "nonexistent@example.com",
			password: "password",
			mockBehavior: func() {
				mockRepo.On("FindByEmail", mock.Anything, "nonexistent@example.com").Return(nil, nil).Once()
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
				mockRepo.On("FindByEmail", mock.Anything, "user@example.com").Return(&domain.User{
					ID:       &userId,
					Email:    "user@example.com",
					Password: hashedPassword,
					Type:     domain.Client,
				}, nil).Once()
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
			mockRepo.AssertExpectations(t)
		})
	}
}

func TestService_ResetPasswordSendToken(t *testing.T) {
	mockRepo := new(MockRepository)
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
				mockRepo.On("FindByEmail", mock.Anything, "user@example.com").Return(&domain.User{
					Email:     "user@example.com",
					FirstName: "John",
				}, nil).Once()
				mockRepo.On("SaveResetCode", mock.Anything, mock.AnythingOfType("string"), "user@example.com").Return(nil).Once()
			},
			wantErr: false,
		},
		{
			name:  "Error: Email not found",
			email: "nonexistent@example.com",
			mockBehavior: func() {
				mockRepo.On("FindByEmail", mock.Anything, "nonexistent@example.com").Return(nil, nil).Once()
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
			mockRepo.AssertExpectations(t)
		})
	}
}

func TestService_ResetPasswordVerifyToken(t *testing.T) {
	mockRepo := new(MockRepository)
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
				mockRepo.On("GetResetCode", mock.Anything, "user@example.com").Return("123456", nil).Once()
			},
			wantErr:    false,
			expectedOk: true,
		},
		{
			name:  "Error: Invalid code",
			email: "user@example.com",
			code:  "wrong",
			mockBehavior: func() {
				mockRepo.On("GetResetCode", mock.Anything, "user@example.com").Return("123456", nil).Once()
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
			mockRepo.AssertExpectations(t)
		})
	}
}

func TestService_ResetPassword(t *testing.T) {
	mockRepo := new(MockRepository)
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
				mockRepo.On("GetResetCode", mock.Anything, "user@example.com").Return("123456", nil).Once()
				mockRepo.On("FindByEmail", mock.Anything, "user@example.com").Return(&domain.User{
					Email: "user@example.com",
				}, nil).Once()
				mockRepo.On("Update", mock.Anything, mock.MatchedBy(func(u *domain.User) bool {
					ok, _ := auth.VerifyArgon2Password("newsecretpassword123", u.Password)
					return ok
				})).Return(nil).Once()
			},
			wantErr: false,
		},
		{
			name:        "Error: Invalid code",
			email:       "user@example.com",
			code:        "wrong",
			newPassword: "newsecretpassword123",
			mockBehavior: func() {
				mockRepo.On("GetResetCode", mock.Anything, "user@example.com").Return("123456", nil).Once()
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
				mockRepo.On("GetResetCode", mock.Anything, "nonexistent@example.com").Return("123456", nil).Once()
				mockRepo.On("FindByEmail", mock.Anything, "nonexistent@example.com").Return(nil, nil).Once()
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
			mockRepo.AssertExpectations(t)
		})
	}
}
