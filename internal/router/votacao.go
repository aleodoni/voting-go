package router

import "github.com/gin-gonic/gin"

func registerVotacaoRoutes(rg *gin.RouterGroup, h *Handlers) {
	rg.POST("/projetos/:projetoId/votacao/abrir", h.AbreVotacao.Handle)
	rg.POST("/projetos/:projetoId/votacao/fechar", h.FechaVotacao.Handle)
	rg.DELETE("/projetos/:projetoId/votacao", h.CancelaVotacao.Handle)
}
