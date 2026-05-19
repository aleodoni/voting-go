// Package votacao contains the use cases related to voting management.
package votacao

import (
	"context"

	"github.com/aleodoni/voting-go/internal/application/shared"
	"github.com/aleodoni/voting-go/internal/domain/usuario"
	"github.com/aleodoni/voting-go/internal/domain/votacao"
)

// RetornaProjetosCompletosInput contém os dados necessários para retornar os projetos completos de uma reunião.
type RetornaProjetosCompletosInput struct {
	LoggedInUserKeycloakID string
	ReuniaoID              string
}

// RetornaProjetosCompletosUseCase retorna a lista completa de projetos de uma reunião,
// incluindo os dados de votação associados a cada projeto.
//
// Regras de negócio:
//   - o usuário autenticado deve ser administrador ativo
//   - a reunião informada em [RetornaProjetosCompletosInput.ReuniaoID] deve existir
type RetornaProjetosCompletosUseCase struct {
	repoUsuario usuario.UsuarioRepository
	repoReuniao votacao.ReuniaoRepository
}

// NewRetornaProjetosCompletosUseCase cria uma nova instância de [RetornaProjetosCompletosUseCase].
func NewRetornaProjetosCompletosUseCase(
	repoUsuario usuario.UsuarioRepository,
	repoReuniao votacao.ReuniaoRepository,
) *RetornaProjetosCompletosUseCase {
	return &RetornaProjetosCompletosUseCase{
		repoUsuario: repoUsuario,
		repoReuniao: repoReuniao,
	}
}

// Execute retorna a lista completa de projetos da reunião informada em
// [RetornaProjetosCompletosInput.ReuniaoID].
func (uc *RetornaProjetosCompletosUseCase) Execute(
	ctx context.Context,
	input RetornaProjetosCompletosInput,
) ([]*votacao.Projeto, error) {
	if err := shared.VerificarAdmin(ctx, uc.repoUsuario, input.LoggedInUserKeycloakID); err != nil {
		return nil, err
	}

	if _, err := uc.repoReuniao.FindReuniaoByID(ctx, input.ReuniaoID); err != nil {
		return nil, err
	}

	return uc.repoReuniao.GetProjetosCompleto(ctx, input.ReuniaoID)
}
