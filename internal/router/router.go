// Package router provides routing functionality for the application.
package router

import (
	usuarioHandler "github.com/aleodoni/voting-go/internal/handler/usuario"
	"github.com/aleodoni/voting-go/internal/middleware"
	"github.com/gin-gonic/gin"
)

type Handlers struct {
	Me *usuarioHandler.MeHandler
}

func SetupRouter(jwtMiddleware *middleware.JWTMiddleware, h *Handlers) *gin.Engine {
	r := gin.New()

	r.Use(gin.Logger())
	r.Use(gin.Recovery())

	api := r.Group("/api/v1")

	registerHealthRoutes(api)
	registerProtectedRoutes(api, jwtMiddleware, h)

	return r
}
