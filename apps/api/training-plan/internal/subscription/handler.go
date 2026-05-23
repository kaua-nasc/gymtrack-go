package subscription

import (
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
		plans.GET("/subscriptions", h.ListSubscription)
		plans.GET("/subscriptions/:id", h.ListSubscriptionByUserId)
		plans.POST("/:id/subscriptions", h.Subscribe)
		plans.DELETE("/:id/subscriptions", h.Unsubscribe)

		plans.POST("/:id/subscriptions/send", h.ChangeSubscriptionStatus)
		plans.POST("/:id/subscriptions/privacy", h.ChangeSubscriptionPrivacy)
		plans.POST("/subscriptions/:id/days/:dayId/complete", h.CompleteDay)
		plans.POST("/subscriptions/:id/days/:dayId/cancel", h.CancelDay)
		plans.POST("/subscriptions/:id/days/:dayId/start", h.StartDay)
		plans.GET("/subscriptions/days/resume", h.ListWeeklyDayProgress)
		plans.GET("/subscriptions/days/next", h.FindNextDay)
	}
}

func (h *Handler) ListSubscription(ctx *gin.Context) {
	userId := ctx.Param("userId")
	user, ok := auth.GetAuthUser(ctx)
	if !ok {
		return
	}

	if userId == "" {
		userId = user.ID
	}

	subscriptions, err := h.srv.ListSubscription(ctx.Request.Context(), userId)
	if err != nil {
		slog.ErrorContext(ctx.Request.Context(), "failed to list subscriptions", slog.Any("error", err), slog.String("user_id", userId))
		ctx.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	ctx.JSON(http.StatusOK, subscriptions)
}

func (h *Handler) ListSubscriptionByUserId(ctx *gin.Context) {
	userId := ctx.Param("id")
	user, ok := auth.GetAuthUser(ctx)
	if !ok {
		return
	}

	if _, err := h.srv.ListSubscriptionByUserId(ctx.Request.Context(), user.ID, userId); err != nil {
		slog.ErrorContext(ctx.Request.Context(), "failed to list subscriptions", slog.Any("error", err), slog.String("user_id", userId))
		ctx.JSON(http.StatusInternalServerError, utils.NewErrorResponse("failed to list subscriptions"))
		return
	}

	ctx.Status(http.StatusNoContent)
}

func (h *Handler) Subscribe(ctx *gin.Context) {
	id := ctx.Param("id")
	user, ok := auth.GetAuthUser(ctx)
	if !ok {
		return
	}

	var body struct {
		Type domain.PlanSubscriptionType `json:"type" binding:"required"`
	}

	if err := ctx.ShouldBindJSON(&body); err != nil {
		ctx.JSON(http.StatusBadRequest, utils.NewErrorResponse(err.Error()))
		return
	}

	if err := h.srv.Subscribe(ctx.Request.Context(), id, user.ID, body.Type); err != nil {
		slog.ErrorContext(ctx.Request.Context(), "failed to subscribe", slog.Any("error", err), slog.String("plan_id", id), slog.String("user_id", user.ID))
		ctx.JSON(http.StatusInternalServerError, utils.NewErrorResponse("failed to subscribe"))
		return
	}

	ctx.Status(http.StatusNoContent)
}

func (h *Handler) ChangeSubscriptionStatus(ctx *gin.Context) {
	planId := ctx.Param("id")

	var body struct {
		Status domain.PlanSubscriptionStatus `json:"status" binding:"required"`
	}

	if err := ctx.ShouldBindJSON(&body); err != nil {
		ctx.JSON(http.StatusBadRequest, utils.NewErrorResponse(err.Error()))
		return
	}

	user, ok := auth.GetAuthUser(ctx)
	if !ok {
		return
	}

	if err := h.srv.ChangeSubscriptionStatus(ctx.Request.Context(), planId, user.ID, body.Status); err != nil {
		slog.ErrorContext(ctx.Request.Context(), "failed to change subscription status", slog.Any("error", err), slog.String("plan_id", planId), slog.String("user_id", user.ID))
		ctx.JSON(http.StatusInternalServerError, utils.NewErrorResponse("failed to change status"))
		return
	}

	ctx.Status(http.StatusNoContent)
}

func (h *Handler) ChangeSubscriptionPrivacy(ctx *gin.Context) {
	planId := ctx.Param("id")

	var body struct {
		SubscriptionType domain.PlanSubscriptionType `json:"type" binding:"required"`
	}

	if err := ctx.ShouldBindJSON(&body); err != nil {
		ctx.JSON(http.StatusBadRequest, utils.NewErrorResponse(err.Error()))
		return
	}

	user, ok := auth.GetAuthUser(ctx)
	if !ok {
		return
	}

	if err := h.srv.ChangeSubscriptionPrivacy(ctx.Request.Context(), planId, user.ID, body.SubscriptionType); err != nil {
		slog.ErrorContext(ctx.Request.Context(), "failed to change subscription privacy", slog.Any("error", err), slog.String("plan_id", planId), slog.String("user_id", user.ID))
		ctx.JSON(http.StatusInternalServerError, utils.NewErrorResponse("failed to change privacy"))
		return
	}

	ctx.Status(http.StatusOK)
}

func (h *Handler) Unsubscribe(ctx *gin.Context) {
	id := ctx.Param("id")
	user, ok := auth.GetAuthUser(ctx)
	if !ok {
		return
	}

	if err := h.srv.Unsubscribe(ctx.Request.Context(), id, user.ID); err != nil {
		slog.ErrorContext(ctx.Request.Context(), "failed to unsubscribe", slog.Any("error", err), slog.String("plan_id", id), slog.String("user_id", user.ID))
		ctx.JSON(http.StatusInternalServerError, utils.NewErrorResponse("failed to unsubscribe"))
		return
	}

	ctx.Status(http.StatusOK)
}

func (h *Handler) CompleteDay(ctx *gin.Context) {
	subsId := ctx.Param("id")
	dayId := ctx.Param("dayId")
	user, ok := auth.GetAuthUser(ctx)
	if !ok {
		return
	}

	if err := h.srv.CompleteDay(ctx.Request.Context(), subsId, user.ID, dayId); err != nil {
		slog.ErrorContext(ctx.Request.Context(), "failed to complete day", slog.Any("error", err), slog.String("subscription_id", subsId), slog.String("day_id", dayId), slog.String("user_id", user.ID))
		ctx.JSON(http.StatusInternalServerError, utils.NewErrorResponse(err.Error()))
		return
	}

	ctx.Status(http.StatusOK)
}

func (h *Handler) CancelDay(ctx *gin.Context) {
	subsId := ctx.Param("id")
	dayId := ctx.Param("dayId")
	user, ok := auth.GetAuthUser(ctx)
	if !ok {
		return
	}

	if err := h.srv.CancelDay(ctx.Request.Context(), subsId, user.ID, dayId); err != nil {
		slog.ErrorContext(ctx.Request.Context(), "failed to cancel day", slog.Any("error", err), slog.String("subscription_id", subsId), slog.String("day_id", dayId), slog.String("user_id", user.ID))
		ctx.JSON(http.StatusInternalServerError, utils.NewErrorResponse(err.Error()))
		return
	}

	ctx.Status(http.StatusOK)
}

func (h *Handler) StartDay(ctx *gin.Context) {
	subsId := ctx.Param("id")
	dayId := ctx.Param("dayId")
	user, ok := auth.GetAuthUser(ctx)
	if !ok {
		return
	}

	if err := h.srv.StartDay(ctx.Request.Context(), subsId, user.ID, dayId); err != nil {
		slog.ErrorContext(ctx.Request.Context(), "failed to start day", slog.Any("error", err), slog.String("subscription_id", subsId), slog.String("day_id", dayId), slog.String("user_id", user.ID))
		ctx.JSON(http.StatusInternalServerError, utils.NewErrorResponse(err.Error()))
		return
	}

	ctx.Status(http.StatusOK)
}

func (h *Handler) ListWeeklyDayProgress(ctx *gin.Context) {
	user, ok := auth.GetAuthUser(ctx)
	if !ok {
		return
	}

	progress, err := h.srv.ListWeeklyDayProgress(ctx.Request.Context(), user.ID)
	if err != nil {
		slog.ErrorContext(ctx.Request.Context(), "failed to list weekly day progress", slog.Any("error", err), slog.String("user_id", user.ID))
		ctx.JSON(http.StatusInternalServerError, utils.NewErrorResponse("failed to list weekly day progress"))
		return
	}

	ctx.JSON(http.StatusOK, progress)
}

func (h *Handler) FindNextDay(ctx *gin.Context) {
	user, ok := auth.GetAuthUser(ctx)
	if !ok {
		return
	}

	progress, err := h.srv.FindNextDay(ctx.Request.Context(), user.ID)
	if err != nil {
		slog.ErrorContext(ctx.Request.Context(), "failed to find next day", slog.Any("error", err), slog.String("user_id", user.ID))
		ctx.JSON(http.StatusInternalServerError, utils.NewErrorResponse("failed to find next day"))
		return
	}

	ctx.JSON(http.StatusOK, progress)
}
