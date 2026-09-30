package middleware

import (
	"errors"
	"testing"
	"time"

	"github.com/MicahParks/keyfunc/v2"
)

var errJWKSIndisponivel = errors.New("keycloak indisponível")

// getterFalso falha as primeiras `falhas` chamadas e depois devolve `jwks`.
// O contador `chamadas` permite verificar quantas tentativas foram feitas.
func getterFalso(falhas int, jwks *keyfunc.JWKS, chamadas *int) jwksGetter {
	return func(string, keyfunc.Options) (*keyfunc.JWKS, error) {
		*chamadas++

		if *chamadas <= falhas {
			return nil, errJWKSIndisponivel
		}

		return jwks, nil
	}
}

func TestCarregaJWKS_SucessoNaPrimeiraTentativa(t *testing.T) {
	esperado := &keyfunc.JWKS{}
	chamadas := 0

	jwks, err := carregaJWKS(getterFalso(0, esperado, &chamadas), "http://jwks.test", keyfunc.Options{}, 3, time.Millisecond, 2*time.Millisecond)
	if err != nil {
		t.Fatalf("esperava nil, obteve: %v", err)
	}

	if jwks != esperado {
		t.Fatal("esperava o JWKS devolvido pelo getter")
	}

	if chamadas != 1 {
		t.Fatalf("esperava 1 chamada, obteve %d", chamadas)
	}
}

func TestCarregaJWKS_RecuperaAposFalhasTransitorias(t *testing.T) {
	esperado := &keyfunc.JWKS{}
	chamadas := 0

	jwks, err := carregaJWKS(getterFalso(2, esperado, &chamadas), "http://jwks.test", keyfunc.Options{}, 5, time.Millisecond, 2*time.Millisecond)
	if err != nil {
		t.Fatalf("esperava nil, obteve: %v", err)
	}

	if jwks != esperado {
		t.Fatal("esperava o JWKS devolvido pelo getter")
	}

	if chamadas != 3 {
		t.Fatalf("esperava 3 chamadas (2 falhas e 1 sucesso), obteve %d", chamadas)
	}
}

func TestCarregaJWKS_EsgotaTentativasRetornaErro(t *testing.T) {
	chamadas := 0

	jwks, err := carregaJWKS(getterFalso(100, nil, &chamadas), "http://jwks.test", keyfunc.Options{}, 4, time.Millisecond, 2*time.Millisecond)
	if err == nil {
		t.Fatal("esperava erro, obteve nil")
	}

	if !errors.Is(err, errJWKSIndisponivel) {
		t.Fatalf("esperava que o erro original fosse preservado, obteve: %v", err)
	}

	if jwks != nil {
		t.Fatal("esperava jwks nil")
	}

	if chamadas != 4 {
		t.Fatalf("esperava 4 chamadas, obteve %d", chamadas)
	}
}

func TestCarregaJWKS_TentativasInvalidasFazemUmaTentativa(t *testing.T) {
	chamadas := 0

	_, err := carregaJWKS(getterFalso(100, nil, &chamadas), "http://jwks.test", keyfunc.Options{}, 0, time.Millisecond, 2*time.Millisecond)
	if !errors.Is(err, errJWKSIndisponivel) {
		t.Fatalf("esperava o erro original, obteve: %v", err)
	}

	if chamadas != 1 {
		t.Fatalf("esperava 1 chamada, obteve %d", chamadas)
	}
}
