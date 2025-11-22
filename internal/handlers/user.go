package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/esteveseverson/go-rest-rbac/internal/dtos"
	"github.com/esteveseverson/go-rest-rbac/internal/store"
	"github.com/esteveseverson/go-rest-rbac/internal/utils"
)

func (h *Handler) CreateUserHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()

		var req dtos.CreateUserRequest
		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			utils.RespondWithError(w, http.StatusBadRequest, "Invalid request payload")
		}

		hashedPassword, err := utils.HashPassword(req.Password)
		if err != nil {
			utils.RespondWithError(w, http.StatusInternalServerError, "Error while hashing password")
			return
		}

		_, err = h.Queries.CreateUser(ctx, store.CreateUserParams{
			Username: req.Username,
			Email:    req.Email,
			Password: hashedPassword,
		})

		if err != nil {
			utils.RespondWithError(w, http.StatusInternalServerError, "Error creating user")
			return
		}

		utils.RespondWithSucess(w, http.StatusCreated, "User created with success", req.Username)
	}
}
