package middleware

import (
	"fmt"
	"log"
	"time"

	"github.com/MicahParks/keyfunc/v2"
)

// Parâmetros da carga inicial do JWKS. Com 6 tentativas e espera de 1s, 2s, 4s,
// 8s e 10s entre elas, a API aguarda cerca de 25s pelo Keycloak antes de desistir.
const (
	jwksTentativas    = 6
	jwksEsperaInicial = time.Second
	jwksEsperaMaxima  = 10 * time.Second
)

// jwksGetter obtém o JWKS. Em produção é keyfunc.Get; os testes injetam um
// substituto para não depender de rede.
type jwksGetter func(jwksURL string, options keyfunc.Options) (*keyfunc.JWKS, error)

// carregaJWKS tenta obter o JWKS até `tentativas` vezes, dobrando a espera a cada
// falha até `esperaMaxima`. Devolve erro se todas as tentativas falharem.
func carregaJWKS(
	get jwksGetter,
	jwksURL string,
	options keyfunc.Options,
	tentativas int,
	espera, esperaMaxima time.Duration,
) (*keyfunc.JWKS, error) {
	tentativas = max(tentativas, 1)

	var ultimoErr error

	for tentativa := 1; tentativa <= tentativas; tentativa++ {
		jwks, err := get(jwksURL, options)
		if err == nil {
			return jwks, nil
		}

		ultimoErr = err

		if tentativa == tentativas {
			break
		}

		log.Printf("JWKS indisponível (tentativa %d/%d): %v; nova tentativa em %s", tentativa, tentativas, err, espera)
		time.Sleep(espera)

		espera = min(espera*2, esperaMaxima)
	}

	return nil, fmt.Errorf("falha ao obter o JWKS de %s após %d tentativas: %w", jwksURL, tentativas, ultimoErr)
}
