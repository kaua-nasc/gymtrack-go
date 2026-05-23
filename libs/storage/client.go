package storage

import (
	"context"
	"fmt"
	"os"

	"github.com/Azure/azure-sdk-for-go/sdk/storage/azblob"
)

func UploadBuffer(ctx context.Context, filename string, bytes []byte) error {
	connString := os.Getenv("AZURE_STORAGE_CONNECTION_STRING")
	if connString == "" {
		return fmt.Errorf("AZURE_STORAGE_CONNECTION_STRING env variable not found")
	}

	containerName := os.Getenv("AZURE_STORAGE_CONTAINER")
	if containerName == "" {
		return fmt.Errorf("AZURE_STORAGE_CONTAINER env variable not found")
	}

	client, err := azblob.NewClientFromConnectionString(connString, nil)
	if err != nil {
		return err
	}
	a, err := client.UploadBuffer(ctx, containerName, filename, bytes, &azblob.UploadBufferOptions{})
	if err != nil {
		return err
	}

	if a.Date == nil {
		return fmt.Errorf("falha no upload: %w", err)
	}

	return nil
}

func GetBlobURL(filename string) *string {
	uri := os.Getenv("AZURE_STORAGE_URL")

	url := fmt.Sprintf("%s/%s", uri, filename)

	return &url
}
