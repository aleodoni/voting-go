CREATE OR REPLACE PROCEDURE reset_staging()
LANGUAGE plpgsql
AS $$
BEGIN
  DELETE FROM voto_contrario;
  DELETE FROM restricao;
  DELETE FROM voto;
  DELETE FROM votacao;
  UPDATE reuniao SET rec_data = NOW()::date, updated_at = NOW();
END;
$$;

SELECT cron.schedule('reset-staging', '0 23 * * *', 'CALL reset_staging()')
WHERE NOT EXISTS (
  SELECT 1 FROM cron.job WHERE jobname = 'reset-staging'
);