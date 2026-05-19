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

// RegistraVotoInput contém os dados necessários para registrar um voto.
type RegistraVotoInput struct {
	LoggedInUserKeycloakID string
	VotacaoID              string
	Voto                   domainVotacao.OpcaoVoto
	Restricao              *domainVotacao.Restricao
	VotoContrario          *domainVotacao.VotoContrario
}

// RegistraVotoPayload é publicado no barramento de eventos quando um voto é registrado.
type RegistraVotoPayload struct {
	VotacaoID string `json:"votacaoId"`
}

// RegistraVotoUseCase registra o voto do usuário autenticado em uma votação aberta.
//
// Regras de negócio:
//   - o usuário autenticado deve estar ativo e possuir permissão de voto
//   - a votação deve existir e estar com status aberto ([domainVotacao.StatusVotacaoA])
//   - o usuário não pode votar mais de uma vez na mesma votação
//   - [RegistraVotoInput.Restricao] e [RegistraVotoInput.VotoContrario] são opcionais;
//     quando informados, recebem IDs gerados automaticamente
//
// Ao concluir com sucesso, publica o evento [event.VotoRegistrado] no barramento.
type RegistraVotoUseCase struct {
	repoUsuario domainUsuario.UsuarioRepository
	repoVotacao domainVotacao.VotacaoRepository
	bus         *event.Bus
}

// NewRegistraVotoUseCase cria uma nova instância de [RegistraVotoUseCase].
func NewRegistraVotoUseCase(
	repoUsuario domainUsuario.UsuarioRepository,
	repoVotacao domainVotacao.VotacaoRepository,
	bus *event.Bus,
) *RegistraVotoUseCase {
	return &RegistraVotoUseCase{
		repoUsuario: repoUsuario,
		repoVotacao: repoVotacao,
		bus:         bus,
	}
}

// Execute registra o voto do usuário autenticado na votação informada em
// [RegistraVotoInput.VotacaoID].
func (uc *RegistraVotoUseCase) Execute(ctx context.Context, input RegistraVotoInput) error {
	u, err := shared.VerificarVota(ctx, uc.repoUsuario, input.LoggedInUserKeycloakID)
	if err != nil {
		return err
	}

	v, err := uc.repoVotacao.BuscaVotacao(ctx, input.VotacaoID)
	if err != nil {
		return err
	}
	if v.Status != domainVotacao.StatusVotacaoA {
		return domainVotacao.ErrVotacaoNaoAberta
	}

	jaVotou, err := uc.repoVotacao.UsuarioJaVotou(ctx, u.ID, input.VotacaoID)
	if err != nil {
		return err
	}
	if jaVotou {
		return domainVotacao.ErrUsuarioJaVotou
	}

	voto := &domainVotacao.Voto{
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

	v.RegistrarVoto(voto)

	publishEvents(uc.bus, v.PullEvents(), func(e domain.DomainEvent) (event.Event, bool) {
		switch evt := e.(type) {
		case domainVotacao.VotoRegistradoEvent:
			return event.Event{
				Type:    event.VotoRegistrado,
				Payload: RegistraVotoPayload{VotacaoID: evt.VotacaoID},
			}, true
		}
		return event.Event{}, false
	})

	return nil
}
