package usuario

import (
	"context"

	"github.com/aleodoni/voting-go/internal/application/shared"
	domainUsuario "github.com/aleodoni/voting-go/internal/domain/usuario"
)

// UpdateCredencialInput contém os dados necessários para atualizar a credencial de um usuário.
type UpdateCredencialInput struct {
	AdminKeycloakID string
	UsuarioID       string
	Ativo           bool
	PodeVotar       bool
	PodeAdministrar bool
}

// UpdateCredencialUseCase atualiza as permissões da credencial de um usuário.
//
// Regras de negócio:
//   - o usuário autenticado deve ser administrador ativo
//   - a credencial do usuário alvo deve existir no sistema
type UpdateCredencialUseCase struct {
	usuarioRepo domainUsuario.UsuarioRepository
}

// NewUpdateCredencialUseCase cria uma nova instância de [UpdateCredencialUseCase].
func NewUpdateCredencialUseCase(
	usuarioRepo domainUsuario.UsuarioRepository,
) *UpdateCredencialUseCase {
	return &UpdateCredencialUseCase{
		usuarioRepo: usuarioRepo,
	}
}

// Execute atualiza a credencial do usuário informado em [UpdateCredencialInput.UsuarioID].
//
// Retorna a [domainUsuario.Credencial] atualizada em caso de sucesso.
func (uc *UpdateCredencialUseCase) Execute(ctx context.Context, input UpdateCredencialInput) (*domainUsuario.Credencial, error) {
	if err := shared.VerificarAdmin(ctx, uc.usuarioRepo, input.AdminKeycloakID); err != nil {
		return nil, err
	}

	err := uc.usuarioRepo.UpdateDisplayNamePermissions(
		ctx,
		input.UsuarioID,
		nil,
		input.Ativo,
		input.PodeAdministrar,
		input.PodeVotar,
	)
	if err != nil {
		return nil, err
	}

	u, err := uc.usuarioRepo.FindByID(ctx, input.UsuarioID)
	if err != nil {
		return nil, err
	}

	return u.Credencial, nil
}
