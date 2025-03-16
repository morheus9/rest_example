package config

import (
	"fmt"
	"os"
)

// Config содержит конфигурационные параметры приложения.
type Config struct {
	DatabaseURL   string // URL для подключения к базе данных
	ServerAddress string // Адрес и порт сервера
}

// LoadConfig загружает конфигурацию из переменных окружения.
// Переменные окружения должны быть установлены из секретов Kubernetes.
func LoadConfig() (*Config, error) {
	// Загрузка параметров подключения к базе данных из переменных окружения
	dbHost, err := getEnv("DB_HOST", "")
	if err != nil {
		return nil, fmt.Errorf("failed to load DB_HOST: %w", err)
	}

	dbPort, err := getEnv("DB_PORT", "")
	if err != nil {
		return nil, fmt.Errorf("failed to load DB_PORT: %w", err)
	}

	dbUser, err := getEnv("DB_USER", "")
	if err != nil {
		return nil, fmt.Errorf("failed to load DB_USER: %w", err)
	}

	dbPassword, err := getEnv("DB_PASSWORD", "")
	if err != nil {
		return nil, fmt.Errorf("failed to load DB_PASSWORD: %w", err)
	}

	dbName, err := getEnv("DB_NAME", "")
	if err != nil {
		return nil, fmt.Errorf("failed to load DB_NAME: %w", err)
	}

	// Формирование DatabaseURL
	databaseURL := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable",
		dbUser,
		dbPassword,
		dbHost,
		dbPort,
		dbName,
	)

	// Загрузка адреса сервера
	serverAddr, err := getEnv("SERVER_ADDRESS", ":8080")
	if err != nil {
		return nil, fmt.Errorf("failed to load SERVER_ADDRESS: %w", err)
	}

	return &Config{
		DatabaseURL:   databaseURL,
		ServerAddress: serverAddr,
	}, nil
}

// getEnv возвращает значение переменной окружения.
// Если переменная не задана и fallback пустой, возвращает ошибку.
func getEnv(key, fallback string) (string, error) {
	val, exists := os.LookupEnv(key)
	if !exists {
		if fallback == "" {
			return "", fmt.Errorf("environment variable %s is required but not set", key)
		}
		return fallback, nil
	}
	return val, nil
}
