DO $$
BEGIN
    IF NOT EXISTS (
        SELECT 1
        FROM pg_type
        WHERE typname = 'opcao_voto'
    ) THEN
        CREATE TYPE opcao_voto AS ENUM ('F', 'R', 'C', 'V', 'A');
    END IF;
END
$$;