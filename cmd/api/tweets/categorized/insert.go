package categorized

import (
	"context"

	"github.com/lhbelfanti/corpus-creator/internal/database"
	"github.com/lhbelfanti/corpus-creator/internal/log"
)

// InsertSingle inserts a new categorized tweet DTO into 'categorized_tweets' table and returns the ID
type InsertSingle func(ctx context.Context, dto DTO) (int, error)

// MakeInsertSingle creates a new InsertSingle
func MakeInsertSingle(db database.Connection) InsertSingle {
	const query string = `
		INSERT INTO categorized_tweets(search_criteria_id, tweet_id, tweet_year, tweet_month, user_id, categorization) 
		VALUES ($1, $2, $3, $4, $5, $6)
		RETURNING id;
	`

	return func(ctx context.Context, dto DTO) (int, error) {
		var categorizedTweetID int

		err := db.QueryRow(
			ctx,
			query,
			dto.SearchCriteriaID,
			dto.TweetID,
			dto.TweetYear,
			dto.TweetMonth,
			dto.UserID,
			dto.Categorization,
		).Scan(&categorizedTweetID)
		if err != nil {
			log.Error(ctx, err.Error())
			return -1, FailedToExecuteInsertCategorizedTweet
		}

		return categorizedTweetID, nil
	}
}
