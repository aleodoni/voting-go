package persistence

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"

	"github.com/aleodoni/voting-go/internal/domain/votacao"
	"github.com/aleodoni/voting-go/internal/infrastructure/persistence/mappers"
	"github.com/aleodoni/voting-go/internal/infrastructure/persistence/sqlc/enums"
	db "github.com/aleodoni/voting-go/internal/infrastructure/persistence/sqlc/generated"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgtype"
	"github.com/jackc/pgx/v5/pgxpool"
)

type votacaoRepositorySQLC struct {
	q *db.Queries
}

func NewVotacaoRepositorySQLC(pool *pgxpool.Pool) votacao.VotacaoRepository {
	return &votacaoRepositorySQLC{
		q: db.New(pool),
	}
}

func (r *votacaoRepositorySQLC) queries(ctx context.Context) *db.Queries {
	if tx, ok := TxFromCtx(ctx); ok {
		return r.q.WithTx(tx)
	}
	return r.q
}

func (r *votacaoRepositorySQLC) SalvaVotacao(ctx context.Context, v *votacao.Votacao) error {
	err := r.queries(ctx).UpsertVotacao(ctx, db.UpsertVotacaoParams{
		ID:        v.ID,
		ProjetoID: pgtype.Text{String: derefOrEmpty(v.ProjetoID), Valid: v.ProjetoID != nil},
		Status:    enums.StatusVotacao(v.Status),
	})
	if err != nil {
		return fmt.Errorf("SalvaVotacao: %w", err)
	}
	return nil
}

func (r *votacaoRepositorySQLC) DeletaVotacao(ctx context.Context, votacaoID string) error {
	if err := r.queries(ctx).DeleteVotacao(ctx, votacaoID); err != nil {
		return fmt.Errorf("DeletaVotacao: %w", err)
	}
	return nil
}

func (r *votacaoRepositorySQLC) BuscaVotacao(ctx context.Context, votacaoID string) (*votacao.Votacao, error) {
	row, err := r.queries(ctx).FindVotacaoByID(ctx, votacaoID)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, votacao.ErrVotacaoNaoEncontrada
		}
		return nil, fmt.Errorf("BuscaVotacao: %w", err)
	}
	return mappers.ToDomainVotacaoFromSQLC(row), nil
}

func (r *votacaoRepositorySQLC) GetVotacaoAberta(ctx context.Context) (*votacao.Votacao, error) {
	row, err := r.queries(ctx).FindVotacaoAberta(ctx, enums.StatusVotacaoA)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, nil
		}
		return nil, fmt.Errorf("GetVotacaoAberta: %w", err)
	}
	return mappers.ToDomainVotacaoFromSQLC(row), nil
}

func (r *votacaoRepositorySQLC) UsuarioJaVotou(ctx context.Context, usuarioID, votacaoID string) (bool, error) {
	already, err := r.queries(ctx).UsuarioJaVotou(ctx, db.UsuarioJaVotouParams{
		UsuarioID: usuarioID,
		VotacaoID: votacaoID,
	})
	if err != nil {
		return false, fmt.Errorf("UsuarioJaVotou: %w", err)
	}
	return already, nil
}

func (r *votacaoRepositorySQLC) SalvaVoto(ctx context.Context, v *votacao.Voto) error {
	restricaoJSON, err := marshalVotoField(v.Restricao, func(res *votacao.Restricao) any {
		return map[string]any{
			"restricao_id": res.ID,
			"restricao":    res.Restricao,
		}
	})
	if err != nil {
		return fmt.Errorf("SalvaVoto marshal restricao: %w", err)
	}

	votoContrarioJSON, err := marshalVotoField(v.VotoContrario, func(vc *votacao.VotoContrario) any {
		return map[string]any{
			"id_voto_contrario": vc.ID,
			"id_texto":          vc.IDTexto,
			"opinion": map[string]any{
				"parecer_id": vc.ParecerID,
			},
		}
	})
	if err != nil {
		return fmt.Errorf("SalvaVoto marshal voto_contrario: %w", err)
	}

	err = r.queries(ctx).SaveVoto(ctx, db.SaveVotoParams{
		PVotoID:        v.ID,
		PUsuarioID:     v.UsuarioID,
		PVotacaoID:     v.VotacaoID,
		PVoto:          enums.OpcaoVoto(v.Voto),
		PRestricao:     restricaoJSON,
		PVotoContrario: votoContrarioJSON,
	})

	if err != nil {
		if IsUniqueViolation(err) {
			return votacao.ErrUsuarioJaVotou
		}
		return fmt.Errorf("SalvaVoto: %w", err)
	}
	return nil
}

// marshalVotoField serializa um campo opcional para JSON ou retorna nil.
func marshalVotoField[T any](v *T, fn func(*T) any) ([]byte, error) {
	if v == nil {
		return nil, nil
	}
	return json.Marshal(fn(v))
}
