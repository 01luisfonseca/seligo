package api

import (
	"log"
	"net/http"
)

type FrameworkAPI struct {
	addr string
}

func NewFrameworkAPI(addr string) *FrameworkAPI {
	return &FrameworkAPI{
		addr: addr,
	}
}

func (f *FrameworkAPI) Prepare() http.Handler {
	router := http.NewServeMux()
	RegisterRoutes(router)
	middleware := NewMiddleware(router)
	return middleware.ApplyMiddleware()
}

func (f *FrameworkAPI) Run() error {
	handler := f.Prepare()
	server := http.Server{
		Addr:    f.addr,
		Handler: handler,
	}
	log.Println("Starting server on", f.addr)
	return server.ListenAndServe()
}
