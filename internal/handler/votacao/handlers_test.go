package votacao_test

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
	"github.com/aleodoni/voting-go/internal/domain/votacao"
	handler "github.com/aleodoni/voting-go/internal/handler/votacao"
	"github.com/aleodoni/voting-go/internal/platform/event"
	"github.com/aleodoni/voting-go/internal/test/fakes"
)

// ─── cenário ────────────────────────────────────────────────────────────────

type cenario struct {
	usuarios *fakes.FakeUsuarioRepository
	reunioes *fakes.FakeReuniaoRepository
	votacoes *fakes.FakeVotacaoRepository
}

func usuario(id, keycloakID string, credencial *domainUsuario.Credencial) *domainUsuario.Usuario {
	return &domainUsuario.Usuario{
		AggregateRoot: domain.NewAggregateRoot(id),
		KeycloakID:    keycloakID,
		Credencial:    credencial,
	}
}

// novoCenario cadastra um admin, um vereador e um usuário ativo sem permissão
// de voto nem de administração.
func novoCenario() *cenario {
	usuarios := fakes.NewFakeUsuarioRepository()
	usuarios.Seed(usuario("user-admin", "keycloak-admin", &domainUsuario.Credencial{Ativo: true, PodeAdministrar: true}))
	usuarios.Seed(usuario("user-vereador", "keycloak-vereador", &domainUsuario.Credencial{Ativo: true, PodeVotar: true}))
	usuarios.Seed(usuario("user-sem-voto", "keycloak-sem-voto", &domainUsuario.Credencial{Ativo: true}))

	return &cenario{
		usuarios: usuarios,
		reunioes: fakes.NewFakeReuniaoRepository(),
		votacoes: fakes.NewFakeVotacaoRepository(),
	}
}

func (c *cenario) abre() *handler.AbreVotacaoHandler {
	return handler.NewAbreVotacaoHandler(ucVotacao.NewAbreVotacaoUseCase(c.usuarios, c.reunioes, c.votacoes, event.NewBus()))
}

func (c *cenario) fecha() *handler.FechaVotacaoHandler {
	return handler.NewFechaVotacaoHandler(ucVotacao.NewFechaVotacaoUseCase(c.usuarios, c.reunioes, c.votacoes, event.NewBus()))
}

func (c *cenario) cancela() *handler.CancelaVotacaoHandler {
	return handler.NewCancelaVotacaoHandler(ucVotacao.NewCancelaVotacaoUseCase(c.usuarios, c.reunioes, c.votacoes, event.NewBus()))
}

func (c *cenario) registraVoto() *handler.RegistraVotoHandler {
	return handler.NewRegistraVotoHandler(ucVotacao.NewRegistraVotoUseCase(c.usuarios, c.votacoes, event.NewBus()))
}

func (c *cenario) stats() *handler.RetornaVotingStatsHandler {
	return handler.NewRetornaVotingStatsHandler(ucVotacao.NewRetornaVotingStatsUseCase(c.usuarios, c.votacoes))
}

func (c *cenario) votacaoAberta() *handler.RetornaProjetoVotacaoAbertaHandler {
	return handler.NewRetornaProjetoVotacaoAbertaHandler(ucVotacao.NewRetornaVotacaoAbertaUseCase(c.usuarios, c.votacoes))
}

// ─── helpers de requisição ──────────────────────────────────────────────────

func chamar(t *testing.T, handle gin.HandlerFunc, keycloakID, corpo string, params ...gin.Param) (int, string) {
	t.Helper()

	gin.SetMode(gin.TestMode)

	rec := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(rec)
	c.Request = httptest.NewRequest(http.MethodPost, "/teste", strings.NewReader(corpo))
	c.Request.Header.Set("Content-Type", "application/json")
	c.Params = params
	c.Set("loggedUserKeycloakID", keycloakID)

	handle(c)

	// c.Status(204) não grava o código no ResponseRecorder; o writer do gin tem o valor certo.
	return c.Writer.Status(), rec.Body.String()
}

func projeto(id string) gin.Param { return gin.Param{Key: "projetoId", Value: id} }

func votacaoID(id string) gin.Param { return gin.Param{Key: "votacaoId", Value: id} }

func mensagemDeErro(t *testing.T, corpo string) string {
	t.Helper()

	var resposta map[string]string
	if err := json.Unmarshal([]byte(corpo), &resposta); err != nil {
		t.Fatalf("corpo não é um JSON {\"error\": ...}: %v (%s)", err, corpo)
	}

	return resposta["error"]
}

func esperaErro(t *testing.T, status int, corpo string, statusEsperado int, mensagemEsperada string) {
	t.Helper()

	if status != statusEsperado {
		t.Fatalf("esperava status %d, obteve %d: %s", statusEsperado, status, corpo)
	}

	if mensagem := mensagemDeErro(t, corpo); mensagem != mensagemEsperada {
		t.Fatalf("esperava a mensagem %q, obteve %q", mensagemEsperada, mensagem)
	}
}

func votacaoCom(id string, status votacao.StatusVotacao) *votacao.Votacao {
	return &votacao.Votacao{AggregateRoot: domain.NewAggregateRoot(id), Status: status}
}

func projetoComVotacao(status votacao.StatusVotacao) *votacao.Projeto {
	return &votacao.Projeto{
		ID:               "projeto-1",
		CodigoProposicao: "001",
		Votacao:          votacaoCom("votacao-1", status),
	}
}

// ─── POST /projetos/{projetoId}/votacao/abrir ───────────────────────────────

func TestAbreVotacao_NaoAdminRecebe403(t *testing.T) {
	c := novoCenario()

	status, corpo := chamar(t, c.abre().Handle, "keycloak-vereador", "", projeto("projeto-1"))

	esperaErro(t, status, corpo, http.StatusForbidden, "usuário não tem permissões de admin")
}

func TestAbreVotacao_ProjetoInexistenteRecebe404(t *testing.T) {
	c := novoCenario()

	status, corpo := chamar(t, c.abre().Handle, "keycloak-admin", "", projeto("projeto-inexistente"))

	esperaErro(t, status, corpo, http.StatusNotFound, "projeto não encontrado")
}

func TestAbreVotacao_JaExisteVotacaoAbertaRecebe422(t *testing.T) {
	c := novoCenario()
	c.votacoes.Seed(votacaoCom("votacao-existente", votacao.StatusVotacaoA))

	status, corpo := chamar(t, c.abre().Handle, "keycloak-admin", "", projeto("projeto-1"))

	esperaErro(t, status, corpo, http.StatusUnprocessableEntity, "já existe uma votação aberta")
}

func TestAbreVotacao_ProjetoJaVotadoRecebe422(t *testing.T) {
	c := novoCenario()
	c.reunioes.SeedProjetos("reuniao-1", []*votacao.Projeto{projetoComVotacao(votacao.StatusVotacaoA)})

	status, corpo := chamar(t, c.abre().Handle, "keycloak-admin", "", projeto("projeto-1"))

	esperaErro(t, status, corpo, http.StatusUnprocessableEntity, "projeto já votado")
}

// ─── POST /projetos/{projetoId}/votacao/fechar ──────────────────────────────

func TestFechaVotacao_VotacaoNaoAbertaRecebe422(t *testing.T) {
	c := novoCenario()
	c.reunioes.SeedProjetos("reuniao-1", []*votacao.Projeto{projetoComVotacao(votacao.StatusVotacaoF)})

	status, corpo := chamar(t, c.fecha().Handle, "keycloak-admin", "", projeto("projeto-1"))

	esperaErro(t, status, corpo, http.StatusUnprocessableEntity, "votação não está aberta")
}

func TestFechaVotacao_ProjetoSemVotacaoRecebe404(t *testing.T) {
	c := novoCenario()
	c.reunioes.SeedProjetos("reuniao-1", []*votacao.Projeto{{ID: "projeto-1", CodigoProposicao: "001"}})

	status, corpo := chamar(t, c.fecha().Handle, "keycloak-admin", "", projeto("projeto-1"))

	esperaErro(t, status, corpo, http.StatusNotFound, "votação não encontrada")
}

// ─── DELETE /projetos/{projetoId}/votacao ───────────────────────────────────

func TestCancelaVotacao_VotacaoNaoFechadaRecebe422(t *testing.T) {
	c := novoCenario()
	c.reunioes.SeedProjetos("reuniao-1", []*votacao.Projeto{projetoComVotacao(votacao.StatusVotacaoA)})

	status, corpo := chamar(t, c.cancela().Handle, "keycloak-admin", "", projeto("projeto-1"))

	esperaErro(t, status, corpo, http.StatusUnprocessableEntity, "votação não está fechada")
}

// ─── POST /votacao/{votacaoId}/voto ─────────────────────────────────────────

const votoF = `{"voto":"F"}`

func TestRegistraVoto_UsuarioJaVotouRecebe422(t *testing.T) {
	c := novoCenario()
	c.votacoes.Seed(votacaoCom("votacao-1", votacao.StatusVotacaoA))

	h := c.registraVoto()

	// primeiro voto: aceito
	status, corpo := chamar(t, h.Handle, "keycloak-vereador", votoF, votacaoID("votacao-1"))
	if status != http.StatusNoContent {
		t.Fatalf("primeiro voto: esperava status %d, obteve %d: %s", http.StatusNoContent, status, corpo)
	}

	// segundo voto: regra de negócio violada
	status, corpo = chamar(t, h.Handle, "keycloak-vereador", votoF, votacaoID("votacao-1"))

	esperaErro(t, status, corpo, http.StatusUnprocessableEntity, "usuário já votou nesta votação")
}

func TestRegistraVoto_VotacaoFechadaRecebe422(t *testing.T) {
	c := novoCenario()
	c.votacoes.Seed(votacaoCom("votacao-1", votacao.StatusVotacaoF))

	status, corpo := chamar(t, c.registraVoto().Handle, "keycloak-vereador", votoF, votacaoID("votacao-1"))

	esperaErro(t, status, corpo, http.StatusUnprocessableEntity, "votação não está aberta")
}

func TestRegistraVoto_UsuarioSemPermissaoDeVotoRecebe403(t *testing.T) {
	c := novoCenario()
	c.votacoes.Seed(votacaoCom("votacao-1", votacao.StatusVotacaoA))

	status, corpo := chamar(t, c.registraVoto().Handle, "keycloak-sem-voto", votoF, votacaoID("votacao-1"))

	esperaErro(t, status, corpo, http.StatusForbidden, "usuário não tem permissões de votante")
}

// A validação do corpo continua sendo 400, não passa pelo httperr.
func TestRegistraVoto_CorpoInvalidoRecebe400(t *testing.T) {
	c := novoCenario()

	status, corpo := chamar(t, c.registraVoto().Handle, "keycloak-vereador", `{}`, votacaoID("votacao-1"))

	if status != http.StatusBadRequest {
		t.Fatalf("esperava status %d, obteve %d: %s", http.StatusBadRequest, status, corpo)
	}
}

// ─── GET /votacao/stats ─────────────────────────────────────────────────────

func TestRetornaStats_NaoAdminRecebe403(t *testing.T) {
	c := novoCenario()

	status, corpo := chamar(t, c.stats().Handle, "keycloak-vereador", "")

	esperaErro(t, status, corpo, http.StatusForbidden, "usuário não tem permissões de admin")
}

// Antes, uma falha do repositório virava 403 com o texto do erro.
func TestRetornaStats_ErroDoRepositorioRecebe500SemVazarDetalhes(t *testing.T) {
	c := novoCenario()
	c.votacoes.GetVotingStatsErr = errors.New("pq: connection refused (host=10.0.0.5)")

	status, corpo := chamar(t, c.stats().Handle, "keycloak-admin", "")

	esperaErro(t, status, corpo, http.StatusInternalServerError, "erro interno do servidor")

	if strings.Contains(corpo, "connection refused") || strings.Contains(corpo, "10.0.0.5") {
		t.Fatalf("a resposta vazou detalhes do erro: %s", corpo)
	}
}

// ─── GET /votacao/aberta ────────────────────────────────────────────────────

func TestRetornaVotacaoAberta_UsuarioInexistenteRecebe404(t *testing.T) {
	c := novoCenario()

	status, corpo := chamar(t, c.votacaoAberta().Handle, "keycloak-inexistente", "")

	esperaErro(t, status, corpo, http.StatusNotFound, "usuário não encontrado")
}
