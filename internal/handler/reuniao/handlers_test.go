package reuniao_test

import (
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/aleodoni/go-ddd/domain"
	"github.com/gin-gonic/gin"

	ucVotacao "github.com/aleodoni/voting-go/internal/application/votacao"
	domainUsuario "github.com/aleodoni/voting-go/internal/domain/usuario"
	handler "github.com/aleodoni/voting-go/internal/handler/reuniao"
	"github.com/aleodoni/voting-go/internal/test/fakes"
)

func admin() *domainUsuario.Usuario {
	return &domainUsuario.Usuario{
		AggregateRoot: domain.NewAggregateRoot("user-admin"),
		KeycloakID:    "keycloak-admin",
		Credencial:    &domainUsuario.Credencial{Ativo: true, PodeAdministrar: true},
	}
}

func vereador() *domainUsuario.Usuario {
	return &domainUsuario.Usuario{
		AggregateRoot: domain.NewAggregateRoot("user-vereador"),
		KeycloakID:    "keycloak-vereador",
		Credencial:    &domainUsuario.Credencial{Ativo: true, PodeVotar: true},
	}
}

// repositorios devolve os fakes já com o admin e o vereador cadastrados.
func repositorios() (*fakes.FakeUsuarioRepository, *fakes.FakeReuniaoRepository) {
	usuarioRepo := fakes.NewFakeUsuarioRepository()
	usuarioRepo.Seed(admin())
	usuarioRepo.Seed(vereador())

	return usuarioRepo, fakes.NewFakeReuniaoRepository()
}

func chamar(t *testing.T, handle gin.HandlerFunc, keycloakID string, params ...gin.Param) (int, string) {
	t.Helper()

	gin.SetMode(gin.TestMode)

	rec := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(rec)
	c.Request = httptest.NewRequest(http.MethodGet, "/teste", nil)
	c.Params = params
	c.Set("loggedUserKeycloakID", keycloakID)

	handle(c)

	return rec.Code, rec.Body.String()
}

func mensagemDeErro(t *testing.T, corpo string) string {
	t.Helper()

	var resposta map[string]string
	if err := json.Unmarshal([]byte(corpo), &resposta); err != nil {
		t.Fatalf("corpo não é um JSON {\"error\": ...}: %v (%s)", err, corpo)
	}

	return resposta["error"]
}

func verificaStatus(t *testing.T, esperado, obtido int) {
	t.Helper()

	if obtido != esperado {
		t.Fatalf("esperava status %d, obteve %d", esperado, obtido)
	}
}

// ─── GET /reunioes-dia ──────────────────────────────────────────────────────

func novoHandlerReunioesDia(usuarioRepo *fakes.FakeUsuarioRepository, reuniaoRepo *fakes.FakeReuniaoRepository) *handler.RetornaReunioesDiaHandler {
	return handler.NewRetornaReunioesDiaHandler(ucVotacao.NewRetornaReunioesDiaUseCase(usuarioRepo, reuniaoRepo))
}

func TestRetornaReunioesDia_AdminRecebe200(t *testing.T) {
	usuarioRepo, reuniaoRepo := repositorios()

	status, corpo := chamar(t, novoHandlerReunioesDia(usuarioRepo, reuniaoRepo).Handle, "keycloak-admin")

	verificaStatus(t, http.StatusOK, status)

	if corpo != "[]" {
		t.Fatalf("esperava lista vazia, obteve %s", corpo)
	}
}

func TestRetornaReunioesDia_NaoAdminRecebe403(t *testing.T) {
	usuarioRepo, reuniaoRepo := repositorios()

	status, corpo := chamar(t, novoHandlerReunioesDia(usuarioRepo, reuniaoRepo).Handle, "keycloak-vereador")

	verificaStatus(t, http.StatusForbidden, status)

	if mensagemDeErro(t, corpo) != "usuário não tem permissões de admin" {
		t.Fatalf("mensagem inesperada: %s", corpo)
	}
}

func TestRetornaReunioesDia_UsuarioInexistenteRecebe404(t *testing.T) {
	usuarioRepo, reuniaoRepo := repositorios()

	status, _ := chamar(t, novoHandlerReunioesDia(usuarioRepo, reuniaoRepo).Handle, "keycloak-inexistente")

	verificaStatus(t, http.StatusNotFound, status)
}

// Um erro inesperado do repositório vira 500 genérico e não vaza o texto
// original (antes virava 403 com o erro bruto).
func TestRetornaReunioesDia_ErroDoRepositorioRecebe500SemVazarDetalhes(t *testing.T) {
	usuarioRepo, reuniaoRepo := repositorios()
	reuniaoRepo.GetReunioesDiaErr = errors.New("pq: connection refused (host=10.0.0.5)")

	status, corpo := chamar(t, novoHandlerReunioesDia(usuarioRepo, reuniaoRepo).Handle, "keycloak-admin")

	verificaStatus(t, http.StatusInternalServerError, status)

	if strings.Contains(corpo, "connection refused") || strings.Contains(corpo, "10.0.0.5") {
		t.Fatalf("a resposta vazou detalhes do erro: %s", corpo)
	}

	if mensagemDeErro(t, corpo) != "erro interno do servidor" {
		t.Fatalf("esperava a mensagem genérica, obteve: %s", corpo)
	}
}

// ─── GET /reunioes/{reuniaoId}/projetos ─────────────────────────────────────

func novoHandlerProjetos(usuarioRepo *fakes.FakeUsuarioRepository, reuniaoRepo *fakes.FakeReuniaoRepository) *handler.RetornaProjetosCompletosHandler {
	return handler.NewRetornaProjetosCompletosHandler(ucVotacao.NewRetornaProjetosCompletosUseCase(usuarioRepo, reuniaoRepo))
}

func TestRetornaProjetosCompletos_NaoAdminRecebe403(t *testing.T) {
	usuarioRepo, reuniaoRepo := repositorios()

	status, _ := chamar(t, novoHandlerProjetos(usuarioRepo, reuniaoRepo).Handle, "keycloak-vereador", gin.Param{Key: "reuniaoId", Value: "r-1"})

	verificaStatus(t, http.StatusForbidden, status)
}

func TestRetornaProjetosCompletos_ReuniaoInexistenteRecebe404(t *testing.T) {
	usuarioRepo, reuniaoRepo := repositorios()

	status, corpo := chamar(t, novoHandlerProjetos(usuarioRepo, reuniaoRepo).Handle, "keycloak-admin", gin.Param{Key: "reuniaoId", Value: "inexistente"})

	verificaStatus(t, http.StatusNotFound, status)

	if mensagemDeErro(t, corpo) != "reunião não encontrada" {
		t.Fatalf("mensagem inesperada: %s", corpo)
	}
}

// ─── GET /projetos/{projetoId} ──────────────────────────────────────────────

func novoHandlerProjeto(usuarioRepo *fakes.FakeUsuarioRepository, reuniaoRepo *fakes.FakeReuniaoRepository) *handler.RetornaProjetoCompletoHandler {
	return handler.NewRetornaProjetoCompletoHandler(ucVotacao.NewRetornaProjetoCompletoUseCase(usuarioRepo, reuniaoRepo))
}

func TestRetornaProjetoCompleto_NaoAdminRecebe403(t *testing.T) {
	usuarioRepo, reuniaoRepo := repositorios()

	status, _ := chamar(t, novoHandlerProjeto(usuarioRepo, reuniaoRepo).Handle, "keycloak-vereador", gin.Param{Key: "projetoId", Value: "p-1"})

	verificaStatus(t, http.StatusForbidden, status)
}

func TestRetornaProjetoCompleto_ProjetoInexistenteRecebe404(t *testing.T) {
	usuarioRepo, reuniaoRepo := repositorios()

	status, corpo := chamar(t, novoHandlerProjeto(usuarioRepo, reuniaoRepo).Handle, "keycloak-admin", gin.Param{Key: "projetoId", Value: "inexistente"})

	verificaStatus(t, http.StatusNotFound, status)

	if mensagemDeErro(t, corpo) != "projeto não encontrado" {
		t.Fatalf("mensagem inesperada: %s", corpo)
	}
}
