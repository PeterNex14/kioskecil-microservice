package database

import (
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type DBSetupParams struct {
	DBHost string
	DBUser string
	DBPassword string
	DBName string
	DBPort string
	DBSslmode string `default:"disable"`
}

func InitDB(db_config DBSetupParams) (*gorm.DB, error) {
	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=%s",
		db_config.DBHost,
		db_config.DBUser,
		db_config.DBPassword,
		db_config.DBName,
		db_config.DBPort,
		db_config.DBSslmode,
	)

	return gorm.Open(postgres.Open(dsn), &gorm.Config{})
}