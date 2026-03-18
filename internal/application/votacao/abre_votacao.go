package votacao

import (
	"context"

	"github.com/aleodoni/voting-go/internal/application/shared"
	"github.com/aleodoni/voting-go/internal/domain/usuario"
	"github.com/aleodoni/voting-go/internal/domain/votacao"
	"github.com/aleodoni/voting-go/internal/platform/event"
	"github.com/aleodoni/voting-go/internal/platform/id"
)

type AbreVotacaoInput struct {
	LoggedInUserKeycloakID string
	ProjetoID              string
}

type AbreVotacaoPayload struct {
	ProjetoID string `json:"projetoId"`
	VotacaoID string `json:"votacaoId"`
}

type AbreVotacaoUseCase struct {
	repoUsuario usuario.UsuarioRepository
	repoReuniao votacao.ReuniaoRepository
	repoVotacao votacao.VotacaoRepository
	bus         *event.Bus
}

func NewAbreVotacaoUseCase(
	repoUsuario usuario.UsuarioRepository,
	repoReuniao votacao.ReuniaoRepository,
	repoVotacao votacao.VotacaoRepository,
	bus *event.Bus,
) *AbreVotacaoUseCase {
	return &AbreVotacaoUseCase{
		repoUsuario: repoUsuario,
		repoReuniao: repoReuniao,
		repoVotacao: repoVotacao,
		bus:         bus,
	}
}

func (uc *AbreVotacaoUseCase) Execute(
	ctx context.Context,
	input AbreVotacaoInput,
) error {
	if err := shared.VerificarAdmin(ctx, uc.repoUsuario, input.LoggedInUserKeycloakID); err != nil {
		return err
	}

	// Verifica se já existe votação aberta
	votacaoExistente, err := uc.repoVotacao.GetVotacaoAberta(ctx)
	if err != nil {
		return err
	}
	if votacaoExistente != nil {
		return votacao.ErrVotacaoAberta
	}

	projeto, err := uc.repoReuniao.GetProjetoCompleto(ctx, input.ProjetoID)
	if err != nil {
		return err
	}

	votacaoNova := &votacao.Votacao{
		ID:        id.New(),
		ProjetoID: &projeto.ID,
		Status:    votacao.StatusVotacaoA,
	}

	if err := uc.repoVotacao.SalvaVotacao(ctx, votacaoNova); err != nil {
		return err
	}

	uc.bus.Publish(event.Event{
		Type: event.VotacaoAberta,
		Payload: AbreVotacaoPayload{
			ProjetoID: projeto.ID,
			VotacaoID: votacaoNova.ID,
		},
	})

	return nil
}
