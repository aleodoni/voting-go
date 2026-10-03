# Runbook — Sincronização SPL

## O que é

Copia reuniões, projetos e pareceres do banco SPL (via `postgres_fdw`) para o banco local. Cada execução grava um registro em `sincronia` (sucesso, mensagem de erro, contadores, início/fim).

## Como executar

| Forma | Como |
| ----- | ---- |
| Manual (admin) | `POST /api/v1/sincronia` (botão no admin); `GET /api/v1/sincronia` lista as 3 últimas |
| Job interno | `POST /internal/jobs/sincronia` com `Authorization: Bearer $JOBS_TOKEN` |
| Direto no banco | `CALL public.p_spl_daily_sync();` |

No `POST /api/v1/sincronia` a permissão é verificada **antes** de responder: um não-admin recebe 403 e nada é iniciado. A sincronização em si roda em segundo plano e a resposta é 202; o resultado aparece em `GET /api/v1/sincronia` e, se falhar, no log (`[SINCRONIA]`). Em staging a execução é ignorada.

Job auxiliar: `POST /internal/jobs/fecha_abertas` fecha votações que ficaram abertas.

## Setup do FDW

```bash
make fdw            # local
voting-cli fdw      # ambientes de deploy
voting-cli fdw-drop # remover
```

Requer `DB_SPL_HOST`, `DB_SPL_PORT`, `DB_SPL_NAME`, `DB_SPL_USER`, `DB_SPL_PASSWORD`.

## Diagnóstico

1. Última sincronia com `sucesso = false`? Ver `mensagem_erro`.
2. Conectividade da API/banco para o host SPL (porta, firewall, credenciais).
3. Servidor FDW existente (`server_spl`) e user mapping válidos.