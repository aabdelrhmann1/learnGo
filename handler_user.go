package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/aabdelrhmann1/rssagg/internal/auth"
	"github.com/aabdelrhmann1/rssagg/internal/database"
	"github.com/google/uuid"
)

func (apiCfg *apiConfig)handlerCreateUser(w http.ResponseWriter, r *http.Request){
	type parameters struct{
		Name string `json:"name"`
	}
	decoder := json.NewDecoder(r.Body)

	params := parameters{}
	err := decoder.Decode(&params)
	if err != nil {
		respondWithError(w, 400, fmt.Sprintf("Error parsing JSON: %v", err))
		return
	}

	user, err := apiCfg.DB.CreateUser(r.Context(), database.CreateUserParams{
		ID: uuid.New(),
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
		Name: params.Name,
	})

	if err != nil{
		respondWithError(w, 400, fmt.Sprintf("Couldn't create user: %v", err))
		return
	}

	ResponseWithJSON(w, 200, databaseUserToUser(user))

}

func (apiCfg *apiConfig)handlerGetUserByAPIKey(w http.ResponseWriter, r *http.Request){
	apiKey, err := auth.GetAPIKey(r.Header)
	if err != nil{
		respondWithError(w, 403, fmt.Sprintf("Auth error: %v", err))
		return
	}
	user, err := apiCfg.DB.GetUserByAPIKey(r.Context(), apiKey)
	if err != nil{
		respondWithError(w, 400, fmt.Sprintf("Couldn't get a User error: %v", err))
		return
	}

	ResponseWithJSON(w, 201, databaseUserToUser(user))
}