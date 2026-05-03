package internal

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/kaua-nasc/gymtrack-go/libs/auth"
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
	protected := r.Group("/training-plans")
	protected.Use(auth.AuthMiddleware())
	{
		protected.GET("", h.ListPlan)
		protected.POST("", h.CreatePlan)
		protected.PUT("/:id", h.UpdatePlan)
		protected.GET("/exists/:id", h.ExistsPlan)
		protected.GET("/:id", h.GetPlan)
		protected.POST("/:id/like", h.LikePlan)
		protected.DELETE("/:id/like", h.UnlikePlan)
		protected.GET("/:id/comments", h.ListPlanComment)
		protected.POST("/:id/comments", h.AddPlanComment)
		protected.DELETE("/:planId/comments/:commentId", h.RemovePlanComment)
		protected.POST("/:id/subscribe", h.Subscribe)
		protected.POST("/:id/days/:dayId/complete", h.CompleteDay)
		protected.POST("/:id/feedback", h.AddFeedback)
	}

	exercises := r.Group("/exercises")
	exercises.Use(auth.AuthMiddleware())
	{
		exercises.POST("/:id/log", h.LogExercise)
	}
}

func (h *TrainingPlanHandler) AddFeedback(ctx *gin.Context) {
	id := ctx.Param("id")
	user, ok := ctx.Value(string(auth.UserContextKey)).(auth.AuthUser)
	if !ok {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "unauthorized"})
		return
	}

	var body struct {
		Rating  float64 `json:"rating" binding:"required,min=0,max=5"`
		Message *string `json:"message"`
	}

	if err := ctx.ShouldBindJSON(&body); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := h.srv.AddFeedback(ctx.Request.Context(), id, user.ID, body.Rating, body.Message); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.Status(http.StatusNoContent)
}

func (h *TrainingPlanHandler) LogExercise(ctx *gin.Context) {
	id := ctx.Param("id")
	user, ok := ctx.Value(string(auth.UserContextKey)).(auth.AuthUser)
	if !ok {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "unauthorized"})
		return
	}

	var body struct {
		Reps   []int     `json:"reps" binding:"required"`
		Weight []float64 `json:"weight" binding:"required"`
		Notes  *string   `json:"notes"`
	}

	if err := ctx.ShouldBindJSON(&body); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := h.srv.LogExercise(ctx.Request.Context(), id, user.ID, body.Reps, body.Weight, body.Notes); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.Status(http.StatusNoContent)
}

func (h *TrainingPlanHandler) Subscribe(ctx *gin.Context) {
	id := ctx.Param("id")
	user, ok := ctx.Value(string(auth.UserContextKey)).(auth.AuthUser)
	if !ok {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "unauthorized"})
		return
	}

	var body struct {
		Type PlanSubscriptionType `json:"type" binding:"required"`
	}

	if err := ctx.ShouldBindJSON(&body); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := h.srv.Subscribe(ctx.Request.Context(), id, user.ID, body.Type); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.Status(http.StatusNoContent)
}

func (h *TrainingPlanHandler) CompleteDay(ctx *gin.Context) {
	planId := ctx.Param("id")
	dayId := ctx.Param("dayId")
	user, ok := ctx.Value(string(auth.UserContextKey)).(auth.AuthUser)
	if !ok {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "unauthorized"})
		return
	}

	if err := h.srv.CompleteDay(ctx.Request.Context(), planId, user.ID, dayId); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.Status(http.StatusNoContent)
}

func (h *TrainingPlanHandler) CreatePlan(ctx *gin.Context) {
	var plan TrainingPlan
	if err := ctx.ShouldBindJSON(&plan); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user, ok := ctx.Value(string(auth.UserContextKey)).(auth.AuthUser)
	if !ok {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "unauthorized: user not found in context"})
		return
	}

	if _, err := h.srv.CreatePlan(ctx.Request.Context(), plan, user); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusCreated, plan)
}

func (h *TrainingPlanHandler) ListPlan(ctx *gin.Context) {
	authorId := ctx.Query("authorId")
	cursor := ctx.Query("cursor")
	limitStr := ctx.DefaultQuery("limit", "20")

	limit, err := strconv.Atoi(limitStr)
	if err != nil {
		limit = 20
	}

	plans, nextCursor, err := h.srv.ListPlan(ctx.Request.Context(), authorId, cursor, limit)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"data":   plans,
		"cursor": nextCursor,
	})
}

func (h *TrainingPlanHandler) UpdatePlan(ctx *gin.Context) {
	id := ctx.Param("id")
	if id == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "plan id is required"})
		return
	}

	var plan TrainingPlan
	if err := ctx.ShouldBindJSON(&plan); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	updatedPlan, err := h.srv.UpdatePlan(ctx.Request.Context(), id, plan)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, updatedPlan)
}

func (h *TrainingPlanHandler) ExistsPlan(ctx *gin.Context) {
	id := ctx.Param("id")
	if id == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "plan id is required"})
		return
	}

	exists, err := h.srv.ExistsPlan(ctx.Request.Context(), id)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"exists": exists})
}

func (h *TrainingPlanHandler) GetPlan(ctx *gin.Context) {
	id := ctx.Param("id")
	if id == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "plan id is required"})
		return
	}

	plan, err := h.srv.GetPlan(ctx.Request.Context(), id)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	if plan == nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "training plan not found"})
		return
	}

	ctx.JSON(http.StatusOK, plan)
}

func (h *TrainingPlanHandler) LikePlan(ctx *gin.Context) {
	id := ctx.Param("id")
	user, ok := ctx.Value(string(auth.UserContextKey)).(auth.AuthUser)
	if !ok {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "unauthorized"})
		return
	}

	if err := h.srv.LikePlan(ctx.Request.Context(), id, user.ID); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.Status(http.StatusNoContent)
}

func (h *TrainingPlanHandler) UnlikePlan(ctx *gin.Context) {
	id := ctx.Param("id")
	user, ok := ctx.Value(string(auth.UserContextKey)).(auth.AuthUser)
	if !ok {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "unauthorized"})
		return
	}

	if err := h.srv.UnlikePlan(ctx.Request.Context(), id, user.ID); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.Status(http.StatusNoContent)
}

func (h *TrainingPlanHandler) ListPlanComment(ctx *gin.Context) {
	id := ctx.Param("id")
	cursor := ctx.Query("cursor")
	limitStr := ctx.DefaultQuery("limit", "20")

	limit, err := strconv.Atoi(limitStr)
	if err != nil {
		limit = 20
	}

	comments, nextCursor, err := h.srv.ListPlanComments(ctx.Request.Context(), id, cursor, limit)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"data":   comments,
		"cursor": nextCursor,
	})
}

func (h *TrainingPlanHandler) AddPlanComment(ctx *gin.Context) {
	id := ctx.Param("id")
	user, ok := ctx.Value(string(auth.UserContextKey)).(auth.AuthUser)
	if !ok {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "unauthorized"})
		return
	}

	var body struct {
		Content string `json:"content" binding:"required"`
	}

	if err := ctx.ShouldBindJSON(&body); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	comment, err := h.srv.AddPlanComment(ctx.Request.Context(), id, body.Content, user.ID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusCreated, comment)
}

func (h *TrainingPlanHandler) RemovePlanComment(ctx *gin.Context) {
	commentId := ctx.Param("commentId")
	user, ok := ctx.Value(string(auth.UserContextKey)).(auth.AuthUser)
	if !ok {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "unauthorized"})
		return
	}

	if err := h.srv.RemovePlanComment(ctx.Request.Context(), commentId, user.ID); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.Status(http.StatusNoContent)
}
