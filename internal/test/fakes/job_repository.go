package fakes

import (
	"context"

	"github.com/aleodoni/voting-go/internal/domain/job"
)

// FakeJobRepository é uma implementação fake de JobRepository para uso em testes.
type FakeJobRepository struct {
	// Erros configuráveis por método
	FecharVotacoesAbertasErr error

	// Chamadas registradas para asserção
	FecharVotacoesAbertasCalls int
}

// Verificação em tempo de compilação.
var _ job.JobRepository = (*FakeJobRepository)(nil)

// NewFakeJobRepository cria um novo FakeJobRepository pronto para uso.
func NewFakeJobRepository() *FakeJobRepository {
	return &FakeJobRepository{}
}

// FecharVotacoesAbertas simula o fechamento de votações abertas.
func (f *FakeJobRepository) FecharVotacoesAbertas(ctx context.Context) error {
	f.FecharVotacoesAbertasCalls++

	if f.FecharVotacoesAbertasErr != nil {
		return f.FecharVotacoesAbertasErr
	}

	return nil
}
