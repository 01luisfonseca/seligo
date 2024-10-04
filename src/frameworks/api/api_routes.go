package api

import (
	"net/http"

	"github.com/01luisfonseca/seligo/src/controllers"
)

func RegisterRoutes(mux *http.ServeMux) {
	mux.HandleFunc("/health", controllers.HealthCheck)
}
