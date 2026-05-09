package internal

import (
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/kaua-nasc/gymtrack-go/libs/auth"
)

type UserHandler struct {
	srv *UserService
}

func NewUserHandler(srv *UserService) *UserHandler {
	return &UserHandler{srv: srv}
}

func (h *UserHandler) RegisterRoutes(r *gin.Engine) {
	r.POST("/identity/auth/register", h.Register)
	r.POST("/identity/auth/login", h.Login)

	// Protected routes
	protected := r.Group("/identity/users")
	protected.Use(auth.AuthMiddleware())
	{
		protected.GET("", h.ListUsers)
		protected.GET("/:id", h.GetUser)

		protected.GET("/:id/followers/count", h.CountFollowers)
		protected.GET("/:id/following/count", h.CountFollowing)
		protected.POST("/:id/follows", h.FollowUser)
		protected.POST("/:id/unfollows", h.UnfollowUser)

		protected.POST("/profile/upgrade", h.ChangeToTrainer)
		protected.POST("/profile/downgrade", h.ChangeToClient)

		protected.PATCH("/trainers/profile/code", h.CreateTrainerCode)
		protected.POST("/trainers/profile/link", h.LinkTrainer)
		protected.POST("/trainers/profile/unlink", h.UnlinkTrainer)

		protected.PATCH("/trainers/body-measurements/:id/notes", h.AddBodyMeasurementNote)
		protected.GET("/trainers/body-measurements/latest", h.FindLastBodyMeasurementNote)

		protected.PATCH("/trainers/weight-log/:id/notes", h.AddBodyMeasurementNote)
		protected.GET("/trainers/weight-history", h.ListWeightHistory)

		protected.POST("/trainers/goals", h.AddGoalMetric)
		protected.GET("/trainers/goals", h.ListGoalsMetric)

		protected.GET("/students", h.ListStudents)
		protected.POST("/students/:id/profile/unlink", h.UnlinkStudant)
	}
}

func (h *UserHandler) Login(ctx *gin.Context) {
	var body struct {
		Email    string `json:"email" binding:"required,email"`
		Password string `json:"password" binding:"required"`
	}

	if err := ctx.ShouldBindJSON(&body); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	token, err := h.srv.Login(ctx.Request.Context(), body.Email, body.Password)
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"accessToken": token})
}

func (h *UserHandler) ListUsers(ctx *gin.Context) {
	// ... rest of the code

	idsStr := ctx.Query("ids")
	if idsStr == "" {
		ctx.JSON(http.StatusOK, []User{})
		return
	}

	ids := strings.Split(idsStr, ",")
	res, err := h.srv.ListUsers(ctx.Request.Context(), ids)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, res)
}

func (h *UserHandler) Register(ctx *gin.Context) {
	var u User
	if err := ctx.ShouldBindJSON(&u); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err := h.srv.Register(ctx.Request.Context(), u)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.Status(http.StatusCreated)
}

func (h *UserHandler) GetUser(ctx *gin.Context) {
	id := ctx.Param("id")
	res, err := h.srv.GetUser(ctx.Request.Context(), id)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	if res == nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "user not found"})
		return
	}

	ctx.JSON(http.StatusOK, res)
}

func (h *UserHandler) CountFollowers(ctx *gin.Context) {
	id := ctx.Param("id")

	count, err := h.srv.CountFollowers(ctx.Request.Context(), id)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"count": count})
}

func (h *UserHandler) CountFollowing(ctx *gin.Context) {
	id := ctx.Param("id")

	count, err := h.srv.CountFollowing(ctx.Request.Context(), id)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"count": count})
}

func (h *UserHandler) FollowUser(ctx *gin.Context) {
	id := ctx.Param("id")

	user, ok := ctx.Value(string(auth.UserContextKey)).(auth.AuthUser)
	if !ok {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "unauthorized"})
		return
	}

	if err := h.srv.FollowUser(ctx.Request.Context(), user.ID, id); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.Status(http.StatusOK)
}

func (h *UserHandler) UnfollowUser(ctx *gin.Context) {
	id := ctx.Param("id")

	user, ok := ctx.Value(string(auth.UserContextKey)).(auth.AuthUser)
	if !ok {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "unauthorized"})
		return
	}

	if err := h.srv.UnfollowUser(ctx.Request.Context(), user.ID, id); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.Status(http.StatusOK)
}

func (h *UserHandler) LinkTrainer(ctx *gin.Context) {
	user, ok := ctx.Value(string(auth.UserContextKey)).(auth.AuthUser)
	if !ok {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "unauthorized"})
		return
	}

	var body struct {
		Code string `json:"id" validate:"required"`
	}

	if err := ctx.ShouldBindJSON(&body); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := h.srv.LinkTrainer(ctx.Request.Context(), user.ID, body.Code); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.Status(http.StatusOK)
}

func (h *UserHandler) CreateTrainerCode(ctx *gin.Context) {
	user, ok := ctx.Value(string(auth.UserContextKey)).(auth.AuthUser)
	if !ok {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "unauthorized"})
		return
	}

	var body struct {
		Code string `json:"id" validate:"required"`
	}

	if err := ctx.ShouldBindJSON(&body); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := h.srv.CreateTrainerCode(ctx.Request.Context(), user.ID, body.Code); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.Status(http.StatusOK)
}

func (h *UserHandler) UnlinkTrainer(ctx *gin.Context) {
	user, ok := ctx.Value(string(auth.UserContextKey)).(auth.AuthUser)
	if !ok {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "unauthorized"})
		return
	}

	if err := h.srv.UnlinkTrainer(ctx.Request.Context(), user.ID); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.Status(http.StatusOK)
}

func (h *UserHandler) ListStudents(ctx *gin.Context) {
	user, ok := ctx.Value(string(auth.UserContextKey)).(auth.AuthUser)
	if !ok {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "unauthorized"})
		return
	}

	if user.Type != "PERSONAL_TRAINER" {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "nao eh o personal, nao consegue ver alunos"})
		return
	}

	cursor := ctx.Query("cursor")
	limitStr := ctx.DefaultQuery("limit", "20")

	limit, err := strconv.Atoi(limitStr)
	if err != nil {
		limit = 20
	}

	users, nextCursor, err := h.srv.ListStudents(ctx.Request.Context(), user.ID, cursor, limit)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"data":       users,
		"nextCursor": nextCursor,
	})
}

func (h *UserHandler) UnlinkStudant(ctx *gin.Context) {
	id := ctx.Param("id")

	if err := h.srv.UnlinkStudant(ctx.Request.Context(), id); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.Status(http.StatusOK)
}

func (h *UserHandler) AddBodyMeasurementNote(ctx *gin.Context) {
	id := ctx.Param("id")

	var body struct {
		Note string `json:"note" validate:"required"`
	}

	if err := ctx.ShouldBindJSON(&body); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := h.srv.AddBodyMeasurementNote(ctx.Request.Context(), id, body.Note); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.Status(http.StatusOK)
}

func (h *UserHandler) FindLastBodyMeasurementNote(ctx *gin.Context) {
	user, ok := ctx.Value(string(auth.UserContextKey)).(auth.AuthUser)
	if !ok {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "unauthorized"})
		return
	}

	measurement, err := h.srv.FindLastBodyMeasurementNote(ctx.Request.Context(), user.ID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, measurement)
}

func (h *UserHandler) AddGoalMetric(ctx *gin.Context) {
	user, ok := ctx.Value(string(auth.UserContextKey)).(auth.AuthUser)
	if !ok {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "unauthorized"})
		return
	}

	var body struct {
		Type        string    `json:"type" validate:"required"`
		TargetValue float64   `json:"targetValue" validate:"required,gt=0"`
		Deadline    time.Time `json:"deadline" validate:"required, gt"`
	}

	if err := ctx.ShouldBindJSON(&body); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := h.srv.AddGoalMetric(ctx.Request.Context(), &MetricGoal{
		UserId:      user.ID,
		Type:        body.Type,
		TargetValue: body.TargetValue,
		Deadline:    &body.Deadline,
	}); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.Status(http.StatusOK)
}

func (h *UserHandler) ListGoalsMetric(ctx *gin.Context) {
	user, ok := ctx.Value(string(auth.UserContextKey)).(auth.AuthUser)
	if !ok {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "unauthorized"})
		return
	}

	goals, err := h.srv.ListGoalsMetric(ctx.Request.Context(), user.ID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, goals)
}

func (h *UserHandler) AddWeightLogNote(ctx *gin.Context) {
	id := ctx.Param("id")

	var body struct {
		Note string `json:"note" validate:"required"`
	}

	if err := ctx.ShouldBindJSON(&body); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := h.srv.AddWeightLogNote(ctx.Request.Context(), id, body.Note); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.Status(http.StatusOK)
}

func (h *UserHandler) ChangeToTrainer(ctx *gin.Context) {
	user, ok := ctx.Value(string(auth.UserContextKey)).(auth.AuthUser)
	if !ok {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "unauthorized"})
		return
	}

	var body struct {
		Cref string `json:"cref" validate:"required"`
	}

	if err := ctx.ShouldBindJSON(&body); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := h.srv.ChangeToTrainer(ctx.Request.Context(), user.ID, body.Cref); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.Status(http.StatusOK)
}

func (h *UserHandler) ChangeToClient(ctx *gin.Context) {
	user, ok := ctx.Value(string(auth.UserContextKey)).(auth.AuthUser)
	if !ok {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "unauthorized"})
		return
	}

	var body struct {
		Cref string `json:"cref" validate:"required"`
	}

	if err := ctx.ShouldBindJSON(&body); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := h.srv.ChangeToClient(ctx.Request.Context(), user.ID); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.Status(http.StatusOK)
}

func (h *UserHandler) ListWeightHistory(ctx *gin.Context) {
	user, ok := ctx.Value(string(auth.UserContextKey)).(auth.AuthUser)
	if !ok {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "unauthorized"})
		return
	}

	cursor := ctx.Query("cursor")
	limitStr := ctx.DefaultQuery("limit", "20")

	limit, err := strconv.Atoi(limitStr)
	if err != nil {
		limit = 20
	}

	history, nextCursor, err := h.srv.ListWeightHistory(ctx.Request.Context(), user.ID, cursor, limit)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"data":       history,
		"nextCursor": nextCursor,
	})
}
