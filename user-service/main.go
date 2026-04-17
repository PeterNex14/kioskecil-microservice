package main

import (
	"log/slog"
	"os"

	"github.com/PeterNex14/kioskecil-microservice/common/config"
	"github.com/PeterNex14/kioskecil-microservice/common/database"
	"github.com/PeterNex14/kioskecil-microservice/common/logger"
	db_users_gen "github.com/PeterNex14/kioskecil-microservice/user-service/db/sqlc"

	_ "github.com/lib/pq"
)

type apiConfig struct {
	config.BaseConfig
	JWTSecret 			string
	dbQueries			*db_users_gen.Queries
}

func main() {
	// 1. Initialize Logger
	env := config.GetEnv("APP_ENV", "development")
	serviceName := config.GetEnv("SERVICE_NAME", "user-service")
	logger.InitLogger(env, serviceName)

	// 2. Initialize Database
	db, err := database.InitDB()
	if err != nil {
		slog.Error("failed to connect to database", "error", err, "db_name", os.Getenv("DB_NAME"))
		os.Exit(1)
	}
	defer db.Close()

	dbQueries := db_users_gen.New(db)

	cfg := &apiConfig{
		BaseConfig: config.NewBaseConfig(db, serviceName, env),
		JWTSecret: os.Getenv("JWT_SECRET"),
		dbQueries: dbQueries,
	}

	slog.Info("Service started", 
		"service", cfg.ServiceName, 
		"env", cfg.Environment,
	)
}
