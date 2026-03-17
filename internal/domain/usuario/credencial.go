package usuario

import "time"

type Credencial struct {
	ID              string
	UsuarioID       string
	Ativo           bool
	PodeAdministrar bool
	PodeVotar       bool
	CreatedAt       time.Time
	UpdatedAt       time.Time
}

func (c *Credencial) IsActive() bool {
	return c.Ativo
}

func (c *Credencial) IsAdmin() bool {
	return c.Ativo && c.PodeAdministrar
}

func (c *Credencial) CanVote() bool {
	return c.Ativo && c.PodeVotar
}
