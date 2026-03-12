package mappers

import (
	"github.com/aleodoni/voting-go/internal/domain/votacao"
	"github.com/aleodoni/voting-go/internal/infrastructure/persistence/models"
)

func ToModelRestricao(r *votacao.Restricao) *models.RestricaoModel {
	return &models.RestricaoModel{
		ID:        r.ID,
		Restricao: r.Restricao,
		VotoID:    r.VotoID,
		CreatedAt: r.CreatedAt,
		UpdatedAt: r.UpdatedAt,
	}
}

func ToDomainRestricao(m *models.RestricaoModel) *votacao.Restricao {
	return &votacao.Restricao{
		ID:        m.ID,
		Restricao: m.Restricao,
		VotoID:    m.VotoID,
		CreatedAt: m.CreatedAt,
		UpdatedAt: m.UpdatedAt,
	}
}
