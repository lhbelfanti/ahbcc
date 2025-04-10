package summary

import (
	"context"
	"errors"

	"github.com/jackc/pgx/v5"

	"ahbcc/internal/log"
)

// Upsert seeks for the ID of the summary from the table search_criteria_executions_summary.
// If it is found it means the summary was already added, and it updates the total_tweets value.
// If it is not found, a new row in the table is inserted.
type Upsert func(ctx context.Context, tx pgx.Tx, executionSummary DAO) error

// MakeUpsert creates a new Upsert
func MakeUpsert(selectIDBySearchCriteriaIDYearAndMonth SelectIDBySearchCriteriaIDYearAndMonth, insertExecutionSummary Insert, updateSummaryTotalTweets UpdateTotalTweets) Upsert {
	return func(ctx context.Context, tx pgx.Tx, executionSummary DAO) error {
		executionSummaryID, err := selectIDBySearchCriteriaIDYearAndMonth(ctx, executionSummary.SearchCriteriaID, executionSummary.Year, executionSummary.Month)

		if err != nil && !errors.Is(err, NoExecutionSummaryFoundForTheGivenCriteria) {
			log.Error(ctx, err.Error())
			return FailedToRetrieveExecutionSummaryID
		}

		if errors.Is(err, NoExecutionSummaryFoundForTheGivenCriteria) {
			_, err = insertExecutionSummary(tx, ctx, executionSummary)
			if err != nil {
				log.Error(ctx, err.Error())
				return FailedToInsertExecutionSummary
			}
		} else {
			err = updateSummaryTotalTweets(tx, ctx, executionSummaryID, executionSummary.Total)
			if err != nil {
				log.Error(ctx, err.Error())
				return FailedToUpdateTotalTweets
			}
		}

		return nil
	}
}
