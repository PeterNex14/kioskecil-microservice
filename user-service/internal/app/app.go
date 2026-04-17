package app

import (
	"context"
	"database/sql"
	"fmt"
	"log/slog"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/PeterNex14/kioskecil-microservice/common/database"
	"github.com/PeterNex14/kioskecil-microservice/common/logger"
	"github.com/PeterNex14/kioskecil-microservice/common/system"
	db_users_gen "github.com/PeterNex14/kioskecil-microservice/user-service/db/sqlc"
	"github.com/PeterNex14/kioskecil-microservice/user-service/internal/config"
	"github.com/PeterNex14/kioskecil-microservice/user-service/internal/handler"
	"github.com/PeterNex14/kioskecil-microservice/user-service/internal/repository"
	"github.com/PeterNex14/kioskecil-microservice/user-service/internal/service"

	_ "github.com/lib/pq"
)

// App holds the dependencies for the User Service
type App struct {
	DB          *sql.DB
	Config      *config.Config
	HttpServer  *http.Server
	UserHandler *handler.UserHandler
}

// New initializes the App with its dependencies based on the provided config
func New(cfg *config.Config) (*App, error) {
	// 1. Initialize Logger
	logger.InitLogger(cfg.Env, cfg.ServiceName)

	// 2. Initialize Database using generic InitDB
	db, err := database.InitDB(cfg.DB)
	if err != nil {
		slog.Error("failed to connect to database", "error", err, "db_name", cfg.DB.DBName)
		return nil, err
	}

	// 3. Dependency Injection Wiring
	dbQueries := db_users_gen.New(db)
	userRepo := repository.NewUserRepository(dbQueries)
	userSvc := service.NewUserService(userRepo)
	userHandler := handler.NewUserHandler(userSvc)

	// 4. Initialize Gin Engine
	if cfg.Env == "production" {
		gin.SetMode(gin.ReleaseMode)
	}

	router := gin.Default()

	// Register global health check
	router.GET("/health", userHandler.HealthCheck)

	// Register API routes
	apiV1 := router.Group("/api/v1")
	userHandler.RegisterRoutes(apiV1)

	// 5. Setup HTTP Server
	// We use port 8080 as defined in docker-compose
	srv := &http.Server{
		Addr:    fmt.Sprintf(":%s", "8080"),
		Handler: router,
	}

	return &App{
		DB:          db,
		Config:      cfg,
		HttpServer:  srv,
		UserHandler: userHandler,
	}, nil
}

// Run starts the application and waits for an exit signal
func (a *App) Run() {
	// Start HTTP server in a goroutine so it doesn't block the main thread
	go func() {
		slog.Info("Starting HTTP server", "addr", a.HttpServer.Addr)
		if err := a.HttpServer.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			slog.Error("failed to start http server", "error", err)
		}
	}()

	slog.Info("Service started",
		"service", a.Config.ServiceName,
		"env", a.Config.Env,
	)

	// Block until signal is received
	system.WaitExitSignal()

	a.Shutdown()
}

// Shutdown handles the cleanup of resources
func (a *App) Shutdown() {
	slog.Info("Shutting down service...")

	// 1. Shutdown HTTP Server with a timeout context
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := a.HttpServer.Shutdown(ctx); err != nil {
		slog.Error("failed to shutdown http server", "error", err)
	}

	// 2. Close Database Connection
	if err := a.DB.Close(); err != nil {
		slog.Error("failed to close database connection", "error", err)
	}

	slog.Info("Service stopped gracefully")
}
