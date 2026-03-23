package votacao

import (
	"context"

	"github.com/aleodoni/go-ddd/domain"
	"github.com/aleodoni/voting-go/internal/application/shared"
	domainUsuario "github.com/aleodoni/voting-go/internal/domain/usuario"
	domainVotacao "github.com/aleodoni/voting-go/internal/domain/votacao"
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
//   - a votação deve estar com status votada ([domainVotacao.StatusVotacaoV])
//
// Ao concluir com sucesso, remove a votação e publica o evento [event.VotacaoCancelada]
// no barramento.
type CancelaVotacaoUseCase struct {
	repoUsuario domainUsuario.UsuarioRepository
	repoReuniao domainVotacao.ReuniaoRepository
	repoVotacao domainVotacao.VotacaoRepository
	bus         *event.Bus
}

// NewCancelaVotacaoUseCase cria uma nova instância de [CancelaVotacaoUseCase].
func NewCancelaVotacaoUseCase(
	repoUsuario domainUsuario.UsuarioRepository,
	repoReuniao domainVotacao.ReuniaoRepository,
	repoVotacao domainVotacao.VotacaoRepository,
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
func (uc *CancelaVotacaoUseCase) Execute(ctx context.Context, input CancelaVotacaoInput) error {
	if err := shared.VerificarAdmin(ctx, uc.repoUsuario, input.LoggedInUserKeycloakID); err != nil {
		return err
	}

	projeto, err := uc.repoReuniao.GetProjetoCompleto(ctx, input.ProjetoID)
	if err != nil {
		return err
	}

	if projeto.Votacao == nil {
		return domainVotacao.ErrVotacaoNaoEncontrada
	}

	if projeto.Votacao.Status != domainVotacao.StatusVotacaoV {
		return domainVotacao.ErrVotacaoNaoFechada
	}

	projeto.Votacao.Cancelar()

	if err := uc.repoVotacao.DeletaVotacao(ctx, projeto.Votacao.ID); err != nil {
		return err
	}

	publishEvents(uc.bus, projeto.Votacao.PullEvents(), func(e domain.DomainEvent) (event.Event, bool) {
		switch evt := e.(type) {
		case domainVotacao.VotacaoCanceladaEvent:
			return event.Event{
				Type:    event.VotacaoCancelada,
				Payload: CancelaVotacaoPayload{ProjetoID: evt.ProjetoID, VotacaoID: evt.VotacaoID},
			}, true
		}
		return event.Event{}, false
	})

	return nil
}
