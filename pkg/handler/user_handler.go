package handler

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/zenmaster911/Game/pkg/model"
)

var validate = validator.New()

// type Handler struct {
// 	service *service.UserService
// }

// func NewUserHandler(service *service.UserService) *Handler {
// 	return &Handler{service: service}
// }

func (h *Handler) SignUp(w http.ResponseWriter, r *http.Request) {
	var input model.CreateUser

	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
	}

	if err := validate.Struct(input); err != nil {
		sendValidationErrors(w, err)
		return
	}

	user, err := h.services.Authorization.Create(&input)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(user)
}

func (h *Handler) SignIn(w http.ResponseWriter, r *http.Request) {
	var input model.SignInInput

	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
	}

	if err := validate.Struct(input); err != nil {
		sendValidationErrors(w, err)
		return
	}

	token, err := h.services.GenerateToken(input.Username, input.Password) //добавить обработчик ошибки
	if err != nil {
		http.Error(w, fmt.Sprintf("token generation error: %s", err), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Authorization", "Bearer"+token)
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]interface{}{
		"token": token,
	})

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
