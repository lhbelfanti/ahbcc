package summary

import (
	"context"

	"ahbcc/internal/database"
	"ahbcc/internal/log"
)

// UpdateTotalTweets updates a counts.DAO
type UpdateTotalTweets func(ctx context.Context, id, totalTweets int) error

// MakeUpdateTotalTweets creates a new UpdateTotalTweets
func MakeUpdateTotalTweets(db database.Connection) UpdateTotalTweets {
	const query string = `
		UPDATE search_criteria_executions_summary
		SET total_tweets = $2
		WHERE id = $1
	`

	return func(ctx context.Context, id, totalTweets int) error {
		_, err := db.Exec(ctx, query, id, totalTweets)
		if err != nil {
			log.Error(ctx, err.Error())
			return FailedToUpdateTotalTweets
		}

		return nil
	}
}
