package persistence

import (
	"context"
	"errors"

	"github.com/aleodoni/voting-go/internal/domain/votacao"
	"github.com/aleodoni/voting-go/internal/infrastructure/persistence/mappers"
	"github.com/aleodoni/voting-go/internal/infrastructure/persistence/models"
	"gorm.io/gorm"
)

type reuniaoRepository struct {
	db *gorm.DB
}

func NewReuniaoRepository(db *gorm.DB) votacao.ReuniaoRepository {
	return &reuniaoRepository{db: db}
}

func (r *reuniaoRepository) FindReuniaoByID(ctx context.Context, reuniaoID string) (*votacao.Reuniao, error) {
	var model models.Reuniao

	db := DBFromCtx(ctx, r.db)

	err := db.Where("id = ?", reuniaoID).First(&model).Error

	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, votacao.ErrReuniaoNotFound
	}
	if err != nil {
		return nil, err
	}

	return mappers.ToDomainReuniao(&model), nil
}

func (r *reuniaoRepository) GetReunioesDia(ctx context.Context) ([]*votacao.Reuniao, error) {
	var model []*models.Reuniao

	db := DBFromCtx(ctx, r.db)

	if err := db.Find(&model).Table("v_reunioes_hoje").Error; err != nil {
		return nil, err
	}

	return mappers.ToDomainReunioes(model), nil
}
