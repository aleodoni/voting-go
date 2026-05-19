package router

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func registerHealthRoutes(api *gin.RouterGroup) {
	api.GET("/health", healthHandler)
}

// healthHandler godoc
//
//	@Summary		Health check
//	@Description	Verifica se a API está no ar
//	@Tags			health
//	@Produce		json
//	@Success		200	{object}	map[string]string
//	@Router			/health [get]
func healthHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"status": "ok",
	})
}
