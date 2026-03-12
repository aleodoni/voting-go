// Package domain provides shared domain errors.
package domain

import "errors"

var ErrForbidden = errors.New("acesso negado")
var ErrUnauthorized = errors.New("não autorizado")
