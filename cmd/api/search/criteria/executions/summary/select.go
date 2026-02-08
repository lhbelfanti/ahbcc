package summary

import (
	"context"

	"github.com/lhbelfanti/corpus-creator/internal/database"
	"github.com/lhbelfanti/corpus-creator/internal/log"
)

type (
	// SelectMonthlyTweetsCountsByYearByCriteriaID returns the count of all the tweets (using the `tweets` table) for each year and month, seeking by search criteria ID
	SelectMonthlyTweetsCountsByYearByCriteriaID func(ctx context.Context, criteriaID int) ([]DAO, error)

	// SelectAll returns the summarization of the tweets retrieved for each month and year, for all the criteria
	SelectAll func(ctx context.Context) ([]DAO, error)
)

// MakeSelectMonthlyTweetsCountsByYearByCriteriaID creates a new SelectMonthlyTweetCountByYearByCriteriaID
func MakeSelectMonthlyTweetsCountsByYearByCriteriaID(db database.Connection, collectRows database.CollectRows[DAO]) SelectMonthlyTweetsCountsByYearByCriteriaID {
	const query string = `
		SELECT 
		    search_criteria_id,
			EXTRACT(YEAR FROM posted_at) AS tweets_year,
			LPAD(EXTRACT(MONTH FROM posted_at)::text, 2, '0')::int AS tweets_month,
			COUNT(*) AS total
		FROM 
			tweets
		WHERE 
			search_criteria_id = $1
		GROUP BY
		    search_criteria_id,
			tweets_year,
			tweets_month
		ORDER BY 
			tweets_year DESC,
			tweets_month DESC;
	`

	return func(ctx context.Context, criteriaID int) ([]DAO, error) {
		rows, err := db.Query(ctx, query, criteriaID)
		if err != nil {
			log.Error(ctx, err.Error())
			return nil, FailedToRetrieveMonthlyTweetsCountsByYear
		}

		tweetsCounts, err := collectRows(rows)
		if err != nil {
			log.Error(ctx, err.Error())
			return nil, FailedToExecuteCollectRowsInSelectMonthlyTweetsCountsByYear
		}

		return tweetsCounts, nil
	}
}

// MakeSelectAll creates a new SelectAll
func MakeSelectAll(db database.Connection, collectRows database.CollectRows[DAO]) SelectAll {
	const query string = `
		SELECT search_criteria_id, tweets_year, tweets_month, total_tweets 
		FROM search_criteria_executions_summary;
	`

	return func(ctx context.Context) ([]DAO, error) {
		rows, err := db.Query(ctx, query)
		if err != nil {
			log.Error(ctx, err.Error())
			return nil, FailedToRetrieveExecutionsSummary
		}

		searchCriteriaExecutionsSummary, err := collectRows(rows)
		if err != nil {
			log.Error(ctx, err.Error())
			return nil, FailedToExecuteCollectRowsInSelectAll
		}

		return searchCriteriaExecutionsSummary, nil
	}
}
