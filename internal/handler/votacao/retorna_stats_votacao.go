package votacao

import (
	"net/http"

	"github.com/gin-gonic/gin"

	ucVotacao "github.com/aleodoni/voting-go/internal/application/votacao"
)

type VotingStatsResponse struct {
	TotalProjects      int64 `json:"total_projects"`
	TotalVotedProjects int64 `json:"total_voted_projects"`
}

type RetornaVotingStatsHandler struct {
	retornaVotingStatsUseCase *ucVotacao.RetornaVotingStatsUseCase
}

func NewRetornaVotingStatsHandler(retornaVotingStatsUseCase *ucVotacao.RetornaVotingStatsUseCase) *RetornaVotingStatsHandler {
	return &RetornaVotingStatsHandler{retornaVotingStatsUseCase: retornaVotingStatsUseCase}
}

// Handle godoc
//
//	@Summary		Retorna estatísticas de votação do dia
//	@Description	Retorna o total de projetos e total de projetos votados no dia atual (requer admin)
//	@Tags			votação
//	@Produce		json
//	@Success		200	{object}	VotingStatsResponse
//	@Failure		403	{object}	ErrorResponse
//	@Security		BearerAuth
//	@Router			/votacao/stats [get]
func (h *RetornaVotingStatsHandler) Handle(c *gin.Context) {
	loggedUserKeycloakID := c.GetString("loggedUserKeycloakID")

	input := ucVotacao.RetornaVotingStatsInput{
		LoggedInUserKeycloakID: loggedUserKeycloakID,
	}

	stats, err := h.retornaVotingStatsUseCase.Execute(c.Request.Context(), input)
	if err != nil {
		c.JSON(http.StatusForbidden, ErrorResponse{Error: err.Error()})
		return
	}

	c.JSON(http.StatusOK, VotingStatsResponse{
		TotalProjects:      stats.TotalProjects,
		TotalVotedProjects: stats.TotalVotedProjects,
	})
}
