package usuario

import (
	"context"

	"github.com/aleodoni/voting-go/internal/application/shared"
	domainUsuario "github.com/aleodoni/voting-go/internal/domain/usuario"
)

type UpdateCredencialUseCase struct {
	usuarioRepo    domainUsuario.UsuarioRepository
	credencialRepo domainUsuario.CredencialRepository
}

func NewUpdateCredencialUseCase(
	usuarioRepo domainUsuario.UsuarioRepository,
	credencialRepo domainUsuario.CredencialRepository,
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

func (uc *UpdateCredencialUseCase) Execute(ctx context.Context, input UpdateCredencialInput) (*domainUsuario.Credencial, error) {
	// Verificar se o usuário logado é admin
	if err := shared.VerificarAdmin(ctx, uc.usuarioRepo, input.AdminKeycloakID); err != nil {
		return nil, err
	}

	// Buscar a credencial do usuário no BD
	cred, err := uc.credencialRepo.FindByUsuarioID(ctx, input.UsuarioID)
	if err != nil {
		return nil, err
	}

	cred.Ativo = input.Ativo
	cred.PodeVotar = input.PodeVotar
	cred.PodeAdministrar = input.PodeAdministrar

	// Atualizar a credencial
	if err := uc.credencialRepo.Update(ctx, cred); err != nil {
		return nil, err
	}

	return cred, nil
}
