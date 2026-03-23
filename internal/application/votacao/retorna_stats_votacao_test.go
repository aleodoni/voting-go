package votacao_test

import (
	"context"
	"errors"
	"testing"

	"github.com/aleodoni/go-ddd/domain"
	usecase "github.com/aleodoni/voting-go/internal/application/votacao"
	domainUsuario "github.com/aleodoni/voting-go/internal/domain/usuario"
	domainVotacao "github.com/aleodoni/voting-go/internal/domain/votacao"
	"github.com/aleodoni/voting-go/internal/test/fakes"
	"github.com/nrednav/cuid2"
)

func TestRetornaVotingStats_AdminRetornaStats(t *testing.T) {
	usuarioRepo := fakes.NewFakeUsuarioRepository()
	votacaoRepo := fakes.NewFakeVotacaoRepository()

	usuarioRepo.Seed(adminUsuario("keycloak-admin", "user-admin"))
	votacaoRepo.SeedVotingStats(&domainVotacao.VotingStats{
		TotalProjects:      10,
		TotalVotedProjects: 4,
	})

	uc := usecase.NewRetornaVotingStatsUseCase(usuarioRepo, votacaoRepo)

	result, err := uc.Execute(context.Background(), usecase.RetornaVotingStatsInput{
		LoggedInUserKeycloakID: "keycloak-admin",
	})

	if err != nil {
		t.Fatalf("esperava nil, got %v", err)
	}
	if result.TotalProjects != 10 {
		t.Errorf("esperava 10, got %d", result.TotalProjects)
	}
	if result.TotalVotedProjects != 4 {
		t.Errorf("esperava 4, got %d", result.TotalVotedProjects)
	}
}

func TestRetornaVotingStats_UsuarioNaoEncontrado(t *testing.T) {
	usuarioRepo := fakes.NewFakeUsuarioRepository()
	votacaoRepo := fakes.NewFakeVotacaoRepository()

	uc := usecase.NewRetornaVotingStatsUseCase(usuarioRepo, votacaoRepo)

	_, err := uc.Execute(context.Background(), usecase.RetornaVotingStatsInput{
		LoggedInUserKeycloakID: "keycloak-inexistente",
	})

	if !errors.Is(err, domainUsuario.ErrUserNotFound) {
		t.Fatalf("esperava ErrUserNotFound, got %v", err)
	}
}

func TestRetornaVotingStats_UsuarioNaoEAdmin(t *testing.T) {
	usuarioRepo := fakes.NewFakeUsuarioRepository()
	votacaoRepo := fakes.NewFakeVotacaoRepository()

	usuarioRepo.Seed(&domainUsuario.Usuario{
		AggregateRoot: domain.NewAggregateRoot(cuid2.Generate()),
		KeycloakID:    "keycloak-comum",
		Credencial: &domainUsuario.Credencial{
			Ativo:           true,
			PodeAdministrar: false,
		},
	})

	uc := usecase.NewRetornaVotingStatsUseCase(usuarioRepo, votacaoRepo)

	_, err := uc.Execute(context.Background(), usecase.RetornaVotingStatsInput{
		LoggedInUserKeycloakID: "keycloak-comum",
	})

	if !errors.Is(err, domainUsuario.ErrUserNotAdmin) {
		t.Fatalf("esperava ErrUserNotAdmin, got %v", err)
	}
}

func TestRetornaVotingStats_UsuarioInativo(t *testing.T) {
	usuarioRepo := fakes.NewFakeUsuarioRepository()
	votacaoRepo := fakes.NewFakeVotacaoRepository()

	usuarioRepo.Seed(&domainUsuario.Usuario{
		AggregateRoot: domain.NewAggregateRoot(cuid2.Generate()),
		KeycloakID:    "keycloak-inativo",
		Credencial: &domainUsuario.Credencial{
			Ativo:           false,
			PodeAdministrar: true,
		},
	})

	uc := usecase.NewRetornaVotingStatsUseCase(usuarioRepo, votacaoRepo)

	_, err := uc.Execute(context.Background(), usecase.RetornaVotingStatsInput{
		LoggedInUserKeycloakID: "keycloak-inativo",
	})

	if !errors.Is(err, domainUsuario.ErrUserNotActive) {
		t.Fatalf("esperava ErrUserNotActive, got %v", err)
	}
}

func TestRetornaVotingStats_ErroNoRepositorio(t *testing.T) {
	usuarioRepo := fakes.NewFakeUsuarioRepository()
	votacaoRepo := fakes.NewFakeVotacaoRepository()

	usuarioRepo.Seed(adminUsuario("keycloak-admin", "user-admin"))
	votacaoRepo.GetVotingStatsErr = errors.New("erro de banco")

	uc := usecase.NewRetornaVotingStatsUseCase(usuarioRepo, votacaoRepo)

	_, err := uc.Execute(context.Background(), usecase.RetornaVotingStatsInput{
		LoggedInUserKeycloakID: "keycloak-admin",
	})

	if err == nil {
		t.Fatal("esperava erro, got nil")
	}
}
