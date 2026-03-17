package votacao

import (
	"net/http"

	"github.com/aleodoni/voting-go/internal/application/votacao"
	"github.com/gin-gonic/gin"
)

type CancelaVotacaoHandler struct {
	cancelaVotacaoUseCase *votacao.CancelaVotacaoUseCase
}

func NewCancelaVotacaoHandler(cancelaVotacaoUseCase *votacao.CancelaVotacaoUseCase) *CancelaVotacaoHandler {
	return &CancelaVotacaoHandler{cancelaVotacaoUseCase: cancelaVotacaoUseCase}
}

func (h *CancelaVotacaoHandler) Handle(c *gin.Context) {
	loggedUserKeycloakID := c.GetString("loggedUserKeycloakID")
	projetoID := c.Param("projetoId")

	input := votacao.CancelaVotacaoInput{
		LoggedInUserKeycloakID: loggedUserKeycloakID,
		ProjetoID:              projetoID,
	}

	if err := h.cancelaVotacaoUseCase.Execute(c.Request.Context(), input); err != nil {
		c.JSON(http.StatusForbidden, gin.H{"error": err.Error()})
		return
	}

	c.Status(http.StatusNoContent)
}
