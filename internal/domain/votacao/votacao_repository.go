package votacao

import "context"

type VotacaoRepository interface {
	SalvaVotacao(ctx context.Context, votacao *Votacao) error
	DeletaVotacao(ctx context.Context, votacaoID string) error
	SalvaVoto(ctx context.Context, voto *Voto) error
	BuscaVotacao(ctx context.Context, votacaoID string) (*Votacao, error)
	UsuarioJaVotou(ctx context.Context, usuarioID, votacaoID string) (bool, error)
	GetVotacaoAberta(ctx context.Context) (*Votacao, error)
}
