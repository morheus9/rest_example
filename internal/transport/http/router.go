package http

import (
	"net/http"
	"strconv"
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

	return mux
}
