package main

import (
	"log"
	"os"

	"github.com/PeterNex14/kioskecil-microservice/common/config"
	"github.com/PeterNex14/kioskecil-microservice/common/database"
)

type apiConfig struct {
	config.BaseConfig
	JWTSecret 			string
}

func main() {
	env := os.Getenv("APP_ENV")
	serviceName := os.Getenv("SERVICE_NAME")

	db := database.InitDB()

	cfg := &apiConfig{
		BaseConfig: config.NewBaseConfig(db, serviceName, env),
		JWTSecret: os.Getenv("JWT_SECRET"),
	}

	log.Printf("Starting [%s] in [%s] mode", cfg.ServiceName, cfg.Environment)
}
