package main

import (
	"context"
	"fmt"
	"log"
)

func main () {
	catService := NewGetFactService("https://catfact.ninja/fact")
	logger := NewLoggingService(catService)

	_, err := logger.GetCatFact(context.TODO())
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Finished")
}