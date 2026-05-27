package router

import (
	"github.com/aleodoni/voting-go/internal/middleware"
	"github.com/gin-gonic/gin"
)

func registerJobsRoutes(api *gin.RouterGroup, jobsMiddleware *middleware.InternalJobMiddleware, h *Handlers) {

	jobs := api.Group("/jobs")
	jobs.Use(jobsMiddleware.Handler())

	jobs.POST("/sincronia", h.ExecutaSincroniaJob.Handle)
	jobs.POST("/fecha_abertas", h.FechaVotacoesAbertasJob.Handle)
}
