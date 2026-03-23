package fakes

import (
	"context"
	"time"

	"github.com/aleodoni/voting-go/internal/domain/votacao"
)

type FakeVotacaoRepository struct {
	votacoes map[string]*votacao.Votacao
	votos    map[string][]string // votacaoID -> []usuarioID  ← novo

	SalvaVotacaoErr            error
	DeletaVotacaoErr           error
	SalvaVotoErr               error
	BuscaVotacaoErr            error
	UsuarioJaVotouErr          error
	GetVotacaoAbertaErr        error
	GetProjetoVotacaoAbertaErr error
	GetVotingStatsErr          error

	SalvaVotacaoCalls    []votacao.Votacao
	DeletaVotacaoCalls   []string
	SalvaVotoCalls       []votacao.Voto
	projetoVotacaoAberta *votacao.Projeto
	votingStats          *votacao.VotingStats
}

var _ votacao.VotacaoRepository = (*FakeVotacaoRepository)(nil)

func NewFakeVotacaoRepository() *FakeVotacaoRepository {
	return &FakeVotacaoRepository{
		votacoes: make(map[string]*votacao.Votacao),
		votos:    make(map[string][]string),
	}
}

func (f *FakeVotacaoRepository) GetProjetoVotacaoAberta(ctx context.Context) (*votacao.Projeto, error) {
	if f.GetProjetoVotacaoAbertaErr != nil {
		return nil, f.GetProjetoVotacaoAbertaErr
	}
	return f.projetoVotacaoAberta, nil
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

func (f *FakeVotacaoRepository) SalvaVoto(ctx context.Context, v *votacao.Voto) error {
	if f.SalvaVotoErr != nil {
		return f.SalvaVotoErr
	}
	f.SalvaVotoCalls = append(f.SalvaVotoCalls, *v)
	// Registra o voto no mapa para UsuarioJaVotou funcionar nos testes
	f.votos[v.VotacaoID] = append(f.votos[v.VotacaoID], v.UsuarioID)
	return nil
}

func (f *FakeVotacaoRepository) BuscaVotacao(ctx context.Context, votacaoID string) (*votacao.Votacao, error) {
	if f.BuscaVotacaoErr != nil {
		return nil, f.BuscaVotacaoErr
	}
	v, ok := f.votacoes[votacaoID]
	if !ok {
		return nil, votacao.ErrVotacaoNaoEncontrada
	}
	return v, nil
}

func (f *FakeVotacaoRepository) UsuarioJaVotou(ctx context.Context, usuarioID, votacaoID string) (bool, error) {
	if f.UsuarioJaVotouErr != nil {
		return false, f.UsuarioJaVotouErr
	}
	for _, uid := range f.votos[votacaoID] {
		if uid == usuarioID {
			return true, nil
		}
	}
	return false, nil
}

func (f *FakeVotacaoRepository) GetVotacaoAberta(ctx context.Context) (*votacao.Votacao, error) {
	if f.GetVotacaoAbertaErr != nil {
		return nil, f.GetVotacaoAbertaErr
	}
	for _, v := range f.votacoes {
		if v.Status == votacao.StatusVotacaoA {
			return v, nil
		}
	}
	return nil, nil
}

func (f *FakeVotacaoRepository) GetVotingStats(ctx context.Context, date time.Time) (*votacao.VotingStats, error) {
	if f.GetVotingStatsErr != nil {
		return nil, f.GetVotingStatsErr
	}
	return f.votingStats, nil
}

func (f *FakeVotacaoRepository) SeedVotingStats(s *votacao.VotingStats) {
	f.votingStats = s
}

func (f *FakeVotacaoRepository) Seed(v *votacao.Votacao) {
	f.votacoes[v.ID] = v
}

func (f *FakeVotacaoRepository) SeedProjeto(p *votacao.Projeto) {
	f.projetoVotacaoAberta = p
}
