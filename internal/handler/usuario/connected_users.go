package usuario

import (
	"net/http"

	"github.com/gin-gonic/gin"

	ucUsuario "github.com/aleodoni/voting-go/internal/application/usuario"
	"github.com/aleodoni/voting-go/internal/handler/httperr"
)

type ConnectedUsersHandler struct {
	listConnectedUsersUseCase *ucUsuario.ListConnectedUsersUseCase
}

func NewConnectedUsersHandler(listConnectedUsersUseCase *ucUsuario.ListConnectedUsersUseCase) *ConnectedUsersHandler {
	return &ConnectedUsersHandler{listConnectedUsersUseCase: listConnectedUsersUseCase}
}

type ConnectedUserResponse struct {
	UserID   string `json:"user_id"`
	Username string `json:"username"`
	IsAdmin  bool   `json:"is_admin"`
}

// Handle godoc
//
//	@Summary		Retorna usuários conectados via SSE
//	@Description	Retorna a lista de usuários com conexão SSE ativa no momento (requer admin)
//	@Tags			usuários
//	@Produce		json
//	@Success		200	{array}		ConnectedUserResponse
//	@Failure		401	{object}	map[string]interface{}
//	@Failure		403	{object}	ErrorResponse
//	@Failure		404	{object}	ErrorResponse
//	@Failure		500	{object}	ErrorResponse
//	@Security		BearerAuth
//	@Router			/usuarios-conectados [get]
func (h *ConnectedUsersHandler) Handle(c *gin.Context) {
	subscribers, err := h.listConnectedUsersUseCase.Execute(c.Request.Context(), ucUsuario.ListConnectedUsersInput{
		LoggedInUserKeycloakID: c.GetString("loggedUserKeycloakID"),
	})
	if err != nil {
		httperr.Respond(c, err)
		return
	}

	users := make([]ConnectedUserResponse, 0, len(subscribers))
	for _, s := range subscribers {
		users = append(users, ConnectedUserResponse{
			UserID:   s.UserID,
			Username: s.Username,
			IsAdmin:  s.IsAdmin,
		})
	}

	c.JSON(http.StatusOK, users)
}
