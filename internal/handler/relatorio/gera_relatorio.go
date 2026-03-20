package relatorio

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"

	ucRelatorio "github.com/aleodoni/voting-go/internal/application/relatorio"
)

type GeraRelatorioReuniaoHandler struct {
	geraRelatorioReuniaoUseCase *ucRelatorio.GeraRelatorioReuniaoUseCase
}

func NewGeraRelatorioReuniaoHandler(geraRelatorioReuniaoUseCase *ucRelatorio.GeraRelatorioReuniaoUseCase) *GeraRelatorioReuniaoHandler {
	return &GeraRelatorioReuniaoHandler{geraRelatorioReuniaoUseCase: geraRelatorioReuniaoUseCase}
}

// Handle godoc
//
//	@Summary		Gera relatório de uma reunião
//	@Description	Gera o relatório PDF com os projetos e votações de uma reunião
//	@Tags			reuniões
//	@Produce		application/pdf
//	@Param			reuniaoId	path	string	true	"ID da reunião"
//	@Success		200
//	@Failure		500	{object}	ErrorResponse
//	@Security		BearerAuth
//	@Router			/reunioes/{reuniaoId}/relatorio [get]
func (h *GeraRelatorioReuniaoHandler) Handle(c *gin.Context) {
	reuniaoID := c.Param("reuniaoId")

	pdf, err := h.geraRelatorioReuniaoUseCase.Execute(c.Request.Context(), ucRelatorio.GeraRelatorioReuniaoInput{
		ReuniaoID: reuniaoID,
	})
	if err != nil {
		c.JSON(http.StatusInternalServerError, ErrorResponse{Error: err.Error()})
		return
	}

	c.Header("Content-Disposition", fmt.Sprintf("inline; filename=reuniao-%s.pdf", reuniaoID))
	c.Data(http.StatusOK, "application/pdf", pdf)
}
