// Package models defines the data models used in the application.
package models

import "time"

type Reuniao struct {
	ID string `gorm:"primaryKey;type:varchar"`

	ConID          int       `gorm:"column:con_id;not null;index:idx_reuniao_unique,unique"`
	RecID          int       `gorm:"column:rec_id;not null;index:idx_reuniao_unique,unique"`
	PacID          int       `gorm:"column:pac_id;not null;index:idx_reuniao_unique,unique"`
	ConDesc        string    `gorm:"column:con_desc;type:varchar;not null"`
	ConSigla       string    `gorm:"column:con_sigla;type:varchar;not null"`
	RecTipoReuniao string    `gorm:"column:rec_tipo_reuniao;type:varchar;not null"`
	RecNumero      string    `gorm:"column:rec_numero;type:varchar;not null"`
	RecData        time.Time `gorm:"column:rec_data;type:date;not null"`
	CreatedAt      time.Time `gorm:"column:created_at;autoCreateTime"`
	UpdatedAt      time.Time `gorm:"column:updated_at;autoUpdateTime"`

	Projetos *[]ProjetoModel `gorm:"foreignKey:ReuniaoID"`
}

func (Reuniao) TableName() string {
	return "reuniao"
}
