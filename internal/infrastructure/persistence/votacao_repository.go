package persistence

import (
	"context"
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
