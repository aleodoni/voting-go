package jobs

import (
	"log"
	"net/http"

	sincronia "github.com/aleodoni/voting-go/internal/handler/sincronia"
	"github.com/gin-gonic/gin"

	ucSincronia "github.com/aleodoni/voting-go/internal/application/sincronia"
)

type ExecutaSincroniaJobHandler struct {
	executaSincroniaUseCase *ucSincronia.ExecutaSincroniaUseCase
	appEnv                  string
}

func NewExecutaSincroniaJobHandler(executaSincroniaUseCase *ucSincronia.ExecutaSincroniaUseCase, appEnv string) *ExecutaSincroniaJobHandler {
	return &ExecutaSincroniaJobHandler{executaSincroniaUseCase: executaSincroniaUseCase, appEnv: appEnv}
}

// Handle godoc
//
//	@Summary		Executa job sincronia
//	@Description	Executa a sincronização de dados (requer admin)
//	@Tags			sincronia
//	@Produce		json
//	@Success		200		{object}	SincroniaResponse
//	@Failure		403		{object}	ErrorResponse
//	@Security		BearerAuth
//	@Router			/usuarios [get]
func (h *ExecutaSincroniaJobHandler) Handle(c *gin.Context) {
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
		c.JSON(http.StatusForbidden, sincronia.ErrorResponse{Error: err.Error()})
		return
	}

	c.JSON(http.StatusOK, sincronia.ToSincroniaResponse(output))
}
