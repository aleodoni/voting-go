// Package usuario defines the interfaces for data access and manipulation.
package usuario

import "context"

type UsuarioRepository interface {
	FindByKeycloakID(ctx context.Context, keycloakID string) (*Usuario, error)
	FindByUsername(ctx context.Context, username string) (*Usuario, error)
	Create(ctx context.Context, usuario *Usuario) error
}
