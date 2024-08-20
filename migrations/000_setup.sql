-- Create a function to check if an index exists and create it if it does not
CREATE OR REPLACE FUNCTION create_index_if_not_exists(
    p_index_name TEXT,
    p_table_name TEXT,
    p_index_def TEXT
) RETURNS VOID AS $$
DECLARE
    index_exists BOOLEAN;
BEGIN
    -- Check if the index exists
    SELECT EXISTS (
        SELECT 1
        FROM pg_indexes
        WHERE schemaname = 'public'
          AND indexname = p_index_name
    ) INTO index_exists;

    -- Create the index if it does not exist
    IF NOT index_exists THEN
        EXECUTE format('CREATE INDEX %I ON %I (%s)', p_index_name, p_table_name, p_index_def);
    END IF;
END $$ LANGUAGE plpgsql;