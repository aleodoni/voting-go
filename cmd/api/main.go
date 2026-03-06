// Package main provides the main entry point for the API server.
package main

import (
	"log"

	"github.com/aleodoni/voting-go/internal/config"
	"github.com/aleodoni/voting-go/internal/database"
	"github.com/aleodoni/voting-go/internal/router"
)

func main() {
	// Load configuration
	cfg := config.LoadConfig()

	database.Connect(cfg)

	// Setup router
	r := router.SetupRouter()

	// Start server
	log.Printf("🚀 %s running on port %s", cfg.AppName, cfg.AppPort)
	err := r.Run(":" + cfg.AppPort)
	if err != nil {
		log.Fatal("Failed to start server:", err)
	}
}
