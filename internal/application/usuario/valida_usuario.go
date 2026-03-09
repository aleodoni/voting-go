// Package usuario contains the use cases related to user management.
package usuario

import (
	"context"
	"errors"
	"time"

	domainCredencial "github.com/aleodoni/voting-go/internal/domain/credencial"
	domainUsuario "github.com/aleodoni/voting-go/internal/domain/usuario"
	"github.com/aleodoni/voting-go/internal/platform/id"
	"github.com/aleodoni/voting-go/internal/platform/transaction"
)

type EnsureUsuarioUseCase struct {
	usuarioRepo    domainUsuario.UsuarioRepository
	credencialRepo domainCredencial.CredencialRepository
	transactor     transaction.Transactor
}

func NewEnsureUsuarioUseCase(
	usuarioRepo domainUsuario.UsuarioRepository,
	credencialRepo domainCredencial.CredencialRepository,
	transactor transaction.Transactor,
) *EnsureUsuarioUseCase {
	return &EnsureUsuarioUseCase{
		usuarioRepo:    usuarioRepo,
		credencialRepo: credencialRepo,
		transactor:     transactor,
	}
}

type EnsureUsuarioInput struct {
	KeycloakID string
	Username   string
	Email      string
	Nome       string
}

func (uc *EnsureUsuarioUseCase) Execute(ctx context.Context, input EnsureUsuarioInput) (*domainUsuario.Usuario, error) {
	u, err := uc.usuarioRepo.FindByKeycloakID(ctx, input.KeycloakID)
	if err == nil {
		return u, nil
	}
	if !errors.Is(err, domainUsuario.ErrNotFound) {
		return nil, err
	}

	now := time.Now()
	u = &domainUsuario.Usuario{
		ID:         id.New(),
		KeycloakID: input.KeycloakID,
		Username:   input.Username,
		Email:      input.Email,
		Nome:       input.Nome,
		CreatedAt:  now,
		UpdatedAt:  now,
	}

	cred := &domainCredencial.Credencial{
		ID:              id.New(),
		UsuarioID:       u.ID,
		Ativo:           true,
		PodeVotar:       false,
		PodeAdministrar: false,
		CreatedAt:       now,
		UpdatedAt:       now,
	}

	err = uc.transactor.WithTransaction(ctx, func(txCtx context.Context) error {
		if err := uc.usuarioRepo.Create(txCtx, u); err != nil {
			return err
		}
		return uc.credencialRepo.Create(txCtx, cred)
	})
	if err != nil {
		return nil, err
	}

	u.Credencial = cred
	return u, nil
}
