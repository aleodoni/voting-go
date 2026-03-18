package votacao

import "github.com/aleodoni/voting-go/internal/domain/votacao"

type RegistraVotoRequest struct {
	Voto          votacao.OpcaoVoto             `json:"voto"          binding:"required" example:"F" enums:"F,R,C,V,A"`
	Restricao     *RegistraRestricaoRequest     `json:"restricao"`
	VotoContrario *RegistraVotoContrarioRequest `json:"votoContrario"`
}

type RegistraRestricaoRequest struct {
	Restricao string `json:"restricao" binding:"required" example:"Impedimento regimental"`
}

type RegistraVotoContrarioRequest struct {
	IDTexto   int    `json:"idTexto"   binding:"required" example:"1"`
	ParecerID string `json:"parecerId" binding:"required" example:"cls1par123"`
}

// --- Errors ---

type ErrorResponse struct {
	Error string `json:"error" example:"mensagem de erro"`
}
