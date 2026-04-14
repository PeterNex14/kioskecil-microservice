package config

import "gorm.io/gorm"

type BaseConfig struct {
	DB 				*gorm.DB
	ServiceName		string
	Environment		string
}

func NewBaseConfig(db *gorm.DB, serviceName string, env string) BaseConfig {
	return BaseConfig{
		DB: db,
		ServiceName: serviceName,
		Environment: env,
	}
}