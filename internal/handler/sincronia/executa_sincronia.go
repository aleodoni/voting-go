package sincronia

import (
	"context"
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"

	ucSincronia "github.com/aleodoni/voting-go/internal/application/sincronia"
	"github.com/aleodoni/voting-go/internal/handler/httperr"
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
//	@Description	Inicia a sincronização de dados de forma assíncrona (requer admin). Em ambiente staging a execução é ignorada.
//	@Tags			sincronia
//	@Produce		json
//	@Success		202	{object}	map[string]interface{}	"Sincronia iniciada"
//	@Failure		403	{object}	ErrorResponse
//	@Failure		404	{object}	ErrorResponse
//	@Failure		500	{object}	ErrorResponse
//	@Security		BearerAuth
//	@Router			/sincronia [post]
func (h *ExecutaSincroniaHandler) Handle(c *gin.Context) {
	loggedUserKeycloakID := c.GetString("loggedUserKeycloakID")

	input := ucSincronia.ExecutaSincroniaInput{
		LoggedInUserKeycloakID: loggedUserKeycloakID,
	}

	// Recusa antes de responder 202: a execução é assíncrona, e um erro de
	// permissão dentro da goroutine só apareceria no log.
	if err := h.executaSincroniaUseCase.Autorizar(c.Request.Context(), input); err != nil {
		httperr.Respond(c, err)
		return
	}

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
