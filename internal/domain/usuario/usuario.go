// Package usuario defines the Usuario entity and its repository interface.
package usuario

import (
	"time"

	"github.com/aleodoni/voting-go/internal/domain/credencial"
)

type Usuario struct {
	ID           string
	KeycloakID   string
	Username     string
	Email        string
	Nome         string
	NomeFantasia *string

	Credencial *credencial.Credencial

	CreatedAt time.Time
	UpdatedAt time.Time
}
