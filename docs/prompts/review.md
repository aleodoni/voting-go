# Prompt — revisão

```
Revise <arquivos/PR/pasta> considerando:
- Aderência às camadas (handler sem regra de negócio; domain sem dependências externas).
- Autorização: admin/pode votar verificados no caso de uso.
- Concorrência e consistência (uma votação por projeto, um voto por usuário).
- Segurança: token, logs, CORS, segredos.
- Testes: cobertura de casos de uso e fakes.
- Front: duplicação entre apps, uso de @voting/shared, consistência SSE/Query.

Responda por severidade (crítico / importante / sugestão), com arquivo e trecho.
Não altere código; proponha as correções e, para as que eu aprovar, entregue os arquivos completos.
```