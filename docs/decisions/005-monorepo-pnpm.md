# ADR 0005 — Monorepo pnpm para os frontends

- **Status**: Aceito
- **Contexto**: dois apps (admin e vereador) compartilham autenticação, cliente HTTP, SSE e UI.
- **Decisão**: workspace pnpm em `web/` com `apps/admin`, `apps/vereador` e `packages/shared` (`@voting/shared`, consumido direto do código-fonte TypeScript).
- **Consequências**: reuso simples; exige disciplina para mover código duplicado para o pacote compartilhado (ver `BACKLOG.md`).