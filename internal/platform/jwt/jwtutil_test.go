package jwtutil_test

import (
	"testing"

	"github.com/golang-jwt/jwt/v5"

	jwtutil "github.com/aleodoni/voting-go/internal/platform/jwt"
)

func TestClaimString(t *testing.T) {
	claims := jwt.MapClaims{
		"texto":  "valor",
		"numero": 42,
		"nulo":   nil,
	}

	casos := []struct {
		nome     string
		claims   jwt.MapClaims
		chave    string
		esperado string
	}{
		{"claim string", claims, "texto", "valor"},
		{"claim ausente", claims, "ausente", ""},
		{"claim que não é string", claims, "numero", ""},
		{"claim nula", claims, "nulo", ""},
		{"claims nil", nil, "texto", ""},
	}

	for _, tc := range casos {
		t.Run(tc.nome, func(t *testing.T) {
			obtido := jwtutil.ClaimString(tc.claims, tc.chave)
			if obtido != tc.esperado {
				t.Fatalf("esperava %q, obteve %q", tc.esperado, obtido)
			}
		})
	}
}

func TestIdentity_ClaimsValidas(t *testing.T) {
	claims := jwt.MapClaims{
		"sub":                "keycloak-123",
		"preferred_username": "maria",
	}

	keycloakID, username, ok := jwtutil.Identity(claims)
	if !ok {
		t.Fatal("esperava ok = true")
	}

	if keycloakID != "keycloak-123" {
		t.Fatalf("esperava keycloakID %q, obteve %q", "keycloak-123", keycloakID)
	}

	if username != "maria" {
		t.Fatalf("esperava username %q, obteve %q", "maria", username)
	}
}

func TestIdentity_ClaimsInvalidas(t *testing.T) {
	casos := []struct {
		nome   string
		claims jwt.MapClaims
	}{
		{"claims nil", nil},
		{"claims vazias", jwt.MapClaims{}},
		{"sem sub", jwt.MapClaims{"preferred_username": "maria"}},
		{"sem preferred_username", jwt.MapClaims{"sub": "keycloak-123"}},
		{"sub vazio", jwt.MapClaims{"sub": "", "preferred_username": "maria"}},
		{"preferred_username vazio", jwt.MapClaims{"sub": "keycloak-123", "preferred_username": ""}},
		{"sub que não é string", jwt.MapClaims{"sub": 123, "preferred_username": "maria"}},
		{"preferred_username que não é string", jwt.MapClaims{"sub": "keycloak-123", "preferred_username": 42}},
		{"sub nulo", jwt.MapClaims{"sub": nil, "preferred_username": "maria"}},
	}

	for _, tc := range casos {
		t.Run(tc.nome, func(t *testing.T) {
			if _, _, ok := jwtutil.Identity(tc.claims); ok {
				t.Fatal("esperava ok = false")
			}
		})
	}
}
