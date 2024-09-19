package criteria

import (
	"context"
	"errors"

	"ahbcc/internal/log"
	"ahbcc/internal/scrapper"
)

// Enqueue retrieves the criteria by ID from the database and enqueues its information
type Enqueue func(ctx context.Context, criteriaID int, forced bool) error

// MakeEnqueue creates a new Enqueue
func MakeEnqueue(selectCriteriaByID SelectByID, selectLastDayExecutedByCriteria SelectLastDayExecutedByCriteriaID, selectExecutionsByStatuses SelectExecutionsByStatuses, enqueueCriteria scrapper.EnqueueCriteria) Enqueue {
	return func(ctx context.Context, criteriaID int, forced bool) error {
		criteriaDAO, err := selectCriteriaByID(ctx, criteriaID)
		if err != nil {
			log.Error(ctx, err.Error())
			return FailedToExecuteSelectCriteriaByID
		}

		if forced {
			lastDayExecutedDate, err := selectLastDayExecutedByCriteria(ctx, criteriaID)
			if err != nil {
				if !errors.Is(err, NoExecutionDaysFoundForTheGivenCriteriaID) {
					log.Error(ctx, err.Error())
					return FailedToExecuteSelectLastDayExecutedByCriteriaID
				}
			} else {
				criteriaDAO.Since = lastDayExecutedDate
			}
		} else {
			executionsDAO, err := selectExecutionsByStatuses(ctx, []string{PendingStatus, InProgressStatus})
			if err != nil {
				log.Error(ctx, err.Error())
				return FailedToExecuteSelectExecutionsByStatuses
			}

			for _, execution := range executionsDAO {
				if execution.SearchCriteriaID == criteriaID {
					return AnExecutionOfThisCriteriaIDIsAlreadyEnqueued
				}
			}
		}

		err = enqueueCriteria(ctx, criteriaDAO.toCriteriaDTO())
		if err != nil {
			log.Error(ctx, err.Error())
			return FailedToExecuteEnqueueCriteria
		}

		return nil
	}
}
