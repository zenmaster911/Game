package handler

import (
	"context"
	"fmt"
	"net/http"
	"strconv"
	"strings"

	"github.com/go-chi/chi/v5"
)

const (
	authHeader = "Authorization"
	userCtx    = "userId"
)

func (h *Handler) userIdentity(next http.Handler) http.Handler {

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		header := r.Header.Get(authHeader)
		if header == "" {
			http.Error(w, "empty auth header", http.StatusUnauthorized)
			return
		}

		headerParts := strings.Split(header, " ")
		if len(headerParts) != 2 /*|| headerParts[0] !="Bearer"*/ {
			http.Error(w, "invalid auth header", http.StatusUnauthorized)
			return
		}

		id, err := h.services.Authorization.ParseToken(headerParts[1])
		if err != nil {
			http.Error(w, fmt.Sprintf("parse token error %s", err), http.StatusUnauthorized)
			return
		}
		userId := id

		ctx := context.WithValue(r.Context(), userCtx, userId)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

func getUserId(w http.ResponseWriter, r *http.Request) (userID int, err error) {
	userIDraw := r.Context().Value(userCtx)
	if userIDraw == nil {
		http.Error(w, "no User ID in context", http.StatusInternalServerError)
		return 0, fmt.Errorf("no user with current ID found")
	}

	userID, ok := userIDraw.(int)
	if !ok {
		http.Error(w, fmt.Sprintf("wrong type of user ID in context %d", userID), http.StatusInternalServerError)
		return 0, fmt.Errorf("wrong type of user ID in context")
	}

	return userID, nil
}

func getCharId(w http.ResponseWriter, r *http.Request) (charID int, err error) {
	param := chi.URLParam(r, "id")

	charID, err = strconv.Atoi(param)
	if err != nil {
		return 0, fmt.Errorf("error in geting character id: %s", err)
	}
	return charID, nil
}
