package persistence

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"

	"github.com/aleodoni/voting-go/internal/domain/votacao"
	"github.com/aleodoni/voting-go/internal/infrastructure/persistence/mappers"
	db "github.com/aleodoni/voting-go/internal/infrastructure/persistence/sqlc/generated"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

type reuniaoRepositorySQLC struct {
	q *db.Queries
}

func NewReuniaoRepositorySQLC(pool *pgxpool.Pool) votacao.ReuniaoRepository {
	return &reuniaoRepositorySQLC{
		q: db.New(pool),
	}
}

func (r *reuniaoRepositorySQLC) queries(ctx context.Context) *db.Queries {
	if tx, ok := TxFromCtx(ctx); ok {
		return r.q.WithTx(tx)
	}
	return r.q
}

func (r *reuniaoRepositorySQLC) FindReuniaoByID(ctx context.Context, reuniaoID string) (*votacao.Reuniao, error) {
	row, err := r.queries(ctx).FindReuniaoByID(ctx, reuniaoID)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, votacao.ErrReuniaoNotFound
		}
		return nil, fmt.Errorf("FindReuniaoByID: %w", err)
	}
	return mappers.ToDomainReuniaoFromSQLC(row), nil
}

func (r *reuniaoRepositorySQLC) GetReunioesDia(ctx context.Context) ([]*votacao.Reuniao, error) {
	rows, err := r.queries(ctx).GetReunioesDia(ctx)
	if err != nil {
		return nil, fmt.Errorf("GetReunioesDia: %w", err)
	}

	reunioes := make([]*votacao.Reuniao, len(rows))
	for i, row := range rows {
		reunioes[i] = mappers.ToDomainReuniaoFromDiaRow(row)
	}
	return reunioes, nil
}

func (r *reuniaoRepositorySQLC) GetProjetosCompleto(ctx context.Context, reuniaoID string) ([]*votacao.Projeto, error) {
	raw, err := r.queries(ctx).GetProjetosCompleto(ctx, reuniaoID)
	if err != nil {
		return nil, fmt.Errorf("GetProjetosCompleto: %w", err)
	}

	if raw == "" || raw == "[]" {
		return []*votacao.Projeto{}, nil
	}

	var items []ProjetoCompletoJSON
	if err := json.Unmarshal([]byte(raw), &items); err != nil {
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

func (r *reuniaoRepositorySQLC) GetProjetoCompleto(ctx context.Context, projetoID string) (*votacao.Projeto, error) {
	raw, err := r.queries(ctx).GetProjetoCompleto(ctx, projetoID)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, votacao.ErrProjetoNotFound
		}
		return nil, fmt.Errorf("GetProjetoCompleto: %w", err)
	}

	if raw == "" || raw == "{}" {
		return nil, votacao.ErrProjetoNotFound
	}

	var item ProjetoCompletoJSON
	if err := json.Unmarshal([]byte(raw), &item); err != nil {
		return nil, fmt.Errorf("GetProjetoCompleto unmarshal: %w", err)
	}

	return ToDomainProjetoFromJSON(item)
}
