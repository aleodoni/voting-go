// Package middleware provides HTTP middleware for authentication and authorization.
package middleware

import (
	"strings"
	"time"

	"github.com/MicahParks/keyfunc/v2"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"

	"github.com/aleodoni/voting-go/internal/config"
)

type JWTMiddleware struct {
	jwks *keyfunc.JWKS
	cfg  *config.Config
}

func NewJWTMiddleware(cfg *config.Config) *JWTMiddleware {
	jwks, err := keyfunc.Get(cfg.JWKSURL, keyfunc.Options{
		RefreshInterval: time.Hour,
	})

	if err != nil {
		panic("failed to get JWKS: " + err.Error())
	}

	return &JWTMiddleware{
		jwks: jwks,
		cfg:  cfg,
	}
}

func (m *JWTMiddleware) Handler() gin.HandlerFunc {
	return func(c *gin.Context) {
		tokenString := extractToken(c)

		if tokenString == "" {
			c.AbortWithStatusJSON(401, gin.H{"error": "Missing token"})
			return
		}

		token, err := jwt.Parse(tokenString, m.jwks.Keyfunc)

		if err != nil || !token.Valid {
			c.AbortWithStatusJSON(401, gin.H{"error": "Invalid token"})
			return
		}

		claims := token.Claims.(jwt.MapClaims)

		loggedUserKeycloakID, ok := claims["sub"].(string)

		if !ok {
			c.AbortWithStatusJSON(401, gin.H{"error": "Subject (sub) not found in token"})
			return
		}

		loggedUserName := claims["preferred_username"].(string)

		c.Set("loggedUserKeycloakID", loggedUserKeycloakID)
		c.Set("loggedUserName", loggedUserName)
		c.Set("claims", claims)

		c.Next()
	}
}

func extractToken(c *gin.Context) string {
	authHeader := c.GetHeader("Authorization")
	if authHeader == "" {
		return ""
	}
	return strings.TrimPrefix(authHeader, "Bearer ")
}
