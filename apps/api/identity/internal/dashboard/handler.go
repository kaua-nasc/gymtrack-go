package dashboard

import (
	"errors"
	"log/slog"
	"net/http"

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
	}
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
