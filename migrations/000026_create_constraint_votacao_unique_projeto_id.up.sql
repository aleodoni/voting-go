ALTER TABLE votacao
ADD CONSTRAINT uq_votacao_projeto UNIQUE (projeto_id);