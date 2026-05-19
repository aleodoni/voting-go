package mappers

import (
	"github.com/aleodoni/go-ddd/domain"
	"github.com/aleodoni/voting-go/internal/domain/votacao"
	db "github.com/aleodoni/voting-go/internal/infrastructure/persistence/sqlc/generated"
	"github.com/jackc/pgx/v5/pgtype"
)

// função base (reutilizável)
func toDomainVotacao(
	id string,
	projetoID pgtype.Text,
	status string,
	createdAt pgtype.Timestamp,
	updatedAt pgtype.Timestamp,
) *votacao.Votacao {

	var pid *string
	if projetoID.Valid {
		pid = &projetoID.String
	}

	return &votacao.Votacao{
		AggregateRoot: domain.NewAggregateRoot(id),
		ProjetoID:     pid,
		Status:        votacao.StatusVotacao(status),
		CreatedAt:     createdAt.Time,
		UpdatedAt:     updatedAt.Time,
	}
}

// mapper para FindVotacaoAberta
func ToDomainVotacaoFromFindAberta(m db.FindVotacaoAbertaRow) *votacao.Votacao {
	return toDomainVotacao(
		m.ID,
		m.ProjetoID,
		m.Status,
		m.CreatedAt,
		m.UpdatedAt,
	)
}

// mapper para FindVotacaoByID
func ToDomainVotacaoFromFindByID(m db.FindVotacaoByIDRow) *votacao.Votacao {
	return toDomainVotacao(
		m.ID,
		m.ProjetoID,
		m.Status,
		m.CreatedAt,
		m.UpdatedAt,
	)
}
