package votacao

import (
	"time"

	domainUsuario "github.com/aleodoni/voting-go/internal/domain/usuario"
)

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

	Usuario       domainUsuario.Usuario
	Restricao     *Restricao
	VotoContrario *VotoContrario
}
