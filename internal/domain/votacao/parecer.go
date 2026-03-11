package votacao

import "time"

type Parecer struct {
	ID               string
	CodigoProposicao string
	TCPNome          string
	Vereador         string
	IDTexto          int
	ProjetoID        string

	CreatedAt time.Time
	UpdatedAt time.Time
}
