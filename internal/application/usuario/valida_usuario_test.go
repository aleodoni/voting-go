package usuario_test

import (
	"context"
	"testing"

	"github.com/aleodoni/go-ddd/domain"
	usecase "github.com/aleodoni/voting-go/internal/application/usuario"
	domainUsuario "github.com/aleodoni/voting-go/internal/domain/usuario"
	"github.com/aleodoni/voting-go/internal/test/fakes"
	"github.com/nrednav/cuid2"
)

//
// helpers
//

func defaultInput() usecase.EnsureUsuarioInput {
	return usecase.EnsureUsuarioInput{
		KeycloakID: "keycloak-test",
		Username:   "user",
		Email:      "user@test.com",
		Nome:       "User Test",
	}
}

type fakeTransactor struct{}

func (f *fakeTransactor) Do(ctx context.Context, fn func(context.Context) error) error {
	return fn(ctx)
}

//
// tests
//

func TestEnsureUsuario_UsuarioJaExiste(t *testing.T) {

	usuarioRepo := fakes.NewFakeUsuarioRepository()

	usuarioRepo.Seed(&domainUsuario.Usuario{
		AggregateRoot: domain.NewAggregateRoot(cuid2.Generate()),
		KeycloakID:    "keycloak-123",
		Username:      "joao",
	})

	uc := usecase.NewEnsureUsuarioUseCase(usuarioRepo, &fakeTransactor{})

	input := defaultInput()
	input.KeycloakID = "keycloak-123"

	u, err := uc.Execute(context.Background(), input)

	if err != nil {
		t.Fatalf("esperava nil, got %v", err)
	}

	if u.KeycloakID != "keycloak-123" {
		t.Errorf("esperava keycloak-123, got %s", u.KeycloakID)
	}
}

func TestEnsureUsuario_UsuarioNaoExiste_CriaUsuarioECredencial(t *testing.T) {

	usuarioRepo := fakes.NewFakeUsuarioRepository()

	uc := usecase.NewEnsureUsuarioUseCase(usuarioRepo, &fakeTransactor{})

	input := defaultInput()
	input.KeycloakID = "keycloak-456"

	u, err := uc.Execute(context.Background(), input)

	if err != nil {
		t.Fatalf("esperava nil, got %v", err)
	}

	if u.Credencial == nil {
		t.Fatal("esperava credencial preenchida")
	}

	saved, err := usuarioRepo.FindByKeycloakID(context.Background(), "keycloak-456")
	if err != nil {
		t.Fatal("usuario não foi criado")
	}

	if saved.Credencial == nil {
		t.Fatal("credencial não foi criada")
	}

	if saved.Credencial.UsuarioID != saved.ID {
		t.Error("credencial não pertence ao usuario criado")
	}
}

func TestEnsureUsuario_CredencialComDefaultsCorretos(t *testing.T) {

	usuarioRepo := fakes.NewFakeUsuarioRepository()

	uc := usecase.NewEnsureUsuarioUseCase(usuarioRepo, &fakeTransactor{})

	u, err := uc.Execute(context.Background(), defaultInput())

	if err != nil {
		t.Fatalf("esperava nil, got %v", err)
	}

	if u.Credencial == nil {
		t.Fatal("credencial não criada")
	}

	if !u.Credencial.Ativo {
		t.Error("esperava Ativo = true")
	}

	if u.Credencial.PodeVotar {
		t.Error("esperava PodeVotar = false")
	}

	if u.Credencial.PodeAdministrar {
		t.Error("esperava PodeAdministrar = false")
	}
}

func TestEnsureUsuario_IDsGerados(t *testing.T) {

	usuarioRepo := fakes.NewFakeUsuarioRepository()

	uc := usecase.NewEnsureUsuarioUseCase(usuarioRepo, &fakeTransactor{})

	u, err := uc.Execute(context.Background(), defaultInput())

	if err != nil {
		t.Fatalf("esperava nil, got %v", err)
	}

	if u.ID == "" {
		t.Error("esperava ID do usuario preenchido")
	}

	if u.Credencial == nil {
		t.Fatal("credencial não criada")
	}

	if u.Credencial.ID == "" {
		t.Error("esperava ID da credencial preenchido")
	}
}
