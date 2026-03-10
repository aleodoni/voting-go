CREATE TABLE IF NOT EXISTS "votacao" (
  "id" varchar PRIMARY KEY NOT NULL,
  "projeto_id" varchar,
  "status" "status_votacao" DEFAULT 'F' NOT NULL,
  "created_at" timestamp DEFAULT now() NOT NULL,
  "updated_at" timestamp DEFAULT now() NOT NULL,
  CONSTRAINT "votacao_projeto_id_unique" UNIQUE("projeto_id"),
  CONSTRAINT "votacao_projeto_id_fk" FOREIGN KEY ("projeto_id") REFERENCES "projeto"("id")
);