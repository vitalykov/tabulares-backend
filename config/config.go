package config

import (
	"os"
)

type DBConfig struct {
	DBHost     string
	DBPort     string
	DBName     string
	DBUser     string
	DBPassword string
}

func NewDBConfig() *DBConfig {
	return &DBConfig{
		DBName:     os.Getenv("DB_NAME"),
		DBPort:     os.Getenv("DB_PORT"),
		DBUser:     os.Getenv("DB_USER"),
		DBPassword: os.Getenv("DB_PASSWORD"),
		DBHost:     os.Getenv("DB_HOST"),
	}
}
