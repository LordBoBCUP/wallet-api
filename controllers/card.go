package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/lordbobcup/wallet-api/config"
	"github.com/lordbobcup/wallet-api/models"
)

func NewCard(w http.ResponseWriter, r *http.Request, config *config.DatabaseConfig) {
	w.Header().Set("content-type", "application/json")
	var card models.Card

	err := json.NewDecoder(r.Body).Decode(&card)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(`{ "message": "` + err.Error() + `" }`))
		return 
	}

	errors := config.DB.Write(config.Collection, card.Id, card);
	if errors != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(`{ "message": "` + errors.Error() + `" }`))
		return 
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(&card)
}

func GetCards(w http.ResponseWriter, r *http.Request, config *config.DatabaseConfig) {
	w.Header().Set("content-type", "application/json")

	result, errors := config.DB.ReadAll(config.Collection)
	if errors != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(`{ "message": "` + errors.Error() + `" }`))
		return 
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(&result)
}

func GetCard(w http.ResponseWriter, r *http.Request, config *config.DatabaseConfig) {
	w.Header().Set("content-type", "application/json")

	params := mux.Vars(r)
	id, _ := params["id"]

	card := models.Card{}

	errors := config.DB.Read(config.Collection, id, &card)
	if errors != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(`{ "message": "` + errors.Error() + `" }`))
		return 
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(&card)
}
