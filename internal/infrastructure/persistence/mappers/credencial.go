package mappers

import (
	"github.com/aleodoni/voting-go/internal/domain/credencial"
	"github.com/aleodoni/voting-go/internal/infrastructure/persistence/models"
)

func ToDomainCredencial(m *models.CredencialModel) *credencial.Credencial {
	return &credencial.Credencial{
		ID:              m.ID,
		UsuarioID:       m.UsuarioID,
		Ativo:           m.Ativo,
		PodeAdministrar: m.PodeAdministrar,
		PodeVotar:       m.PodeVotar,
		CreatedAt:       m.CreatedAt,
		UpdatedAt:       m.UpdatedAt,
	}
}

func ToModelCredencial(c *credencial.Credencial) *models.CredencialModel {
	return &models.CredencialModel{
		ID:              c.ID,
		UsuarioID:       c.UsuarioID,
		Ativo:           c.Ativo,
		PodeAdministrar: c.PodeAdministrar,
		PodeVotar:       c.PodeVotar,
		CreatedAt:       c.CreatedAt,
		UpdatedAt:       c.UpdatedAt,
	}
}
