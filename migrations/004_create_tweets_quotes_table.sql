-- Create the tweets_quotes table
CREATE TABLE IF NOT EXISTS tweets_quotes (
    id            SERIAL PRIMARY KEY,
    author        TEXT NOT NULL,
    avatar        TEXT,
    posted_at     TIMESTAMP WITH TIME ZONE,
    is_a_reply    BOOLEAN NOT NULL,
    text_content  TEXT NULL,
    images        TEXT[] NULL
);

-- Table comments
COMMENT ON TABLE tweets_quotes                IS 'Contains the quotes of tweets scrapped by GoXCrap';
COMMENT ON COLUMN tweets_quotes.id            IS 'Auto-incrementing ID of the quote, agnostic to business logic';
COMMENT ON COLUMN tweets_quotes.author        IS 'The user that wrote the tweet';
COMMENT ON COLUMN tweets_quotes.avatar        IS 'The user profile image';
COMMENT ON COLUMN tweets_quotes.posted_at     IS 'Timestamp indicating when the quoted tweet was posted';
COMMENT ON COLUMN tweets_quotes.is_a_reply    IS 'Boolean indicating if the quoted tweet is a reply to another tweet';
COMMENT ON COLUMN tweets_quotes.text_content  IS 'Text content of the quoted tweet, if any';
COMMENT ON COLUMN tweets_quotes.images        IS 'Array of image URLs associated with the quoted tweet, if any';