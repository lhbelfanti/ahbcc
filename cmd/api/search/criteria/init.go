package criteria

import (
	"context"
	"errors"

	"github.com/lhbelfanti/corpus-creator/cmd/api/search/criteria/executions"
	"github.com/lhbelfanti/corpus-creator/internal/log"
)

// Init retrieves all the criteria in a 'PENDING' or 'IN PROGRESS' state and executes an enqueue of each one
type Init func(ctx context.Context) error

// MakeInit creates a new Init
func MakeInit(selectExecutionsByStatuses executions.SelectExecutionsByStatuses, resume Resume) Init {
	return func(ctx context.Context) error {
		executionsDAO, err := selectExecutionsByStatuses(ctx, []string{executions.PendingStatus, executions.InProgressStatus})
		if err != nil {
			log.Error(ctx, err.Error())
			return FailedToExecuteSelectExecutionsByStatuses
		}

		for _, execution := range executionsDAO {
			err = resume(ctx, execution.SearchCriteriaID)
			if err != nil {
				if !errors.Is(err, FailedToRetrieveSearchCriteriaExecutionID) {
					log.Error(ctx, err.Error())
					return FailedToExecuteEnqueueCriteria
				} else {
					// If the search criteria does not have an active execution, there is nothing to enqueue
					continue
				}
			}
		}

		return nil
	}
}
