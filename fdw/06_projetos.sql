DO $$
BEGIN
  IF NOT EXISTS (
    SELECT 1
    FROM information_schema.foreign_tables
    WHERE foreign_table_name = 'spl_votacao_projetos_foreign'
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