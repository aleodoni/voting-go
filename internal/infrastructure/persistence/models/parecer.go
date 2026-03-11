// Package models defines the data models used in the application.
package models

import "time"

type ParecerModel struct {
	ID               string    `gorm:"primaryKey;column:id;type:varchar"`
	CodigoProposicao string    `gorm:"column:codigo_proposicao;type:varchar;not null;index:idx_parecer_unique,unique"`
	TCPNome          string    `gorm:"column:tcp_nome;type:varchar;not null"`
	Vereador         string    `gorm:"column:vereador;type:varchar;not null"`
	IDTexto          int       `gorm:"column:id_texto;not null;index:idx_parecer_unique,unique"`
	ProjetoID        string    `gorm:"column:projeto_id;type:varchar;not null;index:idx_parecer_unique,unique"`
	CreatedAt        time.Time `gorm:"column:created_at;autoCreateTime"`
	UpdatedAt        time.Time `gorm:"column:updated_at;autoUpdateTime"`
}

func (ParecerModel) TableName() string {
	return "parecer"
}
