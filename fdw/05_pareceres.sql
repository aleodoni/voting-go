DO $$
BEGIN
  IF NOT EXISTS (
    SELECT 1
    FROM information_schema.foreign_tables
    WHERE foreign_table_name = 'spl_votacao_pareceres_foreign'
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