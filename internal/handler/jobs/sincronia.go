package jobs

import (
	"context"
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"

	ucSincroniaJob "github.com/aleodoni/voting-go/internal/application/jobs"
)

type ExecutaSincroniaJobHandler struct {
	executaSincroniaUseCase *ucSincroniaJob.ExecutaSincroniaJobUseCase
	appEnv                  string
}

func NewExecutaSincroniaJobHandler(executaSincroniaUseCase *ucSincroniaJob.ExecutaSincroniaJobUseCase, appEnv string) *ExecutaSincroniaJobHandler {
	return &ExecutaSincroniaJobHandler{executaSincroniaUseCase: executaSincroniaUseCase, appEnv: appEnv}
}

// Handle godoc
//
//	@Summary		Executa job sincronia
//	@Description	Executa a sincronização de dados (requer admin)
//	@Tags			jobs
//	@Produce		json
//	@Success		202
//	@Failure		401
//	@Security		BearerAuth
//	@Router			/internal/jobs/sincronia [post]
func (h *ExecutaSincroniaJobHandler) Handle(c *gin.Context) {
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

		_, err := h.executaSincroniaUseCase.Execute(ctx)
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
		"message":  "sincronia iniciada",
		"executed": true,
	})
}
