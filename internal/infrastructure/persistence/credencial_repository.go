// Package persistence provides the persistence layer for the application.
package persistence

import (
	"context"
	"errors"

	"github.com/aleodoni/voting-go/internal/domain/usuario"
	"github.com/aleodoni/voting-go/internal/infrastructure/persistence/mappers"
	"github.com/aleodoni/voting-go/internal/infrastructure/persistence/models"
	"gorm.io/gorm"
)

type credencialRepository struct {
	db *gorm.DB
}

func NewCredencialRepository(db *gorm.DB) usuario.CredencialRepository {
	return &credencialRepository{db: db}
}

func (r *credencialRepository) FindByUsuarioID(ctx context.Context, usuarioID string) (*usuario.Credencial, error) {
	var model models.CredencialModel

	db := DBFromCtx(ctx, r.db)

	err := db.
		Where("usuario_id = ?", usuarioID).
		First(&model).Error

	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, usuario.ErrCredencialNotFound
	}
	if err != nil {
		return nil, err
	}

	return mappers.ToDomainCredencial(&model), nil
}

func (r *credencialRepository) Create(ctx context.Context, cred *usuario.Credencial) error {
	db := DBFromCtx(ctx, r.db)

	return db.Create(mappers.ToModelCredencial(cred)).Error
}

func (r *credencialRepository) Update(ctx context.Context, cred *usuario.Credencial) error {
	db := DBFromCtx(ctx, r.db)
	model := mappers.ToModelCredencial(cred)

	return db.
		Model(&models.CredencialModel{}).
		Where("usuario_id = ?", cred.UsuarioID).
		Updates(model).Error
}
