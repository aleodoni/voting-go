package mappers

import (
	"github.com/aleodoni/go-ddd/domain"
	"github.com/aleodoni/voting-go/internal/domain/votacao"
	"github.com/aleodoni/voting-go/internal/infrastructure/persistence/models"
)

// ToModelVotacao converte a entidade de domínio [votacao.Votacao] para [models.VotacaoModel].
func ToModelVotacao(v *votacao.Votacao) *models.VotacaoModel {
	return &models.VotacaoModel{
		ID:        v.ID,
		ProjetoID: v.ProjetoID,
		Status:    models.StatusVotacao(v.Status),
		CreatedAt: v.CreatedAt,
		UpdatedAt: v.UpdatedAt,
	}
}

// ToDomainVotacao converte um [models.VotacaoModel] para a entidade de domínio [votacao.Votacao].
func ToDomainVotacao(m *models.VotacaoModel) *votacao.Votacao {
	v := &votacao.Votacao{
		AggregateRoot: domain.NewAggregateRoot(m.ID),
		ProjetoID:     m.ProjetoID,
		Status:        votacao.StatusVotacao(m.Status),
		CreatedAt:     m.CreatedAt,
		UpdatedAt:     m.UpdatedAt,
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
