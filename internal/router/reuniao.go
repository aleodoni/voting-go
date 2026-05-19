package router

import "github.com/gin-gonic/gin"

func registerReuniaoRoutes(rg *gin.RouterGroup, h *Handlers) {
	rg.GET("/reunioes-dia", h.RetornaReunioesDia.Handle)
	rg.GET("/reunioes/:reuniaoId/projetos", h.RetornaProjetosCompletos.Handle)
	rg.GET("/reunioes/:reuniaoId/relatorio", h.GeraRelatorioReuniao.Handle)
}
