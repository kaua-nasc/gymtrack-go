package plan

import (
	"errors"
	"io"
	"log/slog"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/kaua-nasc/gymtrack-go/apps/api/training-plan/internal/domain"
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
		plans.GET("", h.ListPlan)
		plans.POST("/by-ids", h.ListPlansByIds)
		plans.POST("", h.CreatePlan)
		plans.GET("author/:authorId", h.ListPlan)
		plans.GET("/:id", h.GetPlan)
		plans.PUT("/:id", h.UpdatePlan)
		plans.PUT("/:id/max-subscriptions", h.UpdateMaxSubscriptions)
		plans.DELETE("/:id", h.DeletePlan)
		plans.GET("/exists/:id", h.ExistsPlan)

		plans.POST("/:id/days", h.CreateDay)
		plans.DELETE("/:id/days/:dayId", h.DeleteDay)

		plans.POST("/:id/days/:dayId/exercises", h.CreateExercise)
		plans.POST("/:id/days/:dayId/exercises/:exerciseId/media", h.UploadExerciseMedia)
		plans.DELETE("/:id/days/:dayId/exercises/:exerciseId", h.DeleteExercise)
	}
}

func (h *Handler) CreatePlan(ctx *gin.Context) {
	var plan domain.TrainingPlan
	if err := ctx.ShouldBindJSON(&plan); err != nil {
		ctx.JSON(http.StatusBadRequest, utils.NewErrorResponse(err.Error()))
		return
	}

	user, ok := auth.GetAuthUser(ctx)
	if !ok {
		return
	}

	_, err := h.srv.CreatePlan(ctx.Request.Context(), plan, user)
	if err != nil {
		slog.ErrorContext(ctx.Request.Context(), "failed to create plan", slog.Any("error", err), slog.String("authorId", user.ID))
		ctx.JSON(http.StatusInternalServerError, utils.NewErrorResponse(err.Error()))
		return
	}

	ctx.Status(http.StatusCreated)
}

func (h *Handler) ListPlan(ctx *gin.Context) {
	authorId := ctx.Param("authorId")
	cursor, limit := utils.GetPagination(ctx)

	plans, nextCursor, err := h.srv.ListPlan(ctx.Request.Context(), authorId, cursor, limit)
	if err != nil {
		slog.ErrorContext(ctx.Request.Context(), "failed to list plans", slog.Any("error", err), slog.String("authorId", authorId))
		ctx.JSON(http.StatusInternalServerError, utils.NewErrorResponse(err.Error()))
		return
	}

	ctx.JSON(http.StatusOK, utils.NewPaginatedResponse(plans, nextCursor))
}

func (h *Handler) UpdatePlan(ctx *gin.Context) {
	id := ctx.Param("id")
	if id == "" {
		ctx.JSON(http.StatusBadRequest, utils.NewErrorResponse("plan id is required"))
		return
	}

	var plan domain.TrainingPlan
	if err := ctx.ShouldBindJSON(&plan); err != nil {
		ctx.JSON(http.StatusBadRequest, utils.NewErrorResponse(err.Error()))
		return
	}

	updatedPlan, err := h.srv.UpdatePlan(ctx.Request.Context(), id, plan)
	if err != nil {
		slog.ErrorContext(ctx.Request.Context(), "failed to update plan", slog.Any("error", err), slog.String("plan_id", id))
		ctx.JSON(http.StatusInternalServerError, utils.NewErrorResponse(err.Error()))
		return
	}

	ctx.JSON(http.StatusOK, updatedPlan)
}

func (h *Handler) UpdateMaxSubscriptions(ctx *gin.Context) {
	id := ctx.Param("id")
	if id == "" {
		ctx.JSON(http.StatusBadRequest, utils.NewErrorResponse("plan id is required"))
		return
	}

	var payload struct {
		MaxSubscriptions int `json:"maxSubscriptions" binding:"required,min=0"`
	}
	if err := ctx.ShouldBindJSON(&payload); err != nil {
		ctx.JSON(http.StatusBadRequest, utils.NewErrorResponse(err.Error()))
		return
	}

	err := h.srv.UpdateMaxSubscriptions(ctx.Request.Context(), id, payload.MaxSubscriptions)
	switch {
	case errors.Is(err, domain.ErrPlanNotFound):
		ctx.JSON(http.StatusNotFound, utils.NewErrorResponse(err.Error()))
		return
	case errors.Is(err, domain.ErrMaxSubscriptionsBelowCurrent):
		ctx.JSON(http.StatusBadRequest, utils.NewErrorResponse(err.Error()))
		return
	case errors.Is(err, domain.ErrPlanAccessForbidden):
		ctx.JSON(http.StatusForbidden, utils.NewErrorResponse(err.Error()))
		return
	case err != nil:
		slog.ErrorContext(ctx.Request.Context(), "failed to update plan max subscriptions", slog.Any("error", err), slog.String("plan_id", id))
		ctx.JSON(http.StatusInternalServerError, utils.NewErrorResponse("failed to update plan max subscriptions"))
		return
	}

	ctx.Status(http.StatusOK)
}

func (h *Handler) DeletePlan(ctx *gin.Context) {
	id := ctx.Param("id")

	if err := h.srv.DeletePlan(ctx.Request.Context(), id); err != nil {
		slog.ErrorContext(ctx.Request.Context(), "failed to delete plan", slog.Any("error", err), slog.String("plan_id", id))
		ctx.JSON(http.StatusInternalServerError, utils.NewErrorResponse(err.Error()))
		return
	}

	ctx.Status(http.StatusOK)
}

func (h *Handler) ExistsPlan(ctx *gin.Context) {
	id := ctx.Param("id")
	if id == "" {
		ctx.JSON(http.StatusBadRequest, utils.NewErrorResponse("plan id is required"))
		return
	}

	publicOnly := ctx.Query("publicOnly") == "true"

	exists, err := h.srv.ExistsPlan(ctx.Request.Context(), id, publicOnly)
	if err != nil {
		slog.ErrorContext(ctx.Request.Context(), "failed to check plan existence", slog.Any("error", err), slog.String("plan_id", id))
		ctx.JSON(http.StatusInternalServerError, utils.NewErrorResponse(err.Error()))
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"exists": exists})
}

func (h *Handler) GetPlan(ctx *gin.Context) {
	id := ctx.Param("id")
	if id == "" {
		ctx.JSON(http.StatusBadRequest, utils.NewErrorResponse("plan id is required"))
		return
	}

	plan, err := h.srv.GetPlan(ctx.Request.Context(), id)
	if err != nil {
		slog.ErrorContext(ctx.Request.Context(), "failed to get plan", slog.Any("error", err), slog.String("plan_id", id))
		ctx.JSON(http.StatusInternalServerError, utils.NewErrorResponse(err.Error()))
		return
	}

	if plan == nil {
		ctx.JSON(http.StatusNotFound, utils.NewErrorResponse("training plan not found"))
		return
	}

	ctx.JSON(http.StatusOK, plan)
}

func (h *Handler) CreateDay(ctx *gin.Context) {
	var day domain.Day
	if err := ctx.ShouldBindJSON(&day); err != nil {
		ctx.JSON(http.StatusBadRequest, utils.NewErrorResponse(err.Error()))
		return
	}

	if err := h.srv.CreateDay(ctx.Request.Context(), &day); err != nil {
		slog.ErrorContext(ctx.Request.Context(), "failed to create day", slog.Any("error", err))
		ctx.JSON(http.StatusInternalServerError, utils.NewErrorResponse(err.Error()))
		return
	}

	ctx.Status(http.StatusCreated)
}

func (h *Handler) CreateExercise(ctx *gin.Context) {
	var exercise domain.Exercise
	if err := ctx.ShouldBindJSON(&exercise); err != nil {
		ctx.JSON(http.StatusBadRequest, utils.NewErrorResponse(err.Error()))
		return
	}

	if err := h.srv.CreateExercise(ctx.Request.Context(), &exercise); err != nil {
		slog.ErrorContext(ctx.Request.Context(), "failed to create exercise", slog.Any("error", err))
		ctx.JSON(http.StatusInternalServerError, utils.NewErrorResponse(err.Error()))
		return
	}

	ctx.JSON(http.StatusCreated, exercise)
}

func (h *Handler) UploadExerciseMedia(ctx *gin.Context) {
	exerciseId := ctx.Param("exerciseId")
	if exerciseId == "" {
		ctx.JSON(http.StatusBadRequest, utils.NewErrorResponse("exercise id is required"))
		return
	}

	var videoFile *domain.UploadFile
	var imageFile *domain.UploadFile

	// Try to read video file
	vFile, err := ctx.FormFile("video")
	if err == nil && vFile != nil {
		if vFile.Size > 20*1024*1024 {
			ctx.JSON(http.StatusBadRequest, utils.NewErrorResponse("video file size exceeds 20MB limit (approx. 1 minute of video)"))
			return
		}
		f, err := vFile.Open()
		if err == nil {
			defer f.Close()
			data, err := io.ReadAll(f)
			if err == nil {
				videoFile = &domain.UploadFile{
					Data:     data,
					Filename: vFile.Filename,
				}
			}
		}
	}

	// Try to read image file
	iFile, err := ctx.FormFile("image")
	if err == nil && iFile != nil {
		if iFile.Size > 5*1024*1024 {
			ctx.JSON(http.StatusBadRequest, utils.NewErrorResponse("image file size exceeds 5MB limit"))
			return
		}
		f, err := iFile.Open()
		if err == nil {
			defer f.Close()
			data, err := io.ReadAll(f)
			if err == nil {
				imageFile = &domain.UploadFile{
					Data:     data,
					Filename: iFile.Filename,
				}
			}
		}
	}

	if videoFile == nil && imageFile == nil {
		ctx.JSON(http.StatusBadRequest, utils.NewErrorResponse("no media files (video or image) provided"))
		return
	}

	_, _, err = h.srv.UploadExerciseMedia(ctx.Request.Context(), exerciseId, videoFile, imageFile)
	if err != nil {
		slog.ErrorContext(ctx.Request.Context(), "failed to upload exercise media", slog.Any("error", err), slog.String("exerciseId", exerciseId))
		ctx.JSON(http.StatusInternalServerError, utils.NewErrorResponse(err.Error()))
		return
	}

	ctx.Status(http.StatusOK)
}

func (h *Handler) DeleteDay(ctx *gin.Context) {
	dayId := ctx.Param("dayId")

	if err := h.srv.DeleteDay(ctx.Request.Context(), dayId); err != nil {
		slog.ErrorContext(ctx.Request.Context(), "failed to delete day", slog.Any("error", err), slog.String("day_id", dayId))
		ctx.JSON(http.StatusInternalServerError, utils.NewErrorResponse(err.Error()))
		return
	}

	ctx.Status(http.StatusOK)
}

func (h *Handler) DeleteExercise(ctx *gin.Context) {
	id := ctx.Param("exerciseId")

	if err := h.srv.DeleteExercise(ctx.Request.Context(), id); err != nil {
		slog.ErrorContext(ctx.Request.Context(), "failed to delete exercise", slog.Any("error", err), slog.String("exercise_id", id))
		ctx.JSON(http.StatusInternalServerError, utils.NewErrorResponse(err.Error()))
		return
	}

	ctx.Status(http.StatusOK)
}

func (h *Handler) ListPlansByIds(ctx *gin.Context) {
	var ids []string
	if err := ctx.ShouldBindJSON(&ids); err != nil {
		ctx.JSON(http.StatusBadRequest, utils.NewErrorResponse(err.Error()))
		return
	}

	res, err := h.srv.ListPlansByIds(ctx.Request.Context(), ids)
	if err != nil {
		slog.ErrorContext(ctx.Request.Context(), "failed to list training plans", slog.Any("error", err))
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "failed to list training plans"})
		return
	}

	ctx.JSON(http.StatusOK, res)
}
