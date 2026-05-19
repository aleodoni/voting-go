// Package usuario provides HTTP handlers for user management.
package usuario

import (
	"net/http"

	"github.com/gin-gonic/gin"

	ucUsuario "github.com/aleodoni/voting-go/internal/application/usuario"
)

type AtualizaFantasiaHandler struct {
	updateDisplayNameUseCase *ucUsuario.UpdateDisplayNameUseCase
}

func NewAtualizaFantasiaHandler(updateDisplayNameUseCase *ucUsuario.UpdateDisplayNameUseCase) *AtualizaFantasiaHandler {
	return &AtualizaFantasiaHandler{updateDisplayNameUseCase: updateDisplayNameUseCase}
}

// Handle godoc
//
//	@Summary		Atualiza nome fantasia
//	@Description	Atualiza o nome fantasia de um usuário
//	@Tags			usuários
//	@Accept			json
//	@Param			body	body	AtualizaFantasiaRequest	true	"Dados a atualizar"
//	@Success		204
//	@Failure		400	{object}	ErrorResponse
//	@Failure		403	{object}	ErrorResponse
//	@Security		BearerAuth
//	@Router			/usuarios/fantasia [put]
func (h *AtualizaFantasiaHandler) Handle(c *gin.Context) {
	loggedUserKeycloakID := c.GetString("loggedUserKeycloakID")

	var req AtualizaFantasiaRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, ErrorResponse{Error: err.Error()})
		return
	}

	input := ucUsuario.UpdateDisplayNameInput{
		LoggedInUserKeycloakID: loggedUserKeycloakID,
		UserID:                 req.UserID,
		DisplayName:            req.DisplayName,
	}

	if err := h.updateDisplayNameUseCase.Execute(c.Request.Context(), input); err != nil {
		c.JSON(http.StatusForbidden, ErrorResponse{Error: err.Error()})
		return
	}

	c.Status(http.StatusNoContent)
}
