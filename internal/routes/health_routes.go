package routes

import (
	"net/http"

	"github.com/esteveseverson/go-rest-rbac/internal/handlers"
)

func SetupHealthRoute(mux *http.ServeMux, handler *handlers.Handler) {
	mux.HandleFunc("/health", handler.HealthHandler())
}
