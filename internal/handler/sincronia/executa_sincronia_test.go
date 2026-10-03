package sincronia_test

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/aleodoni/go-ddd/domain"
	"github.com/gin-gonic/gin"

	ucSincronia "github.com/aleodoni/voting-go/internal/application/sincronia"
	domainUsuario "github.com/aleodoni/voting-go/internal/domain/usuario"
	handler "github.com/aleodoni/voting-go/internal/handler/sincronia"
	"github.com/aleodoni/voting-go/internal/test/fakes"
)

func novoHandler(t *testing.T, appEnv string, usuarios ...*domainUsuario.Usuario) (*handler.ExecutaSincroniaHandler, *fakes.FakeSincroniaRepository) {
	t.Helper()

	usuarioRepo := fakes.NewFakeUsuarioRepository()
	for _, u := range usuarios {
		usuarioRepo.Seed(u)
	}

	sincroniaRepo := fakes.NewFakeSincroniaRepository()
	uc := ucSincronia.NewExecutaSincroniaUseCase(sincroniaRepo, usuarioRepo)

	return handler.NewExecutaSincroniaHandler(uc, appEnv), sincroniaRepo
}

func chamar(t *testing.T, h *handler.ExecutaSincroniaHandler, keycloakID string) (int, map[string]any) {
	t.Helper()

	gin.SetMode(gin.TestMode)

	rec := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(rec)
	c.Request = httptest.NewRequest(http.MethodPost, "/sincronia", nil)
	c.Set("loggedUserKeycloakID", keycloakID)

	h.Handle(c)

	var corpo map[string]any
	if err := json.Unmarshal(rec.Body.Bytes(), &corpo); err != nil {
		t.Fatalf("corpo não é JSON: %v (%s)", err, rec.Body.String())
	}

	return rec.Code, corpo
}

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

// Antes, o handler respondia 202 "Sincronia iniciada" a um não-admin e a recusa
// só aparecia no log, dentro da goroutine.
func TestExecutaSincronia_NaoAdminRecebe403EnaoIniciaSincronia(t *testing.T) {
	h, sincroniaRepo := novoHandler(t, "production", vereador())

	status, _ := chamar(t, h, "keycloak-vereador")

	if status != http.StatusForbidden {
		t.Fatalf("esperava status %d, obteve %d", http.StatusForbidden, status)
	}

	if sincroniaRepo.SyncCalls != 0 {
		t.Fatalf("não deveria iniciar a sincronização, chamadas a Sync: %d", sincroniaRepo.SyncCalls)
	}
}

func TestExecutaSincronia_UsuarioInexistenteRecebe404(t *testing.T) {
	h, sincroniaRepo := novoHandler(t, "production")

	status, _ := chamar(t, h, "keycloak-inexistente")

	if status != http.StatusNotFound {
		t.Fatalf("esperava status %d, obteve %d", http.StatusNotFound, status)
	}

	if sincroniaRepo.SyncCalls != 0 {
		t.Fatalf("não deveria iniciar a sincronização, chamadas a Sync: %d", sincroniaRepo.SyncCalls)
	}
}

// A permissão é checada antes do atalho de staging: um não-admin recebe 403
// também lá.
func TestExecutaSincronia_StagingNaoAdminRecebe403(t *testing.T) {
	h, _ := novoHandler(t, "staging", vereador())

	status, _ := chamar(t, h, "keycloak-vereador")

	if status != http.StatusForbidden {
		t.Fatalf("esperava status %d, obteve %d", http.StatusForbidden, status)
	}
}

func TestExecutaSincronia_StagingAdminIgnoraExecucao(t *testing.T) {
	h, sincroniaRepo := novoHandler(t, "staging", admin())

	status, corpo := chamar(t, h, "keycloak-admin")

	if status != http.StatusAccepted {
		t.Fatalf("esperava status %d, obteve %d", http.StatusAccepted, status)
	}

	if corpo["executed"] != false {
		t.Fatalf("esperava executed=false, obteve %v", corpo["executed"])
	}

	if sincroniaRepo.SyncCalls != 0 {
		t.Fatalf("staging não deveria executar a sincronização, chamadas a Sync: %d", sincroniaRepo.SyncCalls)
	}
}

func TestExecutaSincronia_AdminRecebe202(t *testing.T) {
	h, _ := novoHandler(t, "production", admin())

	status, corpo := chamar(t, h, "keycloak-admin")

	if status != http.StatusAccepted {
		t.Fatalf("esperava status %d, obteve %d", http.StatusAccepted, status)
	}

	if corpo["executed"] != true {
		t.Fatalf("esperava executed=true, obteve %v", corpo["executed"])
	}
}
