package votacao

import (
	"context"

	"github.com/aleodoni/voting-go/internal/application/shared"
	"github.com/aleodoni/voting-go/internal/domain/usuario"
	"github.com/aleodoni/voting-go/internal/domain/votacao"
)

// RetornaProjetoCompletoInput contém os dados necessários para retornar um projeto completo.
type RetornaProjetoCompletoInput struct {
	LoggedInUserKeycloakID string
	ProjetoID              string
}

// RetornaProjetoCompletoUseCase retorna um projeto completo pelo ID.
type RetornaProjetoCompletoUseCase struct {
	repoUsuario usuario.UsuarioRepository
	repoReuniao votacao.ReuniaoRepository
}

// NewRetornaProjetoCompletoUseCase cria uma nova instância.
func NewRetornaProjetoCompletoUseCase(
	repoUsuario usuario.UsuarioRepository,
	repoReuniao votacao.ReuniaoRepository,
) *RetornaProjetoCompletoUseCase {
	return &RetornaProjetoCompletoUseCase{
		repoUsuario: repoUsuario,
		repoReuniao: repoReuniao,
	}
}

// Execute retorna o projeto completo.
func (uc *RetornaProjetoCompletoUseCase) Execute(
	ctx context.Context,
	input RetornaProjetoCompletoInput,
) (*votacao.Projeto, error) {
	if err := shared.VerificarAdmin(ctx, uc.repoUsuario, input.LoggedInUserKeycloakID); err != nil {
		return nil, err
	}

	return uc.repoReuniao.GetProjetoCompleto(ctx, input.ProjetoID)
}
