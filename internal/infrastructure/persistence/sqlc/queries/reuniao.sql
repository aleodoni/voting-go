-- name: FindReuniaoByID :one
SELECT id,
       con_id,
       rec_id,
       pac_id,
       con_desc,
       con_sigla,
       rec_tipo_reuniao,
       rec_numero,
       rec_data,
       created_at,
       updated_at
FROM reuniao
WHERE id = $1;

-- name: GetReunioesDia :many
SELECT id,
       con_id,
       rec_id,
       pac_id,
       con_desc,
       con_sigla,
       rec_tipo_reuniao,
       rec_numero,
       rec_data,
       created_at,
       updated_at
FROM v_reunioes_hoje;

-- name: GetProjetosCompleto :one
SELECT public.f_get_projetos_completo($1)::text AS result;

-- name: GetProjetoCompleto :one
SELECT public.f_get_projeto_completo($1)::text AS result;