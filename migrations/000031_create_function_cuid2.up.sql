CREATE OR REPLACE FUNCTION public.cuid2()
RETURNS text AS $$
DECLARE
  ts bigint;
  counter int;
  rand_bytes bytea;
  rand_num numeric;
  id text;
BEGIN
  -- timestamp em nanos
  ts := floor(extract(epoch from clock_timestamp()) * 1000000);

  -- contador monotônico (simulado via random)
  counter := trunc(random() * 1000000);

  -- parte aleatória
  rand_bytes := gen_random_bytes(8);
  rand_num := ('x' || encode(rand_bytes, 'hex'))::bit(64)::bigint;

  -- monta id em base36
  id := lpad(to_char(ts, 'FM99999999999999999999'), 12, '0')
        || lpad(to_char(counter, 'FM999999'), 6, '0')
        || lpad(to_char(rand_num, 'FM9999999999'), 6, '0');

  -- converte para base36
  id := lower(to_char(abs(hashtext(id)), 'FM999999999999999999999999999999'));

  -- garante 24 chars e começa com letra
  id := 'c' || substring(md5(id) from 1 for 23);

  RETURN id;
END;
$$ LANGUAGE plpgsql STRICT;