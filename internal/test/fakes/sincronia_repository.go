package fakes

import (
	"context"

	"github.com/aleodoni/voting-go/internal/domain/sincronia"
)

// FakeSincroniaRepository é uma implementação fake de SincroniaRepository para uso em testes.
type FakeSincroniaRepository struct {
	sincronias []*sincronia.Sincronia

	// Erros configuráveis por método
	SyncErr     error
	ListLastErr error

	// Chamadas registradas para asserção
	SyncCalls     int
	ListLastCalls int
}

// Verificação em tempo de compilação.
var _ sincronia.SincroniaRepository = (*FakeSincroniaRepository)(nil)

// NewFakeSincroniaRepository cria um novo FakeSincroniaRepository pronto para uso.
func NewFakeSincroniaRepository() *FakeSincroniaRepository {
	return &FakeSincroniaRepository{}
}

// Seed insere sincronias diretamente no fake.
func (f *FakeSincroniaRepository) Seed(s *sincronia.Sincronia) {
	f.sincronias = append(f.sincronias, s)
}

// Sync simula a execução de uma sincronização.
func (f *FakeSincroniaRepository) Sync(ctx context.Context) (*sincronia.Sincronia, error) {
	f.SyncCalls++

	if f.SyncErr != nil {
		return nil, f.SyncErr
	}

	s := &sincronia.Sincronia{}
	f.sincronias = append(f.sincronias, s)
	return s, nil
}

// ListLastSincronias retorna as sincronias armazenadas ou o erro configurado.
func (f *FakeSincroniaRepository) ListLastSincronias(ctx context.Context) ([]*sincronia.Sincronia, error) {
	f.ListLastCalls++

	if f.ListLastErr != nil {
		return nil, f.ListLastErr
	}

	return f.sincronias, nil
}
