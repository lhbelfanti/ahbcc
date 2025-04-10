package summary

import (
	"context"
	"errors"

	"github.com/jackc/pgx/v5"

	"ahbcc/internal/database"
	"ahbcc/internal/log"
)

type (
	// SelectIDBySearchCriteriaIDYearAndMonth returns the id of a tweets counts row, seeking by its search_criteria_id, tweets_year, and tweets_month
	SelectIDBySearchCriteriaIDYearAndMonth func(ctx context.Context, searchCriteriaID, year, month int) (int, error)

	// SelectMonthlyTweetsCountsByYearByCriteriaID returns the count of all the tweets (using the `tweets` table) for each year and month, seeking by search criteria ID
	SelectMonthlyTweetsCountsByYearByCriteriaID func(ctx context.Context, criteriaID int) ([]DAO, error)
)

// MakeSelectIDBySearchCriteriaIDYearAndMonth creates a new SelectIDBySearchCriteriaIDYearAndMonth
func MakeSelectIDBySearchCriteriaIDYearAndMonth(db database.Connection) SelectIDBySearchCriteriaIDYearAndMonth {
	const query string = `
		SELECT id
		FROM search_criteria_executions_summary
		WHERE search_criteria_id = $1 AND year = $2 AND month = $3;
	`

	return func(ctx context.Context, searchCriteriaID, year, month int) (int, error) {
		var searchCriteriaExecutionSummaryID int
		err := db.QueryRow(ctx, query, searchCriteriaID, year, month).Scan(&searchCriteriaExecutionSummaryID)
		if errors.Is(err, pgx.ErrNoRows) {
			log.Error(ctx, err.Error())
			return 0, NoExecutionSummaryFoundForTheGivenCriteria
		} else if err != nil {
			log.Error(ctx, err.Error())
			return 0, FailedToExecuteQueryToRetrieveExecutionsSummary
		}

		return searchCriteriaExecutionSummaryID, nil
	}
}

// MakeSelectMonthlyTweetsCountsByYearByCriteriaID creates a new SelectMonthlyTweetCountByYearByCriteriaID
func MakeSelectMonthlyTweetsCountsByYearByCriteriaID(db database.Connection, collectRows database.CollectRows[DAO]) SelectMonthlyTweetsCountsByYearByCriteriaID {
	const query string = `
		SELECT 
			EXTRACT(YEAR FROM posted_at) AS posted_at_year,
			LPAD(EXTRACT(MONTH FROM posted_at)::text, 2, '0') AS posted_at_month,
			COUNT(*) AS quantity
		FROM 
			tweets
		WHERE 
			search_criteria_id = $1
		GROUP BY 
			posted_at_year,
			posted_at_month
		ORDER BY 
			posted_at_year DESC,
			posted_at_month DESC;
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
