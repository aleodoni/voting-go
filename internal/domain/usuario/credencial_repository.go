package usuario

import "context"

// CredencialRepository define o contrato de persistência para [Credencial].
type CredencialRepository interface {
	// FindByUsuarioID retorna a credencial associada ao usuário informado.
	FindByUsuarioID(ctx context.Context, usuarioID string) (*Credencial, error)

	// Create persiste uma nova credencial no repositório.
	Create(ctx context.Context, cred *Credencial) error

	// Update atualiza os dados de uma credencial existente no repositório.
	Update(ctx context.Context, cred *Credencial) error
}
