-- ============================================================
-- Função: f_get_voting_stats(p_stats_date date)
-- Retorna o total de projetos e total de projetos votados no dia
-- ============================================================

CREATE OR REPLACE FUNCTION f_get_voting_stats(p_stats_date date)
RETURNS TABLE (
    totalProjects bigint,
    totalVotedProjects bigint
)
LANGUAGE plpgsql
STABLE
AS $$
BEGIN
    RETURN QUERY
    SELECT
        COALESCE(COUNT(DISTINCT p.id), 0) AS totalProjects,
        COALESCE(COUNT(DISTINCT CASE WHEN v.status = 'V' THEN v.projeto_id END), 0) AS totalVotedProjects
    FROM projeto p
    INNER JOIN reuniao r 
        ON p.reuniao_id = r.id
    LEFT JOIN votacao v 
        ON p.id = v.projeto_id
    WHERE r.rec_data = p_stats_date;
END;
$$;
