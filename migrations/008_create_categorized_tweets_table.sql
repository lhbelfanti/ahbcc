-- Create the enum type for categorization
SELECT create_enum_type_if_not_exists('verdict', ARRAY['POSITIVE', 'INDETERMINATE', 'NEGATIVE']);

-- Create the categorized_tweets table
CREATE TABLE IF NOT EXISTS categorized_tweets (
    id                  SERIAL PRIMARY KEY,
    search_criteria_id  INTEGER NOT NULL,
    tweet_id            INTEGER NOT NULL,
    tweet_year          INTEGER NOT NULL,
    tweet_month         INTEGER NOT NULL,
    user_id             INTEGER NOT NULL,
    categorization      verdict NOT NULL,

    CONSTRAINT fk_search_criteria_id FOREIGN KEY(search_criteria_id) REFERENCES search_criteria(id),
    CONSTRAINT fk_tweet_id FOREIGN KEY(tweet_id) REFERENCES tweets(uuid),
    CONSTRAINT fk_user_id FOREIGN KEY(user_id) REFERENCES users(id)
);

-- Table indexes
CREATE INDEX IF NOT EXISTS idx_categorized_tweets_tweet_id ON categorized_tweets(tweet_id);
CREATE INDEX IF NOT EXISTS idx_categorized_tweets_user_id ON categorized_tweets(user_id);
CREATE INDEX IF NOT EXISTS idx_categorized_tweets_user_id_tweet_year ON categorized_tweets(user_id, tweet_year);


-- Table comments
COMMENT ON TABLE categorized_tweets                     IS 'Contains the categorization of tweets by user for adverse behavior';
COMMENT ON COLUMN categorized_tweets.id                 IS 'Auto-incrementing ID of the categorization record, agnostic to business logic';
COMMENT ON COLUMN categorized_tweets.search_criteria_id IS 'ID of the search criteria from where the tweets where obtained. This field can also be obtained from the tweet itself, but it was added in this table for query optimization reasons';
COMMENT ON COLUMN categorized_tweets.tweet_id           IS 'Foreign key referencing the ID of the tweet';
COMMENT ON COLUMN categorized_tweets.tweet_year         IS 'Year when the tweet was published. This field can also be obtained from the tweet itself, but it was added in this table for query optimization reasons';
COMMENT ON COLUMN categorized_tweets.tweet_month        IS 'Month when the tweet was published. This field can also be obtained from the tweet itself, but it was added in this table for query optimization reasons';
COMMENT ON COLUMN categorized_tweets.user_id            IS 'Foreign key referencing the ID of the user who categorized the tweet';
COMMENT ON COLUMN categorized_tweets.categorization     IS 'Indicates the categorization verdict. It can be POSITIVE, INDETERMINATE or NEGATIVE';