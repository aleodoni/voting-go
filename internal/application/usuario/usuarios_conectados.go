package usuario

import (
	"context"

	"github.com/aleodoni/voting-go/internal/application/shared"
	domainUsuario "github.com/aleodoni/voting-go/internal/domain/usuario"
	"github.com/aleodoni/voting-go/internal/platform/event"
)

// ConnectedUsersProvider fornece os usuários com conexão SSE ativa.
// É implementada por [event.Bus].
type ConnectedUsersProvider interface {
	ConnectedUsers() []*event.Subscriber
}

// ListConnectedUsersInput contém os dados necessários para listar os usuários conectados.
type ListConnectedUsersInput struct {
	LoggedInUserKeycloakID string
}

// ListConnectedUsersUseCase retorna os usuários com conexão SSE ativa.
//
// Regras de negócio:
//   - o usuário autenticado deve ser administrador ativo
type ListConnectedUsersUseCase struct {
	repo       domainUsuario.UsuarioRepository
	conectados ConnectedUsersProvider
}

// NewListConnectedUsersUseCase cria uma nova instância de [ListConnectedUsersUseCase].
func NewListConnectedUsersUseCase(repo domainUsuario.UsuarioRepository, conectados ConnectedUsersProvider) *ListConnectedUsersUseCase {
	return &ListConnectedUsersUseCase{repo: repo, conectados: conectados}
}

// Execute verifica se o usuário logado é administrador e retorna os usuários conectados.
func (uc *ListConnectedUsersUseCase) Execute(ctx context.Context, input ListConnectedUsersInput) ([]*event.Subscriber, error) {
	if err := shared.VerificarAdmin(ctx, uc.repo, input.LoggedInUserKeycloakID); err != nil {
		return nil, err
	}

	return uc.conectados.ConnectedUsers(), nil
}
