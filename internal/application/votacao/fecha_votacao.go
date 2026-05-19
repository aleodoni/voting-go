package votacao

import (
	"context"

	"github.com/aleodoni/go-ddd/domain"
	"github.com/aleodoni/voting-go/internal/application/shared"
	domainUsuario "github.com/aleodoni/voting-go/internal/domain/usuario"
	domainVotacao "github.com/aleodoni/voting-go/internal/domain/votacao"
	"github.com/aleodoni/voting-go/internal/platform/event"
)

// FechaVotacaoInput contém os dados necessários para fechar uma votação.
type FechaVotacaoInput struct {
	LoggedInUserKeycloakID string
	ProjetoID              string
}

// FechaVotacaoPayload é publicado no barramento de eventos quando uma votação é fechada.
type FechaVotacaoPayload struct {
	ProjetoID string `json:"projetoId"`
	VotacaoID string `json:"votacaoId"`
}

// FechaVotacaoUseCase fecha a votação aberta de um projeto, alterando seu status para fechado.
//
// Regras de negócio:
//   - o usuário autenticado deve ser administrador ativo
//   - o projeto deve possuir uma votação associada
//   - a votação deve estar com status aberto ([domainVotacao.StatusVotacaoA])
//
// Ao concluir com sucesso, atualiza o status da votação para [domainVotacao.StatusVotacaoF]
// e publica o evento [event.VotacaoFechada] no barramento.
type FechaVotacaoUseCase struct {
	repoUsuario domainUsuario.UsuarioRepository
	repoReuniao domainVotacao.ReuniaoRepository
	repoVotacao domainVotacao.VotacaoRepository
	bus         *event.Bus
}

// NewFechaVotacaoUseCase cria uma nova instância de [FechaVotacaoUseCase].
func NewFechaVotacaoUseCase(
	repoUsuario domainUsuario.UsuarioRepository,
	repoReuniao domainVotacao.ReuniaoRepository,
	repoVotacao domainVotacao.VotacaoRepository,
	bus *event.Bus,
) *FechaVotacaoUseCase {
	return &FechaVotacaoUseCase{
		repoUsuario: repoUsuario,
		repoReuniao: repoReuniao,
		repoVotacao: repoVotacao,
		bus:         bus,
	}
}

// Execute fecha a votação associada ao projeto informado em [FechaVotacaoInput.ProjetoID].
func (uc *FechaVotacaoUseCase) Execute(ctx context.Context, input FechaVotacaoInput) error {
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

	if projeto.Votacao.Status != domainVotacao.StatusVotacaoA {
		return domainVotacao.ErrVotacaoNaoAberta
	}

	projeto.Votacao.Fechar()

	if err := uc.repoVotacao.SalvaVotacao(ctx, projeto.Votacao); err != nil {
		return err
	}

	publishEvents(uc.bus, projeto.Votacao.PullEvents(), func(e domain.DomainEvent) (event.Event, bool) {
		switch evt := e.(type) {
		case domainVotacao.VotacaoFechadaEvent:
			return event.Event{
				Type:    event.VotacaoFechada,
				Payload: FechaVotacaoPayload{ProjetoID: evt.ProjetoID, VotacaoID: evt.VotacaoID},
			}, true
		}
		return event.Event{}, false
	})

	return nil
}
