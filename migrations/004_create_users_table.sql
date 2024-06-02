CREATE TABLE IF NOT EXISTS users (
    id   SERIAL PRIMARY KEY,
    name TEXT
);

COMMENT ON TABLE users          IS 'Contains the users enabled to give a verdict on the tweets';
COMMENT ON COLUMN users.id      IS 'Auto-incrementing ID of the user, agnostic to business logic';
COMMENT ON COLUMN users.name    IS 'Name of the user';