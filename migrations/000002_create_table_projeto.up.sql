CREATE TABLE IF NOT EXISTS projeto (
  "id" varchar PRIMARY KEY NOT NULL,
  "sumula" varchar NOT NULL,
  "relator" varchar NOT NULL,
  "tem_emendas" boolean NOT NULL,
  "pac_id" integer NOT NULL,
  "par_id" integer NOT NULL,
  "codigo_proposicao" varchar NOT NULL,
  "iniciativa" varchar NOT NULL,
  "conclusao_comissao" varchar NOT NULL,
  "conclusao_relator" varchar NOT NULL,
  "created_at" timestamp DEFAULT now() NOT NULL,
  "updated_at" timestamp DEFAULT now() NOT NULL,
  "reuniao_id" varchar NOT NULL,
  CONSTRAINT "unique_reuniao_id_pac_par_codigo" UNIQUE("reuniao_id","pac_id","par_id","codigo_proposicao"),
  CONSTRAINT "projeto_reuniao_id_fk" FOREIGN KEY ("reuniao_id") REFERENCES "reuniao"("id") ON DELETE CASCADE
);