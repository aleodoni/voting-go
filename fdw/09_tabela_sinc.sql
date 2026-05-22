-- CREATE TABLE public.sincronia (
--   id text PRIMARY KEY DEFAULT public.cuid2(),

--   iniciado_em timestamp NOT NULL DEFAULT now(),
--   finalizado_em timestamp,

--   sucesso boolean,
--   mensagem_erro text,

--   reunioes_sincronizadas integer DEFAULT 0,
--   projetos_sincronizados integer DEFAULT 0,
--   pareceres_sincronizados integer DEFAULT 0
-- );