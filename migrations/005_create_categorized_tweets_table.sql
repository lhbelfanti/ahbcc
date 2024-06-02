CREATE TABLE IF NOT EXISTS categorized_tweets (
    id                SERIAL PRIMARY KEY,
    tweet_id          INTEGER NOT NULL,
    user_id           INTEGER NOT NULL,
    adverse_behavior  BOOLEAN,
    CONSTRAINT fk_tweet_id FOREIGN KEY(tweet_id) REFERENCES tweets(id),
    CONSTRAINT fk_user_id FOREIGN KEY(user_id) REFERENCES users(id)
);

COMMENT ON TABLE categorized_tweets                     IS 'Contains the categorization of tweets by users for adverse behavior';
COMMENT ON COLUMN categorized_tweets.id                 IS 'Auto-incrementing ID of the categorization record, agnostic to business logic';
COMMENT ON COLUMN categorized_tweets.tweet_id           IS 'Foreign key referencing the ID of the tweet';
COMMENT ON COLUMN categorized_tweets.user_id            IS 'Foreign key referencing the ID of the user who categorized the tweet';
COMMENT ON COLUMN categorized_tweets.adverse_behavior   IS 'Boolean indicating if the tweet exhibits adverse behavior as determined by the user';