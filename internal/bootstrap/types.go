package bootstrap

import (
	ucRelatorio "github.com/aleodoni/voting-go/internal/application/relatorio"
	ucSincronia "github.com/aleodoni/voting-go/internal/application/sincronia"
	ucUsuario "github.com/aleodoni/voting-go/internal/application/usuario"
	ucVotacao "github.com/aleodoni/voting-go/internal/application/votacao"

	domainShared "github.com/aleodoni/voting-go/internal/domain/shared"
	domainSincronia "github.com/aleodoni/voting-go/internal/domain/sincronia"
	domainUsuario "github.com/aleodoni/voting-go/internal/domain/usuario"
	domainVotacao "github.com/aleodoni/voting-go/internal/domain/votacao"
)

type repositories struct {
	usuario    domainUsuario.UsuarioRepository
	transactor domainShared.UnitOfWork
	reuniao    domainVotacao.ReuniaoRepository
	votacao    domainVotacao.VotacaoRepository
	sincronia  domainSincronia.SincroniaRepository
}

type useCases struct {
	ensureUsuario                *ucUsuario.EnsureUsuarioUseCase
	updateDisplayNamePermissions *ucUsuario.UpdateDisplayNamePermissionsUseCase
	updateDisplayName            *ucUsuario.UpdateDisplayNameUseCase
	updateCredencial             *ucUsuario.UpdateCredencialUseCase
	listUsuarios                 *ucUsuario.ListUsuariosUseCase
	retornaUsuario               *ucUsuario.RetornaUsuarioUseCase
	retornaReunioesDia           *ucVotacao.RetornaReunioesDiaUseCase
	retornaProjetos              *ucVotacao.RetornaProjetosCompletosUseCase
	retornaProjeto               *ucVotacao.RetornaProjetoCompletoUseCase
	abreVotacao                  *ucVotacao.AbreVotacaoUseCase
	fechaVotacao                 *ucVotacao.FechaVotacaoUseCase
	cancelaVotacao               *ucVotacao.CancelaVotacaoUseCase
	registraVoto                 *ucVotacao.RegistraVotoUseCase
	geraRelatorio                *ucRelatorio.GeraRelatorioReuniaoUseCase
	retornaProjetoVotacaoAberta  *ucVotacao.RetornaVotacaoAbertaUseCase
	retornaStatsVotacao          *ucVotacao.RetornaVotingStatsUseCase
	retornaUltimasSincronias     *ucSincronia.RetornaSincroniasUseCase
	executaSincronia             *ucSincronia.ExecutaSincroniaUseCase
}
