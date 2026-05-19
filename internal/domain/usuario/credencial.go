package usuario

import (
	"time"

	"github.com/aleodoni/go-ddd/domain"
)

// Credencial representa as permissões de acesso de um usuário no sistema.
type Credencial struct {
	domain.Entity[string]
	UsuarioID       string
	Ativo           bool
	PodeAdministrar bool
	PodeVotar       bool
	CreatedAt       time.Time
	UpdatedAt       time.Time
}

// IsActive informa se a credencial está ativa.
func (c *Credencial) IsActive() bool {
	return c.Ativo
}

// IsAdmin informa se o usuário possui permissão de administrador e está ativo.
func (c *Credencial) IsAdmin() bool {
	return c.Ativo && c.PodeAdministrar
}

// CanVote informa se o usuário possui permissão para votar e está ativo.
func (c *Credencial) CanVote() bool {
	return c.Ativo && c.PodeVotar
}
