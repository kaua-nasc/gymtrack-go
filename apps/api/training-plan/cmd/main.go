package main

import (
	"context"
	"database/sql"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/kaua-nasc/gymtrack-go/apps/api/training-plan/internal/domain"
	"github.com/kaua-nasc/gymtrack-go/apps/api/training-plan/internal/exerciselog"
	"github.com/kaua-nasc/gymtrack-go/apps/api/training-plan/internal/feedback"
	"github.com/kaua-nasc/gymtrack-go/apps/api/training-plan/internal/plan"
	"github.com/kaua-nasc/gymtrack-go/apps/api/training-plan/internal/subscription"
	"github.com/kaua-nasc/gymtrack-go/libs/config"
	"github.com/kaua-nasc/gymtrack-go/libs/db"
	logutil "github.com/kaua-nasc/gymtrack-go/libs/log"
	"go.uber.org/fx"
)

func NewDatabase() (*sql.DB, error) {
	dsn := os.Getenv("DATABASE_URL")
	if dsn == "" {
		panic("DATABASE_URL not found")
	}
	return db.NewConnection(dsn)
}

func NewHTTPServer(
	lc fx.Lifecycle,
	planHandler *plan.Handler,
	subHandler *subscription.Handler,
	feedbackHandler *feedback.Handler,
	logHandler *exerciselog.Handler,
) *gin.Engine {
	port := os.Getenv("TRAINING_PLAN_PORT")
	if port == "" {
		panic("TRAINING_PLAN_PORT not found")
	}

	r := gin.Default()

	r.Use(cors.New(cors.Config{
		AllowAllOrigins:  true,
		AllowMethods:     []string{"GET", "POST", "PUT", "PATCH", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Accept", "Authorization", "X-Requested-With"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))

	r.GET("/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"status": "ok", "service": "training-plan"})
	})

	planHandler.RegisterRoutes(r)
	subHandler.RegisterRoutes(r)
	feedbackHandler.RegisterRoutes(r)
	logHandler.RegisterRoutes(r)

	server := &http.Server{
		Addr:    port,
		Handler: r,
	}

	lc.Append(fx.Hook{
		OnStart: func(ctx context.Context) error {
			log.Println("API Running, port " + port)
			go server.ListenAndServe()
			return nil
		},
		OnStop: func(ctx context.Context) error {
			return server.Shutdown(ctx)
		},
	})

	return r
}

func main() {
	config.LoadEnvironmentVariable()
	logutil.InitLogger()

	fx.New(
		fx.Provide(
			NewDatabase,
			domain.NewIdentityService,
			plan.NewRepository,
			plan.NewService,
			plan.NewHandler,
			subscription.NewRepository,
			subscription.NewService,
			subscription.NewHandler,
			feedback.NewRepository,
			feedback.NewService,
			feedback.NewHandler,
			exerciselog.NewRepository,
			exerciselog.NewService,
			exerciselog.NewHandler,
			NewHTTPServer,
		),
		fx.Invoke(func(*gin.Engine) {}),
	).Run()
}

