package votacao

import "errors"

var ErrReuniaoNotFound = errors.New("reuniao não encontrada")
var ErrProjetoNotFound = errors.New("projeto não encontrado")
var ErrVotacaoNaoCriada = errors.New("votação não criada")
var ErrVotacaoAlreadyExists = errors.New("votação já existe")
