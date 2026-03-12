CREATE TABLE IF NOT EXISTS "voto_contrario" (
  "id" varchar PRIMARY KEY NOT NULL,
  "id_texto" integer NOT NULL,
  "created_at" timestamp DEFAULT now() NOT NULL,
  "updated_at" timestamp DEFAULT now() NOT NULL,
  "voto_id" varchar NOT NULL,
  "parecer_id" varchar NOT NULL,
  CONSTRAINT "voto_contrario_voto_id_fk" FOREIGN KEY ("voto_id") REFERENCES "voto"("id") ON DELETE CASCADE,
  CONSTRAINT "voto_contrario_parecer_id_fk" FOREIGN KEY ("parecer_id") REFERENCES "parecer"("id") ON DELETE CASCADE
);