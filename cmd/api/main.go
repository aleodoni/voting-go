// Package main provides the main entry point for the API server.
package main

import (
	"log"

	ucCredencial "github.com/aleodoni/voting-go/internal/application/credencial"
	ucValidaUsuario "github.com/aleodoni/voting-go/internal/application/usuario"
	"github.com/aleodoni/voting-go/internal/config"
	"github.com/aleodoni/voting-go/internal/database"
	credencialHandler "github.com/aleodoni/voting-go/internal/handler/credencial"
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
	updateCredencialUC := ucCredencial.NewUpdateCredencialUseCase(usuarioRepo, credencialRepo)

	// Handlers
	meHandler := usuarioHandler.NewMeHandler(validaUsuarioUC)
	updateCredencialHandler := credencialHandler.NewUpdateCredencialHandler(updateCredencialUC)

	// Create middlewares
	jwtMiddleware := middleware.NewJWTMiddleware(cfg)

	// Router
	r := router.SetupRouter(jwtMiddleware, &router.Handlers{
		Me:                meHandler,
		UpdateCredenciais: updateCredencialHandler,
	})

	// Start server
	log.Printf("🚀 %s running on port %s", cfg.AppName, cfg.AppPort)
	err := r.Run(":" + cfg.AppPort)
	if err != nil {
		log.Fatal("Failed to start server:", err)
	}
}
