package config

import (
	"os"
)

type Config struct {
	DatabaseURL   string
	ServerAddress string
}

func LoadConfig() (*Config, error) {
	// Можно улучшить загрузку через env или флаг, здесь простая реализация
	return &Config{
		DatabaseURL:   getEnv("DATABASE_URL", "postgres://user:password@localhost:5432/dbname?sslmode=disable"),
		ServerAddress: getEnv("SERVER_ADDRESS", ":8080"),
	}, nil
}

func getEnv(key, fallback string) string {
	if val, exists := os.LookupEnv(key); exists {
		return val
	}
	return fallback
}
