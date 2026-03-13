package votacao_test

import (
	"context"
	"testing"

	usecase "github.com/aleodoni/voting-go/internal/application/votacao"
	domainUsuario "github.com/aleodoni/voting-go/internal/domain/usuario"
	domainVotacao "github.com/aleodoni/voting-go/internal/domain/votacao"
	"github.com/aleodoni/voting-go/internal/test/fakes"
)

func TestRetornaProjetosCompletos_AdminRetornaProjetos(t *testing.T) {
	usuarioRepo := fakes.NewFakeUsuarioRepository()
	reuniaoRepo := fakes.NewFakeReuniaoRepository()

	usuarioRepo.Seed(adminUsuario("keycloak-admin", "user-admin"))

	// Seed da reunião
	reuniao := &domainVotacao.Reuniao{ID: "reuniao-1"}
	reuniaoRepo.Seed(reuniao)

	// Seed dos projetos
	reuniaoRepo.SeedProjetos("reuniao-1", []*domainVotacao.Projeto{
		{ID: "projeto-1", CodigoProposicao: "001"},
		{ID: "projeto-2", CodigoProposicao: "002"},
	})

	uc := usecase.NewRetornaProjetosCompletosUseCase(usuarioRepo, reuniaoRepo)

	result, err := uc.Execute(context.Background(), usecase.RetornaProjetosCompletosInput{
		LoggedInUserKeycloakID: "keycloak-admin",
		ReuniaoID:              "reuniao-1",
	})

	if err != nil {
		t.Fatalf("esperava nil, got %v", err)
	}
	if len(result) != 2 {
		t.Errorf("esperava 2 projetos, got %d", len(result))
	}
}

func TestRetornaProjetosCompletos_AdminSemProjetos(t *testing.T) {
	usuarioRepo := fakes.NewFakeUsuarioRepository()
	reuniaoRepo := fakes.NewFakeReuniaoRepository()

	usuarioRepo.Seed(adminUsuario("keycloak-admin", "user-admin"))

	reuniaoRepo.Seed(&domainVotacao.Reuniao{ID: "reuniao-1"})

	uc := usecase.NewRetornaProjetosCompletosUseCase(usuarioRepo, reuniaoRepo)

	result, err := uc.Execute(context.Background(), usecase.RetornaProjetosCompletosInput{
		LoggedInUserKeycloakID: "keycloak-admin",
		ReuniaoID:              "reuniao-1",
	})

	if err != nil {
		t.Fatalf("esperava nil, got %v", err)
	}
	if len(result) != 0 {
		t.Errorf("esperava 0 projetos, got %d", len(result))
	}
}

func TestRetornaProjetosCompletos_ReuniaoNaoEncontrada(t *testing.T) {
	usuarioRepo := fakes.NewFakeUsuarioRepository()
	reuniaoRepo := fakes.NewFakeReuniaoRepository()

	usuarioRepo.Seed(adminUsuario("keycloak-admin", "user-admin"))

	uc := usecase.NewRetornaProjetosCompletosUseCase(usuarioRepo, reuniaoRepo)

	_, err := uc.Execute(context.Background(), usecase.RetornaProjetosCompletosInput{
		LoggedInUserKeycloakID: "keycloak-admin",
		ReuniaoID:              "reuniao-inexistente",
	})

	if err != domainVotacao.ErrReuniaoNotFound {
		t.Fatalf("esperava ErrReuniaoNotFound, got %v", err)
	}
}

func TestRetornaProjetosCompletos_UsuarioNaoEncontrado(t *testing.T) {
	usuarioRepo := fakes.NewFakeUsuarioRepository()
	reuniaoRepo := fakes.NewFakeReuniaoRepository()

	uc := usecase.NewRetornaProjetosCompletosUseCase(usuarioRepo, reuniaoRepo)

	_, err := uc.Execute(context.Background(), usecase.RetornaProjetosCompletosInput{
		LoggedInUserKeycloakID: "keycloak-inexistente",
		ReuniaoID:              "reuniao-1",
	})

	if err == nil {
		t.Fatal("esperava erro, got nil")
	}
}

func TestRetornaProjetosCompletos_UsuarioNaoEAdmin(t *testing.T) {
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

	uc := usecase.NewRetornaProjetosCompletosUseCase(usuarioRepo, reuniaoRepo)

	_, err := uc.Execute(context.Background(), usecase.RetornaProjetosCompletosInput{
		LoggedInUserKeycloakID: "keycloak-comum",
		ReuniaoID:              "reuniao-1",
	})

	if err != domainUsuario.ErrUserNotAdmin {
		t.Fatalf("esperava ErrUserNotAdmin, got %v", err)
	}
}

func TestRetornaProjetosCompletos_UsuarioInativo(t *testing.T) {
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

	uc := usecase.NewRetornaProjetosCompletosUseCase(usuarioRepo, reuniaoRepo)

	_, err := uc.Execute(context.Background(), usecase.RetornaProjetosCompletosInput{
		LoggedInUserKeycloakID: "keycloak-inativo",
		ReuniaoID:              "reuniao-1",
	})

	if err != domainUsuario.ErrUserNotActive {
		t.Fatalf("esperava ErrUserNotActive, got %v", err)
	}
}

func TestRetornaProjetosCompletos_ErroNoRepositorio(t *testing.T) {
	usuarioRepo := fakes.NewFakeUsuarioRepository()
	reuniaoRepo := fakes.NewFakeReuniaoRepository()

	usuarioRepo.Seed(adminUsuario("keycloak-admin", "user-admin"))
	reuniaoRepo.Seed(&domainVotacao.Reuniao{ID: "reuniao-1"})
	reuniaoRepo.GetProjetosCompletoErr = domainVotacao.ErrReuniaoNotFound

	uc := usecase.NewRetornaProjetosCompletosUseCase(usuarioRepo, reuniaoRepo)

	_, err := uc.Execute(context.Background(), usecase.RetornaProjetosCompletosInput{
		LoggedInUserKeycloakID: "keycloak-admin",
		ReuniaoID:              "reuniao-1",
	})

	if err == nil {
		t.Fatal("esperava erro, got nil")
	}
}
