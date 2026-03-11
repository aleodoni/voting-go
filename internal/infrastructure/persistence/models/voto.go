// Package models defines the data models used in the application.
package models

import "time"

// OpcaoVoto representa o enum opcao_voto do PostgreSQL.
type OpcaoVoto string

const (
	OpcaoVotoF OpcaoVoto = "F"
	OpcaoVotoR OpcaoVoto = "R"
	OpcaoVotoC OpcaoVoto = "C"
	OpcaoVotoV OpcaoVoto = "V"
	OpcaoVotoA OpcaoVoto = "A"
)

type VotoModel struct {
	ID        string    `gorm:"primaryKey;column:id;type:varchar"`
	Voto      OpcaoVoto `gorm:"column:voto;type:opcao_voto;not null"`
	VotacaoID string    `gorm:"column:votacao_id;type:varchar;not null"`
	UsuarioID string    `gorm:"column:usuario_id;type:varchar;not null"`
	CreatedAt time.Time `gorm:"column:created_at;autoCreateTime"`
	UpdatedAt time.Time `gorm:"column:updated_at;autoUpdateTime"`

	Restricao     *RestricaoModel     `gorm:"foreignKey:VotoID"`
	VotoContrario *VotoContrarioModel `gorm:"foreignKey:VotoID"`
}

func (VotoModel) TableName() string {
	return "voto"
}
