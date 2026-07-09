package user

import (
	"bytes"
	"context"
	"database/sql"
	"os"
	"testing"

	"github.com/Azure/azure-sdk-for-go/sdk/storage/azblob"
	"github.com/kaua-nasc/gymtrack-go/apps/api/identity/internal/domain"
	integrationtests "github.com/kaua-nasc/gymtrack-go/libs/integration-tests"
	"github.com/kaua-nasc/gymtrack-go/libs/utils"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestIntegration_UserUpdateProfile(t *testing.T) {
	db, cleanup := integrationtests.StartPostgres(t)
	integrationtests.StartAzure(t)

	if db == nil {
		return
	}
	defer cleanup()

	repo := NewRepository(db)
	srv := NewService(repo)
	id := createTestUser(t, db)

	t.Run("update all profile fields", func(t *testing.T) {
		firstName := "Jane"
		lastName := "Smith"
		bio := "Updated bio"
		height := 170.5
		weightUnit := domain.KG
		heightUnit := domain.CM
		currentWeight := 65.0

		err := srv.UpdateProfile(
			context.Background(), id,
			&firstName, &lastName, &bio,
			&height, &weightUnit, &heightUnit, &currentWeight,
		)
		require.NoError(t, err)

		user, err := repo.Find(context.Background(), id, id)
		require.NoError(t, err)
		require.NotNil(t, user)

		assert.Equal(t, firstName, user.FirstName)
		assert.Equal(t, lastName, user.LastName)
		assert.Equal(t, &bio, user.Bio)
		assert.Equal(t, height, *user.Height)
		assert.Equal(t, currentWeight, *user.CurrentWeight)
		assert.Equal(t, weightUnit, user.WeightUnit)
		assert.Equal(t, heightUnit, user.HeightUnit)
	})

	t.Run("return error for nonexistent user", func(t *testing.T) {
		fakeID := "00000000-0000-0000-0000-000000000000"
		err := srv.UpdateProfile(
			context.Background(), fakeID,
			nil, nil, nil, nil, nil, nil, nil,
		)
		require.ErrorIs(t, err, domain.ErrUserNotFound)
	})
}

func TestIntegration_UploadProfilePicture(t *testing.T) {
	db, cleanup := integrationtests.StartPostgres(t)
	integrationtests.StartAzure(t)

	if db == nil {
		return
	}
	defer cleanup()

	connStr := os.Getenv("AZURE_STORAGE_CONNECTION_STRING")
	azClient, err := azblob.NewClientFromConnectionString(connStr, nil)
	require.NoError(t, err)
	_, err = azClient.CreateContainer(context.Background(), "testcontainer", nil)
	require.NoError(t, err)

	repo := NewRepository(db)
	srv := NewService(repo)
	id := createTestUser(t, db)

	t.Run("upload profile picture successfully", func(t *testing.T) {
		fileContent := []byte("fake-image-content")
		reader := bytes.NewReader(fileContent)

		err := srv.UploadProfilePicture(context.Background(), id, reader)
		require.NoError(t, err)

		user, err := repo.Find(context.Background(), id, id)
		require.NoError(t, err)
		require.NotNil(t, user)
		require.NotNil(t, user.ProfilePictureUrl)
		require.Contains(t, *user.ProfilePictureUrl, "identity/user/profile/")
	})

	t.Run("return error for nonexistent user", func(t *testing.T) {
		fakeID := "00000000-0000-0000-0000-000000000000"
		reader := bytes.NewReader([]byte("content"))
		err := srv.UploadProfilePicture(context.Background(), fakeID, reader)
		require.ErrorIs(t, err, domain.ErrUserNotFound)
	})
}

func createTestUser(t *testing.T, db *sql.DB) string {
	t.Helper()

	id, err := utils.GenerateUUIDV7String(context.Background())
	require.NoError(t, err)

	_, err = db.Exec(`
		INSERT INTO users (id, "firstName", "lastName", email, password, type, "isVerified")
		VALUES ($1, 'John', 'Doe', 'john@example.com', 'hashed-password', 'CLIENT', true)
	`, id)
	require.NoError(t, err)

	return id
}
