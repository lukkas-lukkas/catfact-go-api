package main

import (
	"context"
	"fmt"
	"log"
)

func main () {
	catService := NewGetFactService("https://catfact.ninja/fact")
	logger := NewLoggingService(catService)

	fact, err := logger.GetCatFact(context.TODO())
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("%+v\n", fact)
}