// Package bootstrap provides the App struct which holds the application's configuration, database connection, and router.
package bootstrap

import (
	"log"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"

	ucCredencial "github.com/aleodoni/voting-go/internal/application/credencial"
	ucUsuario "github.com/aleodoni/voting-go/internal/application/usuario"
	"github.com/aleodoni/voting-go/internal/middleware"
	"github.com/aleodoni/voting-go/internal/router"

	"github.com/aleodoni/voting-go/internal/config"
	"github.com/aleodoni/voting-go/internal/database"
	"github.com/aleodoni/voting-go/internal/infrastructure/persistence"

	credencialHandler "github.com/aleodoni/voting-go/internal/handler/credencial"
	usuarioHandler "github.com/aleodoni/voting-go/internal/handler/usuario"
)

type App struct {
	Config *config.Config
	DB     *gorm.DB
	Router *gin.Engine
}

func NewApp() *App {
	cfg := config.LoadConfig()

	db, err := database.Connect(cfg)
	if err != nil {
		log.Fatal(err)
	}

	// repositories
	usuarioRepo := persistence.NewUsuarioRepository(db)
	credencialRepo := persistence.NewCredencialRepository(db)
	transactor := persistence.NewGormTransactor(db)

	// use cases
	validaUsuarioUC := ucUsuario.NewEnsureUsuarioUseCase(
		usuarioRepo,
		credencialRepo,
		transactor,
	)

	atualizaFantasiaCredencialUC := ucUsuario.NewUpdateDisplayNamePermissionsUseCase(
		usuarioRepo,
	)

	updateCredencialUC := ucCredencial.NewUpdateCredencialUseCase(
		usuarioRepo,
		credencialRepo,
	)

	// handlers
	meHandler := usuarioHandler.NewMeHandler(validaUsuarioUC)

	updateCredencialHandler := credencialHandler.NewUpdateCredencialHandler(
		updateCredencialUC,
	)

	updateFantasiaCredencialHandler := usuarioHandler.NewAtualizaFantasiaCredenciaisHandler(
		atualizaFantasiaCredencialUC,
	)

	// middleware
	jwtMiddleware := middleware.NewJWTMiddleware(cfg)

	// router
	r := router.SetupRouter(jwtMiddleware, &router.Handlers{
		Me:                        meHandler,
		UpdateCredenciais:         updateCredencialHandler,
		UpdateFantasiaCredenciais: updateFantasiaCredencialHandler,
	})

	return &App{
		Config: cfg,
		DB:     db,
		Router: r,
	}

}
