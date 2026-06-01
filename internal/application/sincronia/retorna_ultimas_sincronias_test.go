package sincronia_test

import (
	"context"
	"errors"
	"testing"

	"github.com/aleodoni/go-ddd/domain"
	ucSincronia "github.com/aleodoni/voting-go/internal/application/sincronia"
	domainSincronia "github.com/aleodoni/voting-go/internal/domain/sincronia"
	domainUsuario "github.com/aleodoni/voting-go/internal/domain/usuario"
	"github.com/aleodoni/voting-go/internal/test/fakes"
)

func TestRetornaSincronias_Sucesso(t *testing.T) {
	usuarioRepo := fakes.NewFakeUsuarioRepository()
	sincroniaRepo := fakes.NewFakeSincroniaRepository()

	usuarioRepo.Seed(adminAtivo())
	sincroniaRepo.Seed(&domainSincronia.Sincronia{Entity: domain.Entity[string]{ID: "sync-1"}})
	sincroniaRepo.Seed(&domainSincronia.Sincronia{Entity: domain.Entity[string]{ID: "sync-2"}})
	sincroniaRepo.Seed(&domainSincronia.Sincronia{Entity: domain.Entity[string]{ID: "sync-3"}})

	uc := ucSincronia.NewRetornaSincroniasUseCase(sincroniaRepo, usuarioRepo)
	result, err := uc.Execute(context.Background(), ucSincronia.RetornaSincroniasInput{
		LoggedInUserKeycloakID: "keycloak-admin",
	})

	if err != nil {
		t.Fatalf("esperava nil, obteve: %v", err)
	}
	if result == nil {
		t.Fatal("esperava lista, obteve nil")
	}
	if len(result.Sincronias) != 3 {
		t.Fatalf("esperava 3 sincronias, obteve %d", len(result.Sincronias))
	}
}

func TestRetornaSincronias_ListaVazia(t *testing.T) {
	usuarioRepo := fakes.NewFakeUsuarioRepository()
	sincroniaRepo := fakes.NewFakeSincroniaRepository()

	usuarioRepo.Seed(adminAtivo())

	uc := ucSincronia.NewRetornaSincroniasUseCase(sincroniaRepo, usuarioRepo)
	result, err := uc.Execute(context.Background(), ucSincronia.RetornaSincroniasInput{
		LoggedInUserKeycloakID: "keycloak-admin",
	})

	if err != nil {
		t.Fatalf("esperava nil, obteve: %v", err)
	}
	if len(result.Sincronias) != 0 {
		t.Fatalf("esperava lista vazia, obteve %d itens", len(result.Sincronias))
	}
}

func TestRetornaSincronias_AdminNaoEncontrado(t *testing.T) {
	usuarioRepo := fakes.NewFakeUsuarioRepository()
	sincroniaRepo := fakes.NewFakeSincroniaRepository()

	uc := ucSincronia.NewRetornaSincroniasUseCase(sincroniaRepo, usuarioRepo)
	_, err := uc.Execute(context.Background(), ucSincronia.RetornaSincroniasInput{
		LoggedInUserKeycloakID: "keycloak-inexistente",
	})

	if !errors.Is(err, domainUsuario.ErrUserNotFound) {
		t.Fatalf("esperava ErrUserNotFound, obteve: %v", err)
	}
}

func TestRetornaSincronias_AdminInativo(t *testing.T) {
	usuarioRepo := fakes.NewFakeUsuarioRepository()
	sincroniaRepo := fakes.NewFakeSincroniaRepository()

	usuarioRepo.Seed(&domainUsuario.Usuario{
		AggregateRoot: domain.NewAggregateRoot("user-inativo"),
		KeycloakID:    "keycloak-inativo",
		Credencial:    &domainUsuario.Credencial{Ativo: false, PodeAdministrar: true},
	})

	uc := ucSincronia.NewRetornaSincroniasUseCase(sincroniaRepo, usuarioRepo)
	_, err := uc.Execute(context.Background(), ucSincronia.RetornaSincroniasInput{
		LoggedInUserKeycloakID: "keycloak-inativo",
	})

	if !errors.Is(err, domainUsuario.ErrUserNotActive) {
		t.Fatalf("esperava ErrUserNotActive, obteve: %v", err)
	}
}

func TestRetornaSincronias_AdminSemPermissao(t *testing.T) {
	usuarioRepo := fakes.NewFakeUsuarioRepository()
	sincroniaRepo := fakes.NewFakeSincroniaRepository()

	usuarioRepo.Seed(&domainUsuario.Usuario{
		AggregateRoot: domain.NewAggregateRoot("user-comum"),
		KeycloakID:    "keycloak-comum",
		Credencial:    &domainUsuario.Credencial{Ativo: true, PodeAdministrar: false},
	})

	uc := ucSincronia.NewRetornaSincroniasUseCase(sincroniaRepo, usuarioRepo)
	_, err := uc.Execute(context.Background(), ucSincronia.RetornaSincroniasInput{
		LoggedInUserKeycloakID: "keycloak-comum",
	})

	if !errors.Is(err, domainUsuario.ErrUserNotAdmin) {
		t.Fatalf("esperava ErrUserNotAdmin, obteve: %v", err)
	}
}

func TestRetornaSincronias_ErroRepositorio(t *testing.T) {
	usuarioRepo := fakes.NewFakeUsuarioRepository()
	sincroniaRepo := fakes.NewFakeSincroniaRepository()

	usuarioRepo.Seed(adminAtivo())
	erroEsperado := errors.New("falha ao listar")
	sincroniaRepo.ListLastErr = erroEsperado

	uc := ucSincronia.NewRetornaSincroniasUseCase(sincroniaRepo, usuarioRepo)
	_, err := uc.Execute(context.Background(), ucSincronia.RetornaSincroniasInput{
		LoggedInUserKeycloakID: "keycloak-admin",
	})

	if !errors.Is(err, erroEsperado) {
		t.Fatalf("esperava erro do repositório, obteve: %v", err)
	}
}
