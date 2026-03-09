package persistence

import (
	"context"

	"gorm.io/gorm"
)

type txKey struct{}

func DBFromCtx(ctx context.Context, fallback *gorm.DB) *gorm.DB {
	if tx, ok := ctx.Value(txKey{}).(*gorm.DB); ok {
		return tx
	}
	return fallback
}
