-- Create the users table
CREATE TABLE IF NOT EXISTS users (
    id              SERIAL PRIMARY KEY,
    username        TEXT NOT NULL,
    password_hash   TEXT NOT NULL,
    created_at      TIMESTAMP DEFAULT CURRENT_TIMESTAMP,

    CONSTRAINT uq_username_password_hash UNIQUE (username, password_hash)
);

-- Table indexes
SELECT create_index_if_not_exists('idx_users_username', 'users', 'username');

-- Table comments
COMMENT ON TABLE users                  IS 'Stores user credentials and metadata';
COMMENT ON COLUMN users.id              IS 'Auto-incrementing ID of the user, agnostic to business logic';
COMMENT ON COLUMN users.username        IS 'Username, must be unique';
COMMENT ON COLUMN users.password_hash   IS 'Hashed password for authentication. Must be unique';
COMMENT ON COLUMN users.created_at      IS 'Timestamp of when the user was created';