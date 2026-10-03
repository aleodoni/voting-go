// Package httperr traduz erros da aplicação em respostas HTTP.
package httperr

import (
	"errors"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"

	domainUsuario "github.com/aleodoni/voting-go/internal/domain/usuario"
	domainVotacao "github.com/aleodoni/voting-go/internal/domain/votacao"
)

// mensagemErroInterno é o texto devolvido ao cliente para erros inesperados.
// O erro real vai para o log, não para a resposta.
const mensagemErroInterno = "erro interno do servidor"

// regras associa cada erro conhecido ao status HTTP correspondente. A mensagem
// devolvida é a do próprio erro-alvo (não a de um erro que o embrulhe), pois ela
// é escrita para o usuário.
var regras = []struct {
	alvo   error
	status int
}{
	{domainUsuario.ErrUserNotFound, http.StatusNotFound},
	{domainUsuario.ErrCredencialNotFound, http.StatusNotFound},
	{domainVotacao.ErrReuniaoNotFound, http.StatusNotFound},
	{domainVotacao.ErrProjetoNotFound, http.StatusNotFound},
	{domainUsuario.ErrUserNotAdmin, http.StatusForbidden},
	{domainUsuario.ErrUserNotVoter, http.StatusForbidden},
	{domainUsuario.ErrUserNotActive, http.StatusForbidden},
}

// Respond escreve a resposta JSON {"error": "..."} correspondente ao erro:
// 404 para "não encontrado", 403 para "sem permissão" e 500 com mensagem
// genérica para qualquer outro erro, que é registrado no log.
func Respond(c *gin.Context, err error) {
	status, mensagem := traduz(err)

	if status == http.StatusInternalServerError {
		log.Printf("erro interno em %s %s: %v", c.Request.Method, c.FullPath(), err)
	}

	c.JSON(status, gin.H{"error": mensagem})
}

func traduz(err error) (int, string) {
	for _, r := range regras {
		if errors.Is(err, r.alvo) {
			return r.status, r.alvo.Error()
		}
	}

	return http.StatusInternalServerError, mensagemErroInterno
}
