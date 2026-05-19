DO $$
BEGIN
  -- 🔹 1. Remove a tabela da publicação, se existir
  IF EXISTS (
    SELECT 1
    FROM pg_publication_tables
    WHERE pubname = 'supabase_realtime'
      AND schemaname = 'public'
      AND tablename = 'votacao'
  ) THEN
    EXECUTE 'ALTER PUBLICATION supabase_realtime DROP TABLE public.votacao;';
  ELSE
    RAISE NOTICE 'Tabela public.votacao não estava na publicação supabase_realtime.';
  END IF;

  -- 🔹 2. Restaura REPLICA IDENTITY para default (ou nada)
  BEGIN
    EXECUTE 'ALTER TABLE public.votacao REPLICA IDENTITY DEFAULT;';
  EXCEPTION
    WHEN others THEN
      RAISE NOTICE 'Aviso: não foi possível restaurar REPLICA IDENTITY para public.votacao (%).', SQLERRM;
  END;
END $$;