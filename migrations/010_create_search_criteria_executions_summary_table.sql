-- Create the search_criteria_executions_summary table
CREATE TABLE IF NOT EXISTS search_criteria_executions_summary (
    id                  SERIAL PRIMARY KEY,
    search_criteria_id  INTEGER NOT NULL,
    tweets_year         INTEGER NOT NULL,
    tweets_month        INTEGER NOT NULL,
    total_tweets        INTEGER NOT NULL,

    CONSTRAINT search_criteria_executions_summary_valid_month CHECK (tweets_month BETWEEN 1 AND 12),
    CONSTRAINT fk_search_criteria_id FOREIGN KEY(search_criteria_id) REFERENCES search_criteria(id)
);

-- Table indexes
CREATE INDEX IF NOT EXISTS idx_executions_summary_search_criteria_id ON search_criteria_executions_summary(search_criteria_id);
CREATE INDEX IF NOT EXISTS idx_executions_summary_search_criteria_id_tweets_year ON search_criteria_executions_summary(search_criteria_id, tweets_year);

-- Table comments
COMMENT ON TABLE search_criteria_executions_summary IS 'Records the result of each search criteria execution in a summarized format';
COMMENT ON COLUMN search_criteria_executions_summary.id IS 'Auto-incrementing ID of the search criteria execution summary, agnostic to business logic';
COMMENT ON COLUMN search_criteria_executions_summary.search_criteria_id IS 'ID of the search criteria from where the tweets where obtained';
COMMENT ON COLUMN search_criteria_executions_summary.tweets_year IS 'Year when the tweets were published';
COMMENT ON COLUMN search_criteria_executions_summary.tweets_month IS 'Month when the tweets were published';
COMMENT ON COLUMN search_criteria_executions_summary.total_tweets IS 'The amount of tweets retrieved';
