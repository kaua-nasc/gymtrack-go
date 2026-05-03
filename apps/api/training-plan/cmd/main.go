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
	"github.com/kaua-nasc/gymtrack-go/apps/api/training-plan/internal"
	"github.com/kaua-nasc/gymtrack-go/libs/config"
	"github.com/kaua-nasc/gymtrack-go/libs/db"
	"go.uber.org/fx"
)

func NewDatabase() (*sql.DB, error) {
	dsn := os.Getenv("DATABASE_URL")
	if dsn == "" {
		panic("DATABASE_URL not found")
	}
	return db.NewConnection(dsn)
}

func NewHTTPServer(lc fx.Lifecycle, handler *internal.TrainingPlanHandler) *gin.Engine {
	port := os.Getenv("PORT")
	if port == "" {
		panic("PORT not found")
	}

	r := gin.Default()

	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:4200"},
		AllowMethods:     []string{"GET", "POST", "PUT", "PATCH", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Accept", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))

	handler.RegisterRoutes(r)

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

	fx.New(
		fx.Provide(
			NewDatabase,
			internal.NewTrainingPlanRepository,
			internal.NewIdentityService,
			internal.NewLocalStorageService,
			internal.NewTrainingPlanService,
			internal.NewTrainingPlanHandler,
			NewHTTPServer,
		),
		fx.Invoke(func(*gin.Engine) {}),
	).Run()
}
