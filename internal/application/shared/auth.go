// Package shared contains common utilities and types used across the application.
package shared

import (
	"context"

	domainUsuario "github.com/aleodoni/voting-go/internal/domain/usuario"
)

// VerificarAdmin verifica se o usuário autenticado é administrador ativo.
//
// Regras de negócio:
//   - o usuário deve existir no sistema
//   - o usuário deve estar ativo (ver [VerificarAtivo])
//   - o usuário deve possuir credencial com perfil de administrador
//
// Retorna [domainUsuario.ErrUserNotAdmin] se o usuário não for administrador.
func VerificarAdmin(ctx context.Context, repo domainUsuario.UsuarioRepository, keycloakID string) error {
	loggedUser, err := repo.FindByKeycloakID(ctx, keycloakID)
	if err != nil {
		return err
	}

	erroInativo := VerificarAtivo(loggedUser)

	if erroInativo != nil {
		return erroInativo
	}

	if loggedUser.Credencial == nil || !loggedUser.Credencial.IsAdmin() {
		return domainUsuario.ErrUserNotAdmin
	}

	return nil
}

// VerificarVota verifica se o usuário autenticado possui permissão para votar.
//
// Regras de negócio:
//   - o usuário deve existir no sistema
//   - o usuário deve estar ativo (ver [VerificarAtivo])
//   - o usuário deve possuir credencial com permissão de voto
//
// Retorna o [domainUsuario.Usuario] autenticado em caso de sucesso.
// Retorna [domainUsuario.ErrUserNotVoter] se o usuário não tiver permissão para votar.
func VerificarVota(ctx context.Context, repo domainUsuario.UsuarioRepository, keycloakID string) (*domainUsuario.Usuario, error) {
	loggedUser, err := repo.FindByKeycloakID(ctx, keycloakID)
	if err != nil {
		return nil, err
	}

	erroInativo := VerificarAtivo(loggedUser)

	if erroInativo != nil {
		return nil, erroInativo
	}

	if loggedUser.Credencial == nil || !loggedUser.Credencial.CanVote() {
		return nil, domainUsuario.ErrUserNotVoter
	}

	return loggedUser, nil
}

// VerificarAtivo verifica se o usuário autenticado está ativo no sistema.
//
// Retorna [domainUsuario.ErrUserNotActive] se o usuário não possuir credencial
// ou se a credencial estiver inativa.
func VerificarAtivo(u *domainUsuario.Usuario) error {
	if u.Credencial == nil || !u.Credencial.IsActive() {
		return domainUsuario.ErrUserNotActive
	}
	return nil
}
