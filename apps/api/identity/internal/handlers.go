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
	r.GET("/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"status": "ok", "service": "identity"})
	})
	r.POST("/identity/auth/register", h.Register)
	r.POST("/identity/auth/login", h.Login)
	r.POST("/identity/auth/reset-password/send-token", h.ResetPasswordSendToken)
	r.POST("/identity/auth/reset-password/verify-token", h.ResetPasswordVerifyToken)
	r.POST("/identity/auth/reset-password", h.ResetPassword)

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

		protected.PUT("/profile/picture", h.UploadProfilePicture)
		protected.DELETE("/profile/picture", h.RemoveProfilePicture)

		protected.POST("/profile/upgrade", h.ChangeToTrainer)
		protected.POST("/profile/downgrade", h.ChangeToClient)

		protected.PATCH("/trainers/profile/code", h.CreateTrainerCode)
		protected.POST("/trainers/profile/link", h.LinkTrainer)
		protected.POST("/trainers/profile/unlink", h.UnlinkTrainer)

		protected.PATCH("/trainers/body-measurements/:id/notes", h.AddBodyMeasurementNote)
		protected.GET("/trainers/body-measurements/latest", h.FindLastBodyMeasurementNote)
		protected.GET("/trainers/body-measurements", h.ListBodyMeasurements)
		protected.GET("/trainers/students/:id/body-measurements", h.ListBodyMeasurements)

		protected.PATCH("/trainers/weight-log/:id/notes", h.AddWeightLogNote)
		protected.GET("/trainers/weight-logs", h.ListWeightLogs)
		protected.GET("/trainers/students/:id/weight-logs", h.ListWeightLogs)

		protected.POST("/trainers/goals", h.AddGoalMetric)
		protected.GET("/trainers/goals", h.ListGoalsMetric)
		protected.GET("/trainers/students/:id/goals", h.ListGoalsMetricById)

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

func (h *UserHandler) ResetPasswordSendToken(ctx *gin.Context) {
	var body struct {
		Email string `json:"email" binding:"required,email"`
	}

	if err := ctx.ShouldBindJSON(&body); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := h.srv.ResetPasswordSendToken(ctx.Request.Context(), body.Email); err != nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}

	ctx.Status(http.StatusOK)
}

func (h *UserHandler) ResetPasswordVerifyToken(ctx *gin.Context) {
	var body struct {
		Email string `json:"email" binding:"required,email"`
		Code  string `json:"code" binding:"required"`
	}

	if err := ctx.ShouldBindJSON(&body); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	res, err := h.srv.ResetPasswordVerifyToken(ctx.Request.Context(), body.Email, body.Code)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"valid": res})
}

func (h *UserHandler) ResetPassword(ctx *gin.Context) {
	var body struct {
		Email       string `json:"email" binding:"required,email"`
		Code        string `json:"code" binding:"required"`
		NewPassword string `json:"password" binding:"required"`
	}

	if err := ctx.ShouldBindJSON(&body); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := h.srv.ResetPassword(ctx.Request.Context(), body.Email, body.Code, body.NewPassword); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.Status(http.StatusOK)
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
		Code string `json:"code" validate:"required"`
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

func (h *UserHandler) ListBodyMeasurements(ctx *gin.Context) {
	id := ctx.Param("id")
	if id == "" {
		user, ok := ctx.Value(string(auth.UserContextKey)).(auth.AuthUser)
		if !ok {
			ctx.JSON(http.StatusUnauthorized, gin.H{"error": "unauthorized"})
			return
		}
		id = user.ID
	}

	cursor := ctx.Query("cursor")
	limitStr := ctx.DefaultQuery("limit", "20")

	limit, err := strconv.Atoi(limitStr)
	if err != nil {
		limit = 20
	}

	measurements, nextCursor, err := h.srv.ListBodyMeasurements(ctx.Request.Context(), id, cursor, limit)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"data":       measurements,
		"nextCursor": nextCursor,
	})
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

	cursor := ctx.Query("cursor")
	limitStr := ctx.DefaultQuery("limit", "20")

	limit, err := strconv.Atoi(limitStr)
	if err != nil {
		limit = 20
	}

	goals, nextCursor, err := h.srv.ListGoalsMetric(ctx.Request.Context(), user.ID, cursor, limit)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"data":       goals,
		"nextCursor": nextCursor,
	})
}

func (h *UserHandler) ListGoalsMetricById(ctx *gin.Context) {
	id := ctx.Param("id")

	cursor := ctx.Query("cursor")
	limitStr := ctx.DefaultQuery("limit", "20")

	limit, err := strconv.Atoi(limitStr)
	if err != nil {
		limit = 20
	}

	goals, nextCursor, err := h.srv.ListGoalsMetric(ctx.Request.Context(), id, cursor, limit)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"data":       goals,
		"nextCursor": nextCursor,
	})
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

	if err := h.srv.ChangeToClient(ctx.Request.Context(), user.ID); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.Status(http.StatusOK)
}

func (h *UserHandler) RemoveProfilePicture(ctx *gin.Context) {
	user, ok := ctx.Value(string(auth.UserContextKey)).(auth.AuthUser)
	if !ok {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "unauthorized"})
		return
	}

	if err := h.srv.RemoveProfilePicture(ctx.Request.Context(), user.ID); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.Status(http.StatusNoContent)
}

func (h *UserHandler) ListWeightLogs(ctx *gin.Context) {
	id := ctx.Param("id")
	if id == "" {
		user, ok := ctx.Value(string(auth.UserContextKey)).(auth.AuthUser)
		if !ok {
			ctx.JSON(http.StatusUnauthorized, gin.H{"error": "unauthorized"})
			return
		}
		id = user.ID
	}

	cursor := ctx.Query("cursor")
	limitStr := ctx.DefaultQuery("limit", "20")

	limit, err := strconv.Atoi(limitStr)
	if err != nil {
		limit = 20
	}

	logs, nextCursor, err := h.srv.ListWeightLogs(ctx.Request.Context(), id, cursor, limit)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"data":       logs,
		"nextCursor": nextCursor,
	})
}

func (h *UserHandler) UploadProfilePicture(ctx *gin.Context) {
	user, ok := ctx.Value(string(auth.UserContextKey)).(auth.AuthUser)
	if !ok {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "unauthorized"})
		return
	}

	file, err := ctx.FormFile("file")
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "file is required"})
	}

	openedFile, err := file.Open()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "failed to open file"})
	}
	defer openedFile.Close()

	if err := h.srv.UploadProfilePicture(ctx.Request.Context(), user.ID, openedFile); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.Status(http.StatusOK)
}
