-- ============================================================
-- Extensão: pg_cron
-- Objetivo: Garantir a existência da extensão para agendamentos.
-- ============================================================
CREATE EXTENSION IF NOT EXISTS pg_cron;

-- ============================================================
-- Função: f_fechar_votacoes_abertas()
-- Objetivo: Fechar automaticamente todas as votações abertas 
--            (status = 'A'), marcando-as como 'V' (votadas).
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
-- Nome da tarefa: fechar_votacoes_2am_horario_brasilia
-- Horário: 05:00 UTC = 02:00 horário de Brasília (UTC-3)
-- ============================================================

-- Remove o job antigo, se existir, antes de recriar
DO $$
BEGIN
  PERFORM cron.unschedule('fechar_votacoes_2am_horario_brasilia');
EXCEPTION
  WHEN others THEN
    NULL;
END;
$$;

-- Cria (ou recria) o agendamento
SELECT cron.schedule(
  'fechar_votacoes_2am_horario_brasilia',
  '0 5 * * *',  -- executa diariamente às 05:00 UTC
  $$SELECT f_fechar_votacoes_abertas();$$
);

-- Consulta os jobs registrados, se desejar:
-- SELECT * FROM cron.job;
