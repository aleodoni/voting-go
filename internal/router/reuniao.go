package router

import "github.com/gin-gonic/gin"

func registerReuniaoRoutes(rg *gin.RouterGroup, h *Handlers) {
	rg.GET("/reunioes-dia", h.RetornaReunioesDia.Handle)

}
