-- ===============================================
-- Função: f_get_projetos_completo (idempotente)
-- ===============================================

CREATE OR REPLACE FUNCTION public.f_get_projetos_completo(p_reuniao_id varchar)
RETURNS jsonb
LANGUAGE sql
AS $$
SELECT coalesce(jsonb_agg(projeto_json), '[]'::jsonb)
FROM (
    SELECT jsonb_build_object(
        'projeto', p.*,
        'pareceres', coalesce(pareceres.pareceres_array, '[]'::jsonb),
        'votacoes', coalesce(votacoes.votacoes_array, '[]'::jsonb)
    ) AS projeto_json
    FROM public.projeto p

    -- Pareceres
    LEFT JOIN LATERAL (
        SELECT jsonb_agg(pa.*) AS pareceres_array
        FROM public.parecer pa
        WHERE pa.projeto_id = p.id
    ) pareceres ON true

    -- Votações + votos + restrições + votoContrário
    LEFT JOIN LATERAL (
        SELECT jsonb_agg(
            jsonb_build_object(
                'id', v.id,
                'projeto_id', v.projeto_id,
                'status', v.status,
                'votos', coalesce(votos.votos_array, '[]'::jsonb)
            )
        ) AS votacoes_array
        FROM public.votacao v
        LEFT JOIN LATERAL (
            SELECT jsonb_agg(
                jsonb_build_object(
                    'id', vo.id,
                    'voto', vo.voto,
                    'created_at', vo.created_at,
                    'updated_at', vo.updated_at,
                    'votacao_id', vo.votacao_id,
                    'usuario_id', vo.usuario_id,
                    'usuario', u.*,
                    'restricoes', coalesce(restricoes.restricoes_array, '[]'::jsonb),
                    'votoContrario', coalesce(vc.voto_contrario_array, '[]'::jsonb)
                )
            ) AS votos_array
            FROM public.voto vo
            LEFT JOIN public.usuario u ON u.id = vo.usuario_id

            -- Restrições por voto
            LEFT JOIN LATERAL (
                SELECT jsonb_agg(re.*) AS restricoes_array
                FROM public.restricao re
                WHERE re.voto_id = vo.id
            ) restricoes ON true

            -- Voto contrário + parecer
            LEFT JOIN LATERAL (
                SELECT jsonb_agg(
                    jsonb_build_object(
                        'id', vc.id,
                        'id_texto', vc.id_texto,
                        'parecer_id', vc.parecer_id,
                        'voto_id', vc.voto_id,
                        'parecer', pa2.*
                    )
                ) AS voto_contrario_array
                FROM public.voto_contrario vc
                LEFT JOIN public.parecer pa2 ON pa2.id = vc.parecer_id
                WHERE vc.voto_id = vo.id
            ) vc ON true

            WHERE vo.votacao_id = v.id
        ) votos ON true
        WHERE v.projeto_id = p.id
    ) votacoes ON true

    WHERE p.reuniao_id = p_reuniao_id
) t;
$$;