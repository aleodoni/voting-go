CREATE VIEW v_usuario_permissoes AS
    SELECT 
      u.id,
      u.keycloak_id,
      u.email,
      u.nome,
      u.nome_fantasia,
      u.created_at,
      u.updated_at,
      p.id AS permissoes_id,
      p.ativo,
      p.pode_administrar,
      p.pode_votar,
      p.created_at AS permissoes_created_at,
      p.updated_at AS permissoes_updated_at,
      p.usuario_id
    FROM usuario u
    LEFT JOIN credencial p ON p.usuario_id = u.id;