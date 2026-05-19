-- ============================================================
-- Configura tabela "voto" para replicação no Supabase Realtime
-- ============================================================
DO $$
BEGIN
-- 🔹 1. Configura a tabela para replicação completa
  BEGIN
    EXECUTE 'ALTER TABLE public.voto REPLICA IDENTITY FULL;';
  EXCEPTION
    WHEN others THEN
      RAISE NOTICE 'Aviso: não foi possível alterar REPLICA IDENTITY FULL para public.voto (%).', SQLERRM;
  END;

  -- 🔹 2. Adiciona a tabela à publicação "supabase_realtime", se ainda não estiver
  IF NOT EXISTS (
    SELECT 1
    FROM pg_publication_tables
    WHERE pubname = 'supabase_realtime'
      AND schemaname = 'public'
      AND tablename = 'voto'
  ) THEN
    EXECUTE 'ALTER PUBLICATION supabase_realtime ADD TABLE public.voto;';
  ELSE
    RAISE NOTICE 'Tabela public.voto já está incluída na publicação supabase_realtime.';
  END IF;
END $$;
