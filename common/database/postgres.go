package database

import (
	"database/sql"
	"fmt"

	"github.com/PeterNex14/kioskecil-microservice/common/config"
)

func InitDB() (*sql.DB, error) {
	// Use config.GetEnv to provide default values if environment variables are missing
	host := config.GetEnv("DB_HOST", "localhost")
	user := config.GetEnv("USER_DB_USER", "postgres")
	password := config.GetEnv("USER_DB_PASSWORD", "password")
	dbname := config.GetEnv("USER_DB_NAME", "postgres")
	port := config.GetEnv("DB_PORT", "5432")

	// Construct DSN string safely
	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Asia/Jakarta",
		host, user, password, dbname, port,
	)

	db, err := sql.Open("postgres", dsn)
	if err != nil {
		return nil, fmt.Errorf("error opening database: %v", err)
	}

	if err := db.Ping(); err != nil {
		return nil, fmt.Errorf("error connecting to database [%s]: %v", dbname, err)
	}

	return db, nil
}