package usuario_test

import (
	"context"
	"testing"

	usecase "github.com/aleodoni/voting-go/internal/application/usuario"
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
		return nil, domainUsuario.ErrUserNotFound
	}
	return u, nil
}

func (f *fakeUsuarioRepo) FindByUsername(ctx context.Context, username string) (*domainUsuario.Usuario, error) {
	return nil, domainUsuario.ErrUserNotFound
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
	credenciais map[string]*domainUsuario.Credencial
}

func newFakeCredencialRepo() *fakeCredencialRepo {
	return &fakeCredencialRepo{credenciais: make(map[string]*domainUsuario.Credencial)}
}

func (f *fakeCredencialRepo) FindByUsuarioID(ctx context.Context, usuarioID string) (*domainUsuario.Credencial, error) {
	cred, ok := f.credenciais[usuarioID]
	if !ok {
		return nil, domainUsuario.ErrUserNotFound
	}
	return cred, nil
}

func (f *fakeCredencialRepo) Create(ctx context.Context, cred *domainUsuario.Credencial) error {
	f.credenciais[cred.UsuarioID] = cred
	return nil
}

func (f *fakeCredencialRepo) Update(ctx context.Context, cred *domainUsuario.Credencial) error {
	f.credenciais[cred.UsuarioID] = cred
	return nil
}

// --- helpers ---

func adminAtivo() *domainUsuario.Usuario {
	return &domainUsuario.Usuario{
		ID:         "admin-1",
		KeycloakID: "keycloak-admin",
		Credencial: &domainUsuario.Credencial{
			Ativo:           true,
			PodeAdministrar: true,
		},
	}
}

func usuarioAlvo(credencial *domainUsuario.Credencial) *domainUsuario.Usuario {
	return &domainUsuario.Usuario{
		ID:         "user-1",
		KeycloakID: "keycloak-user",
		Credencial: credencial,
	}
}

// --- testes ---

func TestUpdateCredencial_Sucesso(t *testing.T) {
	usuarioRepo := newFakeUsuarioRepo()
	usuarioRepo.usuarios["keycloak-admin"] = adminAtivo()

	credencialRepo := newFakeCredencialRepo()
	credencialRepo.credenciais["user-1"] = &domainUsuario.Credencial{
		ID:        "cred-1",
		UsuarioID: "user-1",
		Ativo:     false,
		PodeVotar: false,
	}

	uc := usecase.NewUpdateCredencialUseCase(usuarioRepo, credencialRepo)

	cred, err := uc.Execute(context.Background(), usecase.UpdateCredencialInput{
		AdminKeycloakID: "keycloak-admin",
		UsuarioID:       "user-1",
		Ativo:           true,
		PodeVotar:       true,
		PodeAdministrar: false,
	})

	if err != nil {
		t.Fatalf("esperava nil, got %v", err)
	}
	if !cred.Ativo {
		t.Error("esperava Ativo = true")
	}
	if !cred.PodeVotar {
		t.Error("esperava PodeVotar = true")
	}
	if cred.PodeAdministrar {
		t.Error("esperava PodeAdministrar = false")
	}
}

func TestUpdateCredencial_AdminInativo_Inativo(t *testing.T) {
	usuarioRepo := newFakeUsuarioRepo()
	usuarioRepo.usuarios["keycloak-admin"] = &domainUsuario.Usuario{
		ID:         "admin-1",
		KeycloakID: "keycloak-admin",
		Credencial: &domainUsuario.Credencial{
			Ativo:           false,
			PodeAdministrar: true,
		},
	}

	uc := usecase.NewUpdateCredencialUseCase(usuarioRepo, newFakeCredencialRepo())

	_, err := uc.Execute(context.Background(), usecase.UpdateCredencialInput{
		AdminKeycloakID: "keycloak-admin",
		UsuarioID:       "user-1",
	})

	if err != domainUsuario.ErrUserNotActive {
		t.Errorf("esperava ErrUserNotActive, got %v", err)
	}
}

func TestUpdateCredencial_AdminSemPermissao_Forbidden(t *testing.T) {
	usuarioRepo := newFakeUsuarioRepo()
	usuarioRepo.usuarios["keycloak-admin"] = &domainUsuario.Usuario{
		ID:         "admin-1",
		KeycloakID: "keycloak-admin",
		Credencial: &domainUsuario.Credencial{
			Ativo:           true,
			PodeAdministrar: false,
		},
	}

	uc := usecase.NewUpdateCredencialUseCase(usuarioRepo, newFakeCredencialRepo())

	_, err := uc.Execute(context.Background(), usecase.UpdateCredencialInput{
		AdminKeycloakID: "keycloak-admin",
		UsuarioID:       "user-1",
	})

	if err != domainUsuario.ErrUserNotAdmin {
		t.Errorf("esperava ErrUserNotAdmin, got %v", err)
	}
}

func TestUpdateCredencial_AdminNaoEncontrado(t *testing.T) {
	uc := usecase.NewUpdateCredencialUseCase(newFakeUsuarioRepo(), newFakeCredencialRepo())

	_, err := uc.Execute(context.Background(), usecase.UpdateCredencialInput{
		AdminKeycloakID: "keycloak-inexistente",
		UsuarioID:       "user-1",
	})

	if err != domainUsuario.ErrUserNotFound {
		t.Errorf("esperava ErrUserNotFound, got %v", err)
	}
}

func TestUpdateCredencial_UsuarioAlvoNaoEncontrado(t *testing.T) {
	usuarioRepo := newFakeUsuarioRepo()
	usuarioRepo.usuarios["keycloak-admin"] = adminAtivo()

	uc := usecase.NewUpdateCredencialUseCase(usuarioRepo, newFakeCredencialRepo())

	_, err := uc.Execute(context.Background(), usecase.UpdateCredencialInput{
		AdminKeycloakID: "keycloak-admin",
		UsuarioID:       "user-inexistente",
	})

	if err != domainUsuario.ErrUserNotFound {
		t.Errorf("esperava ErrNotFound, got %v", err)
	}
}
