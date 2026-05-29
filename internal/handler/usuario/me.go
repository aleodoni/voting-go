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
	adminGroup    string
}

func NewMeHandler(ensureUseCase *ucUsuario.EnsureUsuarioUseCase, adminGroup string) *MeHandler {
	return &MeHandler{ensureUseCase: ensureUseCase, adminGroup: adminGroup}
}

func extractGroups(claims jwt.MapClaims) []string {
	rawGroups, ok := claims["groups"].([]interface{})
	if !ok {
		return []string{}
	}

	groups := make([]string, 0, len(rawGroups))

	for _, g := range rawGroups {
		if str, ok := g.(string); ok {
			groups = append(groups, str)
		}
	}

	return groups
}

func isAdmin(groups []string, adminGroup string) bool {
	for _, g := range groups {
		if g == adminGroup {
			return true
		}
	}

	return false
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

	groups := extractGroups(claims)

	admin := isAdmin(groups, h.adminGroup)

	input := ucUsuario.EnsureUsuarioInput{
		KeycloakID: jwtutil.ClaimString(claims, "sub"),
		Username:   jwtutil.ClaimString(claims, "preferred_username"),
		Email:      jwtutil.ClaimString(claims, "email"),
		Nome:       jwtutil.ClaimString(claims, "name"),
		IsAdmin:    admin,
	}

	usuario, err := h.ensureUseCase.Execute(c.Request.Context(), input)

	if err != nil {
		c.JSON(http.StatusInternalServerError, ErrorResponse{Error: "failed to ensure user"})
		return
	}

	c.JSON(http.StatusOK, toMeResponse(usuario))
}
