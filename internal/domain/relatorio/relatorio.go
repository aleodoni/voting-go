package relatorio

import "context"

type VotoItem struct {
	UsuarioNome   string
	Opcao         string
	Restricao     string
	VotoContrario string
}

type VotacaoItem struct {
	Totais map[string]int
	Votos  []VotoItem
}

type ProjetoItem struct {
	CodigoProposicao string
	Iniciativa       string
	Relator          string
	Votacao          *VotacaoItem
}

type ReuniaoOutput struct {
	ConDesc        string
	RecNumero      string
	RecTipoReuniao string
	RecData        string
	Projetos       []ProjetoItem
}

type Generator interface {
	Gera(ctx context.Context, data ReuniaoOutput) ([]byte, error)
}
