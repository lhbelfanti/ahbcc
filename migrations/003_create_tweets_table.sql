CREATE TABLE IF NOT EXISTS tweets (
    id                  SERIAL PRIMARY KEY,
    hash                TEXT NOT NULL,
    posted_at           TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    is_a_reply          BOOLEAN NOT NULL,
    text_content        TEXT NULL,
    images              TEXT[] NULL,
    quote_id            INTEGER NULL,
    search_criteria_id  INTEGER NOT NULL,

    CONSTRAINT uq_hash_search_criteria UNIQUE (hash, search_criteria_id),
    CONSTRAINT fk_quote_id FOREIGN KEY(quote_id) REFERENCES tweets_quotes(id),
    CONSTRAINT fk_search_criteria_id FOREIGN KEY(search_criteria_id) REFERENCES search_criteria(id)
);

CREATE INDEX idx_tweets_quote_id ON tweets(quote_id);
CREATE INDEX idx_tweets_search_criteria ON tweets(search_criteria_id);

COMMENT ON TABLE tweets                         IS 'Contains the tweets scrapped by GoXCrap';
COMMENT ON COLUMN tweets.id                     IS 'Auto-incrementing ID of the tweet, agnostic to business logic';
COMMENT ON COLUMN tweets.hash                   IS 'Unique hash identifier for the tweet. It is part of the primary key';
COMMENT ON COLUMN tweets.posted_at              IS 'Timestamp indicating when the tweet was posted';
COMMENT ON COLUMN tweets.is_a_reply             IS 'Boolean indicating if the tweet is a reply to another tweet';
COMMENT ON COLUMN tweets.text_content           IS 'The text content of the tweet, if any';
COMMENT ON COLUMN tweets.images                 IS 'Array of image URLs associated with the tweet, if any';
COMMENT ON COLUMN tweets.quote_id               IS 'Foreign key referencing the ID of the quoted tweet. It is also part of the primary key';
COMMENT ON COLUMN tweets.search_criteria_id     IS 'Foreign key referencing the ID of the search criteria';