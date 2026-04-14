-- name: UpsertVotacao :exec
INSERT INTO votacao (id, projeto_id, status, created_at, updated_at)
VALUES (
    sqlc.arg(id),
    sqlc.arg(projeto_id),
    sqlc.arg(status),
    NOW(),
    NOW()
)
ON CONFLICT (id) DO UPDATE
    SET projeto_id = EXCLUDED.projeto_id,
        status     = EXCLUDED.status,
        updated_at = NOW();

-- name: DeleteVotacao :exec
DELETE FROM votacao
WHERE id = $1;

-- name: FindVotacaoByID :one
SELECT id, projeto_id, status::text as status, created_at, updated_at
FROM votacao
WHERE id = $1;

-- name: FindVotacaoAberta :one
SELECT 
    id, 
    projeto_id, 
    status::text as status, 
    created_at, 
    updated_at
FROM votacao
WHERE status = $1
LIMIT 1;

-- name: UsuarioJaVotou :one
SELECT EXISTS (
    SELECT 1 FROM voto
    WHERE usuario_id = $1
    AND votacao_id = $2
) AS already_voted;

-- name: SaveVoto :exec
SELECT f_save_vote(
    sqlc.arg(p_voto_id),
    sqlc.arg(p_usuario_id),
    sqlc.arg(p_votacao_id),
    sqlc.arg(p_voto)::text,
    sqlc.arg(p_restricao),
    sqlc.arg(p_voto_contrario)
);

-- name: GetProjectOpenVoting :one
SELECT public.f_get_project_open_voting();

-- name: GetVotingStats :one
SELECT public.f_get_voting_stats($1);