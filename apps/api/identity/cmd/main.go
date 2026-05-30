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
	"github.com/kaua-nasc/gymtrack-go/apps/api/identity/internal/auth"
	"github.com/kaua-nasc/gymtrack-go/apps/api/identity/internal/dashboard"
	"github.com/kaua-nasc/gymtrack-go/apps/api/identity/internal/domain"
	"github.com/kaua-nasc/gymtrack-go/apps/api/identity/internal/followers"
	"github.com/kaua-nasc/gymtrack-go/apps/api/identity/internal/metrics"
	"github.com/kaua-nasc/gymtrack-go/apps/api/identity/internal/trainer"
	"github.com/kaua-nasc/gymtrack-go/apps/api/identity/internal/user"
	"github.com/kaua-nasc/gymtrack-go/libs/cache"
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
	authHandler *auth.Handler,
	userHandler *user.Handler,
	trainerHandler *trainer.Handler,
	followersHandler *followers.Handler,
	metricsHandler *metrics.Handler,
	dashboardHandler *dashboard.Handler,
) *gin.Engine {
	port := os.Getenv("IDENTITY_PORT")
	if port == "" {
		panic("IDENTITY_PORT not found")
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

	authHandler.RegisterRoutes(r)
	userHandler.RegisterRoutes(r)
	trainerHandler.RegisterRoutes(r)
	followersHandler.RegisterRoutes(r)
	metricsHandler.RegisterRoutes(r)
	dashboardHandler.RegisterRoutes(r)

	server := &http.Server{
		Addr:    port,
		Handler: r,
	}

	lc.Append(fx.Hook{
		OnStart: func(ctx context.Context) error {
			log.Println("Identity API Running, port " + port)
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
		cache.Module,
		fx.Provide(
			NewDatabase,
			auth.NewRepository,
			auth.NewService,
			auth.NewHandler,
			user.NewRepository,
			user.NewService,
			user.NewHandler,
			trainer.NewRepository,
			trainer.NewService,
			trainer.NewHandler,
			followers.NewRepository,
			followers.NewService,
			followers.NewHandler,
			metrics.NewRepository,
			metrics.NewService,
			metrics.NewHandler,
			dashboard.NewService,
			dashboard.NewHandler,
			domain.NewTrainingPlanService,
			NewHTTPServer,
		),
		fx.Invoke(func(*gin.Engine) {}),
	).Run()
}
