package usuario

import (
	"context"

	domainUsuario "github.com/aleodoni/voting-go/internal/domain/usuario"
)

// UpdateDisplayNameInput contém os dados necessários para atualizar
// o nome de exibição de um usuário.
type UpdateDisplayNameInput struct {
	LoggedInUserKeycloakID string
	UserID                 string
	DisplayName            *string
}

// UpdateDisplayNameUseCase atualiza o nome de exibição e as permissões de um usuário.
//
// Regras de negócio:
//   - o usuário autenticado só pode alterar seu próprio nome de exibição
type UpdateDisplayNameUseCase struct {
	repo domainUsuario.UsuarioRepository
}

// NewUpdateDisplayNameUseCase cria uma nova instância de [UpdateDisplayNameUseCase].
func NewUpdateDisplayNameUseCase(
	repo domainUsuario.UsuarioRepository,
) *UpdateDisplayNameUseCase {
	return &UpdateDisplayNameUseCase{
		repo: repo,
	}
}

// Execute atualiza o nome de exibição e as permissões do usuário informado em
// [UpdateDisplayNameInput.UserID].
func (uc *UpdateDisplayNameUseCase) Execute(
	ctx context.Context,
	input UpdateDisplayNameInput,
) error {
	user, err := uc.repo.FindByID(ctx, input.UserID)
	if err != nil {
		return domainUsuario.ErrUserNotFound
	}

	if user.KeycloakID != input.LoggedInUserKeycloakID {
		return domainUsuario.ErrUserNotFound
	}

	return uc.repo.UpdateDisplayName(
		ctx,
		input.UserID,
		input.DisplayName,
	)
}
