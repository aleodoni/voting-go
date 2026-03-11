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

	println("----------------1")
	loggedUser, err := uc.repo.FindByKeycloakID(ctx, input.LoggedInUserKeycloakID)
	println(loggedUser.Nome)
	if err != nil {
		println("----------------2")
		return err
	}

	println("----------------3")
	if loggedUser.Credencial == nil || !loggedUser.Credencial.IsAdmin() || !loggedUser.Credencial.IsActive() {
		println("----------------4")
		return domainUsuario.ErrNotAdmin
	}

	println("----------------5")
	return uc.repo.UpdateDisplayNamePermissions(
		ctx,
		input.UserID,
		input.DisplayName,
		input.IsActive,
		input.CanAdmin,
		input.CanVote,
	)
}
