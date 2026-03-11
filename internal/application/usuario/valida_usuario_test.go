package usuario_test

import (
	"context"
	"testing"

	usecase "github.com/aleodoni/voting-go/internal/application/usuario"
	domainUsuario "github.com/aleodoni/voting-go/internal/domain/usuario"
	"github.com/aleodoni/voting-go/internal/test/fakes"
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

func (f *fakeTransactor) WithTransaction(ctx context.Context, fn func(context.Context) error) error {
	return fn(ctx)
}

//
// tests
//

func TestEnsureUsuario_UsuarioJaExiste(t *testing.T) {

	usuarioRepo := fakes.NewFakeUsuarioRepository()
	credRepo := fakes.NewFakeCredencialRepository()

	usuarioRepo.Seed(&domainUsuario.Usuario{
		ID:         "user-1",
		KeycloakID: "keycloak-123",
		Username:   "joao",
	})

	uc := usecase.NewEnsureUsuarioUseCase(usuarioRepo, credRepo, &fakeTransactor{})

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
	credRepo := fakes.NewFakeCredencialRepository()

	uc := usecase.NewEnsureUsuarioUseCase(usuarioRepo, credRepo, &fakeTransactor{})

	input := defaultInput()
	input.KeycloakID = "keycloak-456"

	u, err := uc.Execute(context.Background(), input)

	if err != nil {
		t.Fatalf("esperava nil, got %v", err)
	}

	if u.Credencial == nil {
		t.Fatal("esperava credencial preenchida")
	}

	cred, err := credRepo.FindByUsuarioID(context.Background(), u.ID)

	if err != nil {
		t.Fatal("credencial não foi criada")
	}

	if cred.UsuarioID != u.ID {
		t.Error("credencial não pertence ao usuario criado")
	}
}

func TestEnsureUsuario_CredencialComDefaultsCorretos(t *testing.T) {

	usuarioRepo := fakes.NewFakeUsuarioRepository()
	credRepo := fakes.NewFakeCredencialRepository()

	uc := usecase.NewEnsureUsuarioUseCase(usuarioRepo, credRepo, &fakeTransactor{})

	u, err := uc.Execute(context.Background(), defaultInput())

	if err != nil {
		t.Fatalf("esperava nil, got %v", err)
	}

	cred, err := credRepo.FindByUsuarioID(context.Background(), u.ID)

	if err != nil {
		t.Fatal("credencial não criada")
	}

	if !cred.Ativo {
		t.Error("esperava Ativo = true")
	}

	if cred.PodeVotar {
		t.Error("esperava PodeVotar = false")
	}

	if cred.PodeAdministrar {
		t.Error("esperava PodeAdministrar = false")
	}
}

func TestEnsureUsuario_IDsGerados(t *testing.T) {

	usuarioRepo := fakes.NewFakeUsuarioRepository()
	credRepo := fakes.NewFakeCredencialRepository()

	uc := usecase.NewEnsureUsuarioUseCase(usuarioRepo, credRepo, &fakeTransactor{})

	u, err := uc.Execute(context.Background(), defaultInput())

	if err != nil {
		t.Fatalf("esperava nil, got %v", err)
	}

	if u.ID == "" {
		t.Error("esperava ID do usuario preenchido")
	}

	cred, err := credRepo.FindByUsuarioID(context.Background(), u.ID)

	if err != nil {
		t.Fatal("credencial não criada")
	}

	if cred.ID == "" {
		t.Error("esperava ID da credencial preenchido")
	}
}
