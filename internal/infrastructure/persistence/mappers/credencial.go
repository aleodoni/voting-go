// Package mappers contains functions to convert between domain entities and database models.
package mappers

import (
	"github.com/aleodoni/go-ddd/domain"
	"github.com/aleodoni/voting-go/internal/domain/usuario"
	"github.com/aleodoni/voting-go/internal/infrastructure/persistence/models"
)

func ToDomainCredencial(m *models.CredencialModel) *usuario.Credencial {
	return &usuario.Credencial{
		Entity:          domain.Entity[string]{ID: m.ID},
		UsuarioID:       m.UsuarioID,
		Ativo:           m.Ativo,
		PodeAdministrar: m.PodeAdministrar,
		PodeVotar:       m.PodeVotar,
		CreatedAt:       m.CreatedAt,
		UpdatedAt:       m.UpdatedAt,
	}
}

func ToModelCredencial(c *usuario.Credencial) *models.CredencialModel {
	return &models.CredencialModel{
		ID:              c.Entity.ID,
		UsuarioID:       c.UsuarioID,
		Ativo:           c.Ativo,
		PodeAdministrar: c.PodeAdministrar,
		PodeVotar:       c.PodeVotar,
		CreatedAt:       c.CreatedAt,
		UpdatedAt:       c.UpdatedAt,
	}
}
