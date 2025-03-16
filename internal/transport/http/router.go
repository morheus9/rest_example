package http

import (
	"log/slog"
	"net/http"
	"strconv"
	"time"
)

func NewRouter(handler *Handler) http.Handler {
	mux := http.NewServeMux()

	// Регистрация маршрутов
	mux.HandleFunc("POST /users", func(w http.ResponseWriter, r *http.Request) {
		handler.CreateUser(w, r)
	})

	mux.HandleFunc("GET /users/{id}", func(w http.ResponseWriter, r *http.Request) {
		idStr := r.PathValue("id")
		id, err := strconv.ParseInt(idStr, 10, 64)
		if err != nil {
			http.Error(w, "Invalid ID", http.StatusBadRequest)
			return
		}
		handler.GetUser(w, r, id)
	})

	// Обернуть маршрутизатор в middleware для логирования
	return LoggerMiddleware(mux)
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

		// Используем ResponseWriter для захвата статуса ответа
		rw := &responseWriter{w, http.StatusOK}
		next.ServeHTTP(rw, r)

		slog.Info("Request completed",
			"method", r.Method,
			"path", r.URL.Path,
			"status", rw.status,
			"duration", time.Since(start),
		)
	})
}

// responseWriter используется для захвата статуса ответа
type responseWriter struct {
	http.ResponseWriter
	status int
}

func (rw *responseWriter) WriteHeader(code int) {
	rw.status = code
	rw.ResponseWriter.WriteHeader(code)
}
