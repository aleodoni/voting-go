// Package votacao implements the use case
package votacao

import (
	"context"

	"github.com/aleodoni/voting-go/internal/application/shared"
	"github.com/aleodoni/voting-go/internal/domain/usuario"
	"github.com/aleodoni/voting-go/internal/domain/votacao"
)

type RetornaProjetosCompletosInput struct {
	LoggedInUserKeycloakID string
	ReuniaoID              string
}

type RetornaProjetosCompletosUseCase struct {
	repoUsuario usuario.UsuarioRepository
	repoReuniao votacao.ReuniaoRepository
}

func NewRetornaProjetosCompletosUseCase(
	repoUsuario usuario.UsuarioRepository,
	repoReuniao votacao.ReuniaoRepository,
) *RetornaProjetosCompletosUseCase {
	return &RetornaProjetosCompletosUseCase{
		repoUsuario: repoUsuario,
		repoReuniao: repoReuniao,
	}
}

func (uc *RetornaProjetosCompletosUseCase) Execute(
	ctx context.Context,
	input RetornaProjetosCompletosInput,
) ([]*votacao.Projeto, error) {
	// Verificar se o usuário logado é admin
	if err := shared.VerificarAdmin(ctx, uc.repoUsuario, input.LoggedInUserKeycloakID); err != nil {
		return nil, err
	}

	// Verificar se a reunião existe
	if _, err := uc.repoReuniao.FindReuniaoByID(ctx, input.ReuniaoID); err != nil {
		return nil, err
	}

	return uc.repoReuniao.GetProjetosCompleto(ctx, input.ReuniaoID)
}
