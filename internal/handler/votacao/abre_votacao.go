// Package votacao contains the handler for opening a voting session.
package votacao

import (
	"net/http"

	"github.com/gin-gonic/gin"

	ucVotacao "github.com/aleodoni/voting-go/internal/application/votacao"
)

type AbreVotacaoHandler struct {
	abreVotacaoUseCase *ucVotacao.AbreVotacaoUseCase
}

func NewAbreVotacaoHandler(abreVotacaoUseCase *ucVotacao.AbreVotacaoUseCase) *AbreVotacaoHandler {
	return &AbreVotacaoHandler{abreVotacaoUseCase: abreVotacaoUseCase}
}

// Handle godoc
//
//	@Summary		Abre uma votação
//	@Description	Abre a sessão de votação para um projeto
//	@Tags			votação
//	@Param			projetoId	path	string	true	"ID do projeto"
//	@Success		204
//	@Failure		403	{object}	ErrorResponse
//	@Security		BearerAuth
//	@Router			/projetos/{projetoId}/votacao/abrir [post]
func (h *AbreVotacaoHandler) Handle(c *gin.Context) {
	loggedUserKeycloakID := c.GetString("loggedUserKeycloakID")
	projetoID := c.Param("projetoId")

	input := ucVotacao.AbreVotacaoInput{
		LoggedInUserKeycloakID: loggedUserKeycloakID,
		ProjetoID:              projetoID,
	}

	if err := h.abreVotacaoUseCase.Execute(c.Request.Context(), input); err != nil {
		c.JSON(http.StatusForbidden, gin.H{"error": err.Error()})
		return
	}

	c.Status(http.StatusNoContent)
}
