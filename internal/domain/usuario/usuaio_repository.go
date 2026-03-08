// Package usuario defines the interfaces for data access and manipulation.
package usuario

type UsuarioRepository interface {
	FindByKeycloakID(keycloakID string) (*Usuario, error)
	FindByUsername(username string) (*Usuario, error)
	Create(usuario *Usuario) error
}
