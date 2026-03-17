package votacao

import (
	"encoding/json"
	"fmt"
	"io"

	"github.com/aleodoni/voting-go/internal/platform/event"
	"github.com/gin-gonic/gin"
)

type SSEHandler struct {
	bus *event.Bus
}

func NewSSEHandler(bus *event.Bus) *SSEHandler {
	return &SSEHandler{bus: bus}
}

func (h *SSEHandler) Handle(c *gin.Context) {
	c.Header("Content-Type", "text/event-stream")
	c.Header("Cache-Control", "no-cache")
	c.Header("Connection", "keep-alive")
	c.Header("Access-Control-Allow-Origin", "*")

	ch := h.bus.Subscribe()
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
