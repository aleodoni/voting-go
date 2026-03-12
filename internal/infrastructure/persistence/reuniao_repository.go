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
	var raw string
	if err := r.db.WithContext(ctx).
		Raw("SELECT public.f_get_projetos_completo(?)", reuniaoID).
		Scan(&raw).Error; err != nil {
		return nil, fmt.Errorf("GetProjetosCompleto: %w", err)
	}

	var items []ProjetoCompletoJSON
	if err := json.Unmarshal([]byte(raw), &items); err != nil {
		return nil, fmt.Errorf("GetProjetosCompleto unmarshal: %w", err)
	}

	projetos := make([]*votacao.Projeto, 0, len(items))
	for _, item := range items {
		p, err := toDomainProjeto(item)
		if err != nil {
			return nil, err
		}
		projetos = append(projetos, p)
	}

	return projetos, nil
}

func toDomainProjeto(raw ProjetoCompletoJSON) (*votacao.Projeto, error) {
	var pj projetoJSON
	if err := json.Unmarshal(raw.Projeto, &pj); err != nil {
		return nil, fmt.Errorf("unmarshal projeto: %w", err)
	}

	var pareceres []parecerJSONRaw
	if err := json.Unmarshal(raw.Pareceres, &pareceres); err != nil {
		return nil, fmt.Errorf("unmarshal pareceres: %w", err)
	}

	var votacoes []votacaoJSON
	if err := json.Unmarshal(raw.Votacoes, &votacoes); err != nil {
		return nil, fmt.Errorf("unmarshal votacoes: %w", err)
	}

	return &votacao.Projeto{
		ID:                pj.ID,
		Sumula:            pj.Sumula,
		Relator:           pj.Relator,
		TemEmendas:        pj.TemEmendas,
		PacID:             pj.PacID,
		ParID:             pj.ParID,
		CodigoProposicao:  pj.CodigoProposicao,
		Iniciativa:        pj.Iniciativa,
		ConclusaoComissao: pj.ConclusaoComissao,
		ConclusaoRelator:  pj.ConclusaoRelator,
		ReuniaoID:         pj.ReuniaoID,
		CreatedAt:         pj.CreatedAt.Time,
		UpdatedAt:         pj.UpdatedAt.Time,
		Pareceres:         mapParecerSlice(pareceres),
		Votacoes:          mapVotacaoSlice(votacoes),
	}, nil
}
