package handler

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/zenmaster911/Game/pkg/model"
)

func (h *Handler) createChar(w http.ResponseWriter, r *http.Request) {
	userID, err := getUserId(w, r)
	if err != nil {
		http.Error(w, "Extracting userID from context error", http.StatusInternalServerError)
		return
	}
	var input model.Character
	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
	}

	if err := validate.Struct(input); err != nil {
		sendValidationErrors(w, err)
		return
	}
	//fmt.Println(userID)
	//fmt.Println(&input)
	id, err := h.services.Character.CreateChar(userID, &input)
	if err != nil {
		http.Error(w, fmt.Sprintf("Error in Create character: %s", err), http.StatusInternalServerError)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(map[string]interface{}{
		"id": id,
	})
}

func (h *Handler) UserChars(w http.ResponseWriter, r *http.Request) {
	userID, err := getUserId(w, r)

	if err != nil {
		http.Error(w, "Extracting userID from context error", http.StatusInternalServerError)
		return
	}

	chars, err := h.services.UserChars(userID)
	if err != nil {
		http.Error(w, fmt.Sprintf("getting users chars list problem: %s", err), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(chars)
}
