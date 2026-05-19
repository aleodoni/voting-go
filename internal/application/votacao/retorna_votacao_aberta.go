package votacao

import (
	"context"

	"github.com/aleodoni/voting-go/internal/domain/usuario"
	"github.com/aleodoni/voting-go/internal/domain/votacao"
)

// RetornaVotacaoAbertaInput contém os dados necessários para retornar a votação aberta.
type RetornaVotacaoAbertaInput struct {
	LoggedInUserKeycloakID string
}

// RetornaVotacaoAbertaUseCase retorna o projeto com votação atualmente aberta,
// incluindo seus votos, usuários e pareceres.
//
// Regras de negócio:
//   - o usuário autenticado deve existir
//   - deve existir uma votação com status 'A' (aberta)
type RetornaVotacaoAbertaUseCase struct {
	repoUsuario usuario.UsuarioRepository
	repoVotacao votacao.VotacaoRepository
}

// NewRetornaVotacaoAbertaUseCase cria uma nova instância de [RetornaVotacaoAbertaUseCase].
func NewRetornaVotacaoAbertaUseCase(
	repoUsuario usuario.UsuarioRepository,
	repoVotacao votacao.VotacaoRepository,
) *RetornaVotacaoAbertaUseCase {
	return &RetornaVotacaoAbertaUseCase{
		repoUsuario: repoUsuario,
		repoVotacao: repoVotacao,
	}
}

// Execute retorna o projeto com votação aberta e todos os seus dados relacionados.
func (uc *RetornaVotacaoAbertaUseCase) Execute(
	ctx context.Context,
	input RetornaVotacaoAbertaInput,
) (*votacao.Projeto, error) {
	if _, err := uc.repoUsuario.FindByKeycloakID(ctx, input.LoggedInUserKeycloakID); err != nil {
		return nil, err
	}

	return uc.repoVotacao.GetProjetoVotacaoAberta(ctx)
}
