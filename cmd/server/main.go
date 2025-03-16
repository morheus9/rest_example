package main

import (
	"context"
	"log/slog"
	"net/http"
	"os"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"

	"github.com/morheus9/rest_example/internal/config"
	"github.com/morheus9/rest_example/internal/repository"
	"github.com/morheus9/rest_example/internal/service"
	transportHTTP "github.com/morheus9/rest_example/internal/transport/http"
)

func main() {
	// Настройка логгера
	logger := slog.New(slog.NewTextHandler(os.Stdout, nil))
	slog.SetDefault(logger)

	// Загрузка конфигурации
	cfg, err := config.LoadConfig()
	if err != nil {
		slog.Error("Failed to load config", "error", err)
		os.Exit(1)
	}
	slog.Info("Config loaded successfully")

	// Инициализация пула подключений PGX
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	dbpool, err := pgxpool.New(ctx, cfg.DatabaseURL)
	if err != nil {
		slog.Error("Failed to connect to database", "error", err)
		os.Exit(1)
	}
	defer dbpool.Close()
	slog.Info("Database connection established")

	// Инициализируем репозиторий, сервисы и HTTP-обработчики
	userRepo := repository.NewPgUserRepository(dbpool)
	userService := service.NewUserService(userRepo)
	handler := transportHTTP.NewHandler(userService)

	// Инициализируем роутер с middleware для логирования
	router := transportHTTP.NewRouter(handler)
	http.Handle("/", LoggerMiddleware(router))

	// Запуск HTTP-сервера
	address := cfg.ServerAddress
	slog.Info("Starting server", "address", address)
	if err := http.ListenAndServe(address, nil); err != nil {
		slog.Error("Server failed", "error", err)
		os.Exit(1)
	}
}

// LoggerMiddleware добавляет логирование для каждого запроса.
func LoggerMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()

		slog.Info("Request started",
			"method", r.Method,
			"path", r.URL.Path,
			"remote_addr", r.RemoteAddr,
		)

		next.ServeHTTP(w, r)

		slog.Info("Request completed",
			"method", r.Method,
			"path", r.URL.Path,
			"duration", time.Since(start),
		)
	})
}
