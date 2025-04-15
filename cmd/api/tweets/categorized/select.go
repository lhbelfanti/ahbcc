package categorized

import (
	"context"

	"ahbcc/internal/database"
	"ahbcc/internal/log"
)

// SelectAllByUserID returns a struct with all the analyzed tweets divided by year and month
type SelectAllByUserID func(ctx context.Context, userID int) ([]AnalyzedTweetsDAO, error)

// MakeSelectAllByUserID creates a new SelectAllByUserID
func MakeSelectAllByUserID(db database.Connection, collectRows database.CollectRows[AnalyzedTweetsDAO]) SelectAllByUserID {
	const query string = `
		SELECT search_criteria_id, tweet_year, tweet_month, COUNT(*) AS analyzed_tweets
		FROM categorized_tweets
		WHERE user_id = $1
		GROUP BY search_criteria_id, tweet_year, tweet_month
		ORDER BY search_criteria_id, tweet_year, tweet_month;
	`

	return func(ctx context.Context, userID int) ([]AnalyzedTweetsDAO, error) {
		rows, err := db.Query(ctx, query, userID)
		if err != nil {
			log.Error(ctx, err.Error())
			return nil, FailedToExecuteSelectAllCategorizedTweetsByUserID
		}

		analyzedTweets, err := collectRows(rows)
		if err != nil {
			log.Error(ctx, err.Error())
			return nil, FailedToExecuteCollectRowsInSelectAllCategorizedTweetsByUserID
		}

		return analyzedTweets, nil
	}
}
