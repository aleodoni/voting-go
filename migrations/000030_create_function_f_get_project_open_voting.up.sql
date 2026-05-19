CREATE OR REPLACE FUNCTION public.f_get_project_open_voting()
RETURNS json
LANGUAGE plpgsql
AS $$
DECLARE
    result json;
BEGIN
    SELECT json_build_object(
        'projetoId',              p.id,
        'sumula',                 p.sumula,
        'relator',                p.relator,
        'temEmendas',             p.tem_emendas,
        'pacId',                  p.pac_id,
        'parId',                  p.par_id,
        'codigoProposicao',       p.codigo_proposicao,
        'iniciativa',             p.iniciativa,
        'conclusaoComissao',      p.conclusao_comissao,
        'conclusaoRelator',       p.conclusao_relator,
        'createdAt',              p.created_at,
        'updatedAt',              p.updated_at,
        'reuniaoId',              p.reuniao_id,
        'votacaoId',              v.id,
        'votacaoStatus',          v.status,
        'votacaoCreatedAt',       v.created_at,
        'votacaoUpdatedAt',       v.updated_at,
        'votos', (
            SELECT COALESCE(
                json_agg(
                    json_build_object(
                        'id',        vo.id,
                        'voto',      vo.voto,
                        'votacaoId', vo.votacao_id,
                        'usuarioId', vo.usuario_id,
                        'createdAt', vo.created_at,
                        'updatedAt', vo.updated_at,
                        'usuario', json_build_object(
                            'id',           u.id,
                            'keycloakId',   u.keycloak_id,
                            'email',        u.email,
                            'nome',         u.nome,
                            'nomeFantasia', u.nome_fantasia,
                            'username',     u.username,
                            'createdAt',    u.created_at,
                            'updatedAt',    u.updated_at
                        )
                    )
                ),
                '[]'::json
            )
            FROM public.voto vo
            INNER JOIN public.usuario u ON vo.usuario_id = u.id
            WHERE vo.votacao_id = v.id
        ),
        'pareceres', (
            SELECT COALESCE(
                json_agg(
                    json_build_object(
                        'id',               pa.id,
                        'codigoProposicao', pa.codigo_proposicao,
                        'tcpNome',          pa.tcp_nome,
                        'vereador',         pa.vereador,
                        'idTexto',          pa.id_texto,
                        'projetoId',        pa.projeto_id,
                        'createdAt',        pa.created_at,
                        'updatedAt',        pa.updated_at
                    )
                ),
                '[]'::json
            )
            FROM public.parecer pa
            WHERE pa.projeto_id = p.id
        )
    )
    INTO result
    FROM public.projeto p
    INNER JOIN public.votacao v ON p.id = v.projeto_id
    WHERE v.status = 'A'
    LIMIT 1;

    RETURN result;
END;
$$;