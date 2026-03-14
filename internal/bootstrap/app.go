// Package bootstrap provides the App struct which holds the application's configuration, database connection, and router.
package bootstrap

import (
	"log"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"

	ucUsuario "github.com/aleodoni/voting-go/internal/application/usuario"
	ucVotacao "github.com/aleodoni/voting-go/internal/application/votacao"
	"github.com/aleodoni/voting-go/internal/middleware"
	"github.com/aleodoni/voting-go/internal/platform/event"
	"github.com/aleodoni/voting-go/internal/router"

	"github.com/aleodoni/voting-go/internal/config"
	"github.com/aleodoni/voting-go/internal/database"
	"github.com/aleodoni/voting-go/internal/infrastructure/persistence"

	reuniaoHandler "github.com/aleodoni/voting-go/internal/handler/reuniao"
	usuarioHandler "github.com/aleodoni/voting-go/internal/handler/usuario"
	votacaoHandler "github.com/aleodoni/voting-go/internal/handler/votacao"
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

	// event bus
	bus := event.NewBus()

	// repositories
	usuarioRepo := persistence.NewUsuarioRepository(db)
	credencialRepo := persistence.NewCredencialRepository(db)
	transactor := persistence.NewGormTransactor(db)
	reuniaoRepo := persistence.NewReuniaoRepository(db)
	votacaoRepo := persistence.NewVotacaoRepository(db)

	// use cases
	validaUsuarioUC := ucUsuario.NewEnsureUsuarioUseCase(
		usuarioRepo,
		credencialRepo,
		transactor,
	)

	atualizaFantasiaCredencialUC := ucUsuario.NewUpdateDisplayNamePermissionsUseCase(
		usuarioRepo,
	)

	updateCredencialUC := ucUsuario.NewUpdateCredencialUseCase(
		usuarioRepo,
		credencialRepo,
	)

	ucRetornaReunioesDiaUC := ucVotacao.NewRetornaReunioesDiaUseCase(
		usuarioRepo,
		reuniaoRepo,
	)

	ucRetornaProjetosCompletosUC := ucVotacao.NewRetornaProjetosCompletosUseCase(
		usuarioRepo,
		reuniaoRepo,
	)

	ucAbreVotacaoUC := ucVotacao.NewAbreVotacaoUseCase(usuarioRepo, reuniaoRepo, votacaoRepo, bus)
	ucFechaVotacaoUC := ucVotacao.NewFechaVotacaoUseCase(usuarioRepo, reuniaoRepo, votacaoRepo, bus)
	ucCancelaVotacaoUC := ucVotacao.NewCancelaVotacaoUseCase(usuarioRepo, reuniaoRepo, votacaoRepo, bus)

	// handlers
	meHandler := usuarioHandler.NewMeHandler(validaUsuarioUC)

	updateCredencialHandler := usuarioHandler.NewUpdateCredencialHandler(
		updateCredencialUC,
	)

	updateFantasiaCredencialHandler := usuarioHandler.NewAtualizaFantasiaCredenciaisHandler(
		atualizaFantasiaCredencialUC,
	)

	retornaReunioesDiaHandler := reuniaoHandler.NewRetornaReunioesDiaHandler(
		ucRetornaReunioesDiaUC,
	)

	retornaProjetosCompletosHandler := reuniaoHandler.NewRetornaProjetosCompletosHandler(
		ucRetornaProjetosCompletosUC,
	)

	abreVotacaoHandler := votacaoHandler.NewAbreVotacaoHandler(ucAbreVotacaoUC)
	fechaVotacaoHandler := votacaoHandler.NewFechaVotacaoHandler(ucFechaVotacaoUC)
	cancelaVotacaoHandler := votacaoHandler.NewCancelaVotacaoHandler(ucCancelaVotacaoUC)

	sseHandler := votacaoHandler.NewSSEHandler(bus)

	// middleware
	jwtMiddleware := middleware.NewJWTMiddleware(cfg)

	// router
	r := router.SetupRouter(jwtMiddleware, &router.Handlers{
		Me:                        meHandler,
		UpdateCredenciais:         updateCredencialHandler,
		UpdateFantasiaCredenciais: updateFantasiaCredencialHandler,
		RetornaReunioesDia:        retornaReunioesDiaHandler,
		RetornaProjetosCompletos:  retornaProjetosCompletosHandler,
		AbreVotacao:               abreVotacaoHandler,
		FechaVotacao:              fechaVotacaoHandler,
		CancelaVotacao:            cancelaVotacaoHandler,
		SSE:                       sseHandler,
	})

	return &App{
		Config: cfg,
		DB:     db,
		Router: r,
	}
}
