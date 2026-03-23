-- ============================================================
-- Função: f_get_voting_stats(p_stats_date date)
-- Retorna o total de projetos e total de projetos votados no dia
-- ============================================================

CREATE OR REPLACE FUNCTION f_get_voting_stats(p_stats_date date)
RETURNS json
LANGUAGE plpgsql
STABLE
AS $$
DECLARE
    result json;
BEGIN
    SELECT json_build_object(
        'totalProjects',      COALESCE(COUNT(DISTINCT p.id), 0),
        'totalVotedProjects', COALESCE(COUNT(DISTINCT CASE WHEN v.status = 'V' THEN v.projeto_id END), 0)
    )
    INTO result
    FROM projeto p
    INNER JOIN reuniao r 
        ON p.reuniao_id = r.id
    LEFT JOIN votacao v 
        ON p.id = v.projeto_id
    WHERE r.rec_data = p_stats_date;

    RETURN result;
END;
$$;