package config

import (
	"github.com/PeterNex14/kioskecil-microservice/common/config"
	"github.com/PeterNex14/kioskecil-microservice/common/database"
)

// Config holds all configuration for the User Service
type Config struct {
	Env         string
	ServiceName string
	JWTSecret   string
	DB          database.Config
}

// Load populates the Config struct from environment variables
func Load() *Config {
	return &Config{
		Env:         config.GetEnv("APP_ENV", "development"),
		ServiceName: config.GetEnv("SERVICE_NAME", "user-service"),
		JWTSecret:   config.GetEnv("JWT_SECRET", "very-secret-key"),
		DB: database.Config{
			Host:     config.GetEnv("DB_HOST", "localhost"),
			Port:     config.GetEnv("DB_PORT", "5432"),
			User:     config.GetEnv("USER_DB_USER", "postgres"),
			Password: config.GetEnv("USER_DB_PASSWORD", "password"),
			DBName:   config.GetEnv("USER_DB_NAME", "db_users"),
			SSLMode:  config.GetEnv("DB_SSLMODE", "disable"),
		},
	}
}
