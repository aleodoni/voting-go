// Package models defines the data models used in the application.
package models

import "time"

type VotoContrarioModel struct {
	ID        string    `gorm:"primaryKey;column:id;type:varchar"`
	IDTexto   int       `gorm:"column:id_texto;not null"`
	VotoID    string    `gorm:"column:voto_id;type:varchar;not null"`
	ParecerID string    `gorm:"column:parecer_id;type:varchar;not null"`
	CreatedAt time.Time `gorm:"column:created_at;autoCreateTime"`
	UpdatedAt time.Time `gorm:"column:updated_at;autoUpdateTime"`
}

func (VotoContrarioModel) TableName() string {
	return "voto_contrario"
}
