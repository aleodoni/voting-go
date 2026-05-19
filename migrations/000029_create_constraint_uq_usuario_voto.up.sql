ALTER TABLE voto
ADD CONSTRAINT uq_usuario_votacao UNIQUE (usuario_id, votacao_id);