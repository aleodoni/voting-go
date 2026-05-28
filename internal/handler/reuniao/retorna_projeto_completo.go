// Package reuniao contains the handler for returning the meetings of the day.
package reuniao

import (
	"net/http"

	"github.com/gin-gonic/gin"

	ucReuniao "github.com/aleodoni/voting-go/internal/application/votacao"
)

type RetornaProjetoCompletoHandler struct {
	retornaProjetoCompletoUseCase *ucReuniao.RetornaProjetoCompletoUseCase
}

func NewRetornaProjetoCompletoHandler(retornaProjetoCompletoUseCase *ucReuniao.RetornaProjetoCompletoUseCase) *RetornaProjetoCompletoHandler {
	return &RetornaProjetoCompletoHandler{retornaProjetoCompletoUseCase: retornaProjetoCompletoUseCase}
}

// Handle godoc
//
//	@Summary		Retorna projeto de uma reunião
//	@Description	Retorna a lista completa de projetos de uma reunião (requer admin)
//	@Tags			reuniões
//	@Produce		json
//	@Param			projetoId	path		string	true	"ID do projeto"
//	@Success		200			{object}	ProjetoResponse
//	@Failure		403			{object}	ErrorResponse
//	@Security		BearerAuth
//	@Router			/projetos/:projetoId [get]
func (h *RetornaProjetoCompletoHandler) Handle(c *gin.Context) {
	loggedUserKeycloakID := c.GetString("loggedUserKeycloakID")
	projetoID := c.Param("projetoId")

	input := ucReuniao.RetornaProjetoCompletoInput{
		LoggedInUserKeycloakID: loggedUserKeycloakID,
		ProjetoID:              projetoID,
	}

	projeto, err := h.retornaProjetoCompletoUseCase.Execute(c.Request.Context(), input)
	if err != nil {
		c.JSON(http.StatusForbidden, ErrorResponse{Error: err.Error()})
		return
	}

	c.JSON(http.StatusOK, ToProjetoResponse(projeto))
}
