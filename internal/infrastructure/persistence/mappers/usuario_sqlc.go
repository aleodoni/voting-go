package mappers

import (
	"github.com/aleodoni/go-ddd/domain"
	"github.com/aleodoni/voting-go/internal/domain/usuario"
	db "github.com/aleodoni/voting-go/internal/infrastructure/persistence/sqlc/generated"
)

func MapFindByKeycloakIDRowToDomain(row db.FindByKeycloakIDRow) *usuario.Usuario {
	var nomeFantasia *string
	if row.NomeFantasia.Valid {
		nomeFantasia = &row.NomeFantasia.String
	}

	return &usuario.Usuario{
		AggregateRoot: domain.AggregateRoot[string]{
			Entity: domain.Entity[string]{ID: row.ID},
		},
		KeycloakID:   row.KeycloakID,
		Nome:         row.Nome,
		NomeFantasia: nomeFantasia,
		Email:        row.Email,
		Username:     row.Username,
		Credencial: &usuario.Credencial{
			Ativo:           row.Ativo,
			PodeAdministrar: row.PodeAdministrar,
			PodeVotar:       row.PodeVotar,
		},
	}
}

func MapFindByUsernameRowToDomain(row db.FindByUsernameRow) *usuario.Usuario {
	var nomeFantasia *string
	if row.NomeFantasia.Valid {
		nomeFantasia = &row.NomeFantasia.String
	}

	return &usuario.Usuario{
		AggregateRoot: domain.AggregateRoot[string]{
			Entity: domain.Entity[string]{ID: row.ID},
		},
		KeycloakID:   row.KeycloakID,
		Nome:         row.Nome,
		NomeFantasia: nomeFantasia,
		Email:        row.Email,
		Username:     row.Username,
		Credencial: &usuario.Credencial{
			Ativo:           row.Ativo,
			PodeAdministrar: row.PodeAdministrar,
			PodeVotar:       row.PodeVotar,
		},
	}
}

func MapListUsersRowToDomain(row db.ListUsersRow) *usuario.Usuario {
	var nomeFantasia *string
	if row.NomeFantasia.Valid {
		nomeFantasia = &row.NomeFantasia.String
	}

	return &usuario.Usuario{
		AggregateRoot: domain.AggregateRoot[string]{
			Entity: domain.Entity[string]{ID: row.ID},
		},
		KeycloakID:   row.KeycloakID,
		Username:     row.Username,
		Nome:         row.Nome,
		NomeFantasia: nomeFantasia,
		Email:        row.Email,
		Credencial: &usuario.Credencial{
			Ativo:           row.Ativo,
			PodeAdministrar: row.PodeAdministrar,
			PodeVotar:       row.PodeVotar,
		},
	}
}

func MapFindByIDRowToDomain(row db.FindByIDRow) *usuario.Usuario {
	var nomeFantasia *string
	if row.NomeFantasia.Valid {
		nomeFantasia = &row.NomeFantasia.String
	}

	return &usuario.Usuario{
		AggregateRoot: domain.AggregateRoot[string]{
			Entity: domain.Entity[string]{ID: row.ID},
		},
		KeycloakID:   row.KeycloakID,
		Username:     row.Username,
		Nome:         row.Nome,
		NomeFantasia: nomeFantasia,
		Email:        row.Email,
		Credencial: &usuario.Credencial{
			Ativo:           row.Ativo,
			PodeAdministrar: row.PodeAdministrar,
			PodeVotar:       row.PodeVotar,
		},
	}
}
