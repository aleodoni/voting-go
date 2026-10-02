package usuario

import (
	"net/http"

	"github.com/gin-gonic/gin"

	ucCredencial "github.com/aleodoni/voting-go/internal/application/usuario"
	"github.com/aleodoni/voting-go/internal/handler/httperr"
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
//	@Failure		404		{object}	ErrorResponse
//	@Failure		500		{object}	ErrorResponse
//	@Security		BearerAuth
//	@Router			/usuarios/{id}/credencial [patch]
func (h *UpdateCredencialHandler) Handle(c *gin.Context) {
	var req UpdateCredencialRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, ErrorResponse{Error: "invalid request body"})
		return
	}

	input := ucCredencial.UpdateCredencialInput{
		AdminKeycloakID: c.GetString("loggedUserKeycloakID"),
		UsuarioID:       c.Param("id"),
		Ativo:           req.Ativo,
		PodeVotar:       req.PodeVotar,
		PodeAdministrar: req.PodeAdministrar,
	}

	cred, err := h.updateUseCase.Execute(c.Request.Context(), input)
	if err != nil {
		httperr.Respond(c, err)
		return
	}

	c.JSON(http.StatusOK, toCredencialResponse(cred))
}
