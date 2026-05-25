-- ============================================================
-- Remove o agendamento do pg_cron
-- ============================================================

DO $$
DECLARE
  v_jobid integer;
BEGIN
  SELECT jobid
  INTO v_jobid
  FROM cron.job
  WHERE jobname = 'fechar_votacoes_2am_horario_brasilia';

  IF v_jobid IS NOT NULL THEN
    PERFORM cron.unschedule(v_jobid);
  END IF;
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

-- DROP EXTENSION IF EXISTS pg_cron;