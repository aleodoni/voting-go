-- ============================================================
-- Função: f_save_vote()
-- Objetivo: Salvar um voto principal e, opcionalmente, seus 
--            registros de restrição e voto contrário.
-- ============================================================

CREATE OR REPLACE FUNCTION f_save_vote(
    p_voto_id text,
    p_usuario_id text,
    p_votacao_id text,
    p_voto opcao_voto,
    p_restricao jsonb DEFAULT null,
    p_voto_contrario jsonb DEFAULT null
)
RETURNS void
LANGUAGE plpgsql
AS $$
DECLARE
    new_vote_id text;
BEGIN
    -- Inserir o voto principal
    INSERT INTO voto (id, usuario_id, votacao_id, voto)
    VALUES (p_voto_id, p_usuario_id, p_votacao_id, p_voto)
    RETURNING id INTO new_vote_id;

    -- Inserir restrição, se houver
    IF p_restricao IS NOT NULL THEN
        INSERT INTO restricao (id, voto_id, restricao)
        VALUES (
            p_restricao->>'restricao_id',
            new_vote_id,
            p_restricao->>'restricao'
        );
    END IF;

    -- Inserir voto contrário, se houver
    IF p_voto_contrario IS NOT NULL THEN
        INSERT INTO voto_contrario (id, id_texto, voto_id, parecer_id)
        VALUES (
            p_voto_contrario->>'id_voto_contrario',
            (p_voto_contrario->>'id_texto')::int,
            new_vote_id,
            p_voto_contrario #>> '{opinion,parecer_id}'
        );
    END IF;
END;
$$;
