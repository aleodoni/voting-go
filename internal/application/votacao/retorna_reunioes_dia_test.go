package votacao_test

import (
	"context"
	"testing"

	usecase "github.com/aleodoni/voting-go/internal/application/votacao"
	domainUsuario "github.com/aleodoni/voting-go/internal/domain/usuario"
	domainVotacao "github.com/aleodoni/voting-go/internal/domain/votacao"
	"github.com/aleodoni/voting-go/internal/test/fakes"
)

type fakeReuniaoRepository struct {
	reunioes []*domainVotacao.Reuniao

	GetReunioesDiaErr  error
	FindReuniaoByIDErr error
}

var _ domainVotacao.ReuniaoRepository = (*fakeReuniaoRepository)(nil)

func newFakeReuniaoRepository() *fakeReuniaoRepository {
	return &fakeReuniaoRepository{}
}

func (f *fakeReuniaoRepository) FindReuniaoByID(ctx context.Context, reuniaoID string) (*domainVotacao.Reuniao, error) {
	if f.FindReuniaoByIDErr != nil {
		return nil, f.FindReuniaoByIDErr
	}

	for _, r := range f.reunioes {
		if r.ID == reuniaoID {
			return r, nil
		}
	}

	return nil, domainVotacao.ErrReuniaoNotFound
}

func (f *fakeReuniaoRepository) GetReunioesDia(ctx context.Context) ([]*domainVotacao.Reuniao, error) {
	if f.GetReunioesDiaErr != nil {
		return nil, f.GetReunioesDiaErr
	}

	return f.reunioes, nil
}

//
// helpers
//

func adminUsuario(keycloakID, userID string) *domainUsuario.Usuario {
	return &domainUsuario.Usuario{
		ID:         userID,
		KeycloakID: keycloakID,
		Username:   "admin",
		Credencial: &domainUsuario.Credencial{
			ID:              "cred-admin",
			UsuarioID:       userID,
			Ativo:           true,
			PodeAdministrar: true,
		},
	}
}

//
// tests
//

func TestRetornaReunioesDia_AdminRetornaReunioesDodia(t *testing.T) {

	usuarioRepo := fakes.NewFakeUsuarioRepository()
	reuniaoRepo := newFakeReuniaoRepository()

	usuarioRepo.Seed(adminUsuario("keycloak-admin", "user-admin"))

	reuniaoRepo.reunioes = []*domainVotacao.Reuniao{
		{ID: "reuniao-1", RecNumero: "001"},
		{ID: "reuniao-2", RecNumero: "002"},
	}

	uc := usecase.NewRetornaReunioesDiaUseCase(usuarioRepo, reuniaoRepo)

	result, err := uc.Execute(context.Background(), usecase.RetornaReunioesDiaInput{
		LoggedInUserKeycloakID: "keycloak-admin",
	})

	if err != nil {
		t.Fatalf("esperava nil, got %v", err)
	}

	if len(result) != 2 {
		t.Errorf("esperava 2 reunioes, got %d", len(result))
	}
}

func TestRetornaReunioesDia_AdminSemReunioesNoDia(t *testing.T) {

	usuarioRepo := fakes.NewFakeUsuarioRepository()
	reuniaoRepo := newFakeReuniaoRepository()

	usuarioRepo.Seed(adminUsuario("keycloak-admin", "user-admin"))

	uc := usecase.NewRetornaReunioesDiaUseCase(usuarioRepo, reuniaoRepo)

	result, err := uc.Execute(context.Background(), usecase.RetornaReunioesDiaInput{
		LoggedInUserKeycloakID: "keycloak-admin",
	})

	if err != nil {
		t.Fatalf("esperava nil, got %v", err)
	}

	if len(result) != 0 {
		t.Errorf("esperava 0 reunioes, got %d", len(result))
	}
}

func TestRetornaReunioesDia_UsuarioNaoEncontrado(t *testing.T) {

	usuarioRepo := fakes.NewFakeUsuarioRepository()
	reuniaoRepo := newFakeReuniaoRepository()

	uc := usecase.NewRetornaReunioesDiaUseCase(usuarioRepo, reuniaoRepo)

	_, err := uc.Execute(context.Background(), usecase.RetornaReunioesDiaInput{
		LoggedInUserKeycloakID: "keycloak-inexistente",
	})

	if err == nil {
		t.Fatal("esperava erro, got nil")
	}
}

func TestRetornaReunioesDia_UsuarioSemCredencial(t *testing.T) {

	usuarioRepo := fakes.NewFakeUsuarioRepository()
	reuniaoRepo := newFakeReuniaoRepository()

	usuarioRepo.Seed(&domainUsuario.Usuario{
		ID:         "user-sem-cred",
		KeycloakID: "keycloak-sem-cred",
		Credencial: nil,
	})

	uc := usecase.NewRetornaReunioesDiaUseCase(usuarioRepo, reuniaoRepo)

	_, err := uc.Execute(context.Background(), usecase.RetornaReunioesDiaInput{
		LoggedInUserKeycloakID: "keycloak-sem-cred",
	})

	if err != domainUsuario.ErrUserNotActive {
		t.Fatalf("esperava ErrUserNotActive, got %v", err)
	}
}

func TestRetornaReunioesDia_UsuarioNaoEAdmin(t *testing.T) {

	usuarioRepo := fakes.NewFakeUsuarioRepository()
	reuniaoRepo := newFakeReuniaoRepository()

	usuarioRepo.Seed(&domainUsuario.Usuario{
		ID:         "user-comum",
		KeycloakID: "keycloak-comum",
		Credencial: &domainUsuario.Credencial{
			Ativo:           true,
			PodeAdministrar: false,
		},
	})

	uc := usecase.NewRetornaReunioesDiaUseCase(usuarioRepo, reuniaoRepo)

	_, err := uc.Execute(context.Background(), usecase.RetornaReunioesDiaInput{
		LoggedInUserKeycloakID: "keycloak-comum",
	})

	if err != domainUsuario.ErrUserNotAdmin {
		t.Fatalf("esperava ErrNotAdmin, got %v", err)
	}
}

func TestRetornaReunioesDia_UsuarioInativo(t *testing.T) {

	usuarioRepo := fakes.NewFakeUsuarioRepository()
	reuniaoRepo := newFakeReuniaoRepository()

	usuarioRepo.Seed(&domainUsuario.Usuario{
		ID:         "user-inativo",
		KeycloakID: "keycloak-inativo",
		Credencial: &domainUsuario.Credencial{
			Ativo:           false,
			PodeAdministrar: true,
		},
	})

	uc := usecase.NewRetornaReunioesDiaUseCase(usuarioRepo, reuniaoRepo)

	_, err := uc.Execute(context.Background(), usecase.RetornaReunioesDiaInput{
		LoggedInUserKeycloakID: "keycloak-inativo",
	})

	if err != domainUsuario.ErrUserNotActive {
		t.Fatalf("esperava ErrUserNotActive, got %v", err)
	}
}

func TestRetornaReunioesDia_ErroNoRepositorio(t *testing.T) {

	usuarioRepo := fakes.NewFakeUsuarioRepository()
	reuniaoRepo := newFakeReuniaoRepository()

	usuarioRepo.Seed(adminUsuario("keycloak-admin", "user-admin"))
	reuniaoRepo.GetReunioesDiaErr = domainVotacao.ErrReuniaoNotFound

	uc := usecase.NewRetornaReunioesDiaUseCase(usuarioRepo, reuniaoRepo)

	_, err := uc.Execute(context.Background(), usecase.RetornaReunioesDiaInput{
		LoggedInUserKeycloakID: "keycloak-admin",
	})

	if err == nil {
		t.Fatal("esperava erro, got nil")
	}
}
