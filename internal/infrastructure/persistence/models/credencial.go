package models

import "time"

type CredencialModel struct {
	ID string `gorm:"column:id;type:varchar;primaryKey"`

	Ativo           bool `gorm:"column:ativo;not null;default:false"`
	PodeAdministrar bool `gorm:"column:pode_administrar;not null;default:false"`
	PodeVotar       bool `gorm:"column:pode_votar;not null;default:false"`

	UsuarioID string `gorm:"column:usuario_id;type:varchar;not null;uniqueIndex"`

	CreatedAt time.Time `gorm:"column:created_at;autoCreateTime"`
	UpdatedAt time.Time `gorm:"column:updated_at;autoUpdateTime"`
}

func (CredencialModel) TableName() string {
	return "credencial"
}
