//	@title			Voting API
//	@version		1.0.0
//	@description	API de gerenciamento de votações — reuniões, projetos e votos.

//	@host		localhost:8080
//	@BasePath	/api/v1

//	@securityDefinitions.apikey	BearerAuth
//	@in							header
//	@name						Authorization
//	@description				Token JWT emitido pelo Keycloak. Formato: Bearer <token>

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
