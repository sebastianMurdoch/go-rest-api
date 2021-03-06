package main

import (
	"github.com/sebastianMurdoch/go-rest-api/src/api/infrastructure/http"
	"log"
	"os"
)

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	app := http.NewApp()

	if err := app.Run(":"+port); err != nil {
		log.Printf("%s", "Error running server")
	}
}
