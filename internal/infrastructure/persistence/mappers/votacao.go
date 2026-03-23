package mappers

import (
	"github.com/aleodoni/go-ddd/domain"
	"github.com/aleodoni/voting-go/internal/domain/votacao"
	db "github.com/aleodoni/voting-go/internal/infrastructure/persistence/sqlc/generated"
)

// ToDomainVotacaoFromSQLC converte um [db.Votacao] para a entidade de domínio [votacao.Votacao].
func ToDomainVotacaoFromSQLC(m db.Votacao) *votacao.Votacao {
	var projetoID *string
	if m.ProjetoID.Valid {
		projetoID = &m.ProjetoID.String
	}

	return &votacao.Votacao{
		AggregateRoot: domain.NewAggregateRoot(m.ID),
		ProjetoID:     projetoID,
		Status:        votacao.StatusVotacao(m.Status),
		CreatedAt:     m.CreatedAt.Time,
		UpdatedAt:     m.UpdatedAt.Time,
	}
}
