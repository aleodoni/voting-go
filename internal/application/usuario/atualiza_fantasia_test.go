package usuario_test

import (
	"context"
	"testing"

	"github.com/aleodoni/go-ddd/domain"
	ucUsuario "github.com/aleodoni/voting-go/internal/application/usuario"
	domainUsuario "github.com/aleodoni/voting-go/internal/domain/usuario"
	"github.com/aleodoni/voting-go/internal/test/fakes"
)

func strPtr(s string) *string {
	return &s
}

func makeUsuario(id, keycloakID string) *domainUsuario.Usuario {
	u := &domainUsuario.Usuario{
		AggregateRoot: domain.NewAggregateRoot(id),
		KeycloakID:    keycloakID,
		Username:      "usuario_teste",
		Email:         "teste@teste.com",
		Nome:          "Usuário Teste",
	}

	return u
}

func TestUpdateDisplayNameUseCase_Execute(t *testing.T) {
	const (
		userID     = "user-123"
		keycloakID = "keycloak-456"
	)

	t.Run("atualiza nome fantasia com sucesso", func(t *testing.T) {
		repo := fakes.NewFakeUsuarioRepository()
		repo.Seed(makeUsuario(userID, keycloakID))

		uc := ucUsuario.NewUpdateDisplayNameUseCase(repo)

		err := uc.Execute(context.Background(), ucUsuario.UpdateDisplayNameInput{
			LoggedInUserKeycloakID: keycloakID,
			UserID:                 userID,
			DisplayName:            strPtr("Vereador Teste"),
		})

		if err != nil {
			t.Fatalf("esperava nil, got %v", err)
		}

		if len(repo.UpdateDisplayNameCalls) != 1 {
			t.Fatalf("esperava 1 chamada a UpdateDisplayName, got %d", len(repo.UpdateDisplayNameCalls))
		}

		call := repo.UpdateDisplayNameCalls[0]
		if call.UserID != userID {
			t.Errorf("esperava userID %q, got %q", userID, call.UserID)
		}
		if *call.DisplayName != "Vereador Teste" {
			t.Errorf("esperava displayName %q, got %q", "Vereador Teste", *call.DisplayName)
		}
	})

	t.Run("retorna erro quando usuário não existe", func(t *testing.T) {
		repo := fakes.NewFakeUsuarioRepository()
		// repositório vazio — nenhum usuário seedado

		uc := ucUsuario.NewUpdateDisplayNameUseCase(repo)

		err := uc.Execute(context.Background(), ucUsuario.UpdateDisplayNameInput{
			LoggedInUserKeycloakID: keycloakID,
			UserID:                 userID,
			DisplayName:            strPtr("Qualquer Nome"),
		})

		if err != domainUsuario.ErrUserNotFound {
			t.Errorf("esperava ErrUserNotFound, got %v", err)
		}

		if len(repo.UpdateDisplayNameCalls) != 0 {
			t.Errorf("não deveria chamar UpdateDisplayName, mas foi chamado %d vez(es)", len(repo.UpdateDisplayNameCalls))
		}
	})

	t.Run("retorna erro quando keycloakID não corresponde ao usuário logado", func(t *testing.T) {
		repo := fakes.NewFakeUsuarioRepository()
		repo.Seed(makeUsuario(userID, keycloakID))

		uc := ucUsuario.NewUpdateDisplayNameUseCase(repo)

		err := uc.Execute(context.Background(), ucUsuario.UpdateDisplayNameInput{
			LoggedInUserKeycloakID: "keycloak-outro",
			UserID:                 userID,
			DisplayName:            strPtr("Tentativa Inválida"),
		})

		if err != domainUsuario.ErrUserNotFound {
			t.Errorf("esperava ErrUserNotFound, got %v", err)
		}

		if len(repo.UpdateDisplayNameCalls) != 0 {
			t.Errorf("não deveria chamar UpdateDisplayName, mas foi chamado %d vez(es)", len(repo.UpdateDisplayNameCalls))
		}
	})

	t.Run("repassa erro do repositório ao atualizar", func(t *testing.T) {
		repo := fakes.NewFakeUsuarioRepository()
		repo.Seed(makeUsuario(userID, keycloakID))
		repo.UpdateDisplayNameErr = domainUsuario.ErrUserNotFound

		uc := ucUsuario.NewUpdateDisplayNameUseCase(repo)

		err := uc.Execute(context.Background(), ucUsuario.UpdateDisplayNameInput{
			LoggedInUserKeycloakID: keycloakID,
			UserID:                 userID,
			DisplayName:            strPtr("Nome Válido"),
		})

		if err == nil {
			t.Fatal("esperava erro, got nil")
		}
	})
}
