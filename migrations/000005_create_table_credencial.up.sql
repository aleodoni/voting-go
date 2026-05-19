CREATE TABLE IF NOT EXISTS credencial (
  "id" varchar PRIMARY KEY NOT NULL,
  "ativo" boolean DEFAULT false NOT NULL,
  "pode_administrar" boolean DEFAULT false NOT NULL,
  "pode_votar" boolean DEFAULT false NOT NULL,
  "created_at" timestamp DEFAULT now() NOT NULL,
  "updated_at" timestamp DEFAULT now() NOT NULL,
  "usuario_id" varchar NOT NULL UNIQUE,
  CONSTRAINT "credencial_usuario_id_fk" FOREIGN KEY ("usuario_id") REFERENCES "usuario"("id") ON DELETE CASCADE
);