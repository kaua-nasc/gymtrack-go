package internal

import (
	"net/http"
	"strings"

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

		protected.POST("/:id/follows", h.FollowUser)
		protected.POST("/:id/unfollows", h.UnfollowUser)

		protected.POST("/profile/upgrade", h.ChangeToTrainer)
		protected.POST("/profile/downgrade", h.ChangeToClient)

		protected.PATCH("/trainers/profile/code", h.CreateTrainerCode)
		protected.POST("/trainers/profile/link", h.LinkTrainer)
		protected.POST("/trainers/profile/unlink", h.UnlinkTrainer)

		protected.PATCH("/trainers/body-measurement/:id/notes", h.AddBodyMeasurementNote)

		protected.PATCH("/trainers/weight-log/:id/notes", h.AddBodyMeasurementNote)

		protected.PATCH("/trainers/goals/:id/status", h.AddBodyMeasurementNote)

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
