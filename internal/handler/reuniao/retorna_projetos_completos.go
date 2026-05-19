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

// Handle godoc
//
//	@Summary		Retorna projetos de uma reunião
//	@Description	Retorna a lista completa de projetos de uma reunião (requer admin)
//	@Tags			reuniões
//	@Produce		json
//	@Param			reuniaoId	path		string	true	"ID da reunião"
//	@Success		200			{array}		ProjetoResponse
//	@Failure		403			{object}	ErrorResponse
//	@Security		BearerAuth
//	@Router			/reunioes/{reuniaoId}/projetos [get]
func (h *RetornaProjetosCompletosHandler) Handle(c *gin.Context) {
	loggedUserKeycloakID := c.GetString("loggedUserKeycloakID")
	reuniaoID := c.Param("reuniaoId")

	input := ucReuniao.RetornaProjetosCompletosInput{
		LoggedInUserKeycloakID: loggedUserKeycloakID,
		ReuniaoID:              reuniaoID,
	}

	projetos, err := h.retornaProjetosCompletosUseCase.Execute(c.Request.Context(), input)
	if err != nil {
		c.JSON(http.StatusForbidden, ErrorResponse{Error: err.Error()})
		return
	}

	c.JSON(http.StatusOK, toProjetosResponse(projetos))
}
