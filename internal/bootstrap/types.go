package bootstrap

import (
	ucRelatorio "github.com/aleodoni/voting-go/internal/application/relatorio"
	ucUsuario "github.com/aleodoni/voting-go/internal/application/usuario"
	ucVotacao "github.com/aleodoni/voting-go/internal/application/votacao"

	domainShared "github.com/aleodoni/voting-go/internal/domain/shared"
	domainUsuario "github.com/aleodoni/voting-go/internal/domain/usuario"
	domainVotacao "github.com/aleodoni/voting-go/internal/domain/votacao"
)

type repositories struct {
	usuario    domainUsuario.UsuarioRepository
	transactor domainShared.UnitOfWork
	reuniao    domainVotacao.ReuniaoRepository
	votacao    domainVotacao.VotacaoRepository
}

type useCases struct {
	ensureUsuario               *ucUsuario.EnsureUsuarioUseCase
	updateDisplayName           *ucUsuario.UpdateDisplayNamePermissionsUseCase
	updateCredencial            *ucUsuario.UpdateCredencialUseCase
	listUsuarios                *ucUsuario.ListUsuariosUseCase
	retornaReunioesDia          *ucVotacao.RetornaReunioesDiaUseCase
	retornaProjetos             *ucVotacao.RetornaProjetosCompletosUseCase
	abreVotacao                 *ucVotacao.AbreVotacaoUseCase
	fechaVotacao                *ucVotacao.FechaVotacaoUseCase
	cancelaVotacao              *ucVotacao.CancelaVotacaoUseCase
	registraVoto                *ucVotacao.RegistraVotoUseCase
	geraRelatorio               *ucRelatorio.GeraRelatorioReuniaoUseCase
	retornaProjetoVotacaoAberta *ucVotacao.RetornaVotacaoAbertaUseCase
	retornaStatsVotacao         *ucVotacao.RetornaVotingStatsUseCase
}
