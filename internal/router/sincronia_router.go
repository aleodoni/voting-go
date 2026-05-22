package router

import "github.com/gin-gonic/gin"

func registerSincroniaRoutes(rg *gin.RouterGroup, h *Handlers) {
	rg.GET("/sincronia", h.RetornaUltimasSincronias.Handle)
	rg.POST("/sincronia", h.ExecutaSincronia.Handle)
}
