// Package jwtutil provides utilities for extracting claims from JWT tokens.
package jwtutil

import "github.com/golang-jwt/jwt/v5"

// ClaimString retorna o valor textual da claim, ou "" se ela estiver ausente
// ou não for uma string.
func ClaimString(claims jwt.MapClaims, key string) string {
	val, _ := claims[key].(string)
	return val
}

// Identity extrai de um token já validado a identidade que a API exige: o ID
// do usuário no Keycloak ("sub") e o nome de usuário ("preferred_username").
// ok é false se qualquer uma das duas claims estiver ausente, vazia ou não for
// string; nesse caso o token não deve ser aceito.
func Identity(claims jwt.MapClaims) (keycloakID, username string, ok bool) {
	keycloakID = ClaimString(claims, "sub")
	username = ClaimString(claims, "preferred_username")

	return keycloakID, username, keycloakID != "" && username != ""
}
