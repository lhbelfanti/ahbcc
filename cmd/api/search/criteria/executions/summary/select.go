package summary

import (
	"context"
	"errors"

	"github.com/jackc/pgx/v5"

	"ahbcc/internal/database"
	"ahbcc/internal/log"
)

// SelectIDBySearchCriteriaIDYearAndMonth returns the id of a tweets counts row, seeking by its search_criteria_id, tweets_year, and tweets_month
type SelectIDBySearchCriteriaIDYearAndMonth func(ctx context.Context, searchCriteriaID, year, month int) (int, error)

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
