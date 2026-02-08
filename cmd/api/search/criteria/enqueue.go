package criteria

import (
	"context"
	"errors"
	"time"

	"github.com/lhbelfanti/corpus-creator/cmd/api/search/criteria/executions"
	"github.com/lhbelfanti/corpus-creator/internal/log"
	"github.com/lhbelfanti/corpus-creator/internal/scrapper"
)

type (
	// Enqueue retrieves the criteria by ID from the database and enqueues its information
	Enqueue func(ctx context.Context, criteriaID int, forced bool) error

	// Resume retrieves the criteria by ID from the database, searches its last execution day and
	// enqueues data starting from that day
	Resume func(ctx context.Context, criteriaID int) error
)

// MakeEnqueue creates a new Enqueue
func MakeEnqueue(selectCriteriaByID SelectByID, selectExecutionsByStatuses executions.SelectExecutionsByStatuses, insertExecution executions.InsertExecution, enqueueCriteria scrapper.EnqueueCriteria) Enqueue {
	return func(ctx context.Context, criteriaID int, forced bool) error {
		criteriaDAO, err := selectCriteriaByID(ctx, criteriaID)
		if err != nil {
			log.Error(ctx, err.Error())
			return FailedToExecuteSelectCriteriaByID
		}

		if !forced {
			executionsDAO, err := selectExecutionsByStatuses(ctx, []string{executions.PendingStatus, executions.InProgressStatus})
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

		executionID, err := insertExecution(ctx, criteriaID, forced)
		if err != nil {
			log.Error(ctx, err.Error())
			return FailedToInsertSearchCriteriaExecution
		}

		err = enqueueCriteria(ctx, criteriaDAO.toCriteriaDTO(), executionID)
		if err != nil {
			// TODO: Implement transaction: if the criteria was inserted but the enqueue failed, the criteria should be removed from the DB
			// to avoid `AnExecutionOfThisCriteriaIDIsAlreadyEnqueued` if it is needed to enqueue the same criteria again

			log.Error(ctx, err.Error())
			return FailedToExecuteEnqueueCriteria
		}

		return nil
	}
}

// MakeResume creates a new Resume
func MakeResume(selectCriteriaByID SelectByID, selectLastDayExecutedByCriteria executions.SelectLastDayExecutedByCriteriaID, selectExecutionsByStatuses executions.SelectExecutionsByStatuses, enqueueCriteria scrapper.EnqueueCriteria) Resume {
	return func(ctx context.Context, criteriaID int) error {
		criteriaDAO, err := selectCriteriaByID(ctx, criteriaID)
		if err != nil {
			log.Error(ctx, err.Error())
			return FailedToExecuteSelectCriteriaByID
		}

		searchCriteriaExecutionID := -1
		lastExecutionDayExecuted, err := selectLastDayExecutedByCriteria(ctx, criteriaID)
		if err != nil {
			if !errors.Is(err, NoExecutionDaysFoundForTheGivenCriteriaID) {
				log.Error(ctx, err.Error())
				return FailedToExecuteSelectLastDayExecutedByCriteriaID
			}

			// The criterion hasn't started yet, but it was enqueued once before (it is in a PENDING state for example)
			executionsDAO, err := selectExecutionsByStatuses(ctx, []string{executions.PendingStatus, executions.InProgressStatus})
			if err != nil {
				log.Error(ctx, err.Error())
				return FailedToExecuteSelectExecutionsByStatuses
			}

			for _, execution := range executionsDAO {
				if execution.SearchCriteriaID == criteriaID {
					searchCriteriaExecutionID = execution.ID
					break
				}
			}
		} else {
			// The search criteria has been executed once before and is needed to start from the next day of the last day it was executed
			criteriaDAO.Since = lastExecutionDayExecuted.ExecutionDate.Add(24 * time.Hour)
			searchCriteriaExecutionID = lastExecutionDayExecuted.SearchCriteriaExecutionID
		}

		if searchCriteriaExecutionID == -1 {
			return FailedToRetrieveSearchCriteriaExecutionID
		}

		err = enqueueCriteria(ctx, criteriaDAO.toCriteriaDTO(), searchCriteriaExecutionID)
		if err != nil {
			log.Error(ctx, err.Error())
			return FailedToExecuteEnqueueCriteria
		}

		return nil
	}
}
