package http

import (
	"net/http"
	"strconv"
	"strings"
)

func NewRouter(handler *Handler) http.Handler {
	mux := http.NewServeMux()

	// Регистрация маршрутов
	mux.HandleFunc("/users", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodPost:
			handler.CreateUser(w, r)
		default:
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		}
	})

	mux.HandleFunc("/users/", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodGet:
			// Извлекаем ID из URL
			parts := strings.Split(r.URL.Path, "/")
			if len(parts) < 3 {
				http.Error(w, "Invalid URL", http.StatusBadRequest)
				return
			}
			idStr := parts[2]
			id, err := strconv.ParseInt(idStr, 10, 64)
			if err != nil {
				http.Error(w, "Invalid ID", http.StatusBadRequest)
				return
			}

			// Вызываем GetUser с извлеченным ID
			handler.GetUser(w, r, id)
		default:
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		}
	})

	return mux
}
