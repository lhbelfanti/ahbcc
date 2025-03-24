CREATE OR REPLACE FUNCTION create_enum_type_if_not_exists(
    type_name TEXT,
    enum_values TEXT[]
) RETURNS VOID AS $$
BEGIN
    IF NOT EXISTS (SELECT 1 FROM pg_type WHERE typname = type_name) THEN
        EXECUTE format('CREATE TYPE %I AS ENUM (%s)',
                      type_name,
                      array_to_string(enum_values, ', '));
    END IF;
END;
$$ LANGUAGE plpgsql;