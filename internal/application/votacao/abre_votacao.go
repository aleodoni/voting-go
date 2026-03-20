package votacao

import (
	"context"

	"github.com/aleodoni/go-ddd/domain"
	"github.com/aleodoni/voting-go/internal/application/shared"
	domainUsuario "github.com/aleodoni/voting-go/internal/domain/usuario"
	domainVotacao "github.com/aleodoni/voting-go/internal/domain/votacao"
	"github.com/aleodoni/voting-go/internal/platform/event"
	"github.com/aleodoni/voting-go/internal/platform/id"
)

// AbreVotacaoInput contém os dados necessários para abrir uma votação.
type AbreVotacaoInput struct {
	LoggedInUserKeycloakID string
	ProjetoID              string
}

// AbreVotacaoPayload é publicado no barramento de eventos quando uma votação é aberta.
type AbreVotacaoPayload struct {
	ProjetoID string `json:"projetoId"`
	VotacaoID string `json:"votacaoId"`
}

// AbreVotacaoUseCase abre uma votação para um projeto em uma reunião.
//
// Regras de negócio:
//   - o usuário autenticado deve ser administrador ativo
//   - não pode existir nenhuma votação aberta no sistema
//   - o projeto deve existir
//   - o projeto não pode ter uma votação associada
//
// Ao concluir com sucesso, publica o evento [event.VotacaoAberta] no barramento.
type AbreVotacaoUseCase struct {
	repoUsuario domainUsuario.UsuarioRepository
	repoReuniao domainVotacao.ReuniaoRepository
	repoVotacao domainVotacao.VotacaoRepository
	bus         *event.Bus
}

// NewAbreVotacaoUseCase cria uma nova instância de [AbreVotacaoUseCase].
func NewAbreVotacaoUseCase(
	repoUsuario domainUsuario.UsuarioRepository,
	repoReuniao domainVotacao.ReuniaoRepository,
	repoVotacao domainVotacao.VotacaoRepository,
	bus *event.Bus,
) *AbreVotacaoUseCase {
	return &AbreVotacaoUseCase{
		repoUsuario: repoUsuario,
		repoReuniao: repoReuniao,
		repoVotacao: repoVotacao,
		bus:         bus,
	}
}

// Execute abre uma votação para o projeto informado em [AbreVotacaoInput.ProjetoID].
func (uc *AbreVotacaoUseCase) Execute(ctx context.Context, input AbreVotacaoInput) error {
	if err := shared.VerificarAdmin(ctx, uc.repoUsuario, input.LoggedInUserKeycloakID); err != nil {
		return err
	}

	votacaoExistente, err := uc.repoVotacao.GetVotacaoAberta(ctx)
	if err != nil {
		return err
	}
	if votacaoExistente != nil {
		return domainVotacao.ErrVotacaoAberta
	}

	projeto, err := uc.repoReuniao.GetProjetoCompleto(ctx, input.ProjetoID)
	if err != nil {
		return err
	}
	if projeto == nil {
		return domainVotacao.ErrProjetoNotFound
	}
	if projeto.Votacao != nil {
		return domainVotacao.ErrProjetoVoted
	}

	votacaoNova := &domainVotacao.Votacao{
		AggregateRoot: domain.NewAggregateRoot(id.New()),
	}
	votacaoNova.Abrir(projeto.ID)

	if err := uc.repoVotacao.SalvaVotacao(ctx, votacaoNova); err != nil {
		return err
	}

	for _, e := range votacaoNova.PullEvents() {
		switch evt := e.(type) {
		case domainVotacao.VotacaoAbertaEvent:
			uc.bus.Publish(event.Event{
				Type: event.VotacaoAberta,
				Payload: AbreVotacaoPayload{
					ProjetoID: evt.ProjetoID,
					VotacaoID: evt.VotacaoID,
				},
			})
		}
	}

	return nil
}
