package tweets

import (
	"context"

	"ahbcc/cmd/api/tweets/counts"
	"ahbcc/internal/database"
	"ahbcc/internal/log"
)

// SelectMonthlyTweetsCountsByYearByCriteriaID returns the count of all the tweets (using the `tweets` table) for each year and month, seeking by search criteria ID
type SelectMonthlyTweetsCountsByYearByCriteriaID func(ctx context.Context, criteriaID int) ([]counts.DAO, error)

// MakeSelectMonthlyTweetsCountsByYearByCriteriaID creates a new SelectMonthlyTweetCountByYearByCriteriaID
func MakeSelectMonthlyTweetsCountsByYearByCriteriaID(db database.Connection, collectRows database.CollectRows[counts.DAO]) SelectMonthlyTweetsCountsByYearByCriteriaID {
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

	return func(ctx context.Context, criteriaID int) ([]counts.DAO, error) {
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
