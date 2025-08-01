package handler

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/zenmaster911/Game/pkg/model"
)

func (h *Handler) CreateSkill(w http.ResponseWriter, r *http.Request) {
	var input model.Skill

	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
	}

	if err := validate.Struct(input); err != nil {
		sendValidationErrors(w, err)
		return
	}

	err := h.services.Skill.CreateSkill(&input)
	if err != nil {
		http.Error(w, fmt.Sprintf("error occured in cskill creationn: %s", err), http.StatusInternalServerError)
		return
	}
	w.Header().Set("content_type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(input)
}
