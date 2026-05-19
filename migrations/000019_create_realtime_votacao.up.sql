DO $$
BEGIN
  -- 🔹 1. Garante que a publicação exista
  IF NOT EXISTS (SELECT 1 FROM pg_publication WHERE pubname = 'supabase_realtime') THEN
    EXECUTE 'CREATE PUBLICATION supabase_realtime;';
  END IF;

  -- 🔹 2. Configura a tabela para replicação completa
  BEGIN
    EXECUTE 'ALTER TABLE public.votacao REPLICA IDENTITY FULL;';
  EXCEPTION
    WHEN others THEN
      RAISE NOTICE 'Aviso: não foi possível alterar REPLICA IDENTITY FULL para public.votacao (%).', SQLERRM;
  END;

  -- 🔹 3. Adiciona a tabela à publicação se não estiver
  IF NOT EXISTS (
    SELECT 1
    FROM pg_publication_tables
    WHERE pubname = 'supabase_realtime'
      AND schemaname = 'public'
      AND tablename = 'votacao'
  ) THEN
    EXECUTE 'ALTER PUBLICATION supabase_realtime ADD TABLE public.votacao;';
  END IF;
END $$;