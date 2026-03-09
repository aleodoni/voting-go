// Package persistence provides the persistence layer for the application.
package persistence

import (
	"context"
	"errors"

	"github.com/aleodoni/voting-go/internal/domain/credencial"
	"github.com/aleodoni/voting-go/internal/infrastructure/persistence/mappers"
	"github.com/aleodoni/voting-go/internal/infrastructure/persistence/models"
	"gorm.io/gorm"
)

type credencialRepository struct {
	db *gorm.DB
}

func NewCredencialRepository(db *gorm.DB) credencial.CredencialRepository {
	return &credencialRepository{db: db}
}

func (r *credencialRepository) FindByUsuarioID(ctx context.Context, usuarioID string) (*credencial.Credencial, error) {
	var model models.CredencialModel

	err := DBFromCtx(ctx, r.db).
		Where("usuario_id = ?", usuarioID).
		First(&model).Error

	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, credencial.ErrNotFound
	}
	if err != nil {
		return nil, err
	}

	return mappers.ToDomainCredencial(&model), nil
}

func (r *credencialRepository) Create(ctx context.Context, cred *credencial.Credencial) error {
	return DBFromCtx(ctx, r.db).Create(mappers.ToModelCredencial(cred)).Error
}

func (r *credencialRepository) Update(ctx context.Context, cred *credencial.Credencial) error {
	return DBFromCtx(ctx, r.db).
		Model(mappers.ToModelCredencial(cred)).
		Updates(mappers.ToModelCredencial(cred)).Error
}
