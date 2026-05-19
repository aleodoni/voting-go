DO $$
BEGIN
    IF NOT EXISTS (
        SELECT 1
        FROM pg_type
        WHERE typname = 'status_votacao'
    ) THEN
        CREATE TYPE status_votacao AS ENUM ('A', 'F', 'V', 'C');
    END IF;
END
$$;