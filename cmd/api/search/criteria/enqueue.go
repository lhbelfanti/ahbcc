package criteria

import (
	"context"
	"errors"

	"ahbcc/internal/log"
	"ahbcc/internal/scrapper"
)

type (
	// Enqueue retrieves the criteria by ID from the database and enqueues its information
	Enqueue func(ctx context.Context, criteriaID int, forced bool) error

	// Resume retrieves the criteria by ID from the database, searches its last execution day and
	// enqueues data starting from that day
	Resume func(ctx context.Context, criteriaID int) error
)

// MakeEnqueue creates a new Enqueue
func MakeEnqueue(selectCriteriaByID SelectByID, selectExecutionsByStatuses SelectExecutionsByStatuses, enqueueCriteria scrapper.EnqueueCriteria) Enqueue {
	return func(ctx context.Context, criteriaID int, forced bool) error {
		criteriaDAO, err := selectCriteriaByID(ctx, criteriaID)
		if err != nil {
			log.Error(ctx, err.Error())
			return FailedToExecuteSelectCriteriaByID
		}

		if !forced {
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

// MakeResume creates a new Resume
func MakeResume(selectCriteriaByID SelectByID, selectLastDayExecutedByCriteria SelectLastDayExecutedByCriteriaID, enqueueCriteria scrapper.EnqueueCriteria) Resume {
	return func(ctx context.Context, criteriaID int) error {
		criteriaDAO, err := selectCriteriaByID(ctx, criteriaID)
		if err != nil {
			log.Error(ctx, err.Error())
			return FailedToExecuteSelectCriteriaByID
		}

		lastDayExecutedDate, err := selectLastDayExecutedByCriteria(ctx, criteriaID)
		if err != nil {
			// if err == NoExecutionDaysFoundForTheGivenCriteriaID the criteria hasn’t started yet, but it was enqueued once before
			if !errors.Is(err, NoExecutionDaysFoundForTheGivenCriteriaID) {
				log.Error(ctx, err.Error())
				return FailedToExecuteSelectLastDayExecutedByCriteriaID
			}
		} else { // The criteria has been executed once before and is needed to start from the last day it was executed
			criteriaDAO.Since = lastDayExecutedDate
		}

		err = enqueueCriteria(ctx, criteriaDAO.toCriteriaDTO())
		if err != nil {
			log.Error(ctx, err.Error())
			return FailedToExecuteEnqueueCriteria
		}

		return nil
	}
}
