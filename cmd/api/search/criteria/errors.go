package criteria

import "errors"

const (
	InvalidURLParameter     string = "Invalid url parameter"
	FailedToEnqueueCriteria string = "Failed to execute enqueue criteria"
)

var (
	FailedToRetrieveCriteriaData = errors.New("failed to retrieve criteria data")

	FailedToExecuteSelectCriteriaByID = errors.New("failed to execute select criteria by id")
	FailedToExecuteEnqueueCriteria    = errors.New("failed to execute enqueue criteria")

	FailedToInsertSearchCriteriaExecution = errors.New("failed to insert search criteria execution")
	FailedToUpdateSearchCriteriaExecution = errors.New("failed to update search criteria execution")

	FailedToInsertSearchCriteriaExecutionDay = errors.New("failed to insert search criteria execution day")
)
