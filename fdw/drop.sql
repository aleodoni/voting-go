-- ============================================================
-- Remove job do pg_cron
-- ============================================================

-- DO $$
-- DECLARE
--   v_jobid integer;
-- BEGIN
--   SELECT jobid
--   INTO v_jobid
--   FROM cron.job
--   WHERE jobname = 'p_spl_daily_sync';

--   IF v_jobid IS NOT NULL THEN
--     PERFORM cron.unschedule(v_jobid);
--   END IF;
-- END;
-- $$;


DROP PROCEDURE IF EXISTS public.p_spl_daily_sync;
-- DROP FUNCTION IF EXISTS public.cuid2();
-- DROP TABLE IF EXISTS public.sincronia;

DROP FOREIGN TABLE IF EXISTS public.spl_votacao_reunioes_foreign;
DROP FOREIGN TABLE IF EXISTS public.spl_votacao_pareceres_foreign;
DROP FOREIGN TABLE IF EXISTS public.spl_votacao_projetos_foreign;

DROP USER MAPPING IF EXISTS FOR postgres SERVER server_spl;

DROP SERVER IF EXISTS server_spl CASCADE;

DROP EXTENSION IF EXISTS postgres_fdw;