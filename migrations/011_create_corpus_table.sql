-- Create the corpus table
CREATE TABLE IF NOT EXISTS corpus (
    id                  SERIAL PRIMARY KEY,
    tweet_author        TEXT NOT NULL,
    tweet_avatar        TEXT,
    tweet_text          TEXT NULL,
    tweet_images        TEXT[] NULL,
    is_tweet_a_reply    BOOLEAN NOT NULL,
    quote_author        TEXT NULL,
    quote_avatar        TEXT NULL,
    quote_text          TEXT NULL,
    quote_images        TEXT[] NULL,
    is_quote_a_reply    BOOLEAN NULL
);

-- Table comments
COMMENT ON TABLE corpus                     IS 'Records the corpus generated after the tweets categorization';
COMMENT ON COLUMN corpus.id                 IS 'Auto-incrementing id of the entry of the corpus, agnostic to business logic';
COMMENT ON COLUMN corpus.tweet_author       IS 'The user that wrote the tweet';
COMMENT ON COLUMN corpus.tweet_avatar       IS 'The user profile image';
COMMENT ON COLUMN corpus.tweet_text         IS 'The text content of the tweet, if any';
COMMENT ON COLUMN corpus.tweet_images       IS 'Array of image URLs associated with the tweet, if any';
COMMENT ON COLUMN corpus.is_tweet_a_reply   IS 'Boolean indicating if the tweet is a reply to another tweet';
COMMENT ON COLUMN corpus.quote_author       IS 'The quote''s user that wrote the tweet, if any';
COMMENT ON COLUMN corpus.quote_avatar       IS 'The quote''s user profile image, if any';
COMMENT ON COLUMN corpus.quote_text         IS 'The quote''s text content, if any';
COMMENT ON COLUMN corpus.quote_images       IS 'Array of image URLs associated with the quote, if any';
COMMENT ON COLUMN corpus.is_quote_a_reply   IS 'Boolean indicating if the quote is a reply to another tweet, if the quote exists';

