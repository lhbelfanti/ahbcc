-- Create the criteria_execution_days table
CREATE TABLE IF NOT EXISTS tweets_counts (
    id                           SERIAL PRIMARY KEY,
    search_criteria_id           INTEGER NOT NULL,
    tweets_year                  INTEGER NOT NULL,
    tweets_month                 INTEGER NOT NULL,
    total_tweets                 INTEGER NOT NULL,

    CONSTRAINT valid_month CHECK (tweets_month BETWEEN 1 AND 12),
    CONSTRAINT fk_search_criteria_id FOREIGN KEY(search_criteria_id) REFERENCES search_criteria(id)
);

-- Table indexes
CREATE INDEX IF NOT EXISTS idx_tweets_counts_search_criteria_id ON tweets_counts(search_criteria_id);
CREATE INDEX IF NOT EXISTS idx_tweets_counts_search_criteria_id_tweets_year ON tweets_counts(search_criteria_id, tweets_year);

-- Table comments
COMMENT ON TABLE tweets_counts IS 'Records the result of each search criteria execution in a summarized format';
COMMENT ON COLUMN tweets_counts.id IS 'Auto-incrementing ID of the tweets availability, agnostic to business logic';

COMMENT ON COLUMN tweets_counts.search_criteria_id IS 'ID of the search criteria from where the tweets where obtained';
COMMENT ON COLUMN tweets_counts.tweets_year IS 'Year when the tweets where published';
COMMENT ON COLUMN tweets_counts.tweets_month IS 'Month when the tweets where published';
COMMENT ON COLUMN tweets_counts.total_tweets IS 'The amount of tweets retrieved';
