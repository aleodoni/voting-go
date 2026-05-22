package mappers

import (
	"time"

	"github.com/aleodoni/go-ddd/domain"

	"github.com/aleodoni/voting-go/internal/domain/sincronia"
	db "github.com/aleodoni/voting-go/internal/infrastructure/persistence/sqlc/generated"
)

func MapGetLastSincroniaRowToDomain(row db.Sincronium) *sincronia.Sincronia {
	var finalizadoEm *time.Time
	if row.FinalizadoEm.Valid {
		finalizadoEm = &row.FinalizadoEm.Time
	}

	var sucesso *bool
	if row.Sucesso.Valid {
		sucesso = &row.Sucesso.Bool
	}

	var mensagemErro *string
	if row.MensagemErro.Valid {
		mensagemErro = &row.MensagemErro.String
	}

	reunioesSincronizadas := 0

	if row.ReunioesSincronizadas.Valid {
		reunioesSincronizadas = int(row.ReunioesSincronizadas.Int32)
	}

	projetosSincronizados := 0
	if row.ProjetosSincronizados.Valid {
		projetosSincronizados = int(row.ProjetosSincronizados.Int32)
	}

	pareceresSincronizados := 0
	if row.PareceresSincronizados.Valid {
		pareceresSincronizados = int(row.PareceresSincronizados.Int32)
	}

	return &sincronia.Sincronia{
		Entity: domain.Entity[string]{
			ID: row.ID,
		},
		IniciadoEm:             row.IniciadoEm.Time,
		FinalizadoEm:           finalizadoEm,
		Sucesso:                sucesso,
		MensagemErro:           mensagemErro,
		ReunioesSincronizadas:  reunioesSincronizadas,
		ProjetosSincronizados:  projetosSincronizados,
		PareceresSincronizados: pareceresSincronizados,
	}
}

func MapSincroniumToDomain(row db.Sincronium) *sincronia.Sincronia {
	var finalizadoEm *time.Time
	if row.FinalizadoEm.Valid {
		finalizadoEm = &row.FinalizadoEm.Time
	}

	var sucesso *bool
	if row.Sucesso.Valid {
		sucesso = &row.Sucesso.Bool
	}

	var mensagemErro *string
	if row.MensagemErro.Valid {
		mensagemErro = &row.MensagemErro.String
	}

	reunioesSincronizadas := 0
	if row.ReunioesSincronizadas.Valid {
		reunioesSincronizadas = int(row.ReunioesSincronizadas.Int32)
	}

	projetosSincronizados := 0
	if row.ProjetosSincronizados.Valid {
		projetosSincronizados = int(row.ProjetosSincronizados.Int32)
	}

	pareceresSincronizados := 0
	if row.PareceresSincronizados.Valid {
		pareceresSincronizados = int(row.PareceresSincronizados.Int32)
	}

	return &sincronia.Sincronia{
		Entity: domain.Entity[string]{
			ID: row.ID,
		},
		IniciadoEm:             row.IniciadoEm.Time,
		FinalizadoEm:           finalizadoEm,
		Sucesso:                sucesso,
		MensagemErro:           mensagemErro,
		ReunioesSincronizadas:  reunioesSincronizadas,
		ProjetosSincronizados:  projetosSincronizados,
		PareceresSincronizados: pareceresSincronizados,
	}
}
