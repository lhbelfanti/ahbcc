CREATE TABLE IF NOT EXISTS tweets_quotes (
    id                  SERIAL PRIMARY KEY,
    quote_is_a_reply    BOOLEAN,
    quote_has_text      BOOLEAN,
    quote_has_images    BOOLEAN,
    quote_text          TEXT,
    quote_images        TEXT[]
);

COMMENT ON TABLE tweets_quotes                      IS 'Contains the quotes of tweets scrapped by GoXCrap';
COMMENT ON COLUMN tweets_quotes.id                  IS 'Auto-incrementing ID of the quote, agnostic to business logic';
COMMENT ON COLUMN tweets_quotes.quote_is_a_reply    IS 'Boolean indicating if the quoted tweet is a reply to another tweet';
COMMENT ON COLUMN tweets_quotes.quote_has_text      IS 'Boolean indicating if the quoted tweet contains text';
COMMENT ON COLUMN tweets_quotes.quote_has_images    IS 'Boolean indicating if the quoted tweet contains images';
COMMENT ON COLUMN tweets_quotes.quote_text          IS 'Text content of the quoted tweet, if any';
COMMENT ON COLUMN tweets_quotes.quote_images        IS 'Array of image URLs associated with the quoted tweet, if any';