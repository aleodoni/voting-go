package votacao_test

import (
	"context"
	"errors"
	"testing"

	ucVotacao "github.com/aleodoni/voting-go/internal/application/votacao"
	domainUsuario "github.com/aleodoni/voting-go/internal/domain/usuario"
	"github.com/aleodoni/voting-go/internal/domain/votacao"
	"github.com/aleodoni/voting-go/internal/platform/event"
	"github.com/aleodoni/voting-go/internal/test/fakes"
)

// helper para evitar repetição nos testes
func setupUsuarioVereador(repo *fakes.FakeUsuarioRepository) {
	repo.Seed(&domainUsuario.Usuario{
		ID:         "user-vereador",
		KeycloakID: "keycloak-vereador",
		Credencial: &domainUsuario.Credencial{
			Ativo:     true,
			PodeVotar: true,
		},
	})
}

func setupVotacaoAberta(t *testing.T, repo *fakes.FakeVotacaoRepository) {
	t.Helper()
	if err := repo.SalvaVotacao(context.Background(), &votacao.Votacao{
		ID:     "votacao-1",
		Status: votacao.StatusVotacaoA,
	}); err != nil {
		t.Fatalf("setupVotacaoAberta: %v", err)
	}
}

// ── testes existentes (corrigidos com seed da votação) ────────────────────────

func TestRegistraVoto_Sucesso(t *testing.T) {
	usuarioRepo := fakes.NewFakeUsuarioRepository()
	votacaoRepo := fakes.NewFakeVotacaoRepository()
	setupUsuarioVereador(usuarioRepo)
	setupVotacaoAberta(t, votacaoRepo)

	uc := ucVotacao.NewRegistraVotoUseCase(usuarioRepo, votacaoRepo, event.NewBus())

	err := uc.Execute(context.Background(), ucVotacao.RegistraVotoInput{
		LoggedInUserKeycloakID: "keycloak-vereador",
		VotacaoID:              "votacao-1",
		Voto:                   votacao.OpcaoVotoF,
	})

	if err != nil {
		t.Fatalf("esperava nil, got %v", err)
	}
	if len(votacaoRepo.SalvaVotoCalls) != 1 {
		t.Fatalf("esperava que SalvaVoto fosse chamado")
	}

	voto := votacaoRepo.SalvaVotoCalls[0]
	if voto.VotacaoID != "votacao-1" {
		t.Errorf("esperava votacao-1, got %s", voto.VotacaoID)
	}
	if voto.UsuarioID != "user-vereador" {
		t.Errorf("esperava user-vereador, got %s", voto.UsuarioID)
	}
	if voto.Voto != votacao.OpcaoVotoF {
		t.Errorf("esperava OpcaoVotoF, got %s", voto.Voto)
	}
}

func TestRegistraVoto_ComRestricao(t *testing.T) {
	usuarioRepo := fakes.NewFakeUsuarioRepository()
	votacaoRepo := fakes.NewFakeVotacaoRepository()
	setupUsuarioVereador(usuarioRepo)
	setupVotacaoAberta(t, votacaoRepo)

	uc := ucVotacao.NewRegistraVotoUseCase(usuarioRepo, votacaoRepo, event.NewBus())

	err := uc.Execute(context.Background(), ucVotacao.RegistraVotoInput{
		LoggedInUserKeycloakID: "keycloak-vereador",
		VotacaoID:              "votacao-1",
		Voto:                   votacao.OpcaoVotoR,
		Restricao:              &votacao.Restricao{Restricao: "minha restrição"},
	})

	if err != nil {
		t.Fatalf("esperava nil, got %v", err)
	}

	voto := votacaoRepo.SalvaVotoCalls[0]
	if voto.Restricao == nil {
		t.Fatal("esperava restricao preenchida")
	}
	if voto.Restricao.ID == "" {
		t.Error("esperava ID gerado para restricao")
	}
	if voto.Restricao.Restricao != "minha restrição" {
		t.Errorf("esperava 'minha restrição', got %s", voto.Restricao.Restricao)
	}
}

func TestRegistraVoto_ComVotoContrario(t *testing.T) {
	usuarioRepo := fakes.NewFakeUsuarioRepository()
	votacaoRepo := fakes.NewFakeVotacaoRepository()
	setupUsuarioVereador(usuarioRepo)
	setupVotacaoAberta(t, votacaoRepo)

	uc := ucVotacao.NewRegistraVotoUseCase(usuarioRepo, votacaoRepo, event.NewBus())

	err := uc.Execute(context.Background(), ucVotacao.RegistraVotoInput{
		LoggedInUserKeycloakID: "keycloak-vereador",
		VotacaoID:              "votacao-1",
		Voto:                   votacao.OpcaoVotoC,
		VotoContrario: &votacao.VotoContrario{
			IDTexto:   42,
			ParecerID: "parecer-1",
		},
	})

	if err != nil {
		t.Fatalf("esperava nil, got %v", err)
	}

	voto := votacaoRepo.SalvaVotoCalls[0]
	if voto.VotoContrario == nil {
		t.Fatal("esperava voto contrario preenchido")
	}
	if voto.VotoContrario.ID == "" {
		t.Error("esperava ID gerado para voto contrario")
	}
	if voto.VotoContrario.IDTexto != 42 {
		t.Errorf("esperava IDTexto 42, got %d", voto.VotoContrario.IDTexto)
	}
	if voto.VotoContrario.ParecerID != "parecer-1" {
		t.Errorf("esperava parecer-1, got %s", voto.VotoContrario.ParecerID)
	}
}

func TestRegistraVoto_UsuarioNaoEncontrado(t *testing.T) {
	usuarioRepo := fakes.NewFakeUsuarioRepository()
	votacaoRepo := fakes.NewFakeVotacaoRepository()
	// sem seed de usuário — não precisa de votação seedada pois falha antes

	uc := ucVotacao.NewRegistraVotoUseCase(usuarioRepo, votacaoRepo, event.NewBus())

	err := uc.Execute(context.Background(), ucVotacao.RegistraVotoInput{
		LoggedInUserKeycloakID: "keycloak-inexistente",
		VotacaoID:              "votacao-1",
		Voto:                   votacao.OpcaoVotoF,
	})

	if err == nil {
		t.Fatal("esperava erro, got nil")
	}
}

func TestRegistraVoto_ErroSalvaVoto(t *testing.T) {
	usuarioRepo := fakes.NewFakeUsuarioRepository()
	votacaoRepo := fakes.NewFakeVotacaoRepository()
	setupUsuarioVereador(usuarioRepo)
	setupVotacaoAberta(t, votacaoRepo)

	votacaoRepo.SalvaVotoErr = errors.New("db error")

	uc := ucVotacao.NewRegistraVotoUseCase(usuarioRepo, votacaoRepo, event.NewBus())

	err := uc.Execute(context.Background(), ucVotacao.RegistraVotoInput{
		LoggedInUserKeycloakID: "keycloak-vereador",
		VotacaoID:              "votacao-1",
		Voto:                   votacao.OpcaoVotoF,
	})

	if err == nil {
		t.Fatal("esperava erro, got nil")
	}
}

func TestRegistraVoto_VotacaoNaoEncontrada(t *testing.T) {
	usuarioRepo := fakes.NewFakeUsuarioRepository()
	votacaoRepo := fakes.NewFakeVotacaoRepository()
	setupUsuarioVereador(usuarioRepo)

	uc := ucVotacao.NewRegistraVotoUseCase(usuarioRepo, votacaoRepo, event.NewBus())

	err := uc.Execute(context.Background(), ucVotacao.RegistraVotoInput{
		LoggedInUserKeycloakID: "keycloak-vereador",
		VotacaoID:              "votacao-inexistente",
		Voto:                   votacao.OpcaoVotoF,
	})

	if !errors.Is(err, votacao.ErrVotacaoNaoEncontrada) {
		t.Fatalf("esperava ErrVotacaoNaoEncontrada, got %v", err)
	}
}

func TestRegistraVoto_VotacaoFechada(t *testing.T) {
	usuarioRepo := fakes.NewFakeUsuarioRepository()
	votacaoRepo := fakes.NewFakeVotacaoRepository()
	setupUsuarioVereador(usuarioRepo)

	votacaoRepo.SalvaVotacao(context.Background(), &votacao.Votacao{
		ID:     "votacao-1",
		Status: votacao.StatusVotacaoF, // ← fechada
	})

	uc := ucVotacao.NewRegistraVotoUseCase(usuarioRepo, votacaoRepo, event.NewBus())

	err := uc.Execute(context.Background(), ucVotacao.RegistraVotoInput{
		LoggedInUserKeycloakID: "keycloak-vereador",
		VotacaoID:              "votacao-1",
		Voto:                   votacao.OpcaoVotoF,
	})

	if !errors.Is(err, votacao.ErrVotacaoNaoAberta) {
		t.Fatalf("esperava ErrVotacaoNaoAberta, got %v", err)
	}
}

func TestRegistraVoto_UsuarioJaVotou(t *testing.T) {
	usuarioRepo := fakes.NewFakeUsuarioRepository()
	votacaoRepo := fakes.NewFakeVotacaoRepository()
	setupUsuarioVereador(usuarioRepo)
	setupVotacaoAberta(t, votacaoRepo)

	uc := ucVotacao.NewRegistraVotoUseCase(usuarioRepo, votacaoRepo, event.NewBus())

	input := ucVotacao.RegistraVotoInput{
		LoggedInUserKeycloakID: "keycloak-vereador",
		VotacaoID:              "votacao-1",
		Voto:                   votacao.OpcaoVotoF,
	}

	// primeiro voto — deve passar
	if err := uc.Execute(context.Background(), input); err != nil {
		t.Fatalf("primeiro voto: esperava nil, got %v", err)
	}

	// segundo voto — deve falhar
	err := uc.Execute(context.Background(), input)
	if !errors.Is(err, votacao.ErrUsuarioJaVotou) {
		t.Fatalf("esperava ErrUsuarioJaVotou, got %v", err)
	}
}
