package main

import (
	"context"
	"database/sql"
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/kaua-nasc/gymtrack-go/apps/api/training-plan/internal"
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
	r := gin.Default()

	handler.RegisterRoutes(r)

	server := &http.Server{
		Addr:    ":3333",
		Handler: r,
	}

	lc.Append(fx.Hook{
		OnStart: func(ctx context.Context) error {
			log.Println("API Running, port :3333")
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
