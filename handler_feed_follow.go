package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/ain-py/go-webserver/internal/database"
	"github.com/google/uuid"
)

func (apiCofg *apiConfig) handlerCreateFeedFollow(w http.ResponseWriter, r *http.Request, user database.User) {
	type parameters struct {
		FeedId uuid.UUID `json:"feed_id"`
	}
	decoder := json.NewDecoder(r.Body)
	params := parameters{}

	err := decoder.Decode(&params)
	if err != nil {
		respondWithError(w, 400, fmt.Sprintf("Error parsing Json", err))
		return
	}

	feedfollow, err := apiCofg.DB.CreateFeedFollow(r.Context(), database.CreateFeedFollowParams{
		ID:        uuid.New(),
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
		FeedID:    params.FeedId,
		UserID:    user.ID,
	})
	if err != nil {
		respondWithError(w, 400, fmt.Sprintf("Couldn't create follow feed:", err))
		return
	}
	respondWithJSON(w, 201, databaseFeedFollowToFeedFollow(feedfollow))
}
