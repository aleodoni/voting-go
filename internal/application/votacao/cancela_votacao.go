package votacao

import (
	"context"

	"github.com/aleodoni/voting-go/internal/application/shared"
	"github.com/aleodoni/voting-go/internal/domain/usuario"
	"github.com/aleodoni/voting-go/internal/domain/votacao"
	"github.com/aleodoni/voting-go/internal/platform/event"
)

type CancelaVotacaoInput struct {
	LoggedInUserKeycloakID string
	ProjetoID              string
}

type CancelaVotacaoPayload struct {
	ProjetoID string `json:"projetoId"`
	VotacaoID string `json:"votacaoId"`
}

type CancelaVotacaoUseCase struct {
	repoUsuario usuario.UsuarioRepository
	repoReuniao votacao.ReuniaoRepository
	repoVotacao votacao.VotacaoRepository
	bus         *event.Bus
}

func NewCancelaVotacaoUseCase(
	repoUsuario usuario.UsuarioRepository,
	repoReuniao votacao.ReuniaoRepository,
	repoVotacao votacao.VotacaoRepository,
	bus *event.Bus,
) *CancelaVotacaoUseCase {
	return &CancelaVotacaoUseCase{
		repoUsuario: repoUsuario,
		repoReuniao: repoReuniao,
		repoVotacao: repoVotacao,
		bus:         bus,
	}
}

func (uc *CancelaVotacaoUseCase) Execute(
	ctx context.Context,
	input CancelaVotacaoInput,
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

	if projeto.Votacao.Status != votacao.StatusVotacaoF {
		return votacao.ErrVotacaoNaoFechada
	}

	if err := uc.repoVotacao.DeletaVotacao(ctx, projeto.Votacao.ID); err != nil {
		return err
	}

	uc.bus.Publish(event.Event{
		Type: event.VotacaoCancelada,
		Payload: CancelaVotacaoPayload{
			ProjetoID: projeto.ID,
			VotacaoID: projeto.Votacao.ID,
		},
	})

	return nil
}
