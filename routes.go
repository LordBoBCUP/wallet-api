package main

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/lordbobcup/wallet-api/config"
)

func Routes(config *config.Config) *mux.Router {
	router := mux.NewRouter().StrictSlash(true)
	// Open Authentication Routes
	router.HandleFunc("/register", func(w http.ResponseWriter, r *http.Request) {controllers.CreateUser(w,r,config)}).Methods("POST")
	router.HandleFunc("/login", func(w http.ResponseWriter, r *http.Request) {controllers.Login(w,r,config)}).Methods("POST")

	return router
}
