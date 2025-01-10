-- Create the tweets table
CREATE TABLE IF NOT EXISTS users_sessions (
    id          SERIAL PRIMARY KEY,
    user_id     INT,
    token TEXT  NOT NULL,
    expires_at  TIMESTAMP NOT NULL,
    created_at  TIMESTAMP DEFAULT CURRENT_TIMESTAMP,

    CONSTRAINT fk_user_id FOREIGN KEY(user_id) REFERENCES users(id) ON DELETE CASCADE
);

-- Table indexes
SELECT create_index_if_not_exists('idx_users_sessions_user_id', 'users_sessions', 'user_id');
SELECT create_index_if_not_exists('idx_users_sessions_token', 'users_sessions', 'token');

-- Table comments
COMMENT ON TABLE users_sessions             IS 'Stores active sessions with token information';
COMMENT ON COLUMN users_sessions.id         IS 'Auto-incrementing ID of the session, agnostic to business logic';
COMMENT ON COLUMN users_sessions.user_id    IS 'Foreign key referencing users table';
COMMENT ON COLUMN users_sessions.token      IS 'Bearer token for session authentication';
COMMENT ON COLUMN users_sessions.expires_at IS 'Expiration time of the session token';
COMMENT ON COLUMN users_sessions.created_at IS 'Timestamp of when the session was created';