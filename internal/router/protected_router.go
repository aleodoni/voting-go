package router

import (
	"github.com/aleodoni/voting-go/internal/middleware"
	"github.com/gin-gonic/gin"
)

func registerProtectedRoutes(api *gin.RouterGroup, jwtMiddleware *middleware.JWTMiddleware, h *Handlers) {

	protected := api.Group("")
	protected.Use(jwtMiddleware.Handler())

	protected.GET("/test", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"status": "route protected",
		})
	})

	registerUsuarioRoutes(protected, h)

}
