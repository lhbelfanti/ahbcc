-- Create the enum type for status
CREATE TYPE execution_status AS ENUM ('PENDING', 'IN PROGRESS', 'DONE');

-- Create the search_criteria_executions table
CREATE TABLE IF NOT EXISTS search_criteria_executions (
    id                  SERIAL PRIMARY KEY,
    status              execution_status NOT NULL,
    search_criteria_id  INTEGER NOT NULL,

    CONSTRAINT fk_search_criteria_id FOREIGN KEY(search_criteria_id) REFERENCES search_criteria(id)
);

-- Table indexes
CREATE INDEX idx_search_criteria_executions_search_criteria_id ON search_criteria_executions(search_criteria_id);

-- Table comments
COMMENT ON TABLE search_criteria_executions IS 'Tracks the execution status of each search criteria';
COMMENT ON COLUMN search_criteria_executions.id IS 'Auto-incrementing ID of the search criteria execution, agnostic to business logic';
COMMENT ON COLUMN search_criteria_executions.status IS 'Current status of the execution, can be PENDING, IN PROGRESS, or DONE';
COMMENT ON COLUMN search_criteria_executions.search_criteria_id IS 'Foreign key referencing the ID of the search criteria';
