# ADR 0002 — sqlc e lógica pesada em funções SQL

- **Status**: Aceito
- **Contexto**: consultas agregadas (projetos completos, estatísticas, usuários paginados) são naturalmente SQL.
- **Decisão**: queries em `sqlc/queries/*.sql` com código gerado; consultas complexas encapsuladas em funções/views PostgreSQL versionadas em migrations.
- **Consequências**: tipagem forte sem ORM; parte da lógica vive no banco (exige migrations cuidadosas e testes de banco). Nunca editar `sqlc/generated/`.