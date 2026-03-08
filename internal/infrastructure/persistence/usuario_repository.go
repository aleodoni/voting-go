// Package persistence provides repository implementations for persistence layer.
package persistence

import (
	"github.com/aleodoni/voting-go/internal/domain/usuario"
	"gorm.io/gorm"
)

type usuarioRepository struct {
	db *gorm.DB
}

func NewUsuarioRepository(db *gorm.DB) usuario.UsuarioRepository {
	return &usuarioRepository{db: db}
}

func (r *usuarioRepository) FindByKeycloakID(keycloakID string) (*usuario.Usuario, error) {

	var user usuario.Usuario

	err := r.db.
		Preload("Credencial").
		Where("keycloak_id = ?", keycloakID).
		First(&user).Error

	if err != nil {
		return nil, err
	}

	return &user, nil
}

func (r *usuarioRepository) FindByUsername(username string) (*usuario.Usuario, error) {

	var user usuario.Usuario

	err := r.db.
		Preload("Credencial").
		Where("username = ?", username).
		First(&user).Error

	if err != nil {
		return nil, err
	}

	return &user, nil
}

func (r *usuarioRepository) Create(user *usuario.Usuario) error {
	return r.db.Create(user).Error
}
