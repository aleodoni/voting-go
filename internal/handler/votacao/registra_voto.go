// Package votacao contains the handler for registering a vote.
package votacao

import (
	"net/http"

	"github.com/gin-gonic/gin"

	ucVotacao "github.com/aleodoni/voting-go/internal/application/votacao"
	"github.com/aleodoni/voting-go/internal/domain/votacao"
)

type RegistraVotoRequest struct {
	Voto          votacao.OpcaoVoto             `json:"voto" binding:"required"`
	Restricao     *RegistraRestricaoRequest     `json:"restricao"`
	VotoContrario *RegistraVotoContrarioRequest `json:"votoContrario"`
}

type RegistraRestricaoRequest struct {
	Restricao string `json:"restricao" binding:"required"`
}

type RegistraVotoContrarioRequest struct {
	IDTexto   int    `json:"idTexto" binding:"required"`
	ParecerID string `json:"parecerId" binding:"required"`
}

type RegistraVotoHandler struct {
	registraVotoUseCase *ucVotacao.RegistraVotoUseCase
}

func NewRegistraVotoHandler(registraVotoUseCase *ucVotacao.RegistraVotoUseCase) *RegistraVotoHandler {
	return &RegistraVotoHandler{registraVotoUseCase: registraVotoUseCase}
}

func (h *RegistraVotoHandler) Handle(c *gin.Context) {
	loggedUserKeycloakID := c.GetString("loggedUserKeycloakID")
	votacaoID := c.Param("votacaoId")

	var req RegistraVotoRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var restricao *votacao.Restricao
	if req.Restricao != nil {
		restricao = &votacao.Restricao{
			Restricao: req.Restricao.Restricao,
		}
	}

	var votoContrario *votacao.VotoContrario
	if req.VotoContrario != nil {
		votoContrario = &votacao.VotoContrario{
			IDTexto:   req.VotoContrario.IDTexto,
			ParecerID: req.VotoContrario.ParecerID,
		}
	}

	input := ucVotacao.RegistraVotoInput{
		LoggedInUserKeycloakID: loggedUserKeycloakID,
		VotacaoID:              votacaoID,
		Voto:                   req.Voto,
		Restricao:              restricao,
		VotoContrario:          votoContrario,
	}

	if err := h.registraVotoUseCase.Execute(c.Request.Context(), input); err != nil {
		c.JSON(http.StatusForbidden, gin.H{"error": err.Error()})
		return
	}

	c.Status(http.StatusNoContent)
}
