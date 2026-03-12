package mappers

import (
	"github.com/aleodoni/voting-go/internal/domain/votacao"
	"github.com/aleodoni/voting-go/internal/infrastructure/persistence/models"
)

func ToModelVotoContrario(vc *votacao.VotoContrario) *models.VotoContrarioModel {
	return &models.VotoContrarioModel{
		ID:        vc.ID,
		IDTexto:   vc.IDTexto,
		VotoID:    vc.VotoID,
		ParecerID: vc.ParecerID,
		CreatedAt: vc.CreatedAt,
		UpdatedAt: vc.UpdatedAt,
	}
}

func ToDomainVotoContrario(m *models.VotoContrarioModel) *votacao.VotoContrario {
	return &votacao.VotoContrario{
		ID:        m.ID,
		IDTexto:   m.IDTexto,
		VotoID:    m.VotoID,
		ParecerID: m.ParecerID,
		CreatedAt: m.CreatedAt,
		UpdatedAt: m.UpdatedAt,
	}
}
