package usuario_test

import (
	"context"
	"errors"
	"testing"

	"github.com/aleodoni/go-ddd/domain"
	usecase "github.com/aleodoni/voting-go/internal/application/usuario"
	domainUsuario "github.com/aleodoni/voting-go/internal/domain/usuario"
	"github.com/aleodoni/voting-go/internal/test/fakes"
)

func TestPesquisaUsuarios_Sucesso(t *testing.T) {
	repo := fakes.NewFakeUsuarioRepository()
	repo.Seed(adminAtivo())
	repo.Seed(&domainUsuario.Usuario{
		AggregateRoot: domain.NewAggregateRoot("user-2"),
		KeycloakID:    "keycloak-2",
		Nome:          "João Silva",
		Email:         "joao@exemplo.com",
		Credencial:    &domainUsuario.Credencial{Ativo: true},
	})

	uc := usecase.NewListUsuariosUseCase(repo)
	result, err := uc.Execute(context.Background(), usecase.ListUsuariosInput{
		LoggedInUserKeycloakID: "keycloak-admin",
		Page:                   1,
		Limit:                  20,
	})

	if err != nil {
		t.Fatalf("esperava nil, obteve: %v", err)
	}
	if result == nil {
		t.Fatal("esperava resultado, obteve nil")
	}
	if result.Total == 0 {
		t.Fatal("esperava usuários na lista")
	}
}

func TestPesquisaUsuarios_FiltroNome(t *testing.T) {
	repo := fakes.NewFakeUsuarioRepository()
	repo.Seed(adminAtivo())
	repo.Seed(&domainUsuario.Usuario{
		AggregateRoot: domain.NewAggregateRoot("user-joao"),
		KeycloakID:    "keycloak-joao",
		Nome:          "João Silva",
		Credencial:    &domainUsuario.Credencial{Ativo: true},
	})
	repo.Seed(&domainUsuario.Usuario{
		AggregateRoot: domain.NewAggregateRoot("user-maria"),
		KeycloakID:    "keycloak-maria",
		Nome:          "Maria Souza",
		Credencial:    &domainUsuario.Credencial{Ativo: true},
	})

	uc := usecase.NewListUsuariosUseCase(repo)
	result, err := uc.Execute(context.Background(), usecase.ListUsuariosInput{
		LoggedInUserKeycloakID: "keycloak-admin",
		Nome:                   "João",
		Page:                   1,
		Limit:                  20,
	})

	if err != nil {
		t.Fatalf("esperava nil, obteve: %v", err)
	}
	if result.Total != 1 {
		t.Fatalf("esperava 1 resultado, obteve %d", result.Total)
	}
}

// TestPesquisaUsuarios_PaginacaoInvalida verifica o bug corrigido:
// page < 1 e limit > 100 devem ser normalizados antes de chegar ao repositório.
func TestPesquisaUsuarios_PaginacaoInvalida(t *testing.T) {
	repo := fakes.NewFakeUsuarioRepository()
	repo.Seed(adminAtivo())

	uc := usecase.NewListUsuariosUseCase(repo)
	result, err := uc.Execute(context.Background(), usecase.ListUsuariosInput{
		LoggedInUserKeycloakID: "keycloak-admin",
		Page:                   0,   // inválido — deve ser normalizado para 1
		Limit:                  200, // inválido — deve ser normalizado para 20
	})

	if err != nil {
		t.Fatalf("esperava nil, obteve: %v", err)
	}
	// Verifica que o repositório recebeu os valores normalizados
	if len(repo.ListUsersCalls) == 0 {
		t.Fatal("esperava chamada ao repositório")
	}
	call := repo.ListUsersCalls[0]
	if call.Page != 1 {
		t.Fatalf("repositório recebeu Page=%d, esperava 1 (normalizado)", call.Page)
	}
	if call.Limit != 20 {
		t.Fatalf("repositório recebeu Limit=%d, esperava 20 (normalizado)", call.Limit)
	}
	// Verifica também os valores retornados
	if result.Page != 1 {
		t.Fatalf("resultado.Page=%d, esperava 1", result.Page)
	}
	if result.Limit != 20 {
		t.Fatalf("resultado.Limit=%d, esperava 20", result.Limit)
	}
}

func TestPesquisaUsuarios_ListarInativos(t *testing.T) {
	repo := fakes.NewFakeUsuarioRepository()
	repo.Seed(adminAtivo())
	repo.Seed(&domainUsuario.Usuario{
		AggregateRoot: domain.NewAggregateRoot("user-inativo"),
		KeycloakID:    "keycloak-inativo-2",
		Nome:          "Inativo",
		Credencial:    &domainUsuario.Credencial{Ativo: false},
	})

	uc := usecase.NewListUsuariosUseCase(repo)

	// Sem listar inativos — não deve aparecer o inativo
	semInativos, err := uc.Execute(context.Background(), usecase.ListUsuariosInput{
		LoggedInUserKeycloakID: "keycloak-admin",
		ListarInativos:         false,
		Page:                   1,
		Limit:                  20,
	})
	if err != nil {
		t.Fatalf("esperava nil, obteve: %v", err)
	}

	// Com listar inativos — deve aparecer
	comInativos, err := uc.Execute(context.Background(), usecase.ListUsuariosInput{
		LoggedInUserKeycloakID: "keycloak-admin",
		ListarInativos:         true,
		Page:                   1,
		Limit:                  20,
	})
	if err != nil {
		t.Fatalf("esperava nil, obteve: %v", err)
	}

	if comInativos.Total <= semInativos.Total {
		t.Fatal("esperava mais usuários ao listar inativos")
	}
}

func TestPesquisaUsuarios_AdminInativo(t *testing.T) {
	repo := fakes.NewFakeUsuarioRepository()
	repo.Seed(&domainUsuario.Usuario{
		AggregateRoot: domain.NewAggregateRoot("user-inativo"),
		KeycloakID:    "keycloak-inativo",
		Credencial:    &domainUsuario.Credencial{Ativo: false, PodeAdministrar: true},
	})

	uc := usecase.NewListUsuariosUseCase(repo)
	_, err := uc.Execute(context.Background(), usecase.ListUsuariosInput{
		LoggedInUserKeycloakID: "keycloak-inativo",
		Page:                   1,
		Limit:                  20,
	})

	if !errors.Is(err, domainUsuario.ErrUserNotActive) {
		t.Fatalf("esperava ErrUserNotActive, obteve: %v", err)
	}
}

func TestPesquisaUsuarios_AdminSemPermissao(t *testing.T) {
	repo := fakes.NewFakeUsuarioRepository()
	repo.Seed(&domainUsuario.Usuario{
		AggregateRoot: domain.NewAggregateRoot("user-comum"),
		KeycloakID:    "keycloak-comum",
		Credencial:    &domainUsuario.Credencial{Ativo: true, PodeAdministrar: false},
	})

	uc := usecase.NewListUsuariosUseCase(repo)
	_, err := uc.Execute(context.Background(), usecase.ListUsuariosInput{
		LoggedInUserKeycloakID: "keycloak-comum",
		Page:                   1,
		Limit:                  20,
	})

	if !errors.Is(err, domainUsuario.ErrUserNotAdmin) {
		t.Fatalf("esperava ErrUserNotAdmin, obteve: %v", err)
	}
}

func TestPesquisaUsuarios_ErroRepositorio(t *testing.T) {
	repo := fakes.NewFakeUsuarioRepository()
	repo.Seed(adminAtivo())
	erroEsperado := errors.New("falha no banco")
	repo.ListUsersErr = erroEsperado

	uc := usecase.NewListUsuariosUseCase(repo)
	_, err := uc.Execute(context.Background(), usecase.ListUsuariosInput{
		LoggedInUserKeycloakID: "keycloak-admin",
		Page:                   1,
		Limit:                  20,
	})

	if !errors.Is(err, erroEsperado) {
		t.Fatalf("esperava erro do repositório, obteve: %v", err)
	}
}
