package main

import (
	"fmt"
	"net/http"

	"github.com/renatocardoso243/GoBackendServer/internal/auth"
	"github.com/renatocardoso243/GoBackendServer/internal/database"
)

type authedHandler func(http.ResponseWriter, *http.Request, database.User)

func (apiCfg *apiConfig) middlewareAuth(handler authedHandler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		apiKey, err := auth.GetAPIKey(r.Header)
		if err != nil {
			respondWithError(w, 403, fmt.Sprintf("Auth error: %v", err))
			return
		}

		user, err := apiCfg.DB.GetUserbyAPIKey(r.Context(), apiKey)
		if err != nil {
			respondWithError(w, 400, fmt.Sprintf("User not found: %v", err))
			return
		}

		respondWithJSON(w, 200, databaseUserToUser(user))
		
		handler(w, r, user)
	}
}