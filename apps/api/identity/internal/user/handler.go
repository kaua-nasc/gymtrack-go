package user

import (
	"errors"
	"log/slog"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/kaua-nasc/gymtrack-go/apps/api/identity/internal/domain"
	"github.com/kaua-nasc/gymtrack-go/libs/auth"
)

type Handler struct {
	srv *Service
}

func NewHandler(srv *Service) *Handler {
	return &Handler{srv: srv}
}

func (h *Handler) RegisterRoutes(r *gin.Engine) {
	protected := r.Group("/identity/users")
	protected.Use(auth.AuthMiddleware())
	{
		protected.GET("", h.ListUsers)
		protected.GET("/:id", h.GetUser)
		protected.PUT("/:id", h.UpdateProfile)
		protected.PUT("/profile/password", h.ChangePassword)
		protected.PUT("/profile/picture", h.UploadProfilePicture)
		protected.DELETE("/profile/picture", h.RemoveProfilePicture)
		protected.POST("/profile/upgrade", h.ChangeToTrainer)
		protected.POST("/profile/downgrade", h.ChangeToClient)
	}
}

func (h *Handler) ListUsers(ctx *gin.Context) {
	idsStr := ctx.Query("ids")
	if idsStr == "" {
		ctx.JSON(http.StatusOK, []domain.User{})
		return
	}

	ids := strings.Split(idsStr, ",")
	res, err := h.srv.ListUsers(ctx.Request.Context(), ids)
	if err != nil {
		slog.ErrorContext(ctx.Request.Context(), "failed to list users", slog.Any("error", err), slog.String("ids", idsStr))
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "failed to list users"})
		return
	}

	ctx.JSON(http.StatusOK, res)
}

func (h *Handler) GetUser(ctx *gin.Context) {
	id := ctx.Param("id")

	userVal, ok := auth.GetAuthUser(ctx)
	if !ok {
		return
	}

	res, err := h.srv.GetUser(ctx.Request.Context(), id, userVal.ID)
	if err != nil {
		slog.ErrorContext(ctx.Request.Context(), "failed to get user", slog.Any("error", err), slog.String("id", id))
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "failed to get user"})
		return
	}

	if res == nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "user not found"})
		return
	}

	ctx.JSON(http.StatusOK, res)
}

func (h *Handler) UpdateProfile(ctx *gin.Context) {
	userVal, ok := auth.GetAuthUser(ctx)
	if !ok {
		return
	}

	var body struct {
		FirstName     *string            `json:"firstName" binding:"omitempty,min=1,max=255"`
		LastName      *string            `json:"lastName" binding:"omitempty,min=1,max=255"`
		Bio           *string            `json:"bio"`
		Height        *float64           `json:"height"`
		WeightUnit    *domain.WeightUnit `json:"weightUnit"`
		HeightUnit    *domain.HeightUnit `json:"heightUnit"`
		CurrentWeight *float64           `json:"currentWeight"`
	}

	if err := ctx.ShouldBindJSON(&body); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err := h.srv.UpdateProfile(ctx.Request.Context(), userVal.ID, body.FirstName, body.LastName, body.Bio, body.Height, body.WeightUnit, body.HeightUnit, body.CurrentWeight)
	if err != nil {
		if errors.Is(err, domain.ErrUserNotFound) {
			ctx.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
			return
		}
		slog.ErrorContext(ctx.Request.Context(), "failed to update profile", slog.Any("error", err), slog.String("user_id", userVal.ID))
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "failed to update profile"})
		return
	}

	ctx.Status(http.StatusOK)
}

func (h *Handler) ChangePassword(ctx *gin.Context) {
	userVal, ok := auth.GetAuthUser(ctx)
	if !ok {
		return
	}

	var body struct {
		CurrentPassword string `json:"currentPassword" binding:"required"`
		NewPassword     string `json:"newPassword" binding:"required,min=8"`
	}

	if err := ctx.ShouldBindJSON(&body); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err := h.srv.ChangePassword(ctx.Request.Context(), userVal.ID, body.CurrentPassword, body.NewPassword)
	if err != nil {
		if errors.Is(err, domain.ErrUserNotFound) {
			ctx.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
			return
		}
		if errors.Is(err, domain.ErrNotVerified) {
			ctx.JSON(http.StatusForbidden, gin.H{"error": err.Error()})
			return
		}
		if errors.Is(err, domain.ErrInvalidCredentials) {
			ctx.JSON(http.StatusUnauthorized, gin.H{"error": "senha atual incorreta"})
			return
		}
		slog.ErrorContext(ctx.Request.Context(), "failed to change password", slog.Any("error", err), slog.String("user_id", userVal.ID))
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "failed to change password"})
		return
	}

	ctx.Status(http.StatusOK)
}

func (h *Handler) UploadProfilePicture(ctx *gin.Context) {
	userVal, ok := auth.GetAuthUser(ctx)
	if !ok {
		return
	}

	file, err := ctx.FormFile("file")
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "file is required"})
		return
	}

	openedFile, err := file.Open()
	if err != nil {
		slog.ErrorContext(ctx.Request.Context(), "failed to open file", slog.Any("error", err), slog.String("user_id", userVal.ID))
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "failed to open file"})
		return
	}
	defer openedFile.Close()

	if err := h.srv.UploadProfilePicture(ctx.Request.Context(), userVal.ID, openedFile); err != nil {
		slog.ErrorContext(ctx.Request.Context(), "failed to upload profile picture", slog.Any("error", err), slog.String("user_id", userVal.ID))
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "failed to upload picture"})
		return
	}

	ctx.Status(http.StatusOK)
}

func (h *Handler) RemoveProfilePicture(ctx *gin.Context) {
	userVal, ok := auth.GetAuthUser(ctx)
	if !ok {
		return
	}

	if err := h.srv.RemoveProfilePicture(ctx.Request.Context(), userVal.ID); err != nil {
		slog.ErrorContext(ctx.Request.Context(), "failed to remove profile picture", slog.Any("error", err), slog.String("user_id", userVal.ID))
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "failed to remove profile picture"})
		return
	}

	ctx.Status(http.StatusNoContent)
}

func (h *Handler) ChangeToTrainer(ctx *gin.Context) {
	userVal, ok := auth.GetAuthUser(ctx)
	if !ok {
		return
	}

	var body struct {
		Cref string `json:"cref" binding:"required"`
	}

	if err := ctx.ShouldBindJSON(&body); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := h.srv.ChangeToTrainer(ctx.Request.Context(), userVal.ID, body.Cref); err != nil {
		slog.ErrorContext(ctx.Request.Context(), "failed to upgrade profile", slog.Any("error", err), slog.String("user_id", userVal.ID))
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "failed to upgrade profile"})
		return
	}

	ctx.Status(http.StatusOK)
}

func (h *Handler) ChangeToClient(ctx *gin.Context) {
	userVal, ok := auth.GetAuthUser(ctx)
	if !ok {
		return
	}

	if err := h.srv.ChangeToClient(ctx.Request.Context(), userVal.ID); err != nil {
		slog.ErrorContext(ctx.Request.Context(), "failed to downgrade profile", slog.Any("error", err), slog.String("user_id", userVal.ID))
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "failed to downgrade profile"})
		return
	}

	ctx.Status(http.StatusOK)
}
