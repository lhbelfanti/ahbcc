package summary

import (
	"context"
	"errors"

	"github.com/jackc/pgx/v5"

	"ahbcc/internal/database"
	"ahbcc/internal/log"
)

// Insert inserts a new row into tweets_counts table
type Insert func(ctx context.Context, dao DAO) (int, error)

// MakeInsert creates a new Insert
func MakeInsert(db database.Connection) Insert {
	const query string = `
		INSERT INTO search_criteria_executions_summary (search_criteria_id, tweets_year, tweets_month, total_tweets)
		VALUES ($1, $2, $3, $4)
		RETURNING id;
	`

	return func(ctx context.Context, dao DAO) (int, error) {
		var tweetsCountsID int
		err := db.QueryRow(ctx, query, dao.SearchCriteriaID, dao.Year, dao.Month, dao.Total).Scan(&tweetsCountsID)
		if errors.Is(err, pgx.ErrNoRows) {
			log.Error(ctx, err.Error())
			return -1, FailedToInsertExecutionSummary
		}

		return tweetsCountsID, nil
	}
}
