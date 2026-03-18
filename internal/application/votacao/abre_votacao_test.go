package votacao_test

import (
	"context"
	"testing"

	ucVotacao "github.com/aleodoni/voting-go/internal/application/votacao"
	domainUsuario "github.com/aleodoni/voting-go/internal/domain/usuario"
	"github.com/aleodoni/voting-go/internal/domain/votacao"
	"github.com/aleodoni/voting-go/internal/platform/event"
	"github.com/aleodoni/voting-go/internal/test/fakes"
)

func TestAbreVotacao_Sucesso(t *testing.T) {
	usuarioRepo := fakes.NewFakeUsuarioRepository()
	reuniaoRepo := fakes.NewFakeReuniaoRepository()
	votacaoRepo := fakes.NewFakeVotacaoRepository()

	usuarioRepo.Seed(adminUsuario("keycloak-admin", "user-admin"))

	projeto := &votacao.Projeto{ID: "projeto-1", CodigoProposicao: "001"}
	reuniaoRepo.SeedProjetos("reuniao-1", []*votacao.Projeto{projeto})

	uc := ucVotacao.NewAbreVotacaoUseCase(usuarioRepo, reuniaoRepo, votacaoRepo, event.NewBus())

	err := uc.Execute(context.Background(), ucVotacao.AbreVotacaoInput{
		LoggedInUserKeycloakID: "keycloak-admin",
		ProjetoID:              "projeto-1",
	})

	if err != nil {
		t.Fatalf("esperava nil, got %v", err)
	}

	if len(reuniaoRepo.GetProjetoCompletoCalls) != 1 {
		t.Fatalf("esperava que GetProjetoCompleto fosse chamado")
	}
	if reuniaoRepo.GetProjetoCompletoCalls[0] != "projeto-1" {
		t.Errorf("esperava projeto-1, got %s", reuniaoRepo.GetProjetoCompletoCalls[0])
	}

	if len(votacaoRepo.SalvaVotacaoCalls) != 1 {
		t.Fatalf("esperava que SalvaVotacao fosse chamado")
	}
	if votacaoRepo.SalvaVotacaoCalls[0].ProjetoID == nil || *votacaoRepo.SalvaVotacaoCalls[0].ProjetoID != "projeto-1" {
		t.Errorf("esperava ProjetoID projeto-1, got %+v", votacaoRepo.SalvaVotacaoCalls[0])
	}
}

func TestAbreVotacao_UsuarioNaoEncontrado(t *testing.T) {
	usuarioRepo := fakes.NewFakeUsuarioRepository()
	reuniaoRepo := fakes.NewFakeReuniaoRepository()
	votacaoRepo := fakes.NewFakeVotacaoRepository()

	uc := ucVotacao.NewAbreVotacaoUseCase(usuarioRepo, reuniaoRepo, votacaoRepo, event.NewBus())

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
	votacaoRepo := fakes.NewFakeVotacaoRepository()

	usuarioRepo.Seed(&domainUsuario.Usuario{
		ID:         "user-comum",
		KeycloakID: "keycloak-comum",
		Credencial: &domainUsuario.Credencial{
			Ativo:           true,
			PodeAdministrar: false,
		},
	})

	uc := ucVotacao.NewAbreVotacaoUseCase(usuarioRepo, reuniaoRepo, votacaoRepo, event.NewBus())

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
	votacaoRepo := fakes.NewFakeVotacaoRepository()

	usuarioRepo.Seed(&domainUsuario.Usuario{
		ID:         "user-inativo",
		KeycloakID: "keycloak-inativo",
		Credencial: &domainUsuario.Credencial{
			Ativo:           false,
			PodeAdministrar: true,
		},
	})

	uc := ucVotacao.NewAbreVotacaoUseCase(usuarioRepo, reuniaoRepo, votacaoRepo, event.NewBus())

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
	votacaoRepo := fakes.NewFakeVotacaoRepository()

	usuarioRepo.Seed(adminUsuario("keycloak-admin", "user-admin"))

	uc := ucVotacao.NewAbreVotacaoUseCase(usuarioRepo, reuniaoRepo, votacaoRepo, event.NewBus())

	err := uc.Execute(context.Background(), ucVotacao.AbreVotacaoInput{
		LoggedInUserKeycloakID: "keycloak-admin",
		ProjetoID:              "projeto-inexistente",
	})

	if err != votacao.ErrProjetoNotFound {
		t.Fatalf("esperava ErrProjetoNotFound, got %v", err)
	}
}

func TestAbreVotacao_ErroSalvaVotacao(t *testing.T) {
	usuarioRepo := fakes.NewFakeUsuarioRepository()
	reuniaoRepo := fakes.NewFakeReuniaoRepository()
	votacaoRepo := fakes.NewFakeVotacaoRepository()

	usuarioRepo.Seed(adminUsuario("keycloak-admin", "user-admin"))

	projeto := &votacao.Projeto{ID: "projeto-1", CodigoProposicao: "001"}
	reuniaoRepo.SeedProjetos("reuniao-1", []*votacao.Projeto{projeto})

	votacaoRepo.SalvaVotacaoErr = votacao.ErrVotacaoAlreadyExists

	uc := ucVotacao.NewAbreVotacaoUseCase(usuarioRepo, reuniaoRepo, votacaoRepo, event.NewBus())

	err := uc.Execute(context.Background(), ucVotacao.AbreVotacaoInput{
		LoggedInUserKeycloakID: "keycloak-admin",
		ProjetoID:              "projeto-1",
	})

	if err != votacao.ErrVotacaoAlreadyExists {
		t.Fatalf("esperava ErrVotacaoAlreadyExists, got %v", err)
	}
}

func TestAbreVotacao_VotacaoJaAberta(t *testing.T) {
	usuarioRepo := fakes.NewFakeUsuarioRepository()
	reuniaoRepo := fakes.NewFakeReuniaoRepository()
	votacaoRepo := fakes.NewFakeVotacaoRepository()

	usuarioRepo.Seed(adminUsuario("keycloak-admin", "user-admin"))

	votacaoRepo.Seed(&votacao.Votacao{
		ID:     "votacao-existente",
		Status: votacao.StatusVotacaoA,
	})

	uc := ucVotacao.NewAbreVotacaoUseCase(usuarioRepo, reuniaoRepo, votacaoRepo, event.NewBus())

	err := uc.Execute(context.Background(), ucVotacao.AbreVotacaoInput{
		LoggedInUserKeycloakID: "keycloak-admin",
		ProjetoID:              "projeto-1",
	})

	if err != votacao.ErrVotacaoAberta {
		t.Fatalf("esperava ErrVotacaoAberta, got %v", err)
	}

	if len(votacaoRepo.SalvaVotacaoCalls) != 0 {
		t.Errorf("esperava que SalvaVotacao não fosse chamado, mas foi chamado %d vez(es)", len(votacaoRepo.SalvaVotacaoCalls))
	}
}

func TestAbreVotacao_ProjetoJaVotado(t *testing.T) {
	usuarioRepo := fakes.NewFakeUsuarioRepository()
	reuniaoRepo := fakes.NewFakeReuniaoRepository()
	votacaoRepo := fakes.NewFakeVotacaoRepository()

	usuarioRepo.Seed(adminUsuario("keycloak-admin", "user-admin"))

	votacaoExistente := &votacao.Votacao{
		ID:     "votacao-projeto",
		Status: votacao.StatusVotacaoA,
	}
	projeto := &votacao.Projeto{
		ID:               "projeto-1",
		CodigoProposicao: "001",
		Votacao:          votacaoExistente,
	}
	reuniaoRepo.SeedProjetos("reuniao-1", []*votacao.Projeto{projeto})

	uc := ucVotacao.NewAbreVotacaoUseCase(usuarioRepo, reuniaoRepo, votacaoRepo, event.NewBus())

	err := uc.Execute(context.Background(), ucVotacao.AbreVotacaoInput{
		LoggedInUserKeycloakID: "keycloak-admin",
		ProjetoID:              "projeto-1",
	})

	if err != votacao.ErrProjetoVoted {
		t.Fatalf("esperava ErrProjetoVoted, got %v", err)
	}

	if len(votacaoRepo.SalvaVotacaoCalls) != 0 {
		t.Errorf("esperava que SalvaVotacao não fosse chamado, mas foi chamado %d vez(es)", len(votacaoRepo.SalvaVotacaoCalls))
	}
}

func TestAbreVotacao_PublicaEventoAoAbrir(t *testing.T) {
	usuarioRepo := fakes.NewFakeUsuarioRepository()
	reuniaoRepo := fakes.NewFakeReuniaoRepository()
	votacaoRepo := fakes.NewFakeVotacaoRepository()

	usuarioRepo.Seed(adminUsuario("keycloak-admin", "user-admin"))

	projeto := &votacao.Projeto{ID: "projeto-1", CodigoProposicao: "001"}
	reuniaoRepo.SeedProjetos("reuniao-1", []*votacao.Projeto{projeto})

	bus := event.NewBus()
	ch := bus.Subscribe()
	defer bus.Unsubscribe(ch)

	uc := ucVotacao.NewAbreVotacaoUseCase(usuarioRepo, reuniaoRepo, votacaoRepo, bus)

	err := uc.Execute(context.Background(), ucVotacao.AbreVotacaoInput{
		LoggedInUserKeycloakID: "keycloak-admin",
		ProjetoID:              "projeto-1",
	})

	if err != nil {
		t.Fatalf("esperava nil, got %v", err)
	}

	select {
	case e := <-ch:
		if e.Type != event.VotacaoAberta {
			t.Errorf("esperava VotacaoAberta, got %s", e.Type)
		}
		payload, ok := e.Payload.(ucVotacao.AbreVotacaoPayload)
		if !ok {
			t.Fatal("payload com tipo incorreto")
		}
		if payload.ProjetoID != "projeto-1" {
			t.Errorf("esperava ProjetoID projeto-1, got %s", payload.ProjetoID)
		}
	default:
		t.Fatal("esperava evento publicado, nenhum recebido")
	}
}
