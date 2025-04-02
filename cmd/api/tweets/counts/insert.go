package counts

import (
	"context"

	"ahbcc/internal/database"
	"ahbcc/internal/log"
)

// Insert inserts a new row into tweets_counts table
type Insert func(ctx context.Context, tweetsCounts DAO) (int, error)

// MakeInsert creates a new Insert
func MakeInsert(db database.Connection) Insert {
	const query string = `
		INSERT INTO tweets_count (search_criteria_id, tweets_year, tweets_month, total_tweets)
		VALUES ($1, $2, $3, $4)
		RETURNING id;
	`

	return func(ctx context.Context, tweetsCounts DAO) (int, error) {
		var tweetsCountsID int
		err := db.QueryRow(ctx, query).Scan(&tweetsCountsID)
		if err != nil {
			log.Error(ctx, err.Error())
			return -1, FailedToInsertTweetsCounts
		}

		return tweetsCountsID, nil
	}
}
