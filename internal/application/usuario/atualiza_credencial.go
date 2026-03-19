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
	usuarioRepo    domainUsuario.UsuarioRepository
	credencialRepo domainUsuario.CredencialRepository
}

// NewUpdateCredencialUseCase cria uma nova instância de [UpdateCredencialUseCase].
func NewUpdateCredencialUseCase(
	usuarioRepo domainUsuario.UsuarioRepository,
	credencialRepo domainUsuario.CredencialRepository,
) *UpdateCredencialUseCase {
	return &UpdateCredencialUseCase{
		usuarioRepo:    usuarioRepo,
		credencialRepo: credencialRepo,
	}
}

// Execute atualiza a credencial do usuário informado em [UpdateCredencialInput.UsuarioID].
//
// Retorna a [domainUsuario.Credencial] atualizada em caso de sucesso.
func (uc *UpdateCredencialUseCase) Execute(ctx context.Context, input UpdateCredencialInput) (*domainUsuario.Credencial, error) {
	if err := shared.VerificarAdmin(ctx, uc.usuarioRepo, input.AdminKeycloakID); err != nil {
		return nil, err
	}

	cred, err := uc.credencialRepo.FindByUsuarioID(ctx, input.UsuarioID)
	if err != nil {
		return nil, err
	}

	cred.Ativo = input.Ativo
	cred.PodeVotar = input.PodeVotar
	cred.PodeAdministrar = input.PodeAdministrar

	if err := uc.credencialRepo.Update(ctx, cred); err != nil {
		return nil, err
	}

	return cred, nil
}
