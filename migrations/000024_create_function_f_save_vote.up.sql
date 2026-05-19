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
    IF p_voto = 'V' THEN
        -- Remove votos anteriores desta votação
        DELETE FROM restricao
        WHERE voto_id IN (
            SELECT id FROM voto WHERE votacao_id = p_votacao_id
        );

        DELETE FROM voto_contrario
        WHERE voto_id IN (
            SELECT id FROM voto WHERE votacao_id = p_votacao_id
        );

        DELETE FROM voto WHERE votacao_id = p_votacao_id;

        -- Insere apenas o voto vistas
        INSERT INTO voto (id, usuario_id, votacao_id, voto)
        VALUES (p_voto_id, p_usuario_id, p_votacao_id, p_voto);

        -- Encerra a votação
        UPDATE votacao
        SET status = 'V'
        WHERE id = p_votacao_id;

        RETURN;
    END IF;

    -- Fluxo normal para os outros tipos de voto
    INSERT INTO voto (id, usuario_id, votacao_id, voto)
    VALUES (p_voto_id, p_usuario_id, p_votacao_id, p_voto)
    RETURNING id INTO new_vote_id;

    IF p_restricao IS NOT NULL THEN
        INSERT INTO restricao (id, voto_id, restricao)
        VALUES (
            p_restricao->>'restricao_id',
            new_vote_id,
            p_restricao->>'restricao'
        );
    END IF;

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