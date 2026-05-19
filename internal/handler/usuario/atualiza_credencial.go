package usuario

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"

	ucCredencial "github.com/aleodoni/voting-go/internal/application/usuario"
	domain "github.com/aleodoni/voting-go/internal/domain"
	jwtutil "github.com/aleodoni/voting-go/internal/platform/jwt"
)

type UpdateCredencialHandler struct {
	updateUseCase *ucCredencial.UpdateCredencialUseCase
}

func NewUpdateCredencialHandler(updateUseCase *ucCredencial.UpdateCredencialUseCase) *UpdateCredencialHandler {
	return &UpdateCredencialHandler{updateUseCase: updateUseCase}
}

// Handle godoc
//
//	@Summary		Atualiza credencial de um usuário
//	@Description	Atualiza as permissões de um usuário (requer admin)
//	@Tags			usuários
//	@Accept			json
//	@Produce		json
//	@Param			id		path		string					true	"ID do usuário"
//	@Param			body	body		UpdateCredencialRequest	true	"Dados da credencial"
//	@Success		200		{object}	CredencialResponse
//	@Failure		400		{object}	ErrorResponse
//	@Failure		401		{object}	ErrorResponse
//	@Failure		403		{object}	ErrorResponse
//	@Failure		500		{object}	ErrorResponse
//	@Security		BearerAuth
//	@Router			/usuarios/{id}/credencial [patch]
func (h *UpdateCredencialHandler) Handle(c *gin.Context) {
	claims, ok := c.MustGet("claims").(jwt.MapClaims)
	if !ok {
		c.JSON(http.StatusUnauthorized, ErrorResponse{Error: "invalid claims"})
		return
	}

	var req UpdateCredencialRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, ErrorResponse{Error: "invalid request body"})
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
			c.JSON(http.StatusForbidden, ErrorResponse{Error: "acesso negado"})
		default:
			c.JSON(http.StatusInternalServerError, ErrorResponse{Error: "erro ao atualizar credencial"})
		}
		return
	}

	c.JSON(http.StatusOK, toCredencialResponse(cred))
}
