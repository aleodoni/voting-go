# ADR 0001 — Clean Architecture com DDD leve

- **Status**: Aceito
- **Contexto**: regras de votação (quem pode abrir/votar, unicidade de voto, eventos) precisam ser testáveis sem HTTP nem banco.
- **Decisão**: camadas `handler → application → domain ← infrastructure`; agregados e eventos via biblioteca `go-ddd`; um caso de uso por arquivo; fakes de repositório em `internal/test/fakes`.
- **Consequências**: mais arquivos e boilerplate; em troca, casos de uso testáveis isoladamente e infraestrutura substituível.