package comment

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/kaua-nasc/gymtrack-go/apps/api/social/internal/domain"
	"github.com/kaua-nasc/gymtrack-go/libs/auth"
	"github.com/kaua-nasc/gymtrack-go/libs/utils"
)

type Handler struct {
	service *Service
}

func NewHandler(service *Service) *Handler {
	return &Handler{service: service}
}

func (h *Handler) RegisterRoutes(r *gin.Engine) {
	social := r.Group("/social")
	social.Use(auth.AuthMiddleware())
	{
		social.POST("/posts/:id/comments", h.addComment)
		social.GET("/posts/:id/comments", h.getComments)
		social.DELETE("/comments/:commentId", h.deleteComment)
	}
}

func (h *Handler) addComment(ctx *gin.Context) {
	postId := ctx.Param("id")
	user, ok := h.getAuthUser(ctx)
	if !ok {
		return
	}

	var comment domain.Comment
	if err := ctx.ShouldBindJSON(&comment); err != nil {
		ctx.JSON(http.StatusBadRequest, utils.NewErrorResponse(err.Error()))
		return
	}

	comment.PostId = postId
	comment.AuthorId = user.ID

	if err := h.service.AddComment(ctx.Request.Context(), &comment, user.ID); err != nil {
		ctx.JSON(http.StatusInternalServerError, utils.NewErrorResponse(err.Error()))
		return
	}

	ctx.Status(http.StatusCreated)
}

func (h *Handler) deleteComment(ctx *gin.Context) {
	commentId := ctx.Param("commentId")
	user, ok := h.getAuthUser(ctx)
	if !ok {
		return
	}

	if err := h.service.DeleteComment(ctx.Request.Context(), commentId, user.ID); err != nil {
		ctx.JSON(http.StatusInternalServerError, utils.NewErrorResponse(err.Error()))
		return
	}

	ctx.Status(http.StatusNoContent)
}

func (h *Handler) getComments(ctx *gin.Context) {
	postId := ctx.Param("id")
	cursor, limit := h.getPagination(ctx)

	comments, nextCursor, err := h.service.GetComments(ctx.Request.Context(), postId, cursor, limit)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, utils.NewErrorResponse(err.Error()))
		return
	}

	ctx.JSON(http.StatusOK, utils.NewPaginatedResponse(comments, nextCursor))
}

func (h *Handler) getAuthUser(ctx *gin.Context) (auth.AuthUser, bool) {
	user, ok := ctx.Value(string(auth.UserContextKey)).(auth.AuthUser)
	if !ok {
		ctx.JSON(http.StatusUnauthorized, utils.NewErrorResponse("unauthorized"))
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
