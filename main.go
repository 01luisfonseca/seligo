package main

import (
	"log"

	"github.com/01luisfonseca/seligo/src/frameworks/api"
)

func main() {
	api := api.NewFrameworkAPI(":8080")
	err := api.Run()
	if err != nil {
		log.Fatalf("Error while running the API: %s", err.Error())
	}
}
