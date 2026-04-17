package main

import (
	"log"

	"github.com/PeterNex14/kioskecil-microservice/user-service/internal/app"
	"github.com/PeterNex14/kioskecil-microservice/user-service/internal/config"
)

func main() {
	// 1. Load configuration
	cfg := config.Load()

	// 2. Initialize the application with config
	application, err := app.New(cfg)
	if err != nil {
		log.Fatalf("failed to initialize application: %v", err)
	}

	// 3. Start the application (blocks until signal)
	application.Run()
}
