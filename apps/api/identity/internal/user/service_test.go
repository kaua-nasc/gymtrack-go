package user

import (
	"bytes"
	"context"
	"testing"

	"github.com/kaua-nasc/gymtrack-go/apps/api/identity/internal/domain"
	"github.com/kaua-nasc/gymtrack-go/libs/auth"
	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
)

func TestService_ChangePassword(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := NewMockRepository(ctrl)
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
				userIdStr := "user-123"
				mockRepo.EXPECT().Find(gomock.Any(), "user-123", "").Return(&domain.User{ID: &userIdStr, Password: hashedPassword, IsVerified: true}, nil)
				mockRepo.EXPECT().Update(gomock.Any(), gomock.Any()).DoAndReturn(func(ctx context.Context, u *domain.User) error {
					ok, _ := auth.VerifyArgon2Password("newpassword123", u.Password)
					if !ok {
						return domain.ErrInvalidCredentials
					}
					return nil
				})
			},
			wantErr: false,
		},
		{
			name:            "Error: User not verified",
			userId:          "user-123",
			currentPassword: "oldpassword123",
			newPassword:     "newpassword123",
			mockBehavior: func() {
				userIdStr := "user-123"
				mockRepo.EXPECT().Find(gomock.Any(), "user-123", "").Return(&domain.User{ID: &userIdStr, Password: hashedPassword, IsVerified: false}, nil)
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
				userIdStr := "user-123"
				mockRepo.EXPECT().Find(gomock.Any(), "user-123", "").Return(&domain.User{ID: &userIdStr, Password: hashedPassword, IsVerified: true}, nil)
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
		})
	}
}

func TestService_UpdateProfile(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := NewMockRepository(ctrl)
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
			firstName:     func(s string) *string { return &s }("Jane"),
			lastName:      func(s string) *string { return &s }("Doe"),
			bio:           func(s string) *string { return &s }("New bio"),
			height:        func(f float64) *float64 { return &f }(175.5),
			weightUnit:    func(u domain.WeightUnit) *domain.WeightUnit { return &u }(domain.KG),
			heightUnit:    func(u domain.HeightUnit) *domain.HeightUnit { return &u }(domain.CM),
			currentWeight: func(f float64) *float64 { return &f }(70.2),
			mockBehavior: func() {
				userIdStr := "user-123"
				mockRepo.EXPECT().Find(gomock.Any(), "user-123", "").Return(&domain.User{
					ID:        &userIdStr,
					FirstName: "John",
					LastName:  "Doe",
				}, nil)
				mockRepo.EXPECT().Update(gomock.Any(), gomock.Any()).DoAndReturn(func(ctx context.Context, u *domain.User) error {
					if u.FirstName == "Jane" && u.LastName == "Doe" && *u.Bio == "New bio" && *u.Height == 175.5 && u.WeightUnit == domain.KG && u.HeightUnit == domain.CM && *u.CurrentWeight == 70.2 {
						return nil
					}
					return assert.AnError
				})
			},
			wantErr: false,
		},
		{
			name: "Error: User not found",
			id:   "nonexistent",
			mockBehavior: func() {
				mockRepo.EXPECT().Find(gomock.Any(), "nonexistent", "").Return(nil, nil)
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
		})
	}
}

func TestService_GetUser(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := NewMockRepository(ctrl)
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
				userIdStr := "user-123"
				mockRepo.EXPECT().Find(gomock.Any(), "user-123", "user-456").Return(&domain.User{
					ID:                &userIdStr,
					Password:          "hashed_pass",
					ProfilePictureUrl: func(s string) *string { return &s }("uploads/profile.png"),
				}, nil)
				mockRepo.EXPECT().GetPrivacySettings(gomock.Any(), "user-123").Return(nil, nil)
			},
			wantErr: false,
			expectedUser: &domain.User{
				ID:                func(s string) *string { return &s }("user-123"),
				Password:          "",
				ProfilePictureUrl: func(s string) *string { return &s }("/uploads/profile.png"),
			},
		},
		{
			name:          "Success get user with azure storage URL prefix",
			id:            "user-123",
			currentUserId: "user-456",
			mockBehavior: func() {
				t.Setenv("AZURE_STORAGE_URL", "https://mystorage.blob.core.windows.net")
				userIdStr := "user-123"
				mockRepo.EXPECT().Find(gomock.Any(), "user-123", "user-456").Return(&domain.User{
					ID:                &userIdStr,
					Password:          "hashed_pass",
					ProfilePictureUrl: func(s string) *string { return &s }("uploads/profile.png"),
				}, nil)
				mockRepo.EXPECT().GetPrivacySettings(gomock.Any(), "user-123").Return(nil, nil)
			},
			wantErr: false,
			expectedUser: &domain.User{
				ID:                func(s string) *string { return &s }("user-123"),
				Password:          "",
				ProfilePictureUrl: func(s string) *string { return &s }("https://mystorage.blob.core.windows.net/uploads/profile.png"),
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
		})
	}
}

func TestService_ListUsers(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := NewMockRepository(ctrl)
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
				userId1Str := "user-1"
				userId2Str := "user-2"
				mockRepo.EXPECT().ListByIDs(gomock.Any(), []string{"user-1", "user-2"}).Return([]*domain.User{
					{ID: &userId1Str, Password: "pass"},
					{ID: &userId2Str, Password: "pass"},
				}, nil)
				mockRepo.EXPECT().GetPrivacySettings(gomock.Any(), "user-1").Return(nil, nil)
				mockRepo.EXPECT().GetPrivacySettings(gomock.Any(), "user-2").Return(nil, nil)
			},
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.mockBehavior()
			users, err := service.ListUsers(context.Background(), "requester-123", tt.ids)

			if tt.wantErr {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
				assert.Len(t, users, 2)
				assert.Empty(t, users[0].Password)
				assert.Empty(t, users[1].Password)
			}
		})
	}
}

func TestService_ChangeUserType(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := NewMockRepository(ctrl)
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
				userIdStr := "user-123"
				mockRepo.EXPECT().Find(gomock.Any(), "user-123", "").Return(&domain.User{ID: &userIdStr, Type: domain.Client}, nil)
				mockRepo.EXPECT().ChangeUserType(gomock.Any(), gomock.Any(), domain.Trainer).Return(nil)
			},
			wantErr: false,
		},
		{
			name:      "Success change to client",
			id:        "user-123",
			isTrainer: false,
			mockBehavior: func() {
				userIdStr := "user-123"
				mockRepo.EXPECT().Find(gomock.Any(), "user-123", "").Return(&domain.User{ID: &userIdStr, Type: domain.Trainer}, nil)
				mockRepo.EXPECT().ChangeUserType(gomock.Any(), gomock.Any(), domain.Client).Return(nil)
			},
			wantErr: false,
		},
		{
			name:      "Error: User not found (trainer)",
			id:        "nonexistent",
			isTrainer: true,
			mockBehavior: func() {
				mockRepo.EXPECT().Find(gomock.Any(), "nonexistent", "").Return(nil, nil)
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
		})
	}
}

func TestService_ProfilePicture(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := NewMockRepository(ctrl)
	service := NewService(mockRepo)

	// Test RemoveProfilePicture
	mockRepo.EXPECT().RemoveProfilePicture(gomock.Any(), "user-123").Return(nil)
	err := service.RemoveProfilePicture(context.Background(), "user-123")
	assert.NoError(t, err)

	// Test UploadProfilePicture - user not found
	mockRepo.EXPECT().Find(gomock.Any(), "nonexistent", "").Return(nil, nil)
	fileContent := bytes.NewReader([]byte("fake_image_data"))
	err = service.UploadProfilePicture(context.Background(), "nonexistent", fileContent)
	assert.Error(t, err)
	assert.Equal(t, "user not found", err.Error())

	// Test UploadProfilePicture - storage error because AZURE_STORAGE_CONNECTION_STRING is missing
	userIdStr := "user-123"
	mockRepo.EXPECT().Find(gomock.Any(), "user-123", "").Return(&domain.User{ID: &userIdStr}, nil)
	t.Setenv("AZURE_STORAGE_CONNECTION_STRING", "") // Ensure it is empty
	err = service.UploadProfilePicture(context.Background(), "user-123", fileContent)
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "AZURE_STORAGE_CONNECTION_STRING env variable not found")
}

func TestService_SearchUsers(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := NewMockRepository(ctrl)
	service := NewService(mockRepo)

	tests := []struct {
		name         string
		term         string
		requesterId  string
		cursorStr    *string
		limit        int
		mockBehavior func()
		wantErr      bool
		wantCount    int
	}{
		{
			name:         "Return empty list for short term",
			term:         "a",
			requesterId:  "user-1",
			mockBehavior: func() {},
			wantErr:      false,
			wantCount:    0,
		},
		{
			name:        "Success search users",
			term:        "John",
			requesterId: "user-1",
			limit:       10,
			mockBehavior: func() {
				userId2 := "user-2"
				mockRepo.EXPECT().SearchByName(gomock.Any(), "user-1", "John", gomock.Any(), 10).Return([]*domain.User{
					{ID: &userId2, FirstName: "John", LastName: "Doe", Password: "pass"},
				}, nil, nil)
				mockRepo.EXPECT().GetPrivacySettings(gomock.Any(), "user-2").Return(nil, nil)
			},
			wantErr:   false,
			wantCount: 1,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.mockBehavior()
			cursor := ""
			if tt.cursorStr != nil {
				cursor = *tt.cursorStr
			}
			users, _, err := service.SearchUsers(context.Background(), tt.term, tt.requesterId, cursor, tt.limit)

			if tt.wantErr {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
				assert.Len(t, users, tt.wantCount)
				if tt.wantCount > 0 {
					assert.Empty(t, users[0].Password)
				}
			}
		})
	}
}
