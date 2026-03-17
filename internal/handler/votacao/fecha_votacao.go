package votacao

import (
	"net/http"

	"github.com/aleodoni/voting-go/internal/application/votacao"
	"github.com/gin-gonic/gin"
)

type FechaVotacaoHandler struct {
	fechaVotacaoUseCase *votacao.FechaVotacaoUseCase
}

func NewFechaVotacaoHandler(fechaVotacaoUseCase *votacao.FechaVotacaoUseCase) *FechaVotacaoHandler {
	return &FechaVotacaoHandler{fechaVotacaoUseCase: fechaVotacaoUseCase}
}

func (h *FechaVotacaoHandler) Handle(c *gin.Context) {
	loggedUserKeycloakID := c.GetString("loggedUserKeycloakID")
	projetoID := c.Param("projetoId")

	input := votacao.FechaVotacaoInput{
		LoggedInUserKeycloakID: loggedUserKeycloakID,
		ProjetoID:              projetoID,
	}

	if err := h.fechaVotacaoUseCase.Execute(c.Request.Context(), input); err != nil {
		c.JSON(http.StatusForbidden, gin.H{"error": err.Error()})
		return
	}

	c.Status(http.StatusNoContent)
}
