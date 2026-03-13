// Package models defines the data models used in the application.
package models

import "time"

type ProjetoModel struct {
	ID                string    `gorm:"primaryKey;column:id;type:varchar"`
	Sumula            string    `gorm:"column:sumula;type:varchar;not null"`
	Relator           string    `gorm:"column:relator;type:varchar;not null"`
	TemEmendas        bool      `gorm:"column:tem_emendas;not null"`
	PacID             int       `gorm:"column:pac_id;not null;index:idx_projeto_unique,unique"`
	ParID             int       `gorm:"column:par_id;not null;index:idx_projeto_unique,unique"`
	CodigoProposicao  string    `gorm:"column:codigo_proposicao;type:varchar;not null;index:idx_projeto_unique,unique"`
	Iniciativa        string    `gorm:"column:iniciativa;type:varchar;not null"`
	ConclusaoComissao string    `gorm:"column:conclusao_comissao;type:varchar;not null"`
	ConclusaoRelator  string    `gorm:"column:conclusao_relator;type:varchar;not null"`
	ReuniaoID         string    `gorm:"column:reuniao_id;type:varchar;not null;index:idx_projeto_unique,unique"`
	CreatedAt         time.Time `gorm:"column:created_at;autoCreateTime"`
	UpdatedAt         time.Time `gorm:"column:updated_at;autoUpdateTime"`

	Pareceres *[]ParecerModel `gorm:"foreignKey:ProjetoID"`
	Votacao   *VotacaoModel   `gorm:"foreignKey:ProjetoID"`
}

func (ProjetoModel) TableName() string {
	return "projeto"
}
