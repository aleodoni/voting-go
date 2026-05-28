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
//	@Summary		Executa job fechar votações abertas
//	@Description	Executa o fechamento de votações abertas
//	@Tags			jobs
//	@Produce		json
//	@Success		200
//	@Failure		403
//	@Security		BearerAuth
//	@Router			/internal/jobs/fecha_votacoes_abertas [post]
func (h *FechaVotacoesAbertasJobHandler) Handle(c *gin.Context) {

	err := h.fechaVotacoesAbertasUseCase.Execute(c.Request.Context())
	if err != nil {
		c.JSON(http.StatusForbidden, gin.H{"error": "Erro ao executar job: " + err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Job executado com sucesso"})
}
