// Package usuario defines the Usuario entity and its repository interface.
package usuario

import (
	"time"
)

type Usuario struct {
	ID           string
	KeycloakID   string
	Username     string
	Email        string
	Nome         string
	NomeFantasia *string

	Credencial *Credencial

	CreatedAt time.Time
	UpdatedAt time.Time
}
