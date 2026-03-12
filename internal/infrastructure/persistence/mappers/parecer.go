// Package mappers provides functions to convert between domain and persistence models.
package mappers

import (
	"github.com/aleodoni/voting-go/internal/domain/votacao"
	"github.com/aleodoni/voting-go/internal/infrastructure/persistence/models"
)

func ToModelParecer(p *votacao.Parecer) *models.ParecerModel {
	return &models.ParecerModel{
		ID:               p.ID,
		CodigoProposicao: p.CodigoProposicao,
		TCPNome:          p.TCPNome,
		Vereador:         p.Vereador,
		IDTexto:          p.IDTexto,
		ProjetoID:        p.ProjetoID,
		CreatedAt:        p.CreatedAt,
		UpdatedAt:        p.UpdatedAt,
	}
}

func ToDomainParecer(m *models.ParecerModel) *votacao.Parecer {
	return &votacao.Parecer{
		ID:               m.ID,
		CodigoProposicao: m.CodigoProposicao,
		TCPNome:          m.TCPNome,
		Vereador:         m.Vereador,
		IDTexto:          m.IDTexto,
		ProjetoID:        m.ProjetoID,
		CreatedAt:        m.CreatedAt,
		UpdatedAt:        m.UpdatedAt,
	}
}
