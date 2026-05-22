package auth

import (
	"errors"
	"log/slog"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/kaua-nasc/gymtrack-go/apps/api/identity/internal/domain"
	"github.com/kaua-nasc/gymtrack-go/libs/log"
)

type Handler struct {
	srv *Service
}

func NewHandler(srv *Service) *Handler {
	return &Handler{srv: srv}
}

func (h *Handler) RegisterRoutes(r *gin.Engine) {
	r.Use(log.LoggerMiddleware())
	r.GET("/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"status": "ok", "service": "identity"})
	})
	r.POST("/identity/auth/register", h.Register)
	r.POST("/identity/auth/login", h.Login)
	r.POST("/identity/auth/verify/send-token", h.SendVerificationEmail)
	r.POST("/identity/auth/verify", h.VerifyEmail)
	r.POST("/identity/auth/reset-password/send-token", h.ResetPasswordSendToken)
	r.POST("/identity/auth/reset-password/verify-token", h.ResetPasswordVerifyToken)
	r.POST("/identity/auth/reset-password", h.ResetPassword)
}

func (h *Handler) SendVerificationEmail(ctx *gin.Context) {
	var body struct {
		Email string `json:"email" binding:"required,email"`
	}

	if err := ctx.ShouldBindJSON(&body); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := h.srv.SendVerificationEmail(ctx.Request.Context(), body.Email); err != nil {
		if errors.Is(err, domain.ErrEmailNotFound) {
			ctx.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
			return
		}
		if errors.Is(err, domain.ErrAlreadyVerified) {
			ctx.JSON(http.StatusConflict, gin.H{"error": err.Error()})
			return
		}
		slog.ErrorContext(ctx.Request.Context(), "failed to send verification token", slog.Any("error", err), slog.String("email", body.Email))
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "failed to send verification token"})
		return
	}

	ctx.Status(http.StatusOK)
}

func (h *Handler) VerifyEmail(ctx *gin.Context) {
	var body struct {
		Email string `json:"email" binding:"required,email"`
		Code  string `json:"code" binding:"required"`
	}

	if err := ctx.ShouldBindJSON(&body); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := h.srv.VerifyEmail(ctx.Request.Context(), body.Email, body.Code); err != nil {
		if errors.Is(err, domain.ErrInvalidCode) {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		if errors.Is(err, domain.ErrUserNotFound) {
			ctx.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
			return
		}
		if errors.Is(err, domain.ErrAlreadyVerified) {
			ctx.JSON(http.StatusConflict, gin.H{"error": err.Error()})
			return
		}
		slog.ErrorContext(ctx.Request.Context(), "failed to verify email", slog.Any("error", err), slog.String("email", body.Email))
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "failed to verify email"})
		return
	}

	ctx.Status(http.StatusOK)
}

func (h *Handler) Login(ctx *gin.Context) {
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
		if errors.Is(err, domain.ErrInvalidCredentials) {
			ctx.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
			return
		}
		slog.ErrorContext(ctx.Request.Context(), "login failed", slog.Any("error", err), slog.String("email", body.Email))
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "failed to login"})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"accessToken": token})
}

func (h *Handler) ResetPasswordSendToken(ctx *gin.Context) {
	var body struct {
		Email string `json:"email" binding:"required,email"`
	}

	if err := ctx.ShouldBindJSON(&body); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := h.srv.ResetPasswordSendToken(ctx.Request.Context(), body.Email); err != nil {
		if errors.Is(err, domain.ErrEmailNotFound) {
			ctx.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
			return
		}
		slog.ErrorContext(ctx.Request.Context(), "failed to send reset token", slog.Any("error", err), slog.String("email", body.Email))
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "failed to send reset token"})
		return
	}

	ctx.Status(http.StatusOK)
}

func (h *Handler) ResetPasswordVerifyToken(ctx *gin.Context) {
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
		if errors.Is(err, domain.ErrInvalidCode) {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		slog.ErrorContext(ctx.Request.Context(), "failed to verify reset token", slog.Any("error", err), slog.String("email", body.Email))
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "failed to verify reset token"})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"valid": res})
}

func (h *Handler) ResetPassword(ctx *gin.Context) {
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
		if errors.Is(err, domain.ErrInvalidCode) {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		if errors.Is(err, domain.ErrUserNotFound) {
			ctx.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
			return
		}
		slog.ErrorContext(ctx.Request.Context(), "failed to reset password", slog.Any("error", err), slog.String("email", body.Email))
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "failed to reset password"})
		return
	}

	ctx.Status(http.StatusOK)
}

func (h *Handler) Register(ctx *gin.Context) {
	var body struct {
		FirstName string `json:"firstName" binding:"required,min=1,max=255"`
		LastName  string `json:"lastName" binding:"required,min=1,max=255"`
		Email     string `json:"email" binding:"required,email"`
		Password  string `json:"password" binding:"required,min=8"`
	}

	if err := ctx.ShouldBindJSON(&body); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err := h.srv.Register(ctx.Request.Context(), domain.User{
		FirstName: body.FirstName,
		LastName:  body.LastName,
		Email:     body.Email,
		Password:  body.Password,
	})

	if err != nil {
		if errors.Is(err, domain.ErrUserAlreadyExists) {
			ctx.JSON(http.StatusConflict, gin.H{"error": err.Error()})
			return
		}
		slog.ErrorContext(ctx.Request.Context(), "user registration failed", slog.Any("error", err), slog.String("email", body.Email))
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "failed to register user"})
		return
	}

	ctx.Status(http.StatusCreated)
}
