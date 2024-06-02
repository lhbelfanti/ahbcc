CREATE TABLE IF NOT EXISTS tweets (
    id          SERIAL PRIMARY KEY,
    hash        TEXT NOT NULL,
    posted_at   TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    is_a_reply  BOOLEAN,
    has_text    BOOLEAN,
    has_images  BOOLEAN,
    text        TEXT,
    images      TEXT[],
    has_quote   BOOLEAN,
    quote_id    INTEGER,
    CONSTRAINT fk_quote_id FOREIGN KEY(quote_id) REFERENCES tweets_quotes(id)
);

COMMENT ON TABLE tweets             IS 'Contains the tweets scrapped by GoXCrap';
COMMENT ON COLUMN tweets.id         IS 'Auto-incrementing ID of the tweet, agnostic to business logic';
COMMENT ON COLUMN tweets.hash       IS 'Unique hash identifier for the tweet';
COMMENT ON COLUMN tweets.posted_at  IS 'Timestamp indicating when the tweet was posted';
COMMENT ON COLUMN tweets.is_a_reply IS 'Boolean indicating if the tweet is a reply to another tweet';
COMMENT ON COLUMN tweets.has_text   IS 'Boolean indicating if the tweet contains text';
COMMENT ON COLUMN tweets.has_images IS 'Boolean indicating if the tweet contains images';
COMMENT ON COLUMN tweets.text       IS 'Text content of the tweet, if any';
COMMENT ON COLUMN tweets.images     IS 'Array of image URLs associated with the tweet, if any';
COMMENT ON COLUMN tweets.has_quote  IS 'Boolean indicating if the tweet contains a quoted tweet';
COMMENT ON COLUMN tweets.quote_id   IS 'Foreign key referencing the ID of the quoted tweet';