package httperr_test

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/gin-gonic/gin"

	domainUsuario "github.com/aleodoni/voting-go/internal/domain/usuario"
	domainVotacao "github.com/aleodoni/voting-go/internal/domain/votacao"
	"github.com/aleodoni/voting-go/internal/handler/httperr"
)

func responder(t *testing.T, err error) (int, string) {
	t.Helper()

	gin.SetMode(gin.TestMode)

	rec := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(rec)
	c.Request = httptest.NewRequest(http.MethodGet, "/teste", nil)

	httperr.Respond(c, err)

	var corpo map[string]string
	if erro := json.Unmarshal(rec.Body.Bytes(), &corpo); erro != nil {
		t.Fatalf("corpo não é um JSON {\"error\": ...}: %v (%s)", erro, rec.Body.String())
	}

	return rec.Code, corpo["error"]
}

func TestRespond_ErrosConhecidos(t *testing.T) {
	casos := []struct {
		nome     string
		err      error
		status   int
		mensagem string
	}{
		{"usuário não encontrado", domainUsuario.ErrUserNotFound, http.StatusNotFound, "usuário não encontrado"},
		{"credencial não encontrada", domainUsuario.ErrCredencialNotFound, http.StatusNotFound, "credencial não encontrada"},
		{"reunião não encontrada", domainVotacao.ErrReuniaoNotFound, http.StatusNotFound, "reunião não encontrada"},
		{"projeto não encontrado", domainVotacao.ErrProjetoNotFound, http.StatusNotFound, "projeto não encontrado"},
		{"usuário não é admin", domainUsuario.ErrUserNotAdmin, http.StatusForbidden, "usuário não tem permissões de admin"},
		{"usuário não é votante", domainUsuario.ErrUserNotVoter, http.StatusForbidden, "usuário não tem permissões de votante"},
		{"usuário inativo", domainUsuario.ErrUserNotActive, http.StatusForbidden, "usuário não está ativo"},
	}

	for _, tc := range casos {
		t.Run(tc.nome, func(t *testing.T) {
			status, mensagem := responder(t, tc.err)

			if status != tc.status {
				t.Fatalf("esperava status %d, obteve %d", tc.status, status)
			}

			if mensagem != tc.mensagem {
				t.Fatalf("esperava mensagem %q, obteve %q", tc.mensagem, mensagem)
			}
		})
	}
}

// Erros embrulhados com %w são reconhecidos, mas a mensagem devolvida é a do
// erro conhecido, sem o contexto interno que o embrulha.
func TestRespond_ErroEmbrulhado(t *testing.T) {
	err := fmt.Errorf("falha ao carregar o usuário logado: %w", domainUsuario.ErrUserNotAdmin)

	status, mensagem := responder(t, err)

	if status != http.StatusForbidden {
		t.Fatalf("esperava status %d, obteve %d", http.StatusForbidden, status)
	}

	if mensagem != "usuário não tem permissões de admin" {
		t.Fatalf("mensagem inesperada: %q", mensagem)
	}
}

// Erros desconhecidos viram 500 genérico e não vazam o texto original.
func TestRespond_ErroDesconhecidoNaoVazaDetalhes(t *testing.T) {
	err := errors.New("pq: connection refused (host=10.0.0.5)")

	status, mensagem := responder(t, err)

	if status != http.StatusInternalServerError {
		t.Fatalf("esperava status %d, obteve %d", http.StatusInternalServerError, status)
	}

	if mensagem != "erro interno do servidor" {
		t.Fatalf("esperava a mensagem genérica, obteve %q", mensagem)
	}

	if strings.Contains(mensagem, "connection refused") || strings.Contains(mensagem, "10.0.0.5") {
		t.Fatalf("a resposta vazou detalhes do erro: %q", mensagem)
	}
}
