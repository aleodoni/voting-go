-- ============================================================
-- Remove o agendamento do pg_cron
-- ============================================================

DO $$
BEGIN
  PERFORM cron.unschedule('fechar_votacoes_2am_horario_brasilia');
EXCEPTION
  WHEN others THEN
    NULL;
END;
$$;

-- ============================================================
-- Remove a função
-- ============================================================

DROP FUNCTION IF EXISTS f_fechar_votacoes_abertas();

-- ============================================================
-- (Opcional) Remove a extensão
-- Só faça isso se tiver certeza que nenhuma outra migration usa
-- ============================================================

DROP EXTENSION IF EXISTS pg_cron;