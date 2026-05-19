ALTER TABLE usuario
ADD COLUMN username varchar NOT NULL;

ALTER TABLE usuario
ADD CONSTRAINT usuario_username_unique UNIQUE(username);