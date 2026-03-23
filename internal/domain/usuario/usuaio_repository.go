// Package usuario defines the interfaces for data access and manipulation.
package usuario

import "context"

type UsuarioRepository interface {
	FindByKeycloakID(ctx context.Context, keycloakID string) (*Usuario, error)
	FindByUsername(ctx context.Context, username string) (*Usuario, error)
	Create(ctx context.Context, usuario *Usuario) error
	UpdateDisplayNamePermissions(
		ctx context.Context,
		userID string,
		displayName *string,
		isActive bool,
		canAdmin bool,
		canVote bool,
	) error
	ListUsers(ctx context.Context, search string, page, limit int) ([]*Usuario, int64, error)
	FindByID(ctx context.Context, id string) (*Usuario, error)
}
