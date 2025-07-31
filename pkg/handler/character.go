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
		return
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

func (h *Handler) DeleteCharByNickname(w http.ResponseWriter, r *http.Request) {
	UserID, err := getUserId(w, r)
	if err != nil {
		http.Error(w, "Extracting userID from context error", http.StatusInternalServerError)
		return
	}
	type deleteNickname struct {
		DeleteNickname string `json:"delete_nickname"`
	}
	var nickname deleteNickname
	if err := json.NewDecoder(r.Body).Decode(&nickname); err != nil {
		http.Error(w, "invalid request body", http.StatusBadRequest)
		return
	}

	if err := h.services.Character.DeleteCharByNickname(UserID, nickname.DeleteNickname); err != nil {
		http.Error(w, fmt.Sprintf("get character destruction failed: %s", err), http.StatusInternalServerError)
		return
	}
	fmt.Println(UserID, nickname.DeleteNickname)
	w.WriteHeader(http.StatusNoContent)

}

func (h *Handler) GetCharById(w http.ResponseWriter, r *http.Request) {

	UserId, err := getUserId(w, r)
	if err != nil {
		http.Error(w, "extracting userID from context error", http.StatusInternalServerError)
		return
	}

	CharID, err := getCharId(w, r)
	if err != nil {
		http.Error(w, fmt.Sprintf("extracting character ID error: %s", err), http.StatusInternalServerError)
	}

	char, err := h.services.Character.GetCharById(UserId, CharID)
	if err != nil {
		http.Error(w, fmt.Sprintf("error in getting char id: %s", err), http.StatusInternalServerError)
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(char)
}
