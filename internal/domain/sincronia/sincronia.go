package sincronia

import (
	"time"

	"github.com/aleodoni/go-ddd/domain"
)

type Sincronia struct {
	domain.Entity[string]

	Sucesso                *bool
	MensagemErro           *string
	ReunioesSincronizadas  int
	ProjetosSincronizados  int
	PareceresSincronizados int

	IniciadoEm   time.Time
	FinalizadoEm *time.Time
}
