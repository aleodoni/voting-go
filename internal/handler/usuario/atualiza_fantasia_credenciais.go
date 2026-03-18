// Package usuario provides HTTP handlers for user management.
package usuario

import (
	"net/http"

	"github.com/gin-gonic/gin"

	ucUsuario "github.com/aleodoni/voting-go/internal/application/usuario"
)

type AtualizaFantasiaCredenciaisHandler struct {
	updateDisplayNamePermissionsUseCase *ucUsuario.UpdateDisplayNamePermissionsUseCase
}

func NewAtualizaFantasiaCredenciaisHandler(updateDisplayNamePermissionsUseCase *ucUsuario.UpdateDisplayNamePermissionsUseCase) *AtualizaFantasiaCredenciaisHandler {
	return &AtualizaFantasiaCredenciaisHandler{updateDisplayNamePermissionsUseCase: updateDisplayNamePermissionsUseCase}
}

// Handle godoc
//
//	@Summary		Atualiza nome fantasia e permissões
//	@Description	Atualiza o nome fantasia e as permissões de credencial de um usuário (requer admin)
//	@Tags			usuários
//	@Accept			json
//	@Param			body	body	AtualizaFantasiaCredenciaisRequest	true	"Dados a atualizar"
//	@Success		204
//	@Failure		400	{object}	ErrorResponse
//	@Failure		403	{object}	ErrorResponse
//	@Security		BearerAuth
//	@Router			/usuarios/fantasia-credenciais [put]
func (h *AtualizaFantasiaCredenciaisHandler) Handle(c *gin.Context) {
	loggedUserKeycloakID := c.GetString("loggedUserKeycloakID")

	var req AtualizaFantasiaCredenciaisRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, ErrorResponse{Error: err.Error()})
		return
	}

	input := ucUsuario.UpdateDisplayNamePermissionsInput{
		LoggedInUserKeycloakID: loggedUserKeycloakID,
		UserID:                 req.UserID,
		DisplayName:            req.DisplayName,
		IsActive:               req.IsActive,
		CanAdmin:               req.CanAdmin,
		CanVote:                req.CanVote,
	}

	if err := h.updateDisplayNamePermissionsUseCase.Execute(c.Request.Context(), input); err != nil {
		c.JSON(http.StatusForbidden, ErrorResponse{Error: err.Error()})
		return
	}

	c.Status(http.StatusNoContent)
}
