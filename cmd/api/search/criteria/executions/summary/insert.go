package summary

import (
	"context"
	"errors"

	"github.com/jackc/pgx/v5"

	"github.com/lhbelfanti/corpus-creator/internal/database"
	"github.com/lhbelfanti/corpus-creator/internal/log"
)

// Insert inserts a new row into search_criteria_executions_summary table
type Insert func(tx pgx.Tx, ctx context.Context, dao DAO) (int, error)

// MakeInsert creates a new Insert
func MakeInsert(db database.Connection) Insert {
	const query string = `
		INSERT INTO search_criteria_executions_summary (search_criteria_id, tweets_year, tweets_month, total_tweets)
		VALUES ($1, $2, $3, $4)
		RETURNING id;
	`

	return func(tx pgx.Tx, ctx context.Context, dao DAO) (int, error) {
		if tx != nil {
			db = tx
		}

		var tweetsCountsID int
		err := db.QueryRow(ctx, query, dao.SearchCriteriaID, dao.Year, dao.Month, dao.Total).Scan(&tweetsCountsID)
		if errors.Is(err, pgx.ErrNoRows) {
			log.Error(ctx, err.Error())
			return -1, FailedToInsertExecutionSummary
		}

		return tweetsCountsID, nil
	}
}
