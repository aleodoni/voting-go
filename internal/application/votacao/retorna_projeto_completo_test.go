package votacao_test

import (
	"context"
	"errors"
	"testing"

	"github.com/aleodoni/go-ddd/domain"
	ucVotacao "github.com/aleodoni/voting-go/internal/application/votacao"
	domainUsuario "github.com/aleodoni/voting-go/internal/domain/usuario"
	"github.com/aleodoni/voting-go/internal/domain/votacao"
	"github.com/aleodoni/voting-go/internal/test/fakes"
)

func TestRetornaProjetoCompleto_Sucesso(t *testing.T) {
	usuarioRepo := fakes.NewFakeUsuarioRepository()
	reuniaoRepo := fakes.NewFakeReuniaoRepository()

	usuarioRepo.Seed(adminUsuario("keycloak-admin", "user-admin"))
	reuniaoRepo.SeedProjetos("reuniao-1", []*votacao.Projeto{
		{ID: "projeto-1", CodigoProposicao: "001"},
	})

	uc := ucVotacao.NewRetornaProjetoCompletoUseCase(usuarioRepo, reuniaoRepo)
	result, err := uc.Execute(context.Background(), ucVotacao.RetornaProjetoCompletoInput{
		LoggedInUserKeycloakID: "keycloak-admin",
		ProjetoID:              "projeto-1",
	})

	if err != nil {
		t.Fatalf("esperava nil, obteve: %v", err)
	}
	if result == nil {
		t.Fatal("esperava projeto, obteve nil")
	}
	if result.ID != "projeto-1" {
		t.Fatalf("esperava ID 'projeto-1', obteve '%s'", result.ID)
	}
}

func TestRetornaProjetoCompleto_ProjetoNaoEncontrado(t *testing.T) {
	usuarioRepo := fakes.NewFakeUsuarioRepository()
	reuniaoRepo := fakes.NewFakeReuniaoRepository()

	usuarioRepo.Seed(adminUsuario("keycloak-admin", "user-admin"))

	uc := ucVotacao.NewRetornaProjetoCompletoUseCase(usuarioRepo, reuniaoRepo)
	_, err := uc.Execute(context.Background(), ucVotacao.RetornaProjetoCompletoInput{
		LoggedInUserKeycloakID: "keycloak-admin",
		ProjetoID:              "projeto-inexistente",
	})

	if !errors.Is(err, votacao.ErrProjetoNotFound) {
		t.Fatalf("esperava ErrProjetoNotFound, obteve: %v", err)
	}
}

func TestRetornaProjetoCompleto_AdminNaoEncontrado(t *testing.T) {
	usuarioRepo := fakes.NewFakeUsuarioRepository()
	reuniaoRepo := fakes.NewFakeReuniaoRepository()

	uc := ucVotacao.NewRetornaProjetoCompletoUseCase(usuarioRepo, reuniaoRepo)
	_, err := uc.Execute(context.Background(), ucVotacao.RetornaProjetoCompletoInput{
		LoggedInUserKeycloakID: "keycloak-inexistente",
		ProjetoID:              "qualquer",
	})

	if !errors.Is(err, domainUsuario.ErrUserNotFound) {
		t.Fatalf("esperava ErrUserNotFound, obteve: %v", err)
	}
}

func TestRetornaProjetoCompleto_AdminInativo(t *testing.T) {
	usuarioRepo := fakes.NewFakeUsuarioRepository()
	reuniaoRepo := fakes.NewFakeReuniaoRepository()

	usuarioRepo.Seed(&domainUsuario.Usuario{
		AggregateRoot: domain.NewAggregateRoot("user-inativo"),
		KeycloakID:    "keycloak-inativo",
		Credencial:    &domainUsuario.Credencial{Ativo: false, PodeAdministrar: true},
	})

	uc := ucVotacao.NewRetornaProjetoCompletoUseCase(usuarioRepo, reuniaoRepo)
	_, err := uc.Execute(context.Background(), ucVotacao.RetornaProjetoCompletoInput{
		LoggedInUserKeycloakID: "keycloak-inativo",
		ProjetoID:              "qualquer",
	})

	if !errors.Is(err, domainUsuario.ErrUserNotActive) {
		t.Fatalf("esperava ErrUserNotActive, obteve: %v", err)
	}
}

func TestRetornaProjetoCompleto_AdminSemPermissao(t *testing.T) {
	usuarioRepo := fakes.NewFakeUsuarioRepository()
	reuniaoRepo := fakes.NewFakeReuniaoRepository()

	usuarioRepo.Seed(&domainUsuario.Usuario{
		AggregateRoot: domain.NewAggregateRoot("user-comum"),
		KeycloakID:    "keycloak-comum",
		Credencial:    &domainUsuario.Credencial{Ativo: true, PodeAdministrar: false},
	})

	uc := ucVotacao.NewRetornaProjetoCompletoUseCase(usuarioRepo, reuniaoRepo)
	_, err := uc.Execute(context.Background(), ucVotacao.RetornaProjetoCompletoInput{
		LoggedInUserKeycloakID: "keycloak-comum",
		ProjetoID:              "qualquer",
	})

	if !errors.Is(err, domainUsuario.ErrUserNotAdmin) {
		t.Fatalf("esperava ErrUserNotAdmin, obteve: %v", err)
	}
}
