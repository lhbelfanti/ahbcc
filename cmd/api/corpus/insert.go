package corpus

import (
	"context"

	"ahbcc/internal/database"
	"ahbcc/internal/log"
)

// Insert inserts a new entry in the corpus table
type Insert func(ctx context.Context, entry DTO) error

// MakeInsert creates a new Insert function
func MakeInsert(db database.Connection) Insert {
	const query string = `INSERT INTO corpus (
		tweet_author, tweet_avatar, tweet_text, tweet_images, is_tweet_a_reply,
		quote_author, quote_avatar, quote_text, quote_images, is_quote_a_reply
	) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10)`

	return func(ctx context.Context, entry DTO) error {
		_, err := db.Exec(
			ctx,
			query,
			entry.TweetAuthor,
			entry.TweetAvatar,
			entry.TweetText,
			entry.TweetImages,
			entry.IsTweetAReply,
			entry.QuoteAuthor,
			entry.QuoteAvatar,
			entry.QuoteText,
			entry.QuoteImages,
			entry.IsQuoteAReply,
		)
		if err != nil {
			log.Error(ctx, err.Error())
			return FailedToInsertCorpusEntry
		}

		return nil
	}
}
