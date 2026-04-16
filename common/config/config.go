package config

import (
	"database/sql"
)

type BaseConfig struct {
	DB 				*sql.DB
	ServiceName		string
	Environment		string
}

func NewBaseConfig(db *sql.DB, serviceName string, env string) BaseConfig {
	return BaseConfig{
		DB: db,
		ServiceName: serviceName,
		Environment: env,
	}
}