package sincronia

import (
	"net/http"

	"github.com/gin-gonic/gin"

	ucSincronia "github.com/aleodoni/voting-go/internal/application/sincronia"
)

type ExecutaSincroniaHandler struct {
	executaSincroniaUseCase *ucSincronia.ExecutaSincroniaUseCase
}

func NewExecutaSincroniaHandler(executaSincroniaUseCase *ucSincronia.ExecutaSincroniaUseCase) *ExecutaSincroniaHandler {
	return &ExecutaSincroniaHandler{executaSincroniaUseCase: executaSincroniaUseCase}
}

// Handle godoc
//
//	@Summary		Executa sincronização
//	@Description	Executa a sincronização de dados (requer admin)
//	@Tags			sincronia
//	@Produce		json
//	@Success		200		{object}	ExecutaSincroniaResponse
//	@Failure		403		{object}	ErrorResponse
//	@Security		BearerAuth
//	@Router			/usuarios [get]
func (h *ExecutaSincroniaHandler) Handle(c *gin.Context) {
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
