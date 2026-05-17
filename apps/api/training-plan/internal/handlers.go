package internal

import (
	"log/slog"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/kaua-nasc/gymtrack-go/libs/auth"
	"github.com/kaua-nasc/gymtrack-go/libs/log"
	"github.com/kaua-nasc/gymtrack-go/libs/utils"
)

type TrainingPlanHandler struct {
	srv *TrainingPlanService
}

func NewTrainingPlanHandler(srv *TrainingPlanService) *TrainingPlanHandler {
	return &TrainingPlanHandler{
		srv: srv,
	}
}

func (h *TrainingPlanHandler) RegisterRoutes(r *gin.Engine) {
	r.Use(log.LoggerMiddleware())
	r.GET("/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"status": "ok", "service": "training-plan"})
	})
	plans := r.Group("/training-plans")
	plans.Use(auth.AuthMiddleware())
	{
		plans.GET("", h.ListPlan)
		plans.POST("/by-ids", h.ListPlansByIds)
		plans.POST("", h.CreatePlan)
		plans.GET("author/:authorId", h.ListPlan)
		plans.GET("/:id", h.GetPlan)
		plans.PUT("/:id", h.UpdatePlan)
		plans.DELETE("/:id", h.DeletePlan)
		plans.GET("/exists/:id", h.ExistsPlan)

		plans.GET("/subscriptions", h.ListSubscription)
		plans.GET("/subscriptions/:userId", h.ListSubscriptionByUserId)
		plans.POST("/:id/subscriptions", h.Subscribe)
		plans.DELETE("/:id/subscriptions", h.Unsubscribe)

		plans.POST("/:id/subscriptions/send", h.ChangeSubscriptionStatus)
		plans.POST("/:id/subscriptions/privacy", h.ChangeSubscriptionPrivacy)

		plans.POST("/:id/days", h.CreateDay)
		plans.DELETE("/:id/days/:dayId", h.DeleteDay)
		plans.PUT("/:id/days/:dayId/complete", h.CompleteDay)

		plans.POST("/:id/feedback", h.AddFeedback)

		plans.POST("/:id/days/:dayId/exercises", h.CreateExercise)
		plans.DELETE("/:id/days/:dayId/exercises/:exerciseId", h.DeleteExercise)
		plans.POST("/:id/days/:dayId/exercises/:exerciseId/logs", h.LogExercise)

		plans.GET("/activity/weekly", h.ListActivityWeekly)
	}
}

func (h *TrainingPlanHandler) AddFeedback(ctx *gin.Context) {
	id := ctx.Param("id")
	user, ok := h.getAuthUser(ctx)
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
		slog.ErrorContext(ctx.Request.Context(), "failed to add feedback", slog.Any("error", err), slog.String("plan_id", id), slog.String("user_id", user.ID))
		ctx.JSON(http.StatusInternalServerError, utils.NewErrorResponse("failed to add feedback"))
		return
	}

	ctx.Status(http.StatusNoContent)
}

func (h *TrainingPlanHandler) LogExercise(ctx *gin.Context) {
	id := ctx.Param("exerciseId")
	user, ok := h.getAuthUser(ctx)
	if !ok {
		return
	}

	var body struct {
		Reps   []int     `json:"reps" binding:"required"`
		Weight []float64 `json:"weight" binding:"required"`
		Notes  *string   `json:"notes"`
	}

	if err := ctx.ShouldBindJSON(&body); err != nil {
		ctx.JSON(http.StatusBadRequest, utils.NewErrorResponse(err.Error()))
		return
	}

	if err := h.srv.LogExercise(ctx.Request.Context(), id, user.ID, body.Reps, body.Weight, body.Notes); err != nil {
		slog.ErrorContext(ctx.Request.Context(), "failed to log exercise", slog.Any("error", err), slog.String("exercise_id", id), slog.String("user_id", user.ID))
		ctx.JSON(http.StatusInternalServerError, utils.NewErrorResponse("failed to log exercise"))
		return
	}

	ctx.Status(http.StatusNoContent)
}

func (h *TrainingPlanHandler) ListSubscription(ctx *gin.Context) {
	userId := ctx.Param("userId")
	user, ok := h.getAuthUser(ctx)
	if !ok {
		return
	}

	if userId == "" {
		userId = user.ID
	}

	subscriptions, err := h.srv.ListSubscription(ctx.Request.Context(), userId)
	if err != nil {
		slog.ErrorContext(ctx.Request.Context(), "failed to list subscriptions", slog.Any("error", err), slog.String("user_id", userId))
		ctx.JSON(http.StatusInternalServerError, utils.NewErrorResponse("failed to list subscriptions"))
		return
	}

	ctx.JSON(http.StatusOK, subscriptions)
}

func (h *TrainingPlanHandler) ListSubscriptionByUserId(ctx *gin.Context) {
	userId := ctx.Param("userId")
	user, ok := h.getAuthUser(ctx)
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

func (h *TrainingPlanHandler) Subscribe(ctx *gin.Context) {
	id := ctx.Param("id")
	user, ok := h.getAuthUser(ctx)
	if !ok {
		return
	}

	var body struct {
		Type PlanSubscriptionType `json:"type" binding:"required"`
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

func (h *TrainingPlanHandler) ChangeSubscriptionStatus(ctx *gin.Context) {
	planId := ctx.Param("id")

	var body struct {
		Status PlanSubscriptionStatus `json:"status" binding:"required"`
	}

	if err := ctx.ShouldBindJSON(&body); err != nil {
		ctx.JSON(http.StatusBadRequest, utils.NewErrorResponse(err.Error()))
		return
	}

	user, ok := h.getAuthUser(ctx)
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

func (h *TrainingPlanHandler) ChangeSubscriptionPrivacy(ctx *gin.Context) {
	planId := ctx.Param("id")

	var body struct {
		SubscriptionType PlanSubscriptionType `json:"type" binding:"required"`
	}

	if err := ctx.ShouldBindJSON(&body); err != nil {
		ctx.JSON(http.StatusBadRequest, utils.NewErrorResponse(err.Error()))
		return
	}

	user, ok := h.getAuthUser(ctx)
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

func (h *TrainingPlanHandler) Unsubscribe(ctx *gin.Context) {
	id := ctx.Param("id")
	user, ok := h.getAuthUser(ctx)
	if !ok {
		return
	}

	if err := h.srv.Unsubscribe(ctx.Request.Context(), id, user.ID); err != nil {
		slog.ErrorContext(ctx.Request.Context(), "failed to unsubscribe", slog.Any("error", err), slog.String("plan_id", id), slog.String("user_id", user.ID))
		ctx.JSON(http.StatusInternalServerError, utils.NewErrorResponse("failed to unsubscribe"))
		return
	}

	ctx.Status(http.StatusNoContent)
}

func (h *TrainingPlanHandler) CompleteDay(ctx *gin.Context) {
	planId := ctx.Param("id")
	dayId := ctx.Param("dayId")
	user, ok := h.getAuthUser(ctx)
	if !ok {
		return
	}

	if err := h.srv.CompleteDay(ctx.Request.Context(), planId, user.ID, dayId); err != nil {
		slog.ErrorContext(ctx.Request.Context(), "failed to complete day", slog.Any("error", err), slog.String("plan_id", planId), slog.String("day_id", dayId), slog.String("user_id", user.ID))
		ctx.JSON(http.StatusInternalServerError, utils.NewErrorResponse(err.Error()))
		return
	}

	ctx.Status(http.StatusNoContent)
}

func (h *TrainingPlanHandler) CreatePlan(ctx *gin.Context) {
	var plan TrainingPlan
	if err := ctx.ShouldBindJSON(&plan); err != nil {
		ctx.JSON(http.StatusBadRequest, utils.NewErrorResponse(err.Error()))
		return
	}

	user, ok := h.getAuthUser(ctx)
	if !ok {
		return
	}

	_, err := h.srv.CreatePlan(ctx.Request.Context(), plan, user)
	if err != nil {
		slog.ErrorContext(ctx.Request.Context(), "failed to create plan", slog.Any("error", err), slog.String("authorId", user.ID))
		ctx.JSON(http.StatusInternalServerError, utils.NewErrorResponse("failed to create plan"))
		return
	}

	ctx.Status(http.StatusCreated)
}

func (h *TrainingPlanHandler) ListPlan(ctx *gin.Context) {
	authorId := ctx.Param("authorId")
	cursor, limit := h.getPagination(ctx)

	plans, nextCursor, err := h.srv.ListPlan(ctx.Request.Context(), authorId, cursor, limit)
	if err != nil {
		slog.ErrorContext(ctx.Request.Context(), "failed to list plans", slog.Any("error", err), slog.String("authorId", authorId))
		ctx.JSON(http.StatusInternalServerError, utils.NewErrorResponse("failed to list plans"))
		return
	}

	ctx.JSON(http.StatusOK, utils.NewPaginatedResponse(plans, nextCursor))
}

func (h *TrainingPlanHandler) UpdatePlan(ctx *gin.Context) {
	id := ctx.Param("id")
	if id == "" {
		ctx.JSON(http.StatusBadRequest, utils.NewErrorResponse("plan id is required"))
		return
	}

	var plan TrainingPlan
	if err := ctx.ShouldBindJSON(&plan); err != nil {
		ctx.JSON(http.StatusBadRequest, utils.NewErrorResponse(err.Error()))
		return
	}

	updatedPlan, err := h.srv.UpdatePlan(ctx.Request.Context(), id, plan)
	if err != nil {
		slog.ErrorContext(ctx.Request.Context(), "failed to update plan", slog.Any("error", err), slog.String("plan_id", id))
		ctx.JSON(http.StatusInternalServerError, utils.NewErrorResponse("failed to update plan"))
		return
	}

	ctx.JSON(http.StatusOK, updatedPlan)
}

func (h *TrainingPlanHandler) DeletePlan(ctx *gin.Context) {
	id := ctx.Param("id")

	if err := h.srv.DeletePlan(ctx.Request.Context(), id); err != nil {
		slog.ErrorContext(ctx.Request.Context(), "failed to delete plan", slog.Any("error", err), slog.String("plan_id", id))
		ctx.JSON(http.StatusInternalServerError, utils.NewErrorResponse("failed to delete plan"))
		return
	}

	ctx.Status(http.StatusOK)
}

func (h *TrainingPlanHandler) ExistsPlan(ctx *gin.Context) {
	id := ctx.Param("id")
	if id == "" {
		ctx.JSON(http.StatusBadRequest, utils.NewErrorResponse("plan id is required"))
		return
	}

	exists, err := h.srv.ExistsPlan(ctx.Request.Context(), id)
	if err != nil {
		slog.ErrorContext(ctx.Request.Context(), "failed to check plan existence", slog.Any("error", err), slog.String("plan_id", id))
		ctx.JSON(http.StatusInternalServerError, utils.NewErrorResponse("failed to check plan existence"))
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"exists": exists})
}

func (h *TrainingPlanHandler) GetPlan(ctx *gin.Context) {
	id := ctx.Param("id")
	if id == "" {
		ctx.JSON(http.StatusBadRequest, utils.NewErrorResponse("plan id is required"))
		return
	}

	plan, err := h.srv.GetPlan(ctx.Request.Context(), id)
	if err != nil {
		slog.ErrorContext(ctx.Request.Context(), "failed to get plan", slog.Any("error", err), slog.String("plan_id", id))
		ctx.JSON(http.StatusInternalServerError, utils.NewErrorResponse("failed to get plan"))
		return
	}

	if plan == nil {
		ctx.JSON(http.StatusNotFound, utils.NewErrorResponse("training plan not found"))
		return
	}

	ctx.JSON(http.StatusOK, plan)
}

func (h *TrainingPlanHandler) CreateDay(ctx *gin.Context) {
	var day Day
	if err := ctx.ShouldBindJSON(&day); err != nil {
		ctx.JSON(http.StatusBadRequest, utils.NewErrorResponse(err.Error()))
		return
	}

	if err := h.srv.CreateDay(ctx.Request.Context(), &day); err != nil {
		slog.ErrorContext(ctx.Request.Context(), "failed to create day", slog.Any("error", err))
		ctx.JSON(http.StatusInternalServerError, utils.NewErrorResponse("failed to create day"))
		return
	}

	ctx.Status(http.StatusCreated)
}

func (h *TrainingPlanHandler) CreateExercise(ctx *gin.Context) {
	var exercise Exercise
	if err := ctx.ShouldBindJSON(&exercise); err != nil {
		ctx.JSON(http.StatusBadRequest, utils.NewErrorResponse(err.Error()))
		return
	}

	if err := h.srv.CreateExercise(ctx.Request.Context(), &exercise); err != nil {
		slog.ErrorContext(ctx.Request.Context(), "failed to create exercise", slog.Any("error", err))
		ctx.JSON(http.StatusInternalServerError, utils.NewErrorResponse("failed to create exercise"))
		return
	}

	ctx.Status(http.StatusCreated)
}

func (h *TrainingPlanHandler) DeleteDay(ctx *gin.Context) {
	dayId := ctx.Param("dayId")

	if err := h.srv.DeleteDay(ctx.Request.Context(), dayId); err != nil {
		slog.ErrorContext(ctx.Request.Context(), "failed to delete day", slog.Any("error", err), slog.String("day_id", dayId))
		ctx.JSON(http.StatusInternalServerError, utils.NewErrorResponse("failed to delete day"))
		return
	}

	ctx.Status(http.StatusOK)
}

func (h *TrainingPlanHandler) DeleteExercise(ctx *gin.Context) {
	id := ctx.Param("exerciseId")

	if err := h.srv.DeleteExercise(ctx.Request.Context(), id); err != nil {
		slog.ErrorContext(ctx.Request.Context(), "failed to delete exercise", slog.Any("error", err), slog.String("exercise_id", id))
		ctx.JSON(http.StatusInternalServerError, utils.NewErrorResponse("failed to delete exercise"))
		return
	}

	ctx.Status(http.StatusOK)
}

func (h *TrainingPlanHandler) ListActivityWeekly(ctx *gin.Context) {
	user, ok := h.getAuthUser(ctx)
	if !ok {
		return
	}

	activity, err := h.srv.ListActivityWeekly(ctx.Request.Context(), user.ID)
	if err != nil {
		slog.ErrorContext(ctx.Request.Context(), "failed to list weekly activity", slog.Any("error", err), slog.String("user_id", user.ID))
		ctx.JSON(http.StatusInternalServerError, utils.NewErrorResponse("failed to list weekly activity"))
		return
	}

	ctx.JSON(http.StatusOK, activity)
}

func (h *TrainingPlanHandler) ListPlansByIds(ctx *gin.Context) {
	var ids []string
	if err := ctx.ShouldBindJSON(&ids); err != nil {
		ctx.JSON(http.StatusBadRequest, utils.NewErrorResponse(err.Error()))
		return
	}

	res, err := h.srv.ListPlansByIds(ctx.Request.Context(), ids)
	if err != nil {
		slog.ErrorContext(ctx.Request.Context(), "failed to list training plans", slog.Any("error", err))
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "failed to list training plans"})
		return
	}

	ctx.JSON(http.StatusOK, res)
}

func (h *TrainingPlanHandler) getAuthUser(ctx *gin.Context) (auth.AuthUser, bool) {
	user, ok := ctx.Value(string(auth.UserContextKey)).(auth.AuthUser)
	if !ok {
		ctx.JSON(http.StatusUnauthorized, utils.NewErrorResponse("unauthorized"))
		return auth.AuthUser{}, false
	}

	return user, true
}

func (h *TrainingPlanHandler) getPagination(ctx *gin.Context) (string, int) {
	cursor := ctx.Query("cursor")
	limit, _ := strconv.Atoi(ctx.DefaultQuery("limit", "20"))
	if limit <= 0 {
		limit = 20
	}
	return cursor, limit
}
