# BACKLOG — dívida técnica

Features ficam em `docs/roadmap.md`. Aqui: correções, limpezas e riscos conhecidos.
Prioridade: **P0** segurança/quebra · **P1** importante · **P2** melhoria.

Itens resolvidos são **removidos** daqui — o commit que resolveu é o registro.

## Repositório e segurança


## Backend

- [ ] **P2** `Dockerfile` da API não usava `--parseInternal` (Makefile usava); alinhado na reestruturação — validar Swagger da imagem
- [ ] **P2** `domain/votacao/errors.go` tem erros sem uso em lugar nenhum: `ErrVotacaoNaoCriada`, `ErrVotacaoAlreadyExists` e `ErrVotacaoFechada` (esta repete a mensagem de `ErrVotacaoNaoAberta`); remover
- [ ] **P2** `handler/votacao/sse.go` responde 401 quando o usuário logado não existe no banco (o front trata 401 com `login()`) e tem respostas próprias; avaliar usar `httperr.Respond`
- [ ] **P2** Decidir se `GET /reunioes/{id}/relatorio` deve exigir admin: o caso de uso não verifica permissão e qualquer usuário autenticado gera o PDF com os votos nominais (só o app admin chama a rota hoje)
- [ ] **P2** `PATCH /usuarios/{id}/credencial` não é chamado pelo front e duplica `PUT /usuarios/fantasia-credenciais` (as duas chamam `UpdateDisplayNamePermissions`); avaliar remover
- [ ] **P2** Agrupar `fdw/`, `spl/`, `seeds/`, `migrations/` em `db/` (opcional; ajustar Makefile, CLI e Dockerfile)

## Testes

- [ ] **P2** Ligar no `make test-api` os testes k6 que alteram estado (`sincronia`, `atualiza-fantasia-credenciais`, `reunioes-dia`), com a base de dev recarregada antes de cada rodada
- [ ] **P2** Testes de handler/router (hoje só casos de uso e o middleware JWT têm testes)

## Frontend

- [ ] **P2** Unificar `MeetingSelect/ProjectCard`, `StatusBadge`, `ProjectVotingPanel` com os equivalentes de `VotingPanel/` no admin
- [ ] **P2** Remover `web/packages/shared/pnpm-lock.yaml` (lockfile único na raiz do workspace) e o campo `workspaces` de `web/package.json` (redundante com `pnpm-workspace.yaml`)
- [ ] **P2** Padronizar tipagem de `useIsProjectVoting.ts` (SSE + TanStack Query)
- [ ] **P2** Os apps não mostram a mensagem do servidor: o toast de voto é genérico ("Erro ao registrar voto") e o `VotingCard` exibe a mensagem do axios ("Request failed with status code 422"). Criar um helper em `@voting/shared` que extraia `error.response.data.error` e usá-lo nesses pontos (o 422 traz, por exemplo, "usuário já votou nesta votação")
- [ ] **P2** O app admin não tem tela de "acesso negado": um usuário sem permissão de admin recebe 403 em todas as chamadas e vê telas de erro vazias (antes o `login()` em 403 o prendia num loop de redirecionamentos). Avaliar um guard de rota que explique a falta de permissão
- [ ] **P2** `ProjectDTO` em `web/apps/admin/src/types/meeting.ts` é idêntico a `ProjetoDTO` de `@voting/shared`; usar o do shared e remover a cópia
- [ ] **P2** Ao subir o Biome para uma versão com `css.parser.tailwindDirectives`, trocar o override de `noUnknownAtRules` em `web/biome.json` por essa opção (ADR 0006)

## CI/CD

- [ ] **P1** Documentar qual caminho de deploy é o oficial (compose/SSH × AWS/ECR) e desativar o que não for usado
- [ ] **P2** `infra/Makefile` referencia `docker-compose.production.yml` e `docker-compose.orange.yml`, que não estão no repo; alinhar ou documentar

## Documentação

- [ ] **P2** Publicar `docs/` no GitHub Pages junto com o Swagger (`docs.yml`)
- [ ] **P2** Confirmar e preencher em `docs/domain.md` o significado dos status `F`/`V` e das opções de voto