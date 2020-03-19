package main

import (
	"log"
	"os"
	"github.com/lordbobcup/wallet-api/config"
)

func main() {
	config, err := config.New()
	if err != nil {
		log.Fatal("Error during configuration. Application closing.")
		os.Exit(1)
	}
	
	router := Routes(config) //
	fmt.Println("Starting webserver...")
	//log.Fatal(http.ListenAndServe(":1337", handlers.CORS(allowedHeaders, allowedOrigins, allowedMethods)(router)))
	log.Fatal(http.ListenAndServe(":1337", corsMiddleware(router)))

}