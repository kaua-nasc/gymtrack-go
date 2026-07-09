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
	"github.com/kaua-nasc/gymtrack-go/apps/api/social/internal/comment"
	"github.com/kaua-nasc/gymtrack-go/apps/api/social/internal/domain"
	"github.com/kaua-nasc/gymtrack-go/apps/api/social/internal/like"
	"github.com/kaua-nasc/gymtrack-go/apps/api/social/internal/post"
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

func NewCache(lc fx.Lifecycle) (cache.Cache, error) {
	redisURL := os.Getenv("REDIS_CACHE_CONNECTION_STRING")
	if redisURL == "" {
		panic("REDIS_CACHE_CONNECTION_STRING environment variable is required")
	}

	client, err := cache.NewFxCache(lc, redisURL)
	if err != nil {
		panic("Error to connect to client")
	}

	return client, nil
}

func NewHTTPServer(
	lc fx.Lifecycle,
	postHandler *post.Handler,
	likeHandler *like.Handler,
	commentHandler *comment.Handler,
) *gin.Engine {
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

	postHandler.RegisterRoutes(r)
	likeHandler.RegisterRoutes(r)
	commentHandler.RegisterRoutes(r)

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
		fx.Provide(
			NewDatabase,
			NewCache,
			domain.NewIdentityService,
			domain.NewTrainingPlanClient,
			post.NewRepository,
			post.NewService,
			post.NewHandler,
			like.NewRepository,
			like.NewService,
			like.NewHandler,
			comment.NewRepository,
			comment.NewService,
			comment.NewHandler,
			NewHTTPServer,
		),
		fx.Invoke(func(*gin.Engine) {}),
	).Run()
}
