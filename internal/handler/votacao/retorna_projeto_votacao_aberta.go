package votacao

import (
	"net/http"

	"github.com/gin-gonic/gin"

	ucVotacao "github.com/aleodoni/voting-go/internal/application/votacao"
	mappers "github.com/aleodoni/voting-go/internal/handler/reuniao"
)

type RetornaProjetoVotacaoAbertaHandler struct {
	retornaProjetoVotacaoAbertaUseCase *ucVotacao.RetornaVotacaoAbertaUseCase
}

func NewRetornaProjetoVotacaoAbertaHandler(retornaProjetoVotacaoAbertaUseCase *ucVotacao.RetornaVotacaoAbertaUseCase) *RetornaProjetoVotacaoAbertaHandler {
	return &RetornaProjetoVotacaoAbertaHandler{retornaProjetoVotacaoAbertaUseCase: retornaProjetoVotacaoAbertaUseCase}
}

// Handle godoc
//
//	@Summary		Retorna projeto com votação aberta
//	@Description	Retorna a lista completa de um projeto que tenha a votação aberta (requer usuario logado)
//	@Tags			votação
//	@Produce		json
//	@Success		200	{object}	reuniao.ProjetoResponse
//	@Failure		403	{object}	ErrorResponse
//	@Security		BearerAuth
//	@Router			/votacao/aberta [get]
func (h *RetornaProjetoVotacaoAbertaHandler) Handle(c *gin.Context) {
	loggedUserKeycloakID := c.GetString("loggedUserKeycloakID")

	input := ucVotacao.RetornaVotacaoAbertaInput{
		LoggedInUserKeycloakID: loggedUserKeycloakID,
	}

	projeto, err := h.retornaProjetoVotacaoAbertaUseCase.Execute(c.Request.Context(), input)
	if err != nil {
		c.JSON(http.StatusForbidden, ErrorResponse{Error: err.Error()})
		return
	}

	c.JSON(http.StatusOK, mappers.ToProjetoResponse(projeto))
}
