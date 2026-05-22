package exerciselog

import (
	"log/slog"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/kaua-nasc/gymtrack-go/libs/auth"
	"github.com/kaua-nasc/gymtrack-go/libs/log"
	"github.com/kaua-nasc/gymtrack-go/libs/utils"
)

type Handler struct {
	srv *Service
}

func NewHandler(srv *Service) *Handler {
	return &Handler{
		srv: srv,
	}
}

func (h *Handler) RegisterRoutes(r *gin.Engine) {
	r.Use(log.LoggerMiddleware())
	plans := r.Group("/training-plans")
	plans.Use(auth.AuthMiddleware())
	{
		plans.POST("/:id/days/:dayId/exercises/:exerciseId/logs", h.LogExercise)
	}
}

func (h *Handler) LogExercise(ctx *gin.Context) {
	id := ctx.Param("exerciseId")
	user, ok := h.getAuthUser(ctx)
	if !ok {
		return
	}

	var body struct {
		Reps   []int     `json:"reps" binding:"required"`
		Weight []float64 `json:"weight" binding:"required"`
		Notes  *string   `json:"notes"`
	}

	if err := ctx.ShouldBindJSON(&body); err != nil {
		ctx.JSON(http.StatusBadRequest, utils.NewErrorResponse(err.Error()))
		return
	}

	if err := h.srv.LogExercise(ctx.Request.Context(), id, user.ID, body.Reps, body.Weight, body.Notes); err != nil {
		slog.ErrorContext(ctx.Request.Context(), "failed to log exercise", slog.Any("error", err), slog.String("exercise_id", id), slog.String("user_id", user.ID))
		ctx.JSON(http.StatusInternalServerError, utils.NewErrorResponse("failed to log exercise"))
		return
	}

	ctx.Status(http.StatusNoContent)
}

func (h *Handler) getAuthUser(ctx *gin.Context) (auth.AuthUser, bool) {
	user, ok := ctx.Value(string(auth.UserContextKey)).(auth.AuthUser)
	if !ok {
		ctx.JSON(http.StatusUnauthorized, utils.NewErrorResponse("unauthorized"))
		return auth.AuthUser{}, false
	}

	return user, true
}
