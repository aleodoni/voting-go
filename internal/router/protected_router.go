package router

import (
	"github.com/aleodoni/voting-go/internal/middleware"
	"github.com/gin-gonic/gin"
)

func registerProtectedRoutes(api *gin.RouterGroup, jwtMiddleware *middleware.JWTMiddleware, h *Handlers) {

	protected := api.Group("")
	protected.Use(jwtMiddleware.Handler())

	registerUsuarioRoutes(protected, h)
	registerCredencialRoutes(protected, h)
	registerReuniaoRoutes(protected, h)
	registerVotacaoRoutes(protected, h)
	registerSincroniaRoutes(protected, h)
}
