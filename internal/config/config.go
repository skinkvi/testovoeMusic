package config

import (
	"os"

	"go.uber.org/zap"
)

type Config struct {
	ServerPort string
	DbHost     string
	DbPort     string
	DbUser     string
	DbPassword string
	DbName     string
	Logger     *zap.Logger
}

func GetConfig() (*Config, error) {
	return &Config{
		ServerPort: os.Getenv("SERVER_PORT"),
		DbHost:     os.Getenv("DB_HOST"),
		DbPort:     os.Getenv("DB_PORT"),
		DbUser:     os.Getenv("DB_USER"),
		DbPassword: os.Getenv("DB_PASSWORD"),
		DbName:     os.Getenv("DB_NAME"),
		Logger:     zap.NewExample(),
	}, nil
}
