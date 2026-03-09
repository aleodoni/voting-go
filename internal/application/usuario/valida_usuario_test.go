package usuario_test

import (
	"context"
	"testing"

	usecase "github.com/aleodoni/voting-go/internal/application/usuario"
	domainCredencial "github.com/aleodoni/voting-go/internal/domain/credencial"
	domainUsuario "github.com/aleodoni/voting-go/internal/domain/usuario"
)

// --- fakes ---

type fakeUsuarioRepo struct {
	usuarios map[string]*domainUsuario.Usuario
}

func newFakeUsuarioRepo() *fakeUsuarioRepo {
	return &fakeUsuarioRepo{usuarios: make(map[string]*domainUsuario.Usuario)}
}

func (f *fakeUsuarioRepo) FindByKeycloakID(ctx context.Context, keycloakID string) (*domainUsuario.Usuario, error) {
	u, ok := f.usuarios[keycloakID]
	if !ok {
		return nil, domainUsuario.ErrNotFound
	}
	return u, nil
}

func (f *fakeUsuarioRepo) FindByUsername(ctx context.Context, username string) (*domainUsuario.Usuario, error) {
	return nil, domainUsuario.ErrNotFound
}

func (f *fakeUsuarioRepo) Create(ctx context.Context, u *domainUsuario.Usuario) error {
	f.usuarios[u.KeycloakID] = u
	return nil
}

type fakeCredencialRepo struct {
	credenciais []*domainCredencial.Credencial
}

func (f *fakeCredencialRepo) FindByUsuarioID(ctx context.Context, usuarioID string) (*domainCredencial.Credencial, error) {
	return nil, nil
}

func (f *fakeCredencialRepo) Create(ctx context.Context, cred *domainCredencial.Credencial) error {
	f.credenciais = append(f.credenciais, cred)
	return nil
}

type fakeTransactor struct{}

func (f *fakeTransactor) WithTransaction(ctx context.Context, fn func(context.Context) error) error {
	return fn(ctx)
}

// --- testes ---

func TestEnsureUsuario_UsuarioJaExiste(t *testing.T) {
	usuarioRepo := newFakeUsuarioRepo()
	usuarioRepo.usuarios["keycloak-123"] = &domainUsuario.Usuario{
		ID:         "user-1",
		KeycloakID: "keycloak-123",
		Username:   "joao",
	}

	uc := usecase.NewEnsureUsuarioUseCase(usuarioRepo, &fakeCredencialRepo{}, &fakeTransactor{})

	u, err := uc.Execute(context.Background(), usecase.EnsureUsuarioInput{
		KeycloakID: "keycloak-123",
		Username:   "joao",
		Email:      "joao@email.com",
		Nome:       "João",
	})

	if err != nil {
		t.Fatalf("esperava nil, got %v", err)
	}
	if u.KeycloakID != "keycloak-123" {
		t.Errorf("esperava keycloak-123, got %s", u.KeycloakID)
	}
}

func TestEnsureUsuario_UsuarioNaoExiste_CriaUsuarioECredencial(t *testing.T) {
	usuarioRepo := newFakeUsuarioRepo()
	credencialRepo := &fakeCredencialRepo{}

	uc := usecase.NewEnsureUsuarioUseCase(usuarioRepo, credencialRepo, &fakeTransactor{})

	u, err := uc.Execute(context.Background(), usecase.EnsureUsuarioInput{
		KeycloakID: "keycloak-456",
		Username:   "maria",
		Email:      "maria@email.com",
		Nome:       "Maria",
	})

	if err != nil {
		t.Fatalf("esperava nil, got %v", err)
	}
	if u.KeycloakID != "keycloak-456" {
		t.Errorf("esperava keycloak-456, got %s", u.KeycloakID)
	}
	if u.Credencial == nil {
		t.Error("esperava credencial preenchida")
	}
	if len(credencialRepo.credenciais) != 1 {
		t.Errorf("esperava 1 credencial criada, got %d", len(credencialRepo.credenciais))
	}
	if credencialRepo.credenciais[0].UsuarioID != u.ID {
		t.Error("credencial não pertence ao usuario criado")
	}
}
