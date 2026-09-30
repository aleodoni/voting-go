// Package middleware provides HTTP middleware for authentication and authorization.
package middleware

import (
	"errors"
	"strings"
	"time"

	"github.com/MicahParks/keyfunc/v2"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"

	"github.com/aleodoni/voting-go/internal/config"
	jwtutil "github.com/aleodoni/voting-go/internal/platform/jwt"
)

// ErrInvalidToken indica um token que não pôde ser validado ou que não traz a
// identidade exigida pela API.
var ErrInvalidToken = errors.New("invalid token")

type JWTMiddleware struct {
	jwks *keyfunc.JWKS
	cfg  *config.Config
}

// NewJWTMiddleware carrega o JWKS do Keycloak, com novas tentativas se ele
// estiver indisponível, e devolve erro se não conseguir.
func NewJWTMiddleware(cfg *config.Config) (*JWTMiddleware, error) {
	options := keyfunc.Options{
		RefreshInterval:   5 * time.Minute, // Atualiza periodicamente
		RefreshTimeout:    10 * time.Second,
		RefreshUnknownKID: true, // Chave nova dispara refresh imediato
	}

	jwks, err := carregaJWKS(keyfunc.Get, cfg.JWKSURL, options, jwksTentativas, jwksEsperaInicial, jwksEsperaMaxima)
	if err != nil {
		return nil, err
	}

	return &JWTMiddleware{
		jwks: jwks,
		cfg:  cfg,
	}, nil
}

func (m *JWTMiddleware) Handler() gin.HandlerFunc {
	return func(c *gin.Context) {
		tokenString := extractToken(c)
		if tokenString == "" {
			c.AbortWithStatusJSON(401, gin.H{"error": "Missing token"})
			return
		}

		claims, err := m.ValidateToken(tokenString)
		if err != nil {
			c.AbortWithStatusJSON(401, gin.H{"error": "Invalid token"})
			return
		}

		keycloakID, username, ok := jwtutil.Identity(claims)
		if !ok {
			c.AbortWithStatusJSON(401, gin.H{"error": "Invalid token"})
			return
		}

		c.Set("claims", claims)
		c.Set("loggedUserKeycloakID", keycloakID)
		c.Set("loggedUserName", username)
		c.Next()
	}
}

// extractToken lê o JWT do header Authorization. O token não é aceito por query
// string: URLs aparecem em logs de proxy e de servidor.
func extractToken(c *gin.Context) string {
	return strings.TrimPrefix(c.GetHeader("Authorization"), "Bearer ")
}

func (m *JWTMiddleware) ValidateToken(tokenString string) (jwt.MapClaims, error) {
	token, err := jwt.Parse(tokenString, m.jwks.Keyfunc,
		jwt.WithIssuer(m.cfg.KeycloakIssuer),
		jwt.WithAudience(m.cfg.KeycloakClientID),
	)
	if err != nil {
		return nil, err
	}

	if token == nil || !token.Valid {
		return nil, ErrInvalidToken
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return nil, ErrInvalidToken
	}

	return claims, nil
}
