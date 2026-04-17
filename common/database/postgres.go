package database

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

// Config represents the database connection parameters
type Config struct {
	Host     string
	Port     string
	User     string
	Password string
	DBName   string
	SSLMode  string
	TimeZone string
}

// InitDB initializes a PostgreSQL connection using the provided configuration.
func InitDB(cfg Config) (*sql.DB, error) {
	// Apply default values if not specified
	if cfg.SSLMode == "" {
		cfg.SSLMode = "disable"
	}
	if cfg.TimeZone == "" {
		cfg.TimeZone = "Asia/Jakarta"
	}

	// Construct DSN string
	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=%s TimeZone=%s",
		cfg.Host, cfg.User, cfg.Password, cfg.DBName, cfg.Port, cfg.SSLMode, cfg.TimeZone,
	)

	db, err := sql.Open("postgres", dsn)
	if err != nil {
		return nil, fmt.Errorf("error opening database: %w", err)
	}

	if err := db.Ping(); err != nil {
		return nil, fmt.Errorf("error connecting to database [%s]: %w", cfg.DBName, err)
	}

	return db, nil
}