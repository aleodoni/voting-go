// Package bootstrap provides the App struct which holds the application's
// configuration, database connection, and router.
package bootstrap

import (
	"log"

	"github.com/gin-gonic/gin"

	"github.com/aleodoni/voting-go/internal/middleware"
	"github.com/aleodoni/voting-go/internal/platform/event"
	"github.com/aleodoni/voting-go/internal/router"

	"github.com/aleodoni/voting-go/internal/config"
	"github.com/aleodoni/voting-go/internal/database"
)

// App holds the application's top-level dependencies.
type App struct {
	Config *config.Config
	Router *gin.Engine
}

// NewApp initializes and wires the application, connecting to the database,
// building all repositories, use cases, handlers, and the HTTP router.
func NewApp() *App {
	cfg := config.LoadConfig()

	if err := database.RunMigrations(cfg); err != nil {
		log.Fatal(err)
	}

	// if cfg.AppEnv == "staging" || cfg.AppEnv == "production" {
	// 	if err := database.RunFDW(cfg); err != nil {
	// 		log.Println("FDW warning:", err)
	// 	}
	// }

	// Só rodar o FDW em production, para evitar problemas development e staging
	if cfg.AppEnv == "production" {
		if err := database.RunFDW(cfg); err != nil {
			log.Println("FDW warning:", err)
		}
	}

	pgxPool, err := database.ConnectPGX(cfg)
	if err != nil {
		log.Fatal(err)
	}

	if cfg.AppEnv == "development" || cfg.AppEnv == "staging" {
		if err := database.RunSeed(pgxPool); err != nil {
			log.Fatal(err)
		}
	}

	if cfg.AppEnv == "staging" {
		if err := database.RunStagingCron(pgxPool); err != nil {
			log.Println("staging cron warning:", err)
		}
	}

	bus := event.NewBus()
	jwtMiddleware := middleware.NewJWTMiddleware(cfg)
	jobsMiddleware := middleware.NewInternalJobMiddleware(cfg)

	repos := buildRepositories(pgxPool)
	useCases := buildUseCases(repos, bus)
	handlers := buildHandlers(cfg, useCases, repos, bus, jwtMiddleware)

	r := router.SetupRouter(cfg, jwtMiddleware, jobsMiddleware, handlers)

	return &App{
		Config: cfg,
		Router: r,
	}
}
