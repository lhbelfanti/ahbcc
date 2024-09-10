package criteria

import "errors"

const (
	InvalidURLParameter     string = "Invalid url parameter"
	FailedToEnqueueCriteria string = "Failed to execute enqueue criteria"
)

var (
	FailedToRetrieveCriteriaData      = errors.New("failed to retrieve criteria data")
	FailedToExecuteSelectCriteriaByID = errors.New("failed to execute select criteria by id")
	FailedToExecuteEnqueueCriteria    = errors.New("failed to execute enqueue criteria")
)
