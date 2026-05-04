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
	r.POST("/users/register", h.Register)
	r.POST("/identity/auth/login", h.Login)

	// Protected routes
	protected := r.Group("/identity/users")
	protected.Use(auth.AuthMiddleware())
	{
		protected.GET("", h.ListUsers)
		protected.GET("/:id", h.GetUser)
	}
}

func (h *UserHandler) Login(ctx *gin.Context) {
	// ... (implementation same, just ensure it uses body.Email etc)
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

	res, err := h.srv.Register(ctx.Request.Context(), u)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusCreated, res)
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
