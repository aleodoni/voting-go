package usuario

import (
	"context"

	"github.com/aleodoni/voting-go/internal/application/shared"
	"github.com/aleodoni/voting-go/internal/domain/usuario"
)

// RetornaUsuarioInput contém os dados necessários para buscar usuário.
type RetornaUsuarioInput struct {
	LoggedInUserKeycloakID string
	UsuarioID              string
}

// RetornaUsuarioUseCase retorna os dados do usuário
//
// Regras de negócio:
//   - o usuário autenticado deve ser administrador ativo
//   - [RetornaUsuarioInput.UsuarioID] deve ser um id válido
type RetornaUsuarioUseCase struct {
	repoUsuario usuario.UsuarioRepository
}

// NewRetornaUsuarioUseCase cria uma nova instância de [RetornaUsuarioUseCase].
func NewRetornaUsuarioUseCase(repoUsuario usuario.UsuarioRepository) *RetornaUsuarioUseCase {
	return &RetornaUsuarioUseCase{repoUsuario: repoUsuario}
}

// Execute retorna os dados do usuário identificado por [RetornaUsuarioInput.UsuarioID].
//
// Retorna um [usuario.Usuario] contendo o usuário encontrado
func (uc *RetornaUsuarioUseCase) Execute(
	ctx context.Context,
	input RetornaUsuarioInput,
) (*usuario.Usuario, error) {
	if err := shared.VerificarAdmin(ctx, uc.repoUsuario, input.LoggedInUserKeycloakID); err != nil {
		return nil, err
	}

	usuario, err := uc.repoUsuario.FindByID(ctx, input.UsuarioID)
	if err != nil {
		return nil, err
	}

	return usuario, nil
}
