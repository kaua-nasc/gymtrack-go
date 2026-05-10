package log

import (
	"log/slog"
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/kaua-nasc/gymtrack-go/libs/auth"
)

// LoggerMiddleware returns a gin.HandlerFunc that logs requests using slog
func LoggerMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		start := time.Now()
		path := c.Request.URL.Path
		raw := c.Request.URL.RawQuery

		c.Next()

		latency := time.Since(start)
		status := c.Writer.Status()
		method := c.Request.Method

		if raw != "" {
			path = path + "?" + raw
		}

		// Extract user from context if available (from auth middleware)
		var userID string
		if userVal, exists := c.Get(string(auth.UserContextKey)); exists {
			if user, ok := userVal.(auth.AuthUser); ok {
				userID = user.ID
			}
		}

		attributes := []slog.Attr{
			slog.Int("status", status),
			slog.String("method", method),
			slog.String("path", path),
			slog.String("ip", c.ClientIP()),
			slog.Duration("latency", latency),
			slog.String("user_agent", c.Request.UserAgent()),
		}

		if userID != "" {
			attributes = append(attributes, slog.String("user_id", userID))
		}

		msg := "request processed"
		level := slog.LevelInfo

		if status >= http.StatusInternalServerError {
			level = slog.LevelError
			msg = "request failed"
		} else if status >= http.StatusBadRequest {
			level = slog.LevelWarn
			msg = "client error"
		}

		slog.LogAttrs(c.Request.Context(), level, msg, attributes...)
	}
}

// InitLogger initializes the default slog logger
func InitLogger() {
	level := slog.LevelInfo
	if os.Getenv("DEBUG") == "true" {
		level = slog.LevelDebug
	}

	logger := slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{
		Level: level,
	}))
	slog.SetDefault(logger)
}
