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

	erroInativo := VerificarAtivo(ctx, repo, keycloakID)

	if erroInativo != nil {
		return erroInativo
	}

	if loggedUser.Credencial == nil || !loggedUser.Credencial.IsAdmin() {
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

	erroInativo := VerificarAtivo(ctx, repo, keycloakID)

	if erroInativo != nil {
		return erroInativo
	}

	if loggedUser.Credencial == nil || !loggedUser.Credencial.CanVote() {
		return domainUsuario.ErrUserNotVoter
	}

	return nil
}

func VerificarAtivo(ctx context.Context, repo domainUsuario.UsuarioRepository, keycloakID string) error {
	loggedUser, err := repo.FindByKeycloakID(ctx, keycloakID)
	if err != nil {
		return err
	}

	if loggedUser.Credencial == nil || !loggedUser.Credencial.IsActive() {
		return domainUsuario.ErrUserNotActive
	}

	return nil
}
