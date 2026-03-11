// Package usuario provides HTTP handlers for user management.
package usuario

import (
	"github.com/gin-gonic/gin"

	ucUsuario "github.com/aleodoni/voting-go/internal/application/usuario"
)

type AtualizaFantasiaCredenciaisHandler struct {
	updateDisplayNamePermissionsUseCase *ucUsuario.UpdateDisplayNamePermissionsUseCase
}

func NewAtualizaFantasiaCredenciaisHandler(updateDisplayNamePermissionsUseCase *ucUsuario.UpdateDisplayNamePermissionsUseCase) *AtualizaFantasiaCredenciaisHandler {
	return &AtualizaFantasiaCredenciaisHandler{updateDisplayNamePermissionsUseCase: updateDisplayNamePermissionsUseCase}
}

func (h *AtualizaFantasiaCredenciaisHandler) Handle(c *gin.Context) {
	loggedUserKeycloakID := c.GetString("loggedUserKeycloakID")

	var req AtualizaFantasiaCredenciaisRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
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

	err := h.updateDisplayNamePermissionsUseCase.Execute(
		c.Request.Context(),
		input,
	)

	if err != nil {
		println("Error in handler:", err)
		c.JSON(403, gin.H{"error": err.Error()})
		return
	}

	c.Status(204)
}
