package usuario

import "errors"

var ErrUserNotFound = errors.New("usuário não encontrado")
var ErrUserNotAdmin = errors.New("usuário não tem permissões de admin")
var ErrUserNotVoter = errors.New("usuário não tem permissões de votante")
var ErrUserNotActive = errors.New("usuário não está ativo")

var ErrCredencialNotFound = errors.New("credencial não encontrada")
