package comment

import (
	"net/http"

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
	social.Use(auth.AuthMiddleware(), auth.RolesMiddleware(auth.Client, auth.Trainer))
	{
		social.POST("/posts/:id/comments", h.addComment)
		social.GET("/posts/:id/comments", h.getComments)
		social.DELETE("/comments/:commentId", h.deleteComment)
	}
}

func (h *Handler) addComment(ctx *gin.Context) {
	postId := ctx.Param("id")
	user, ok := auth.GetAuthUser(ctx)
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
	user, ok := auth.GetAuthUser(ctx)
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
	cursor, limit := utils.GetPagination(ctx)

	comments, nextCursor, err := h.service.GetComments(ctx.Request.Context(), postId, cursor, limit)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, utils.NewErrorResponse(err.Error()))
		return
	}

	ctx.JSON(http.StatusOK, utils.NewPaginatedResponse(comments, nextCursor))
}
