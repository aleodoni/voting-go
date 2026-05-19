CREATE TABLE IF NOT EXISTS reuniao (
  id varchar PRIMARY KEY NOT NULL,
  con_id integer NOT NULL,
  con_desc varchar NOT NULL,
  rec_id integer NOT NULL,
  con_sigla varchar NOT NULL,
  rec_tipo_reuniao varchar NOT NULL,
  rec_numero varchar NOT NULL,
  pac_id integer NOT NULL,
  rec_data date NOT NULL,
  created_at timestamp DEFAULT now() NOT NULL,
  updated_at timestamp DEFAULT now() NOT NULL,
  CONSTRAINT reuniao_rec_id_con_id_pac_id_unique UNIQUE(rec_id, con_id, pac_id)
);