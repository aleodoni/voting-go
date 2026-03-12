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

	if loggedUser.Credencial == nil || !loggedUser.Credencial.IsAdmin() || !loggedUser.Credencial.IsActive() {
		return domainUsuario.ErrUserNotAdmin
	}

	return nil
}

// VerificarVota verifica se o usuário logado pode votar.
func VerificarVota(ctx context.Context, repo domainUsuario.UsuarioRepository, keycloakID string) error {
	loggedUser, err := repo.FindByKeycloakID(ctx, keycloakID)
	if err != nil {
		return err
	}

	if loggedUser.Credencial == nil || !loggedUser.Credencial.CanVote() || !loggedUser.Credencial.IsActive() {
		return domainUsuario.ErrUserNotVoter
	}

	return nil
}
