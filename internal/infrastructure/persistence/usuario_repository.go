package persistence

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/aleodoni/go-ddd/domain"
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

func (r *usuarioRepository) ListUsers(ctx context.Context, search string, page, limit int) ([]*usuario.Usuario, int64, error) {
	db := DBFromCtx(ctx, r.db)

	type row struct {
		ID              string    `gorm:"column:id"`
		KeycloakID      string    `gorm:"column:keycloak_id"`
		Nome            string    `gorm:"column:nome"`
		NomeFantasia    string    `gorm:"column:nome_fantasia"`
		Email           string    `gorm:"column:email"`
		Ativo           bool      `gorm:"column:ativo"`
		PodeAdministrar bool      `gorm:"column:pode_administrar"`
		PodeVotar       bool      `gorm:"column:pode_votar"`
		UpdatedAt       time.Time `gorm:"column:updated_at"`
		TotalCount      int64     `gorm:"column:total_count"`
	}

	var rows []row
	offset := (page - 1) * limit

	if err := db.Raw(`
		SELECT * FROM public.f_get_users_paginated_with_total(?, ?, ?, ?)
	`, search, search, offset, limit).Scan(&rows).Error; err != nil {
		return nil, 0, fmt.Errorf("ListUsuarios: %w", err)
	}

	if len(rows) == 0 {
		return []*usuario.Usuario{}, 0, nil
	}

	usuarios := make([]*usuario.Usuario, len(rows))
	for i, r := range rows {
		var nomeFantasia *string
		if r.NomeFantasia != "" {
			nomeFantasia = &r.NomeFantasia
		}

		usuarios[i] = &usuario.Usuario{
			AggregateRoot: domain.AggregateRoot[string]{
				Entity: domain.Entity[string]{ID: r.ID},
			},
			KeycloakID:   r.KeycloakID,
			Nome:         r.Nome,
			NomeFantasia: nomeFantasia,
			Email:        r.Email,
			UpdatedAt:    r.UpdatedAt,
			Credencial: &usuario.Credencial{
				Ativo:           r.Ativo,
				PodeAdministrar: r.PodeAdministrar,
				PodeVotar:       r.PodeVotar,
			},
		}
	}

	return usuarios, rows[0].TotalCount, nil
}
