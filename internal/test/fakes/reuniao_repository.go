package fakes

import (
	"context"
	"time"

	"github.com/aleodoni/voting-go/internal/application/shared"
	"github.com/aleodoni/voting-go/internal/domain/votacao"
)

type FakeReuniaoRepository struct {
	// Dados armazenados internamente (simulam o banco)
	reunioes map[string]*votacao.Reuniao // chave: reuniaoID

	// Erros configuráveis por método
	FindReuniaoByIDErr error
	GetReunioesDiaErr  error

	// Chamadas registradas para asserção nos testes
	FindReuniaoByIDCalls []string
	GetReunioesDiaCalls  []time.Time
}

// Verificação em tempo de compilação: garante que FakeReuniaoRepository implementa ReuniaoRepository.
var _ votacao.ReuniaoRepository = (*FakeReuniaoRepository)(nil)

// NewFakeReuniaoRepository cria um novo FakeReuniaoRepository pronto para uso.
func NewFakeReuniaoRepository() *FakeReuniaoRepository {
	return &FakeReuniaoRepository{
		reunioes: make(map[string]*votacao.Reuniao),
	}
}

// Seed insere usuários diretamente no fake (útil para preparar cenários de teste).
func (f *FakeReuniaoRepository) Seed(r *votacao.Reuniao) {
	f.reunioes[r.ID] = r
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
