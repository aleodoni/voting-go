package jobs

import (
	"net/http"

	"github.com/gin-gonic/gin"

	ucJobs "github.com/aleodoni/voting-go/internal/application/jobs"
)

type FechaVotacoesAbertasJobHandler struct {
	fechaVotacoesAbertasUseCase *ucJobs.FechaVotacoesAbertasJobUseCase
	appEnv                      string
}

func NewFechaVotacoesAbertasJobHandler(fechaVotacoesAbertasUseCase *ucJobs.FechaVotacoesAbertasJobUseCase) *FechaVotacoesAbertasJobHandler {
	return &FechaVotacoesAbertasJobHandler{fechaVotacoesAbertasUseCase: fechaVotacoesAbertasUseCase}
}

// Handle godoc
//
//	@Summary		Executa job de fechamento de votações abertas
//	@Description	Fecha todas as votações que estejam abertas. Autenticado por token estático (JOBS_TOKEN).
//	@Tags			jobs
//	@Produce		json
//	@Success		200	{object}	map[string]interface{}	"Job executado com sucesso"
//	@Failure		401	{object}	map[string]interface{}	"Token inválido ou ausente"
//	@Failure		403	{object}	map[string]interface{}	"Erro ao executar job"
//	@Security		BearerAuth
//	@Router			/internal/jobs/fecha_abertas [post]
func (h *FechaVotacoesAbertasJobHandler) Handle(c *gin.Context) {

	err := h.fechaVotacoesAbertasUseCase.Execute(c.Request.Context())
	if err != nil {
		c.JSON(http.StatusForbidden, gin.H{"error": "Erro ao executar job: " + err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Job executado com sucesso"})
}
