package app

import (
	"database/sql"
	"log/slog"

	"github.com/PeterNex14/kioskecil-microservice/common/database"
	"github.com/PeterNex14/kioskecil-microservice/common/logger"
	"github.com/PeterNex14/kioskecil-microservice/common/system"
	db_users_gen "github.com/PeterNex14/kioskecil-microservice/user-service/db/sqlc"
	"github.com/PeterNex14/kioskecil-microservice/user-service/internal/config"

	_ "github.com/lib/pq"
)

// App holds the dependencies for the User Service
type App struct {
	DB      *sql.DB
	Queries *db_users_gen.Queries
	Config  *config.Config
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

	dbQueries := db_users_gen.New(db)

	return &App{
		DB:      db,
		Queries: dbQueries,
		Config:  cfg,
	}, nil
}

// Run starts the application and waits for an exit signal
func (a *App) Run() {
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

	if err := a.DB.Close(); err != nil {
		slog.Error("failed to close database connection", "error", err)
	}

	slog.Info("Service stopped gracefully")
}
