// Package reuniao contains the handler for returning the meetings of the day.
package reuniao

import (
	"net/http"

	"github.com/gin-gonic/gin"

	ucReuniao "github.com/aleodoni/voting-go/internal/application/votacao"
)

type RetornaProjetosCompletosHandler struct {
	retornaProjetosCompletosUseCase *ucReuniao.RetornaProjetosCompletosUseCase
}

func NewRetornaProjetosCompletosHandler(retornaProjetosCompletosUseCase *ucReuniao.RetornaProjetosCompletosUseCase) *RetornaProjetosCompletosHandler {
	return &RetornaProjetosCompletosHandler{retornaProjetosCompletosUseCase: retornaProjetosCompletosUseCase}
}

func (h *RetornaProjetosCompletosHandler) Handle(c *gin.Context) {
	loggedUserKeycloakID := c.GetString("loggedUserKeycloakID")
	reuniaoID := c.Param("reuniaoId")

	input := ucReuniao.RetornaProjetosCompletosInput{
		LoggedInUserKeycloakID: loggedUserKeycloakID,
		ReuniaoID:              reuniaoID,
	}

	projetos, err := h.retornaProjetosCompletosUseCase.Execute(c.Request.Context(), input)

	if err != nil {
		println("Error in handler:", err)
		c.JSON(403, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, projetos)
}
