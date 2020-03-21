package config

import (
	"log"
	"os"

	scribble "github.com/nanobox-io/golang-scribble"
)

type DatabaseConfig struct {
	DB 	*scribble.Driver //Database
	Collection string
}


func New() (*DatabaseConfig) {
	db := DatabaseConfig{}

	dbLocation, err := os.Getwd() 
	if err != nil {
		log.Fatal("Unable to get current working directory")
	}
	log.Print(dbLocation)
	dbLocation = "C:\\Users\\alex.massey\\go\\src\\github.com\\lordbobcup\\wallet-api\\db"

	log.Print(os.Getwd())
	temp, err := scribble.New(dbLocation, nil)
	db.Collection = "cards"
	db.DB = temp

	if err != nil {
		log.Fatalf("Unable to open Database. error: %v", err)
	}

	return &db
}