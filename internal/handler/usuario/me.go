// Package usuario provides HTTP handlers for user management.
package usuario

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"

	ucUsuario "github.com/aleodoni/voting-go/internal/application/usuario"
	jwtutil "github.com/aleodoni/voting-go/internal/platform/jwt"
)

type MeHandler struct {
	ensureUseCase *ucUsuario.EnsureUsuarioUseCase
}

func NewMeHandler(ensureUseCase *ucUsuario.EnsureUsuarioUseCase) *MeHandler {
	return &MeHandler{ensureUseCase: ensureUseCase}
}

// Handle godoc
//
//	@Summary		Retorna o usuário autenticado
//	@Description	Garante que o usuário do token JWT existe na base e retorna seus dados
//	@Tags			usuários
//	@Produce		json
//	@Success		200	{object}	MeResponse
//	@Failure		401	{object}	ErrorResponse
//	@Failure		500	{object}	ErrorResponse
//	@Security		BearerAuth
//	@Router			/me [get]
func (h *MeHandler) Handle(c *gin.Context) {
	claims, ok := c.MustGet("claims").(jwt.MapClaims)
	if !ok {
		c.JSON(http.StatusUnauthorized, ErrorResponse{Error: "invalid claims"})
		return
	}

	input := ucUsuario.EnsureUsuarioInput{
		KeycloakID: jwtutil.ClaimString(claims, "sub"),
		Username:   jwtutil.ClaimString(claims, "preferred_username"),
		Email:      jwtutil.ClaimString(claims, "email"),
		Nome:       jwtutil.ClaimString(claims, "name"),
	}

	usuario, err := h.ensureUseCase.Execute(c.Request.Context(), input)
	if err != nil {
		c.JSON(http.StatusInternalServerError, ErrorResponse{Error: "failed to ensure user"})
		return
	}

	c.JSON(http.StatusOK, toMeResponse(usuario))
}
