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

	db := DBFromCtx(ctx, r.db)

	err := db.
		Preload("Credencial").
		Where("keycloak_id = ?", keycloakID).
		First(&model).Error

	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, usuario.ErrUserNotFound
	}
	if err != nil {
		return nil, err
	}

	return mappers.ToDomainUsuario(&model), nil
}

func (r *usuarioRepository) FindByUsername(ctx context.Context, username string) (*usuario.Usuario, error) {
	var model models.UsuarioModel

	db := DBFromCtx(ctx, r.db)

	err := db.
		Preload("Credencial").
		Where("username = ?", username).
		First(&model).Error

	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, usuario.ErrUserNotFound
	}
	if err != nil {
		return nil, err
	}

	return mappers.ToDomainUsuario(&model), nil
}

func (r *usuarioRepository) Create(ctx context.Context, u *usuario.Usuario) error {
	db := DBFromCtx(ctx, r.db)

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
	return db.Create(model).Error
}

func (r *usuarioRepository) UpdateDisplayNamePermissions(
	ctx context.Context,
	userID string,
	displayName *string,
	isActive bool,
	canAdmin bool,
	canVote bool,
) error {
	db := DBFromCtx(ctx, r.db)

	return db.Exec(`
		SELECT public.f_update_user_with_permissions(?, ?, ?, ?, ?)
	`,
		userID,
		displayName,
		isActive,
		canAdmin,
		canVote,
	).Error
}
