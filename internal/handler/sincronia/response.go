package sincronia

import "time"

// ExecutaSincroniaResponse representa os dados de uma resposta de execução de sincronização
//
//	@name	ExecutaSincroniaResponse
type SincroniaResponse struct {
	ID                     string     `json:"id"             example:"cls1abc123"`
	IniciadoEm             time.Time  `json:"iniciado_em"         example:"2026-03-18T00:00:00Z"`
	FinalizadoEm           *time.Time `json:"finalizado_em,omitempty"         example:"2026-03-18T00:00:00Z"`
	Sucesso                *bool      `json:"sucesso,omitempty"       example:"true"`
	MensagemErro           *string    `json:"mensagem_erro,omitempty" example:""`
	ReunioesSincronizadas  int        `json:"reunioes_sincronizadas" example:"42"`
	ProjetosSincronizados  int        `json:"projetos_sincronizados" example:"42"`
	PareceresSincronizados int        `json:"pareceres_sincronizados" example:"42"`
}

// ListUltimasSincroniasResponse representa os dados de uma listagem de últimas sincronias retornadas pela API
//
//	@name	ListUltimasSincroniasResponse
type ListUltimasSincroniasResponse struct {
	Sincronias []*SincroniaResponse `json:"sincronias"`
}

// ErrorResponse representa os dados de um erro retornado pela API
//
//	@name	ErrorResponse
type ErrorResponse struct {
	Error string `json:"error" example:"mensagem de erro"`
}
