package persistence

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"

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

func (r *reuniaoRepository) GetProjetosCompleto(ctx context.Context, reuniaoID string) ([]*votacao.Projeto, error) {
	var raw *string
	db := DBFromCtx(ctx, r.db)

	if err := db.
		Raw("SELECT public.f_get_projetos_completo(?)", reuniaoID).
		Scan(&raw).Error; err != nil {
		return nil, fmt.Errorf("GetProjetosCompleto: %w", err)
	}

	if raw == nil {
		return []*votacao.Projeto{}, nil
	}

	var items []ProjetoCompletoJSON

	if err := json.Unmarshal([]byte(*raw), &items); err != nil {
		return nil, fmt.Errorf("GetProjetosCompleto unmarshal: %w", err)
	}

	projetos := make([]*votacao.Projeto, 0, len(items))
	for _, item := range items {
		p, err := ToDomainProjetoFromJSON(item)
		if err != nil {
			return nil, err
		}
		projetos = append(projetos, p)
	}

	return projetos, nil
}

func (r *reuniaoRepository) GetProjetoCompleto(ctx context.Context, projetoID string) (*votacao.Projeto, error) {
	var raw *string // ponteiro para aceitar NULL
	db := DBFromCtx(ctx, r.db)

	if err := db.
		Raw("SELECT public.f_get_projeto_completo(?)", projetoID).
		Scan(&raw).Error; err != nil {
		return nil, fmt.Errorf("GetProjetoCompleto: %w", err)
	}

	if raw == nil {
		return nil, votacao.ErrProjetoNotFound
	}

	var item ProjetoCompletoJSON

	if err := json.Unmarshal([]byte(*raw), &item); err != nil {
		return nil, fmt.Errorf("GetProjetoCompleto unmarshal: %w", err)
	}

	projeto, err := ToDomainProjetoFromJSON(item)
	if err != nil {
		return nil, err
	}

	return projeto, nil
}
