package votacao

import (
	"context"

	"github.com/aleodoni/voting-go/internal/application/shared"
	"github.com/aleodoni/voting-go/internal/domain/usuario"
	"github.com/aleodoni/voting-go/internal/domain/votacao"
	"github.com/aleodoni/voting-go/internal/platform/event"
)

type FechaVotacaoInput struct {
	LoggedInUserKeycloakID string
	ProjetoID              string
}

type FechaVotacaoPayload struct {
	ProjetoID string `json:"projetoId"`
	VotacaoID string `json:"votacaoId"`
}

type FechaVotacaoUseCase struct {
	repoUsuario usuario.UsuarioRepository
	repoReuniao votacao.ReuniaoRepository
	repoVotacao votacao.VotacaoRepository
	bus         *event.Bus
}

func NewFechaVotacaoUseCase(
	repoUsuario usuario.UsuarioRepository,
	repoReuniao votacao.ReuniaoRepository,
	repoVotacao votacao.VotacaoRepository,
	bus *event.Bus,
) *FechaVotacaoUseCase {
	return &FechaVotacaoUseCase{
		repoUsuario: repoUsuario,
		repoReuniao: repoReuniao,
		repoVotacao: repoVotacao,
		bus:         bus,
	}
}

func (uc *FechaVotacaoUseCase) Execute(
	ctx context.Context,
	input FechaVotacaoInput,
) error {
	if err := shared.VerificarAdmin(ctx, uc.repoUsuario, input.LoggedInUserKeycloakID); err != nil {
		return err
	}

	projeto, err := uc.repoReuniao.GetProjetoCompleto(ctx, input.ProjetoID)
	if err != nil {
		return err
	}

	if projeto.Votacao == nil {
		return votacao.ErrVotacaoNaoEncontrada
	}

	if projeto.Votacao.Status != votacao.StatusVotacaoA {
		return votacao.ErrVotacaoNaoAberta
	}

	projeto.Votacao.Status = votacao.StatusVotacaoF

	if err := uc.repoVotacao.SalvaVotacao(ctx, projeto.Votacao); err != nil {
		return err
	}

	uc.bus.Publish(event.Event{
		Type: event.VotacaoFechada,
		Payload: FechaVotacaoPayload{
			ProjetoID: projeto.ID,
			VotacaoID: projeto.Votacao.ID,
		},
	})

	return nil
}
