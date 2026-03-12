package mappers

import (
	"github.com/aleodoni/voting-go/internal/domain/votacao"
	"github.com/aleodoni/voting-go/internal/infrastructure/persistence/models"
)

func ToModelVotacao(v *votacao.Votacao) *models.VotacaoModel {
	return &models.VotacaoModel{
		ID:        v.ID,
		ProjetoID: v.ProjetoID,
		Status:    models.StatusVotacao(v.Status),
		CreatedAt: v.CreatedAt,
		UpdatedAt: v.UpdatedAt,
	}
}

func ToDomainVotacao(m *models.VotacaoModel) *votacao.Votacao {
	v := &votacao.Votacao{
		ID:        m.ID,
		ProjetoID: m.ProjetoID,
		Status:    votacao.StatusVotacao(m.Status),
		CreatedAt: m.CreatedAt,
		UpdatedAt: m.UpdatedAt,
	}

	if m.Votos != nil {
		votos := make([]votacao.Voto, len(*m.Votos))
		for i, voto := range *m.Votos {
			votos[i] = *ToDomainVoto(&voto)
		}
		v.Votos = &votos
	}

	return v
}
