package votacao

import (
	"context"

	"github.com/aleodoni/voting-go/internal/application/shared"
	"github.com/aleodoni/voting-go/internal/domain/usuario"
	"github.com/aleodoni/voting-go/internal/domain/votacao"
	"github.com/aleodoni/voting-go/internal/platform/event"
	"github.com/aleodoni/voting-go/internal/platform/id"
)

type RegistraVotoInput struct {
	LoggedInUserKeycloakID string
	VotacaoID              string
	Voto                   votacao.OpcaoVoto
	Restricao              *votacao.Restricao
	VotoContrario          *votacao.VotoContrario
}

type RegistraVotoPayload struct {
	VotacaoID string `json:"votacaoId"`
}

type RegistraVotoUseCase struct {
	repoUsuario usuario.UsuarioRepository
	repoVotacao votacao.VotacaoRepository
	bus         *event.Bus
}

func NewRegistraVotoUseCase(
	repoUsuario usuario.UsuarioRepository,
	repoVotacao votacao.VotacaoRepository,
	bus *event.Bus,
) *RegistraVotoUseCase {
	return &RegistraVotoUseCase{
		repoUsuario: repoUsuario,
		repoVotacao: repoVotacao,
		bus:         bus,
	}
}

func (uc *RegistraVotoUseCase) Execute(ctx context.Context, input RegistraVotoInput) error {
	u, err := shared.VerificarVota(ctx, uc.repoUsuario, input.LoggedInUserKeycloakID)
	if err != nil {
		return err
	}

	// 1. Votação existe e está aberta?
	v, err := uc.repoVotacao.BuscaVotacao(ctx, input.VotacaoID)
	if err != nil {
		return err
	}
	if v.Status != votacao.StatusVotacaoA {
		return votacao.ErrVotacaoNaoAberta
	}

	// 2. Usuário já votou?
	jaVotou, err := uc.repoVotacao.UsuarioJaVotou(ctx, u.ID, input.VotacaoID)
	if err != nil {
		return err
	}
	if jaVotou {
		return votacao.ErrUsuarioJaVotou
	}

	voto := &votacao.Voto{
		ID:            id.New(),
		VotacaoID:     input.VotacaoID,
		UsuarioID:     u.ID,
		Voto:          input.Voto,
		Restricao:     input.Restricao,
		VotoContrario: input.VotoContrario,
	}

	if voto.Restricao != nil {
		voto.Restricao.ID = id.New()
	}
	if voto.VotoContrario != nil {
		voto.VotoContrario.ID = id.New()
	}

	if err := uc.repoVotacao.SalvaVoto(ctx, voto); err != nil {
		return err
	}

	uc.bus.Publish(event.Event{
		Type:    event.VotoRegistrado,
		Payload: RegistraVotoPayload{VotacaoID: input.VotacaoID},
	})

	return nil
}
