package followers

import (
	"log/slog"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/kaua-nasc/gymtrack-go/apps/api/identity/internal/domain"
	"github.com/kaua-nasc/gymtrack-go/libs/auth"
	"github.com/kaua-nasc/gymtrack-go/libs/utils"
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
		protected.GET("/:id/followers", h.ListFollower)
		protected.GET("/:id/following", h.ListFollowing)
		protected.GET("/:id/followers/count", h.CountFollowers)
		protected.GET("/:id/following/count", h.CountFollowing)
		protected.POST("/:id/follows", h.FollowUser)
		protected.POST("/:id/unfollows", h.UnfollowUser)
	}
}

func (h *Handler) ListFollower(ctx *gin.Context) {
	id := ctx.Param("id")
	cursor, limit := h.getPagination(ctx)

	users, nextCursor, err := h.srv.ListFollower(ctx.Request.Context(), id, cursor, limit)
	if err != nil {
		slog.ErrorContext(ctx.Request.Context(), "failed to list followers", slog.Any("error", err), slog.String("id", id))
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "failed to list followers"})
		return
	}

	type followerResponse struct {
		Id        string          `json:"id"`
		FirstName string          `json:"firstName"`
		LastName  string          `json:"lastName"`
		Email     string          `json:"email"`
		Type      domain.UserType `json:"type"`
		CreatedAt time.Time       `json:"createdAt"`
	}

	data := make([]followerResponse, len(users))
	for i, u := range users {
		data[i] = followerResponse{
			Id:        *u.ID,
			FirstName: u.FirstName,
			LastName:  u.LastName,
			Email:     u.Email,
			Type:      u.Type,
			CreatedAt: u.CreatedAt,
		}
	}

	ctx.JSON(http.StatusOK, utils.NewPaginatedResponse(data, nextCursor))
}

func (h *Handler) ListFollowing(ctx *gin.Context) {
	id := ctx.Param("id")
	cursor, limit := h.getPagination(ctx)

	users, nextCursor, err := h.srv.ListFollowing(ctx.Request.Context(), id, cursor, limit)
	if err != nil {
		slog.ErrorContext(ctx.Request.Context(), "failed to list following", slog.Any("error", err), slog.String("id", id))
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "failed to list following"})
		return
	}

	type followerResponse struct {
		Id        string          `json:"id"`
		FirstName string          `json:"firstName"`
		LastName  string          `json:"lastName"`
		Email     string          `json:"email"`
		Type      domain.UserType `json:"type"`
		CreatedAt time.Time       `json:"createdAt"`
	}

	data := make([]followerResponse, len(users))
	for i, u := range users {
		data[i] = followerResponse{
			Id:        *u.ID,
			FirstName: u.FirstName,
			LastName:  u.LastName,
			Email:     u.Email,
			Type:      u.Type,
			CreatedAt: u.CreatedAt,
		}
	}

	ctx.JSON(http.StatusOK, utils.NewPaginatedResponse(data, nextCursor))
}

func (h *Handler) CountFollowers(ctx *gin.Context) {
	id := ctx.Param("id")

	count, err := h.srv.CountFollowers(ctx.Request.Context(), id)
	if err != nil {
		slog.ErrorContext(ctx.Request.Context(), "failed to count followers", slog.Any("error", err), slog.String("id", id))
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "failed to count followers"})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"count": count})
}

func (h *Handler) CountFollowing(ctx *gin.Context) {
	id := ctx.Param("id")

	count, err := h.srv.CountFollowing(ctx.Request.Context(), id)
	if err != nil {
		slog.ErrorContext(ctx.Request.Context(), "failed to count following", slog.Any("error", err), slog.String("id", id))
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "failed to count following"})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"count": count})
}

func (h *Handler) FollowUser(ctx *gin.Context) {
	id := ctx.Param("id")

	user, ok := h.getAuthUser(ctx)
	if !ok {
		return
	}

	if err := h.srv.FollowUser(ctx.Request.Context(), user.ID, id); err != nil {
		slog.ErrorContext(ctx.Request.Context(), "failed to follow user", slog.Any("error", err), slog.String("follower_id", user.ID), slog.String("following_id", id))
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "failed to follow user"})
		return
	}

	ctx.Status(http.StatusOK)
}

func (h *Handler) UnfollowUser(ctx *gin.Context) {
	id := ctx.Param("id")

	user, ok := h.getAuthUser(ctx)
	if !ok {
		return
	}

	if err := h.srv.UnfollowUser(ctx.Request.Context(), user.ID, id); err != nil {
		slog.ErrorContext(ctx.Request.Context(), "failed to unfollow user", slog.Any("error", err), slog.String("follower_id", user.ID), slog.String("following_id", id))
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "failed to unfollow user"})
		return
	}

	ctx.Status(http.StatusOK)
}

func (h *Handler) getAuthUser(ctx *gin.Context) (auth.AuthUser, bool) {
	user, ok := ctx.Value(string(auth.UserContextKey)).(auth.AuthUser)
	if !ok {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "unauthorized"})
		return auth.AuthUser{}, false
	}

	return user, true
}

func (h *Handler) getPagination(ctx *gin.Context) (string, int) {
	cursor := ctx.Query("cursor")
	limit, _ := strconv.Atoi(ctx.DefaultQuery("limit", "20"))
	if limit <= 0 {
		limit = 20
	}
	return cursor, limit
}
