package router

import "github.com/gin-gonic/gin"

func registerUsuarioRoutes(rg *gin.RouterGroup, h *Handlers) {
	rg.GET("/me", h.Me.Handle)
	rg.PUT(
		"/usuarios/fantasia-credenciais",
		h.UpdateFantasiaCredenciais.Handle,
	)
}
