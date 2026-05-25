package sincronia

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"

	ucSincronia "github.com/aleodoni/voting-go/internal/application/sincronia"
)

type ExecutaSincroniaHandler struct {
	executaSincroniaUseCase *ucSincronia.ExecutaSincroniaUseCase
	appEnv                  string
}

func NewExecutaSincroniaHandler(executaSincroniaUseCase *ucSincronia.ExecutaSincroniaUseCase, appEnv string) *ExecutaSincroniaHandler {
	return &ExecutaSincroniaHandler{executaSincroniaUseCase: executaSincroniaUseCase, appEnv: appEnv}
}

// Handle godoc
//
//	@Summary		Executa sincronização
//	@Description	Executa a sincronização de dados (requer admin)
//	@Tags			sincronia
//	@Produce		json
//	@Success		200		{object}	SincroniaResponse
//	@Failure		403		{object}	ErrorResponse
//	@Security		BearerAuth
//	@Router			/usuarios [get]
func (h *ExecutaSincroniaHandler) Handle(c *gin.Context) {
	if h.appEnv != "production" {
		log.Printf("Sincronia executada em ambiente %s, retornando 204 No Content", h.appEnv)
		c.Status(http.StatusNoContent)
		return
	}

	loggedUserKeycloakID := c.GetString("loggedUserKeycloakID")

	input := ucSincronia.ExecutaSincroniaInput{
		LoggedInUserKeycloakID: loggedUserKeycloakID,
	}

	output, err := h.executaSincroniaUseCase.Execute(c.Request.Context(), input)
	if err != nil {
		c.JSON(http.StatusForbidden, ErrorResponse{Error: err.Error()})
		return
	}

	c.JSON(http.StatusOK, toSincroniaResponse(output))
}
