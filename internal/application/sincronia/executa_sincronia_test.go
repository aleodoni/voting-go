package sincronia_test

import (
	"context"
	"errors"
	"testing"

	"github.com/aleodoni/go-ddd/domain"
	ucSincronia "github.com/aleodoni/voting-go/internal/application/sincronia"
	domainUsuario "github.com/aleodoni/voting-go/internal/domain/usuario"
	"github.com/aleodoni/voting-go/internal/test/fakes"
)

// ─── helpers ────────────────────────────────────────────────────────────────

func adminAtivo() *domainUsuario.Usuario {
	return &domainUsuario.Usuario{
		AggregateRoot: domain.NewAggregateRoot("user-admin"),
		KeycloakID:    "keycloak-admin",
		Credencial: &domainUsuario.Credencial{
			Ativo:           true,
			PodeAdministrar: true,
		},
	}
}

// ─── testes ──────────────────────────────────────────────────────────────────

func TestExecutaSincronia_Sucesso(t *testing.T) {
	usuarioRepo := fakes.NewFakeUsuarioRepository()
	sincroniaRepo := fakes.NewFakeSincroniaRepository()

	usuarioRepo.Seed(adminAtivo())

	uc := ucSincronia.NewExecutaSincroniaUseCase(sincroniaRepo, usuarioRepo)
	result, err := uc.Execute(context.Background(), ucSincronia.ExecutaSincroniaInput{
		LoggedInUserKeycloakID: "keycloak-admin",
	})

	if err != nil {
		t.Fatalf("esperava nil, obteve: %v", err)
	}
	if result == nil {
		t.Fatal("esperava sincronia retornada, obteve nil")
	}
	if sincroniaRepo.SyncCalls != 1 {
		t.Fatalf("esperava 1 chamada a Sync, obteve %d", sincroniaRepo.SyncCalls)
	}
}

func TestExecutaSincronia_AdminNaoEncontrado(t *testing.T) {
	usuarioRepo := fakes.NewFakeUsuarioRepository()
	sincroniaRepo := fakes.NewFakeSincroniaRepository()

	uc := ucSincronia.NewExecutaSincroniaUseCase(sincroniaRepo, usuarioRepo)
	_, err := uc.Execute(context.Background(), ucSincronia.ExecutaSincroniaInput{
		LoggedInUserKeycloakID: "keycloak-inexistente",
	})

	if !errors.Is(err, domainUsuario.ErrUserNotFound) {
		t.Fatalf("esperava ErrUserNotFound, obteve: %v", err)
	}
	if sincroniaRepo.SyncCalls != 0 {
		t.Fatal("Sync não deve ser chamado quando admin não é encontrado")
	}
}

func TestExecutaSincronia_AdminInativo(t *testing.T) {
	usuarioRepo := fakes.NewFakeUsuarioRepository()
	sincroniaRepo := fakes.NewFakeSincroniaRepository()

	usuarioRepo.Seed(&domainUsuario.Usuario{
		AggregateRoot: domain.NewAggregateRoot("user-inativo"),
		KeycloakID:    "keycloak-inativo",
		Credencial:    &domainUsuario.Credencial{Ativo: false, PodeAdministrar: true},
	})

	uc := ucSincronia.NewExecutaSincroniaUseCase(sincroniaRepo, usuarioRepo)
	_, err := uc.Execute(context.Background(), ucSincronia.ExecutaSincroniaInput{
		LoggedInUserKeycloakID: "keycloak-inativo",
	})

	if !errors.Is(err, domainUsuario.ErrUserNotActive) {
		t.Fatalf("esperava ErrUserNotActive, obteve: %v", err)
	}
}

func TestExecutaSincronia_AdminSemPermissao(t *testing.T) {
	usuarioRepo := fakes.NewFakeUsuarioRepository()
	sincroniaRepo := fakes.NewFakeSincroniaRepository()

	usuarioRepo.Seed(&domainUsuario.Usuario{
		AggregateRoot: domain.NewAggregateRoot("user-comum"),
		KeycloakID:    "keycloak-comum",
		Credencial:    &domainUsuario.Credencial{Ativo: true, PodeAdministrar: false},
	})

	uc := ucSincronia.NewExecutaSincroniaUseCase(sincroniaRepo, usuarioRepo)
	_, err := uc.Execute(context.Background(), ucSincronia.ExecutaSincroniaInput{
		LoggedInUserKeycloakID: "keycloak-comum",
	})

	if !errors.Is(err, domainUsuario.ErrUserNotAdmin) {
		t.Fatalf("esperava ErrUserNotAdmin, obteve: %v", err)
	}
}

func TestExecutaSincronia_ErroRepositorio(t *testing.T) {
	usuarioRepo := fakes.NewFakeUsuarioRepository()
	sincroniaRepo := fakes.NewFakeSincroniaRepository()

	usuarioRepo.Seed(adminAtivo())
	erroEsperado := errors.New("falha na sincronização")
	sincroniaRepo.SyncErr = erroEsperado

	uc := ucSincronia.NewExecutaSincroniaUseCase(sincroniaRepo, usuarioRepo)
	_, err := uc.Execute(context.Background(), ucSincronia.ExecutaSincroniaInput{
		LoggedInUserKeycloakID: "keycloak-admin",
	})

	if !errors.Is(err, erroEsperado) {
		t.Fatalf("esperava erro do repositório, obteve: %v", err)
	}
}

func TestAutorizar_AdminAtivo(t *testing.T) {
	usuarioRepo := fakes.NewFakeUsuarioRepository()
	sincroniaRepo := fakes.NewFakeSincroniaRepository()
	usuarioRepo.Seed(adminAtivo())

	uc := ucSincronia.NewExecutaSincroniaUseCase(sincroniaRepo, usuarioRepo)
	err := uc.Autorizar(context.Background(), ucSincronia.ExecutaSincroniaInput{
		LoggedInUserKeycloakID: "keycloak-admin",
	})

	if err != nil {
		t.Fatalf("esperava nil, obteve: %v", err)
	}

	if sincroniaRepo.SyncCalls != 0 {
		t.Fatalf("Autorizar não deveria executar a sincronização, chamadas a Sync: %d", sincroniaRepo.SyncCalls)
	}
}

func TestAutorizar_UsuarioNaoEncontrado(t *testing.T) {
	uc := ucSincronia.NewExecutaSincroniaUseCase(fakes.NewFakeSincroniaRepository(), fakes.NewFakeUsuarioRepository())

	err := uc.Autorizar(context.Background(), ucSincronia.ExecutaSincroniaInput{
		LoggedInUserKeycloakID: "keycloak-inexistente",
	})

	if !errors.Is(err, domainUsuario.ErrUserNotFound) {
		t.Fatalf("esperava ErrUserNotFound, obteve: %v", err)
	}
}

func TestAutorizar_UsuarioSemPermissaoAdmin(t *testing.T) {
	usuarioRepo := fakes.NewFakeUsuarioRepository()
	usuarioRepo.Seed(&domainUsuario.Usuario{
		AggregateRoot: domain.NewAggregateRoot("user-vereador"),
		KeycloakID:    "keycloak-vereador",
		Credencial: &domainUsuario.Credencial{
			Ativo:           true,
			PodeAdministrar: false,
		},
	})

	uc := ucSincronia.NewExecutaSincroniaUseCase(fakes.NewFakeSincroniaRepository(), usuarioRepo)

	err := uc.Autorizar(context.Background(), ucSincronia.ExecutaSincroniaInput{
		LoggedInUserKeycloakID: "keycloak-vereador",
	})

	if !errors.Is(err, domainUsuario.ErrUserNotAdmin) {
		t.Fatalf("esperava ErrUserNotAdmin, obteve: %v", err)
	}
}
