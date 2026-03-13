package mappers

import (
	"github.com/aleodoni/voting-go/internal/domain/votacao"
	"github.com/aleodoni/voting-go/internal/infrastructure/persistence/models"
)

func ToModelProjeto(p *votacao.Projeto) *models.ProjetoModel {
	model := &models.ProjetoModel{
		ID:                p.ID,
		Sumula:            p.Sumula,
		Relator:           p.Relator,
		TemEmendas:        p.TemEmendas,
		PacID:             p.PacID,
		ParID:             p.ParID,
		CodigoProposicao:  p.CodigoProposicao,
		Iniciativa:        p.Iniciativa,
		ConclusaoComissao: p.ConclusaoComissao,
		ConclusaoRelator:  p.ConclusaoRelator,
		ReuniaoID:         p.ReuniaoID,
		CreatedAt:         p.CreatedAt,
		UpdatedAt:         p.UpdatedAt,
	}

	if p.Votacao != nil {
		model.Votacao = ToModelVotacao(p.Votacao)
	}

	return model
}

func ToDomainProjeto(m *models.ProjetoModel) *votacao.Projeto {
	p := &votacao.Projeto{
		ID:                m.ID,
		Sumula:            m.Sumula,
		Relator:           m.Relator,
		TemEmendas:        m.TemEmendas,
		PacID:             m.PacID,
		ParID:             m.ParID,
		CodigoProposicao:  m.CodigoProposicao,
		Iniciativa:        m.Iniciativa,
		ConclusaoComissao: m.ConclusaoComissao,
		ConclusaoRelator:  m.ConclusaoRelator,
		ReuniaoID:         m.ReuniaoID,
		CreatedAt:         m.CreatedAt,
		UpdatedAt:         m.UpdatedAt,
	}

	if m.Pareceres != nil {
		pareceres := make([]votacao.Parecer, len(*m.Pareceres))
		for i, par := range *m.Pareceres {
			pareceres[i] = *ToDomainParecer(&par)
		}
		p.Pareceres = &pareceres
	}

	if m.Votacao != nil {
		p.Votacao = ToDomainVotacao(m.Votacao)
	}

	return p
}
