package executions

import (
	"context"

	"ahbcc/cmd/api/search/criteria/executions/summary"
	"ahbcc/internal/database"
	"ahbcc/internal/log"
)

// SummarizeExecutions creates a summary of the search criteria executions. The summary is saved for each month of each year
// from where the tweets were retrieved
type SummarizeExecutions func(ctx context.Context) error

// MakeSummarizeExecutions creates a new SummarizeExecutions
func MakeSummarizeExecutions(db database.Connection, selectExecutionsByStatuses SelectExecutionsByStatuses, selectMonthlyTweetsCountsByYear summary.SelectMonthlyTweetsCountsByYearByCriteriaID, upsertExecutionSummary summary.Upsert) SummarizeExecutions {
	return func(ctx context.Context) error {
		searchCriteriaExecutions, err := selectExecutionsByStatuses(ctx, []string{"DONE"})
		if err != nil {
			log.Error(ctx, err.Error())
			return FailedToExecuteSelectSearchCriteriaExecutionByState
		}

		tx, err := db.Begin(ctx)
		if err != nil {
			log.Error(ctx, err.Error())
			return FailedToBeginTransaction
		}

		defer tx.Rollback(ctx)

		for _, searchCriteriaExecution := range searchCriteriaExecutions {
			searchCriteriaExecutionsSummary, err := selectMonthlyTweetsCountsByYear(ctx, searchCriteriaExecution.SearchCriteriaID)
			if err != nil {
				log.Error(ctx, err.Error())
				return FailedToExecuteSelectMonthlyTweetsCountsByYear
			}

			for _, executionSummary := range searchCriteriaExecutionsSummary {
				err = upsertExecutionSummary(ctx, tx, executionSummary)
				if err != nil {
					log.Error(ctx, err.Error())
					return FailedToExecuteUpsertExecutionSummary
				}
			}
		}

		err = tx.Commit(ctx)
		if err != nil {
			log.Error(ctx, err.Error())
			return FailedToCommitTransaction
		}

		return nil
	}
}
