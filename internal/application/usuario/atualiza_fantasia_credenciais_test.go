package usuario_test

import (
	"context"
	"testing"

	usecase "github.com/aleodoni/voting-go/internal/application/usuario"
	domainUsuario "github.com/aleodoni/voting-go/internal/domain/usuario"
	"github.com/aleodoni/voting-go/internal/test/fakes"
)

//
// helpers
//

func adminUsuario(keycloakID, userID string) *domainUsuario.Usuario {
	return &domainUsuario.Usuario{
		ID:         userID,
		KeycloakID: keycloakID,
		Username:   "admin",
		Credencial: &domainUsuario.Credencial{
			ID:              "cred-admin",
			UsuarioID:       userID,
			Ativo:           true,
			PodeAdministrar: true,
		},
	}
}

func ptr(s string) *string {
	return &s
}

//
// tests
//

func TestUpdateDisplayNamePermissions_AdminAtualiza(t *testing.T) {

	repo := fakes.NewFakeUsuarioRepository()

	repo.Seed(adminUsuario("keycloak-admin", "user-admin"))
	repo.Seed(&domainUsuario.Usuario{
		ID:         "user-alvo",
		KeycloakID: "keycloak-alvo",
		Username:   "alvo",
	})

	uc := usecase.NewUpdateDisplayNamePermissionsUseCase(repo)

	err := uc.Execute(context.Background(), usecase.UpdateDisplayNamePermissionsInput{
		LoggedInUserKeycloakID: "keycloak-admin",
		UserID:                 "user-alvo",
		DisplayName:            ptr("Novo Nome"),
		IsActive:               true,
		CanAdmin:               false,
		CanVote:                true,
	})

	if err != nil {
		t.Fatalf("esperava nil, got %v", err)
	}
}

func TestUpdateDisplayNamePermissions_UsuarioLogadoNaoEncontrado(t *testing.T) {

	repo := fakes.NewFakeUsuarioRepository()

	uc := usecase.NewUpdateDisplayNamePermissionsUseCase(repo)

	err := uc.Execute(context.Background(), usecase.UpdateDisplayNamePermissionsInput{
		LoggedInUserKeycloakID: "keycloak-inexistente",
		UserID:                 "user-alvo",
	})

	if err == nil {
		t.Fatal("esperava erro, got nil")
	}
}

func TestUpdateDisplayNamePermissions_UsuarioLogadoSemCredencial(t *testing.T) {

	repo := fakes.NewFakeUsuarioRepository()

	repo.Seed(&domainUsuario.Usuario{
		ID:         "user-sem-cred",
		KeycloakID: "keycloak-sem-cred",
		Username:   "semcred",
		Credencial: nil,
	})

	uc := usecase.NewUpdateDisplayNamePermissionsUseCase(repo)

	err := uc.Execute(context.Background(), usecase.UpdateDisplayNamePermissionsInput{
		LoggedInUserKeycloakID: "keycloak-sem-cred",
		UserID:                 "user-alvo",
	})

	if err != domainUsuario.ErrUserNotAdmin {
		t.Fatalf("esperava ErrUserNotAdmin, got %v", err)
	}
}

func TestUpdateDisplayNamePermissions_UsuarioLogadoNaoEAdmin(t *testing.T) {

	repo := fakes.NewFakeUsuarioRepository()

	repo.Seed(&domainUsuario.Usuario{
		ID:         "user-comum",
		KeycloakID: "keycloak-comum",
		Username:   "comum",
		Credencial: &domainUsuario.Credencial{
			ID:              "cred-comum",
			UsuarioID:       "user-comum",
			Ativo:           true,
			PodeAdministrar: false,
		},
	})

	uc := usecase.NewUpdateDisplayNamePermissionsUseCase(repo)

	err := uc.Execute(context.Background(), usecase.UpdateDisplayNamePermissionsInput{
		LoggedInUserKeycloakID: "keycloak-comum",
		UserID:                 "user-alvo",
	})

	if err != domainUsuario.ErrUserNotAdmin {
		t.Fatalf("esperava ErrNotAdmin, got %v", err)
	}
}

func TestUpdateDisplayNamePermissions_UsuarioLogadoInativo(t *testing.T) {

	repo := fakes.NewFakeUsuarioRepository()

	repo.Seed(&domainUsuario.Usuario{
		ID:         "user-inativo",
		KeycloakID: "keycloak-inativo",
		Username:   "inativo",
		Credencial: &domainUsuario.Credencial{
			ID:              "cred-inativo",
			UsuarioID:       "user-inativo",
			Ativo:           false,
			PodeAdministrar: true,
		},
	})

	uc := usecase.NewUpdateDisplayNamePermissionsUseCase(repo)

	err := uc.Execute(context.Background(), usecase.UpdateDisplayNamePermissionsInput{
		LoggedInUserKeycloakID: "keycloak-inativo",
		UserID:                 "user-alvo",
	})

	if err != domainUsuario.ErrUserNotAdmin {
		t.Fatalf("esperava ErrNotAdmin, got %v", err)
	}
}

func TestUpdateDisplayNamePermissions_ErroNoRepositorio(t *testing.T) {

	repo := fakes.NewFakeUsuarioRepository()

	repo.Seed(adminUsuario("keycloak-admin", "user-admin"))
	repo.UpdateDisplayNamePermissionsErr = domainUsuario.ErrUserNotFound

	uc := usecase.NewUpdateDisplayNamePermissionsUseCase(repo)

	err := uc.Execute(context.Background(), usecase.UpdateDisplayNamePermissionsInput{
		LoggedInUserKeycloakID: "keycloak-admin",
		UserID:                 "user-inexistente",
	})

	if err == nil {
		t.Fatal("esperava erro, got nil")
	}
}
