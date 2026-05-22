package metrics

import (
	"log/slog"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/kaua-nasc/gymtrack-go/apps/api/identity/internal/domain"
	"github.com/kaua-nasc/gymtrack-go/libs/auth"
	"github.com/kaua-nasc/gymtrack-go/libs/utils"
)

type Handler struct {
	srv *Service
}

func NewHandler(srv *Service) *Handler {
	return &Handler{srv: srv}
}

func (h *Handler) RegisterRoutes(r *gin.Engine) {
	protected := r.Group("/identity/users")
	protected.Use(auth.AuthMiddleware())
	{
		trainers := protected.Group("/trainers")
		{
			trainers.PATCH("/body-measurements/:id/notes", h.AddBodyMeasurementNote)
			trainers.GET("/body-measurements/latest", h.FindLastBodyMeasurementNote)
			trainers.GET("/body-measurements", h.ListBodyMeasurements)
			trainers.GET("/students/:id/body-measurements", h.ListBodyMeasurements)

			trainers.PATCH("/weight-log/:id/notes", h.AddWeightLogNote)
			trainers.GET("/weight-logs", h.ListWeightLogs)
			trainers.GET("/students/:id/weight-logs", h.ListWeightLogs)

			trainers.POST("/goals", h.AddGoalMetric)
			trainers.GET("/goals", h.ListGoalsMetric)
			trainers.GET("/students/:id/goals", h.ListGoalsMetricById)
		}
	}
}

func (h *Handler) AddBodyMeasurementNote(ctx *gin.Context) {
	id := ctx.Param("id")

	var body struct {
		Note string `json:"note" binding:"required"`
	}

	if err := ctx.ShouldBindJSON(&body); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := h.srv.AddBodyMeasurementNote(ctx.Request.Context(), id, body.Note); err != nil {
		slog.ErrorContext(ctx.Request.Context(), "failed to add body measurement note", slog.Any("error", err), slog.String("measurement_id", id))
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "failed to add note"})
		return
	}

	ctx.Status(http.StatusOK)
}

func (h *Handler) FindLastBodyMeasurementNote(ctx *gin.Context) {
	user, ok := auth.GetAuthUser(ctx)
	if !ok {
		return
	}

	measurement, err := h.srv.FindLastBodyMeasurementNote(ctx.Request.Context(), user.ID)
	if err != nil {
		slog.ErrorContext(ctx.Request.Context(), "failed to fetch last measurement", slog.Any("error", err), slog.String("user_id", user.ID))
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "failed to fetch last measurement"})
		return
	}

	ctx.JSON(http.StatusOK, measurement)
}

func (h *Handler) ListBodyMeasurements(ctx *gin.Context) {
	id := ctx.Param("id")
	if id == "" {
		user, ok := ctx.Value(string(auth.UserContextKey)).(auth.AuthUser)
		if !ok {
			ctx.JSON(http.StatusUnauthorized, gin.H{"error": "unauthorized"})
			return
		}
		id = user.ID
	}

	cursor, limit := auth.GetPagination(ctx)

	measurements, nextCursor, err := h.srv.ListBodyMeasurements(ctx.Request.Context(), id, cursor, limit)
	if err != nil {
		slog.ErrorContext(ctx.Request.Context(), "failed to list body measurements", slog.Any("error", err), slog.String("user_id", id))
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "failed to list measurements"})
		return
	}

	ctx.JSON(http.StatusOK, utils.NewPaginatedResponse(measurements, nextCursor))
}

func (h *Handler) AddWeightLogNote(ctx *gin.Context) {
	id := ctx.Param("id")

	var body struct {
		Note string `json:"note" binding:"required"`
	}

	if err := ctx.ShouldBindJSON(&body); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := h.srv.AddWeightLogNote(ctx.Request.Context(), id, body.Note); err != nil {
		slog.ErrorContext(ctx.Request.Context(), "failed to add weight log note", slog.Any("error", err), slog.String("log_id", id))
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "failed to add note"})
		return
	}

	ctx.Status(http.StatusOK)
}

func (h *Handler) ListWeightLogs(ctx *gin.Context) {
	id := ctx.Param("id")
	if id == "" {
		user, ok := auth.GetAuthUser(ctx)
		if !ok {
			return
		}
		id = user.ID
	}

	cursor, limit := auth.GetPagination(ctx)

	logs, nextCursor, err := h.srv.ListWeightLogs(ctx.Request.Context(), id, cursor, limit)
	if err != nil {
		slog.ErrorContext(ctx.Request.Context(), "failed to list weight logs", slog.Any("error", err), slog.String("user_id", id))
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "failed to list weight logs"})
		return
	}

	ctx.JSON(http.StatusOK, utils.NewPaginatedResponse(logs, nextCursor))
}

func (h *Handler) AddGoalMetric(ctx *gin.Context) {
	user, ok := auth.GetAuthUser(ctx)
	if !ok {
		return
	}

	var body struct {
		Type        string    `json:"type" binding:"required"`
		TargetValue float64   `json:"targetValue" binding:"required,gt=0"`
		Deadline    time.Time `json:"deadline" binding:"required"`
	}

	if err := ctx.ShouldBindJSON(&body); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := h.srv.AddGoalMetric(ctx.Request.Context(), &domain.MetricGoal{
		UserId:      user.ID,
		Type:        body.Type,
		TargetValue: body.TargetValue,
		Deadline:    &body.Deadline,
	}); err != nil {
		slog.ErrorContext(ctx.Request.Context(), "failed to add goal metric", slog.Any("error", err), slog.String("user_id", user.ID))
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "failed to add goal metric"})
		return
	}

	ctx.Status(http.StatusOK)
}

func (h *Handler) ListGoalsMetric(ctx *gin.Context) {
	user, ok := auth.GetAuthUser(ctx)
	if !ok {
		return
	}
	cursor, limit := auth.GetPagination(ctx)

	goals, nextCursor, err := h.srv.ListGoalsMetric(ctx.Request.Context(), user.ID, cursor, limit)
	if err != nil {
		slog.ErrorContext(ctx.Request.Context(), "failed to list goals", slog.Any("error", err), slog.String("user_id", user.ID))
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "failed to list goals"})
		return
	}

	ctx.JSON(http.StatusOK, utils.NewPaginatedResponse(goals, nextCursor))
}

func (h *Handler) ListGoalsMetricById(ctx *gin.Context) {
	id := ctx.Param("id")
	cursor, limit := auth.GetPagination(ctx)

	goals, nextCursor, err := h.srv.ListGoalsMetric(ctx.Request.Context(), id, cursor, limit)
	if err != nil {
		slog.ErrorContext(ctx.Request.Context(), "failed to list goals by id", slog.Any("error", err), slog.String("user_id", id))
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "failed to list goals"})
		return
	}

	ctx.JSON(http.StatusOK, utils.NewPaginatedResponse(goals, nextCursor))
}
