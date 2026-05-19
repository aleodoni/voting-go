DO $$
BEGIN
  IF NOT EXISTS (
    SELECT 1 FROM pg_foreign_server WHERE srvname = 'server_spl'
  ) THEN
    CREATE SERVER server_spl
      FOREIGN DATA WRAPPER postgres_fdw
      OPTIONS (
        host '${DB_SPL_HOST}',
        dbname '${DB_SPL_NAME}',
        port '${DB_SPL_PORT}',
        sslmode 'disable'
      );
  END IF;
END
$$;