-- Create the criteria_execution_days table
CREATE TABLE IF NOT EXISTS search_criteria_execution_days (
    id                              SERIAL PRIMARY KEY,
    execution_date                  DATE NOT NULL,
    tweets_quantity                 INTEGER,
    error_reason                    TEXT,
    search_criteria_execution_id    INTEGER NOT NULL,

    CONSTRAINT fk_search_criteria_executions_days_search_criteria_execution_id FOREIGN KEY(search_criteria_execution_id) REFERENCES search_criteria_executions(id)
);

-- Table indexes
CREATE INDEX IF NOT EXISTS idx_search_criteria_execution_days_execution_date ON search_criteria_execution_days(execution_date);
CREATE INDEX IF NOT EXISTS idx_search_criteria_execution_days_execution_date ON search_criteria_execution_days(search_criteria_execution_id);

-- Table comments
COMMENT ON TABLE search_criteria_execution_days IS 'Records daily results of each criteria execution';
COMMENT ON COLUMN search_criteria_execution_days.id IS 'Auto-incrementing ID of the search criteria execution day, agnostic to business logic';
COMMENT ON COLUMN search_criteria_execution_days.execution_date IS 'Date in which the search criteria was executed';
COMMENT ON COLUMN search_criteria_execution_days.tweets_quantity IS 'Number of tweets obtained during the execution on this day';
COMMENT ON COLUMN search_criteria_execution_days.error_reason IS 'Error message if any issues occurred during the execution on this day';
COMMENT ON COLUMN search_criteria_execution_days.search_criteria_execution_id IS 'Foreign key linking to the criteria_execution table';
