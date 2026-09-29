# Prompt — nova feature no front

```
Feature: <tela/componente/hook> no app <admin|vereador>

Endpoint(s) usados: <rota + resposta esperada>
Eventos SSE relevantes (se houver): <votacao_aberta | votacao_fechada | votacao_cancelada | voto_registrado>

Requisitos:
- TanStack Router (rota em src/routes) e TanStack Query (hook em src/hooks).
- Se o código servir aos dois apps, coloque em web/packages/shared e exporte no index.
- shadcn/ui + Tailwind, Biome (tabs, aspas simples), tipos em src/types ou shared/types.
- Invalidar/atualizar queries a partir de eventos SSE quando fizer sentido.
- Não duplique componentes que já existem em shared.

Entregue os arquivos completos e a lista de comandos que devo rodar (pnpm lint, pnpm build).
```