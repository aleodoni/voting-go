package persistence

import (
	"context"
	"errors"
	"fmt"

	"github.com/aleodoni/voting-go/internal/domain/votacao"
	"github.com/aleodoni/voting-go/internal/infrastructure/persistence/mappers"
	"github.com/aleodoni/voting-go/internal/infrastructure/persistence/models"
	"gorm.io/gorm"
)

type votacaoRepository struct {
	db *gorm.DB
}

func NewVotacaoRepository(db *gorm.DB) votacao.VotacaoRepository {
	return &votacaoRepository{db: db}
}

func (r *votacaoRepository) SalvaVotacao(ctx context.Context, v *votacao.Votacao) error {
	db := DBFromCtx(ctx, r.db)

	model := mappers.ToModelVotacao(v)

	if err := db.Save(&model).Error; err != nil {
		return fmt.Errorf("SalvaVotacao: %w", err)
	}

	return nil
}

func (r *votacaoRepository) DeletaVotacao(ctx context.Context, votacaoID string) error {
	db := DBFromCtx(ctx, r.db)

	if err := db.Delete(&models.VotacaoModel{}, "id = ?", votacaoID).Error; err != nil {
		return fmt.Errorf("DeletaVotacao: %w", err)
	}

	return nil
}

func (r *votacaoRepository) SalvaVoto(ctx context.Context, v *votacao.Voto) error {
	db := DBFromCtx(ctx, r.db)

	var restricao interface{}
	if v.Restricao != nil {
		restricao = map[string]interface{}{
			"restricao_id": v.Restricao.ID,
			"restricao":    v.Restricao.Restricao,
		}
	}

	var votoContrario interface{}
	if v.VotoContrario != nil {
		votoContrario = map[string]interface{}{
			"id_voto_contrario": v.VotoContrario.ID,
			"id_texto":          v.VotoContrario.IDTexto,
			"opinion": map[string]interface{}{
				"parecer_id": v.VotoContrario.ParecerID,
			},
		}
	}

	if err := db.Exec(
		"SELECT f_save_vote(?, ?, ?, ?, ?, ?)",
		v.ID,
		v.UsuarioID,
		v.VotacaoID,
		v.Voto,
		restricao,
		votoContrario,
	).Error; err != nil {
		if IsUniqueViolation(err) {
			return votacao.ErrUsuarioJaVotou
		}
		return fmt.Errorf("SalvaVoto: %w", err)
	}

	return nil
}

func (r *votacaoRepository) BuscaVotacao(ctx context.Context, votacaoID string) (*votacao.Votacao, error) {
	db := DBFromCtx(ctx, r.db)

	var model models.VotacaoModel
	if err := db.First(&model, "id = ?", votacaoID).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, votacao.ErrVotacaoNaoEncontrada
		}
		return nil, fmt.Errorf("BuscaVotacao: %w", err)
	}

	return mappers.ToDomainVotacao(&model), nil
}

func (r *votacaoRepository) UsuarioJaVotou(ctx context.Context, usuarioID, votacaoID string) (bool, error) {
	db := DBFromCtx(ctx, r.db)

	var count int64
	if err := db.Model(&models.VotoModel{}).
		Where("usuario_id = ? AND votacao_id = ?", usuarioID, votacaoID).
		Count(&count).Error; err != nil {
		return false, fmt.Errorf("UsuarioJaVotou: %w", err)
	}

	return count > 0, nil
}
