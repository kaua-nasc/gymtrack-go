package utils

import (
	"context"
	"fmt"
	"log/slog"

	"github.com/google/uuid"
)

// GenerateUUIDV7 generates a new UUID v7 as a string.
// It returns a pointer to the string to maintain compatibility with existing service structures.
func GenerateUUIDV7(ctx context.Context) (*string, error) {
	id, err := uuid.NewV7()
	if err != nil {
		slog.ErrorContext(ctx, "failed to generate uuid v7", slog.Any("error", err))
		return nil, fmt.Errorf("error on generate uuid: %w", err)
	}

	idStr := id.String()
	return &idStr, nil
}

// GenerateUUIDV7String generates a new UUID v7 and returns it as a value string.
func GenerateUUIDV7String(ctx context.Context) (string, error) {
	id, err := uuid.NewV7()
	if err != nil {
		slog.ErrorContext(ctx, "failed to generate uuid v7", slog.Any("error", err))
		return "", fmt.Errorf("error on generate uuid: %w", err)
	}
	return id.String(), nil
}
