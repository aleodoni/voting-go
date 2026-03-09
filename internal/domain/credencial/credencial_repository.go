// Package credencial defines the interfaces for data access and manipulation.
package credencial

import "context"

type CredencialRepository interface {
	FindByUsuarioID(ctx context.Context, usuarioID string) (*Credencial, error)
	Create(ctx context.Context, cred *Credencial) error
	Update(ctx context.Context, cred *Credencial) error
}
