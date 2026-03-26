package router

import "github.com/gin-gonic/gin"

func registerUsuarioRoutes(rg *gin.RouterGroup, h *Handlers) {
	rg.GET("/me", h.Me.Handle)
	rg.PUT(
		"/usuarios/fantasia-credenciais",
		h.UpdateFantasiaCredenciais.Handle,
	)
	rg.PUT(
		"/usuarios/fantasia",
		h.UpdateFantasia.Handle,
	)
	rg.GET("/usuarios", h.PesquisaUsuarios.Handle)
	rg.GET("/usuarios-conectados", h.ConnectedUsers.Handle)
}
