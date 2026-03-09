package persistence

import (
	"context"
	"errors"

	"github.com/aleodoni/voting-go/internal/domain/usuario"
	"github.com/aleodoni/voting-go/internal/infrastructure/persistence/mappers"
	"github.com/aleodoni/voting-go/internal/infrastructure/persistence/models"
	"gorm.io/gorm"
)

type usuarioRepository struct {
	db *gorm.DB
}

func NewUsuarioRepository(db *gorm.DB) usuario.UsuarioRepository {
	return &usuarioRepository{db: db}
}

func (r *usuarioRepository) FindByKeycloakID(ctx context.Context, keycloakID string) (*usuario.Usuario, error) {
	var model models.UsuarioModel

	err := DBFromCtx(ctx, r.db).
		Preload("Credencial").
		Where("keycloak_id = ?", keycloakID).
		First(&model).Error

	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, usuario.ErrNotFound
	}
	if err != nil {
		return nil, err
	}

	return mappers.ToDomainUsuario(&model), nil
}

func (r *usuarioRepository) FindByUsername(ctx context.Context, username string) (*usuario.Usuario, error) {
	var model models.UsuarioModel

	err := DBFromCtx(ctx, r.db).
		Preload("Credencial").
		Where("username = ?", username).
		First(&model).Error

	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, usuario.ErrNotFound
	}
	if err != nil {
		return nil, err
	}

	return mappers.ToDomainUsuario(&model), nil
}

func (r *usuarioRepository) Create(ctx context.Context, u *usuario.Usuario) error {
	model := &models.UsuarioModel{
		ID:           u.ID,
		KeycloakID:   u.KeycloakID,
		Username:     u.Username,
		Email:        u.Email,
		Nome:         u.Nome,
		NomeFantasia: u.NomeFantasia,
		CreatedAt:    u.CreatedAt,
		UpdatedAt:    u.UpdatedAt,
	}
	return DBFromCtx(ctx, r.db).Create(model).Error
}
