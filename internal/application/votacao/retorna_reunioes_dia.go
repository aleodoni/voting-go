// Package votacao contains the use cases related to voting management.
package votacao

import (
	"context"

	"github.com/aleodoni/voting-go/internal/application/shared"
	"github.com/aleodoni/voting-go/internal/domain/usuario"
	"github.com/aleodoni/voting-go/internal/domain/votacao"
)

// RetornaReunioesDiaInput contém os dados necessários para retornar as reuniões do dia.
type RetornaReunioesDiaInput struct {
	LoggedInUserKeycloakID string
}

// RetornaReunioesDiaUseCase retorna a lista de reuniões agendadas para o dia atual.
//
// Regras de negócio:
//   - o usuário autenticado deve ser administrador ativo
type RetornaReunioesDiaUseCase struct {
	repoUsuario usuario.UsuarioRepository
	repoReuniao votacao.ReuniaoRepository
}

// NewRetornaReunioesDiaUseCase cria uma nova instância de [RetornaReunioesDiaUseCase].
func NewRetornaReunioesDiaUseCase(
	repoUsuario usuario.UsuarioRepository,
	repoReuniao votacao.ReuniaoRepository,
) *RetornaReunioesDiaUseCase {
	return &RetornaReunioesDiaUseCase{
		repoUsuario: repoUsuario,
		repoReuniao: repoReuniao,
	}
}

// Execute retorna a lista de reuniões agendadas para o dia atual.
func (uc *RetornaReunioesDiaUseCase) Execute(
	ctx context.Context,
	input RetornaReunioesDiaInput,
) ([]*votacao.Reuniao, error) {
	if err := shared.VerificarAdmin(ctx, uc.repoUsuario, input.LoggedInUserKeycloakID); err != nil {
		return nil, err
	}

	return uc.repoReuniao.GetReunioesDia(ctx)
}
