// Package router provides routing functionality for the application.
package router

import (
	"github.com/aleodoni/voting-go/internal/config"
	relatorioHandler "github.com/aleodoni/voting-go/internal/handler/relatorio"
	reuniaoHandler "github.com/aleodoni/voting-go/internal/handler/reuniao"
	usuarioHandler "github.com/aleodoni/voting-go/internal/handler/usuario"
	votacaoHandler "github.com/aleodoni/voting-go/internal/handler/votacao"
	"github.com/aleodoni/voting-go/internal/middleware"
	"github.com/gin-gonic/gin"

	// Swagger
	_ "github.com/aleodoni/voting-go/docs"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

type Handlers struct {
	Me                        *usuarioHandler.MeHandler
	UpdateCredenciais         *usuarioHandler.UpdateCredencialHandler
	UpdateFantasiaCredenciais *usuarioHandler.AtualizaFantasiaCredenciaisHandler
	UpdateFantasia            *usuarioHandler.AtualizaFantasiaHandler
	PesquisaUsuarios          *usuarioHandler.PesquisaUsuariosHandler
	RetornaUsuario            *usuarioHandler.RetornaUsuarioHandler
	ConnectedUsers            *usuarioHandler.ConnectedUsersHandler

	RetornaReunioesDia       *reuniaoHandler.RetornaReunioesDiaHandler
	RetornaProjetosCompletos *reuniaoHandler.RetornaProjetosCompletosHandler
	RetornaProjetoCompleto   *reuniaoHandler.RetornaProjetoCompletoHandler

	AbreVotacao                 *votacaoHandler.AbreVotacaoHandler
	FechaVotacao                *votacaoHandler.FechaVotacaoHandler
	CancelaVotacao              *votacaoHandler.CancelaVotacaoHandler
	RegistraVoto                *votacaoHandler.RegistraVotoHandler
	RetornaProjetoVotacaoAberta *votacaoHandler.RetornaProjetoVotacaoAbertaHandler
	RetornaStatsVotacao         *votacaoHandler.RetornaVotingStatsHandler

	SSE *votacaoHandler.SSEHandler

	GeraRelatorioReuniao *relatorioHandler.GeraRelatorioReuniaoHandler
}

func SetupRouter(cfg *config.Config, jwtMiddleware *middleware.JWTMiddleware, h *Handlers) *gin.Engine {
	r := gin.New()

	r.Use(gin.Logger())
	r.Use(gin.Recovery())
	r.Use(middleware.CORSMiddleware(cfg)) // usa o CORS configurado

	r.OPTIONS("/*any", func(c *gin.Context) {
		c.Status(204)
	})

	// Swagger UI — sem autenticação
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	api := r.Group("/api/v1")
	registerHealthRoutes(api)
	registerProtectedRoutes(api, jwtMiddleware, h)

	// SSE recebe token via query string
	api.GET("/eventos", h.SSE.Handle)

	return r
}
