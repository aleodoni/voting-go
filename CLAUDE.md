# CLAUDE.md — contexto do projeto Voting

Sistema de votação legislativa: API Go + dois apps React/Vite. Leia este arquivo no início de toda sessão; detalhes em `docs/`.

## Stack

- **API**: Go, Gin, PostgreSQL (pgx), sqlc, Keycloak (JWT via JWKS), swaggo, go-pdf/fpdf
- **Front** (`web/`): monorepo pnpm — `apps/admin`, `apps/vereador`, `packages/shared` (`@voting/shared`); React, Vite, TanStack Router/Query, react-hook-form + zod, shadcn/ui + Tailwind, Biome
- **Infra**: Docker Compose, Keycloak 24 (realm import), GitHub Actions, GHCR/ECR
- **Banco**: migrations em `migrations/` (golang-migrate); dados de reuniões/projetos/pareceres vêm do banco SPL via `postgres_fdw` (`fdw/`)

## Arquitetura (resumo)

`handler → application (casos de uso) → domain ← infrastructure`. Regras de negócio no domínio e nos casos de uso; handlers só fazem HTTP; repositórios implementados com sqlc em `internal/infrastructure/persistence`. Composição em `internal/bootstrap`. Eventos de domínio → `platform/event.Bus` → SSE (`/api/v1/eventos`). Detalhes: `docs/architecture.md`.

## Convenções

- Idioma do domínio e do código de negócio: **português** (votacao, reuniao, projeto, parecer, credencial). Comentários/doc de pacotes em português.
- Casos de uso: um arquivo por caso de uso (`<verbo>_<objeto>.go`), tipo `XxxUseCase`, `NewXxxUseCase`, `Execute(ctx, input)`, com `Input` próprio. Verificação de admin via `shared.VerificarAdmin`.
- Cada caso de uso tem `_test.go` usando fakes de `internal/test/fakes`.
- Handlers: um arquivo por endpoint, com anotações swaggo; request/response/mapper no pacote do handler.
- SQL novo: `internal/infrastructure/persistence/sqlc/queries/*.sql` + `sqlc generate`. Nunca editar `sqlc/generated/`.
- Migrations: sempre par `.up.sql`/`.down.sql`, numeração sequencial (`make migrate-create name=...`). Nunca editar migration já aplicada.
- Commits: conventional commits com corpo detalhado. Um commit combinado quando a mudança cruza back e front.
- Front: Biome (tabs, aspas simples). Código usado por mais de um app vai para `@voting/shared`.

## Comandos úteis

```bash
make bootstrap       # migrate + seed + fdw
make run | dev       # API (dev = air)
make swagger         # gera ./swagger (não versionado; NÃO usa docs/)
make test            # testes Go
make dev-web         # front (admin :5173, vereador :5174)
make build-web | lint-web
```

## Regras de trabalho com o Claude

- Não rode build/test/compile: o Alexandre roda localmente e cola o resultado.
- Entregar arquivos completos para arquivos pequenos/médios; diff apenas para arquivos muito grandes.
- Ele aplica as mudanças e faz o push; não assuma acesso de escrita ao repo.
- Trabalho iterativo e metódico; confirme premissas quando o pedido for ambíguo.
- Não versionar segredos. `infra/.env.*` reais são ignorados; use os `.example`.

## Onde ficam as coisas

| Assunto | Arquivo |
| ------- | ------- |
| Roadmap (features) | `docs/roadmap.md` |
| Dívida técnica | `BACKLOG.md` |
| Arquitetura e fluxos | `docs/architecture.md` |
| Glossário e regras de negócio | `docs/domain.md` |
| Decisões (ADRs) | `docs/decisions/` |
| Deploy e operação | `docs/runbooks/` |
| Prompts de sessão | `docs/prompts/` |