package router

import "github.com/gin-gonic/gin"

func registerCredencialRoutes(rg *gin.RouterGroup, h *Handlers) {
	rg.PATCH("/usuarios/:id/credencial", h.UpdateCredenciais.Handle)
}
