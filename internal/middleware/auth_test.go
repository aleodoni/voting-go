package middleware

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/MicahParks/keyfunc/v2"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"

	"github.com/aleodoni/voting-go/internal/config"
)

const (
	testKID      = "test-key"
	testSecret   = "segredo-de-teste-com-32-bytes-ok!"
	testIssuer   = "http://keycloak.test/realms/voting-realm"
	testClientID = "voting-api"
)

func novoJWTMiddlewareDeTeste() *JWTMiddleware {
	jwks := keyfunc.NewGiven(map[string]keyfunc.GivenKey{
		testKID: keyfunc.NewGivenHMAC([]byte(testSecret), keyfunc.GivenKeyOptions{Algorithm: "HS256"}),
	})

	return &JWTMiddleware{
		jwks: jwks,
		cfg: &config.Config{
			KeycloakIssuer:   testIssuer,
			KeycloakClientID: testClientID,
		},
	}
}

func claimsValidas() jwt.MapClaims {
	return jwt.MapClaims{
		"iss":                testIssuer,
		"aud":                testClientID,
		"exp":                time.Now().Add(time.Hour).Unix(),
		"sub":                "keycloak-123",
		"preferred_username": "maria",
	}
}

func tokenAssinado(t *testing.T, secret string, claims jwt.MapClaims) string {
	t.Helper()

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	token.Header["kid"] = testKID

	assinado, err := token.SignedString([]byte(secret))
	if err != nil {
		t.Fatalf("falha ao assinar o token de teste: %v", err)
	}

	return assinado
}

// requisitar executa a requisição contra uma rota protegida pelo middleware.
// O router não usa gin.Recovery: um panic no middleware falha o teste em vez
// de virar um 500 silencioso.
func requisitar(t *testing.T, m *JWTMiddleware, authorization string) *httptest.ResponseRecorder {
	t.Helper()

	defer func() {
		if r := recover(); r != nil {
			t.Fatalf("o middleware entrou em panic: %v", r)
		}
	}()

	gin.SetMode(gin.TestMode)

	router := gin.New()
	router.GET("/protegido", m.Handler(), func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"id":   c.GetString("loggedUserKeycloakID"),
			"nome": c.GetString("loggedUserName"),
		})
	})

	req := httptest.NewRequest(http.MethodGet, "/protegido", nil)
	if authorization != "" {
		req.Header.Set("Authorization", authorization)
	}

	rec := httptest.NewRecorder()
	router.ServeHTTP(rec, req)

	return rec
}

func TestHandler_TokenValidoLiberaAcessoEPopulaContexto(t *testing.T) {
	m := novoJWTMiddlewareDeTeste()
	token := tokenAssinado(t, testSecret, claimsValidas())

	rec := requisitar(t, m, "Bearer "+token)

	if rec.Code != http.StatusOK {
		t.Fatalf("esperava status %d, obteve %d: %s", http.StatusOK, rec.Code, rec.Body.String())
	}

	esperado := `{"id":"keycloak-123","nome":"maria"}`
	if rec.Body.String() != esperado {
		t.Fatalf("esperava corpo %s, obteve %s", esperado, rec.Body.String())
	}
}

func TestHandler_SemTokenRetorna401(t *testing.T) {
	m := novoJWTMiddlewareDeTeste()

	rec := requisitar(t, m, "")

	if rec.Code != http.StatusUnauthorized {
		t.Fatalf("esperava status %d, obteve %d", http.StatusUnauthorized, rec.Code)
	}
}

func TestHandler_TokenSemIdentidadeCompletaRetorna401(t *testing.T) {
	casos := []struct {
		nome      string
		modificar func(claims jwt.MapClaims)
	}{
		{"sem sub", func(c jwt.MapClaims) { delete(c, "sub") }},
		{"sub vazio", func(c jwt.MapClaims) { c["sub"] = "" }},
		{"sub que não é string", func(c jwt.MapClaims) { c["sub"] = 12345 }},
		{"sem preferred_username", func(c jwt.MapClaims) { delete(c, "preferred_username") }},
		{"preferred_username vazio", func(c jwt.MapClaims) { c["preferred_username"] = "" }},
		{"preferred_username que não é string", func(c jwt.MapClaims) { c["preferred_username"] = 42 }},
	}

	for _, tc := range casos {
		t.Run(tc.nome, func(t *testing.T) {
			m := novoJWTMiddlewareDeTeste()

			claims := claimsValidas()
			tc.modificar(claims)

			rec := requisitar(t, m, "Bearer "+tokenAssinado(t, testSecret, claims))

			if rec.Code != http.StatusUnauthorized {
				t.Fatalf("esperava status %d, obteve %d: %s", http.StatusUnauthorized, rec.Code, rec.Body.String())
			}
		})
	}
}

func TestHandler_TokenInvalidoRetorna401(t *testing.T) {
	m := novoJWTMiddlewareDeTeste()

	rec := requisitar(t, m, "Bearer isto-nao-e-um-jwt")

	if rec.Code != http.StatusUnauthorized {
		t.Fatalf("esperava status %d, obteve %d", http.StatusUnauthorized, rec.Code)
	}
}

func TestValidateToken_Rejeicoes(t *testing.T) {
	casos := []struct {
		nome   string
		secret string
		claims func() jwt.MapClaims
	}{
		{"assinatura de outra chave", "outro-segredo-com-32-bytes-de-teste", claimsValidas},
		{"issuer incorreto", testSecret, func() jwt.MapClaims {
			c := claimsValidas()
			c["iss"] = "http://outro-emissor.test/realms/x"

			return c
		}},
		{"audience incorreta", testSecret, func() jwt.MapClaims {
			c := claimsValidas()
			c["aud"] = "outro-client"

			return c
		}},
		{"token expirado", testSecret, func() jwt.MapClaims {
			c := claimsValidas()
			c["exp"] = time.Now().Add(-time.Hour).Unix()

			return c
		}},
	}

	for _, tc := range casos {
		t.Run(tc.nome, func(t *testing.T) {
			m := novoJWTMiddlewareDeTeste()

			claims, err := m.ValidateToken(tokenAssinado(t, tc.secret, tc.claims()))
			if err == nil {
				t.Fatal("esperava erro, obteve nil")
			}

			if claims != nil {
				t.Fatalf("esperava claims nil, obteve %v", claims)
			}
		})
	}
}

func TestValidateToken_TokenValidoRetornaClaims(t *testing.T) {
	m := novoJWTMiddlewareDeTeste()

	claims, err := m.ValidateToken(tokenAssinado(t, testSecret, claimsValidas()))
	if err != nil {
		t.Fatalf("esperava nil, obteve: %v", err)
	}

	if claims["sub"] != "keycloak-123" {
		t.Fatalf("esperava sub %q, obteve %v", "keycloak-123", claims["sub"])
	}
}
