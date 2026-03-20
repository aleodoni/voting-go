// Package router provides routing functionality for the application.
package router

import (
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
	PesquisaUsuarios          *usuarioHandler.PesquisaUsuariosHandler

	RetornaReunioesDia       *reuniaoHandler.RetornaReunioesDiaHandler
	RetornaProjetosCompletos *reuniaoHandler.RetornaProjetosCompletosHandler

	AbreVotacao    *votacaoHandler.AbreVotacaoHandler
	FechaVotacao   *votacaoHandler.FechaVotacaoHandler
	CancelaVotacao *votacaoHandler.CancelaVotacaoHandler
	RegistraVoto   *votacaoHandler.RegistraVotoHandler

	SSE *votacaoHandler.SSEHandler

	GeraRelatorioReuniao *relatorioHandler.GeraRelatorioReuniaoHandler
}

func SetupRouter(jwtMiddleware *middleware.JWTMiddleware, h *Handlers) *gin.Engine {
	r := gin.New()

	r.Use(gin.Logger())
	r.Use(gin.Recovery())

	// Swagger UI — sem autenticação
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	api := r.Group("/api/v1")

	registerHealthRoutes(api)
	registerProtectedRoutes(api, jwtMiddleware, h)
	api.GET("/eventos", h.SSE.Handle)

	return r
}
