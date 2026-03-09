// Package usuario provides HTTP handlers for user management.
package usuario

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"

	ucUsuario "github.com/aleodoni/voting-go/internal/application/usuario"
)

type MeHandler struct {
	ensureUseCase *ucUsuario.EnsureUsuarioUseCase
}

func NewMeHandler(ensureUseCase *ucUsuario.EnsureUsuarioUseCase) *MeHandler {
	return &MeHandler{ensureUseCase: ensureUseCase}
}

func (h *MeHandler) Handle(c *gin.Context) {
	claims, ok := c.MustGet("claims").(jwt.MapClaims)
	if !ok {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "invalid claims"})
		return
	}

	input := ucUsuario.EnsureUsuarioInput{
		KeycloakID: claimString(claims, "sub"),
		Username:   claimString(claims, "preferred_username"),
		Email:      claimString(claims, "email"),
		Nome:       claimString(claims, "name"),
	}

	usuario, err := h.ensureUseCase.Execute(c.Request.Context(), input)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to ensure user"})
		return
	}

	c.JSON(http.StatusOK, usuario)
}

func claimString(claims jwt.MapClaims, key string) string {
	val, _ := claims[key].(string)
	return val
}
