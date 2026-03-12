package persistence

import (
	"context"

	"gorm.io/gorm"
)

type GormTransactor struct {
	db *gorm.DB
}

func NewGormTransactor(db *gorm.DB) *GormTransactor {
	return &GormTransactor{db: db}
}

func (t *GormTransactor) WithTransaction(
	ctx context.Context,
	fn func(ctx context.Context) error,
) error {

	return t.db.WithContext(ctx).Transaction(func(tx *gorm.DB) error {

		txCtx := context.WithValue(ctx, txKey, tx)

		return fn(txCtx)
	})
}
