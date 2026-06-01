package storage

import (
	"context"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"os"
	"path/filepath"
	"time"

	"github.com/Azure/azure-sdk-for-go/sdk/storage/azblob"
)

func UploadBuffer(ctx context.Context, path string, bytes []byte) (string, error) {
	connString := os.Getenv("AZURE_STORAGE_CONNECTION_STRING")
	if connString == "" {
		return "", fmt.Errorf("AZURE_STORAGE_CONNECTION_STRING env variable not found")
	}

	containerName := os.Getenv("AZURE_STORAGE_CONTAINER")
	if containerName == "" {
		return "", fmt.Errorf("AZURE_STORAGE_CONTAINER env variable not found")
	}

	client, err := azblob.NewClientFromConnectionString(connString, nil)
	if err != nil {
		return "", err
	}

	// Generate hashed filename
	dir := filepath.Dir(path)
	ext := filepath.Ext(path)
	base := filepath.Base(path)
	nameWithoutExt := base[:len(base)-len(ext)]

	hashInput := fmt.Sprintf("%s-%d", nameWithoutExt, time.Now().UnixNano())
	hash := sha256.Sum256([]byte(hashInput))
	hashedName := hex.EncodeToString(hash[:])

	finalFilename := filepath.Join(dir, hashedName+ext)

	a, err := client.UploadBuffer(ctx, containerName, finalFilename, bytes, &azblob.UploadBufferOptions{})
	if err != nil {
		return "", err
	}

	if a.Date == nil {
		return "", fmt.Errorf("falha no upload: %w", err)
	}

	return finalFilename, nil
}

func GetBlobURL(filename string) *string {
	uri := os.Getenv("AZURE_STORAGE_URL")
	if uri == "" {
		panic("AZURE_STORAGE_URL env variable not found")
	}

	url := fmt.Sprintf("%s/%s", uri, filename)

	return &url
}
