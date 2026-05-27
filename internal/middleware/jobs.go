package middleware

import (
	"crypto/subtle"
	"net/http"
	"strings"

	"github.com/aleodoni/voting-go/internal/config"
	"github.com/gin-gonic/gin"
)

type InternalJobMiddleware struct {
	cfg *config.Config
}

func NewInternalJobMiddleware(cfg *config.Config) *InternalJobMiddleware {
	return &InternalJobMiddleware{
		cfg: cfg,
	}
}

func (m *InternalJobMiddleware) Handler() gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")

		if authHeader == "" {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"error": "missing authorization header",
			})
			return
		}

		token := strings.TrimPrefix(authHeader, "Bearer ")

		if subtle.ConstantTimeCompare(
			[]byte(token),
			[]byte(m.cfg.JobsToken),
		) != 1 {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"error": "invalid token",
			})
			return
		}

		c.Next()
	}
}
