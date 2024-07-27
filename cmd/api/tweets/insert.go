package tweets

import (
	"context"
	"fmt"
	"log/slog"
	"strings"

	"ahbcc/internal/database"
)

// Insert inserts a new TweetDTO into tweets table
type Insert func(tweet []TweetDTO) error

// MakeInsert creates a new Insert
func MakeInsert(db database.Connection) Insert {
	const query string = `INSERT INTO tweets(hash, is_a_reply, has_text, has_images, text_content, images, has_quote, quote_id, search_criteria_id) VALUES `

	return func(tweets []TweetDTO) error {
		var valueStrings []string
		var values []any
		for i, tweet := range tweets {
			valueStrings = append(valueStrings, fmt.Sprintf("($%d, $%d, $%d, $%d, $%d, $%d, $%d, $%d, $%d)", i*9+1, i*9+2, i*9+3, i*9+4, i*9+5, i*9+6, i*9+7, i*9+8, i*9+9))
			values = append(values, tweet.Hash, tweet.IsAReply, tweet.HasText, tweet.HasImages, tweet.TextContent, tweet.Images, tweet.HasQuote, tweet.QuoteID, tweet.SearchCriteriaID)
		}

		query := query + strings.Join(valueStrings, ",")

		_, err := db.Exec(context.Background(), query, values...)
		if err != nil {
			slog.Error(err.Error())
			return FailedToInsertTweets
		}

		return nil
	}
}
