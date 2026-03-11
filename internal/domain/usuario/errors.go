package usuario

import "errors"

var ErrNotFound = errors.New("usuario not found")
var ErrNotAdmin = errors.New("usuario does not have admin permissions")
