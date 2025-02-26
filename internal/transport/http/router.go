package http

import (
	"github.com/julienschmidt/httprouter"
)

func NewRouter(handler *Handler) *httprouter.Router {
	router := httprouter.New()

	router.POST("/users", handler.CreateUser)
	router.GET("/users/:id", handler.GetUser)
	// Добавить другие роуты по необходимости

	return router
}
