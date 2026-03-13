package votacao_test

import (
	"context"
	"testing"

	ucVotacao "github.com/aleodoni/voting-go/internal/application/votacao"
	domainUsuario "github.com/aleodoni/voting-go/internal/domain/usuario"
	"github.com/aleodoni/voting-go/internal/domain/votacao"
	"github.com/aleodoni/voting-go/internal/test/fakes"
)

// helper para criar usuário admin

func TestAbreVotacao_Sucesso(t *testing.T) {
	usuarioRepo := fakes.NewFakeUsuarioRepository()
	reuniaoRepo := fakes.NewFakeReuniaoRepository()

	// adiciona usuário admin
	usuarioRepo.Seed(adminUsuario("keycloak-admin", "user-admin"))

	// adiciona um projeto no fake
	projeto := &votacao.Projeto{ID: "projeto-1", CodigoProposicao: "001"}
	reuniaoRepo.SeedProjetos("reuniao-1", []*votacao.Projeto{projeto})

	uc := ucVotacao.NewAbreVotacaoUseCase(usuarioRepo, reuniaoRepo)

	err := uc.Execute(context.Background(), ucVotacao.AbreVotacaoInput{
		LoggedInUserKeycloakID: "keycloak-admin",
		ProjetoID:              "projeto-1",
	})

	if err != nil {
		t.Fatalf("esperava nil, got %v", err)
	}

	// verifica se GetProjetoCompleto foi chamado
	if len(reuniaoRepo.GetProjetoCompletoCalls) != 1 {
		t.Fatalf("esperava que GetProjetoCompleto fosse chamado")
	}
	if reuniaoRepo.GetProjetoCompletoCalls[0] != "projeto-1" {
		t.Errorf("esperava que GetProjetoCompleto fosse chamado com projeto-1, got %s", reuniaoRepo.GetProjetoCompletoCalls[0])
	}

	// verifica se CriaVotacao foi chamado
	if len(reuniaoRepo.CriaVotacaoCalls) != 1 {
		t.Fatalf("esperava que CriaVotacao fosse chamado")
	}
	if reuniaoRepo.CriaVotacaoCalls[0].ProjetoID == nil || *reuniaoRepo.CriaVotacaoCalls[0].ProjetoID != "projeto-1" {
		t.Errorf("esperava que CriaVotacao fosse criado com projeto-1, got %+v", reuniaoRepo.CriaVotacaoCalls[0])
	}
}

func TestAbreVotacao_UsuarioNaoEncontrado(t *testing.T) {
	usuarioRepo := fakes.NewFakeUsuarioRepository()
	reuniaoRepo := fakes.NewFakeReuniaoRepository()

	uc := ucVotacao.NewAbreVotacaoUseCase(usuarioRepo, reuniaoRepo)

	err := uc.Execute(context.Background(), ucVotacao.AbreVotacaoInput{
		LoggedInUserKeycloakID: "keycloak-inexistente",
		ProjetoID:              "projeto-1",
	})

	if err == nil {
		t.Fatal("esperava erro, got nil")
	}
}

func TestAbreVotacao_UsuarioNaoAdmin(t *testing.T) {
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

	uc := ucVotacao.NewAbreVotacaoUseCase(usuarioRepo, reuniaoRepo)

	err := uc.Execute(context.Background(), ucVotacao.AbreVotacaoInput{
		LoggedInUserKeycloakID: "keycloak-comum",
		ProjetoID:              "projeto-1",
	})

	if err != domainUsuario.ErrUserNotAdmin {
		t.Fatalf("esperava ErrUserNotAdmin, got %v", err)
	}
}

func TestAbreVotacao_UsuarioInativo(t *testing.T) {
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

	uc := ucVotacao.NewAbreVotacaoUseCase(usuarioRepo, reuniaoRepo)

	err := uc.Execute(context.Background(), ucVotacao.AbreVotacaoInput{
		LoggedInUserKeycloakID: "keycloak-inativo",
		ProjetoID:              "projeto-1",
	})

	if err != domainUsuario.ErrUserNotActive {
		t.Fatalf("esperava ErrUserNotActive, got %v", err)
	}
}

func TestAbreVotacao_ProjetoNaoEncontrado(t *testing.T) {
	usuarioRepo := fakes.NewFakeUsuarioRepository()
	reuniaoRepo := fakes.NewFakeReuniaoRepository()

	usuarioRepo.Seed(adminUsuario("keycloak-admin", "user-admin"))

	uc := ucVotacao.NewAbreVotacaoUseCase(usuarioRepo, reuniaoRepo)

	err := uc.Execute(context.Background(), ucVotacao.AbreVotacaoInput{
		LoggedInUserKeycloakID: "keycloak-admin",
		ProjetoID:              "projeto-inexistente",
	})

	if err != votacao.ErrProjetoNotFound {
		t.Fatalf("esperava ErrProjetoNotFound, got %v", err)
	}
}

func TestAbreVotacao_ErroCriaVotacao(t *testing.T) {
	usuarioRepo := fakes.NewFakeUsuarioRepository()
	reuniaoRepo := fakes.NewFakeReuniaoRepository()

	usuarioRepo.Seed(adminUsuario("keycloak-admin", "user-admin"))

	// adiciona projeto
	projeto := &votacao.Projeto{ID: "projeto-1", CodigoProposicao: "001"}
	reuniaoRepo.SeedProjetos("reuniao-1", []*votacao.Projeto{projeto})

	// força erro ao criar votação
	reuniaoRepo.CriaVotacaoErr = votacao.ErrVotacaoAlreadyExists

	uc := ucVotacao.NewAbreVotacaoUseCase(usuarioRepo, reuniaoRepo)

	err := uc.Execute(context.Background(), ucVotacao.AbreVotacaoInput{
		LoggedInUserKeycloakID: "keycloak-admin",
		ProjetoID:              "projeto-1",
	})

	if err != votacao.ErrVotacaoAlreadyExists {
		t.Fatalf("esperava ErrVotacaoAlreadyExists, got %v", err)
	}
}
