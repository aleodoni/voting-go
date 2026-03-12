-- ============================================================
-- Função: f_get_project_open_voting()
-- Retorna o projeto com votação aberta (status = 'A'),
-- incluindo seus votos, usuários e pareceres.
-- ============================================================

CREATE OR REPLACE FUNCTION f_get_project_open_voting()
RETURNS json
LANGUAGE plpgsql
AS $$
DECLARE
    result json;
BEGIN
    SELECT json_build_object(
        'projeto', p.*,
        'votacao', json_build_object(
            'id', v.id,
            'projetoId', v.projeto_id,
            'status', v.status,
            'createdAt', v.created_at,
            'updatedAt', v.updated_at,
            'votos', (
                SELECT COALESCE(
                    json_agg(
                        json_build_object(
                            'id', vo.id,
                            'voto', vo.voto,
                            'votacaoId', vo.votacao_id,
                            'usuarioId', vo.usuario_id,
                            'createdAt', vo.created_at,
                            'updatedAt', vo.updated_at,
                            'usuario', json_build_object(
                                'id', u.id,
                                'sub', u.sub,
                                'email', u.email,
                                'nome', u.nome,
                                'nomeFantasia', u.nome_fantasia,
                                'createdAt', u.created_at,
                                'updatedAt', u.updated_at
                            )
                        )
                    ),
                    '[]'::json
                )
                FROM voto vo
                INNER JOIN usuario u 
                    ON vo.usuario_id = u.id
                WHERE vo.votacao_id = v.id
            )
        ),
        'pareceres', (
            SELECT COALESCE(
                json_agg(
                    json_build_object(
                        'id', pa.id,
                        'codigoProposicao', pa.codigo_proposicao,
                        'tcpNome', pa.tcp_nome,
                        'vereador', pa.vereador,
                        'idTexto', pa.id_texto,
                        'projetoId', pa.projeto_id,
                        'createdAt', pa.created_at,
                        'updatedAt', pa.updated_at
                    )
                ),
                '[]'::json
            )
            FROM parecer pa
            WHERE pa.projeto_id = p.id
        )
    )
    INTO result
    FROM projeto p
    INNER JOIN votacao v 
        ON p.id = v.projeto_id
    WHERE v.status = 'A'
    LIMIT 1;

    RETURN result;
END;
$$;
