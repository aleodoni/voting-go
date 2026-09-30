# ADR 0003 — SSE com event bus em memória

- **Status**: Aceito (com limitação conhecida; a autenticação por query string foi substituída pela ADR 0007)
- **Contexto**: apuração e painel em tempo real durante a sessão, com poucos clientes simultâneos.
- **Decisão**: eventos de domínio publicados num `event.Bus` em memória, entregues por SSE (`/api/v1/eventos`). Token via query string porque `EventSource` não envia headers.
- **Consequências**: simples e sem infraestrutura extra; o estado dos assinantes é por instância (não escala horizontalmente sem broker) e o token na URL pode aparecer em logs. Ver `BACKLOG.md` e `roadmap.md`.