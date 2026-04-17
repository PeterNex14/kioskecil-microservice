package app

import (
	"database/sql"
	"log/slog"
	"os"

	"github.com/PeterNex14/kioskecil-microservice/common/config"
	"github.com/PeterNex14/kioskecil-microservice/common/database"
	"github.com/PeterNex14/kioskecil-microservice/common/logger"
	"github.com/PeterNex14/kioskecil-microservice/common/system"
	db_users_gen "github.com/PeterNex14/kioskecil-microservice/user-service/db/sqlc"

	_ "github.com/lib/pq"
)

// App holds the dependencies for the User Service
type App struct {
	DB        *sql.DB
	Queries   *db_users_gen.Queries
	Config    *config.BaseConfig
	JWTSecret string
}

// New initializes the App with its dependencies
func New() (*App, error) {
	env := config.GetEnv("APP_ENV", "development")
	serviceName := config.GetEnv("SERVICE_NAME", "user-service")

	// 1. Initialize Logger
	logger.InitLogger(env, serviceName)

	// 2. Initialize Database
	db, err := database.InitDB()
	if err != nil {
		slog.Error("failed to connect to database", "error", err, "db_name", os.Getenv("DB_NAME"))
		return nil, err
	}

	dbQueries := db_users_gen.New(db)
	baseCfg := config.NewBaseConfig(db, serviceName, env)

	return &App{
		DB:        db,
		Queries:   dbQueries,
		Config:    &baseCfg,
		JWTSecret: os.Getenv("JWT_SECRET"),
	}, nil
}

// Run starts the application and waits for an exit signal
func (a *App) Run() {
	slog.Info("Service started",
		"service", a.Config.ServiceName,
		"env", a.Config.Environment,
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
