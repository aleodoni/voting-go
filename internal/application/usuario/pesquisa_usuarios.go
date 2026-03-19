package usuario

import (
	"context"

	"github.com/aleodoni/voting-go/internal/application/shared"
	"github.com/aleodoni/voting-go/internal/domain/usuario"
)

// ListUsuariosInput contém os dados necessários para listar usuários.
type ListUsuariosInput struct {
	LoggedInUserKeycloakID string
	Search                 string
	Page                   int
	Limit                  int
}

// ListUsuariosUseCase retorna uma lista paginada de usuários.
//
// Regras de negócio:
//   - o usuário autenticado deve ser administrador ativo
//   - [ListUsuariosInput.Page] menor que 1 é corrigido para 1
//   - [ListUsuariosInput.Limit] menor que 1 ou maior que 100 é corrigido para 20
type ListUsuariosUseCase struct {
	repoUsuario usuario.UsuarioRepository
}

// NewListUsuariosUseCase cria uma nova instância de [ListUsuariosUseCase].
func NewListUsuariosUseCase(repoUsuario usuario.UsuarioRepository) *ListUsuariosUseCase {
	return &ListUsuariosUseCase{repoUsuario: repoUsuario}
}

// Execute retorna uma lista paginada de usuários, opcionalmente filtrada pelo termo
// informado em [ListUsuariosInput.Search].
//
// Retorna um [usuario.ListUsuario] contendo os usuários encontrados, o total de registros
// e os parâmetros de paginação efetivamente aplicados.
func (uc *ListUsuariosUseCase) Execute(
	ctx context.Context,
	input ListUsuariosInput,
) (*usuario.ListUsuario, error) {
	if err := shared.VerificarAdmin(ctx, uc.repoUsuario, input.LoggedInUserKeycloakID); err != nil {
		return nil, err
	}

	page := input.Page
	if page < 1 {
		page = 1
	}

	limit := input.Limit
	if limit < 1 || limit > 100 {
		limit = 20
	}

	usuarios, total, err := uc.repoUsuario.ListUsers(ctx, input.Search, page, limit)
	if err != nil {
		return nil, err
	}

	return &usuario.ListUsuario{
		Usuarios: usuarios,
		Total:    total,
		Page:     page,
		Limit:    limit,
	}, nil
}
