package votacao

import (
	"context"
	"time"

	"github.com/aleodoni/voting-go/internal/application/shared"
	"github.com/aleodoni/voting-go/internal/domain/usuario"
	"github.com/aleodoni/voting-go/internal/domain/votacao"
)

// RetornaVotingStatsInput contém os dados necessários para retornar as estatísticas de votação.
type RetornaVotingStatsInput struct {
	LoggedInUserKeycloakID string
}

// RetornaVotingStatsUseCase retorna as estatísticas de votação do dia atual.
//
// Regras de negócio:
//   - o usuário autenticado deve ser administrador ativo
type RetornaVotingStatsUseCase struct {
	repoUsuario usuario.UsuarioRepository
	repoVotacao votacao.VotacaoRepository
}

// NewRetornaVotingStatsUseCase cria uma nova instância de [RetornaVotingStatsUseCase].
func NewRetornaVotingStatsUseCase(
	repoUsuario usuario.UsuarioRepository,
	repoVotacao votacao.VotacaoRepository,
) *RetornaVotingStatsUseCase {
	return &RetornaVotingStatsUseCase{
		repoUsuario: repoUsuario,
		repoVotacao: repoVotacao,
	}
}

// Execute retorna as estatísticas de votação do dia atual.
func (uc *RetornaVotingStatsUseCase) Execute(
	ctx context.Context,
	input RetornaVotingStatsInput,
) (*votacao.VotingStats, error) {
	if err := shared.VerificarAdmin(ctx, uc.repoUsuario, input.LoggedInUserKeycloakID); err != nil {
		return nil, err
	}

	return uc.repoVotacao.GetVotingStats(ctx, time.Now())
}
