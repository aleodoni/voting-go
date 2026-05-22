package sincronia

import (
	"time"

	"github.com/aleodoni/go-ddd/domain"
)

type Sincronia struct {
	domain.Entity[string]

	sucesso                 *bool
	mensagem_erro           *string
	reunioes_sincronizadas  int
	projetos_sincronizados  int
	pareceres_sincronizados int

	iniciado_em   time.Time
	finalizado_em *time.Time
}
