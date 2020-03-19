package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/lordbobcup/wallet-api/config"
	"github.com/lordbobcup/wallet-api/models"
	"github.com/lordbobcup/wallet-api/utils"
)

func NewCalendar(w http.ResponseWriter, r *http.Request, config *config.Config) {
	w.Header().Set("content-type", "application/json")
	var card models.Card

	//errors := config.DB.Debug().Where("id = ?", id).First(&card).GetErrors()
	errors := config.DB.Write(config.DB.collection, card.ID, card);
	utils.CheckError(&w, &errors)

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(&card)
}

func GetCards(w http.ResponseWriter, r *http.Request, config *config.Config) {
	w.Header().Set("content-type", "application/json")

	//errors := config.DB.Debug().Preload("Appointments").Where("id = ?", id).First(&cal).GetErrors()
	result, errors := config.DB.Read(config.DB.Collection)
	utils.CheckError(&w, &errors)

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(&result)
}

