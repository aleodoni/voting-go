// Package models defines the data models used in the application.
package models

import "time"

type RestricaoModel struct {
	ID        string    `gorm:"primaryKey;column:id;type:varchar"`
	Restricao string    `gorm:"column:restricao;type:varchar(500);not null"`
	VotoID    string    `gorm:"column:voto_id;type:varchar;not null"`
	CreatedAt time.Time `gorm:"column:created_at;autoCreateTime"`
	UpdatedAt time.Time `gorm:"column:updated_at;autoUpdateTime"`
}

func (RestricaoModel) TableName() string {
	return "restricao"
}
