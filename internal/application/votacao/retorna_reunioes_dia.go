// Package votacao implements the use case
package votacao

import (
	"context"

	"github.com/aleodoni/voting-go/internal/application/shared"
	"github.com/aleodoni/voting-go/internal/domain/usuario"
	"github.com/aleodoni/voting-go/internal/domain/votacao"
)

type RetornaReunioesDiaInput struct {
	LoggedInUserKeycloakID string
}

type RetornaReunioesDiaUseCase struct {
	repoUsuario usuario.UsuarioRepository
	repoReuniao votacao.ReuniaoRepository
}

func NewRetornaReunioesDiaUseCase(
	repoUsuario usuario.UsuarioRepository,
	repoReuniao votacao.ReuniaoRepository,
) *RetornaReunioesDiaUseCase {
	return &RetornaReunioesDiaUseCase{
		repoUsuario: repoUsuario,
		repoReuniao: repoReuniao,
	}
}

func (uc *RetornaReunioesDiaUseCase) Execute(
	ctx context.Context,
	input RetornaReunioesDiaInput,
) ([]*votacao.Reuniao, error) {
	// Verificar se o usuário logado é admin
	if err := shared.VerificarAdmin(ctx, uc.repoUsuario, input.LoggedInUserKeycloakID); err != nil {
		return nil, err
	}

	return uc.repoReuniao.GetReunioesDia(ctx)
}
