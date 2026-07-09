package integrationtests

import (
	"context"
	"fmt"
	"os"
	"testing"

	"github.com/testcontainers/testcontainers-go"
	"github.com/testcontainers/testcontainers-go/modules/azure/azurite"
)

func StartAzure(t *testing.T) {
	t.Helper()

	ctx := context.Background()

	azuriteContainer, err := azurite.Run(ctx, "mcr.microsoft.com/azure-storage/azurite:3.33.0",
		testcontainers.WithCmd("--blobHost", "0.0.0.0", "--blobPort", "10000",
			"--queueHost", "0.0.0.0", "--queuePort", "10001",
			"--tableHost", "0.0.0.0", "--tablePort", "10002",
			"--skipApiVersionCheck"),
	)

	t.Log(azuriteContainer.BlobServiceURL(ctx))

	if err != nil {
		t.Fatalf("start azurite container: %v", err)
	}
	t.Cleanup(func() {
		if err := azuriteContainer.Terminate(ctx); err != nil {
			t.Logf("failed to terminate azurite container: %s", err)
		}
	})

	url, err := azuriteContainer.BlobServiceURL(ctx)
	if err != nil {
		t.Fatalf("get azure service url: %s", err)
	}

	blobEndpoint := fmt.Sprintf("%s/%s", url, azurite.AccountName)
	b := fmt.Sprintf("DefaultEndpointsProtocol=%s;AccountName=%s;AccountKey=%s;BlobEndpoint=%s", "http", "devstoreaccount1", "Eby8vdM02xNOcqFlqUwJPLlmEtlCDXJ1OUzFT50uSRZ6IFsuFq2UVErCz4I6tq/K1SZFPTOtr/KBHBeksoGMGw==", blobEndpoint)

	os.Setenv("AZURE_STORAGE_CONNECTION_STRING", b)
	os.Setenv("AZURE_STORAGE_CONTAINER", "testcontainer")
	os.Setenv("AZURE_STORAGE_URL", blobEndpoint)
}
