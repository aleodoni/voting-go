// Package usuario contains the handler for searching users.
package usuario

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"

	ucUsuario "github.com/aleodoni/voting-go/internal/application/usuario"
)

type PesquisaUsuariosHandler struct {
	listUsuariosUseCase *ucUsuario.ListUsuariosUseCase
}

func NewPesquisaUsuariosHandler(listUsuariosUseCase *ucUsuario.ListUsuariosUseCase) *PesquisaUsuariosHandler {
	return &PesquisaUsuariosHandler{listUsuariosUseCase: listUsuariosUseCase}
}

func (h *PesquisaUsuariosHandler) Handle(c *gin.Context) {
	loggedUserKeycloakID := c.GetString("loggedUserKeycloakID")

	page, err := strconv.Atoi(c.DefaultQuery("page", "1"))
	if err != nil || page < 1 {
		page = 1
	}

	limit, err := strconv.Atoi(c.DefaultQuery("limit", "20"))
	if err != nil || limit < 1 || limit > 100 {
		limit = 20
	}

	input := ucUsuario.ListUsuariosInput{
		LoggedInUserKeycloakID: loggedUserKeycloakID,
		Search:                 c.Query("search"),
		Page:                   page,
		Limit:                  limit,
	}

	output, err := h.listUsuariosUseCase.Execute(c.Request.Context(), input)
	if err != nil {
		println("Error in handler:", err)
		c.JSON(403, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, output)
}
