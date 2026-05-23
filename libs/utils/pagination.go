package utils

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

// CursorData represents the structure of a pagination cursor.
type CursorData struct {
	ID        string    `json:"id"`
	CreatedAt time.Time `json:"createdAt"`
}

// EncodeCursor takes any data structure and returns a Base64 encoded JSON string.
// This is typically used to generate "next cursor" strings for pagination.
func EncodeCursor(cursor any) (string, error) {
	if cursor == nil {
		return "", nil
	}

	b, err := json.Marshal(cursor)
	if err != nil {
		return "", fmt.Errorf("failed to marshal cursor: %w", err)
	}

	return base64.StdEncoding.EncodeToString(b), nil
}

// DecodeCursor takes a Base64 encoded JSON string and decodes it into the target structure.
// If cursorStr is empty, it does nothing and returns no error.
func DecodeCursor(cursorStr string, target any) error {
	if cursorStr == "" {
		return nil
	}

	b, err := base64.StdEncoding.DecodeString(cursorStr)
	if err != nil {
		return fmt.Errorf("failed to decode base64 cursor: %w", err)
	}

	if err := json.Unmarshal(b, target); err != nil {
		return fmt.Errorf("failed to unmarshal cursor json: %w", err)
	}

	return nil
}

func GetPagination(ctx *gin.Context) (string, int) {
	cursor := ctx.Query("cursor")
	limit, _ := strconv.Atoi(ctx.DefaultQuery("limit", "20"))
	if limit <= 0 {
		limit = 20
	}
	return cursor, limit
}
