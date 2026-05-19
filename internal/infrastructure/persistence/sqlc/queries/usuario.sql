-- ============================================
-- USUARIO
-- ============================================

-- name: FindByKeycloakID :one
SELECT
    u.id,
    u.keycloak_id,
    u.username,
    u.email,
    u.nome,
    u.nome_fantasia,
    c.ativo,
    c.pode_administrar,
    c.pode_votar
FROM usuario u
JOIN credencial c ON c.usuario_id = u.id
WHERE u.keycloak_id = sqlc.arg(keycloak_id);

-- name: FindByID :one
SELECT
    u.id,
    u.keycloak_id,
    u.username,
    u.email,
    u.nome,
    u.nome_fantasia,
    c.ativo,
    c.pode_administrar,
    c.pode_votar
FROM usuario u
JOIN credencial c ON c.usuario_id = u.id
WHERE u.id = sqlc.arg(id);

-- name: FindByEmail :one
SELECT
    u.id,
    u.keycloak_id,
    u.username,
    u.email,
    u.nome,
    u.nome_fantasia,
    c.ativo,
    c.pode_administrar,
    c.pode_votar
FROM usuario u
JOIN credencial c ON c.usuario_id = u.id
WHERE u.email = sqlc.arg(email);

-- name: FindByUsername :one
SELECT
    u.id,
    u.keycloak_id,
    u.username,
    u.email,
    u.nome,
    u.nome_fantasia,
    c.ativo,
    c.pode_administrar,
    c.pode_votar
FROM usuario u
JOIN credencial c ON c.usuario_id = u.id
WHERE u.username = sqlc.arg(username);

-- name: CreateUsuario :exec
INSERT INTO usuario (
    id,
    keycloak_id,
    username,
    email,
    nome,
    nome_fantasia
)
VALUES (
    sqlc.arg(id),
    sqlc.arg(keycloak_id),
    sqlc.arg(username),
    sqlc.arg(email),
    sqlc.arg(nome),
    sqlc.arg(nome_fantasia)
);

-- name: CreateCredencial :exec
INSERT INTO credencial (
    id,
    usuario_id,
    ativo,
    pode_administrar,
    pode_votar
)
VALUES (
    sqlc.arg(id),
    sqlc.arg(usuario_id),
    sqlc.arg(ativo),
    sqlc.arg(pode_administrar),
    sqlc.arg(pode_votar)
);

-- name: ListUsers :many
SELECT
    u.id,
    u.keycloak_id,
    u.username,
    u.email,
    u.nome,
    u.nome_fantasia,
    c.ativo,
    c.pode_administrar,
    c.pode_votar,
    COUNT(*) OVER() AS total_count
FROM usuario u
JOIN credencial c ON c.usuario_id = u.id
WHERE
    (
        sqlc.arg(nome) = ''
        OR u.nome ILIKE '%' || sqlc.arg(nome) || '%'
    )
AND (
        sqlc.arg(email) = ''
        OR u.email ILIKE '%' || sqlc.arg(email) || '%'
    )
AND (
        sqlc.arg(listar_inativos)::bool = true
        OR c.ativo = true
    )
ORDER BY u.nome
LIMIT sqlc.arg(limit_rows)
OFFSET sqlc.arg(offset_rows);

-- name: UpdateDisplayNamePermissions :exec
SELECT public.f_update_user_with_permissions(
    sqlc.arg(user_id),
    sqlc.arg(display_name),
    sqlc.arg(is_active),
    sqlc.arg(can_admin),
    sqlc.arg(can_vote)
);

-- name: UpdateDisplayName :exec
UPDATE usuario
SET nome_fantasia = sqlc.arg(display_name)
WHERE id = sqlc.arg(user_id);


