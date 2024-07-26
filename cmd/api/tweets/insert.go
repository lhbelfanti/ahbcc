package tweets

import (
	"context"
	"log/slog"

	"ahbcc/internal/database"
)

// InsertTweet inserts a new TweetDTO into tweets table
type InsertTweet func(tweet TweetDTO) error

// MakeInsertTweet creates a new InsertTweet
func MakeInsertTweet(db database.Connection) InsertTweet {
	const query string = `
		INSERT INTO tweets(hash, is_a_reply, has_text, has_images, text_content, images, has_quote, quote_id, search_criteria_id) 
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9)
	`

	return func(tweet TweetDTO) error {
		_, err := db.Exec(
			context.Background(),
			query,
			tweet.Hash,
			tweet.IsAReply,
			tweet.HasText,
			tweet.HasImages,
			tweet.TextContent,
			tweet.Images,
			tweet.HasQuote,
			tweet.QuoteID,
			tweet.SearchCriteriaID,
		)

		if err != nil {
			slog.Error(err.Error())
			return FailedToInsertTweet
		}

		return nil
	}
}
