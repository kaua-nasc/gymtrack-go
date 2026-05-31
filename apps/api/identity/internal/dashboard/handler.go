package dashboard

import (
	"errors"
	"log/slog"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/kaua-nasc/gymtrack-go/apps/api/identity/internal/domain"
	"github.com/kaua-nasc/gymtrack-go/libs/auth"
)

type Handler struct {
	srv *Service
}

func NewHandler(srv *Service) *Handler {
	return &Handler{srv: srv}
}

func (h *Handler) RegisterRoutes(r *gin.Engine) {
	protected := r.Group("/identity/users/trainers/students/:id/dashboard")
	protected.Use(auth.AuthMiddleware())
	{
		protected.GET("/engagement", h.GetStudentEngagement)
		protected.GET("/biometrics", h.GetStudentBiometrics)
		protected.GET("/insights", h.GetStudentInsights)
	}
}

func (h *Handler) GetStudentInsights(ctx *gin.Context) {
	studentId := ctx.Param("id")
	userVal, ok := auth.GetAuthUser(ctx)
	if !ok {
		return
	}

	insights, err := h.srv.GetStudentInsights(ctx.Request.Context(), userVal.ID, studentId)
	if err != nil {
		if errors.Is(err, domain.ErrUnauthorizedTrainerAccess) {
			ctx.JSON(http.StatusForbidden, gin.H{"error": err.Error()})
			return
		}
		slog.ErrorContext(ctx.Request.Context(), "failed to get student insights", slog.Any("error", err), slog.String("trainer_id", userVal.ID), slog.String("student_id", studentId))
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "failed to get insights dashboard data"})
		return
	}

	ctx.JSON(http.StatusOK, insights)
}

func (h *Handler) GetStudentBiometrics(ctx *gin.Context) {
	studentId := ctx.Param("id")
	userVal, ok := auth.GetAuthUser(ctx)
	if !ok {
		return
	}

	startStr := ctx.Query("start")
	endStr := ctx.Query("end")

	start, err := time.Parse(time.RFC3339, startStr)
	if err != nil {
		// Default to 30 days if invalid
		start = time.Now().UTC().AddDate(0, 0, -30)
	}

	end, err := time.Parse(time.RFC3339, endStr)
	if err != nil {
		end = time.Now().UTC()
	}

	summary, err := h.srv.GetStudentBiometrics(ctx.Request.Context(), userVal.ID, studentId, start, end)
	if err != nil {
		if errors.Is(err, domain.ErrUnauthorizedTrainerAccess) || errors.Is(err, domain.ErrPrivacySettingsForbidden) {
			ctx.JSON(http.StatusForbidden, gin.H{"error": err.Error()})
			return
		}
		if errors.Is(err, domain.ErrInvalidPeriod) {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		slog.ErrorContext(ctx.Request.Context(), "failed to get student biometrics", slog.Any("error", err), slog.String("trainer_id", userVal.ID), slog.String("student_id", studentId))
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "failed to get biometrics dashboard data"})
		return
	}

	ctx.JSON(http.StatusOK, summary)
}

func (h *Handler) GetStudentEngagement(ctx *gin.Context) {
	studentId := ctx.Param("id")
	userVal, ok := auth.GetAuthUser(ctx)
	if !ok {
		return
	}

	summary, err := h.srv.GetStudentEngagement(ctx.Request.Context(), userVal.ID, studentId)
	if err != nil {
		if errors.Is(err, domain.ErrUnauthorizedTrainerAccess) || errors.Is(err, domain.ErrPrivacySettingsForbidden) {
			ctx.JSON(http.StatusForbidden, gin.H{"error": err.Error()})
			return
		}
		slog.ErrorContext(ctx.Request.Context(), "failed to get student engagement", slog.Any("error", err), slog.String("trainer_id", userVal.ID), slog.String("student_id", studentId))
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "failed to get dashboard data"})
		return
	}

	ctx.JSON(http.StatusOK, summary)
}
