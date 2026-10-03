# BACKLOG — dívida técnica

Features ficam em `docs/roadmap.md`. Aqui: correções, limpezas e riscos conhecidos.
Prioridade: **P0** segurança/quebra · **P1** importante · **P2** melhoria.

Itens resolvidos são **removidos** daqui — o commit que resolveu é o registro.

## Repositório e segurança


## Backend

- [ ] **P2** `cmd/cli/main/` com nível extra; padronizar para `cmd/cli/`
- [ ] **P2** `Dockerfile` da API não usava `--parseInternal` (Makefile usava); alinhado na reestruturação — validar Swagger da imagem
- [ ] **P2** Migrar os handlers de `votacao` (6) para `httperr.Respond`: respondem 403 a qualquer erro e devolvem `err.Error()` ao cliente. Pede estender as regras do `httperr` com os erros do domínio (votação já aberta, não aberta, já votou)
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
- [ ] **P2** `web/packages/shared/src/api-client.ts` chama `getKeycloak().login()` em 401 **e** 403. Um 403 legítimo (sem permissão) não é sessão expirada; avaliar limitar o login ao 401, depois que o backend parar de usar 403 para qualquer erro
- [ ] **P2** `ProjectDTO` em `web/apps/admin/src/types/meeting.ts` é idêntico a `ProjetoDTO` de `@voting/shared`; usar o do shared e remover a cópia
- [ ] **P2** Ao subir o Biome para uma versão com `css.parser.tailwindDirectives`, trocar o override de `noUnknownAtRules` em `web/biome.json` por essa opção (ADR 0006)

## CI/CD

- [ ] **P1** Documentar qual caminho de deploy é o oficial (compose/SSH × AWS/ECR) e desativar o que não for usado
- [ ] **P2** A versão do pnpm (10.33.0) está fixada em três lugares (`ci.yml` e os Dockerfiles de `admin` e `vereador`), e os Dockerfiles instalam com `--no-frozen-lockfile` enquanto o CI usa `--frozen-lockfile`. Fixar `packageManager` em `web/package.json` e alinhar os Dockerfiles
- [ ] **P2** `infra/Makefile` referencia `docker-compose.production.yml` e `docker-compose.orange.yml`, que não estão no repo; alinhar ou documentar

## Documentação

- [ ] **P2** Publicar `docs/` no GitHub Pages junto com o Swagger (`docs.yml`)
- [ ] **P2** Confirmar e preencher em `docs/domain.md` o significado dos status `F`/`V` e das opções de voto