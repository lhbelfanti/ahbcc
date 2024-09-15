package criteria

import "errors"

const (
	InvalidURLParameter                        string = "Invalid url parameter"
	InvalidQueryParameterFormat                string = "Invalid query parameter format"
	FailedToEnqueueCriteria                    string = "Failed to execute enqueue criteria"
	ExecutionWithSameCriteriaIDAlreadyEnqueued string = "An execution with the same criteria id is already enqueued"
)

var (
	FailedToRetrieveCriteriaData = errors.New("failed to retrieve criteria data")

	FailedToExecuteSelectCriteriaByID            = errors.New("failed to execute select criteria by id")
	FailedToExecuteSelectExecutionsByStatuses    = errors.New("failed to execute select executions by statuses")
	AnExecutionOfThisCriteriaIDIsAlreadyEnqueued = errors.New("an execution of this criteria is already enqueued")
	FailedToExecuteEnqueueCriteria               = errors.New("failed to execute enqueue criteria")

	FailedToInsertSearchCriteriaExecution                    = errors.New("failed to insert search criteria execution")
	FailedToUpdateSearchCriteriaExecution                    = errors.New("failed to update search criteria execution")
	FailedToExecuteSelectSearchCriteriaExecutionByState      = errors.New("failed to execute select search criteria execution by state")
	FailedToExecuteSelectCollectRowsInSelectExecutionByState = errors.New("failed to execute select collect rows in select criteria execution by state")

	FailedToInsertSearchCriteriaExecutionDay = errors.New("failed to insert search criteria execution day")
)
