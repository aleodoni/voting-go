# BACKLOG — dívida técnica

Features ficam em `docs/roadmap.md`. Aqui: correções, limpezas e riscos conhecidos.
Prioridade: **P0** segurança/quebra · **P1** importante · **P2** melhoria.

## Repositório e segurança

- [ ] **P0** Remover do índice `infra/.env.aws`, `.env.deploy`, `.env.staging` (`git rm --cached`), manter apenas `*.example`; revisar histórico e rotacionar credenciais se algum valor real já foi commitado
- [ ] **P1** Remover `node_modules/.pnpm-workspace-state-v1.json` e `web/apps/admin/tsconfig.tsbuildinfo` do índice
- [ ] **P1** Corrigir `paths` do `build_api.yml` (`build-api.yml` → `build_api.yml`); revisar `build_admin.yml` e `build_vereador.yml` pelo mesmo problema
- [ ] **P2** Verificar se `infra/keycloak/realm-import/*.json` contêm segredos de client ou senhas de usuário e, se houver, tirar do versionamento

## Backend

- [ ] **P1** `middleware/auth.go`: asserções de tipo sem verificação (`claims["sub"].(string)`, `claims["preferred_username"].(string)`) podem causar panic com token malformado
- [ ] **P1** Token JWT na query string do SSE aparece em logs de proxy/servidor; avaliar ticket curto de uso único para o SSE
- [ ] **P1** `NewJWTMiddleware` faz `panic` se o JWKS estiver indisponível na subida; avaliar retry/erro controlado
- [ ] **P2** `Votacao.Abrir` levanta `VotacaoAbertaEvent` sem `domainEvent` (OccurredAt zerado), diferente de Fechar/Cancelar — verificar e padronizar
- [ ] **P2** `domain/votacao/status_votacao.go` contém `VotingStats`, não o status; renomear/mover
- [ ] **P2** Typos em nomes de arquivo: `usuaio_repository.go`, `unity_of_work_sqlc.go`
- [ ] **P2** `cmd/cli/main/` com nível extra; padronizar para `cmd/cli/`
- [ ] **P2** `Dockerfile` da API não usava `--parseInternal` (Makefile usava); alinhado na reestruturação — validar Swagger da imagem
- [ ] **P2** Handlers `usuario/*` misturam regra de "quem pode" (admin) entre handler e caso de uso; padronizar no caso de uso
- [ ] **P2** Agrupar `fdw/`, `spl/`, `seeds/`, `migrations/` em `db/` (opcional; ajustar Makefile, CLI e Dockerfile)

## Testes

- [ ] **P1** Separar `tests/api/betha-*.test.js` (serviço externo, exige segredos) em `tests/integrations/`
- [ ] **P1** `test-api` no Makefile só roda `health`; cobrir os demais `tests/api/*.test.js` ou documentar por que não
- [ ] **P2** Testes de handler/router (hoje só casos de uso têm testes)

## Frontend

- [ ] **P1** Mover para `@voting/shared` o que está duplicado entre admin e vereador: `useIsProjectVoting`, `useUser`, `FormUserInfo`, `LoggedUserCard`
- [ ] **P1** Remover `react-router-dom` dos apps se só o TanStack Router é usado
- [ ] **P1** Escolher um linter: `lint` dos apps roda `eslint`, mas a raiz `web/` configura Biome
- [ ] **P2** Unificar `MeetingSelect/ProjectCard`, `StatusBadge`, `ProjectVotingPanel` com os equivalentes de `VotingPanel/` no admin
- [ ] **P2** Remover `web/packages/shared/pnpm-lock.yaml` (lockfile único na raiz do workspace) e o campo `workspaces` de `web/package.json` (redundante com `pnpm-workspace.yaml`)
- [ ] **P2** Padronizar tipagem de `useIsProjectVoting.ts` (SSE + TanStack Query)

## CI/CD

- [ ] **P1** `ci.yml` roda só em `develop`; incluir PRs e pushes em `main`
- [ ] **P1** Adicionar job de lint + build do front no CI
- [ ] **P1** Documentar qual caminho de deploy é o oficial (compose/SSH × AWS/ECR) e desativar o que não for usado
- [ ] **P2** `infra/Makefile` referencia `docker-compose.production.yml` e `docker-compose.orange.yml`, que não estão no repo; alinhar ou documentar

## Documentação

- [x] README atualizado (front, CLI, jobs, SSE, infra, testes)
- [ ] **P2** Publicar `docs/` no GitHub Pages junto com o Swagger (`docs.yml`)