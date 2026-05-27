package sincronia

import (
	"context"
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"

	ucSincronia "github.com/aleodoni/voting-go/internal/application/sincronia"
)

type ExecutaSincroniaHandler struct {
	executaSincroniaUseCase *ucSincronia.ExecutaSincroniaUseCase
	appEnv                  string
}

func NewExecutaSincroniaHandler(executaSincroniaUseCase *ucSincronia.ExecutaSincroniaUseCase, appEnv string) *ExecutaSincroniaHandler {
	return &ExecutaSincroniaHandler{executaSincroniaUseCase: executaSincroniaUseCase, appEnv: appEnv}
}

// Handle godoc
//
//	@Summary		Executa sincronização
//	@Description	Executa a sincronização de dados (requer admin)
//	@Tags			sincronia
//	@Produce		json
//	@Success		200	{object}	SincroniaResponse
//	@Failure		403	{object}	ErrorResponse
//	@Security		BearerAuth
//	@Router			/sincronia [post]
func (h *ExecutaSincroniaHandler) Handle(c *gin.Context) {
	if h.appEnv == "staging" {
		log.Printf(
			"Sincronia ignorada em ambiente %s",
			h.appEnv,
		)

		c.JSON(http.StatusAccepted, gin.H{
			"message":  "Sincronia ignorada fora de produção/development",
			"executed": false,
		})

		return
	}

	loggedUserKeycloakID := c.GetString("loggedUserKeycloakID")

	input := ucSincronia.ExecutaSincroniaInput{
		LoggedInUserKeycloakID: loggedUserKeycloakID,
	}

	// EXECUTA ASYNC
	go func() {
		defer func() {
			if r := recover(); r != nil {
				log.Printf(
					"[SINCRONIA] panic recovered: %v",
					r,
				)
			}
		}()

		start := time.Now()

		// contexto independente da request HTTP
		ctx, cancel := context.WithTimeout(
			context.Background(),
			30*time.Minute,
		)
		defer cancel()

		_, err := h.executaSincroniaUseCase.Execute(ctx, input)
		if err != nil {
			log.Printf(
				"[SINCRONIA] erro ao executar: %v",
				err,
			)

			return
		}

		log.Printf(
			"[SINCRONIA] finalizada em %s",
			time.Since(start),
		)
	}()

	// RESPONDE IMEDIATAMENTE
	c.JSON(http.StatusAccepted, gin.H{
		"message":  "Sincronia iniciada",
		"executed": true,
	})
}
