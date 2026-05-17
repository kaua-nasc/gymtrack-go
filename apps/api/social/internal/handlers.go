package internal

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/kaua-nasc/gymtrack-go/libs/auth"
	"github.com/kaua-nasc/gymtrack-go/libs/utils"
)

type PostHandler struct {
	service *PostService
}

func NewPostHandler(service *PostService) *PostHandler {
	return &PostHandler{service: service}
}

func (h *PostHandler) RegisterRoutes(r *gin.Engine) {
	social := r.Group("/social")
	social.Use(auth.AuthMiddleware())
	{
		social.POST("/posts", h.createPost)
		social.GET("/posts", h.getFeed)
		social.PUT("/posts/:id", h.updatePost)
		social.DELETE("/posts/:id", h.deletePost)
		social.POST("/posts/:id/like", h.toggleLike)
		social.POST("/posts/:id/comments", h.addComment)
		social.GET("/posts/:id/comments", h.getComments)
		social.DELETE("/comments/:commentId", h.deleteComment)
	}
}

func (h *PostHandler) createPost(ctx *gin.Context) {
	var post Post
	if err := ctx.ShouldBindJSON(&post); err != nil {
		ctx.JSON(http.StatusBadRequest, utils.NewErrorResponse(err.Error()))
		return
	}

	user, ok := h.getAuthUser(ctx)
	if !ok {
		return
	}

	if err := h.service.CreatePost(ctx.Request.Context(), &post, user.ID); err != nil {
		ctx.JSON(http.StatusInternalServerError, utils.NewErrorResponse(err.Error()))
		return
	}

	ctx.Status(http.StatusCreated)
}

func (h *PostHandler) updatePost(ctx *gin.Context) {
	postId := ctx.Param("id")
	user, ok := h.getAuthUser(ctx)
	if !ok {
		return
	}

	var body struct {
		Content string `json:"content" binding:"required"`
	}
	if err := ctx.ShouldBindJSON(&body); err != nil {
		ctx.JSON(http.StatusBadRequest, utils.NewErrorResponse(err.Error()))
		return
	}

	if err := h.service.UpdatePost(ctx.Request.Context(), postId, user.ID, body.Content); err != nil {
		ctx.JSON(http.StatusInternalServerError, utils.NewErrorResponse(err.Error()))
		return
	}

	ctx.Status(http.StatusNoContent)
}

func (h *PostHandler) deletePost(ctx *gin.Context) {
	postId := ctx.Param("id")
	user, ok := h.getAuthUser(ctx)
	if !ok {
		return
	}

	if err := h.service.DeletePost(ctx.Request.Context(), postId, user.ID); err != nil {
		ctx.JSON(http.StatusInternalServerError, utils.NewErrorResponse(err.Error()))
		return
	}

	ctx.Status(http.StatusNoContent)
}

func (h *PostHandler) getFeed(ctx *gin.Context) {
	user, ok := h.getAuthUser(ctx)
	if !ok {
		return
	}

	cursor, limit := h.getPagination(ctx)
	posts, nextCursor, err := h.service.GetFeed(ctx.Request.Context(), user.ID, cursor, limit)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, utils.NewErrorResponse(err.Error()))
		return
	}

	ctx.JSON(http.StatusOK, utils.NewPaginatedResponse(posts, nextCursor))
}

func (h *PostHandler) toggleLike(ctx *gin.Context) {
	postId := ctx.Param("id")
	user, ok := h.getAuthUser(ctx)
	if !ok {
		return
	}

	if err := h.service.ToggleLike(ctx.Request.Context(), postId, user.ID); err != nil {
		ctx.JSON(http.StatusInternalServerError, utils.NewErrorResponse(err.Error()))
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "success"})
}

func (h *PostHandler) addComment(ctx *gin.Context) {
	postId := ctx.Param("id")
	user, ok := h.getAuthUser(ctx)
	if !ok {
		return
	}

	var comment Comment
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

func (h *PostHandler) deleteComment(ctx *gin.Context) {
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

func (h *PostHandler) getComments(ctx *gin.Context) {
	postId := ctx.Param("id")
	cursor, limit := h.getPagination(ctx)

	comments, nextCursor, err := h.service.GetComments(ctx.Request.Context(), postId, cursor, limit)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, utils.NewErrorResponse(err.Error()))
		return
	}

	ctx.JSON(http.StatusOK, utils.NewPaginatedResponse(comments, nextCursor))
}

func (h *PostHandler) getAuthUser(ctx *gin.Context) (auth.AuthUser, bool) {
	user, ok := ctx.Value(string(auth.UserContextKey)).(auth.AuthUser)
	if !ok {
		ctx.JSON(http.StatusUnauthorized, utils.NewErrorResponse("unauthorized"))
		return auth.AuthUser{}, false
	}

	return user, true
}

func (h *PostHandler) getPagination(ctx *gin.Context) (string, int) {
	cursor := ctx.Query("cursor")
	limit, _ := strconv.Atoi(ctx.DefaultQuery("limit", "20"))
	if limit <= 0 {
		limit = 20
	}
	return cursor, limit
}
