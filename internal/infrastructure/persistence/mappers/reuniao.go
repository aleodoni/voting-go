// Package mappers provides functions to convert between domain and persistence models.
package mappers

import (
	"github.com/aleodoni/voting-go/internal/domain/votacao"
	"github.com/aleodoni/voting-go/internal/infrastructure/persistence/models"
)

func ToModelReuniao(r *votacao.Reuniao) *models.Reuniao {
	return &models.Reuniao{
		ID:             r.ID,
		ConID:          r.ConID,
		RecID:          r.RecID,
		PacID:          r.PacID,
		ConDesc:        r.ConDesc,
		ConSigla:       r.ConSigla,
		RecTipoReuniao: r.RecTipoReuniao,
		RecNumero:      r.RecNumero,
		RecData:        r.RecData,
		CreatedAt:      r.CreatedAt,
		UpdatedAt:      r.UpdatedAt,
	}
}

func ToDomainReuniao(m *models.Reuniao) *votacao.Reuniao {
	r := &votacao.Reuniao{
		ID:             m.ID,
		ConID:          m.ConID,
		RecID:          m.RecID,
		PacID:          m.PacID,
		ConDesc:        m.ConDesc,
		ConSigla:       m.ConSigla,
		RecTipoReuniao: m.RecTipoReuniao,
		RecNumero:      m.RecNumero,
		RecData:        m.RecData,
		CreatedAt:      m.CreatedAt,
		UpdatedAt:      m.UpdatedAt,
	}

	if m.Projetos != nil {
		projetos := make([]votacao.Projeto, len(*m.Projetos))
		for i, p := range *m.Projetos {
			projetos[i] = *ToDomainProjeto(&p)
		}
		r.Projetos = &projetos
	}

	return r
}

func ToDomainReunioes(models []*models.Reuniao) []*votacao.Reuniao {
	reunioes := make([]*votacao.Reuniao, len(models))
	for i, m := range models {
		reunioes[i] = ToDomainReuniao(m)
	}
	return reunioes
}
