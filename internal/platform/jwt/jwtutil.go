// Package jwtutil provides utilities for extracting claims from JWT tokens.
package jwtutil

import "github.com/golang-jwt/jwt/v5"

func ClaimString(claims jwt.MapClaims, key string) string {
	val, _ := claims[key].(string)
	return val
}
