package main

import (
	"fmt"
	"net/http"

	"github.com/ain-py/go-webserver/internal/auth"
	"github.com/ain-py/go-webserver/internal/database"
)

type authHandler func(http.ResponseWriter, *http.Request, database.User)

func (apiCofg *apiConfig) middleWareAuth(handler authHandler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		apikey, err := auth.GetAPIKey(r.Header)
		if err != nil {
			respondWithError(w, 403, fmt.Sprintf("bad auth key", err))
			return
		}
		user, err := apiCofg.DB.GetUserByApiKey(r.Context(), apikey)
		if err != nil {
			respondWithError(w, 403, fmt.Sprintf("cant get user with api key", err))
			return
		}
		handler(w, r, user)
	}
}
