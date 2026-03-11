package usuario

import "errors"

var ErrUserNotFound = errors.New("usuario not found")
var ErrUserNotAdmin = errors.New("usuario does not have admin permissions")

var ErrCredencialNotFound = errors.New("credencial not found")
