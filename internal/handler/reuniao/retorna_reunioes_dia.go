// Package reuniao contains the handler for returning the meetings of the day.
package reuniao

import (
	"net/http"

	"github.com/gin-gonic/gin"

	ucReuniao "github.com/aleodoni/voting-go/internal/application/votacao"
)

type RetornaReunioesDiaHandler struct {
	retornaReunioesDiaUseCase *ucReuniao.RetornaReunioesDiaUseCase
}

func NewRetornaReunioesDiaHandler(retornaReunioesDiaUseCase *ucReuniao.RetornaReunioesDiaUseCase) *RetornaReunioesDiaHandler {
	return &RetornaReunioesDiaHandler{retornaReunioesDiaUseCase: retornaReunioesDiaUseCase}
}

func (h *RetornaReunioesDiaHandler) Handle(c *gin.Context) {
	loggedUserKeycloakID := c.GetString("loggedUserKeycloakID")

	input := ucReuniao.RetornaReunioesDiaInput{
		LoggedInUserKeycloakID: loggedUserKeycloakID,
	}

	reunioes, err := h.retornaReunioesDiaUseCase.Execute(c.Request.Context(), input)

	if err != nil {
		println("Error in handler:", err)
		c.JSON(403, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, reunioes)
}
