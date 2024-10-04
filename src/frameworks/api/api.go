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

func (f *FrameworkAPI) Run() error {
	router := http.NewServeMux()
	RegisterRoutes(router)
	middleware := NewMiddleware(router)
	server := http.Server{
		Addr:    f.addr,
		Handler: middleware.ApplyMiddleware(),
	}
	log.Println("Starting server on", f.addr)

	return server.ListenAndServe()
}
