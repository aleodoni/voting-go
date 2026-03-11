// Package fakes contém implementações fake de interfaces para uso em testes.
package fakes

import (
	"context"

	"github.com/aleodoni/voting-go/internal/domain/usuario"
)

// FakeCredencialRepository é uma implementação fake de CredencialRepository para uso em testes.
// Todos os campos de controle são públicos para permitir configuração direta nos testes.
type FakeCredencialRepository struct {
	// Dados armazenados internamente (simulam o banco)
	credenciais map[string]*usuario.Credencial // chave: usuarioID

	// Erros configuráveis por método
	FindByUsuarioIDErr error
	CreateErr          error
	UpdateErr          error

	// Chamadas registradas para asserção nos testes
	FindByUsuarioIDCalls []string
	CreateCalls          []*usuario.Credencial
	UpdateCalls          []*usuario.Credencial
}

// Verificação em tempo de compilação: garante que FakeCredencialRepository implementa CredencialRepository.
var _ usuario.CredencialRepository = (*FakeCredencialRepository)(nil)

// NewFakeCredencialRepository cria um novo FakeCredencialRepository pronto para uso.
func NewFakeCredencialRepository() *FakeCredencialRepository {
	return &FakeCredencialRepository{
		credenciais: make(map[string]*usuario.Credencial),
	}
}

// Seed insere credenciais diretamente no fake (útil para preparar cenários de teste).
func (f *FakeCredencialRepository) Seed(c *usuario.Credencial) {
	f.credenciais[c.UsuarioID] = c
}

// FindByUsuarioID retorna a credencial correspondente ao usuarioID ou o erro configurado.
func (f *FakeCredencialRepository) FindByUsuarioID(ctx context.Context, usuarioID string) (*usuario.Credencial, error) {
	f.FindByUsuarioIDCalls = append(f.FindByUsuarioIDCalls, usuarioID)

	if f.FindByUsuarioIDErr != nil {
		return nil, f.FindByUsuarioIDErr
	}

	cred, ok := f.credenciais[usuarioID]
	if !ok {
		return nil, usuario.ErrCredencialNotFound
	}

	return cred, nil
}

// Create armazena a credencial internamente ou retorna o erro configurado.
func (f *FakeCredencialRepository) Create(ctx context.Context, cred *usuario.Credencial) error {
	f.CreateCalls = append(f.CreateCalls, cred)

	if f.CreateErr != nil {
		return f.CreateErr
	}

	f.credenciais[cred.UsuarioID] = cred
	return nil
}

// Update atualiza a credencial armazenada ou retorna o erro configurado.
func (f *FakeCredencialRepository) Update(ctx context.Context, cred *usuario.Credencial) error {
	f.UpdateCalls = append(f.UpdateCalls, cred)

	if f.UpdateErr != nil {
		return f.UpdateErr
	}

	f.credenciais[cred.UsuarioID] = cred
	return nil
}
