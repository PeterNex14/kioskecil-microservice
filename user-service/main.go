package main

import (
	"log"
	"os"

	"github.com/PeterNex14/kioskecil-microservice/common/config"
	"github.com/PeterNex14/kioskecil-microservice/common/database"
	db_users_gen "github.com/PeterNex14/kioskecil-microservice/user-service/db/sqlc"

	_ "github.com/lib/pq"
)

type apiConfig struct {
	config.BaseConfig
	JWTSecret 			string
	dbQueries			*db_users_gen.Queries
}

func main() {
	env := os.Getenv("APP_ENV")
	serviceName := os.Getenv("SERVICE_NAME")

	db, err := database.InitDB()
	if err != nil {
		log.Fatalf("Gagal koneksi ke database [%s]: %v", os.Getenv("DB_NAME"), err)
	}
	defer db.Close()

	dbQueries := db_users_gen.New(db)

	cfg := &apiConfig{
		BaseConfig: config.NewBaseConfig(db, serviceName, env),
		JWTSecret: os.Getenv("JWT_SECRET"),
		dbQueries: dbQueries,
	}

	log.Printf("Starting [%s] in [%s] mode", cfg.ServiceName, cfg.Environment)
}
