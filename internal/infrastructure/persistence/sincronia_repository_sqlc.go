package persistence

import (
	"context"

	"github.com/aleodoni/voting-go/internal/domain/sincronia"
	"github.com/aleodoni/voting-go/internal/infrastructure/persistence/mappers"
	db "github.com/aleodoni/voting-go/internal/infrastructure/persistence/sqlc/generated"
	"github.com/jackc/pgx/v5/pgxpool"
)

type sincroniaRepositorySQLC struct {
	q *db.Queries
}

func NewSincroniaRepositorySQLC(pool *pgxpool.Pool) sincronia.SincroniaRepository {
	return &sincroniaRepositorySQLC{
		q: db.New(pool),
	}
}

func (r *sincroniaRepositorySQLC) queries(ctx context.Context) *db.Queries {
	if tx, ok := TxFromCtx(ctx); ok {
		return r.q.WithTx(tx)
	}
	return r.q
}

func (r *sincroniaRepositorySQLC) Sync(ctx context.Context) (*sincronia.Sincronia, error) {
	q := r.queries(ctx)

	err := q.ExecuteDailySync(ctx)
	if err != nil {
		return nil, err
	}

	row, err := q.GetLastSincronia(ctx)
	if err != nil {
		return nil, err
	}

	return mappers.MapGetLastSincroniaRowToDomain(row), nil
}

func (r *sincroniaRepositorySQLC) ListLastSincronias(
	ctx context.Context,
) ([]*sincronia.Sincronia, error) {
	rows, err := r.queries(ctx).ListLastSincronias(ctx)
	if err != nil {
		return nil, err
	}

	items := make([]*sincronia.Sincronia, 0, len(rows))

	for _, row := range rows {
		items = append(items, mappers.MapSincroniumToDomain(row))
	}

	return items, nil
}
