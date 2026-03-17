package votacao_test

import (
	"context"
	"testing"

	"github.com/aleodoni/voting-go/internal/application/shared"
	usecase "github.com/aleodoni/voting-go/internal/application/votacao"
	domainUsuario "github.com/aleodoni/voting-go/internal/domain/usuario"
	domainVotacao "github.com/aleodoni/voting-go/internal/domain/votacao"
	"github.com/aleodoni/voting-go/internal/test/fakes"
)

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

//
// TESTES
//

func TestRetornaReunioesDia_AdminRetornaReunioesDodia(t *testing.T) {
	usuarioRepo := fakes.NewFakeUsuarioRepository()
	reuniaoRepo := fakes.NewFakeReuniaoRepository()

	usuarioRepo.Seed(adminUsuario("keycloak-admin", "user-admin"))

	hoje := shared.GetCurrentDate()
	reuniaoRepo.Seed(&domainVotacao.Reuniao{ID: "reuniao-1", RecNumero: "001", RecData: hoje})
	reuniaoRepo.Seed(&domainVotacao.Reuniao{ID: "reuniao-2", RecNumero: "002", RecData: hoje})

	uc := usecase.NewRetornaReunioesDiaUseCase(usuarioRepo, reuniaoRepo)

	result, err := uc.Execute(context.Background(), usecase.RetornaReunioesDiaInput{
		LoggedInUserKeycloakID: "keycloak-admin",
	})

	if err != nil {
		t.Fatalf("esperava nil, got %v", err)
	}

	if len(result) != 2 {
		t.Errorf("esperava 2 reunioes, got %d", len(result))
	}
}

func TestRetornaReunioesDia_AdminSemReunioesNoDia(t *testing.T) {
	usuarioRepo := fakes.NewFakeUsuarioRepository()
	reuniaoRepo := fakes.NewFakeReuniaoRepository()

	usuarioRepo.Seed(adminUsuario("keycloak-admin", "user-admin"))

	// Nenhuma reunião semeada -> deve retornar 0
	uc := usecase.NewRetornaReunioesDiaUseCase(usuarioRepo, reuniaoRepo)

	result, err := uc.Execute(context.Background(), usecase.RetornaReunioesDiaInput{
		LoggedInUserKeycloakID: "keycloak-admin",
	})

	if err != nil {
		t.Fatalf("esperava nil, got %v", err)
	}

	if len(result) != 0 {
		t.Errorf("esperava 0 reunioes, got %d", len(result))
	}
}

func TestRetornaReunioesDia_UsuarioNaoEncontrado(t *testing.T) {
	usuarioRepo := fakes.NewFakeUsuarioRepository()
	reuniaoRepo := fakes.NewFakeReuniaoRepository()

	uc := usecase.NewRetornaReunioesDiaUseCase(usuarioRepo, reuniaoRepo)

	_, err := uc.Execute(context.Background(), usecase.RetornaReunioesDiaInput{
		LoggedInUserKeycloakID: "keycloak-inexistente",
	})

	if err == nil {
		t.Fatal("esperava erro, got nil")
	}
}

func TestRetornaReunioesDia_UsuarioSemCredencial(t *testing.T) {
	usuarioRepo := fakes.NewFakeUsuarioRepository()
	reuniaoRepo := fakes.NewFakeReuniaoRepository()

	usuarioRepo.Seed(&domainUsuario.Usuario{
		ID:         "user-sem-cred",
		KeycloakID: "keycloak-sem-cred",
		Credencial: nil,
	})

	uc := usecase.NewRetornaReunioesDiaUseCase(usuarioRepo, reuniaoRepo)

	_, err := uc.Execute(context.Background(), usecase.RetornaReunioesDiaInput{
		LoggedInUserKeycloakID: "keycloak-sem-cred",
	})

	if err != domainUsuario.ErrUserNotActive {
		t.Fatalf("esperava ErrUserNotActive, got %v", err)
	}
}

func TestRetornaReunioesDia_UsuarioNaoEAdmin(t *testing.T) {
	usuarioRepo := fakes.NewFakeUsuarioRepository()
	reuniaoRepo := fakes.NewFakeReuniaoRepository()

	usuarioRepo.Seed(&domainUsuario.Usuario{
		ID:         "user-comum",
		KeycloakID: "keycloak-comum",
		Credencial: &domainUsuario.Credencial{
			Ativo:           true,
			PodeAdministrar: false,
		},
	})

	uc := usecase.NewRetornaReunioesDiaUseCase(usuarioRepo, reuniaoRepo)

	_, err := uc.Execute(context.Background(), usecase.RetornaReunioesDiaInput{
		LoggedInUserKeycloakID: "keycloak-comum",
	})

	if err != domainUsuario.ErrUserNotAdmin {
		t.Fatalf("esperava ErrUserNotAdmin, got %v", err)
	}
}

func TestRetornaReunioesDia_UsuarioInativo(t *testing.T) {
	usuarioRepo := fakes.NewFakeUsuarioRepository()
	reuniaoRepo := fakes.NewFakeReuniaoRepository()

	usuarioRepo.Seed(&domainUsuario.Usuario{
		ID:         "user-inativo",
		KeycloakID: "keycloak-inativo",
		Credencial: &domainUsuario.Credencial{
			Ativo:           false,
			PodeAdministrar: true,
		},
	})

	uc := usecase.NewRetornaReunioesDiaUseCase(usuarioRepo, reuniaoRepo)

	_, err := uc.Execute(context.Background(), usecase.RetornaReunioesDiaInput{
		LoggedInUserKeycloakID: "keycloak-inativo",
	})

	if err != domainUsuario.ErrUserNotActive {
		t.Fatalf("esperava ErrUserNotActive, got %v", err)
	}
}

func TestRetornaReunioesDia_ErroNoRepositorio(t *testing.T) {
	usuarioRepo := fakes.NewFakeUsuarioRepository()
	reuniaoRepo := fakes.NewFakeReuniaoRepository()

	usuarioRepo.Seed(adminUsuario("keycloak-admin", "user-admin"))
	reuniaoRepo.GetReunioesDiaErr = domainVotacao.ErrReuniaoNotFound

	uc := usecase.NewRetornaReunioesDiaUseCase(usuarioRepo, reuniaoRepo)

	_, err := uc.Execute(context.Background(), usecase.RetornaReunioesDiaInput{
		LoggedInUserKeycloakID: "keycloak-admin",
	})

	if err == nil {
		t.Fatal("esperava erro, got nil")
	}
}
