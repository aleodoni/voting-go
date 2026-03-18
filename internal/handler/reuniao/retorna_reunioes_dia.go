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

// Handle godoc
//
//	@Summary		Retorna reuniões do dia
//	@Description	Retorna a lista de reuniões agendadas para o dia atual (requer admin)
//	@Tags			reuniões
//	@Produce		json
//	@Success		200	{array}		ReuniaoResponse
//	@Failure		403	{object}	ErrorResponse
//	@Security		BearerAuth
//	@Router			/reunioes-dia [get]
func (h *RetornaReunioesDiaHandler) Handle(c *gin.Context) {
	loggedUserKeycloakID := c.GetString("loggedUserKeycloakID")

	input := ucReuniao.RetornaReunioesDiaInput{
		LoggedInUserKeycloakID: loggedUserKeycloakID,
	}

	reunioes, err := h.retornaReunioesDiaUseCase.Execute(c.Request.Context(), input)
	if err != nil {
		c.JSON(http.StatusForbidden, ErrorResponse{Error: err.Error()})
		return
	}

	c.JSON(http.StatusOK, toReunioesDiaResponse(reunioes))
}
