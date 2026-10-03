package relatorio_test

import (
	"context"
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"time"

	"github.com/gin-gonic/gin"

	ucRelatorio "github.com/aleodoni/voting-go/internal/application/relatorio"
	domainRelatorio "github.com/aleodoni/voting-go/internal/domain/relatorio"
	"github.com/aleodoni/voting-go/internal/domain/votacao"
	handler "github.com/aleodoni/voting-go/internal/handler/relatorio"
	"github.com/aleodoni/voting-go/internal/test/fakes"
)

// fakeGenerator implementa domainRelatorio.Generator.
type fakeGenerator struct {
	pdf      []byte
	err      error
	chamadas int
}

func (f *fakeGenerator) Gera(_ context.Context, _ domainRelatorio.ReuniaoOutput) ([]byte, error) {
	f.chamadas++

	return f.pdf, f.err
}

func chamar(t *testing.T, reuniaoRepo *fakes.FakeReuniaoRepository, gerador *fakeGenerator, reuniaoID string) *httptest.ResponseRecorder {
	t.Helper()

	gin.SetMode(gin.TestMode)

	h := handler.NewGeraRelatorioReuniaoHandler(ucRelatorio.NewGeraRelatorioReuniaoUseCase(reuniaoRepo, gerador))

	rec := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(rec)
	c.Request = httptest.NewRequest(http.MethodGet, "/teste", nil)
	c.Params = gin.Params{{Key: "reuniaoId", Value: reuniaoID}}

	h.Handle(c)

	return rec
}

func mensagemDeErro(t *testing.T, corpo string) string {
	t.Helper()

	var resposta map[string]string
	if err := json.Unmarshal([]byte(corpo), &resposta); err != nil {
		t.Fatalf("corpo não é um JSON {\"error\": ...}: %v (%s)", err, corpo)
	}

	return resposta["error"]
}

func TestGeraRelatorio_ReuniaoExistenteRetornaPDF(t *testing.T) {
	reuniaoRepo := fakes.NewFakeReuniaoRepository()
	reuniaoRepo.Seed(&votacao.Reuniao{ID: "r-1", RecData: time.Date(2026, 10, 1, 0, 0, 0, 0, time.UTC)})

	gerador := &fakeGenerator{pdf: []byte("%PDF-conteudo-de-teste")}

	rec := chamar(t, reuniaoRepo, gerador, "r-1")

	if rec.Code != http.StatusOK {
		t.Fatalf("esperava status %d, obteve %d: %s", http.StatusOK, rec.Code, rec.Body.String())
	}

	if tipo := rec.Header().Get("Content-Type"); tipo != "application/pdf" {
		t.Fatalf("esperava Content-Type application/pdf, obteve %q", tipo)
	}

	if disp := rec.Header().Get("Content-Disposition"); disp != "inline; filename=reuniao-r-1.pdf" {
		t.Fatalf("Content-Disposition inesperado: %q", disp)
	}

	if rec.Body.String() != "%PDF-conteudo-de-teste" {
		t.Fatalf("o corpo deveria ser o PDF gerado, obteve %q", rec.Body.String())
	}
}

// Antes, uma reunião inexistente virava 500 com o texto do erro.
func TestGeraRelatorio_ReuniaoInexistenteRecebe404(t *testing.T) {
	gerador := &fakeGenerator{}

	rec := chamar(t, fakes.NewFakeReuniaoRepository(), gerador, "inexistente")

	if rec.Code != http.StatusNotFound {
		t.Fatalf("esperava status %d, obteve %d", http.StatusNotFound, rec.Code)
	}

	if mensagemDeErro(t, rec.Body.String()) != "reunião não encontrada" {
		t.Fatalf("mensagem inesperada: %s", rec.Body.String())
	}

	if gerador.chamadas != 0 {
		t.Fatalf("não deveria gerar o PDF de uma reunião inexistente, chamadas: %d", gerador.chamadas)
	}
}

func TestGeraRelatorio_ErroDoGeradorRecebe500SemVazarDetalhes(t *testing.T) {
	reuniaoRepo := fakes.NewFakeReuniaoRepository()
	reuniaoRepo.Seed(&votacao.Reuniao{ID: "r-1"})

	gerador := &fakeGenerator{err: errors.New(`exec: "wkhtmltopdf": executable file not found in $PATH`)}

	rec := chamar(t, reuniaoRepo, gerador, "r-1")

	if rec.Code != http.StatusInternalServerError {
		t.Fatalf("esperava status %d, obteve %d", http.StatusInternalServerError, rec.Code)
	}

	if strings.Contains(rec.Body.String(), "wkhtmltopdf") {
		t.Fatalf("a resposta vazou detalhes do erro: %s", rec.Body.String())
	}

	if mensagemDeErro(t, rec.Body.String()) != "erro interno do servidor" {
		t.Fatalf("esperava a mensagem genérica, obteve: %s", rec.Body.String())
	}
}
