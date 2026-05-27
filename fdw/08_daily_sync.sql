CREATE OR REPLACE PROCEDURE public.p_spl_daily_sync()
LANGUAGE plpgsql
AS $$
DECLARE
  v_execucao_id text;

  v_reunioes integer := 0;
  v_projetos integer := 0;
  v_pareceres integer := 0;

BEGIN

  -- =========================================================
  -- EVITA EXECUÇÕES CONCORRENTES
  -- =========================================================
  IF NOT pg_try_advisory_lock(987654) THEN
    RAISE EXCEPTION 'Sincronização SPL já está em execução';
  END IF;

  -- =========================================================
  -- REGISTRA INÍCIO DA EXECUÇÃO
  -- =========================================================
  INSERT INTO public.sincronia (
    id,
    iniciado_em,
    sucesso
  )
  VALUES (
    public.cuid2(),
    now(),
    null
  )
  RETURNING id INTO v_execucao_id;

  BEGIN

    -- =========================================================
    -- SINCRONIZA REUNIÕES
    -- =========================================================

    -- Remove reuniões do dia que não existem mais na origem
    -- CASCADE remove projetos, pareceres e votações
    DELETE FROM public.reuniao r
    WHERE r.rec_data::date = CURRENT_DATE
      AND NOT EXISTS (
        SELECT 1
        FROM spl_votacao_reunioes_foreign s
        WHERE s.rec_id = r.rec_id
          AND s.con_id = r.con_id
          AND s.pac_id = r.pac_id
      );

    -- Insere ou atualiza reuniões
    INSERT INTO public.reuniao (
      id,
      con_id,
      con_desc,
      rec_id,
      con_sigla,
      rec_tipo_reuniao,
      rec_numero,
      pac_id,
      rec_data,
      created_at,
      updated_at
    )
    SELECT
      public.cuid2(),
      s.con_id,
      s.con_desc,
      s.rec_id,
      s.con_sigla,
      s.rec_tipo_reuniao,
      s.rec_numero,
      s.pac_id,
      s.rec_data,
      now(),
      now()
    FROM spl_votacao_reunioes_foreign s
    WHERE s.rec_data::date = CURRENT_DATE
    ON CONFLICT (rec_id, con_id, pac_id)
    DO UPDATE SET
      con_desc = EXCLUDED.con_desc,
      con_sigla = EXCLUDED.con_sigla,
      rec_tipo_reuniao = EXCLUDED.rec_tipo_reuniao,
      rec_numero = EXCLUDED.rec_numero,
      rec_data = EXCLUDED.rec_data,
      updated_at = now();

    GET DIAGNOSTICS v_reunioes = ROW_COUNT;

    -- =========================================================
    -- SINCRONIZA PROJETOS
    -- =========================================================

    -- Remove projetos do dia que não existem mais na origem
    -- CASCADE remove pareceres e votações
    DELETE FROM public.projeto p
    WHERE EXISTS (
      SELECT 1
      FROM public.reuniao r
      WHERE r.id = p.reuniao_id
        AND r.rec_data::date = CURRENT_DATE
    )
    AND NOT EXISTS (
      SELECT 1
      FROM spl_votacao_projetos_foreign s
      WHERE s.pac_id = p.pac_id
        AND s.par_id = p.par_id
        AND s.codigo_proposicao = p.codigo_proposicao
    );

    -- Insere ou atualiza projetos
    INSERT INTO public.projeto (
      id,
      sumula,
      relator,
      tem_emendas,
      pac_id,
      par_id,
      codigo_proposicao,
      iniciativa,
      conclusao_comissao,
      conclusao_relator,
      created_at,
      updated_at,
      reuniao_id
    )
    SELECT
      public.cuid2(),
      p.sumula,
      p.relator,
      p.tem_emendas,
      p.pac_id,
      p.par_id,
      p.codigo_proposicao,
      p.iniciativa,
      p.conclusao_comissao,
      p.conclusao_relator,
      now(),
      now(),
      r.id
    FROM spl_votacao_projetos_foreign p
    JOIN public.reuniao r
      ON r.pac_id = p.pac_id
    WHERE r.rec_data::date = CURRENT_DATE
    ON CONFLICT (reuniao_id, pac_id, par_id, codigo_proposicao)
    DO UPDATE SET
      sumula = EXCLUDED.sumula,
      relator = EXCLUDED.relator,
      tem_emendas = EXCLUDED.tem_emendas,
      iniciativa = EXCLUDED.iniciativa,
      conclusao_comissao = EXCLUDED.conclusao_comissao,
      conclusao_relator = EXCLUDED.conclusao_relator,
      updated_at = now();

    GET DIAGNOSTICS v_projetos = ROW_COUNT;

    -- =========================================================
    -- SINCRONIZA PARECERES
    -- =========================================================

    -- Remove pareceres do dia que não existem mais na origem
    -- VERIFICAR QUANDO OS PARECERES SÃO EXCLUÍDOS
    DELETE FROM public.parecer pa
    WHERE EXISTS (
      SELECT 1
      FROM public.projeto p
      JOIN public.reuniao r
        ON r.id = p.reuniao_id
      WHERE p.id = pa.projeto_id
        AND r.rec_data::date = CURRENT_DATE
    )
    AND NOT EXISTS (
      SELECT 1
      FROM spl_votacao_pareceres_foreign s
      WHERE s.pro_codigo = pa.codigo_proposicao
        AND s.txt_id = pa.id_texto
    );

    -- Insere ou atualiza pareceres
    INSERT INTO public.parecer (
      id,
      codigo_proposicao,
      tcp_nome,
      vereador,
      id_texto,
      created_at,
      updated_at,
      projeto_id
    )
    SELECT
      public.cuid2(),
      pa.pro_codigo,
      pa.tcp_nome,
      pa.vereador,
      pa.txt_id,
      now(),
      now(),
      p.id
    FROM spl_votacao_pareceres_foreign pa
    JOIN public.projeto p
      ON p.codigo_proposicao = pa.pro_codigo
    JOIN public.reuniao r
      ON r.id = p.reuniao_id
    WHERE r.rec_data::date = CURRENT_DATE
    ON CONFLICT (projeto_id, codigo_proposicao, id_texto)
    DO UPDATE SET
      tcp_nome = EXCLUDED.tcp_nome,
      vereador = EXCLUDED.vereador,
      id_texto = EXCLUDED.id_texto,
      updated_at = now();

    GET DIAGNOSTICS v_pareceres = ROW_COUNT;

    -- =========================================================
    -- FINALIZA EXECUÇÃO COM SUCESSO
    -- =========================================================
    UPDATE public.sincronia
    SET
      finalizado_em = now(),
      sucesso = true,
      reunioes_sincronizadas = v_reunioes,
      projetos_sincronizados = v_projetos,
      pareceres_sincronizados = v_pareceres
    WHERE id = v_execucao_id;

  EXCEPTION
    WHEN OTHERS THEN

      UPDATE public.sincronia
      SET
        finalizado_em = now(),
        sucesso = false,
        mensagem_erro = SQLERRM
      WHERE id = v_execucao_id;

      PERFORM pg_advisory_unlock(987654);

      RAISE;
  END;

  -- =========================================================
  -- LIBERA LOCK
  -- =========================================================
  PERFORM pg_advisory_unlock(987654);

END;
$$;

-- ============================================================
-- AGENDAMENTO VIA PG_CRON
--
-- Nome:
--   p_spl_daily_sync
--
-- Horário:
--   06:00 UTC = 03:00 horário de Brasília (UTC-3)
--
-- AWS RDS:
--   Executa explicitamente no banco voting_db.
-- ============================================================

-- Remove job antigo, se existir
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

-- Cria o agendamento
-- SELECT cron.schedule(
--   'p_spl_daily_sync',
--   '0 6 * * *',
--   $$CALL public.p_spl_daily_sync();$$
-- );

-- Garante execução no banco correto (AWS RDS)
-- UPDATE cron.job
-- SET database = 'voting_db'
-- WHERE jobname = 'p_spl_daily_sync';

-- ============================================================
-- CONSULTA ÚTIL
--
-- SELECT
--   jobid,
--   jobname,
--   database,
--   schedule,
--   active
-- FROM cron.job;
-- ============================================================