package votacao

import "time"

type StatusVotacao string

const (
	StatusVotacaoA StatusVotacao = "A"
	StatusVotacaoF StatusVotacao = "F"
	StatusVotacaoV StatusVotacao = "V"
	StatusVotacaoC StatusVotacao = "C"
)

type Votacao struct {
	ID        string
	ProjetoID *string
	Status    StatusVotacao
	CreatedAt time.Time
	UpdatedAt time.Time

	Projeto *Projeto
	Votos   *[]Voto
}
