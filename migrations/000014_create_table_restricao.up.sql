CREATE TABLE IF NOT EXISTS "restricao" (
  "id" varchar PRIMARY KEY NOT NULL,
  "restricao" varchar(500) NOT NULL,
  "created_at" timestamp DEFAULT now() NOT NULL,
  "updated_at" timestamp DEFAULT now() NOT NULL,
  "voto_id" varchar NOT NULL,
  CONSTRAINT "restricao_voto_id_fk" FOREIGN KEY ("voto_id") REFERENCES "voto"("id") ON DELETE CASCADE
);