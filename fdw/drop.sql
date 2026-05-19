DROP PROCEDURE IF EXISTS public.p_spl_daily_sync;
DROP FUNCTION IF EXISTS public.cuid2();
DROP TABLE IF EXISTS public.sincronia;

DROP FOREIGN TABLE IF EXISTS public.spl_votacao_reunioes_foreign;
DROP FOREIGN TABLE IF EXISTS public.spl_votacao_pareceres_foreign;
DROP FOREIGN TABLE IF EXISTS public.spl_votacao_projetos_foreign;

DROP USER MAPPING IF EXISTS FOR postgres SERVER server_spl;

DROP SERVER IF EXISTS server_spl CASCADE;

DROP EXTENSION IF EXISTS postgres_fdw;