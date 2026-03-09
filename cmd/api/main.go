// Package main provides the main entry point for the API server.
package main

import (
	"log"

	ucValidaUsuario "github.com/aleodoni/voting-go/internal/application/usuario"
	"github.com/aleodoni/voting-go/internal/config"
	"github.com/aleodoni/voting-go/internal/database"
	usuarioHandler "github.com/aleodoni/voting-go/internal/handler/usuario"
	"github.com/aleodoni/voting-go/internal/infrastructure/persistence"
	"github.com/aleodoni/voting-go/internal/middleware"
	"github.com/aleodoni/voting-go/internal/router"
)

func main() {
	// Load configuration
	cfg := config.LoadConfig()

	database.Connect(cfg)

	// Repositórios
	usuarioRepo := persistence.NewUsuarioRepository(database.DB)
	credencialRepo := persistence.NewCredencialRepository(database.DB)
	transactor := persistence.NewGormTransactor(database.DB)

	// Use cases
	validaUsuarioUC := ucValidaUsuario.NewEnsureUsuarioUseCase(usuarioRepo, credencialRepo, transactor)

	// Handlers
	meHandler := usuarioHandler.NewMeHandler(validaUsuarioUC)

	// Create middlewares
	jwtMiddleware := middleware.NewJWTMiddleware(cfg)

	// Router
	r := router.SetupRouter(jwtMiddleware, &router.Handlers{
		Me: meHandler,
	})

	// Start server
	log.Printf("🚀 %s running on port %s", cfg.AppName, cfg.AppPort)
	err := r.Run(":" + cfg.AppPort)
	if err != nil {
		log.Fatal("Failed to start server:", err)
	}
}
