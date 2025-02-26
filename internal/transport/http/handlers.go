package http

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/julienschmidt/httprouter"

	"github.com/morheus9/rest_example/internal/domain"
	userService "github.com/morheus9/rest_example/internal/service"
)

type Handler struct {
	userService userService.UserService
}

func NewHandler(us userService.UserService) *Handler {
	return &Handler{userService: us}
}

func (h *Handler) CreateUser(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	// Декодируем JSON-запрос в структуру запроса
	var userRequest struct {
		Name  string `json:"name"`
		Email string `json:"email"`
	}
	if err := json.NewDecoder(r.Body).Decode(&userRequest); err != nil {
		http.Error(w, "Bad request", http.StatusBadRequest)
		return
	}

	// Маппинг входных данных в модель domain.User
	user := &domain.User{
		Name:  userRequest.Name,
		Email: userRequest.Email,
	}

	// Передача в сервис для создания пользователя
	created, err := h.userService.CreateUser(r.Context(), user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Формируем JSON-ответ
	resp, err := json.Marshal(created)
	if err != nil {
		http.Error(w, "Internal error", http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	w.Write(resp)
}

func (h *Handler) GetUser(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	idStr := ps.ByName("id")
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		http.Error(w, "Invalid id", http.StatusBadRequest)
		return
	}

	// Получаем пользователя через сервис
	user, err := h.userService.GetUser(r.Context(), id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	// Формируем JSON-ответ
	resp, err := json.Marshal(user)
	if err != nil {
		http.Error(w, "Internal error", http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(resp)
}
