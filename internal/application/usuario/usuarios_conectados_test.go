package usuario_test

import (
	"context"
	"errors"
	"testing"

	"github.com/aleodoni/go-ddd/domain"

	usecase "github.com/aleodoni/voting-go/internal/application/usuario"
	domainUsuario "github.com/aleodoni/voting-go/internal/domain/usuario"
	"github.com/aleodoni/voting-go/internal/platform/event"
	"github.com/aleodoni/voting-go/internal/test/fakes"
)

// fakeConectados implementa usecase.ConnectedUsersProvider e conta as consultas.
type fakeConectados struct {
	usuarios []*event.Subscriber
	chamadas int
}

func (f *fakeConectados) ConnectedUsers() []*event.Subscriber {
	f.chamadas++

	return f.usuarios
}

func TestListConnectedUsers_Sucesso(t *testing.T) {
	repo := fakes.NewFakeUsuarioRepository()
	repo.Seed(adminAtivo())

	conectados := &fakeConectados{usuarios: []*event.Subscriber{
		{UserID: "user-1", Username: "maria", IsAdmin: true},
		{UserID: "user-2", Username: "joao", IsAdmin: false},
	}}

	uc := usecase.NewListConnectedUsersUseCase(repo, conectados)
	resultado, err := uc.Execute(context.Background(), usecase.ListConnectedUsersInput{
		LoggedInUserKeycloakID: "keycloak-admin",
	})
	if err != nil {
		t.Fatalf("esperava nil, obteve: %v", err)
	}

	if len(resultado) != 2 {
		t.Fatalf("esperava 2 usuários conectados, obteve %d", len(resultado))
	}

	if resultado[0].Username != "maria" || resultado[1].Username != "joao" {
		t.Fatalf("usuários inesperados: %+v, %+v", resultado[0], resultado[1])
	}
}

func TestListConnectedUsers_UsuarioNaoAdmin(t *testing.T) {
	repo := fakes.NewFakeUsuarioRepository()
	repo.Seed(&domainUsuario.Usuario{
		AggregateRoot: domain.NewAggregateRoot("user-vereador"),
		KeycloakID:    "keycloak-vereador",
		Credencial: &domainUsuario.Credencial{
			Ativo:           true,
			PodeVotar:       true,
			PodeAdministrar: false,
		},
	})

	conectados := &fakeConectados{}

	uc := usecase.NewListConnectedUsersUseCase(repo, conectados)
	_, err := uc.Execute(context.Background(), usecase.ListConnectedUsersInput{
		LoggedInUserKeycloakID: "keycloak-vereador",
	})

	if !errors.Is(err, domainUsuario.ErrUserNotAdmin) {
		t.Fatalf("esperava ErrUserNotAdmin, obteve: %v", err)
	}

	if conectados.chamadas != 0 {
		t.Fatalf("não deveria consultar os conectados sem permissão, chamadas: %d", conectados.chamadas)
	}
}

func TestListConnectedUsers_AdminInativo(t *testing.T) {
	repo := fakes.NewFakeUsuarioRepository()
	repo.Seed(&domainUsuario.Usuario{
		AggregateRoot: domain.NewAggregateRoot("user-inativo"),
		KeycloakID:    "keycloak-inativo",
		Credencial: &domainUsuario.Credencial{
			Ativo:           false,
			PodeAdministrar: true,
		},
	})

	conectados := &fakeConectados{}

	uc := usecase.NewListConnectedUsersUseCase(repo, conectados)
	_, err := uc.Execute(context.Background(), usecase.ListConnectedUsersInput{
		LoggedInUserKeycloakID: "keycloak-inativo",
	})

	if !errors.Is(err, domainUsuario.ErrUserNotActive) {
		t.Fatalf("esperava ErrUserNotActive, obteve: %v", err)
	}

	if conectados.chamadas != 0 {
		t.Fatalf("não deveria consultar os conectados sem permissão, chamadas: %d", conectados.chamadas)
	}
}

func TestListConnectedUsers_UsuarioNaoEncontrado(t *testing.T) {
	repo := fakes.NewFakeUsuarioRepository()
	conectados := &fakeConectados{}

	uc := usecase.NewListConnectedUsersUseCase(repo, conectados)
	_, err := uc.Execute(context.Background(), usecase.ListConnectedUsersInput{
		LoggedInUserKeycloakID: "keycloak-inexistente",
	})

	if !errors.Is(err, domainUsuario.ErrUserNotFound) {
		t.Fatalf("esperava ErrUserNotFound, obteve: %v", err)
	}

	if conectados.chamadas != 0 {
		t.Fatalf("não deveria consultar os conectados sem permissão, chamadas: %d", conectados.chamadas)
	}
}
