// internal/handler/http/router.go
package http

import "net/http"

type RouteHandler struct {
	routes map[string]http.Handler
}

func NewRouter() *RouteHandler {
	return &RouteHandler{
		routes: make(map[string]http.Handler),
	}
}

func (rh *RouteHandler) Handle(pattern string, handler http.Handler) {
	rh.routes[pattern] = handler
}

func (rh *RouteHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	for pattern, handler := range rh.routes {
		if pathMatch(r.URL.Path, pattern) {
			handler.ServeHTTP(w, r)
			return
		}
	}
	http.NotFound(w, r)
}

func pathMatch(path, pattern string) bool {
	// Реализация проверки соответствия пути и шаблона
	// Можно добавить поддержку параметров в пути
	return path == pattern
}
