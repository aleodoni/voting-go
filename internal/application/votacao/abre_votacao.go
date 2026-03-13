// Package votacao implements the use case
package votacao

import (
	"context"

	"github.com/aleodoni/voting-go/internal/application/shared"
	"github.com/aleodoni/voting-go/internal/domain/usuario"
	"github.com/aleodoni/voting-go/internal/domain/votacao"
	"github.com/aleodoni/voting-go/internal/platform/id"
)

type AbreVotacaoInput struct {
	LoggedInUserKeycloakID string
	ProjetoID              string
}

type AbreVotacaoUseCase struct {
	repoUsuario usuario.UsuarioRepository
	repoReuniao votacao.ReuniaoRepository
}

func NewAbreVotacaoUseCase(
	repoUsuario usuario.UsuarioRepository,
	repoReuniao votacao.ReuniaoRepository,
) *AbreVotacaoUseCase {
	return &AbreVotacaoUseCase{
		repoUsuario: repoUsuario,
		repoReuniao: repoReuniao,
	}
}

func (uc *AbreVotacaoUseCase) Execute(
	ctx context.Context,
	input AbreVotacaoInput,
) error {
	// Verificar se o usuário logado é admin
	if err := shared.VerificarAdmin(ctx, uc.repoUsuario, input.LoggedInUserKeycloakID); err != nil {
		return err
	}

	// Busca o projeto
	projeto, err := uc.repoReuniao.GetProjetoCompleto(ctx, input.ProjetoID)
	if err != nil {
		return err
	}

	// Cria a votação
	votacaoNova := &votacao.Votacao{
		ID:        id.New(),
		ProjetoID: &projeto.ID,
		Status:    votacao.StatusVotacaoA,
	}

	if err := uc.repoReuniao.CriaVotacao(ctx, votacaoNova); err != nil {
		return err
	}

	return nil
}
