package main

import (
	"context"
	"log"
	"net/http"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"

	"github.com/morheus9/rest_example/internal/config"
	"github.com/morheus9/rest_example/internal/repository"
	"github.com/morheus9/rest_example/internal/service"
	transportHTTP "github.com/morheus9/rest_example/internal/transport/http" // алиас для внутреннего http-пакета
)

func main() {
	// Загрузка конфигурации
	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatalf("error loading config: %v", err)
	}

	// Инициализация пула подключений PGX
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	dbpool, err := pgxpool.New(ctx, cfg.DatabaseURL)
	if err != nil {
		log.Fatalf("unable to connect to database: %v", err)
	}
	defer dbpool.Close()

	// Инициализируем репозиторий, сервисы и HTTP-обработчики
	userRepo := repository.NewPgUserRepository(dbpool)
	userService := service.NewUserService(userRepo)
	handler := transportHTTP.NewHandler(userService)

	// Инициализируем роутер
	router := transportHTTP.NewRouter(handler)

	// Запуск HTTP-сервера
	address := cfg.ServerAddress
	log.Printf("Server starting on %s", address)
	if err := http.ListenAndServe(address, router); err != nil {
		log.Fatalf("Server failed: %v", err)
	}
}
