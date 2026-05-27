package sincronia

import (
	"net/http"

	"github.com/gin-gonic/gin"

	ucSincronia "github.com/aleodoni/voting-go/internal/application/sincronia"
)

type RetornaUltimasSincroniasHandler struct {
	retornaUltimasSincroniasUseCase *ucSincronia.RetornaSincroniasUseCase
}

func NewRetornaUltimasSincroniasHandler(retornaUltimasSincroniasUseCase *ucSincronia.RetornaSincroniasUseCase) *RetornaUltimasSincroniasHandler {
	return &RetornaUltimasSincroniasHandler{retornaUltimasSincroniasUseCase: retornaUltimasSincroniasUseCase}
}

// Handle godoc
//
//	@Summary		Retorna últimas 3 sincronias
//	@Description	Retorna as últimas 3 sincronias executadas (requer admin)
//	@Tags			sincronia
//	@Produce		json
//	@Success		200	{object}	ListUltimasSincroniasResponse
//	@Failure		403	{object}	ErrorResponse
//	@Security		BearerAuth
//	@Router			/usuarios [get]
func (h *RetornaUltimasSincroniasHandler) Handle(c *gin.Context) {
	loggedUserKeycloakID := c.GetString("loggedUserKeycloakID")

	input := ucSincronia.RetornaSincroniasInput{
		LoggedInUserKeycloakID: loggedUserKeycloakID,
	}

	output, err := h.retornaUltimasSincroniasUseCase.Execute(c.Request.Context(), input)
	if err != nil {
		c.JSON(http.StatusForbidden, ErrorResponse{Error: err.Error()})
		return
	}

	c.JSON(http.StatusOK, ToListUltimasSincroniasResponse(output))
}
