package criteria

import (
	"context"

	"ahbcc/internal/log"
)

// Init retrieves all the criteria in a 'PENDING' or 'IN PROGRESS' state and executes an enqueue of each one
type Init func(ctx context.Context) error

// MakeInit creates a new Init
func MakeInit(selectExecutionsByStatuses SelectExecutionsByStatuses, enqueue Enqueue) Init {
	return func(ctx context.Context) error {
		executionsDAO, err := selectExecutionsByStatuses(ctx, []string{PendingStatus, InProgressStatus})
		if err != nil {
			log.Error(ctx, err.Error())
			return FailedToExecuteSelectExecutionsByStatuses
		}

		for _, execution := range executionsDAO {
			err = enqueue(ctx, execution.SearchCriteriaID, true)
			if err != nil {
				log.Error(ctx, err.Error())
				return FailedToExecuteEnqueueCriteria
			}
		}

		return nil
	}
}
