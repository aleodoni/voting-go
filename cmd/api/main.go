// Package main provides the main entry point for the API server.
package main

import (
	"log"

	"github.com/aleodoni/voting-go/internal/bootstrap"
)

func main() {

	app := bootstrap.NewApp()

	// Start server
	log.Printf("🚀 %s running on port %s", app.Config.AppName, app.Config.AppPort)
	err := app.Router.Run(":" + app.Config.AppPort)
	if err != nil {
		log.Fatal("Failed to start server:", err)
	}
}
