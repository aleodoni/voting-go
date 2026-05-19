CREATE TABLE IF NOT EXISTS "voto" (
  "id" varchar PRIMARY KEY NOT NULL,
  "voto" "opcao_voto" NOT NULL,
  "created_at" timestamp DEFAULT now() NOT NULL,
  "updated_at" timestamp DEFAULT now() NOT NULL,
  "votacao_id" varchar NOT NULL,
  "usuario_id" varchar NOT NULL,
  CONSTRAINT "voto_votacao_id_fk" FOREIGN KEY ("votacao_id") REFERENCES "votacao"("id") ON DELETE CASCADE,
  CONSTRAINT "voto_usuario_id_fk" FOREIGN KEY ("usuario_id") REFERENCES "usuario"("id") ON DELETE CASCADE
);