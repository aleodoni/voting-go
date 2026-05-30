package votacao_test

import (
	"context"
	"errors"
	"testing"

	"github.com/aleodoni/go-ddd/domain"
	ucVotacao "github.com/aleodoni/voting-go/internal/application/votacao"
	domainUsuario "github.com/aleodoni/voting-go/internal/domain/usuario"
	"github.com/aleodoni/voting-go/internal/domain/votacao"
	"github.com/aleodoni/voting-go/internal/platform/event"
	"github.com/aleodoni/voting-go/internal/test/fakes"
	"github.com/nrednav/cuid2"
)

func TestFechaVotacao_Sucesso(t *testing.T) {
	usuarioRepo := fakes.NewFakeUsuarioRepository()
	reuniaoRepo := fakes.NewFakeReuniaoRepository()
	votacaoRepo := fakes.NewFakeVotacaoRepository()

	usuarioRepo.Seed(adminUsuario("keycloak-admin", "user-admin"))

	projeto := &votacao.Projeto{
		ID:               "projeto-1",
		CodigoProposicao: "001",
		Votacao: &votacao.Votacao{
			AggregateRoot: domain.NewAggregateRoot("votacao-1"),
			Status:        votacao.StatusVotacaoA,
		},
	}
	reuniaoRepo.SeedProjetos("reuniao-1", []*votacao.Projeto{projeto})

	uc := ucVotacao.NewFechaVotacaoUseCase(usuarioRepo, reuniaoRepo, votacaoRepo, event.NewBus())

	err := uc.Execute(context.Background(), ucVotacao.FechaVotacaoInput{
		LoggedInUserKeycloakID: "keycloak-admin",
		ProjetoID:              "projeto-1",
	})

	if !errors.Is(err, nil) {
		t.Fatalf("esperava nil, got %v", err)
	}

	if len(votacaoRepo.SalvaVotacaoCalls) != 1 {
		t.Fatalf("esperava que SalvaVotacao fosse chamado")
	}
	if votacaoRepo.SalvaVotacaoCalls[0].Status != votacao.StatusVotacaoV {
		t.Errorf("esperava status V, got %s", votacaoRepo.SalvaVotacaoCalls[0].Status)
	}
}

func TestFechaVotacao_UsuarioNaoEncontrado(t *testing.T) {
	usuarioRepo := fakes.NewFakeUsuarioRepository()
	reuniaoRepo := fakes.NewFakeReuniaoRepository()
	votacaoRepo := fakes.NewFakeVotacaoRepository()

	uc := ucVotacao.NewFechaVotacaoUseCase(usuarioRepo, reuniaoRepo, votacaoRepo, event.NewBus())

	err := uc.Execute(context.Background(), ucVotacao.FechaVotacaoInput{
		LoggedInUserKeycloakID: "keycloak-inexistente",
		ProjetoID:              "projeto-1",
	})

	if err == nil {
		t.Fatal("esperava erro, got nil")
	}
}

func TestFechaVotacao_UsuarioNaoAdmin(t *testing.T) {
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

	uc := ucVotacao.NewFechaVotacaoUseCase(usuarioRepo, reuniaoRepo, votacaoRepo, event.NewBus())

	err := uc.Execute(context.Background(), ucVotacao.FechaVotacaoInput{
		LoggedInUserKeycloakID: "keycloak-comum",
		ProjetoID:              "projeto-1",
	})

	if !errors.Is(err, domainUsuario.ErrUserNotAdmin) {
		t.Fatalf("esperava ErrUserNotAdmin, got %v", err)
	}
}

func TestFechaVotacao_UsuarioInativo(t *testing.T) {
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

	uc := ucVotacao.NewFechaVotacaoUseCase(usuarioRepo, reuniaoRepo, votacaoRepo, event.NewBus())

	err := uc.Execute(context.Background(), ucVotacao.FechaVotacaoInput{
		LoggedInUserKeycloakID: "keycloak-inativo",
		ProjetoID:              "projeto-1",
	})

	if !errors.Is(err, domainUsuario.ErrUserNotActive) {
		t.Fatalf("esperava ErrUserNotActive, got %v", err)
	}
}

func TestFechaVotacao_ProjetoNaoEncontrado(t *testing.T) {
	usuarioRepo := fakes.NewFakeUsuarioRepository()
	reuniaoRepo := fakes.NewFakeReuniaoRepository()
	votacaoRepo := fakes.NewFakeVotacaoRepository()

	usuarioRepo.Seed(adminUsuario("keycloak-admin", "user-admin"))

	uc := ucVotacao.NewFechaVotacaoUseCase(usuarioRepo, reuniaoRepo, votacaoRepo, event.NewBus())

	err := uc.Execute(context.Background(), ucVotacao.FechaVotacaoInput{
		LoggedInUserKeycloakID: "keycloak-admin",
		ProjetoID:              "projeto-inexistente",
	})

	if !errors.Is(err, votacao.ErrProjetoNotFound) {
		t.Fatalf("esperava ErrProjetoNotFound, got %v", err)
	}
}

func TestFechaVotacao_VotacaoNaoEncontrada(t *testing.T) {
	usuarioRepo := fakes.NewFakeUsuarioRepository()
	reuniaoRepo := fakes.NewFakeReuniaoRepository()
	votacaoRepo := fakes.NewFakeVotacaoRepository()

	usuarioRepo.Seed(adminUsuario("keycloak-admin", "user-admin"))

	// projeto sem votação
	projeto := &votacao.Projeto{ID: "projeto-1", CodigoProposicao: "001"}
	reuniaoRepo.SeedProjetos("reuniao-1", []*votacao.Projeto{projeto})

	uc := ucVotacao.NewFechaVotacaoUseCase(usuarioRepo, reuniaoRepo, votacaoRepo, event.NewBus())

	err := uc.Execute(context.Background(), ucVotacao.FechaVotacaoInput{
		LoggedInUserKeycloakID: "keycloak-admin",
		ProjetoID:              "projeto-1",
	})

	if !errors.Is(err, votacao.ErrVotacaoNaoEncontrada) {
		t.Fatalf("esperava ErrVotacaoNaoEncontrada, got %v", err)
	}
}

func TestFechaVotacao_VotacaoNaoAberta(t *testing.T) {
	usuarioRepo := fakes.NewFakeUsuarioRepository()
	reuniaoRepo := fakes.NewFakeReuniaoRepository()
	votacaoRepo := fakes.NewFakeVotacaoRepository()

	usuarioRepo.Seed(adminUsuario("keycloak-admin", "user-admin"))

	// votação já fechada
	projeto := &votacao.Projeto{
		ID:               "projeto-1",
		CodigoProposicao: "001",
		Votacao: &votacao.Votacao{
			AggregateRoot: domain.NewAggregateRoot("votacao-1"),
			Status:        votacao.StatusVotacaoF,
		},
	}
	reuniaoRepo.SeedProjetos("reuniao-1", []*votacao.Projeto{projeto})

	uc := ucVotacao.NewFechaVotacaoUseCase(usuarioRepo, reuniaoRepo, votacaoRepo, event.NewBus())

	err := uc.Execute(context.Background(), ucVotacao.FechaVotacaoInput{
		LoggedInUserKeycloakID: "keycloak-admin",
		ProjetoID:              "projeto-1",
	})

	if !errors.Is(err, votacao.ErrVotacaoNaoAberta) {
		t.Fatalf("esperava ErrVotacaoNaoAberta, got %v", err)
	}
}

func TestFechaVotacao_ErroSalvaVotacao(t *testing.T) {
	usuarioRepo := fakes.NewFakeUsuarioRepository()
	reuniaoRepo := fakes.NewFakeReuniaoRepository()
	votacaoRepo := fakes.NewFakeVotacaoRepository()

	usuarioRepo.Seed(adminUsuario("keycloak-admin", "user-admin"))

	projeto := &votacao.Projeto{
		ID:               "projeto-1",
		CodigoProposicao: "001",
		Votacao: &votacao.Votacao{
			AggregateRoot: domain.NewAggregateRoot("votacao-1"),
			Status:        votacao.StatusVotacaoA,
		},
	}
	reuniaoRepo.SeedProjetos("reuniao-1", []*votacao.Projeto{projeto})

	votacaoRepo.SalvaVotacaoErr = votacao.ErrVotacaoNaoEncontrada

	uc := ucVotacao.NewFechaVotacaoUseCase(usuarioRepo, reuniaoRepo, votacaoRepo, event.NewBus())

	err := uc.Execute(context.Background(), ucVotacao.FechaVotacaoInput{
		LoggedInUserKeycloakID: "keycloak-admin",
		ProjetoID:              "projeto-1",
	})

	if !errors.Is(err, votacao.ErrVotacaoNaoEncontrada) {
		t.Fatalf("esperava ErrVotacaoNaoEncontrada, got %v", err)
	}
}

func TestFechaVotacao_PublicaEvento(t *testing.T) {
	usuarioRepo := fakes.NewFakeUsuarioRepository()
	reuniaoRepo := fakes.NewFakeReuniaoRepository()
	votacaoRepo := fakes.NewFakeVotacaoRepository()

	usuarioRepo.Seed(adminUsuario("keycloak-admin", "user-admin"))

	projeto := &votacao.Projeto{
		ID:               "projeto-1",
		CodigoProposicao: "001",
		Votacao: &votacao.Votacao{
			AggregateRoot: domain.NewAggregateRoot("votacao-1"),
			Status:        votacao.StatusVotacaoA,
		},
	}
	reuniaoRepo.SeedProjetos("reuniao-1", []*votacao.Projeto{projeto})

	bus := event.NewBus()
	ch := bus.Subscribe("test-user", "test", false)
	defer bus.Unsubscribe(ch)

	uc := ucVotacao.NewFechaVotacaoUseCase(usuarioRepo, reuniaoRepo, votacaoRepo, bus)
	if err := uc.Execute(context.Background(), ucVotacao.FechaVotacaoInput{
		LoggedInUserKeycloakID: "keycloak-admin",
		ProjetoID:              "projeto-1",
	}); err != nil {
		t.Fatalf("esperava nil, obteve: %v", err)
	}

	select {
	case e := <-ch:
		if e.Type != event.VotacaoFechada {
			t.Errorf("esperava VotacaoFechada, obteve %s", e.Type)
		}
	default:
		t.Fatal("esperava evento publicado, nenhum recebido")
	}
}
