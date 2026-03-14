package fakes

import (
	"context"
	"time"

	"github.com/aleodoni/voting-go/internal/application/shared"
	"github.com/aleodoni/voting-go/internal/domain/votacao"
)

type FakeReuniaoRepository struct {
	reunioes map[string]*votacao.Reuniao
	projetos map[string][]*votacao.Projeto

	FindReuniaoByIDErr     error
	GetReunioesDiaErr      error
	GetProjetosCompletoErr error
	GetProjetoCompletoErr  error

	FindReuniaoByIDCalls     []string
	GetReunioesDiaCalls      []time.Time
	GetProjetosCompletoCalls []string
	GetProjetoCompletoCalls  []string
}

var _ votacao.ReuniaoRepository = (*FakeReuniaoRepository)(nil)

func NewFakeReuniaoRepository() *FakeReuniaoRepository {
	return &FakeReuniaoRepository{
		reunioes: make(map[string]*votacao.Reuniao),
	}
}

func (f *FakeReuniaoRepository) Seed(r *votacao.Reuniao) {
	f.reunioes[r.ID] = r
}

func (f *FakeReuniaoRepository) SeedProjetos(reuniaoID string, projetos []*votacao.Projeto) {
	if f.projetos == nil {
		f.projetos = make(map[string][]*votacao.Projeto)
	}
	f.projetos[reuniaoID] = projetos
}

func (f *FakeReuniaoRepository) FindReuniaoByID(ctx context.Context, reuniaoID string) (*votacao.Reuniao, error) {
	f.FindReuniaoByIDCalls = append(f.FindReuniaoByIDCalls, reuniaoID)
	if f.FindReuniaoByIDErr != nil {
		return nil, f.FindReuniaoByIDErr
	}
	r, ok := f.reunioes[reuniaoID]
	if !ok {
		return nil, votacao.ErrReuniaoNotFound
	}
	return r, nil
}

func (f *FakeReuniaoRepository) GetReunioesDia(ctx context.Context) ([]*votacao.Reuniao, error) {
	f.GetReunioesDiaCalls = append(f.GetReunioesDiaCalls, shared.GetCurrentDate())
	hoje := shared.GetCurrentDate()
	if f.GetReunioesDiaErr != nil {
		return nil, f.GetReunioesDiaErr
	}
	var reunioes []*votacao.Reuniao
	for _, r := range f.reunioes {
		if r.RecData.Equal(hoje) {
			reunioes = append(reunioes, r)
		}
	}
	return reunioes, nil
}

func (f *FakeReuniaoRepository) GetProjetosCompleto(ctx context.Context, reuniaoID string) ([]*votacao.Projeto, error) {
	f.GetProjetosCompletoCalls = append(f.GetProjetosCompletoCalls, reuniaoID)
	if f.GetProjetosCompletoErr != nil {
		return nil, f.GetProjetosCompletoErr
	}
	projetos, ok := f.projetos[reuniaoID]
	if !ok {
		return []*votacao.Projeto{}, nil
	}
	return projetos, nil
}

func (f *FakeReuniaoRepository) GetProjetoCompleto(ctx context.Context, projetoID string) (*votacao.Projeto, error) {
	f.GetProjetoCompletoCalls = append(f.GetProjetoCompletoCalls, projetoID)
	for _, projetos := range f.projetos {
		for _, p := range projetos {
			if p.ID == projetoID {
				return p, nil
			}
		}
	}
	return nil, votacao.ErrProjetoNotFound
}
