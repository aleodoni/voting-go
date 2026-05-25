-- ============================================================
-- Extensão: pg_cron
-- Objetivo: Utilizar agendamentos automáticos via pg_cron.
--
-- Observação:
-- A criação da extensão foi removida porque alguns ambientes
-- gerenciados (ex.: AWS RDS) não permitem CREATE EXTENSION
-- durante migrations da aplicação.
-- ============================================================

-- CREATE EXTENSION IF NOT EXISTS pg_cron;

-- ============================================================
-- Função: f_fechar_votacoes_abertas()
-- Objetivo:
-- Fechar automaticamente todas as votações abertas
-- (status = 'A'), marcando-as como 'V' (votadas).
-- ============================================================

CREATE OR REPLACE FUNCTION f_fechar_votacoes_abertas()
RETURNS void AS $$
BEGIN
  UPDATE votacao
  SET status = 'V'
  WHERE status = 'A';
END;
$$ LANGUAGE plpgsql;

-- ============================================================
-- Agendamento via pg_cron
--
-- Nome do job:
--   fechar_votacoes_2am_horario_brasilia
--
-- Horário:
--   05:00 UTC = 02:00 horário de Brasília (UTC-3)
--
-- Observação AWS RDS:
-- O pg_cron registra jobs no banco "postgres" por padrão.
-- Por isso, após criar o job, ajustamos explicitamente
-- o banco de execução para "voting_db".
-- ============================================================

-- Remove o job antigo, se existir
-- DO $$
-- DECLARE
--   v_jobid integer;
-- BEGIN
--   SELECT jobid
--   INTO v_jobid
--   FROM cron.job
--   WHERE jobname = 'fechar_votacoes_2am_horario_brasilia';

--   IF v_jobid IS NOT NULL THEN
--     PERFORM cron.unschedule(v_jobid);
--   END IF;
-- END;
-- $$;

-- Cria o agendamento
-- SELECT cron.schedule(
--   'fechar_votacoes_2am_horario_brasilia',
--   '0 5 * * *',
--   $$SELECT f_fechar_votacoes_abertas();$$
-- );

-- Garante execução no banco correto (AWS RDS)
-- UPDATE cron.job
-- SET database = 'voting_db'
-- WHERE jobname = 'fechar_votacoes_2am_horario_brasilia';

-- ============================================================
-- Consulta útil para validação:
--
-- SELECT
--   jobid,
--   jobname,
--   database,
--   schedule,
--   active
-- FROM cron.job;
-- ============================================================