package internal

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/kaua-nasc/gymtrack-go/libs/auth"
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
		social.POST("/posts/:id/like", h.toggleLike)
		social.POST("/posts/:id/comments", h.addComment)
		social.GET("/posts/:id/comments", h.getComments)
	}
}

func (h *PostHandler) createPost(ctx *gin.Context) {
	var post Post
	if err := ctx.ShouldBindJSON(&post); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user, ok := h.getAuthUser(ctx)
	if !ok {
		return
	}

	post.AuthorId = user.ID

	if err := h.service.CreatePost(ctx.Request.Context(), &post, user.ID); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusCreated, post)
}

func (h *PostHandler) getFeed(ctx *gin.Context) {
	user, ok := h.getAuthUser(ctx)
	if !ok {
		return
	}

	posts, err := h.service.GetFeed(ctx.Request.Context(), user.ID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, posts)
}

func (h *PostHandler) toggleLike(ctx *gin.Context) {
	postId := ctx.Param("id")
	user, ok := h.getAuthUser(ctx)
	if !ok {
		return
	}

	if err := h.service.ToggleLike(ctx.Request.Context(), postId, user.ID); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
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
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	comment.PostId = postId
	comment.AuthorId = user.ID

	if err := h.service.AddComment(ctx.Request.Context(), &comment, user.ID); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusCreated, comment)
}

func (h *PostHandler) getComments(ctx *gin.Context) {
	postId := ctx.Param("id")

	comments, err := h.service.GetComments(ctx.Request.Context(), postId)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, comments)
}

func (h *PostHandler) getAuthUser(ctx *gin.Context) (auth.AuthUser, bool) {
	user, ok := ctx.Value(string(auth.UserContextKey)).(auth.AuthUser)
	if !ok {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "unauthorized"})
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
