package usuario_test

import (
	"context"
	"errors"
	"testing"

	"github.com/aleodoni/go-ddd/domain"
	usecase "github.com/aleodoni/voting-go/internal/application/usuario"
	domainUsuario "github.com/aleodoni/voting-go/internal/domain/usuario"
	"github.com/aleodoni/voting-go/internal/test/fakes"
)

func TestRetornaUsuario_Sucesso(t *testing.T) {
	repo := fakes.NewFakeUsuarioRepository()

	admin := adminAtivo()
	alvo := &domainUsuario.Usuario{
		AggregateRoot: domain.NewAggregateRoot("user-alvo"),
		KeycloakID:    "keycloak-alvo",
		Nome:          "Usuário Alvo",
	}
	repo.Seed(admin)
	repo.Seed(alvo)

	uc := usecase.NewRetornaUsuarioUseCase(repo)
	result, err := uc.Execute(context.Background(), usecase.RetornaUsuarioInput{
		LoggedInUserKeycloakID: admin.KeycloakID,
		UsuarioID:              "user-alvo",
	})

	if err != nil {
		t.Fatalf("esperava nil, obteve: %v", err)
	}
	if result == nil {
		t.Fatal("esperava usuário, obteve nil")
	}
	if result.ID != "user-alvo" {
		t.Fatalf("esperava ID 'user-alvo', obteve '%s'", result.ID)
	}
}

func TestRetornaUsuario_UsuarioAlvoNaoEncontrado(t *testing.T) {
	repo := fakes.NewFakeUsuarioRepository()
	repo.Seed(adminAtivo())

	uc := usecase.NewRetornaUsuarioUseCase(repo)
	_, err := uc.Execute(context.Background(), usecase.RetornaUsuarioInput{
		LoggedInUserKeycloakID: "keycloak-admin",
		UsuarioID:              "id-inexistente",
	})

	if !errors.Is(err, domainUsuario.ErrUserNotFound) {
		t.Fatalf("esperava ErrUserNotFound, obteve: %v", err)
	}
}

func TestRetornaUsuario_AdminNaoEncontrado(t *testing.T) {
	repo := fakes.NewFakeUsuarioRepository()

	uc := usecase.NewRetornaUsuarioUseCase(repo)
	_, err := uc.Execute(context.Background(), usecase.RetornaUsuarioInput{
		LoggedInUserKeycloakID: "keycloak-inexistente",
		UsuarioID:              "qualquer",
	})

	if !errors.Is(err, domainUsuario.ErrUserNotFound) {
		t.Fatalf("esperava ErrUserNotFound, obteve: %v", err)
	}
}

func TestRetornaUsuario_AdminInativo(t *testing.T) {
	repo := fakes.NewFakeUsuarioRepository()
	repo.Seed(&domainUsuario.Usuario{
		AggregateRoot: domain.NewAggregateRoot("user-inativo"),
		KeycloakID:    "keycloak-inativo",
		Credencial:    &domainUsuario.Credencial{Ativo: false, PodeAdministrar: true},
	})

	uc := usecase.NewRetornaUsuarioUseCase(repo)
	_, err := uc.Execute(context.Background(), usecase.RetornaUsuarioInput{
		LoggedInUserKeycloakID: "keycloak-inativo",
		UsuarioID:              "qualquer",
	})

	if !errors.Is(err, domainUsuario.ErrUserNotActive) {
		t.Fatalf("esperava ErrUserNotActive, obteve: %v", err)
	}
}

func TestRetornaUsuario_AdminSemPermissao(t *testing.T) {
	repo := fakes.NewFakeUsuarioRepository()
	repo.Seed(&domainUsuario.Usuario{
		AggregateRoot: domain.NewAggregateRoot("user-comum"),
		KeycloakID:    "keycloak-comum",
		Credencial:    &domainUsuario.Credencial{Ativo: true, PodeAdministrar: false},
	})

	uc := usecase.NewRetornaUsuarioUseCase(repo)
	_, err := uc.Execute(context.Background(), usecase.RetornaUsuarioInput{
		LoggedInUserKeycloakID: "keycloak-comum",
		UsuarioID:              "qualquer",
	})

	if !errors.Is(err, domainUsuario.ErrUserNotAdmin) {
		t.Fatalf("esperava ErrUserNotAdmin, obteve: %v", err)
	}
}
