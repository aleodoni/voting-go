// Package votacao defines the Votacao entity and its related methods.
package votacao

import "time"

type Reuniao struct {
	ID             string
	ConID          int
	RecID          int
	PacID          int
	ConDesc        string
	ConSigla       string
	RecTipoReuniao string
	RecNumero      string
	RecData        time.Time

	Projetos *[]Projeto

	CreatedAt time.Time
	UpdatedAt time.Time
}
