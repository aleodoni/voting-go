-- ============================================================
-- FDW SPL Setup
-- Ambiente: staging / produção
-- ============================================================

-- ============================================================
-- 1. Extensão
-- ============================================================
CREATE EXTENSION IF NOT EXISTS postgres_fdw WITH SCHEMA public;

-- ============================================================
-- 2. Foreign Server
-- ============================================================
DO $$
BEGIN
  IF NOT EXISTS (
    SELECT 1
    FROM pg_foreign_server
    WHERE srvname = 'server_spl'
  ) THEN
    CREATE SERVER server_spl
      FOREIGN DATA WRAPPER postgres_fdw
      OPTIONS (
        host '${DB_SPL_HOST}',
        dbname '${DB_SPL_NAME}',
        port '${DB_SPL_PORT}',
        sslmode 'disable'
      );
  END IF;
END
$$;

-- ============================================================
-- 3. User Mapping
-- ============================================================
DO $$
BEGIN
  IF NOT EXISTS (
    SELECT 1
    FROM pg_user_mappings
    WHERE srvname = 'server_spl'
      AND umuser = (SELECT usesysid FROM pg_user WHERE usename = 'postgres')
  ) THEN
    CREATE USER MAPPING FOR postgres
      SERVER server_spl
      OPTIONS (
        user '${DB_SPL_USER}',
        password '${DB_SPL_PASSWORD}'
      );
  END IF;
END
$$;

-- ============================================================
-- 4. Foreign Table: reuniões
-- ============================================================
DO $$
BEGIN
  IF NOT EXISTS (
    SELECT 1
    FROM information_schema.foreign_tables
    WHERE foreign_table_schema = 'public'
      AND foreign_table_name = 'spl_votacao_reunioes_foreign'
  ) THEN
    CREATE FOREIGN TABLE public.spl_votacao_reunioes_foreign (
      versao integer,
      rec_id integer,
      rec_tipo_reuniao varchar,
      rec_numero varchar,
      con_desc varchar,
      con_id integer,
      pac_id integer,
      rec_data date,
      con_sigla varchar
    )
    SERVER server_spl
    OPTIONS (
      schema_name 'spl',
      table_name 'v_votacao_reunioes2'
    );
  END IF;
END
$$;

-- ============================================================
-- 5. Foreign Table: pareceres
-- ============================================================
DO $$
BEGIN
  IF NOT EXISTS (
    SELECT 1
    FROM information_schema.foreign_tables
    WHERE foreign_table_schema = 'public'
      AND foreign_table_name = 'spl_votacao_pareceres_foreign'
  ) THEN
    CREATE FOREIGN TABLE public.spl_votacao_pareceres_foreign (
      pro_codigo varchar(23),
      txt_id integer,
      vereador varchar(100),
      tcp_nome varchar(40)
    )
    SERVER server_spl
    OPTIONS (
      schema_name 'spl',
      table_name 'v_projetos_pareceres2'
    );
  END IF;
END
$$;

-- ============================================================
-- 6. Foreign Table: projetos
-- ============================================================
DO $$
BEGIN
  IF NOT EXISTS (
    SELECT 1
    FROM information_schema.foreign_tables
    WHERE foreign_table_schema = 'public'
      AND foreign_table_name = 'spl_votacao_projetos_foreign'
  ) THEN
    CREATE FOREIGN TABLE public.spl_votacao_projetos_foreign (
      sumula text,
      comissao varchar,
      paralelo boolean,
      relator varchar(100),
      tem_emendas boolean,
      pac_id integer,
      par_id integer,
      codigo_proposicao varchar(23),
      iniciativa text,
      conclusao_comissao varchar,
      conclusao_relator varchar(40)
    )
    SERVER server_spl
    OPTIONS (
      schema_name 'spl',
      table_name 'v_projetos_comissao2'
    );
  END IF;
END
$$;