package persistence

import (
	"context"

	"github.com/jackc/pgx/v5"
)

type sqlcTxKey struct{}

func ContextWithTx(ctx context.Context, tx pgx.Tx) context.Context {
	return context.WithValue(ctx, sqlcTxKey{}, tx)
}

func TxFromCtx(ctx context.Context) (pgx.Tx, bool) {
	tx, ok := ctx.Value(sqlcTxKey{}).(pgx.Tx)
	return tx, ok
}
