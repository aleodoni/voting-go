package votacao

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	domainUsuario "github.com/aleodoni/voting-go/internal/domain/usuario"

	"github.com/aleodoni/voting-go/internal/middleware"
	"github.com/aleodoni/voting-go/internal/platform/event"
	"github.com/gin-gonic/gin"
)

type SSEHandler struct {
	bus           *event.Bus
	jwtMiddleware *middleware.JWTMiddleware
	usuarioRepo   domainUsuario.UsuarioRepository
}

func NewSSEHandler(bus *event.Bus, jwtMiddleware *middleware.JWTMiddleware, usuarioRepo domainUsuario.UsuarioRepository) *SSEHandler {
	return &SSEHandler{
		bus:           bus,
		jwtMiddleware: jwtMiddleware,
		usuarioRepo:   usuarioRepo,
	}
}

func (h *SSEHandler) Handle(c *gin.Context) {
	// Headers CORS explícitos para SSE
	origin := c.Request.Header.Get("Origin")
	if origin != "" {
		c.Header("Access-Control-Allow-Origin", origin)
		c.Header("Access-Control-Allow-Credentials", "true")
	}

	token := c.Query("token")
	if token == "" {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "missing token"})
		return
	}

	claims, err := h.jwtMiddleware.ValidateToken(token)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "invalid token"})
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

	c.Header("Content-Type", "text/event-stream")
	c.Header("Cache-Control", "no-cache")
	c.Header("Connection", "keep-alive")
	c.Header("X-Accel-Buffering", "no")
	c.Status(http.StatusOK)

	fmt.Fprintf(c.Writer, ": connected\n\n")
	c.Writer.Flush()

	heartbeat := time.NewTicker(25 * time.Second)
	defer heartbeat.Stop()
	ctx := c.Request.Context()

	for {
		select {
		case <-ctx.Done():
			return
		case e, ok := <-ch:
			if !ok {
				return
			}
			payload, _ := json.Marshal(e.Payload)
			fmt.Fprintf(c.Writer, "event: %s\ndata: %s\n\n", e.Type, payload)
			c.Writer.Flush()
		case <-heartbeat.C:
			fmt.Fprintf(c.Writer, ": ping\n\n")
			c.Writer.Flush()
		}
	}
}
