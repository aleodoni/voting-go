package models

import "time"

type UsuarioModel struct {
	ID           string  `gorm:"column:id;type:varchar;primaryKey"`
	KeycloakID   string  `gorm:"column:keycloak_id;type:varchar;not null;uniqueIndex"`
	Email        string  `gorm:"column:email;type:varchar;not null;uniqueIndex"`
	Nome         string  `gorm:"column:nome;type:varchar;not null"`
	NomeFantasia *string `gorm:"column:nome_fantasia;type:varchar"`
	Username     string  `gorm:"column:username;type:varchar;not null;uniqueIndex"`

	Credencial *CredencialModel `gorm:"foreignKey:UsuarioID;references:ID"`

	CreatedAt time.Time `gorm:"column:created_at;autoCreateTime"`
	UpdatedAt time.Time `gorm:"column:updated_at;autoUpdateTime"`
}

func (UsuarioModel) TableName() string {
	return "usuario"
}
