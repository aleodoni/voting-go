// Package credencial contains the use cases related to credencial management.
package credencial

import (
	"context"

	domain "github.com/aleodoni/voting-go/internal/domain"
	domainCredencial "github.com/aleodoni/voting-go/internal/domain/credencial"
	domainUsuario "github.com/aleodoni/voting-go/internal/domain/usuario"
)

type UpdateCredencialUseCase struct {
	usuarioRepo    domainUsuario.UsuarioRepository
	credencialRepo domainCredencial.CredencialRepository
}

func NewUpdateCredencialUseCase(
	usuarioRepo domainUsuario.UsuarioRepository,
	credencialRepo domainCredencial.CredencialRepository,
) *UpdateCredencialUseCase {
	return &UpdateCredencialUseCase{
		usuarioRepo:    usuarioRepo,
		credencialRepo: credencialRepo,
	}
}

type UpdateCredencialInput struct {
	AdminKeycloakID string
	UsuarioID       string
	Ativo           bool
	PodeVotar       bool
	PodeAdministrar bool
}

func (uc *UpdateCredencialUseCase) Execute(ctx context.Context, input UpdateCredencialInput) (*domainCredencial.Credencial, error) {
	admin, err := uc.usuarioRepo.FindByKeycloakID(ctx, input.AdminKeycloakID)
	if err != nil {
		return nil, err
	}
	if !admin.Credencial.IsActive() || !admin.Credencial.IsAdmin() {
		return nil, domain.ErrForbidden
	}

	cred, err := uc.credencialRepo.FindByUsuarioID(ctx, input.UsuarioID)
	if err != nil {
		return nil, err
	}

	cred.Ativo = input.Ativo
	cred.PodeVotar = input.PodeVotar
	cred.PodeAdministrar = input.PodeAdministrar

	if err := uc.credencialRepo.Update(ctx, cred); err != nil {
		return nil, err
	}

	return cred, nil
}
