// Package usuario contains the handler for searching users.
package usuario

import (
	"net/http"

	"github.com/gin-gonic/gin"

	ucUsuario "github.com/aleodoni/voting-go/internal/application/usuario"
)

type RetornaUsuarioHandler struct {
	retornaUsuarioUseCase *ucUsuario.RetornaUsuarioUseCase
}

func NewRetornaUsuarioHandler(retornaUsuarioUseCase *ucUsuario.RetornaUsuarioUseCase) *RetornaUsuarioHandler {
	return &RetornaUsuarioHandler{retornaUsuarioUseCase: retornaUsuarioUseCase}
}

// Handle godoc
//
//	@Summary		Retorna usuário
//	@Description	Retorna dados do usuário (requer admin)
//	@Tags			usuários
//	@Produce		json
//	@Param			usuarioId	path		string	true	"ID do usuário"
//	@Success		200			{object}	UsuarioResponse
//	@Failure		403			{object}	ErrorResponse
//	@Security		BearerAuth
//	@Router			/usuarios/{usuarioId} [get]
func (h *RetornaUsuarioHandler) Handle(c *gin.Context) {
	loggedUserKeycloakID := c.GetString("loggedUserKeycloakID")

	input := ucUsuario.RetornaUsuarioInput{
		LoggedInUserKeycloakID: loggedUserKeycloakID,
		UsuarioID:              c.Param("usuarioId"),
	}

	output, err := h.retornaUsuarioUseCase.Execute(c.Request.Context(), input)
	if err != nil {
		c.JSON(http.StatusForbidden, ErrorResponse{Error: err.Error()})
		return
	}

	c.JSON(http.StatusOK, toUsuarioResponse(output))
}
