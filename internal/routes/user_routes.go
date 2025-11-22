package routes

import (
	"net/http"

	"github.com/esteveseverson/go-rest-rbac/internal/handlers"
)

func SetUpUserRoutes(mux *http.ServeMux, handler *handlers.Handler) {
	mux.HandleFunc("POST /user/register", handler.CreateUserHandler())
}
