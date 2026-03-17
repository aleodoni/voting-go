// Package shared contains common utilities and types used across the application.
package shared

import (
	"context"

	domainUsuario "github.com/aleodoni/voting-go/internal/domain/usuario"
)

// VerificarAdmin verifica se o usuário logado é admin e ativo.
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

// VerificarVota verifica se o usuário logado pode votar.
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

func VerificarAtivo(u *domainUsuario.Usuario) error {
	if u.Credencial == nil || !u.Credencial.IsActive() {
		return domainUsuario.ErrUserNotActive
	}
	return nil
}
