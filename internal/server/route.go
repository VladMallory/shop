package http_server

import (
	core_middleware "authTest/internal/middleware"
	"net/http"
)

type Route struct {
	Method     string
	Path       string
	Handler    http.HandlerFunc
	Middleware []core_middleware.Middleware
}

func NewRoute(method string, path string, handler http.HandlerFunc) *Route {
	return &Route{
		Method:  method,
		Path:    path,
		Handler: handler,
	}
}

func (r *Route) WithMiddleware(middleware ...core_middleware.Middleware) http.Handler {
	return core_middleware.ChainMiddleware(r.Handler, r.Middleware...)
}
