package summary

import (
	"context"

	"github.com/jackc/pgx/v5"

	"ahbcc/internal/database"
	"ahbcc/internal/log"
)

// UpdateTotalTweets updates a counts.DAO
type UpdateTotalTweets func(tx pgx.Tx, ctx context.Context, id, totalTweets int) error

// MakeUpdateTotalTweets creates a new UpdateTotalTweets
func MakeUpdateTotalTweets(db database.Connection) UpdateTotalTweets {
	const query string = `
		UPDATE search_criteria_executions_summary
		SET total_tweets = $2
		WHERE id = $1
	`

	return func(tx pgx.Tx, ctx context.Context, id, totalTweets int) error {
		if tx != nil {
			db = tx
		}

		_, err := db.Exec(ctx, query, id, totalTweets)
		if err != nil {
			log.Error(ctx, err.Error())
			return FailedToUpdateTotalTweets
		}

		return nil
	}
}
