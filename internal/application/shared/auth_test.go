package shared_test

import (
	"context"
	"errors"
	"testing"

	"github.com/aleodoni/voting-go/internal/application/shared"
	domainUsuario "github.com/aleodoni/voting-go/internal/domain/usuario"
	"github.com/aleodoni/voting-go/internal/test/fakes"
)

// ─────────────────────────────────────────────
// VerificarAdmin
// ─────────────────────────────────────────────

func TestVerificarAdmin_AdminAtivoRetornaNil(t *testing.T) {
	repo := fakes.NewFakeUsuarioRepository()
	repo.Seed(adminAtivo("keycloak-admin"))

	err := shared.VerificarAdmin(context.Background(), repo, "keycloak-admin")
	if err != nil {
		t.Fatalf("esperava nil, obteve: %v", err)
	}
}

func TestVerificarAdmin_UsuarioNaoEncontrado(t *testing.T) {
	repo := fakes.NewFakeUsuarioRepository()

	err := shared.VerificarAdmin(context.Background(), repo, "keycloak-inexistente")
	if !errors.Is(err, domainUsuario.ErrUserNotFound) {
		t.Fatalf("esperava ErrUserNotFound, obteve: %v", err)
	}
}

func TestVerificarAdmin_UsuarioInativo(t *testing.T) {
	repo := fakes.NewFakeUsuarioRepository()
	repo.Seed(usuarioInativo("keycloak-inativo"))

	err := shared.VerificarAdmin(context.Background(), repo, "keycloak-inativo")
	if !errors.Is(err, domainUsuario.ErrUserNotActive) {
		t.Fatalf("esperava ErrUserNotActive, obteve: %v", err)
	}
}

func TestVerificarAdmin_UsuarioSemCredencial(t *testing.T) {
	repo := fakes.NewFakeUsuarioRepository()
	repo.Seed(usuarioSemCredencial("keycloak-sem-cred"))

	err := shared.VerificarAdmin(context.Background(), repo, "keycloak-sem-cred")
	if !errors.Is(err, domainUsuario.ErrUserNotActive) {
		t.Fatalf("esperava ErrUserNotActive, obteve: %v", err)
	}
}

func TestVerificarAdmin_UsuarioSemPermissaoAdmin(t *testing.T) {
	repo := fakes.NewFakeUsuarioRepository()
	repo.Seed(usuarioSemAdmin("keycloak-comum"))

	err := shared.VerificarAdmin(context.Background(), repo, "keycloak-comum")
	if !errors.Is(err, domainUsuario.ErrUserNotAdmin) {
		t.Fatalf("esperava ErrUserNotAdmin, obteve: %v", err)
	}
}

// ─────────────────────────────────────────────
// VerificarVota
// ─────────────────────────────────────────────

func TestVerificarVota_VotanteAtivoRetornaUsuario(t *testing.T) {
	repo := fakes.NewFakeUsuarioRepository()
	repo.Seed(adminAtivo("keycloak-votante"))

	u, err := shared.VerificarVota(context.Background(), repo, "keycloak-votante")
	if err != nil {
		t.Fatalf("esperava nil, obteve: %v", err)
	}
	if u == nil {
		t.Fatal("esperava usuário retornado, obteve nil")
	}
}

func TestVerificarVota_UsuarioNaoEncontrado(t *testing.T) {
	repo := fakes.NewFakeUsuarioRepository()

	_, err := shared.VerificarVota(context.Background(), repo, "keycloak-inexistente")
	if !errors.Is(err, domainUsuario.ErrUserNotFound) {
		t.Fatalf("esperava ErrUserNotFound, obteve: %v", err)
	}
}

func TestVerificarVota_UsuarioInativo(t *testing.T) {
	repo := fakes.NewFakeUsuarioRepository()
	repo.Seed(usuarioInativo("keycloak-inativo"))

	_, err := shared.VerificarVota(context.Background(), repo, "keycloak-inativo")
	if !errors.Is(err, domainUsuario.ErrUserNotActive) {
		t.Fatalf("esperava ErrUserNotActive, obteve: %v", err)
	}
}

func TestVerificarVota_UsuarioSemCredencial(t *testing.T) {
	repo := fakes.NewFakeUsuarioRepository()
	repo.Seed(usuarioSemCredencial("keycloak-sem-cred"))

	_, err := shared.VerificarVota(context.Background(), repo, "keycloak-sem-cred")
	if !errors.Is(err, domainUsuario.ErrUserNotActive) {
		t.Fatalf("esperava ErrUserNotActive, obteve: %v", err)
	}
}

func TestVerificarVota_UsuarioSemPermissaoVoto(t *testing.T) {
	repo := fakes.NewFakeUsuarioRepository()
	repo.Seed(usuarioSemVoto("keycloak-sem-voto"))

	_, err := shared.VerificarVota(context.Background(), repo, "keycloak-sem-voto")
	if !errors.Is(err, domainUsuario.ErrUserNotVoter) {
		t.Fatalf("esperava ErrUserNotVoter, obteve: %v", err)
	}
}

// ─────────────────────────────────────────────
// VerificarAtivo
// ─────────────────────────────────────────────

func TestVerificarAtivo_UsuarioAtivoRetornaNil(t *testing.T) {
	u := adminAtivo("qualquer")
	if err := shared.VerificarAtivo(u); err != nil {
		t.Fatalf("esperava nil, obteve: %v", err)
	}
}

func TestVerificarAtivo_UsuarioInativoRetornaErro(t *testing.T) {
	u := usuarioInativo("qualquer")
	if !errors.Is(shared.VerificarAtivo(u), domainUsuario.ErrUserNotActive) {
		t.Fatal("esperava ErrUserNotActive para usuário inativo")
	}
}

func TestVerificarAtivo_SemCredencialRetornaErro(t *testing.T) {
	u := usuarioSemCredencial("qualquer")
	if !errors.Is(shared.VerificarAtivo(u), domainUsuario.ErrUserNotActive) {
		t.Fatal("esperava ErrUserNotActive para usuário sem credencial")
	}
}
