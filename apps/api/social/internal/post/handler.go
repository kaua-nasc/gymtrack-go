package post

import (
	"io"
	"log/slog"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/kaua-nasc/gymtrack-go/apps/api/social/internal/domain"
	"github.com/kaua-nasc/gymtrack-go/libs/auth"
	"github.com/kaua-nasc/gymtrack-go/libs/log"
	"github.com/kaua-nasc/gymtrack-go/libs/utils"
)

type Handler struct {
	service *Service
}

func NewHandler(service *Service) *Handler {
	return &Handler{service: service}
}

func (h *Handler) RegisterRoutes(r *gin.Engine) {
	r.Use(log.LoggerMiddleware())
	r.GET("/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"status": "ok", "service": "identity"})
	})
	r.GET("/health2", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"status": "ok", "service": "identity"})
	})
	social := r.Group("/social")
	social.Use(auth.AuthMiddleware())
	{
		app := social.Group("")
		app.Use(auth.RolesMiddleware(auth.Client, auth.Trainer))
		{
			app.POST("/posts", h.createPost)
			app.GET("/posts", h.getFeed)
			app.GET("/posts/author/:authorId", h.getPostsByAuthor)
			app.POST("/posts/media", h.uploadMedia)
			app.PUT("/posts/:id", h.updatePost)
			app.DELETE("/posts/:id", h.deletePost)
		}

		admin := social.Group("/admin")
		admin.Use(auth.RolesMiddleware(auth.Admin))
		{
			admin.GET("/posts/pending", h.getPendingPosts)
			admin.GET("/posts/history", h.getAuditHistory)
			admin.PATCH("/posts/:id/status", h.updatePostStatus)
		}
	}
}

func (h *Handler) createPost(ctx *gin.Context) {
	var post domain.Post
	if err := ctx.ShouldBindJSON(&post); err != nil {
		ctx.JSON(http.StatusBadRequest, utils.NewErrorResponse(err.Error()))
		return
	}

	user, ok := auth.GetAuthUser(ctx)
	if !ok {
		return
	}

	if err := h.service.CreatePost(ctx, &post, user.ID); err != nil {
		slog.ErrorContext(ctx.Request.Context(), "failed to create post", slog.Any("error", err), slog.String("authorId", user.ID))
		ctx.JSON(http.StatusInternalServerError, utils.NewErrorResponse(err.Error()))
		return
	}

	ctx.Status(http.StatusCreated)
}

func (h *Handler) uploadMedia(ctx *gin.Context) {
	user, ok := auth.GetAuthUser(ctx)
	if !ok {
		return
	}

	form, err := ctx.MultipartForm()
	if err != nil {
		ctx.JSON(http.StatusBadRequest, utils.NewErrorResponse("failed to parse multipart form"))
		return
	}

	files := form.File["files"]
	if len(files) == 0 {
		ctx.JSON(http.StatusBadRequest, utils.NewErrorResponse("no files provided"))
		return
	}

	var readers []io.Reader
	var filenames []string

	for _, fileHeader := range files {
		f, err := fileHeader.Open()
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, utils.NewErrorResponse("failed to open file"))
			return
		}
		defer f.Close()
		readers = append(readers, f)
		filenames = append(filenames, fileHeader.Filename)
	}

	urls, err := h.service.UploadMedia(ctx.Request.Context(), user.ID, readers, filenames)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, utils.NewErrorResponse(err.Error()))
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"mediaUrls": urls})
}

func (h *Handler) updatePost(ctx *gin.Context) {
	postId := ctx.Param("id")
	user, ok := auth.GetAuthUser(ctx)
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

func (h *Handler) deletePost(ctx *gin.Context) {
	postId := ctx.Param("id")
	user, ok := auth.GetAuthUser(ctx)
	if !ok {
		return
	}

	if err := h.service.DeletePost(ctx.Request.Context(), postId, user.ID); err != nil {
		ctx.JSON(http.StatusInternalServerError, utils.NewErrorResponse(err.Error()))
		return
	}

	ctx.Status(http.StatusNoContent)
}

func (h *Handler) getFeed(ctx *gin.Context) {
	user, ok := auth.GetAuthUser(ctx)
	if !ok {
		return
	}

	cursor, limit := utils.GetPagination(ctx)
	posts, nextCursor, err := h.service.GetFeed(ctx.Request.Context(), user.ID, cursor, limit)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, utils.NewErrorResponse(err.Error()))
		return
	}

	ctx.JSON(http.StatusOK, utils.NewPaginatedResponse(posts, nextCursor))
}

func (h *Handler) getPendingPosts(ctx *gin.Context) {
	user, ok := auth.GetAuthUser(ctx)
	if !ok {
		return
	}

	cursor, limit := utils.GetPagination(ctx)
	posts, nextCursor, err := h.service.GetPendingPosts(ctx.Request.Context(), user.ID, cursor, limit)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, utils.NewErrorResponse(err.Error()))
		return
	}

	ctx.JSON(http.StatusOK, utils.NewPaginatedResponse(posts, nextCursor))
}

func (h *Handler) getAuditHistory(ctx *gin.Context) {
	user, ok := auth.GetAuthUser(ctx)
	if !ok {
		return
	}

	cursor, limit := utils.GetPagination(ctx)
	status := ctx.Query("status")
	startDate := ctx.Query("startDate")
	endDate := ctx.Query("endDate")

	posts, nextCursor, err := h.service.GetAuditHistory(ctx.Request.Context(), user.ID, status, startDate, endDate, cursor, limit)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, utils.NewErrorResponse(err.Error()))
		return
	}

	ctx.JSON(http.StatusOK, utils.NewPaginatedResponse(posts, nextCursor))
}

func (h *Handler) updatePostStatus(ctx *gin.Context) {
	postId := ctx.Param("id")
	user, ok := auth.GetAuthUser(ctx)
	if !ok {
		return
	}

	var body struct {
		Status domain.PostStatus `json:"status" binding:"required"`
		Reason *string           `json:"reason" binding:"omitempty,max=1024"`
	}
	if err := ctx.ShouldBindJSON(&body); err != nil {
		ctx.JSON(http.StatusBadRequest, utils.NewErrorResponse(err.Error()))
		return
	}

	if body.Status != domain.PostApproved && body.Status != domain.PostRejected {
		ctx.JSON(http.StatusBadRequest, utils.NewErrorResponse("invalid status. must be APPROVED or REJECTED"))
		return
	}

	if body.Status == domain.PostRejected && (body.Reason == nil || *body.Reason == "") {
		ctx.JSON(http.StatusBadRequest, utils.NewErrorResponse("reason is required when rejecting a post"))
		return
	}

	if err := h.service.UpdatePostStatus(ctx.Request.Context(), user.ID, postId, body.Status, body.Reason); err != nil {
		ctx.JSON(http.StatusInternalServerError, utils.NewErrorResponse(err.Error()))
		return
	}

	ctx.Status(http.StatusNoContent)
}

func (h *Handler) getPostsByAuthor(ctx *gin.Context) {
	authorId := ctx.Param("authorId")
	user, ok := auth.GetAuthUser(ctx)
	if !ok {
		return
	}

	cursor, limit := utils.GetPagination(ctx)
	posts, nextCursor, err := h.service.GetPostsByAuthor(ctx.Request.Context(), authorId, user.ID, cursor, limit)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, utils.NewErrorResponse(err.Error()))
		return
	}

	ctx.JSON(http.StatusOK, utils.NewPaginatedResponse(posts, nextCursor))
}
