package votacao

import (
	"context"

	"github.com/aleodoni/voting-go/internal/application/shared"
	"github.com/aleodoni/voting-go/internal/domain/usuario"
	"github.com/aleodoni/voting-go/internal/domain/votacao"
	"github.com/aleodoni/voting-go/internal/platform/event"
)

// CancelaVotacaoInput contém os dados necessários para cancelar uma votação.
type CancelaVotacaoInput struct {
	LoggedInUserKeycloakID string
	ProjetoID              string
}

// CancelaVotacaoPayload é publicado no barramento de eventos quando uma votação é cancelada.
type CancelaVotacaoPayload struct {
	ProjetoID string `json:"projetoId"`
	VotacaoID string `json:"votacaoId"`
}

// CancelaVotacaoUseCase cancela a votação associada a um projeto.
//
// Regras de negócio:
//   - o usuário autenticado deve ser administrador ativo
//   - o projeto deve possuir uma votação associada
//   - a votação deve estar com status fechado ([votacao.StatusVotacaoF])
//
// Ao concluir com sucesso, remove a votação e publica o evento [event.VotacaoCancelada]
// no barramento.
type CancelaVotacaoUseCase struct {
	repoUsuario usuario.UsuarioRepository
	repoReuniao votacao.ReuniaoRepository
	repoVotacao votacao.VotacaoRepository
	bus         *event.Bus
}

// NewCancelaVotacaoUseCase cria uma nova instância de [CancelaVotacaoUseCase].
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

// Execute cancela a votação associada ao projeto informado em [CancelaVotacaoInput.ProjetoID].
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
