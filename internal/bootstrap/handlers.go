package bootstrap

import (
	relatorioHandler "github.com/aleodoni/voting-go/internal/handler/relatorio"
	reuniaoHandler "github.com/aleodoni/voting-go/internal/handler/reuniao"
	usuarioHandler "github.com/aleodoni/voting-go/internal/handler/usuario"
	votacaoHandler "github.com/aleodoni/voting-go/internal/handler/votacao"
	"github.com/aleodoni/voting-go/internal/middleware"

	"github.com/aleodoni/voting-go/internal/platform/event"
	"github.com/aleodoni/voting-go/internal/router"
)

func buildHandlers(uc *useCases, repos *repositories, bus *event.Bus, jwtMiddleware *middleware.JWTMiddleware) *router.Handlers {
	return &router.Handlers{
		Me:                          usuarioHandler.NewMeHandler(uc.ensureUsuario),
		UpdateCredenciais:           usuarioHandler.NewUpdateCredencialHandler(uc.updateCredencial),
		UpdateFantasiaCredenciais:   usuarioHandler.NewAtualizaFantasiaCredenciaisHandler(uc.updateDisplayNamePermissions),
		UpdateFantasia:              usuarioHandler.NewAtualizaFantasiaHandler(uc.updateDisplayName),
		RetornaReunioesDia:          reuniaoHandler.NewRetornaReunioesDiaHandler(uc.retornaReunioesDia),
		RetornaProjetosCompletos:    reuniaoHandler.NewRetornaProjetosCompletosHandler(uc.retornaProjetos),
		RetornaProjetoCompleto:      reuniaoHandler.NewRetornaProjetoCompletoHandler(uc.retornaProjeto),
		AbreVotacao:                 votacaoHandler.NewAbreVotacaoHandler(uc.abreVotacao),
		FechaVotacao:                votacaoHandler.NewFechaVotacaoHandler(uc.fechaVotacao),
		CancelaVotacao:              votacaoHandler.NewCancelaVotacaoHandler(uc.cancelaVotacao),
		RegistraVoto:                votacaoHandler.NewRegistraVotoHandler(uc.registraVoto),
		PesquisaUsuarios:            usuarioHandler.NewPesquisaUsuariosHandler(uc.listUsuarios),
		RetornaUsuario:              usuarioHandler.NewRetornaUsuarioHandler(uc.retornaUsuario),
		SSE:                         votacaoHandler.NewSSEHandler(bus, jwtMiddleware, repos.usuario),
		GeraRelatorioReuniao:        relatorioHandler.NewGeraRelatorioReuniaoHandler(uc.geraRelatorio),
		RetornaProjetoVotacaoAberta: votacaoHandler.NewRetornaProjetoVotacaoAbertaHandler(uc.retornaProjetoVotacaoAberta),
		RetornaStatsVotacao:         votacaoHandler.NewRetornaVotingStatsHandler(uc.retornaStatsVotacao),
		ConnectedUsers:              usuarioHandler.NewConnectedUsersHandler(bus),
	}
}
