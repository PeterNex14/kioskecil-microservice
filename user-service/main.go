package main

import (
	"log"

	"github.com/PeterNex14/kioskecil-microservice/user-service/internal/app"
)

func main() {
	// Initialize the application
	application, err := app.New()
	if err != nil {
		log.Fatal(err)
	}

	// Start the application (this blocks until an exit signal is received)
	application.Run()
}
