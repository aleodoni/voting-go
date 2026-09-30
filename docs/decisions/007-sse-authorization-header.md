# ADR 0007 — SSE autenticado pelo header Authorization

- **Status**: Aceito (substitui a parte de autenticação da ADR 0003)
- **Contexto**: a ADR 0003 adotou o token JWT na query string de `/api/v1/eventos`, com a premissa de que `EventSource` não envia headers. O front nunca usou `EventSource`: o `useSSE` lê o stream com `fetch`, que envia `Authorization` normalmente. Com isso o JWT aparecia na URL e, portanto, em logs de proxy e de servidor.
- **Decisão**: `/api/v1/eventos` passa para o grupo protegido e é autenticado pelo `JWTMiddleware`, com `Authorization: Bearer <token>` como nas demais rotas. O middleware deixa de aceitar o token por query string em qualquer rota.
- **Consequências**: o JWT não aparece mais em URLs, e o handler do SSE deixa de validar token por conta própria. Quem consumir o stream precisa de um cliente que envie headers (como `fetch`); o `EventSource` nativo não serve. Tokens já gravados em logs antigos expiram com o tempo de vida do access token.
