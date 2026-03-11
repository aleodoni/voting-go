package usuario_test

import (
	"context"
	"testing"

	usecase "github.com/aleodoni/voting-go/internal/application/usuario"
	domainCredencial "github.com/aleodoni/voting-go/internal/domain/credencial"
	domainUsuario "github.com/aleodoni/voting-go/internal/domain/usuario"
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

//
// fakes
//

type fakeUsuarioRepo struct {
	usuarios map[string]*domainUsuario.Usuario
}

func newFakeUsuarioRepo() *fakeUsuarioRepo {
	return &fakeUsuarioRepo{
		usuarios: make(map[string]*domainUsuario.Usuario),
	}
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

func (f *fakeUsuarioRepo) UpdateDisplayNamePermissions(
	ctx context.Context,
	userID string,
	displayName *string,
	isActive bool,
	canAdmin bool,
	canVote bool,
) error {

	for _, u := range f.usuarios {

		if u.ID == userID {

			if displayName != nil {
				u.NomeFantasia = displayName
			}

			if u.Credencial != nil {
				u.Credencial.Ativo = isActive
				u.Credencial.PodeAdministrar = canAdmin
				u.Credencial.PodeVotar = canVote
			}
		}
	}

	return nil
}

type fakeCredencialRepo struct {
	credenciais map[string]*domainCredencial.Credencial
}

func newFakeCredencialRepo() *fakeCredencialRepo {
	return &fakeCredencialRepo{
		credenciais: make(map[string]*domainCredencial.Credencial),
	}
}

func (f *fakeCredencialRepo) FindByUsuarioID(ctx context.Context, usuarioID string) (*domainCredencial.Credencial, error) {

	cred, ok := f.credenciais[usuarioID]
	if !ok {
		return nil, domainCredencial.ErrNotFound
	}

	return cred, nil
}

func (f *fakeCredencialRepo) Create(ctx context.Context, cred *domainCredencial.Credencial) error {
	f.credenciais[cred.UsuarioID] = cred
	return nil
}

func (f *fakeCredencialRepo) Update(ctx context.Context, cred *domainCredencial.Credencial) error {
	f.credenciais[cred.UsuarioID] = cred
	return nil
}

type fakeTransactor struct{}

func (f *fakeTransactor) WithTransaction(ctx context.Context, fn func(context.Context) error) error {
	return fn(ctx)
}

//
// tests
//

func TestEnsureUsuario_UsuarioJaExiste(t *testing.T) {

	usuarioRepo := newFakeUsuarioRepo()
	credRepo := newFakeCredencialRepo()

	usuarioRepo.usuarios["keycloak-123"] = &domainUsuario.Usuario{
		ID:         "user-1",
		KeycloakID: "keycloak-123",
		Username:   "joao",
	}

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

	usuarioRepo := newFakeUsuarioRepo()
	credRepo := newFakeCredencialRepo()

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

	cred, ok := credRepo.credenciais[u.ID]

	if !ok {
		t.Fatal("credencial não foi criada")
	}

	if cred.UsuarioID != u.ID {
		t.Error("credencial não pertence ao usuario criado")
	}
}

func TestEnsureUsuario_CredencialComDefaultsCorretos(t *testing.T) {

	usuarioRepo := newFakeUsuarioRepo()
	credRepo := newFakeCredencialRepo()

	uc := usecase.NewEnsureUsuarioUseCase(usuarioRepo, credRepo, &fakeTransactor{})

	u, err := uc.Execute(context.Background(), defaultInput())

	if err != nil {
		t.Fatalf("esperava nil, got %v", err)
	}

	cred, ok := credRepo.credenciais[u.ID]

	if !ok {
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

	usuarioRepo := newFakeUsuarioRepo()
	credRepo := newFakeCredencialRepo()

	uc := usecase.NewEnsureUsuarioUseCase(usuarioRepo, credRepo, &fakeTransactor{})

	u, err := uc.Execute(context.Background(), defaultInput())

	if err != nil {
		t.Fatalf("esperava nil, got %v", err)
	}

	if u.ID == "" {
		t.Error("esperava ID do usuario preenchido")
	}

	cred, ok := credRepo.credenciais[u.ID]

	if !ok {
		t.Fatal("credencial não criada")
	}

	if cred.ID == "" {
		t.Error("esperava ID da credencial preenchido")
	}
}
