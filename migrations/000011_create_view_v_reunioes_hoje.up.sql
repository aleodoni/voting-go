-- ===============================================
-- View: v_reunioes_hoje (idempotente)
-- ===============================================
CREATE OR REPLACE VIEW public.v_reunioes_hoje AS
SELECT
  *
FROM
  public.reuniao
WHERE
  rec_data = current_date;
