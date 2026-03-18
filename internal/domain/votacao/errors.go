package votacao

import "errors"

var ErrReuniaoNotFound = errors.New("reuniao não encontrada")
var ErrProjetoNotFound = errors.New("projeto não encontrado")
var ErrVotacaoNaoCriada = errors.New("votação não criada")
var ErrVotacaoAlreadyExists = errors.New("votação já existe")
var ErrVotacaoNaoAberta = errors.New("votação não está aberta")
var ErrVotacaoNaoFechada = errors.New("votação não está fechada")
var ErrVotacaoNaoEncontrada = errors.New("votação não encontrada")
var ErrVotacaoFechada = errors.New("votação não está aberta")
var ErrVotacaoAberta = errors.New("já existe uma votação aberta")
var ErrUsuarioJaVotou = errors.New("usuário já votou nesta votação")
