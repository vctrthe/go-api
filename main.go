package main

import (
	"go-api/app"
	"log"
)

func main() {
	// App Init
	appInstance := app.InitApp()

	// Start Server
	if err := appInstance.Router.Run(); err != nil {
		log.Fatal("Error starting the server:", err)
	}
}
