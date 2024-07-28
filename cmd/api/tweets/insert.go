package tweets

import (
	"context"
	"fmt"
	"log/slog"
	"strings"

	"ahbcc/cmd/api/tweets/quotes"
	"ahbcc/internal/database"
)

// Insert inserts a new TweetDTO into 'tweets' table
type Insert func(tweet []TweetDTO) error

// MakeInsert creates a new Insert
func MakeInsert(db database.Connection, insertQuote quotes.InsertSingle) Insert {
	const (
		query string = `
			INSERT INTO tweets(hash, is_a_reply, text_content, images, quote_id, search_criteria_id) 
			VALUES %s
		    ON CONFLICT (hash, search_criteria_id) DO NOTHING;
		`
		parameters = 6
	)

	return func(tweets []TweetDTO) error {
		var valueStrings []string
		var values []any
		for i, tweet := range tweets {
			idx := i * parameters
			valueStrings = append(valueStrings, fmt.Sprintf("($%d, $%d, $%d, $%d, $%d, $%d)", idx+1, idx+2, idx+3, idx+4, idx+5, idx+6))
			values = append(values, tweet.Hash, tweet.IsAReply, tweet.TextContent, tweet.Images)

			if tweet.Quote != nil {
				quoteID, err := insertQuote(*tweet.Quote)
				if err != nil {
					values = append(values, nil)
				} else {
					values = append(values, quoteID)
				}
			}

			values = append(values, tweet.SearchCriteriaID)
		}

		queryToExecute := fmt.Sprintf(query, strings.Join(valueStrings, ","))

		_, err := db.Exec(context.Background(), queryToExecute, values...)
		if err != nil {
			slog.Error(err.Error())
			return FailedToInsertTweets
		}

		return nil
	}
}
