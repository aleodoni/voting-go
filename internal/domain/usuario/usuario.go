// Package usuario defines the Usuario entity and its repository interface.
package usuario

import (
	"time"

	"github.com/aleodoni/go-ddd/domain"
)

// Usuario representa o agregado raiz do domínio de usuário.
type Usuario struct {
	domain.AggregateRoot[string]

	KeycloakID   string
	Username     string
	Email        string
	Nome         string
	NomeFantasia *string

	Credencial *Credencial

	CreatedAt time.Time
	UpdatedAt time.Time
}
