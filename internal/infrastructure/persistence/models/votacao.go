// Package models defines the data models used in the application.
package models

import "time"

// StatusVotacao representa o enum status_votacao do PostgreSQL.
type StatusVotacao string

const (
	StatusVotacaoA StatusVotacao = "A"
	StatusVotacaoF StatusVotacao = "F"
	StatusVotacaoV StatusVotacao = "V"
	StatusVotacaoC StatusVotacao = "C"
)

type VotacaoModel struct {
	ID        string        `gorm:"primaryKey;column:id;type:varchar"`
	ProjetoID *string       `gorm:"column:projeto_id;type:varchar;uniqueIndex"`
	Status    StatusVotacao `gorm:"column:status;type:status_votacao;not null;default:F"`
	CreatedAt time.Time     `gorm:"column:created_at;autoCreateTime"`
	UpdatedAt time.Time     `gorm:"column:updated_at;autoUpdateTime"`

	Projeto *ProjetoModel `gorm:"foreignKey:ProjetoID"`
	Votos   *[]VotoModel  `gorm:"foreignKey:VotacaoID"`
}

func (VotacaoModel) TableName() string {
	return "votacao"
}
