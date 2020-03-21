package main

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/lordbobcup/wallet-api/config"
	"github.com/lordbobcup/wallet-api/controllers"
)

func Routes(config *config.DatabaseConfig) *mux.Router {
	router := mux.NewRouter().StrictSlash(true)
	// Open Authentication Routes
	router.HandleFunc("/cards", func(w http.ResponseWriter, r *http.Request) {controllers.NewCard(w,r,config)}).Methods("POST")
	router.HandleFunc("/cards", func(w http.ResponseWriter, r *http.Request) {controllers.GetCards(w,r,config)}).Methods("GET")
	router.HandleFunc("/card/{id}", func(w http.ResponseWriter, r *http.Request) {controllers.GetCard(w,r,config)}).Methods("GET")
	router.HandleFunc("/card/{id}", func(w http.ResponseWriter, r *http.Request) {controllers.RemoveCard(w,r,config)}).Methods("DELETE")
	
	return router
}
