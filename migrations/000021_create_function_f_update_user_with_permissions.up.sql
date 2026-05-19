-- ===============================================
-- Função: f_update_user_with_permissions (idempotente)
-- ===============================================

CREATE OR REPLACE FUNCTION public.f_update_user_with_permissions(
  p_user_id text,
  p_display_name text,
  p_is_active boolean,
  p_can_admin boolean,
  p_can_vote boolean
)
RETURNS void
AS $$
BEGIN
  -- Atualiza nome_fantasia do usuário
  UPDATE public.usuario
    SET nome_fantasia = p_display_name,
        updated_at = now()
  WHERE id = p_user_id;

  -- Atualiza permissões associadas
  UPDATE public.credencial
    SET ativo = p_is_active,
        pode_administrar = p_can_admin,
        pode_votar = p_can_vote,
        updated_at = now()
  WHERE usuario_id = p_user_id;
END;
$$ LANGUAGE plpgsql;
