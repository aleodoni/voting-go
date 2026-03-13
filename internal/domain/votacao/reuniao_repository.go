package votacao

import (
	"context"
)

type ReuniaoRepository interface {
	FindReuniaoByID(ctx context.Context, reuniaoID string) (*Reuniao, error)
	GetReunioesDia(ctx context.Context) ([]*Reuniao, error)
	GetProjetosCompleto(ctx context.Context, projetoID string) ([]*Projeto, error)
	GetProjetoCompleto(ctx context.Context, projetoID string) (*Projeto, error)
	CriaVotacao(ctx context.Context, votacao *Votacao) error
}
