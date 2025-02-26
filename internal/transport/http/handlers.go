package http

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/julienschmidt/httprouter"

	userService "github.com/morheus9/rest_example/internal/service"
)

type Handler struct {
	userService userService.UserService
}

func NewHandler(us userService.UserService) *Handler {
	return &Handler{
		userService: us,
	}
}

func (h *Handler) CreateUser(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	var user struct {
		Name  string `json:"name"`
		Email string `json:"email"`
	}
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		http.Error(w, "Bad request", http.StatusBadRequest)
		return
	}

	created, err := h.userService.CreateUser(r.Context(), &userServiceUser{
		Name:  user.Name,
		Email: user.Email,
	})
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

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

	user, err := h.userService.GetUser(r.Context(), id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	resp, err := json.Marshal(user)
	if err != nil {
		http.Error(w, "Internal error", http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(resp)
}

// userServiceUser используется для адаптации структуры domain.User к тому, что принимает сервис.
// Если структура domain.User отличается от ожидаемой в сервисе, можно сделать маппинг.
type userServiceUser struct {
	Name  string
	Email string
}

func (u *userServiceUser) ToDomain() *struct {
	Name  string
	Email string
} {
	return &struct {
		Name  string
		Email string
	}{
		Name:  u.Name,
		Email: u.Email,
	}
}
