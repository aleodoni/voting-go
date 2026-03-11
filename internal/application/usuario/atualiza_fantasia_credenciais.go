package usuario

import (
	"context"

	domainUsuario "github.com/aleodoni/voting-go/internal/domain/usuario"
)

type UpdateDisplayNamePermissionsInput struct {
	LoggedInUserKeycloakID string
	UserID                 string
	DisplayName            *string
	IsActive               bool
	CanAdmin               bool
	CanVote                bool
}

type UpdateDisplayNamePermissionsUseCase struct {
	repo domainUsuario.UsuarioRepository
}

func NewUpdateDisplayNamePermissionsUseCase(
	repo domainUsuario.UsuarioRepository,
) *UpdateDisplayNamePermissionsUseCase {
	return &UpdateDisplayNamePermissionsUseCase{
		repo: repo,
	}
}

func (uc *UpdateDisplayNamePermissionsUseCase) Execute(
	ctx context.Context,
	input UpdateDisplayNamePermissionsInput,
) error {

	// Verificar se o usuário logado é admin
	loggedUser, err := uc.repo.FindByKeycloakID(ctx, input.LoggedInUserKeycloakID)
	if err != nil {
		return err
	}

	if loggedUser.Credencial == nil || !loggedUser.Credencial.IsAdmin() || !loggedUser.Credencial.IsActive() {
		return domainUsuario.ErrNotAdmin
	}

	// Atualizar o display name e as permissões do usuário
	return uc.repo.UpdateDisplayNamePermissions(
		ctx,
		input.UserID,
		input.DisplayName,
		input.IsActive,
		input.CanAdmin,
		input.CanVote,
	)
}
