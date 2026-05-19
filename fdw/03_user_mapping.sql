DO $$
BEGIN
  IF NOT EXISTS (
    SELECT 1
    FROM pg_user_mappings
    WHERE srvname = 'server_spl'
      AND umuser = (SELECT usesysid FROM pg_user WHERE usename = 'postgres')
  ) THEN
    CREATE USER MAPPING FOR postgres
      SERVER server_spl
      OPTIONS (
        user '${DB_SPL_USER}',
        password '${DB_SPL_PASSWORD}'
      );
  END IF;
END
$$;