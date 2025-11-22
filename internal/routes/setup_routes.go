package routes

import (
	"net/http"

	"github.com/esteveseverson/go-rest-rbac/internal/handlers"
)

func SetupRoutes(mux *http.ServeMux, handler *handlers.Handler) {
	SetupHealthRoute(mux, handler)
	SetUpUserRoutes(mux, handler)
}
