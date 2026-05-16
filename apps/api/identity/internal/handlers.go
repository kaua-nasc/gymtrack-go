package internal

import (
	"errors"
	"log/slog"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/kaua-nasc/gymtrack-go/libs/auth"
	"github.com/kaua-nasc/gymtrack-go/libs/log"
)

type UserHandler struct {
	srv *UserService
}

func NewUserHandler(srv *UserService) *UserHandler {
	return &UserHandler{srv: srv}
}

func (h *UserHandler) RegisterRoutes(r *gin.Engine) {
	r.Use(log.LoggerMiddleware())
	r.GET("/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"status": "ok", "service": "identity"})
	})
	r.POST("/identity/auth/register", h.Register)
	r.POST("/identity/auth/login", h.Login)
	r.POST("/identity/auth/verify/send-token", h.SendVerificationEmail)
	r.POST("/identity/auth/verify", h.VerifyEmail)
	r.POST("/identity/auth/reset-password/send-token", h.ResetPasswordSendToken)
	r.POST("/identity/auth/reset-password/verify-token", h.ResetPasswordVerifyToken)
	r.POST("/identity/auth/reset-password", h.ResetPassword)

	// Protected routes
	protected := r.Group("/identity/users")
	protected.Use(auth.AuthMiddleware())
	{
		protected.GET("", h.ListUsers)
		protected.GET("/:id", h.GetUser)
		protected.PUT("/:id", h.UpdateProfile)
		protected.PUT("/profile/password", h.ChangePassword)

		protected.GET("/:id/followers", h.ListFollower)
		protected.GET("/:id/following", h.ListFollowing)
		protected.GET("/:id/followers/count", h.CountFollowers)
		protected.GET("/:id/following/count", h.CountFollowing)
		protected.POST("/:id/follows", h.FollowUser)
		protected.POST("/:id/unfollows", h.UnfollowUser)

		protected.PUT("/profile/picture", h.UploadProfilePicture)
		protected.DELETE("/profile/picture", h.RemoveProfilePicture)

		protected.POST("/profile/upgrade", h.ChangeToTrainer)
		protected.POST("/profile/downgrade", h.ChangeToClient)

		// Trainer routes
		trainers := protected.Group("/trainers")
		// trainers.Use(auth.RolesMiddleware(string(Trainer)))
		{
			trainers.PATCH("/profile/code", h.CreateTrainerCode)
			trainers.POST("/profile/link", h.LinkTrainer)
			trainers.POST("/profile/unlink", h.UnlinkTrainer)

			trainers.PATCH("/body-measurements/:id/notes", h.AddBodyMeasurementNote)
			trainers.GET("/body-measurements/latest", h.FindLastBodyMeasurementNote)
			trainers.GET("/body-measurements", h.ListBodyMeasurements)
			trainers.GET("/students/:id/body-measurements", h.ListBodyMeasurements)

			trainers.PATCH("/weight-log/:id/notes", h.AddWeightLogNote)
			trainers.GET("/weight-logs", h.ListWeightLogs)
			trainers.GET("/students/:id/weight-logs", h.ListWeightLogs)

			trainers.POST("/goals", h.AddGoalMetric)
			trainers.GET("/goals", h.ListGoalsMetric)
			trainers.GET("/students/:id/goals", h.ListGoalsMetricById)

			trainers.GET("/students", h.ListStudents)
			trainers.POST("/students/:id/profile/unlink", h.UnlinkStudent)
		}
	}
}

func (h *UserHandler) SendVerificationEmail(ctx *gin.Context) {
	var body struct {
		Email string `json:"email" binding:"required,email"`
	}

	if err := ctx.ShouldBindJSON(&body); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := h.srv.SendVerificationEmail(ctx.Request.Context(), body.Email); err != nil {
		if errors.Is(err, ErrEmailNotFound) {
			ctx.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
			return
		}
		if errors.Is(err, ErrAlreadyVerified) {
			ctx.JSON(http.StatusConflict, gin.H{"error": err.Error()})
			return
		}
		slog.ErrorContext(ctx.Request.Context(), "failed to send verification token", slog.Any("error", err), slog.String("email", body.Email))
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "failed to send verification token"})
		return
	}

	ctx.Status(http.StatusOK)
}

func (h *UserHandler) VerifyEmail(ctx *gin.Context) {
	var body struct {
		Email string `json:"email" binding:"required,email"`
		Code  string `json:"code" binding:"required"`
	}

	if err := ctx.ShouldBindJSON(&body); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := h.srv.VerifyEmail(ctx.Request.Context(), body.Email, body.Code); err != nil {
		if errors.Is(err, ErrInvalidCode) {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		if errors.Is(err, ErrUserNotFound) {
			ctx.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
			return
		}
		if errors.Is(err, ErrAlreadyVerified) {
			ctx.JSON(http.StatusConflict, gin.H{"error": err.Error()})
			return
		}
		slog.ErrorContext(ctx.Request.Context(), "failed to verify email", slog.Any("error", err), slog.String("email", body.Email))
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "failed to verify email"})
		return
	}

	ctx.Status(http.StatusOK)
}

func (h *UserHandler) Login(ctx *gin.Context) {
	var body struct {
		Email    string `json:"email" binding:"required,email"`
		Password string `json:"password" binding:"required"`
	}

	if err := ctx.ShouldBindJSON(&body); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	token, err := h.srv.Login(ctx.Request.Context(), body.Email, body.Password)
	if err != nil {
		if errors.Is(err, ErrInvalidCredentials) {
			ctx.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
			return
		}
		slog.ErrorContext(ctx.Request.Context(), "login failed", slog.Any("error", err), slog.String("email", body.Email))
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "failed to login"})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"accessToken": token})
}

func (h *UserHandler) ResetPasswordSendToken(ctx *gin.Context) {
	var body struct {
		Email string `json:"email" binding:"required,email"`
	}

	if err := ctx.ShouldBindJSON(&body); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := h.srv.ResetPasswordSendToken(ctx.Request.Context(), body.Email); err != nil {
		if errors.Is(err, ErrEmailNotFound) {
			ctx.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
			return
		}
		slog.ErrorContext(ctx.Request.Context(), "failed to send reset token", slog.Any("error", err), slog.String("email", body.Email))
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "failed to send reset token"})
		return
	}

	ctx.Status(http.StatusOK)
}

func (h *UserHandler) ResetPasswordVerifyToken(ctx *gin.Context) {
	var body struct {
		Email string `json:"email" binding:"required,email"`
		Code  string `json:"code" binding:"required"`
	}

	if err := ctx.ShouldBindJSON(&body); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	res, err := h.srv.ResetPasswordVerifyToken(ctx.Request.Context(), body.Email, body.Code)
	if err != nil {
		if errors.Is(err, ErrInvalidCode) {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		slog.ErrorContext(ctx.Request.Context(), "failed to verify reset token", slog.Any("error", err), slog.String("email", body.Email))
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "failed to verify reset token"})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"valid": res})
}

func (h *UserHandler) ResetPassword(ctx *gin.Context) {
	var body struct {
		Email       string `json:"email" binding:"required,email"`
		Code        string `json:"code" binding:"required"`
		NewPassword string `json:"password" binding:"required"`
	}

	if err := ctx.ShouldBindJSON(&body); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := h.srv.ResetPassword(ctx.Request.Context(), body.Email, body.Code, body.NewPassword); err != nil {
		if errors.Is(err, ErrInvalidCode) {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		if errors.Is(err, ErrUserNotFound) {
			ctx.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
			return
		}
		slog.ErrorContext(ctx.Request.Context(), "failed to reset password", slog.Any("error", err), slog.String("email", body.Email))
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "failed to reset password"})
		return
	}

	ctx.Status(http.StatusOK)
}

func (h *UserHandler) ListUsers(ctx *gin.Context) {
	idsStr := ctx.Query("ids")
	if idsStr == "" {
		ctx.JSON(http.StatusOK, []User{})
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

func (h *UserHandler) Register(ctx *gin.Context) {
	var body struct {
		FirstName string `json:"firstName" binding:"required,min=1,max=255"`
		LastName  string `json:"lastName" binding:"required,min=1,max=255"`
		Email     string `json:"email" binding:"required,email"`
		Password  string `json:"password" binding:"required,min=8"`
	}

	if err := ctx.ShouldBindJSON(&body); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err := h.srv.Register(ctx.Request.Context(), User{
		FirstName: body.FirstName,
		LastName:  body.LastName,
		Email:     body.Email,
		Password:  body.Password,
	})

	if err != nil {
		if errors.Is(err, ErrUserAlreadyExists) {
			ctx.JSON(http.StatusConflict, gin.H{"error": err.Error()})
			return
		}
		slog.ErrorContext(ctx.Request.Context(), "user registration failed", slog.Any("error", err), slog.String("email", body.Email))
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "failed to register user"})
		return
	}

	ctx.Status(http.StatusCreated)
}

func (h *UserHandler) ChangePassword(ctx *gin.Context) {
	user, ok := h.getAuthUser(ctx)
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

	err := h.srv.ChangePassword(ctx.Request.Context(), user.ID, body.CurrentPassword, body.NewPassword)
	if err != nil {
		if errors.Is(err, ErrUserNotFound) {
			ctx.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
			return
		}
		if errors.Is(err, ErrNotVerified) {
			ctx.JSON(http.StatusForbidden, gin.H{"error": err.Error()})
			return
		}
		if errors.Is(err, ErrInvalidCredentials) {
			ctx.JSON(http.StatusUnauthorized, gin.H{"error": "senha atual incorreta"})
			return
		}
		slog.ErrorContext(ctx.Request.Context(), "failed to change password", slog.Any("error", err), slog.String("user_id", user.ID))
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "failed to change password"})
		return
	}

	ctx.Status(http.StatusOK)
}

func (h *UserHandler) UpdateProfile(ctx *gin.Context) {
	user, ok := h.getAuthUser(ctx)
	if !ok {
		return
	}

	var body struct {
		FirstName     *string     `json:"firstName" binding:"omitempty,min=1,max=255"`
		LastName      *string     `json:"lastName" binding:"omitempty,min=1,max=255"`
		Bio           *string     `json:"bio"`
		Height        *float64    `json:"height"`
		WeightUnit    *WeightUnit `json:"weightUnit"`
		HeightUnit    *HeightUnit `json:"heightUnit"`
		CurrentWeight *float64    `json:"currentWeight"`
	}

	if err := ctx.ShouldBindJSON(&body); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err := h.srv.UpdateProfile(ctx.Request.Context(), user.ID, body.FirstName, body.LastName, body.Bio, body.Height, body.WeightUnit, body.HeightUnit, body.CurrentWeight)
	if err != nil {
		if errors.Is(err, ErrUserNotFound) {
			ctx.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
			return
		}
		slog.ErrorContext(ctx.Request.Context(), "failed to update profile", slog.Any("error", err), slog.String("user_id", user.ID))
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "failed to update profile"})
		return
	}

	ctx.Status(http.StatusOK)
}

func (h *UserHandler) GetUser(ctx *gin.Context) {
	id := ctx.Param("id")

	user, ok := h.getAuthUser(ctx)
	if !ok {
		return
	}

	res, err := h.srv.GetUser(ctx.Request.Context(), id, user.ID)
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

func (h *UserHandler) ListFollower(ctx *gin.Context) {
	id := ctx.Param("id")
	cursor, limit := h.getPagination(ctx)

	users, nextCursor, err := h.srv.ListFollower(ctx.Request.Context(), id, cursor, limit)
	if err != nil {
		slog.ErrorContext(ctx.Request.Context(), "failed to list followers", slog.Any("error", err), slog.String("id", id))
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "failed to list followers"})
		return
	}

	type followerResponse struct {
		Id        string    `json:"id"`
		FirstName string    `json:"firstName"`
		LastName  string    `json:"lastName"`
		Email     string    `json:"email"`
		Type      UserType  `json:"type"`
		CreatedAt time.Time `json:"createdAt"`
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

	ctx.JSON(http.StatusOK, gin.H{
		"data":       data,
		"nextCursor": nextCursor,
	})
}

func (h *UserHandler) ListFollowing(ctx *gin.Context) {
	id := ctx.Param("id")
	cursor, limit := h.getPagination(ctx)

	users, nextCursor, err := h.srv.ListFollowing(ctx.Request.Context(), id, cursor, limit)
	if err != nil {
		slog.ErrorContext(ctx.Request.Context(), "failed to list following", slog.Any("error", err), slog.String("id", id))
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "failed to list following"})
		return
	}

	type followerResponse struct {
		Id        string    `json:"id"`
		FirstName string    `json:"firstName"`
		LastName  string    `json:"lastName"`
		Email     string    `json:"email"`
		Type      UserType  `json:"type"`
		CreatedAt time.Time `json:"createdAt"`
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

	ctx.JSON(http.StatusOK, gin.H{
		"data":       data,
		"nextCursor": nextCursor,
	})
}
func (h *UserHandler) CountFollowers(ctx *gin.Context) {
	id := ctx.Param("id")

	count, err := h.srv.CountFollowers(ctx.Request.Context(), id)
	if err != nil {
		slog.ErrorContext(ctx.Request.Context(), "failed to count followers", slog.Any("error", err), slog.String("id", id))
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "failed to count followers"})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"count": count})
}

func (h *UserHandler) CountFollowing(ctx *gin.Context) {
	id := ctx.Param("id")

	count, err := h.srv.CountFollowing(ctx.Request.Context(), id)
	if err != nil {
		slog.ErrorContext(ctx.Request.Context(), "failed to count following", slog.Any("error", err), slog.String("id", id))
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "failed to count following"})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"count": count})
}

func (h *UserHandler) FollowUser(ctx *gin.Context) {
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

func (h *UserHandler) UnfollowUser(ctx *gin.Context) {
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

func (h *UserHandler) LinkTrainer(ctx *gin.Context) {
	user, ok := h.getAuthUser(ctx)
	if !ok {
		return
	}

	var body struct {
		Code string `json:"code" binding:"required"`
	}

	if err := ctx.ShouldBindJSON(&body); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := h.srv.LinkTrainer(ctx.Request.Context(), user.ID, body.Code); err != nil {
		if errors.Is(err, ErrTrainerNotFound) {
			ctx.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
			return
		}
		slog.ErrorContext(ctx.Request.Context(), "failed to link trainer", slog.Any("error", err), slog.String("user_id", user.ID), slog.String("code", body.Code))
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "failed to link trainer"})
		return
	}

	ctx.Status(http.StatusOK)
}

func (h *UserHandler) CreateTrainerCode(ctx *gin.Context) {
	user, ok := h.getAuthUser(ctx)
	if !ok {
		return
	}

	var body struct {
		Code string `json:"code" binding:"required"`
	}

	if err := ctx.ShouldBindJSON(&body); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := h.srv.CreateTrainerCode(ctx.Request.Context(), user.ID, body.Code); err != nil {
		slog.ErrorContext(ctx.Request.Context(), "failed to create trainer code", slog.Any("error", err), slog.String("user_id", user.ID))
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "failed to create trainer code"})
		return
	}

	ctx.Status(http.StatusOK)
}

func (h *UserHandler) UnlinkTrainer(ctx *gin.Context) {
	user, ok := h.getAuthUser(ctx)
	if !ok {
		return
	}

	if err := h.srv.UnlinkTrainer(ctx.Request.Context(), user.ID); err != nil {
		slog.ErrorContext(ctx.Request.Context(), "failed to unlink trainer", slog.Any("error", err), slog.String("user_id", user.ID))
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "failed to unlink trainer"})
		return
	}

	ctx.Status(http.StatusOK)
}

func (h *UserHandler) ListStudents(ctx *gin.Context) {
	user, ok := h.getAuthUser(ctx)
	if !ok {
		return
	}

	cursor, limit := h.getPagination(ctx)

	users, nextCursor, err := h.srv.ListStudents(ctx.Request.Context(), user.ID, cursor, limit)
	if err != nil {
		slog.ErrorContext(ctx.Request.Context(), "failed to list students", slog.Any("error", err), slog.String("trainer_id", user.ID))
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "failed to list students"})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"data":       users,
		"nextCursor": nextCursor,
	})
}

func (h *UserHandler) UnlinkStudent(ctx *gin.Context) {
	id := ctx.Param("id")

	if err := h.srv.UnlinkStudent(ctx.Request.Context(), id); err != nil {
		slog.ErrorContext(ctx.Request.Context(), "failed to unlink student", slog.Any("error", err), slog.String("student_id", id))
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "failed to unlink student"})
		return
	}

	ctx.Status(http.StatusOK)
}

func (h *UserHandler) AddBodyMeasurementNote(ctx *gin.Context) {
	id := ctx.Param("id")

	var body struct {
		Note string `json:"note" binding:"required"`
	}

	if err := ctx.ShouldBindJSON(&body); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := h.srv.AddBodyMeasurementNote(ctx.Request.Context(), id, body.Note); err != nil {
		slog.ErrorContext(ctx.Request.Context(), "failed to add body measurement note", slog.Any("error", err), slog.String("measurement_id", id))
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "failed to add note"})
		return
	}

	ctx.Status(http.StatusOK)
}

func (h *UserHandler) FindLastBodyMeasurementNote(ctx *gin.Context) {
	user, ok := h.getAuthUser(ctx)
	if !ok {
		return
	}

	measurement, err := h.srv.FindLastBodyMeasurementNote(ctx.Request.Context(), user.ID)
	if err != nil {
		slog.ErrorContext(ctx.Request.Context(), "failed to fetch last measurement", slog.Any("error", err), slog.String("user_id", user.ID))
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "failed to fetch last measurement"})
		return
	}

	ctx.JSON(http.StatusOK, measurement)
}

func (h *UserHandler) ListBodyMeasurements(ctx *gin.Context) {
	id := ctx.Param("id")
	if id == "" {
		user, ok := ctx.Value(string(auth.UserContextKey)).(auth.AuthUser)
		if !ok {
			ctx.JSON(http.StatusUnauthorized, gin.H{"error": "unauthorized"})
			return
		}
		id = user.ID
	}

	cursor, limit := h.getPagination(ctx)

	measurements, nextCursor, err := h.srv.ListBodyMeasurements(ctx.Request.Context(), id, cursor, limit)
	if err != nil {
		slog.ErrorContext(ctx.Request.Context(), "failed to list body measurements", slog.Any("error", err), slog.String("user_id", id))
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "failed to list measurements"})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"data":       measurements,
		"nextCursor": nextCursor,
	})
}

func (h *UserHandler) AddGoalMetric(ctx *gin.Context) {
	user, ok := h.getAuthUser(ctx)
	if !ok {
		return
	}

	var body struct {
		Type        string    `json:"type" binding:"required"`
		TargetValue float64   `json:"targetValue" binding:"required,gt=0"`
		Deadline    time.Time `json:"deadline" binding:"required"`
	}

	if err := ctx.ShouldBindJSON(&body); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := h.srv.AddGoalMetric(ctx.Request.Context(), &MetricGoal{
		UserId:      user.ID,
		Type:        body.Type,
		TargetValue: body.TargetValue,
		Deadline:    &body.Deadline,
	}); err != nil {
		slog.ErrorContext(ctx.Request.Context(), "failed to add goal metric", slog.Any("error", err), slog.String("user_id", user.ID))
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "failed to add goal metric"})
		return
	}

	ctx.Status(http.StatusOK)
}

func (h *UserHandler) ListGoalsMetric(ctx *gin.Context) {
	user, ok := h.getAuthUser(ctx)
	if !ok {
		return
	}
	cursor, limit := h.getPagination(ctx)

	goals, nextCursor, err := h.srv.ListGoalsMetric(ctx.Request.Context(), user.ID, cursor, limit)
	if err != nil {
		slog.ErrorContext(ctx.Request.Context(), "failed to list goals", slog.Any("error", err), slog.String("user_id", user.ID))
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "failed to list goals"})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"data":       goals,
		"nextCursor": nextCursor,
	})
}

func (h *UserHandler) ListGoalsMetricById(ctx *gin.Context) {
	id := ctx.Param("id")
	cursor, limit := h.getPagination(ctx)

	goals, nextCursor, err := h.srv.ListGoalsMetric(ctx.Request.Context(), id, cursor, limit)
	if err != nil {
		slog.ErrorContext(ctx.Request.Context(), "failed to list goals by id", slog.Any("error", err), slog.String("user_id", id))
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "failed to list goals"})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"data":       goals,
		"nextCursor": nextCursor,
	})
}

func (h *UserHandler) AddWeightLogNote(ctx *gin.Context) {
	id := ctx.Param("id")

	var body struct {
		Note string `json:"note" binding:"required"`
	}

	if err := ctx.ShouldBindJSON(&body); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := h.srv.AddWeightLogNote(ctx.Request.Context(), id, body.Note); err != nil {
		slog.ErrorContext(ctx.Request.Context(), "failed to add weight log note", slog.Any("error", err), slog.String("log_id", id))
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "failed to add note"})
		return
	}

	ctx.Status(http.StatusOK)
}

func (h *UserHandler) ChangeToTrainer(ctx *gin.Context) {
	user, ok := h.getAuthUser(ctx)
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

	if err := h.srv.ChangeToTrainer(ctx.Request.Context(), user.ID, body.Cref); err != nil {
		slog.ErrorContext(ctx.Request.Context(), "failed to upgrade profile", slog.Any("error", err), slog.String("user_id", user.ID))
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "failed to upgrade profile"})
		return
	}

	ctx.Status(http.StatusOK)
}

func (h *UserHandler) ChangeToClient(ctx *gin.Context) {
	user, ok := h.getAuthUser(ctx)
	if !ok {
		return
	}

	if err := h.srv.ChangeToClient(ctx.Request.Context(), user.ID); err != nil {
		slog.ErrorContext(ctx.Request.Context(), "failed to downgrade profile", slog.Any("error", err), slog.String("user_id", user.ID))
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "failed to downgrade profile"})
		return
	}

	ctx.Status(http.StatusOK)
}

func (h *UserHandler) RemoveProfilePicture(ctx *gin.Context) {
	user, ok := h.getAuthUser(ctx)
	if !ok {
		return
	}

	if err := h.srv.RemoveProfilePicture(ctx.Request.Context(), user.ID); err != nil {
		slog.ErrorContext(ctx.Request.Context(), "failed to remove profile picture", slog.Any("error", err), slog.String("user_id", user.ID))
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "failed to remove profile picture"})
		return
	}

	ctx.Status(http.StatusNoContent)
}

func (h *UserHandler) ListWeightLogs(ctx *gin.Context) {
	id := ctx.Param("id")
	if id == "" {
		user, ok := h.getAuthUser(ctx)
		if !ok {
			return
		}
		id = user.ID
	}

	cursor, limit := h.getPagination(ctx)

	logs, nextCursor, err := h.srv.ListWeightLogs(ctx.Request.Context(), id, cursor, limit)
	if err != nil {
		slog.ErrorContext(ctx.Request.Context(), "failed to list weight logs", slog.Any("error", err), slog.String("user_id", id))
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "failed to list weight logs"})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"data":       logs,
		"nextCursor": nextCursor,
	})
}

func (h *UserHandler) UploadProfilePicture(ctx *gin.Context) {
	user, ok := h.getAuthUser(ctx)
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
		slog.ErrorContext(ctx.Request.Context(), "failed to open file", slog.Any("error", err), slog.String("user_id", user.ID))
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "failed to open file"})
		return
	}
	defer openedFile.Close()

	if err := h.srv.UploadProfilePicture(ctx.Request.Context(), user.ID, openedFile); err != nil {
		slog.ErrorContext(ctx.Request.Context(), "failed to upload profile picture", slog.Any("error", err), slog.String("user_id", user.ID))
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "failed to upload picture"})
		return
	}

	ctx.Status(http.StatusOK)
}

func (h *UserHandler) getAuthUser(ctx *gin.Context) (auth.AuthUser, bool) {
	user, ok := ctx.Value(string(auth.UserContextKey)).(auth.AuthUser)
	if !ok {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "unauthorized"})
		return auth.AuthUser{}, false
	}

	return user, true
}

func (h *UserHandler) getPagination(ctx *gin.Context) (string, int) {
	cursor := ctx.Query("cursor")
	limit, _ := strconv.Atoi(ctx.DefaultQuery("limit", "20"))
	if limit <= 0 {
		limit = 20
	}
	return cursor, limit
}
