package reuniao

import "time"

// ParecerResponse representa os dados de uma reunião retornados pela API
//
//	@name	ParecerResponse
type ParecerResponse struct {
	ID               string    `json:"id"                example:"cls1par123"`
	CodigoProposicao string    `json:"codigo_proposicao" example:"PL-001-2026"`
	TCPNome          string    `json:"tcp_nome"          example:"Favorável"`
	Vereador         string    `json:"vereador"          example:"João Silva"`
	IDTexto          int       `json:"id_texto"          example:"1"`
	ProjetoID        string    `json:"projeto_id"        example:"cls1abc123"`
	CreatedAt        time.Time `json:"created_at"        example:"2026-03-18T09:00:00Z"`
	UpdatedAt        time.Time `json:"updated_at"        example:"2026-03-18T09:00:00Z"`
}

// VotacaoResponse representa os dados de uma votação retornados pela API
//
//	@name	VotacaoResponse
type VotacaoResponse struct {
	ID        string         `json:"id"         example:"cls1vot123"`
	ProjetoID *string        `json:"projeto_id" example:"cls1abc123"`
	Status    string         `json:"status"     example:"A"`
	Votos     []VotoResponse `json:"votos"`
	CreatedAt time.Time      `json:"created_at" example:"2026-03-18T09:00:00Z"`
	UpdatedAt time.Time      `json:"updated_at" example:"2026-03-18T09:00:00Z"`
}

// ReuniaoResponse representa os dados de uma reunião retornados pela API
//
//	@name	ReuniaoResponse
type ReuniaoResponse struct {
	ID             string             `json:"id"               example:"cls1abc123"`
	ConID          int                `json:"con_id"           example:"1"`
	RecID          int                `json:"rec_id"           example:"42"`
	PacID          int                `json:"pac_id"           example:"7"`
	ConDesc        string             `json:"con_desc"         example:"Câmara Municipal"`
	ConSigla       string             `json:"con_sigla"        example:"CM"`
	RecTipoReuniao string             `json:"rec_tipo_reuniao" example:"Ordinária"`
	RecNumero      string             `json:"rec_numero"       example:"001/2026"`
	RecData        time.Time          `json:"rec_data"         example:"2026-03-18T00:00:00Z"`
	Projetos       *[]ProjetoResponse `json:"projetos"`
	CreatedAt      time.Time          `json:"created_at"       example:"2026-03-18T09:00:00Z"`
	UpdatedAt      time.Time          `json:"updated_at"       example:"2026-03-18T09:00:00Z"`
}

// ProjetoResponse representa os dados de um projeto retornados pela API
//
//	@name	ProjetoResponse
type ProjetoResponse struct {
	ID                string             `json:"id"                  example:"cls1abc123"`
	Sumula            string             `json:"sumula"              example:"Projeto de Lei nº 001/2026"`
	Relator           string             `json:"relator"             example:"Vereador João"`
	TemEmendas        bool               `json:"tem_emendas"         example:"false"`
	PacID             int                `json:"pac_id"              example:"7"`
	ParID             int                `json:"par_id"              example:"3"`
	CodigoProposicao  string             `json:"codigo_proposicao"   example:"PL-001-2026"`
	Iniciativa        string             `json:"iniciativa"          example:"Poder Executivo"`
	ConclusaoComissao string             `json:"conclusao_comissao"  example:"Favorável"`
	ConclusaoRelator  string             `json:"conclusao_relator"   example:"Favorável"`
	ReuniaoID         string             `json:"reuniao_id"          example:"cls1xyz456"`
	Pareceres         *[]ParecerResponse `json:"pareceres"`
	Votacao           *VotacaoResponse   `json:"votacao"`
	CreatedAt         time.Time          `json:"created_at"          example:"2026-03-18T09:00:00Z"`
	UpdatedAt         time.Time          `json:"updated_at"          example:"2026-03-18T09:00:00Z"`
}

type UsuarioVotoResponse struct {
	ID           string  `json:"id"`
	Nome         string  `json:"nome"`
	NomeFantasia *string `json:"nome_fantasia"`
}

type RestricaoResponse struct {
	ID        string `json:"id"`
	Restricao string `json:"restricao"`
}

type VotoContrarioResponse struct {
	ID        string           `json:"id"`
	IDTexto   int              `json:"id_texto"`
	ParecerID string           `json:"parecer_id"`
	Parecer   *ParecerResponse `json:"parecer"`
}

type VotoResponse struct {
	ID            string                 `json:"id"`
	Voto          string                 `json:"voto"`
	UsuarioID     string                 `json:"usuario_id"`
	Usuario       UsuarioVotoResponse    `json:"usuario"`
	Restricao     *RestricaoResponse     `json:"restricao"`
	VotoContrario *VotoContrarioResponse `json:"voto_contrario"`
}

// ErrorResponse representa os dados de um erro retornado pela API
//
//	@name	ErrorResponse
type ErrorResponse struct {
	Error string `json:"error" example:"mensagem de erro"`
}
