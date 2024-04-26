package main

import (
	"log"
)

func main () {
	catService := NewGetFactService("https://catfact.ninja/fact")
	logger := NewLoggingService(catService)

	server := NewApiServer(logger)
	err := server.Start(":8080")
	if err != nil {
		log.Fatal(err)
	}
}
