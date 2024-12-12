-- Create the tweets table
CREATE TABLE IF NOT EXISTS tweets (
    id                  SERIAL PRIMARY KEY,
    uuid                TEXT NOT NULL,
    author              TEXT NOT NULL,
    avatar              TEXT,
    posted_at           TIMESTAMP WITH TIME ZONE,
    is_a_reply          BOOLEAN NOT NULL,
    text_content        TEXT NULL,
    images              TEXT[] NULL,
    quote_id            INTEGER NULL,
    search_criteria_id  INTEGER NOT NULL,

    CONSTRAINT uq_uuid_search_criteria UNIQUE (uuid, search_criteria_id),
    CONSTRAINT fk_quote_id FOREIGN KEY(quote_id) REFERENCES tweets_quotes(id),
    CONSTRAINT fk_search_criteria_id FOREIGN KEY(search_criteria_id) REFERENCES search_criteria(id)
);

-- Table indexes
SELECT create_index_if_not_exists('idx_tweets_quote_id', 'tweets', 'quote_id');
SELECT create_index_if_not_exists('idx_tweets_search_criteria', 'tweets', 'search_criteria_id');

-- Table comments
COMMENT ON TABLE tweets                         IS 'Contains the tweets scrapped by GoXCrap';
COMMENT ON COLUMN tweets.id                     IS 'Auto-incrementing ID of the tweet, agnostic to business logic';
COMMENT ON COLUMN tweets.uuid                   IS 'UUID identifier for the tweet. It is part of the primary key';
COMMENT ON COLUMN tweets.author                 IS 'The user that wrote the tweet';
COMMENT ON COLUMN tweets.avatar                 IS 'The user profile image';
COMMENT ON COLUMN tweets.posted_at              IS 'Timestamp indicating when the tweet was posted';
COMMENT ON COLUMN tweets.is_a_reply             IS 'Boolean indicating if the tweet is a reply to another tweet';
COMMENT ON COLUMN tweets.text_content           IS 'The text content of the tweet, if any';
COMMENT ON COLUMN tweets.images                 IS 'Array of image URLs associated with the tweet, if any';
COMMENT ON COLUMN tweets.quote_id               IS 'Foreign key referencing the ID of the quoted tweet';
COMMENT ON COLUMN tweets.search_criteria_id     IS 'Foreign key referencing the ID of the search criteria';