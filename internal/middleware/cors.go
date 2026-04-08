package middleware

import (
	"strings"
	"time"

	"github.com/aleodoni/voting-go/internal/config"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func CORSMiddleware(cfg *config.Config) gin.HandlerFunc {
	allowOrigins := cfg.AllowOrigins
	for i := range allowOrigins {
		allowOrigins[i] = strings.TrimSpace(allowOrigins[i])
	}

	return cors.New(cors.Config{
		AllowOrigins:           allowOrigins,
		AllowMethods:           []string{"GET", "POST", "PUT", "PATCH", "DELETE", "OPTIONS"},
		AllowHeaders:           []string{"Origin", "Content-Type", "Authorization"},
		ExposeHeaders:          []string{"Content-Length"},
		AllowCredentials:       true,
		MaxAge:                 12 * time.Hour,
		AllowBrowserExtensions: true, // evita rejeição de origens inesperadas
	})
}
