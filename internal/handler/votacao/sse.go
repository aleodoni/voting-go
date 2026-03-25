package votacao

import (
	"encoding/json"
	"fmt"
	"io"

	domainUsuario "github.com/aleodoni/voting-go/internal/domain/usuario"
	"github.com/aleodoni/voting-go/internal/middleware"
	"github.com/aleodoni/voting-go/internal/platform/event"
	"github.com/gin-gonic/gin"
)

type SSEHandler struct {
	bus           *event.Bus
	usuarioRepo   domainUsuario.UsuarioRepository
	jwtMiddleware *middleware.JWTMiddleware
}

func NewSSEHandler(bus *event.Bus, usuarioRepo domainUsuario.UsuarioRepository, jwtMiddleware *middleware.JWTMiddleware) *SSEHandler {
	return &SSEHandler{bus: bus, usuarioRepo: usuarioRepo, jwtMiddleware: jwtMiddleware}
}

func (h *SSEHandler) Handle(c *gin.Context) {
	c.Header("Content-Type", "text/event-stream")
	c.Header("Cache-Control", "no-cache")
	c.Header("Connection", "keep-alive")
	c.Header("Access-Control-Allow-Origin", "*")

	token := c.Query("token")
	if token == "" {
		c.AbortWithStatusJSON(401, gin.H{"error": "token não informado"})
		return
	}

	claims, err := h.jwtMiddleware.ValidateToken(token)
	if err != nil {
		c.AbortWithStatusJSON(401, gin.H{"error": "token inválido"})
		return
	}

	keycloakID := claims["sub"].(string)
	username := claims["preferred_username"].(string)

	u, err := h.usuarioRepo.FindByKeycloakID(c.Request.Context(), keycloakID)
	if err != nil {
		c.AbortWithStatusJSON(401, gin.H{"error": "usuário não encontrado"})
		return
	}

	isAdmin := u.Credencial != nil && u.Credencial.IsAdmin()
	ch := h.bus.Subscribe(u.ID, username, isAdmin)
	defer h.bus.Unsubscribe(ch)

	c.Stream(func(w io.Writer) bool {
		select {
		case e, ok := <-ch:
			if !ok {
				return false
			}
			payload, err := json.Marshal(e.Payload)
			if err != nil {
				return true
			}
			fmt.Fprintf(w, "event: %s\ndata: %s\n\n", e.Type, payload)
			return true
		case <-c.Request.Context().Done():
			return false
		}
	})
}
