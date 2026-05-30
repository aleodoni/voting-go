package usuario

import (
	"net/http"

	"github.com/aleodoni/voting-go/internal/platform/event"
	"github.com/gin-gonic/gin"
)

type ConnectedUsersHandler struct {
	bus *event.Bus
}

func NewConnectedUsersHandler(bus *event.Bus) *ConnectedUsersHandler {
	return &ConnectedUsersHandler{bus: bus}
}

type ConnectedUserResponse struct {
	UserID   string `json:"user_id"`
	Username string `json:"username"`
	IsAdmin  bool   `json:"is_admin"`
}

// Handle godoc
//
//	@Summary		Retorna usuários conectados via SSE
//	@Description	Retorna a lista de usuários com conexão SSE ativa no momento
//	@Tags			usuários
//	@Produce		json
//	@Success		200	{array}		ConnectedUserResponse
//	@Failure		401	{object}	map[string]interface{}
//	@Security		BearerAuth
//	@Router			/usuarios-conectados [get]
func (h *ConnectedUsersHandler) Handle(c *gin.Context) {
	subscribers := h.bus.ConnectedUsers()

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
