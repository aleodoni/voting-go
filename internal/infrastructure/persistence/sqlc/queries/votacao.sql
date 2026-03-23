-- name: UpsertVotacao :exec
INSERT INTO votacao (id, projeto_id, status, created_at, updated_at)
VALUES ($1, $2, $3, NOW(), NOW())
ON CONFLICT (id) DO UPDATE
    SET projeto_id = EXCLUDED.projeto_id,
        status     = EXCLUDED.status,
        updated_at = NOW();

-- name: DeleteVotacao :exec
DELETE FROM votacao
WHERE id = $1;

-- name: FindVotacaoByID :one
SELECT id, projeto_id, status, created_at, updated_at
FROM votacao
WHERE id = $1;

-- name: FindVotacaoAberta :one
SELECT id, projeto_id, status, created_at, updated_at
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
SELECT f_save_vote($1, $2, $3, $4, $5, $6);

-- name: GetProjectOpenVoting :one
SELECT public.f_get_project_open_voting();

-- name: GetVotingStats :one
SELECT public.f_get_voting_stats($1);