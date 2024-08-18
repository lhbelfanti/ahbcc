-- Create the search_criteria table
CREATE TABLE IF NOT EXISTS search_criteria (
    id                    SERIAL PRIMARY KEY,
    name                  TEXT NOT NULL,
    all_of_these_words    TEXT[] NULL,
    this_exact_phrase     TEXT NULL,
    any_of_these_words    TEXT[] NULL,
    none_of_these_words   TEXT[] NULL,
    these_hashtags        TEXT[] NULL,
    language              TEXT NOT NULL,
    since_date            DATE NOT NULL,
    until_date            DATE NOT NULL
);

-- Table comments
COMMENT ON TABLE search_criteria                        IS 'Contains the criteria used for searching tweets';
COMMENT ON COLUMN search_criteria.id                    IS 'Auto-incrementing ID of the search criteria, agnostic to business logic';
COMMENT ON COLUMN search_criteria.name                  IS 'Name of the search criteria for easy identification';
COMMENT ON COLUMN search_criteria.all_of_these_words    IS 'Array of words that must all be present in the tweet';
COMMENT ON COLUMN search_criteria.this_exact_phrase     IS 'Exact phrase that must be present in the tweet';
COMMENT ON COLUMN search_criteria.any_of_these_words    IS 'Array of words, any of which can be present in the tweet';
COMMENT ON COLUMN search_criteria.none_of_these_words   IS 'Array of words that must not be present in the tweet';
COMMENT ON COLUMN search_criteria.these_hashtags        IS 'Array of hashtags that must be present in the tweet';
COMMENT ON COLUMN search_criteria.language              IS 'Language of the tweet';
COMMENT ON COLUMN search_criteria.since_date            IS 'Date from which to start the search';
COMMENT ON COLUMN search_criteria.until_date            IS 'Date until which to perform the search';
