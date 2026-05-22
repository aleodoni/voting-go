-- name: ExecuteDailySync :exec
CALL public.p_spl_daily_sync();

-- name: ListLastSincronias :many
SELECT
  id,
  iniciado_em,
  finalizado_em,
  sucesso,
  mensagem_erro,
  reunioes_sincronizadas,
  projetos_sincronizados,
  pareceres_sincronizados
FROM sincronia
ORDER BY iniciado_em DESC
LIMIT 3;

-- name: GetLastSincronia :one
SELECT *
FROM sincronia
ORDER BY iniciado_em DESC
LIMIT 1;