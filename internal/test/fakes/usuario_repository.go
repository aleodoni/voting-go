// Package fakes contém implementações fake de interfaces para uso em testes.
package fakes

import (
	"context"
	"strings"

	"github.com/aleodoni/voting-go/internal/domain/usuario"
)

// FakeUsuarioRepository é uma implementação fake de UsuarioRepository para uso em testes.
// Todos os campos de controle são públicos para permitir configuração direta nos testes.
type FakeUsuarioRepository struct {
	// Dados armazenados internamente (simulam o banco)
	usuarios map[string]*usuario.Usuario // chave: keycloakID

	// Erros configuráveis por método
	FindByKeycloakIDErr             error
	FindByUsernameErr               error
	CreateErr                       error
	UpdateDisplayNamePermissionsErr error
	ListUsersErr                    error

	// Chamadas registradas para asserção nos testes
	FindByKeycloakIDCalls             []string
	FindByUsernameCalls               []string
	CreateCalls                       []*usuario.Usuario
	UpdateDisplayNamePermissionsCalls []UpdateDisplayNamePermissionsArgs
	ListUsersCalls                    []ListUsersArgs
}

// UpdateDisplayNamePermissionsArgs registra os argumentos de cada chamada ao método.
type UpdateDisplayNamePermissionsArgs struct {
	UserID      string
	DisplayName *string
	IsActive    bool
	CanAdmin    bool
	CanVote     bool
}

// ListUsersArgs registra os argumentos de cada chamada ao método.
type ListUsersArgs struct {
	Search string
	Page   int
	Limit  int
}

// Verificação em tempo de compilação: garante que FakeUsuarioRepository implementa UsuarioRepository.
var _ usuario.UsuarioRepository = (*FakeUsuarioRepository)(nil)

// NewFakeUsuarioRepository cria um novo FakeUsuarioRepository pronto para uso.
func NewFakeUsuarioRepository() *FakeUsuarioRepository {
	return &FakeUsuarioRepository{
		usuarios: make(map[string]*usuario.Usuario),
	}
}

// Seed insere usuários diretamente no fake (útil para preparar cenários de teste).
func (f *FakeUsuarioRepository) Seed(u *usuario.Usuario) {
	f.usuarios[u.KeycloakID] = u
}

// FindByKeycloakID retorna o usuário correspondente ao keycloakID ou o erro configurado.
func (f *FakeUsuarioRepository) FindByKeycloakID(ctx context.Context, keycloakID string) (*usuario.Usuario, error) {
	f.FindByKeycloakIDCalls = append(f.FindByKeycloakIDCalls, keycloakID)

	if f.FindByKeycloakIDErr != nil {
		return nil, f.FindByKeycloakIDErr
	}

	u, ok := f.usuarios[keycloakID]
	if !ok {
		return nil, usuario.ErrUserNotFound
	}

	return u, nil
}

// FindByUsername retorna o primeiro usuário cujo Username corresponda ou o erro configurado.
func (f *FakeUsuarioRepository) FindByUsername(ctx context.Context, username string) (*usuario.Usuario, error) {
	f.FindByUsernameCalls = append(f.FindByUsernameCalls, username)

	if f.FindByUsernameErr != nil {
		return nil, f.FindByUsernameErr
	}

	for _, u := range f.usuarios {
		if u.Username == username {
			return u, nil
		}
	}

	return nil, usuario.ErrUserNotFound
}

// Create armazena o usuário internamente ou retorna o erro configurado.
func (f *FakeUsuarioRepository) Create(ctx context.Context, u *usuario.Usuario) error {
	f.CreateCalls = append(f.CreateCalls, u)

	if f.CreateErr != nil {
		return f.CreateErr
	}

	f.usuarios[u.KeycloakID] = u
	return nil
}

// UpdateDisplayNamePermissions atualiza o usuário armazenado ou retorna o erro configurado.
func (f *FakeUsuarioRepository) UpdateDisplayNamePermissions(
	ctx context.Context,
	userID string,
	displayName *string,
	isActive bool,
	canAdmin bool,
	canVote bool,
) error {
	f.UpdateDisplayNamePermissionsCalls = append(f.UpdateDisplayNamePermissionsCalls, UpdateDisplayNamePermissionsArgs{
		UserID:      userID,
		DisplayName: displayName,
		IsActive:    isActive,
		CanAdmin:    canAdmin,
		CanVote:     canVote,
	})

	if f.UpdateDisplayNamePermissionsErr != nil {
		return f.UpdateDisplayNamePermissionsErr
	}

	for _, u := range f.usuarios {
		if u.ID == userID {
			if displayName != nil {
				u.NomeFantasia = displayName
			}

			if u.Credencial != nil {
				u.Credencial.Ativo = isActive
				u.Credencial.PodeAdministrar = canAdmin
				u.Credencial.PodeVotar = canVote
			}

			return nil
		}
	}

	return usuario.ErrUserNotFound
}

// ListUsers retorna os usuários filtrados pelo search ou o erro configurado.
func (f *FakeUsuarioRepository) ListUsers(ctx context.Context, search string, page, limit int) ([]*usuario.Usuario, int64, error) {
	f.ListUsersCalls = append(f.ListUsersCalls, ListUsersArgs{
		Search: search,
		Page:   page,
		Limit:  limit,
	})

	if f.ListUsersErr != nil {
		return nil, 0, f.ListUsersErr
	}

	var result []*usuario.Usuario
	for _, u := range f.usuarios {
		if search == "" || strings.Contains(u.Username, search) {
			result = append(result, u)
		}
	}

	return result, int64(len(result)), nil
}
