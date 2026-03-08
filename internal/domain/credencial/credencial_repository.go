// Package credencial defines the interfaces for data access and manipulation.
package credencial

type CredencialRepository interface {
	FindByUsuarioID(usuarioID string) (*Credencial, error)
	Create(cred *Credencial) error
}
