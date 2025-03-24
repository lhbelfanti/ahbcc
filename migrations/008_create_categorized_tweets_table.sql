-- Create the categorized_tweets table
CREATE TABLE IF NOT EXISTS categorized_tweets (
    id                SERIAL PRIMARY KEY,
    tweet_id          INTEGER NOT NULL,
    adverse_behavior  BOOLEAN NOT NULL,

    CONSTRAINT fk_tweet_id FOREIGN KEY(tweet_id) REFERENCES tweets(uuid)
);

-- Table indexes
CREATE INDEX IF NOT EXISTS idx_categorized_tweets_tweet_id ON categorized_tweets(tweet_id);

-- Table comments
COMMENT ON TABLE categorized_tweets                     IS 'Contains the categorization of tweets for adverse behavior';
COMMENT ON COLUMN categorized_tweets.id                 IS 'Auto-incrementing ID of the categorization record, agnostic to business logic';
COMMENT ON COLUMN categorized_tweets.tweet_id           IS 'Foreign key referencing the ID of the tweet';
COMMENT ON COLUMN categorized_tweets.adverse_behavior   IS 'Boolean indicating if the tweet exhibits adverse behavior as determined by the user';