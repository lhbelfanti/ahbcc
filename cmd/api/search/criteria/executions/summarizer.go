package executions

import (
	"context"

	"github.com/lhbelfanti/corpus-creator/cmd/api/search/criteria/executions/summary"
	"github.com/lhbelfanti/corpus-creator/internal/database"
	"github.com/lhbelfanti/corpus-creator/internal/log"
)

// Summarize creates a summary of the search criteria executions. The summary is saved for each month of each year
// from where the tweets were retrieved
type Summarize func(ctx context.Context) error

// MakeSummarize creates a new Summarize
func MakeSummarize(db database.Connection, selectExecutionsByStatuses SelectExecutionsByStatuses, deleteAllExecutionsSummary summary.DeleteAll, selectMonthlyTweetsCountsByYear summary.SelectMonthlyTweetsCountsByYearByCriteriaID, insertExecutionSummary summary.Insert) Summarize {
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

		err = deleteAllExecutionsSummary(ctx)
		if err != nil {
			log.Error(ctx, err.Error())
			return FailedToClearOldSummary
		}

		for _, searchCriteriaExecution := range searchCriteriaExecutions {
			searchCriteriaExecutionsSummary, err := selectMonthlyTweetsCountsByYear(ctx, searchCriteriaExecution.SearchCriteriaID)
			if err != nil {
				log.Error(ctx, err.Error())
				return FailedToExecuteSelectMonthlyTweetsCountsByYear
			}

			for _, executionSummary := range searchCriteriaExecutionsSummary {
				_, err = insertExecutionSummary(tx, ctx, executionSummary)
				if err != nil {
					log.Error(ctx, err.Error())
					return FailedToExecuteInsertExecutionSummary
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
