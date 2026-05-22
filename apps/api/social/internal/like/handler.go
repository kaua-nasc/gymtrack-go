package like

import (
	"net/http"

	"github.com/gin-gonic/gin"
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
		social.POST("/posts/:id/like", h.toggleLike)
	}
}

func (h *Handler) toggleLike(ctx *gin.Context) {
	postId := ctx.Param("id")
	user, ok := auth.GetAuthUser(ctx)
	if !ok {
		return
	}

	if err := h.service.ToggleLike(ctx.Request.Context(), postId, user.ID); err != nil {
		ctx.JSON(http.StatusInternalServerError, utils.NewErrorResponse(err.Error()))
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "success"})
}
