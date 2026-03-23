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

func usuarioAutenticado(keycloakID string) *domainUsuario.Usuario {
	return &domainUsuario.Usuario{
		AggregateRoot: domain.NewAggregateRoot(cuid2.Generate()),
		KeycloakID:    keycloakID,
		Username:      "vereador",
		Credencial: &domainUsuario.Credencial{
			Ativo:     true,
			PodeVotar: true,
		},
	}
}

func projetoComVotacaoAberta() *domainVotacao.Projeto {
	return &domainVotacao.Projeto{
		ID:               "projeto-1",
		Sumula:           "Projeto de teste",
		Relator:          "Relator Teste",
		CodigoProposicao: "001.00001.2026",
		Votacao: &domainVotacao.Votacao{
			AggregateRoot: domain.NewAggregateRoot("votacao-1"),
			Status:        domainVotacao.StatusVotacaoA,
		},
		Pareceres: &[]domainVotacao.Parecer{},
	}
}

func TestRetornaVotacaoAberta_UsuarioAutenticadoRetornaProjeto(t *testing.T) {
	usuarioRepo := fakes.NewFakeUsuarioRepository()
	votacaoRepo := fakes.NewFakeVotacaoRepository()

	usuarioRepo.Seed(usuarioAutenticado("keycloak-user"))
	votacaoRepo.SeedProjeto(projetoComVotacaoAberta())

	uc := usecase.NewRetornaVotacaoAbertaUseCase(usuarioRepo, votacaoRepo)

	result, err := uc.Execute(context.Background(), usecase.RetornaVotacaoAbertaInput{
		LoggedInUserKeycloakID: "keycloak-user",
	})

	if err != nil {
		t.Fatalf("esperava nil, got %v", err)
	}
	if result == nil {
		t.Fatal("esperava projeto, got nil")
	}
	if result.ID != "projeto-1" {
		t.Errorf("esperava projeto-1, got %s", result.ID)
	}
}

func TestRetornaVotacaoAberta_SemVotacaoAberta(t *testing.T) {
	usuarioRepo := fakes.NewFakeUsuarioRepository()
	votacaoRepo := fakes.NewFakeVotacaoRepository()

	usuarioRepo.Seed(usuarioAutenticado("keycloak-user"))
	// nenhum projeto seedado — retorna nil

	uc := usecase.NewRetornaVotacaoAbertaUseCase(usuarioRepo, votacaoRepo)

	result, err := uc.Execute(context.Background(), usecase.RetornaVotacaoAbertaInput{
		LoggedInUserKeycloakID: "keycloak-user",
	})

	if err != nil {
		t.Fatalf("esperava nil, got %v", err)
	}
	if result != nil {
		t.Errorf("esperava nil, got %v", result)
	}
}

func TestRetornaVotacaoAberta_UsuarioNaoEncontrado(t *testing.T) {
	usuarioRepo := fakes.NewFakeUsuarioRepository()
	votacaoRepo := fakes.NewFakeVotacaoRepository()

	uc := usecase.NewRetornaVotacaoAbertaUseCase(usuarioRepo, votacaoRepo)

	_, err := uc.Execute(context.Background(), usecase.RetornaVotacaoAbertaInput{
		LoggedInUserKeycloakID: "keycloak-inexistente",
	})

	if !errors.Is(err, domainUsuario.ErrUserNotFound) {
		t.Fatalf("esperava ErrUserNotFound, got %v", err)
	}
}

func TestRetornaVotacaoAberta_ErroNoRepositorio(t *testing.T) {
	usuarioRepo := fakes.NewFakeUsuarioRepository()
	votacaoRepo := fakes.NewFakeVotacaoRepository()

	usuarioRepo.Seed(usuarioAutenticado("keycloak-user"))
	votacaoRepo.GetProjetoVotacaoAbertaErr = domainVotacao.ErrVotacaoNaoEncontrada

	uc := usecase.NewRetornaVotacaoAbertaUseCase(usuarioRepo, votacaoRepo)

	_, err := uc.Execute(context.Background(), usecase.RetornaVotacaoAbertaInput{
		LoggedInUserKeycloakID: "keycloak-user",
	})

	if !errors.Is(err, domainVotacao.ErrVotacaoNaoEncontrada) {
		t.Fatalf("esperava ErrVotacaoNaoEncontrada, got %v", err)
	}
}
