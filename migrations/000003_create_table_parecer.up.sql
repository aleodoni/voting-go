CREATE TABLE IF NOT EXISTS parecer (
  "id" varchar PRIMARY KEY NOT NULL,
  "codigo_proposicao" varchar NOT NULL,
  "tcp_nome" varchar NOT NULL,
  "vereador" varchar NOT NULL,
  "id_texto" integer NOT NULL,
  "created_at" timestamp DEFAULT now() NOT NULL,
  "updated_at" timestamp DEFAULT now() NOT NULL,
  "projeto_id" varchar NOT NULL,
  CONSTRAINT "unique_parecer_projeto_id_codigo_proposicao_id_texto" UNIQUE("projeto_id", "codigo_proposicao", "id_texto"),
  CONSTRAINT "parecer_projeto_id_fk" FOREIGN KEY ("projeto_id") REFERENCES "projeto"("id") ON DELETE CASCADE
);