DO $$
BEGIN
  IF NOT EXISTS (
    SELECT 1
    FROM information_schema.foreign_tables
    WHERE foreign_table_name = 'spl_votacao_reunioes_foreign'
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