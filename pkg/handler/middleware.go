package handler

import (
	"context"
	"fmt"
	"net/http"
	"strconv"
	"strings"
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
		userId := strconv.Itoa(id)

		ctx := context.WithValue(r.Context(), userCtx, userId)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
