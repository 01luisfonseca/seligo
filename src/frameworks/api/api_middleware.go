package api

import (
	"net/http"

	"github.com/01luisfonseca/seligo/src/frameworks/api/middleware"
)

type Middleware struct {
	handler        http.Handler
	MiddlewareList []middleware.MiddlewareFunc
}

func NewMiddleware(h http.Handler) Middleware {
	return Middleware{
		h,
		middleware.AppliedMiddleware(),
	}
}

func (m Middleware) ApplyMiddleware() http.Handler {
	handler := m.handler
	for _, m := range m.MiddlewareList {
		handler = m(handler) // Se envuelve el manejador en el middleware
	}
	return handler
}
