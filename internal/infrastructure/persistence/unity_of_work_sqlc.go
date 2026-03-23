package persistence

import (
	"context"

	"github.com/jackc/pgx/v5/pgxpool"
)

type UnitOfWorkSQLC struct {
	db *pgxpool.Pool
}

func NewUnitOfWorkSQLC(db *pgxpool.Pool) *UnitOfWorkSQLC {
	return &UnitOfWorkSQLC{
		db: db,
	}
}

func (u *UnitOfWorkSQLC) Do(
	ctx context.Context,
	fn func(ctx context.Context) error,
) (err error) {
	tx, err := u.db.Begin(ctx)
	if err != nil {
		return err
	}

	defer func() {
		if err != nil {
			_ = tx.Rollback(ctx)
		}
	}()

	ctx = ContextWithTx(ctx, tx)

	if err = fn(ctx); err != nil {
		return err
	}

	err = tx.Commit(ctx)
	return err
}
