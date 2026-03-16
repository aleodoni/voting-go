package usuario

import (
	"context"

	"github.com/aleodoni/voting-go/internal/application/shared"
	"github.com/aleodoni/voting-go/internal/domain/usuario"
)

type ListUsuariosInput struct {
	LoggedInUserKeycloakID string
	Search                 string
	Page                   int
	Limit                  int
}

type ListUsuariosOutput struct {
	Usuarios []*usuario.Usuario
	Total    int64
	Page     int
	Limit    int
}

type ListUsuariosUseCase struct {
	repoUsuario usuario.UsuarioRepository
}

func NewListUsuariosUseCase(repoUsuario usuario.UsuarioRepository) *ListUsuariosUseCase {
	return &ListUsuariosUseCase{repoUsuario: repoUsuario}
}

func (uc *ListUsuariosUseCase) Execute(
	ctx context.Context,
	input ListUsuariosInput,
) (*ListUsuariosOutput, error) {
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

	return &ListUsuariosOutput{
		Usuarios: usuarios,
		Total:    total,
		Page:     page,
		Limit:    limit,
	}, nil
}
