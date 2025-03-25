CREATE OR REPLACE FUNCTION create_enum_type_if_not_exists(
    type_name TEXT,
    enum_values TEXT[]
) RETURNS VOID AS $$
DECLARE
    quoted_values TEXT;
BEGIN
    -- Quote each array element and join with commas
    SELECT string_agg(quote_literal(value), ', ')
    INTO quoted_values
    FROM unnest(enum_values) AS value;

    IF NOT EXISTS (SELECT 1 FROM pg_type WHERE typname = type_name) THEN
        EXECUTE format('CREATE TYPE %I AS ENUM (%s)',
                       type_name,
                       quoted_values);
    END IF;
END;
$$ LANGUAGE plpgsql;