package metrics

import (
	"errors"
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
		protected.POST("/body-measurements", h.CreateBodyMeasurement)
		protected.POST("/weight-logs", h.CreateWeightLog)

		trainers := protected.Group("/trainers")
		{
			trainers.POST("/students/:id/body-measurements", h.CreateBodyMeasurement)
			trainers.POST("/students/:id/weight-logs", h.CreateWeightLog)

			trainers.PATCH("/body-measurements/:id/notes", h.AddBodyMeasurementNote)
			trainers.GET("/body-measurements/latest", h.FindLastBodyMeasurementNote)
			trainers.GET("/body-measurements", h.ListBodyMeasurements)
			trainers.GET("/students/:id/body-measurements", h.ListBodyMeasurements)

			trainers.PATCH("/weight-log/:id/notes", h.AddWeightLogNote)
			trainers.GET("/weight-logs", h.ListWeightLogs)
			trainers.GET("/students/:id/weight-logs", h.ListWeightLogs)
		}
	}
}

func (h *Handler) CreateBodyMeasurement(ctx *gin.Context) {
	id := ctx.Param("id")
	userVal, ok := auth.GetAuthUser(ctx)
	if !ok {
		return
	}
	if id == "" {
		id = userVal.ID
	}

	var body struct {
		Type       domain.BodyMeasurementType `json:"type" binding:"required"`
		Value      float64                    `json:"value" binding:"required"`
		MeasuredAt time.Time                  `json:"measuredAt" binding:"required"`
	}

	if err := ctx.ShouldBindJSON(&body); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	measurement := &domain.BodyMeasurement{
		Type:       body.Type,
		Value:      body.Value,
		MeasuredAt: body.MeasuredAt,
		UserId:     id,
	}

	if err := h.srv.CreateBodyMeasurement(ctx.Request.Context(), userVal.ID, measurement); err != nil {
		if errors.Is(err, domain.ErrUnauthorizedAccess) || errors.Is(err, domain.ErrUnauthorizedTrainerAccess) || errors.Is(err, domain.ErrPrivacySettingsForbidden) {
			ctx.JSON(http.StatusForbidden, gin.H{"error": err.Error()})
			return
		}
		slog.ErrorContext(ctx.Request.Context(), "failed to create body measurement", slog.Any("error", err), slog.String("user_id", id))
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "failed to create measurement"})
		return
	}

	ctx.Status(http.StatusCreated)
}

func (h *Handler) AddBodyMeasurementNote(ctx *gin.Context) {
	id := ctx.Param("id")
	userVal, ok := auth.GetAuthUser(ctx)
	if !ok {
		return
	}

	var body struct {
		Note string `json:"note" binding:"required"`
	}

	if err := ctx.ShouldBindJSON(&body); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := h.srv.AddBodyMeasurementNote(ctx.Request.Context(), userVal.ID, id, body.Note); err != nil {
		if errors.Is(err, domain.ErrUnauthorizedAccess) || errors.Is(err, domain.ErrUnauthorizedTrainerAccess) || errors.Is(err, domain.ErrPrivacySettingsForbidden) {
			ctx.JSON(http.StatusForbidden, gin.H{"error": err.Error()})
			return
		}
		slog.ErrorContext(ctx.Request.Context(), "failed to add body measurement note", slog.Any("error", err), slog.String("measurement_id", id))
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "failed to add note"})
		return
	}

	ctx.Status(http.StatusOK)
}

func (h *Handler) FindLastBodyMeasurementNote(ctx *gin.Context) {
	userVal, ok := auth.GetAuthUser(ctx)
	if !ok {
		return
	}

	measurement, err := h.srv.FindLastBodyMeasurementNote(ctx.Request.Context(), userVal.ID, userVal.ID)
	if err != nil {
		if errors.Is(err, domain.ErrUnauthorizedAccess) || errors.Is(err, domain.ErrUnauthorizedTrainerAccess) || errors.Is(err, domain.ErrPrivacySettingsForbidden) {
			ctx.JSON(http.StatusForbidden, gin.H{"error": err.Error()})
			return
		}
		slog.ErrorContext(ctx.Request.Context(), "failed to fetch last measurement", slog.Any("error", err), slog.String("user_id", userVal.ID))
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "failed to fetch last measurement"})
		return
	}

	ctx.JSON(http.StatusOK, measurement)
}

func (h *Handler) ListBodyMeasurements(ctx *gin.Context) {
	id := ctx.Param("id")
	userVal, ok := auth.GetAuthUser(ctx)
	if !ok {
		return
	}
	if id == "" {
		id = userVal.ID
	}

	cursor, limit := utils.GetPagination(ctx)

	measurements, nextCursor, err := h.srv.ListBodyMeasurements(ctx.Request.Context(), userVal.ID, id, cursor, limit)
	if err != nil {
		if errors.Is(err, domain.ErrUnauthorizedAccess) || errors.Is(err, domain.ErrUnauthorizedTrainerAccess) || errors.Is(err, domain.ErrPrivacySettingsForbidden) {
			ctx.JSON(http.StatusForbidden, gin.H{"error": err.Error()})
			return
		}
		slog.ErrorContext(ctx.Request.Context(), "failed to list body measurements", slog.Any("error", err), slog.String("user_id", id))
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "failed to list measurements"})
		return
	}

	ctx.JSON(http.StatusOK, utils.NewPaginatedResponse(measurements, nextCursor))
}

func (h *Handler) CreateWeightLog(ctx *gin.Context) {
	id := ctx.Param("id")
	userVal, ok := auth.GetAuthUser(ctx)
	if !ok {
		return
	}
	if id == "" {
		id = userVal.ID
	}

	var body struct {
		Weight     float64   `json:"weight" binding:"required"`
		MeasuredAt time.Time `json:"measuredAt" binding:"required"`
	}

	if err := ctx.ShouldBindJSON(&body); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	log := &domain.WeightLog{
		Weight:     body.Weight,
		MeasuredAt: body.MeasuredAt,
		UserId:     id,
	}

	if err := h.srv.CreateWeightLog(ctx.Request.Context(), userVal.ID, log); err != nil {
		if errors.Is(err, domain.ErrUnauthorizedAccess) || errors.Is(err, domain.ErrUnauthorizedTrainerAccess) || errors.Is(err, domain.ErrPrivacySettingsForbidden) {
			ctx.JSON(http.StatusForbidden, gin.H{"error": err.Error()})
			return
		}
		slog.ErrorContext(ctx.Request.Context(), "failed to create weight log", slog.Any("error", err), slog.String("user_id", id))
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "failed to create weight log"})
		return
	}

	ctx.Status(http.StatusCreated)
}

func (h *Handler) AddWeightLogNote(ctx *gin.Context) {
	id := ctx.Param("id")
	userVal, ok := auth.GetAuthUser(ctx)
	if !ok {
		return
	}

	var body struct {
		Note string `json:"note" binding:"required"`
	}

	if err := ctx.ShouldBindJSON(&body); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := h.srv.AddWeightLogNote(ctx.Request.Context(), userVal.ID, id, body.Note); err != nil {
		if errors.Is(err, domain.ErrUnauthorizedAccess) || errors.Is(err, domain.ErrUnauthorizedTrainerAccess) || errors.Is(err, domain.ErrPrivacySettingsForbidden) {
			ctx.JSON(http.StatusForbidden, gin.H{"error": err.Error()})
			return
		}
		slog.ErrorContext(ctx.Request.Context(), "failed to add weight log note", slog.Any("error", err), slog.String("log_id", id))
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "failed to add note"})
		return
	}

	ctx.Status(http.StatusOK)
}

func (h *Handler) ListWeightLogs(ctx *gin.Context) {
	id := ctx.Param("id")
	userVal, ok := auth.GetAuthUser(ctx)
	if !ok {
		return
	}
	if id == "" {
		id = userVal.ID
	}

	cursor, limit := utils.GetPagination(ctx)

	logs, nextCursor, err := h.srv.ListWeightLogs(ctx.Request.Context(), userVal.ID, id, cursor, limit)
	if err != nil {
		if errors.Is(err, domain.ErrUnauthorizedAccess) || errors.Is(err, domain.ErrUnauthorizedTrainerAccess) || errors.Is(err, domain.ErrPrivacySettingsForbidden) {
			ctx.JSON(http.StatusForbidden, gin.H{"error": err.Error()})
			return
		}
		slog.ErrorContext(ctx.Request.Context(), "failed to list weight logs", slog.Any("error", err), slog.String("user_id", id))
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "failed to list weight logs"})
		return
	}

	ctx.JSON(http.StatusOK, utils.NewPaginatedResponse(logs, nextCursor))
}
