package subscription

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
		plans.GET("/subscriptions", h.ListSubscription)
		plans.GET("/subscriptions/:id", h.ListSubscriptionByUserId)
		plans.GET("/my", h.ListSubscribedPlans)
		plans.POST("/:id/subscriptions", h.Subscribe)
		plans.DELETE("/:id/subscriptions", h.Unsubscribe)

		plans.POST("/:id/subscriptions/send", h.ChangeSubscriptionStatus)
		plans.POST("/:id/subscriptions/privacy", h.ChangeSubscriptionPrivacy)
		plans.POST("/subscriptions/:id/days/:dayId/complete", h.CompleteDay)
		plans.POST("/subscriptions/:id/days/:dayId/cancel", h.CancelDay)
		plans.POST("/subscriptions/:id/days/:dayId/start", h.StartDay)
		plans.GET("/subscriptions/days/resume", h.ListWeeklyDayProgress)
		plans.GET("/subscriptions/days/next", h.FindNextDay)
		plans.GET("/subscriptions/engagement/:id", h.GetEngagementSummary)
	}
}

func (h *Handler) GetEngagementSummary(ctx *gin.Context) {
	userId := ctx.Param("id")

	summary, err := h.srv.GetEngagementSummary(ctx.Request.Context(), userId)
	if err != nil {
		slog.ErrorContext(ctx.Request.Context(), "failed to get engagement summary", slog.Any("error", err), slog.String("user_id", userId))
		ctx.JSON(http.StatusInternalServerError, utils.NewErrorResponse("failed to get engagement summary"))
		return
	}

	ctx.JSON(http.StatusOK, summary)
}

func (h *Handler) ListSubscribedPlans(ctx *gin.Context) {
	user, ok := auth.GetAuthUser(ctx)
	if !ok {
		return
	}

	cursor, limit := utils.GetPagination(ctx)
	filters := h.parseFilters(ctx)

	plans, nextCursor, err := h.srv.ListSubscribedPlans(ctx.Request.Context(), user.ID, filters, cursor, limit)
	if err != nil {
		slog.ErrorContext(ctx.Request.Context(), "failed to list subscribed plans", slog.Any("error", err), slog.String("user_id", user.ID))
		ctx.JSON(http.StatusInternalServerError, utils.NewErrorResponse("failed to list subscribed plans"))
		return
	}

	ctx.JSON(http.StatusOK, utils.NewPaginatedResponse(plans, nextCursor))
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

	filters := h.parseFilters(ctx)

	subscriptions, err := h.srv.ListSubscription(ctx.Request.Context(), userId, filters)
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

	filters := h.parseFilters(ctx)

	res, err := h.srv.ListSubscriptionByUserId(ctx.Request.Context(), user.ID, userId, filters)
	if err != nil {
		if errors.Is(err, domain.ErrPrivacySettingsForbidden) {
			ctx.JSON(http.StatusForbidden, utils.NewErrorResponse(err.Error()))
			return
		}
		slog.ErrorContext(ctx.Request.Context(), "failed to list subscriptions", slog.Any("error", err), slog.String("user_id", userId))
		ctx.JSON(http.StatusInternalServerError, utils.NewErrorResponse("failed to list subscriptions"))
		return
	}

	ctx.JSON(http.StatusOK, res)
}

func (h *Handler) parseFilters(ctx *gin.Context) domain.ListSubscriptionFilters {
	var filters domain.ListSubscriptionFilters

	if status := ctx.Query("status"); status != "" {
		s := domain.PlanSubscriptionStatus(status)
		filters.Status = &s
	}

	if subType := ctx.Query("type"); subType != "" {
		t := domain.PlanSubscriptionType(subType)
		filters.Type = &t
	}

	if planType := ctx.Query("planType"); planType != "" {
		pt := domain.TrainingPlanType(planType)
		filters.PlanType = &pt
	}

	if visibility := ctx.Query("visibility"); visibility != "" {
		v := domain.TrainingPlanVisibility(visibility)
		filters.Visibility = &v
	}

	if level := ctx.Query("level"); level != "" {
		l := domain.TrainingPlanLevel(level)
		filters.Level = &l
	}

	if authorId := ctx.Query("authorId"); authorId != "" {
		filters.AuthorId = &authorId
	}

	return filters
}

func (h *Handler) Subscribe(ctx *gin.Context) {
	id := ctx.Param("id")
	user, ok := auth.GetAuthUser(ctx)
	if !ok {
		return
	}

	if err := h.srv.Subscribe(ctx.Request.Context(), id, user.ID); err != nil {
		switch {
		case errors.Is(err, domain.ErrPlanNotFound):
			ctx.JSON(http.StatusNotFound, utils.NewErrorResponse(err.Error()))
			return
		case errors.Is(err, domain.ErrAlreadySubscribed), errors.Is(err, domain.ErrMaxSubscriptionsReached):
			ctx.JSON(http.StatusConflict, utils.NewErrorResponse(err.Error()))
			return
		case errors.Is(err, domain.ErrPlanIncomplete):
			ctx.JSON(http.StatusBadRequest, utils.NewErrorResponse(err.Error()))
			return
		case errors.Is(err, domain.ErrSubscriptionForbidden):
			ctx.JSON(http.StatusForbidden, utils.NewErrorResponse(err.Error()))
			return
		default:
			slog.ErrorContext(ctx.Request.Context(), "failed to subscribe", slog.Any("error", err), slog.String("plan_id", id), slog.String("user_id", user.ID))
			ctx.JSON(http.StatusInternalServerError, utils.NewErrorResponse("failed to subscribe"))
			return
		}
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
