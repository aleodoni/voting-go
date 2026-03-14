package fakes

import (
	"context"

	"github.com/aleodoni/voting-go/internal/domain/votacao"
)

type FakeVotacaoRepository struct {
	votacoes map[string]*votacao.Votacao

	SalvaVotacaoErr  error
	DeletaVotacaoErr error

	SalvaVotacaoCalls  []votacao.Votacao
	DeletaVotacaoCalls []string
}

var _ votacao.VotacaoRepository = (*FakeVotacaoRepository)(nil)

func NewFakeVotacaoRepository() *FakeVotacaoRepository {
	return &FakeVotacaoRepository{
		votacoes: make(map[string]*votacao.Votacao),
	}
}

func (f *FakeVotacaoRepository) SalvaVotacao(ctx context.Context, v *votacao.Votacao) error {
	if f.SalvaVotacaoErr != nil {
		return f.SalvaVotacaoErr
	}
	f.SalvaVotacaoCalls = append(f.SalvaVotacaoCalls, *v)
	f.votacoes[v.ID] = v
	return nil
}

func (f *FakeVotacaoRepository) DeletaVotacao(ctx context.Context, votacaoID string) error {
	if f.DeletaVotacaoErr != nil {
		return f.DeletaVotacaoErr
	}
	f.DeletaVotacaoCalls = append(f.DeletaVotacaoCalls, votacaoID)
	delete(f.votacoes, votacaoID)
	return nil
}
