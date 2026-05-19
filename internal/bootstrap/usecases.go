package bootstrap

import (
	ucRelatorio "github.com/aleodoni/voting-go/internal/application/relatorio"
	ucUsuario "github.com/aleodoni/voting-go/internal/application/usuario"
	ucVotacao "github.com/aleodoni/voting-go/internal/application/votacao"

	infraRelatorio "github.com/aleodoni/voting-go/internal/infrastructure/report"
	"github.com/aleodoni/voting-go/internal/platform/event"
)

func buildUseCases(r *repositories, bus *event.Bus) *useCases {
	pdfGenerator := infraRelatorio.NewPDFRelatorioReuniaoGenerator()

	return &useCases{
		ensureUsuario:                ucUsuario.NewEnsureUsuarioUseCase(r.usuario, r.transactor),
		updateDisplayName:            ucUsuario.NewUpdateDisplayNameUseCase(r.usuario),
		updateDisplayNamePermissions: ucUsuario.NewUpdateDisplayNamePermissionsUseCase(r.usuario),
		updateCredencial:             ucUsuario.NewUpdateCredencialUseCase(r.usuario),
		listUsuarios:                 ucUsuario.NewListUsuariosUseCase(r.usuario),
		retornaUsuario:               ucUsuario.NewRetornaUsuarioUseCase(r.usuario),
		retornaReunioesDia:           ucVotacao.NewRetornaReunioesDiaUseCase(r.usuario, r.reuniao),
		retornaProjetos:              ucVotacao.NewRetornaProjetosCompletosUseCase(r.usuario, r.reuniao),
		retornaProjeto:               ucVotacao.NewRetornaProjetoCompletoUseCase(r.usuario, r.reuniao),
		abreVotacao:                  ucVotacao.NewAbreVotacaoUseCase(r.usuario, r.reuniao, r.votacao, bus),
		fechaVotacao:                 ucVotacao.NewFechaVotacaoUseCase(r.usuario, r.reuniao, r.votacao, bus),
		cancelaVotacao:               ucVotacao.NewCancelaVotacaoUseCase(r.usuario, r.reuniao, r.votacao, bus),
		registraVoto:                 ucVotacao.NewRegistraVotoUseCase(r.usuario, r.votacao, bus),
		geraRelatorio:                ucRelatorio.NewGeraRelatorioReuniaoUseCase(r.reuniao, pdfGenerator),
		retornaProjetoVotacaoAberta:  ucVotacao.NewRetornaVotacaoAbertaUseCase(r.usuario, r.votacao),
		retornaStatsVotacao:          ucVotacao.NewRetornaVotingStatsUseCase(r.usuario, r.votacao),
	}
}
