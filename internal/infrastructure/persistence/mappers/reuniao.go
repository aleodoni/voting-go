// Package mappers provides functions to convert between domain and persistence models.
package mappers

import (
	"github.com/aleodoni/voting-go/internal/domain/votacao"
	db "github.com/aleodoni/voting-go/internal/infrastructure/persistence/sqlc/generated"
)

func ToDomainReuniaoFromSQLC(m db.FindReuniaoByIDRow) *votacao.Reuniao {
	return &votacao.Reuniao{
		ID:             m.ID,
		ConID:          int(m.ConID),
		RecID:          int(m.RecID),
		PacID:          int(m.PacID),
		ConDesc:        m.ConDesc,
		ConSigla:       m.ConSigla,
		RecTipoReuniao: m.RecTipoReuniao,
		RecNumero:      m.RecNumero,
		RecData:        m.RecData.Time,
		CreatedAt:      m.CreatedAt.Time,
		UpdatedAt:      m.UpdatedAt.Time,
	}
}

func ToDomainReuniaoFromDiaRow(m db.GetReunioesDiaRow) *votacao.Reuniao {
	return &votacao.Reuniao{
		ID:             m.ID,
		ConID:          int(m.ConID),
		RecID:          int(m.RecID),
		PacID:          int(m.PacID),
		ConDesc:        m.ConDesc,
		ConSigla:       m.ConSigla,
		RecTipoReuniao: m.RecTipoReuniao,
		RecNumero:      m.RecNumero,
		RecData:        m.RecData.Time,
		CreatedAt:      m.CreatedAt.Time,
		UpdatedAt:      m.UpdatedAt.Time,
	}
}
