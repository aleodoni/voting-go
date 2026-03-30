package router

import "github.com/gin-gonic/gin"

func registerVotacaoRoutes(rg *gin.RouterGroup, h *Handlers) {
	rg.GET("/projetos/:projetoId", h.RetornaProjetoCompleto.Handle)
	rg.POST("/projetos/:projetoId/votacao/abrir", h.AbreVotacao.Handle)
	rg.POST("/projetos/:projetoId/votacao/fechar", h.FechaVotacao.Handle)
	rg.DELETE("/projetos/:projetoId/votacao", h.CancelaVotacao.Handle)
	rg.POST("/votacao/:votacaoId/voto", h.RegistraVoto.Handle)
	rg.GET("/votacao/aberta", h.RetornaProjetoVotacaoAberta.Handle)
	rg.GET("/votacao/stats", h.RetornaStatsVotacao.Handle)
}
