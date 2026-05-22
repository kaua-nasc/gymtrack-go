package user

import (
	"bytes"
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

func (m *MockRepository) Find(ctx context.Context, id string, currentUserId string) (*domain.User, error) {
	args := m.Called(ctx, id, currentUserId)
	user, _ := args.Get(0).(*domain.User)
	return user, args.Error(1)
}

func (m *MockRepository) Update(ctx context.Context, u *domain.User) error {
	args := m.Called(ctx, u)
	return args.Error(0)
}

func (m *MockRepository) ListByIDs(ctx context.Context, ids []string) ([]*domain.User, error) {
	args := m.Called(ctx, ids)
	users, _ := args.Get(0).([]*domain.User)
	return users, args.Error(1)
}

func (m *MockRepository) ChangeUserType(ctx context.Context, u domain.User, newType domain.UserType) error {
	args := m.Called(ctx, u, newType)
	return args.Error(0)
}

func (m *MockRepository) RemoveProfilePicture(ctx context.Context, userId string) error {
	args := m.Called(ctx, userId)
	return args.Error(0)
}

func (m *MockRepository) ChangeProfileImage(ctx context.Context, u domain.User, pictureUrl string) error {
	args := m.Called(ctx, u, pictureUrl)
	return args.Error(0)
}

func TestService_ChangePassword(t *testing.T) {
	mockRepo := new(MockRepository)
	service := NewService(mockRepo)

	hashedPassword, _ := auth.HashArgon2Password("oldpassword123")

	tests := []struct {
		name            string
		userId          string
		currentPassword string
		newPassword     string
		mockBehavior    func()
		wantErr         bool
		expectedError   string
	}{
		{
			name:            "Success change password",
			userId:          "user-123",
			currentPassword: "oldpassword123",
			newPassword:     "newpassword123",
			mockBehavior: func() {
				userId := "user-123"
				mockRepo.On("Find", mock.Anything, "user-123", "").Return(&domain.User{ID: &userId, Password: hashedPassword, IsVerified: true}, nil).Once()
				mockRepo.On("Update", mock.Anything, mock.MatchedBy(func(u *domain.User) bool {
					ok, _ := auth.VerifyArgon2Password("newpassword123", u.Password)
					return ok
				})).Return(nil).Once()
			},
			wantErr: false,
		},
		{
			name:            "Error: User not verified",
			userId:          "user-123",
			currentPassword: "oldpassword123",
			newPassword:     "newpassword123",
			mockBehavior: func() {
				userId := "user-123"
				mockRepo.On("Find", mock.Anything, "user-123", "").Return(&domain.User{ID: &userId, Password: hashedPassword, IsVerified: false}, nil).Once()
			},
			wantErr:       true,
			expectedError: "user not verified",
		},
		{
			name:            "Error: Incorrect current password",
			userId:          "user-123",
			currentPassword: "wrongpassword",
			newPassword:     "newpassword123",
			mockBehavior: func() {
				userId := "user-123"
				mockRepo.On("Find", mock.Anything, "user-123", "").Return(&domain.User{ID: &userId, Password: hashedPassword, IsVerified: true}, nil).Once()
			},
			wantErr:       true,
			expectedError: "invalid credentials",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.mockBehavior()
			err := service.ChangePassword(context.Background(), tt.userId, tt.currentPassword, tt.newPassword)

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

func TestService_UpdateProfile(t *testing.T) {
	mockRepo := new(MockRepository)
	service := NewService(mockRepo)

	tests := []struct {
		name          string
		id            string
		firstName     *string
		lastName      *string
		bio           *string
		height        *float64
		weightUnit    *domain.WeightUnit
		heightUnit    *domain.HeightUnit
		currentWeight *float64
		mockBehavior  func()
		wantErr       bool
		expectedError string
	}{
		{
			name:          "Success update profile",
			id:            "user-123",
			firstName:     new("Jane"),
			lastName:      new("Doe"),
			bio:           new("New bio"),
			height:        new(175.5),
			weightUnit:    new(domain.KG),
			heightUnit:    new(domain.CM),
			currentWeight: new(70.2),
			mockBehavior: func() {
				userId := "user-123"
				mockRepo.On("Find", mock.Anything, "user-123", "").Return(&domain.User{
					ID:        &userId,
					FirstName: "John",
					LastName:  "Doe",
				}, nil).Once()
				mockRepo.On("Update", mock.Anything, mock.MatchedBy(func(u *domain.User) bool {
					return u.FirstName == "Jane" && u.LastName == "Doe" && *u.Bio == "New bio" && *u.Height == 175.5 && u.WeightUnit == domain.KG && u.HeightUnit == domain.CM && *u.CurrentWeight == 70.2
				})).Return(nil).Once()
			},
			wantErr: false,
		},
		{
			name: "Error: User not found",
			id:   "nonexistent",
			mockBehavior: func() {
				mockRepo.On("Find", mock.Anything, "nonexistent", "").Return(nil, nil).Once()
			},
			wantErr:       true,
			expectedError: "user not found",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.mockBehavior()
			err := service.UpdateProfile(context.Background(), tt.id, tt.firstName, tt.lastName, tt.bio, tt.height, tt.weightUnit, tt.heightUnit, tt.currentWeight)

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

func TestService_GetUser(t *testing.T) {
	mockRepo := new(MockRepository)
	service := NewService(mockRepo)

	tests := []struct {
		name          string
		id            string
		currentUserId string
		mockBehavior  func()
		wantErr       bool
		expectedUser  *domain.User
	}{
		{
			name:          "Success get user with sanitization",
			id:            "user-123",
			currentUserId: "user-456",
			mockBehavior: func() {
				userId := "user-123"
				mockRepo.On("Find", mock.Anything, "user-123", "user-456").Return(&domain.User{
					ID:                &userId,
					Password:          "hashed_pass",
					ProfilePictureUrl: new("uploads/profile.png"),
				}, nil).Once()
			},
			wantErr: false,
			expectedUser: &domain.User{
				ID:                new("user-123"),
				Password:          "",
				ProfilePictureUrl: new("/uploads/profile.png"),
			},
		},
		{
			name:          "Success get user with azure storage URL prefix",
			id:            "user-123",
			currentUserId: "user-456",
			mockBehavior: func() {
				t.Setenv("AZURE_STORAGE_URL", "https://mystorage.blob.core.windows.net")
				userId := "user-123"
				mockRepo.On("Find", mock.Anything, "user-123", "user-456").Return(&domain.User{
					ID:                &userId,
					Password:          "hashed_pass",
					ProfilePictureUrl: new("uploads/profile.png"),
				}, nil).Once()
			},
			wantErr: false,
			expectedUser: &domain.User{
				ID:                new("user-123"),
				Password:          "",
				ProfilePictureUrl: new("https://mystorage.blob.core.windows.net/uploads/profile.png"),
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.mockBehavior()
			user, err := service.GetUser(context.Background(), tt.id, tt.currentUserId)

			if tt.wantErr {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
				assert.NotNil(t, user)
				assert.Equal(t, tt.expectedUser.Password, user.Password)
				if tt.expectedUser.ProfilePictureUrl != nil {
					assert.Equal(t, *tt.expectedUser.ProfilePictureUrl, *user.ProfilePictureUrl)
				}
			}
			mockRepo.AssertExpectations(t)
		})
	}
}

func TestService_ListUsers(t *testing.T) {
	mockRepo := new(MockRepository)
	service := NewService(mockRepo)

	tests := []struct {
		name         string
		ids          []string
		mockBehavior func()
		wantErr      bool
	}{
		{
			name: "Success list users",
			ids:  []string{"user-1", "user-2"},
			mockBehavior: func() {
				userId1 := "user-1"
				userId2 := "user-2"
				mockRepo.On("ListByIDs", mock.Anything, []string{"user-1", "user-2"}).Return([]*domain.User{
					{ID: &userId1, Password: "pass"},
					{ID: &userId2, Password: "pass"},
				}, nil).Once()
			},
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.mockBehavior()
			users, err := service.ListUsers(context.Background(), tt.ids)

			if tt.wantErr {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
				assert.Len(t, users, 2)
				assert.Empty(t, users[0].Password)
				assert.Empty(t, users[1].Password)
			}
			mockRepo.AssertExpectations(t)
		})
	}
}

func TestService_ChangeUserType(t *testing.T) {
	mockRepo := new(MockRepository)
	service := NewService(mockRepo)

	tests := []struct {
		name          string
		id            string
		isTrainer     bool
		mockBehavior  func()
		wantErr       bool
		expectedError string
	}{
		{
			name:      "Success change to trainer",
			id:        "user-123",
			isTrainer: true,
			mockBehavior: func() {
				userId := "user-123"
				mockRepo.On("Find", mock.Anything, "user-123", "").Return(&domain.User{ID: &userId, Type: domain.Client}, nil).Once()
				mockRepo.On("ChangeUserType", mock.Anything, mock.Anything, domain.Trainer).Return(nil).Once()
			},
			wantErr: false,
		},
		{
			name:      "Success change to client",
			id:        "user-123",
			isTrainer: false,
			mockBehavior: func() {
				userId := "user-123"
				mockRepo.On("Find", mock.Anything, "user-123", "").Return(&domain.User{ID: &userId, Type: domain.Trainer}, nil).Once()
				mockRepo.On("ChangeUserType", mock.Anything, mock.Anything, domain.Client).Return(nil).Once()
			},
			wantErr: false,
		},
		{
			name:      "Error: User not found (trainer)",
			id:        "nonexistent",
			isTrainer: true,
			mockBehavior: func() {
				mockRepo.On("Find", mock.Anything, "nonexistent", "").Return(nil, nil).Once()
			},
			wantErr:       true,
			expectedError: "user not found",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.mockBehavior()
			var err error
			if tt.isTrainer {
				err = service.ChangeToTrainer(context.Background(), tt.id, "CREF123")
			} else {
				err = service.ChangeToClient(context.Background(), tt.id)
			}

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

func TestService_ProfilePicture(t *testing.T) {
	mockRepo := new(MockRepository)
	service := NewService(mockRepo)

	// Test RemoveProfilePicture
	mockRepo.On("RemoveProfilePicture", mock.Anything, "user-123").Return(nil).Once()
	err := service.RemoveProfilePicture(context.Background(), "user-123")
	assert.NoError(t, err)

	// Test UploadProfilePicture - user not found
	mockRepo.On("Find", mock.Anything, "nonexistent", "").Return(nil, nil).Once()
	fileContent := bytes.NewReader([]byte("fake_image_data"))
	err = service.UploadProfilePicture(context.Background(), "nonexistent", fileContent)
	assert.Error(t, err)
	assert.Equal(t, "user not found", err.Error())

	// Test UploadProfilePicture - storage error because AZURE_STORAGE_CONNECTION_STRING is missing
	userId := "user-123"
	mockRepo.On("Find", mock.Anything, "user-123", "").Return(&domain.User{ID: &userId}, nil).Once()
	t.Setenv("AZURE_STORAGE_CONNECTION_STRING", "") // Ensure it is empty
	err = service.UploadProfilePicture(context.Background(), "user-123", fileContent)
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "AZURE_STORAGE_CONNECTION_STRING env variable not found")

	mockRepo.AssertExpectations(t)
}
