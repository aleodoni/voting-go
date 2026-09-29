# Prompt — mudança de schema

```
Mudança: <descrição>

- Crie o par .up.sql/.down.sql com numeração sequencial (última em migrations/: <número>).
- Nunca edite migration já aplicada.
- O .down deve desfazer completamente o .up.
- Considere dados existentes (default/backfill) e impacto em funções/views que dependem da tabela
  (f_get_projetos_completo, f_get_voting_stats, v_reunioes_hoje, etc.).
- Liste as queries sqlc que precisam mudar e me diga para rodar `sqlc generate`.
- Diga como validar: make migrate, make migrate-down, make migrate.
```