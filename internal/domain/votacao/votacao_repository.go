package votacao

import "context"

type VotacaoRepository interface {
	SalvaVotacao(ctx context.Context, votacao *Votacao) error
	DeletaVotacao(ctx context.Context, votacaoID string) error
}
