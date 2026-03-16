-- ===============================================
-- Função: f_get_users_paginated_with_total (idempotente)
-- ===============================================
CREATE OR REPLACE FUNCTION f_get_users_paginated_with_total(
  p_name text DEFAULT NULL,
  p_email text DEFAULT NULL,
  p_skip int DEFAULT 0,
  p_take int DEFAULT 10
)
RETURNS TABLE (
  id varchar,
  keycloak_id varchar,
  nome varchar,
  nome_fantasia varchar,
  email varchar,
  ativo boolean,
  pode_administrar boolean,
  pode_votar boolean,
  updated_at timestamp,
  total_count int
)
AS $$
BEGIN
  RETURN QUERY
    SELECT 
      up.id::varchar,
      up.keycloak_id::varchar,
      COALESCE(up.nome, '')::varchar,
      COALESCE(up.nome_fantasia, '')::varchar,
      up.email::varchar,
      COALESCE(up.ativo, false)::boolean,
      COALESCE(up.pode_administrar, false)::boolean,
      COALESCE(up.pode_votar, false)::boolean,
      COALESCE(up.updated_at, now())::timestamp,
      COUNT(*) OVER()::int AS total_count
    FROM public.v_usuario_permissoes up
    WHERE 
      (p_name IS NULL OR p_name = '' OR (COALESCE(up.nome_fantasia, up.nome) ILIKE '%' || p_name || '%'))
      AND (p_email IS NULL OR p_email = '' OR up.email ILIKE '%' || p_email || '%')
    ORDER BY COALESCE(up.nome_fantasia, up.nome)
    OFFSET p_skip
    LIMIT p_take;
END;
$$ LANGUAGE plpgsql STABLE;
