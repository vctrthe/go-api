package main

import (
	"go-api/app"
	"log"

	"github.com/joho/godotenv"
)

func main() {
	// Reading database config file and database init
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	// App Init
	appInstance := app.InitApp()

	// Start Server
	if err := appInstance.Router.Run(); err != nil {
		log.Fatal("Error starting the server:", err)
	}
}
