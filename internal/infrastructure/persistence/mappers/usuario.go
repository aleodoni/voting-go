// Package mappers provides functions to convert between persistence models and domain entities.
package mappers

import (
	"github.com/aleodoni/voting-go/internal/domain/usuario"
	"github.com/aleodoni/voting-go/internal/infrastructure/persistence/models"
)

func ToDomainUsuario(m *models.UsuarioModel) *usuario.Usuario {

	u := &usuario.Usuario{
		ID:           m.ID,
		KeycloakID:   m.KeycloakID,
		Username:     m.Username,
		Email:        m.Email,
		Nome:         m.Nome,
		NomeFantasia: m.NomeFantasia,
		CreatedAt:    m.CreatedAt,
		UpdatedAt:    m.UpdatedAt,
	}

	if m.Credencial != nil {
		u.Credencial = ToDomainCredencial(m.Credencial)
	}

	return u
}
