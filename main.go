package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/lordbobcup/wallet-api/config"
)

func main() {
	config := config.New()

	fmt.Print(config.Collection);
	
	router := Routes(config) //
	fmt.Println("Starting webserver...")
	//log.Fatal(http.ListenAndServe(":1337", handlers.CORS(allowedHeaders, allowedOrigins, allowedMethods)(router)))
	log.Fatal(http.ListenAndServe(":1337", corsMiddleware(router)))

}