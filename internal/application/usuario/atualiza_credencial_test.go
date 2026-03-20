package usuario_test

import (
	"context"
	"testing"

	"github.com/aleodoni/go-ddd/domain"
	usecase "github.com/aleodoni/voting-go/internal/application/usuario"
	domainUsuario "github.com/aleodoni/voting-go/internal/domain/usuario"
	"github.com/aleodoni/voting-go/internal/platform/id"
	"github.com/aleodoni/voting-go/internal/test/fakes"
	"github.com/nrednav/cuid2"
)

// --- helpers ---

func adminAtivo() *domainUsuario.Usuario {
	return &domainUsuario.Usuario{
		AggregateRoot: domain.NewAggregateRoot(cuid2.Generate()),
		KeycloakID:    "keycloak-admin",
		Credencial: &domainUsuario.Credencial{
			Ativo:           true,
			PodeAdministrar: true,
		},
	}
}

// --- testes ---

func TestUpdateCredencial_Sucesso(t *testing.T) {
	usuarioRepo := fakes.NewFakeUsuarioRepository()
	usuarioRepo.Seed(adminAtivo())

	credencialRepo := fakes.NewFakeCredencialRepository()
	credencialRepo.Seed(&domainUsuario.Credencial{
		Entity:    domain.Entity[string]{ID: id.New()},
		UsuarioID: "user-1",
		Ativo:     false,
		PodeVotar: false,
	})

	uc := usecase.NewUpdateCredencialUseCase(usuarioRepo, credencialRepo)

	cred, err := uc.Execute(context.Background(), usecase.UpdateCredencialInput{
		AdminKeycloakID: "keycloak-admin",
		UsuarioID:       "user-1",
		Ativo:           true,
		PodeVotar:       true,
		PodeAdministrar: false,
	})

	if err != nil {
		t.Fatalf("esperava nil, got %v", err)
	}
	if !cred.Ativo {
		t.Error("esperava Ativo = true")
	}
	if !cred.PodeVotar {
		t.Error("esperava PodeVotar = true")
	}
	if cred.PodeAdministrar {
		t.Error("esperava PodeAdministrar = false")
	}
}

func TestUpdateCredencial_AdminInativo_Inativo(t *testing.T) {
	usuarioRepo := fakes.NewFakeUsuarioRepository()
	usuarioRepo.Seed(&domainUsuario.Usuario{
		AggregateRoot: domain.NewAggregateRoot(cuid2.Generate()),
		KeycloakID:    "keycloak-admin",
		Credencial: &domainUsuario.Credencial{
			Ativo:           false,
			PodeAdministrar: true,
		},
	})

	uc := usecase.NewUpdateCredencialUseCase(usuarioRepo, fakes.NewFakeCredencialRepository())

	_, err := uc.Execute(context.Background(), usecase.UpdateCredencialInput{
		AdminKeycloakID: "keycloak-admin",
		UsuarioID:       "user-1",
	})

	if err != domainUsuario.ErrUserNotActive {
		t.Errorf("esperava ErrUserNotActive, got %v", err)
	}
}

func TestUpdateCredencial_AdminSemPermissao_Forbidden(t *testing.T) {
	usuarioRepo := fakes.NewFakeUsuarioRepository()
	usuarioRepo.Seed(&domainUsuario.Usuario{
		AggregateRoot: domain.NewAggregateRoot(cuid2.Generate()),
		KeycloakID:    "keycloak-admin",
		Credencial: &domainUsuario.Credencial{
			Ativo:           true,
			PodeAdministrar: false,
		},
	})

	uc := usecase.NewUpdateCredencialUseCase(usuarioRepo, fakes.NewFakeCredencialRepository())

	_, err := uc.Execute(context.Background(), usecase.UpdateCredencialInput{
		AdminKeycloakID: "keycloak-admin",
		UsuarioID:       "user-1",
	})

	if err != domainUsuario.ErrUserNotAdmin {
		t.Errorf("esperava ErrUserNotAdmin, got %v", err)
	}
}

func TestUpdateCredencial_AdminNaoEncontrado(t *testing.T) {
	uc := usecase.NewUpdateCredencialUseCase(fakes.NewFakeUsuarioRepository(), fakes.NewFakeCredencialRepository())

	_, err := uc.Execute(context.Background(), usecase.UpdateCredencialInput{
		AdminKeycloakID: "keycloak-inexistente",
		UsuarioID:       "user-1",
	})

	if err != domainUsuario.ErrUserNotFound {
		t.Errorf("esperava ErrUserNotFound, got %v", err)
	}
}

func TestUpdateCredencial_UsuarioAlvoNaoEncontrado(t *testing.T) {
	usuarioRepo := fakes.NewFakeUsuarioRepository()
	usuarioRepo.Seed(adminAtivo())

	uc := usecase.NewUpdateCredencialUseCase(usuarioRepo, fakes.NewFakeCredencialRepository())

	_, err := uc.Execute(context.Background(), usecase.UpdateCredencialInput{
		AdminKeycloakID: "keycloak-admin",
		UsuarioID:       "user-inexistente",
	})

	if err != domainUsuario.ErrCredencialNotFound {
		t.Errorf("esperava ErrCredencialNotFound, got %v", err)
	}
}
