package handler

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/zenmaster911/Game/pkg/model"
	"github.com/zenmaster911/Game/pkg/service"
)

var validate = validator.New()

type UserHandler struct {
	service *service.UserService
}

func NewUserHandler(service *service.UserService) *UserHandler {
	return &UserHandler{service: service}
}

func (h *UserHandler) SignUp(w http.ResponseWriter, r *http.Request) {
	var input model.CreateUser

	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
	}

	if err := validate.Struct(input); err != nil {
		sendValidationErrors(w, err)
		return
	}

	user, err := h.service.SignUp(&input)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(user)
}

func sendValidationErrors(w http.ResponseWriter, err error) {
	var errors []string
	for _, err := range err.(validator.ValidationErrors) {
		errors = append(errors, fmt.Sprintf("%s is %s", err.Field(), err.Tag()))
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusBadRequest)
	json.NewEncoder(w).Encode((map[string]interface{}{
		"error":  "Validation failed",
		"fields": errors,
	}))
}
