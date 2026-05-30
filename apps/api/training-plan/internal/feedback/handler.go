package feedback

import (
	"errors"
	"log/slog"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/kaua-nasc/gymtrack-go/apps/api/training-plan/internal/domain"
	"github.com/kaua-nasc/gymtrack-go/libs/auth"
	"github.com/kaua-nasc/gymtrack-go/libs/log"
	"github.com/kaua-nasc/gymtrack-go/libs/utils"
)

type Handler struct {
	srv *Service
}

func NewHandler(srv *Service) *Handler {
	return &Handler{
		srv: srv,
	}
}

func (h *Handler) RegisterRoutes(r *gin.Engine) {
	r.Use(log.LoggerMiddleware())
	plans := r.Group("/training-plans")
	plans.Use(auth.AuthMiddleware())
	{
		plans.POST("/:id/feedback", h.AddFeedback)
		plans.GET("/:id/feedback", h.ListFeedback)
	}
}

func (h *Handler) AddFeedback(ctx *gin.Context) {
	id := ctx.Param("id")
	user, ok := auth.GetAuthUser(ctx)
	if !ok {
		return
	}

	var body struct {
		Rating  float64 `json:"rating" binding:"required,min=0,max=5"`
		Message *string `json:"message"`
	}

	if err := ctx.ShouldBindJSON(&body); err != nil {
		ctx.JSON(http.StatusBadRequest, utils.NewErrorResponse(err.Error()))
		return
	}

	if err := h.srv.AddFeedback(ctx.Request.Context(), id, user.ID, body.Rating, body.Message); err != nil {
		if errors.Is(err, domain.ErrFeedbackNotAllowed) {
			ctx.JSON(http.StatusForbidden, utils.NewErrorResponse(err.Error()))
			return
		}
		slog.ErrorContext(ctx.Request.Context(), "failed to add feedback", slog.Any("error", err), slog.String("plan_id", id), slog.String("user_id", user.ID))
		ctx.JSON(http.StatusInternalServerError, utils.NewErrorResponse("failed to add feedback"))
		return
	}

	ctx.Status(http.StatusNoContent)
}

func (h *Handler) ListFeedback(ctx *gin.Context) {
	id := ctx.Param("id")
	if id == "" {
		ctx.JSON(http.StatusBadRequest, utils.NewErrorResponse("plan id is required"))
		return
	}

	cursor, limit := utils.GetPagination(ctx)

	feedbacks, nextCursor, err := h.srv.ListFeedback(ctx.Request.Context(), id, cursor, limit)
	if err != nil {
		slog.ErrorContext(ctx.Request.Context(), "failed to list feedback", slog.Any("error", err), slog.String("plan_id", id))
		ctx.JSON(http.StatusInternalServerError, utils.NewErrorResponse("failed to list feedback"))
		return
	}

	ctx.JSON(http.StatusOK, utils.NewPaginatedResponse(feedbacks, nextCursor))
}
