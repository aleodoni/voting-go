package votacao

import "time"

type OpcaoVoto string

const (
	OpcaoVotoF OpcaoVoto = "F"
	OpcaoVotoR OpcaoVoto = "R"
	OpcaoVotoC OpcaoVoto = "C"
	OpcaoVotoV OpcaoVoto = "V"
	OpcaoVotoA OpcaoVoto = "A"
)

type Voto struct {
	ID        string
	Voto      OpcaoVoto
	VotacaoID string
	UsuarioID string
	CreatedAt time.Time
	UpdatedAt time.Time

	Restricao     *Restricao
	VotoContrario *VotoContrario
}
