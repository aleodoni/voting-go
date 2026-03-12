package votacao

import "time"

type Projeto struct {
	ID                string
	Sumula            string
	Relator           string
	TemEmendas        bool
	PacID             int
	ParID             int
	CodigoProposicao  string
	Iniciativa        string
	ConclusaoComissao string
	ConclusaoRelator  string
	ReuniaoID         string

	Pareceres *[]Parecer
	Votacoes  *[]Votacao

	CreatedAt time.Time
	UpdatedAt time.Time
}
