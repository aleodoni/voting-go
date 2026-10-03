package jobs_test

import (
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/gin-gonic/gin"

	ucJobs "github.com/aleodoni/voting-go/internal/application/jobs"
	handler "github.com/aleodoni/voting-go/internal/handler/jobs"
	"github.com/aleodoni/voting-go/internal/test/fakes"
)

func chamar(t *testing.T, repo *fakes.FakeJobRepository) (int, string) {
	t.Helper()

	gin.SetMode(gin.TestMode)

	h := handler.NewFechaVotacoesAbertasJobHandler(ucJobs.NewFechaVotacoesAbertasJobUseCase(repo))

	rec := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(rec)
	c.Request = httptest.NewRequest(http.MethodPost, "/internal/jobs/fecha_abertas", nil)

	h.Handle(c)

	return rec.Code, rec.Body.String()
}

func TestFechaVotacoesAbertas_SucessoRecebe200(t *testing.T) {
	repo := fakes.NewFakeJobRepository()

	status, corpo := chamar(t, repo)

	if status != http.StatusOK {
		t.Fatalf("esperava status %d, obteve %d: %s", http.StatusOK, status, corpo)
	}

	var resposta map[string]string
	if err := json.Unmarshal([]byte(corpo), &resposta); err != nil {
		t.Fatalf("corpo não é JSON: %v (%s)", err, corpo)
	}

	if resposta["message"] != "Job executado com sucesso" {
		t.Fatalf("mensagem inesperada: %s", corpo)
	}

	if repo.FecharVotacoesAbertasCalls != 1 {
		t.Fatalf("esperava 1 chamada ao repositório, obteve %d", repo.FecharVotacoesAbertasCalls)
	}
}

// Antes, uma falha do job virava 403 e devolvia o texto do erro ao chamador.
func TestFechaVotacoesAbertas_ErroRecebe500SemVazarDetalhes(t *testing.T) {
	repo := fakes.NewFakeJobRepository()
	repo.FecharVotacoesAbertasErr = errors.New("pq: deadlock detected (relation votacao)")

	status, corpo := chamar(t, repo)

	if status != http.StatusInternalServerError {
		t.Fatalf("esperava status %d, obteve %d", http.StatusInternalServerError, status)
	}

	if strings.Contains(corpo, "deadlock") || strings.Contains(corpo, "votacao") {
		t.Fatalf("a resposta vazou detalhes do erro: %s", corpo)
	}

	var resposta map[string]string
	if err := json.Unmarshal([]byte(corpo), &resposta); err != nil {
		t.Fatalf("corpo não é JSON: %v (%s)", err, corpo)
	}

	if resposta["error"] != "erro interno do servidor" {
		t.Fatalf("esperava a mensagem genérica, obteve: %s", corpo)
	}
}
