// Package credencial provides HTTP handlers for credencial management.
package credencial

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"

	ucCredencial "github.com/aleodoni/voting-go/internal/application/credencial"
	domain "github.com/aleodoni/voting-go/internal/domain"
	jwtutil "github.com/aleodoni/voting-go/internal/platform/jwt"
)

type UpdateCredencialHandler struct {
	updateUseCase *ucCredencial.UpdateCredencialUseCase
}

func NewUpdateCredencialHandler(updateUseCase *ucCredencial.UpdateCredencialUseCase) *UpdateCredencialHandler {
	return &UpdateCredencialHandler{updateUseCase: updateUseCase}
}

type updateCredencialRequest struct {
	Ativo           bool `json:"ativo"`
	PodeVotar       bool `json:"pode_votar"`
	PodeAdministrar bool `json:"pode_administrar"`
}

func (h *UpdateCredencialHandler) Handle(c *gin.Context) {
	claims, ok := c.MustGet("claims").(jwt.MapClaims)
	if !ok {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "invalid claims"})
		return
	}

	var req updateCredencialRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid request body"})
		return
	}

	input := ucCredencial.UpdateCredencialInput{
		AdminKeycloakID: jwtutil.ClaimString(claims, "sub"),
		UsuarioID:       c.Param("id"),
		Ativo:           req.Ativo,
		PodeVotar:       req.PodeVotar,
		PodeAdministrar: req.PodeAdministrar,
	}

	cred, err := h.updateUseCase.Execute(c.Request.Context(), input)
	if err != nil {
		switch err {
		case domain.ErrForbidden:
			c.JSON(http.StatusForbidden, gin.H{"error": "acesso negado"})
		default:
			c.JSON(http.StatusInternalServerError, gin.H{"error": "erro ao atualizar credencial"})
		}
		return
	}

	c.JSON(http.StatusOK, cred)
}
