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
	"github.com/kaua-nasc/gymtrack-go/apps/api/social/internal"
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

func NewHTTPServer(lc fx.Lifecycle, handler *internal.PostHandler) *gin.Engine {
	port := os.Getenv("SOCIAL_PORT")
	if port == "" {
		port = os.Getenv("PORT")
	}
	if port == "" {
		port = ":8082"
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

	handler.RegisterRoutes(r)

	server := &http.Server{
		Addr:    port,
		Handler: r,
	}

	lc.Append(fx.Hook{
		OnStart: func(ctx context.Context) error {
			log.Println("Social API Running, port " + port)
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
			internal.NewIdentityService,
			internal.NewTrainingPlanClient,
			internal.NewPostRepository,
			internal.NewPostService,
			internal.NewPostHandler,
			NewHTTPServer,
		),
		fx.Invoke(func(*gin.Engine) {}),
	).Run()
}
