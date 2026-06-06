package http_server

import (
	core_middleware "authTest/internal/middleware"
	"fmt"
	"net/http"
)

type APIVersion string

const (
	APIVersionV1 = "v1"
)

type APIVersionRouter struct {
	*http.ServeMux
	apiVersion  APIVersion
	middlewares []core_middleware.Middleware
}

func NewAPIVersionRouter(apiVersion APIVersion, middleware ...core_middleware.Middleware) *APIVersionRouter {
	return &APIVersionRouter{
		ServeMux:    http.NewServeMux(),
		apiVersion:  apiVersion,
		middlewares: middleware,
	}
}

func (r *APIVersionRouter) RegisterRoutes(routes ...Route) {
	for _, route := range routes {
		pattern := fmt.Sprintf("%s %s", route.Method, route.Path)

		h := route.WithMiddleware()

		r.Handle(pattern, h)
	}
}

func (r *APIVersionRouter) WithMiddleware() http.Handler {
	return core_middleware.ChainMiddleware(r, r.middlewares...)
}
