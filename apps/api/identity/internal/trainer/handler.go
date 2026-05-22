package trainer

import (
	"errors"
	"log/slog"
	"net/http"

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
			trainers.PATCH("/profile/code", h.CreateTrainerCode)
			trainers.POST("/profile/link", h.LinkTrainer)
			trainers.POST("/profile/unlink", h.UnlinkTrainer)
			trainers.GET("/students", h.ListStudents)
			trainers.POST("/students/:id/profile/unlink", h.UnlinkStudent)
		}
	}
}

func (h *Handler) CreateTrainerCode(ctx *gin.Context) {
	user, ok := auth.GetAuthUser(ctx)
	if !ok {
		return
	}

	var body struct {
		Code string `json:"code" binding:"required"`
	}

	if err := ctx.ShouldBindJSON(&body); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := h.srv.CreateTrainerCode(ctx.Request.Context(), user.ID, body.Code); err != nil {
		slog.ErrorContext(ctx.Request.Context(), "failed to create trainer code", slog.Any("error", err), slog.String("user_id", user.ID))
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "failed to create trainer code"})
		return
	}

	ctx.Status(http.StatusOK)
}

func (h *Handler) LinkTrainer(ctx *gin.Context) {
	user, ok := auth.GetAuthUser(ctx)
	if !ok {
		return
	}

	var body struct {
		Code string `json:"code" binding:"required"`
	}

	if err := ctx.ShouldBindJSON(&body); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := h.srv.LinkTrainer(ctx.Request.Context(), user.ID, body.Code); err != nil {
		if errors.Is(err, domain.ErrTrainerNotFound) {
			ctx.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
			return
		}
		slog.ErrorContext(ctx.Request.Context(), "failed to link trainer", slog.Any("error", err), slog.String("user_id", user.ID), slog.String("code", body.Code))
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "failed to link trainer"})
		return
	}

	ctx.Status(http.StatusOK)
}

func (h *Handler) UnlinkTrainer(ctx *gin.Context) {
	user, ok := auth.GetAuthUser(ctx)
	if !ok {
		return
	}

	if err := h.srv.UnlinkTrainer(ctx.Request.Context(), user.ID); err != nil {
		slog.ErrorContext(ctx.Request.Context(), "failed to unlink trainer", slog.Any("error", err), slog.String("user_id", user.ID))
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "failed to unlink trainer"})
		return
	}

	ctx.Status(http.StatusOK)
}

func (h *Handler) ListStudents(ctx *gin.Context) {
	user, ok := auth.GetAuthUser(ctx)
	if !ok {
		return
	}

	cursor, limit := auth.GetPagination(ctx)

	users, nextCursor, err := h.srv.ListStudents(ctx.Request.Context(), user.ID, cursor, limit)
	if err != nil {
		slog.ErrorContext(ctx.Request.Context(), "failed to list students", slog.Any("error", err), slog.String("trainer_id", user.ID))
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "failed to list students"})
		return
	}

	ctx.JSON(http.StatusOK, utils.NewPaginatedResponse(users, nextCursor))
}

func (h *Handler) UnlinkStudent(ctx *gin.Context) {
	id := ctx.Param("id")

	if err := h.srv.UnlinkStudent(ctx.Request.Context(), id); err != nil {
		slog.ErrorContext(ctx.Request.Context(), "failed to unlink student", slog.Any("error", err), slog.String("student_id", id))
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "failed to unlink student"})
		return
	}

	ctx.Status(http.StatusOK)
}
