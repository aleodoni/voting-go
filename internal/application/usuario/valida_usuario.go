// Package usuario contains the use cases related to user management.
package usuario

import (
	"context"
	"errors"
	"time"

	"github.com/aleodoni/go-ddd/domain"
	domainShared "github.com/aleodoni/voting-go/internal/domain/shared"
	domainUsuario "github.com/aleodoni/voting-go/internal/domain/usuario"
	"github.com/aleodoni/voting-go/internal/platform/id"
	"github.com/nrednav/cuid2"
)

// EnsureUsuarioInput contém os dados necessários para garantir a existência de um usuário.
type EnsureUsuarioInput struct {
	KeycloakID string
	Username   string
	Email      string
	Nome       string
}

// EnsureUsuarioUseCase garante que um usuário autenticado pelo Keycloak existe no sistema.
//
// Regras de negócio:
//   - se o usuário já existir, retorna o registro existente sem alterações
//   - se o usuário não existir, cria o usuário e sua credencial em uma única transação
//   - a credencial é criada com [domainUsuario.Credencial.Ativo] true, sem permissão de voto
//     ou administração — as permissões devem ser concedidas posteriormente por um administrador
type EnsureUsuarioUseCase struct {
	usuarioRepo domainUsuario.UsuarioRepository
	transactor  domainShared.UnitOfWork
}

// NewEnsureUsuarioUseCase cria uma nova instância de [EnsureUsuarioUseCase].
func NewEnsureUsuarioUseCase(
	usuarioRepo domainUsuario.UsuarioRepository,
	transactor domainShared.UnitOfWork,
) *EnsureUsuarioUseCase {
	return &EnsureUsuarioUseCase{
		usuarioRepo: usuarioRepo,
		transactor:  transactor,
	}
}

// Execute garante que o usuário identificado por [EnsureUsuarioInput.KeycloakID] existe no sistema.
//
// Retorna o [domainUsuario.Usuario] existente ou recém-criado, sempre com a
// [domainUsuario.Credencial] associada preenchida.
func (uc *EnsureUsuarioUseCase) Execute(ctx context.Context, input EnsureUsuarioInput) (*domainUsuario.Usuario, error) {
	u, err := uc.usuarioRepo.FindByKeycloakID(ctx, input.KeycloakID)

	if err == nil {
		return u, nil
	}

	if !errors.Is(err, domainUsuario.ErrUserNotFound) {
		return nil, err
	}

	now := time.Now()
	u = &domainUsuario.Usuario{
		AggregateRoot: domain.NewAggregateRoot(cuid2.Generate()),
		KeycloakID:    input.KeycloakID,
		Username:      input.Username,
		Email:         input.Email,
		Nome:          input.Nome,
		CreatedAt:     now,
		UpdatedAt:     now,
	}

	cred := &domainUsuario.Credencial{
		Entity:          domain.Entity[string]{ID: id.New()},
		UsuarioID:       u.ID,
		Ativo:           true,
		PodeVotar:       false,
		PodeAdministrar: false,
		CreatedAt:       now,
		UpdatedAt:       now,
	}

	u.Credencial = cred

	if err := uc.transactor.Do(ctx, func(txCtx context.Context) error {
		return uc.usuarioRepo.Create(txCtx, u)
	}); err != nil {
		return nil, err
	}

	return u, nil
}
