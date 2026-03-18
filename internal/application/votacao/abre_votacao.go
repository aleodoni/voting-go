package votacao

import (
	"context"

	"github.com/aleodoni/voting-go/internal/application/shared"
	"github.com/aleodoni/voting-go/internal/domain/usuario"
	"github.com/aleodoni/voting-go/internal/domain/votacao"
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
//
// Ao concluir com sucesso, publica o evento [event.VotacaoAberta] no barramento.
type AbreVotacaoUseCase struct {
	repoUsuario usuario.UsuarioRepository
	repoReuniao votacao.ReuniaoRepository
	repoVotacao votacao.VotacaoRepository
	bus         *event.Bus
}

// NewAbreVotacaoUseCase cria uma nova instância de [AbreVotacaoUseCase].
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
	// Verificar se o usuário logado é admin ativo
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

	// Busca projeto para garantir que ele existe
	projeto, err := uc.repoReuniao.GetProjetoCompleto(ctx, input.ProjetoID)
	if err != nil {
		return err
	}

	// Se nao encontrar o projeto, retorna erro
	if projeto == nil {
		return votacao.ErrProjetoNotFound
	}

	// Verifica se já existe uma votação associada a esse projeto
	if projeto.Votacao != nil {
		return votacao.ErrProjetoVoted
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
