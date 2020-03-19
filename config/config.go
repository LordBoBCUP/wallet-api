package config

import (
	"github.com/nanobox-io/golang-scribble"
	"log"
)

type DatabaseConfig struct {
	DB 	*scribble.Driver //Database
	collection string
}

var dbLocation string = "db.json"

func New() (*DatabaseConfig) {
	db := DatabaseConfig{}

	db.DB, err := scribble.New(dbLocation, nil)
	db.collection = "cards"

	if err != nil {
		log.Fatalf("Unable to open Database. error: %v", err)
	}

	return &db
}