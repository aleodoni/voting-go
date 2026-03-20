package votacao_test

import (
	"context"
	"testing"

	"github.com/aleodoni/go-ddd/domain"
	ucVotacao "github.com/aleodoni/voting-go/internal/application/votacao"
	domainUsuario "github.com/aleodoni/voting-go/internal/domain/usuario"
	"github.com/aleodoni/voting-go/internal/domain/votacao"
	"github.com/aleodoni/voting-go/internal/platform/event"
	"github.com/aleodoni/voting-go/internal/test/fakes"
	"github.com/nrednav/cuid2"
)

func TestCancelaVotacao_Sucesso(t *testing.T) {
	usuarioRepo := fakes.NewFakeUsuarioRepository()
	reuniaoRepo := fakes.NewFakeReuniaoRepository()
	votacaoRepo := fakes.NewFakeVotacaoRepository()

	usuarioRepo.Seed(adminUsuario("keycloak-admin", "user-admin"))

	projetoID := "projeto-1"
	projeto := &votacao.Projeto{
		ID:               "projeto-1",
		CodigoProposicao: "001",
		Votacao: &votacao.Votacao{
			AggregateRoot: domain.NewAggregateRoot("votacao-1"),
			ProjetoID:     &projetoID,
			Status:        votacao.StatusVotacaoF,
		},
	}
	reuniaoRepo.SeedProjetos("reuniao-1", []*votacao.Projeto{projeto})

	uc := ucVotacao.NewCancelaVotacaoUseCase(usuarioRepo, reuniaoRepo, votacaoRepo, event.NewBus())

	err := uc.Execute(context.Background(), ucVotacao.CancelaVotacaoInput{
		LoggedInUserKeycloakID: "keycloak-admin",
		ProjetoID:              "projeto-1",
	})

	if err != nil {
		t.Fatalf("esperava nil, got %v", err)
	}

	if len(votacaoRepo.DeletaVotacaoCalls) != 1 {
		t.Fatalf("esperava que DeletaVotacao fosse chamado")
	}
	if votacaoRepo.DeletaVotacaoCalls[0] != "votacao-1" {
		t.Errorf("esperava votacao-1, got %s", votacaoRepo.DeletaVotacaoCalls[0])
	}
}

func TestCancelaVotacao_UsuarioNaoEncontrado(t *testing.T) {
	usuarioRepo := fakes.NewFakeUsuarioRepository()
	reuniaoRepo := fakes.NewFakeReuniaoRepository()
	votacaoRepo := fakes.NewFakeVotacaoRepository()

	uc := ucVotacao.NewCancelaVotacaoUseCase(usuarioRepo, reuniaoRepo, votacaoRepo, event.NewBus())

	err := uc.Execute(context.Background(), ucVotacao.CancelaVotacaoInput{
		LoggedInUserKeycloakID: "keycloak-inexistente",
		ProjetoID:              "projeto-1",
	})

	if err == nil {
		t.Fatal("esperava erro, got nil")
	}
}

func TestCancelaVotacao_UsuarioNaoAdmin(t *testing.T) {
	usuarioRepo := fakes.NewFakeUsuarioRepository()
	reuniaoRepo := fakes.NewFakeReuniaoRepository()
	votacaoRepo := fakes.NewFakeVotacaoRepository()

	usuarioRepo.Seed(&domainUsuario.Usuario{
		AggregateRoot: domain.NewAggregateRoot(cuid2.Generate()),
		KeycloakID:    "keycloak-comum",
		Credencial: &domainUsuario.Credencial{
			Ativo:           true,
			PodeAdministrar: false,
		},
	})

	uc := ucVotacao.NewCancelaVotacaoUseCase(usuarioRepo, reuniaoRepo, votacaoRepo, event.NewBus())

	err := uc.Execute(context.Background(), ucVotacao.CancelaVotacaoInput{
		LoggedInUserKeycloakID: "keycloak-comum",
		ProjetoID:              "projeto-1",
	})

	if err != domainUsuario.ErrUserNotAdmin {
		t.Fatalf("esperava ErrUserNotAdmin, got %v", err)
	}
}

func TestCancelaVotacao_UsuarioInativo(t *testing.T) {
	usuarioRepo := fakes.NewFakeUsuarioRepository()
	reuniaoRepo := fakes.NewFakeReuniaoRepository()
	votacaoRepo := fakes.NewFakeVotacaoRepository()

	usuarioRepo.Seed(&domainUsuario.Usuario{
		AggregateRoot: domain.NewAggregateRoot(cuid2.Generate()),
		KeycloakID:    "keycloak-inativo",
		Credencial: &domainUsuario.Credencial{
			Ativo:           false,
			PodeAdministrar: true,
		},
	})

	uc := ucVotacao.NewCancelaVotacaoUseCase(usuarioRepo, reuniaoRepo, votacaoRepo, event.NewBus())

	err := uc.Execute(context.Background(), ucVotacao.CancelaVotacaoInput{
		LoggedInUserKeycloakID: "keycloak-inativo",
		ProjetoID:              "projeto-1",
	})

	if err != domainUsuario.ErrUserNotActive {
		t.Fatalf("esperava ErrUserNotActive, got %v", err)
	}
}

func TestCancelaVotacao_ProjetoNaoEncontrado(t *testing.T) {
	usuarioRepo := fakes.NewFakeUsuarioRepository()
	reuniaoRepo := fakes.NewFakeReuniaoRepository()
	votacaoRepo := fakes.NewFakeVotacaoRepository()

	usuarioRepo.Seed(adminUsuario("keycloak-admin", "user-admin"))

	uc := ucVotacao.NewCancelaVotacaoUseCase(usuarioRepo, reuniaoRepo, votacaoRepo, event.NewBus())

	err := uc.Execute(context.Background(), ucVotacao.CancelaVotacaoInput{
		LoggedInUserKeycloakID: "keycloak-admin",
		ProjetoID:              "projeto-inexistente",
	})

	if err != votacao.ErrProjetoNotFound {
		t.Fatalf("esperava ErrProjetoNotFound, got %v", err)
	}
}

func TestCancelaVotacao_VotacaoNaoEncontrada(t *testing.T) {
	usuarioRepo := fakes.NewFakeUsuarioRepository()
	reuniaoRepo := fakes.NewFakeReuniaoRepository()
	votacaoRepo := fakes.NewFakeVotacaoRepository()

	usuarioRepo.Seed(adminUsuario("keycloak-admin", "user-admin"))

	projeto := &votacao.Projeto{ID: "projeto-1", CodigoProposicao: "001"}
	reuniaoRepo.SeedProjetos("reuniao-1", []*votacao.Projeto{projeto})

	uc := ucVotacao.NewCancelaVotacaoUseCase(usuarioRepo, reuniaoRepo, votacaoRepo, event.NewBus())

	err := uc.Execute(context.Background(), ucVotacao.CancelaVotacaoInput{
		LoggedInUserKeycloakID: "keycloak-admin",
		ProjetoID:              "projeto-1",
	})

	if err != votacao.ErrVotacaoNaoEncontrada {
		t.Fatalf("esperava ErrVotacaoNaoEncontrada, got %v", err)
	}
}

func TestCancelaVotacao_VotacaoNaoFechada(t *testing.T) {
	usuarioRepo := fakes.NewFakeUsuarioRepository()
	reuniaoRepo := fakes.NewFakeReuniaoRepository()
	votacaoRepo := fakes.NewFakeVotacaoRepository()

	usuarioRepo.Seed(adminUsuario("keycloak-admin", "user-admin"))

	// votação ainda aberta — não pode cancelar
	projeto := &votacao.Projeto{
		ID:               "projeto-1",
		CodigoProposicao: "001",
		Votacao: &votacao.Votacao{
			AggregateRoot: domain.NewAggregateRoot("votacao-1"),
			Status:        votacao.StatusVotacaoA,
		},
	}
	reuniaoRepo.SeedProjetos("reuniao-1", []*votacao.Projeto{projeto})

	uc := ucVotacao.NewCancelaVotacaoUseCase(usuarioRepo, reuniaoRepo, votacaoRepo, event.NewBus())

	err := uc.Execute(context.Background(), ucVotacao.CancelaVotacaoInput{
		LoggedInUserKeycloakID: "keycloak-admin",
		ProjetoID:              "projeto-1",
	})

	if err != votacao.ErrVotacaoNaoFechada {
		t.Fatalf("esperava ErrVotacaoNaoFechada, got %v", err)
	}
}

func TestCancelaVotacao_ErroDeletaVotacao(t *testing.T) {
	usuarioRepo := fakes.NewFakeUsuarioRepository()
	reuniaoRepo := fakes.NewFakeReuniaoRepository()
	votacaoRepo := fakes.NewFakeVotacaoRepository()

	usuarioRepo.Seed(adminUsuario("keycloak-admin", "user-admin"))

	projeto := &votacao.Projeto{
		ID:               "projeto-1",
		CodigoProposicao: "001",
		Votacao: &votacao.Votacao{
			AggregateRoot: domain.NewAggregateRoot("votacao-1"),
			Status:        votacao.StatusVotacaoF,
		},
	}
	reuniaoRepo.SeedProjetos("reuniao-1", []*votacao.Projeto{projeto})

	votacaoRepo.DeletaVotacaoErr = votacao.ErrVotacaoNaoEncontrada

	uc := ucVotacao.NewCancelaVotacaoUseCase(usuarioRepo, reuniaoRepo, votacaoRepo, event.NewBus())

	err := uc.Execute(context.Background(), ucVotacao.CancelaVotacaoInput{
		LoggedInUserKeycloakID: "keycloak-admin",
		ProjetoID:              "projeto-1",
	})

	if err != votacao.ErrVotacaoNaoEncontrada {
		t.Fatalf("esperava ErrVotacaoNaoEncontrada, got %v", err)
	}
}
