// Package bootstrap provides the App struct which holds the application's
// configuration, database connection, and router.
package bootstrap

import (
	"log"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"

	ucRelatorio "github.com/aleodoni/voting-go/internal/application/relatorio"
	ucUsuario "github.com/aleodoni/voting-go/internal/application/usuario"
	ucVotacao "github.com/aleodoni/voting-go/internal/application/votacao"
	domainUsuario "github.com/aleodoni/voting-go/internal/domain/usuario"
	domainVotacao "github.com/aleodoni/voting-go/internal/domain/votacao"
	"github.com/aleodoni/voting-go/internal/middleware"
	"github.com/aleodoni/voting-go/internal/platform/event"
	"github.com/aleodoni/voting-go/internal/router"

	"github.com/aleodoni/voting-go/internal/config"
	"github.com/aleodoni/voting-go/internal/database"
	persistence "github.com/aleodoni/voting-go/internal/infrastructure/persistence"
	infraRelatorio "github.com/aleodoni/voting-go/internal/infrastructure/report"

	relatorioHandler "github.com/aleodoni/voting-go/internal/handler/relatorio"
	reuniaoHandler "github.com/aleodoni/voting-go/internal/handler/reuniao"
	usuarioHandler "github.com/aleodoni/voting-go/internal/handler/usuario"
	votacaoHandler "github.com/aleodoni/voting-go/internal/handler/votacao"
)

// App holds the application's top-level dependencies.
type App struct {
	Config *config.Config
	DB     *gorm.DB
	Router *gin.Engine
}

// NewApp initializes and wires the application, connecting to the database,
// building all repositories, use cases, handlers, and the HTTP router.
func NewApp() *App {
	cfg := config.LoadConfig()

	db, err := database.Connect(cfg)
	if err != nil {
		log.Fatal(err)
	}

	bus := event.NewBus()

	repos := buildRepositories(db)
	useCases := buildUseCases(repos, bus)
	handlers := buildHandlers(useCases, repos, bus)

	r := router.SetupRouter(middleware.NewJWTMiddleware(cfg), handlers)

	return &App{
		Config: cfg,
		DB:     db,
		Router: r,
	}
}

// repositories groups all repository instances.
type repositories struct {
	usuario    domainUsuario.UsuarioRepository
	credencial domainUsuario.CredencialRepository
	transactor *persistence.GormTransactor
	reuniao    domainVotacao.ReuniaoRepository
	votacao    domainVotacao.VotacaoRepository
}

// buildRepositories creates all repository instances from the given database connection.
func buildRepositories(db *gorm.DB) *repositories {
	return &repositories{
		usuario:    persistence.NewUsuarioRepository(db),
		credencial: persistence.NewCredencialRepository(db),
		transactor: persistence.NewGormTransactor(db),
		reuniao:    persistence.NewReuniaoRepository(db),
		votacao:    persistence.NewVotacaoRepository(db),
	}
}

// useCases groups all use case instances.
type useCases struct {
	ensureUsuario      *ucUsuario.EnsureUsuarioUseCase
	updateDisplayName  *ucUsuario.UpdateDisplayNamePermissionsUseCase
	updateCredencial   *ucUsuario.UpdateCredencialUseCase
	listUsuarios       *ucUsuario.ListUsuariosUseCase
	retornaReunioesDia *ucVotacao.RetornaReunioesDiaUseCase
	retornaProjetos    *ucVotacao.RetornaProjetosCompletosUseCase
	abreVotacao        *ucVotacao.AbreVotacaoUseCase
	fechaVotacao       *ucVotacao.FechaVotacaoUseCase
	cancelaVotacao     *ucVotacao.CancelaVotacaoUseCase
	registraVoto       *ucVotacao.RegistraVotoUseCase
	geraRelatorio      *ucRelatorio.GeraRelatorioReuniaoUseCase
}

// buildUseCases creates all use case instances from the given repositories and event bus.
func buildUseCases(r *repositories, bus *event.Bus) *useCases {
	pdfGenerator := infraRelatorio.NewPDFRelatorioReuniaoGenerator()

	return &useCases{
		ensureUsuario:      ucUsuario.NewEnsureUsuarioUseCase(r.usuario, r.credencial, r.transactor),
		updateDisplayName:  ucUsuario.NewUpdateDisplayNamePermissionsUseCase(r.usuario),
		updateCredencial:   ucUsuario.NewUpdateCredencialUseCase(r.usuario, r.credencial),
		listUsuarios:       ucUsuario.NewListUsuariosUseCase(r.usuario),
		retornaReunioesDia: ucVotacao.NewRetornaReunioesDiaUseCase(r.usuario, r.reuniao),
		retornaProjetos:    ucVotacao.NewRetornaProjetosCompletosUseCase(r.usuario, r.reuniao),
		abreVotacao:        ucVotacao.NewAbreVotacaoUseCase(r.usuario, r.reuniao, r.votacao, bus),
		fechaVotacao:       ucVotacao.NewFechaVotacaoUseCase(r.usuario, r.reuniao, r.votacao, bus),
		cancelaVotacao:     ucVotacao.NewCancelaVotacaoUseCase(r.usuario, r.reuniao, r.votacao, bus),
		registraVoto:       ucVotacao.NewRegistraVotoUseCase(r.usuario, r.votacao, bus),
		geraRelatorio:      ucRelatorio.NewGeraRelatorioReuniaoUseCase(r.reuniao, pdfGenerator),
	}
}

// buildHandlers creates all HTTP handler instances and returns the [router.Handlers]
// struct used to configure the application routes.
func buildHandlers(uc *useCases, repos *repositories, bus *event.Bus) *router.Handlers {
	return &router.Handlers{
		Me:                        usuarioHandler.NewMeHandler(uc.ensureUsuario),
		UpdateCredenciais:         usuarioHandler.NewUpdateCredencialHandler(uc.updateCredencial),
		UpdateFantasiaCredenciais: usuarioHandler.NewAtualizaFantasiaCredenciaisHandler(uc.updateDisplayName),
		RetornaReunioesDia:        reuniaoHandler.NewRetornaReunioesDiaHandler(uc.retornaReunioesDia),
		RetornaProjetosCompletos:  reuniaoHandler.NewRetornaProjetosCompletosHandler(uc.retornaProjetos),
		AbreVotacao:               votacaoHandler.NewAbreVotacaoHandler(uc.abreVotacao),
		FechaVotacao:              votacaoHandler.NewFechaVotacaoHandler(uc.fechaVotacao),
		CancelaVotacao:            votacaoHandler.NewCancelaVotacaoHandler(uc.cancelaVotacao),
		RegistraVoto:              votacaoHandler.NewRegistraVotoHandler(uc.registraVoto),
		PesquisaUsuarios:          usuarioHandler.NewPesquisaUsuariosHandler(uc.listUsuarios),
		SSE:                       votacaoHandler.NewSSEHandler(bus, repos.usuario),
		GeraRelatorioReuniao:      relatorioHandler.NewGeraRelatorioReuniaoHandler(uc.geraRelatorio),
	}
}
