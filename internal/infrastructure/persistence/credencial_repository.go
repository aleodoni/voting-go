// Package persistence provides repository implementations for persistence layer.
package persistence

import (
	"github.com/aleodoni/voting-go/internal/domain/credencial"
	"gorm.io/gorm"
)

type credencialRepository struct {
	db *gorm.DB
}

func NewCredencialRepository(db *gorm.DB) credencial.CredencialRepository {
	return &credencialRepository{db: db}
}

func (r *credencialRepository) FindByUsuarioID(usuarioID string) (*credencial.Credencial, error) {

	var cred credencial.Credencial

	err := r.db.
		Where("usuario_id = ?", usuarioID).
		First(&cred).Error

	if err != nil {
		return nil, err
	}

	return &cred, nil
}

func (r *credencialRepository) Create(cred *credencial.Credencial) error {
	return r.db.Create(cred).Error
}
