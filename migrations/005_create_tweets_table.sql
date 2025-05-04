-- Create the tweets table
CREATE TABLE IF NOT EXISTS tweets (
    id                  SERIAL PRIMARY KEY,
    status_id           TEXT NOT NULL,
    author              TEXT NOT NULL,
    avatar              TEXT,
    posted_at           TIMESTAMP WITH TIME ZONE,
    is_a_reply          BOOLEAN NOT NULL,
    text_content        TEXT NULL,
    images              TEXT[] NULL,
    quote_id            INTEGER NULL,
    search_criteria_id  INTEGER NOT NULL,

    CONSTRAINT uq_id_posted_at_search_criteria UNIQUE (status_id, posted_at, search_criteria_id),
    CONSTRAINT fk_quote_id FOREIGN KEY(quote_id) REFERENCES tweets_quotes(id),
    CONSTRAINT fk_search_criteria_id FOREIGN KEY(search_criteria_id) REFERENCES search_criteria(id)
);

-- Table indexes
CREATE INDEX IF NOT EXISTS idx_tweets_quote_id ON tweets(quote_id);
CREATE INDEX IF NOT EXISTS idx_tweets_search_criteria_id ON tweets(search_criteria_id);
CREATE INDEX IF NOT EXISTS idx_tweets_criteria_posted ON tweets(search_criteria_id, posted_at);

-- Table comments
COMMENT ON TABLE tweets                         IS 'Contains the tweets scrapped by GoXCrap';
COMMENT ON COLUMN tweets.id                     IS 'Auto-incrementing id of the tweet, agnostic to business logic';
COMMENT ON COLUMN tweets.status_id              IS 'The number after the /status/ of the tweet url';
COMMENT ON COLUMN tweets.author                 IS 'The user that wrote the tweet';
COMMENT ON COLUMN tweets.avatar                 IS 'The user profile image';
COMMENT ON COLUMN tweets.posted_at              IS 'Timestamp indicating when the tweet was posted';
COMMENT ON COLUMN tweets.is_a_reply             IS 'Boolean indicating if the tweet is a reply to another tweet';
COMMENT ON COLUMN tweets.text_content           IS 'The text content of the tweet, if any';
COMMENT ON COLUMN tweets.images                 IS 'Array of image URLs associated with the tweet, if any';
COMMENT ON COLUMN tweets.quote_id               IS 'Foreign key referencing the ID of the quoted tweet';
COMMENT ON COLUMN tweets.search_criteria_id     IS 'Foreign key referencing the ID of the search criteria';