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

// Handle godoc
//
//	@Summary		Cancela uma votação
//	@Description	Cancela a sessão de votação de um projeto
//	@Tags			votação
//	@Param			projetoId	path	string	true	"ID do projeto"
//	@Success		204
//	@Failure		403	{object}	ErrorResponse
//	@Security		BearerAuth
//	@Router			/projetos/{projetoId}/votacao [delete]
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
