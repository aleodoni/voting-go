CREATE TABLE IF NOT EXISTS usuario (
  "id" varchar PRIMARY KEY NOT NULL,
  "keycloak_id" varchar NOT NULL,
  "email" varchar NOT NULL,
  "nome" varchar NOT NULL,
  "nome_fantasia" varchar,
  "created_at" timestamp DEFAULT now() NOT NULL,
  "updated_at" timestamp DEFAULT now() NOT NULL,
  CONSTRAINT "usuario_keycloak_id_unique" UNIQUE("keycloak_id"),
  CONSTRAINT "usuario_email_unique" UNIQUE("email")
);