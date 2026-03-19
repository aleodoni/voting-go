package usuario

import (
	"context"

	"github.com/aleodoni/voting-go/internal/application/shared"
	domainUsuario "github.com/aleodoni/voting-go/internal/domain/usuario"
)

// UpdateDisplayNamePermissionsInput contém os dados necessários para atualizar
// o nome de exibição e as permissões de um usuário.
type UpdateDisplayNamePermissionsInput struct {
	LoggedInUserKeycloakID string
	UserID                 string
	DisplayName            *string
	IsActive               bool
	CanAdmin               bool
	CanVote                bool
}

// UpdateDisplayNamePermissionsUseCase atualiza o nome de exibição e as permissões de um usuário.
//
// Regras de negócio:
//   - o usuário autenticado deve ser administrador ativo
type UpdateDisplayNamePermissionsUseCase struct {
	repo domainUsuario.UsuarioRepository
}

// NewUpdateDisplayNamePermissionsUseCase cria uma nova instância de [UpdateDisplayNamePermissionsUseCase].
func NewUpdateDisplayNamePermissionsUseCase(
	repo domainUsuario.UsuarioRepository,
) *UpdateDisplayNamePermissionsUseCase {
	return &UpdateDisplayNamePermissionsUseCase{
		repo: repo,
	}
}

// Execute atualiza o nome de exibição e as permissões do usuário informado em
// [UpdateDisplayNamePermissionsInput.UserID].
func (uc *UpdateDisplayNamePermissionsUseCase) Execute(
	ctx context.Context,
	input UpdateDisplayNamePermissionsInput,
) error {
	if err := shared.VerificarAdmin(ctx, uc.repo, input.LoggedInUserKeycloakID); err != nil {
		return err
	}

	return uc.repo.UpdateDisplayNamePermissions(
		ctx,
		input.UserID,
		input.DisplayName,
		input.IsActive,
		input.CanAdmin,
		input.CanVote,
	)
}
