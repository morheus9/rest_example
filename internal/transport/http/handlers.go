package http

import (
	"encoding/json"
	"log/slog"
	"net/http"

	"github.com/morheus9/rest_example/internal/domain"
	userService "github.com/morheus9/rest_example/internal/service"
)

// Handler структура для обработчиков HTTP-запросов
type Handler struct {
	userService userService.UserService
}

// NewHandler создает новый экземпляр Handler
func NewHandler(us userService.UserService) *Handler {
	return &Handler{userService: us}
}

// CreateUser обрабатывает запрос на создание пользователя
func (h *Handler) CreateUser(w http.ResponseWriter, r *http.Request) {
	// Логируем начало обработки запроса
	slog.Info("CreateUser: started handling request")

	// Декодируем JSON-запрос в структуру запроса
	var userRequest struct {
		Name  string `json:"name"`
		Email string `json:"email"`
	}
	if err := json.NewDecoder(r.Body).Decode(&userRequest); err != nil {
		slog.Error("CreateUser: failed to decode request body", "error", err)
		http.Error(w, "Bad request", http.StatusBadRequest)
		return
	}

	// Логируем успешное декодирование запроса
	slog.Info("CreateUser: request body decoded successfully", "name", userRequest.Name, "email", userRequest.Email)

	// Маппинг входных данных в модель domain.User
	user := &domain.User{
		Name:  userRequest.Name,
		Email: userRequest.Email,
	}

	// Передача в сервис для создания пользователя
	created, err := h.userService.CreateUser(r.Context(), user)
	if err != nil {
		slog.Error("CreateUser: failed to create user", "error", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Логируем успешное создание пользователя
	slog.Info("CreateUser: user created successfully", "user_id", created.ID)

	// Формируем JSON-ответ
	resp, err := json.Marshal(created)
	if err != nil {
		slog.Error("CreateUser: failed to marshal response", "error", err)
		http.Error(w, "Internal error", http.StatusInternalServerError)
		return
	}

	// Логируем успешное формирование ответа
	slog.Info("CreateUser: response marshaled successfully")

	// Отправляем ответ
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	w.Write(resp)
}

// GetUser обрабатывает запрос на получение пользователя по ID
func (h *Handler) GetUser(w http.ResponseWriter, r *http.Request, id int64) {
	// Логируем начало обработки запроса
	slog.Info("GetUser: started handling request", "user_id", id)

	// Получаем пользователя через сервис
	user, err := h.userService.GetUser(r.Context(), id)
	if err != nil {
		slog.Error("GetUser: failed to get user", "error", err, "user_id", id)
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	// Логируем успешное получение пользователя
	slog.Info("GetUser: user retrieved successfully", "user_id", id)

	// Формируем JSON-ответ
	resp, err := json.Marshal(user)
	if err != nil {
		slog.Error("GetUser: failed to marshal response", "error", err)
		http.Error(w, "Internal error", http.StatusInternalServerError)
		return
	}

	// Логируем успешное формирование ответа
	slog.Info("GetUser: response marshaled successfully")

	// Отправляем ответ
	w.Header().Set("Content-Type", "application/json")
	w.Write(resp)
}
