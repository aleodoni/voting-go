package mappers

import (
	"github.com/aleodoni/voting-go/internal/domain/votacao"
	"github.com/aleodoni/voting-go/internal/infrastructure/persistence/models"
)

func ToModelVoto(v *votacao.Voto) *models.VotoModel {
	return &models.VotoModel{
		ID:        v.ID,
		Voto:      models.OpcaoVoto(v.Voto),
		VotacaoID: v.VotacaoID,
		UsuarioID: v.UsuarioID,
		CreatedAt: v.CreatedAt,
		UpdatedAt: v.UpdatedAt,
	}
}

func ToDomainVoto(m *models.VotoModel) *votacao.Voto {
	v := &votacao.Voto{
		ID:        m.ID,
		Voto:      votacao.OpcaoVoto(m.Voto),
		VotacaoID: m.VotacaoID,
		UsuarioID: m.UsuarioID,
		CreatedAt: m.CreatedAt,
		UpdatedAt: m.UpdatedAt,
	}

	if m.Restricao != nil {
		v.Restricao = ToDomainRestricao(m.Restricao)
	}

	if m.VotoContrario != nil {
		v.VotoContrario = ToDomainVotoContrario(m.VotoContrario)
	}

	return v
}
